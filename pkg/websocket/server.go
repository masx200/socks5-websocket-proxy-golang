package websocket

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/upstream"
)

// WebSocketServer WebSocket服务端实现
type WebSocketServer struct {
	config     interfaces.ServerConfig
	httpServer *http.Server
	upgrader   *websocket.Upgrader
	shutdown   chan struct{}
	wg         sync.WaitGroup
	authUsers  map[string]string
	selector   *upstream.UpstreamSelector
}

// NewWebSocketServer 创建新的WebSocket服务端
func NewWebSocketServer(config interfaces.ServerConfig) *WebSocketServer {
	var selector *upstream.UpstreamSelector
	if config.EnableUpstream && len(config.UpstreamConfig) > 0 {
		selector = upstream.NewUpstreamSelector(&config.UpstreamConfig[0])
	}
	
	return &WebSocketServer{
		config:   config,
		shutdown: make(chan struct{}),
		upgrader: &websocket.Upgrader{
			HandshakeTimeout: config.Timeout,
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			CheckOrigin: func(r *http.Request) bool {
				// 允许所有来源，生产环境中应该限制
				return true
			},
		},
		authUsers: config.AuthUsers,
		selector:  selector,
	}
}

// Listen 开始监听连接
func (s *WebSocketServer) Listen() error {
	// 创建HTTP路由
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleWebSocketConnection)

	// 创建HTTP服务器
	s.httpServer = &http.Server{
		Addr:         s.config.ListenAddr,
		Handler:      mux,
		ReadTimeout:  s.config.Timeout,
		WriteTimeout: s.config.Timeout,
	}

	fmt.Printf("WebSocket server listening on %s\n", s.config.ListenAddr)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	return nil
}

// handleWebSocketConnection 处理WebSocket连接
func (s *WebSocketServer) handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	// 从HTTP Headers中解析认证信息
	username, password, targetHost, targetPort, err := s.parseAuthInfo(r)
	if err != nil {
		fmt.Printf("auth info parse error: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 验证用户名密码
	if !s.Authenticate(username, password) {
		fmt.Printf("authentication failed for user: %s\n", username)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 升级HTTP连接为WebSocket连接
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade error: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("WebSocket connection established for target %s:%d\n", targetHost, targetPort)

	// 处理WebSocket连接
	s.wg.Add(1)
	defer s.wg.Done()

	// 创建包装的net.Conn
	wsConn := &websocketNetConn{conn: conn, targetHost: targetHost, targetPort: targetPort}

	// 使用统一的连接处理逻辑
	if err := s.HandleConnection(wsConn); err != nil {
		fmt.Printf("connection handling error: %v\n", err)
	}
}

// HandleConnection 处理客户端连接（实现ProxyServer接口）
func (s *WebSocketServer) HandleConnection(conn net.Conn) error {
	// 从连接中获取目标信息
	wsConn, ok := conn.(*websocketNetConn)
	if !ok {
		return errors.New("invalid connection type")
	}
	
	targetHost := wsConn.targetHost
	targetPort := wsConn.targetPort

	// 选择上游连接
	upstreamConn, err := s.SelectUpstreamConnection(targetHost, targetPort)
	if err != nil {
		return fmt.Errorf("failed to select upstream connection: %w", err)
	}
	defer upstreamConn.Close()

	// 开始数据转发
	return s.forwardData(conn, upstreamConn)
}

// Authenticate 验证用户名密码
func (s *WebSocketServer) Authenticate(username, password string) bool {
	if s.authUsers == nil {
		return true // 如果没有配置用户，则允许所有连接
	}

	storedPassword, exists := s.authUsers[username]
	if !exists {
		return false
	}

	return storedPassword == password
}

// SelectUpstreamConnection 选择上游连接
func (s *WebSocketServer) SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error) {
	if s.selector != nil {
		return s.selector.SelectConnection(targetHost, targetPort)
	}
	
	// 默认直连
	return net.DialTimeout("tcp", fmt.Sprintf("%s:%d", targetHost, targetPort), s.config.Timeout)
}

// Shutdown 优雅关闭服务端
func (s *WebSocketServer) Shutdown() error {
	close(s.shutdown)
	if s.httpServer != nil {
		s.httpServer.Close()
	}
	s.wg.Wait()
	return nil
}

// parseAuthInfo 从HTTP请求中解析认证信息
func (s *WebSocketServer) parseAuthInfo(r *http.Request) (username, password, targetHost string, targetPort int, err error) {
	// 优先从HTTP Headers中解析
	if authHeader := r.Header.Get("X-Proxy-Auth"); authHeader != "" {
		return s.parseAuthHeader(authHeader)
	}

	// 从URL查询参数中解析
	if username = r.URL.Query().Get("username"); username != "" {
		password = r.URL.Query().Get("password")
		targetHost = r.URL.Query().Get("host")
		portStr := r.URL.Query().Get("port")
		if portStr != "" {
			_, err := fmt.Sscanf(portStr, "%d", &targetPort)
			if err != nil {
				return "", "", "", 0, fmt.Errorf("invalid port: %w", err)
			}
		}
		return username, password, targetHost, targetPort, nil
	}

	// 如果没有找到认证信息，返回默认值
	return "", "", "", 0, errors.New("no authentication information found")
}

// parseAuthHeader 解析HTTP Header中的认证信息
func (s *WebSocketServer) parseAuthHeader(authHeader string) (username, password, targetHost string, targetPort int, err error) {
	// Base64解码
	decoded, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return "", "", "", 0, fmt.Errorf("failed to decode auth header: %w", err)
	}

	// 解析查询参数
	values, err := url.ParseQuery(string(decoded))
	if err != nil {
		return "", "", "", 0, fmt.Errorf("failed to parse auth info: %w", err)
	}

	username = values.Get("username")
	password = values.Get("password")
	targetHost = values.Get("host")

	if portStr := values.Get("port"); portStr != "" {
		_, err := fmt.Sscanf(portStr, "%d", &targetPort)
		if err != nil {
			return "", "", "", 0, fmt.Errorf("invalid port: %w", err)
		}
	}

	return username, password, targetHost, targetPort, nil
}

