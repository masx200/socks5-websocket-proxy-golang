package websocket

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/masx200/http-proxy-go-server/resolver"
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
	resolver   resolver.NameResolver
}

// NewWebSocketServer 创建新的WebSocket服务端
func NewWebSocketServer(config interfaces.ServerConfig) *WebSocketServer {
	return NewWebSocketServerWithResolver(config, nil)
}

// NewWebSocketServerWithResolver 创建带有解析器的WebSocket服务端
func NewWebSocketServerWithResolver(config interfaces.ServerConfig, resolver resolver.NameResolver) *WebSocketServer {
	var selector interface{}
	if config.EnableUpstream && len(config.UpstreamConfig) > 0 {
		if resolver != nil {
			selector = upstream.NewDynamicUpstreamSelectorWithResolver(config.UpstreamConfig, upstream.StrategyRoundRobin, resolver)
		} else {
			selector = upstream.NewUpstreamSelector(&config.UpstreamConfig[0])
		}
	} else {
		selector = nil
	}

	return &WebSocketServer{
		config:   config,
		shutdown: make(chan struct{}),
		upgrader: &websocket.Upgrader{
			EnableCompression: true,
			HandshakeTimeout:  config.Timeout,
			ReadBufferSize:    1024,
			WriteBufferSize:   1024,
			CheckOrigin: func(r *http.Request) bool {
				// 允许所有来源，生产环境中应该限制
				return true
			},
		},
		authUsers: config.AuthUsers,
		selector:  selector,
		resolver:  resolver,
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

	log.Printf("[WEBSOCKET-SERVER] Server started successfully, listening on %s\n", s.config.ListenAddr)
	log.Printf("[WEBSOCKET-SERVER] Authentication enabled: %t (%d users configured)\n",
		len(s.authUsers) > 0, len(s.authUsers))
	log.Printf("[WEBSOCKET-SERVER] Upstream selector enabled: %t\n", s.config.EnableUpstream)
	log.Printf("[WEBSOCKET-SERVER] Read timeout: %v, Write timeout: %v\n", s.config.Timeout, s.config.Timeout)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[WEBSOCKET-SERVER] HTTP server error: %v\n", err)
		}
	}()

	return nil
}

