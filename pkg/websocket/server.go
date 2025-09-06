package websocket

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
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
	selector   interface{}
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
	clientAddr := r.RemoteAddr
	fmt.Printf("WebSocket connection attempt from %s\n", clientAddr)

	// 从HTTP Headers中解析认证信息
	username, password, targetHost, targetPort, err := s.parseAuthInfo(r)
	if err != nil {
		fmt.Printf("WebSocket auth info parse error from %s: %v\n", clientAddr, err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 验证用户名密码
	if !s.Authenticate(username, password) {
		fmt.Printf("WebSocket authentication failed for user %s from %s\n", username, clientAddr)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Printf("WebSocket authentication successful for user %s from %s\n", username, clientAddr)

	// 升级HTTP连接为WebSocket连接
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade failed for client %s: %v\n", clientAddr, err)
		return
	}
	defer conn.Close()

	fmt.Printf("WebSocket connection established successfully for target %s:%d from %s\n", targetHost, targetPort, clientAddr)

	// 处理WebSocket连接
	s.wg.Add(1)
	defer s.wg.Done()

	// 创建包装的net.Conn
	wsConn := &websocketNetConn{conn: conn, targetHost: targetHost, targetPort: targetPort}

	// 使用统一的连接处理逻辑
	if err := s.HandleConnection(wsConn); err != nil {
		fmt.Printf("WebSocket connection handling failed for target %s:%d from %s: %v\n", targetHost, targetPort, clientAddr, err)
	} else {
		fmt.Printf("WebSocket connection completed successfully for target %s:%d from %s\n", targetHost, targetPort, clientAddr)
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
		// 尝试类型断言
		if selector, ok := s.selector.(*upstream.UpstreamSelector); ok {
			return selector.SelectConnection(targetHost, targetPort)
		} else if selector, ok := s.selector.(*upstream.DynamicUpstreamSelector); ok {
			return selector.SelectConnection(targetHost, targetPort)
		}
	}

	// 默认直连
	return net.DialTimeout("tcp", net.JoinHostPort(targetHost, fmt.Sprint(targetPort)), s.config.Timeout)
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

// ReloadConfig 重新加载配置
func (s *WebSocketServer) ReloadConfig(newConfig interfaces.ServerConfig) error {
	fmt.Printf("[WEBSOCKET-SERVER] Reloading configuration...\n")
	
	// 更新配置
	s.config = newConfig
	s.authUsers = newConfig.AuthUsers
	
	// 更新上游选择器
	if newConfig.EnableUpstream {
		selector := upstream.NewDynamicUpstreamSelector(newConfig.UpstreamConfig, upstream.StrategyRoundRobin)
		s.selector = selector
		fmt.Printf("[WEBSOCKET-SERVER] Upstream selector updated with %d configurations\n", len(newConfig.UpstreamConfig))
	} else {
		s.selector = nil
		fmt.Printf("[WEBSOCKET-SERVER] Upstream selector disabled\n")
	}
	
	fmt.Printf("[WEBSOCKET-SERVER] Configuration reloaded successfully\n")
	fmt.Printf("[WEBSOCKET-SERVER] Authentication users: %d\n", len(newConfig.AuthUsers))
	fmt.Printf("[WEBSOCKET-SERVER] Upstream enabled: %t\n", newConfig.EnableUpstream)
	
	return nil
}

// parseAuthInfo 从HTTP请求中解析认证信息
func (s *WebSocketServer) parseAuthInfo(r *http.Request) (username, password, targetHost string, targetPort int, err error) {
	// 直接从HTTP Headers中获取用户名和密码
	username = r.Header.Get("X-Proxy-Username")
	password = r.Header.Get("X-Proxy-Password")
	targetHost = r.URL.Query().Get("host")

	portStr := r.URL.Query().Get("port")
	if portStr != "" {
		_, err := fmt.Sscanf(portStr, "%d", &targetPort)
		if err != nil {
			return "", "", "", 0, fmt.Errorf("invalid port: %w", err)
		}
	}

	// 允许没有认证信息的情况，返回空字符串而不是错误
	// 实际的认证验证会在Authenticate方法中进行
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