// forwardData 转发数据
func (s *WebSocketServer) forwardData(clientConn, targetConn net.Conn) error {
	done := make(chan error, 2)

	go func() {
		_, err := io.Copy(targetConn, clientConn)
		done <- err
	}()

	go func() {
		_, err := io.Copy(clientConn, targetConn)
		done <- err
	}()

	return <-done
}

// websocketNetConn 将websocket.Conn包装为net.Conn
type websocketNetConn struct {
	conn       *websocket.Conn
	readBuffer []byte
	targetHost string
	targetPort int
}

func (w *websocketNetConn) Read(b []byte) (n int, err error) {
	if len(w.readBuffer) > 0 {
		n = copy(b, w.readBuffer)
		w.readBuffer = w.readBuffer[n:]
		return n, nil
	}

	_, message, err := w.conn.ReadMessage()
	if err != nil {
		return 0, err
	}

	if len(message) <= len(b) {
		copy(b, message)
		return len(message), nil
	}

	copy(b, message)
	w.readBuffer = message[len(b):]
	return len(b), nil
}

func (w *websocketNetConn) Write(b []byte) (n int, err error) {
	err = w.conn.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (w *websocketNetConn) Close() error {
	return w.conn.Close()
}

func (w *websocketNetConn) LocalAddr() net.Addr {
	return w.conn.LocalAddr()
}

func (w *websocketNetConn) RemoteAddr() net.Addr {
	return w.conn.RemoteAddr()
}

func (w *websocketNetConn) SetDeadline(t time.Time) error {
	return w.conn.SetReadDeadline(t)
}

func (w *websocketNetConn) SetReadDeadline(t time.Time) error {
	return w.conn.SetReadDeadline(t)
}

func (w *websocketNetConn) SetWriteDeadline(t time.Time) error {
	return w.conn.SetWriteDeadline(t)
}

// 注册WebSocket服务端创建函数
func init() {
	interfaces.RegisterServer("websocket", func(config interfaces.ServerConfig) (interfaces.ProxyServer, error) {
		return NewWebSocketServer(config), nil
	})
}