// handleWebSocketConnection 处理WebSocket连接
func (s *WebSocketServer) handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {

	log.Println("url", r.URL.String())

	log.Println("headers", r.Header)
	clientAddr := r.RemoteAddr
	startTime := time.Now()
	log.Printf("[WEBSOCKET-CONN] New connection attempt from %s at %s\n", clientAddr, startTime.Format("2006-01-02 15:04:05"))

	// 记录认证配置状态
	if len(s.authUsers) > 0 {
		log.Printf("[WEBSOCKET-CONN] Authentication required for client %s (%d users configured)\n", clientAddr, len(s.authUsers))
	} else {
		log.Printf("[WEBSOCKET-CONN] No authentication required for client %s\n", clientAddr)
	}

	// 从HTTP Headers中解析认证信息
	username, password, targetHost, targetPort, err := s.parseAuthInfo(r)
	if err != nil {
		log.Printf("[WEBSOCKET-CONN] Auth info parse error from %s: %v\n", clientAddr, err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 验证用户名密码
	if !s.Authenticate(username, password) {
		log.Printf("[WEBSOCKET-AUTH] Authentication failed for user '%s' from %s\n", username, clientAddr)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("[WEBSOCKET-AUTH] Authentication successful for user '%s' from %s\n", username, clientAddr)

	// 升级HTTP连接为WebSocket连接
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[WEBSOCKET-CONN] WebSocket upgrade failed for client %s: %v\n", clientAddr, err)
		return
	}
	defer conn.Close()

	log.Printf("[WEBSOCKET-CONN] WebSocket connection established successfully for target %s:%d from %s\n", targetHost, targetPort, clientAddr)

	// 处理WebSocket连接
	s.wg.Add(1)
	defer s.wg.Done()

	// 创建包装的net.Conn
	wsConn := &websocketNetConn{conn: conn, targetHost: targetHost, targetPort: targetPort}

	// 使用统一的连接处理逻辑
	if err := s.HandleConnection(wsConn); err != nil {
		duration := time.Since(startTime)
		log.Printf("[WEBSOCKET-CONN] Connection handling failed for target %s:%d from %s after %v: %v\n", targetHost, targetPort, clientAddr, duration, err)
	} else {
		duration := time.Since(startTime)
		log.Printf("[WEBSOCKET-CONN] Connection completed successfully for target %s:%d from %s, duration: %v\n", targetHost, targetPort, clientAddr, duration)
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
	if s.authUsers == nil || len(s.authUsers) == 0 {
		log.Printf("[WEBSOCKET-AUTH] No authentication configured, allowing access for user '%s'\n", username)
		return true // 如果没有配置用户，则允许所有连接
	}

	storedPassword, exists := s.authUsers[username]
	if !exists {
		log.Printf("[WEBSOCKET-AUTH] Authentication failed: user '%s' not found in configured users\n", username)
		return false
	}

	if storedPassword == password {
		log.Printf("[WEBSOCKET-AUTH] Authentication successful for user '%s'\n", username)
		return true
	} else {
		log.Printf("[WEBSOCKET-AUTH] Authentication failed: invalid password for user '%s'\n", username)
		return false
	}
}

// SelectUpstreamConnection 选择上游连接
func (s *WebSocketServer) SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error) {
	targetAddr := fmt.Sprintf("%s:%d", targetHost, targetPort)

	// 检查selector是否真正有效（不是nil且包含有效值）
	if s.selector != nil && !isNilInterface(s.selector) {
		log.Printf("[WEBSOCKET-UPSTREAM] Using upstream selector for target %s\n", targetAddr)
		var conn net.Conn
		var err error

		// 尝试类型断言
		if selector, ok := s.selector.(*upstream.UpstreamSelector); ok && selector != nil {
			conn, err = selector.SelectConnection(targetHost, targetPort)
		} else if selector, ok := s.selector.(*upstream.DynamicUpstreamSelector); ok && selector != nil {
			conn, err = selector.SelectConnection(targetHost, targetPort)
		} else {
			err = fmt.Errorf("unknown or nil selector type")
			log.Printf("[WEBSOCKET-UPSTREAM] Selector type check failed for target %s: %v\n", targetAddr, err)
			return conn, err
		}

		if err != nil {
			log.Printf("[WEBSOCKET-UPSTREAM] Upstream selector failed for target %s: %v\n", targetAddr, err)
		} else {
			log.Printf("[WEBSOCKET-UPSTREAM] Upstream selector succeeded for target %s\n", targetAddr)
		}
		return conn, err
	}

	// 默认直连
	timeout := s.config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	log.Printf("[WEBSOCKET-UPSTREAM] Using direct connection for target %s (timeout: %v)\n", targetAddr, timeout)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetHost, fmt.Sprint(targetPort)), timeout)
	if err != nil {
		log.Printf("[WEBSOCKET-UPSTREAM] Direct connection failed for target %s: %v\n", targetAddr, err)
	} else {
		log.Printf("[WEBSOCKET-UPSTREAM] Direct connection established for target %s\n", targetAddr)
	}
	return conn, err
}

// Shutdown 优雅关闭服务端
func (s *WebSocketServer) Shutdown() error {
	log.Printf("[WEBSOCKET-SERVER] Initiating graceful shutdown...\n")
	close(s.shutdown)
	if s.httpServer != nil {
		log.Printf("[WEBSOCKET-SERVER] Closing HTTP server...\n")
		s.httpServer.Close()
	}
	log.Printf("[WEBSOCKET-SERVER] Waiting for active connections to close...\n")
	s.wg.Wait()
	log.Printf("[WEBSOCKET-SERVER] Server shutdown completed successfully\n")
	return nil
}

// ReloadConfig 重新加载配置
func (s *WebSocketServer) ReloadConfig(newConfig interfaces.ServerConfig) error {
	log.Printf("[WEBSOCKET-SERVER] Reloading configuration...\n")

	// 更新配置
	s.config = newConfig
	s.authUsers = newConfig.AuthUsers

	// 更新上游选择器
	if newConfig.EnableUpstream {
		var selector interface{}
		if s.resolver != nil {
			selector = upstream.NewDynamicUpstreamSelectorWithResolver(newConfig.UpstreamConfig, upstream.StrategyRoundRobin, s.resolver)
		} else {
			selector = upstream.NewDynamicUpstreamSelector(newConfig.UpstreamConfig, upstream.StrategyRoundRobin)
		}
		s.selector = selector
		log.Printf("[WEBSOCKET-SERVER] Upstream selector updated with %d configurations\n", len(newConfig.UpstreamConfig))
	} else {
		s.selector = nil
		log.Printf("[WEBSOCKET-SERVER] Upstream selector disabled\n")
	}

	log.Printf("[WEBSOCKET-SERVER] Configuration reloaded successfully\n")
	log.Printf("[WEBSOCKET-SERVER] Authentication users: %d\n", len(newConfig.AuthUsers))
	log.Printf("[WEBSOCKET-SERVER] Upstream enabled: %t\n", newConfig.EnableUpstream)
	log.Printf("[WEBSOCKET-SERVER] Listen address: %s\n", newConfig.ListenAddr)
	log.Printf("[WEBSOCKET-SERVER] Timeout: %v\n", newConfig.Timeout)

	return nil
}

// isNilInterface 检查interface{}是否真正为nil（包括内部值为nil的情况）
func isNilInterface(i interface{}) bool {
	if i == nil {
		return true
	}
	// 使用反射检查interface内部的值是否为nil
	v := reflect.ValueOf(i)
	return v.IsNil()
}

// parseAuthInfo 从HTTP请求中解析认证信息
func (s *WebSocketServer) parseAuthInfo(r *http.Request) (username, password, targetHost string, targetPort int, err error) {
	// 直接从HTTP Headers中获取用户名和密码
	username = r.Header.Get("X-Proxy-Username")
	password = r.Header.Get("X-Proxy-Password")

	// 从HTTP Headers中获取目标主机和端口
	targetHost = r.Header.Get("X-Proxy-Target-Host")

	portStr := r.Header.Get("X-Proxy-Target-Port")
	if portStr != "" {
		_, err := fmt.Sscanf(portStr, "%d", &targetPort)
		if err != nil {
			return "", "", "", 0, fmt.Errorf("invalid port: %w", err)
		}
	}

	// 记录解析的认证信息
	log.Printf("[WEBSOCKET-AUTH] Parsed auth info - username: '%s', password: '%s', targetHost: '%s', targetPort: %d\n",
		username, password, targetHost, targetPort)

	// 允许没有认证信息的情况，返回空字符串而不是错误
	// 实际的认证验证会在Authenticate方法中进行
	return username, password, targetHost, targetPort, nil
}

// forwardData 转发数据
func (s *WebSocketServer) forwardData(clientConn, targetConn net.Conn) error {
	log.Printf("[WEBSOCKET-FORWARD] Starting data forwarding between connections\n")
	done := make(chan error, 2)

	go func() {
		bytesWritten, err := io.Copy(targetConn, clientConn)
		log.Printf("[WEBSOCKET-FORWARD] Client -> Target: %d bytes transferred, error: %v\n", bytesWritten, err)
		done <- err
	}()

	go func() {
		bytesWritten, err := io.Copy(clientConn, targetConn)
		log.Printf("[WEBSOCKET-FORWARD] Target -> Client: %d bytes transferred, error: %v\n", bytesWritten, err)
		done <- err
	}()

	err := <-done
	log.Printf("[WEBSOCKET-FORWARD] Data forwarding completed with error: %v\n", err)
	return err
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
