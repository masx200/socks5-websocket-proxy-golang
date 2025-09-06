package socks5

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/armon/go-socks5"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/upstream"
)

// SOCKS5Server SOCKS5服务端实现
type SOCKS5Server struct {
	config    interfaces.ServerConfig
	server    *socks5.Server
	listener  net.Listener
	shutdown  chan struct{}
	wg        sync.WaitGroup
	authUsers map[string]string
	selector  *upstream.UpstreamSelector
	logger    *log.Logger
}

// NewSOCKS5Server 创建新的SOCKS5服务端
func NewSOCKS5Server(config interfaces.ServerConfig) *SOCKS5Server {
	var selector *upstream.UpstreamSelector
	if config.EnableUpstream && len(config.UpstreamConfig) > 0 {
		selector = upstream.NewUpstreamSelector(&config.UpstreamConfig[0])
	}

	// 创建自定义日志记录器
	logger := log.New(os.Stdout, "[SOCKS5-LIB] ", log.LstdFlags)

	// 创建SOCKS5服务器配置
	socks5Config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{},
		Logger:      logger,
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dailer := net.Dialer{}

			log.Printf("Dial %s %s", network, addr)
			return dailer.DialContext(ctx, network, addr)
		},
	}

	// 如果有认证用户，添加用户名密码认证
	if len(config.AuthUsers) > 0 {
		authenticator := socks5.UserPassAuthenticator{Credentials: socks5.StaticCredentials(config.AuthUsers)}
		socks5Config.AuthMethods = append(socks5Config.AuthMethods, authenticator)
	}

	// 创建SOCKS5服务器
	server, err := socks5.New(socks5Config)
	if err != nil {
		// 如果创建失败，返回nil
		fmt.Printf("Failed to create SOCKS5 server: %v\n", err)
		return nil
	}

	return &SOCKS5Server{
		config:    config,
		server:    server,
		shutdown:  make(chan struct{}),
		authUsers: config.AuthUsers,
		selector:  selector,
		logger:    logger,
	}
}

// Listen 开始监听连接
func (s *SOCKS5Server) Listen() error {
	var err error
	s.listener, err = net.Listen("tcp", s.config.ListenAddr)
	if err != nil {
		fmt.Printf("[SOCKS5-SERVER] Failed to listen on %s: %v\n", s.config.ListenAddr, err)
		return fmt.Errorf("failed to listen on %s: %w", s.config.ListenAddr, err)
	}

	fmt.Printf("[SOCKS5-SERVER] Server started successfully, listening on %s\n", s.config.ListenAddr)
	fmt.Printf("[SOCKS5-SERVER] Authentication enabled: %t (%d users configured)\n",
		len(s.authUsers) > 0, len(s.authUsers))
	fmt.Printf("[SOCKS5-SERVER] Upstream selector enabled: %t\n", s.config.EnableUpstream)

	s.wg.Add(1)
	go s.acceptConnections()

	return nil
}

// acceptConnections 接受连接
func (s *SOCKS5Server) acceptConnections() {
	defer s.wg.Done()
	connectionCount := 0

	for {
		select {
		case <-s.shutdown:
			fmt.Printf("[SOCKS5-SERVER] Shutdown signal received, stopping connection acceptance. Total connections handled: %d\n", connectionCount)
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				select {
				case <-s.shutdown:
					return
				default:
					fmt.Printf("[SOCKS5-SERVER] Accept error: %v\n", err)
					continue
				}
			}

			connectionCount++
			clientAddr := conn.RemoteAddr().String()
			fmt.Printf("[SOCKS5-SERVER] New connection accepted from %s (total: %d)\n", clientAddr, connectionCount)

			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				s.HandleConnection(conn)
			}()
		}
	}
}

// HandleConnection 处理客户端连接
func (s *SOCKS5Server) HandleConnection(conn net.Conn) error {
	defer func() {
		conn.Close()
		clientAddr := conn.RemoteAddr().String()
		fmt.Printf("[SOCKS5-CONN] Connection closed for client %s\n", clientAddr)
	}()

	clientAddr := conn.RemoteAddr().String()
	startTime := time.Now()
	fmt.Printf("[SOCKS5-CONN] New connection from %s at %s\n", clientAddr, startTime.Format("2006-01-02 15:04:05"))

	// 设置超时
	if s.config.Timeout > 0 {
		timeoutDuration := s.config.Timeout
		conn.SetDeadline(time.Now().Add(timeoutDuration))
		fmt.Printf("[SOCKS5-CONN] Connection timeout set to %v for client %s\n", timeoutDuration, clientAddr)
	}

	// 记录认证配置状态
	if len(s.authUsers) > 0 {
		fmt.Printf("[SOCKS5-CONN] Authentication required for client %s (%d users configured)\n", clientAddr, len(s.authUsers))
	} else {
		fmt.Printf("[SOCKS5-CONN] No authentication required for client %s\n", clientAddr)
	}

	// 使用github.com/armon/go-socks5库处理连接
	err := s.server.ServeConn(conn)
	duration := time.Since(startTime)

	if err != nil {
		fmt.Printf("[SOCKS5-CONN] Connection failed for client %s after %v: %v\n", clientAddr, duration, err)
	} else {
		fmt.Printf("[SOCKS5-CONN] Connection completed successfully for client %s, duration: %v\n", clientAddr, duration)
	}
	return err
}

// Authenticate 验证用户名密码
func (s *SOCKS5Server) Authenticate(username, password string) bool {
	if s.authUsers == nil || len(s.authUsers) == 0 {
		fmt.Printf("[SOCKS5-AUTH] No authentication configured, allowing access for user '%s'\n", username)
		return true // 如果没有配置用户，则允许所有连接
	}

	storedPassword, exists := s.authUsers[username]
	if !exists {
		fmt.Printf("[SOCKS5-AUTH] Authentication failed: user '%s' not found in configured users\n", username)
		return false
	}

	if storedPassword == password {
		fmt.Printf("[SOCKS5-AUTH] Authentication successful for user '%s'\n", username)
		return true
	} else {
		fmt.Printf("[SOCKS5-AUTH] Authentication failed: invalid password for user '%s'\n", username)
		return false
	}
}

// SelectUpstreamConnection 选择上游连接
func (s *SOCKS5Server) SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error) {
	targetAddr := fmt.Sprintf("%s:%d", targetHost, targetPort)

	if s.selector != nil {
		fmt.Printf("[SOCKS5-UPSTREAM] Using upstream selector for target %s\n", targetAddr)
		conn, err := s.selector.SelectConnection(targetHost, targetPort)
		if err != nil {
			fmt.Printf("[SOCKS5-UPSTREAM] Upstream selector failed for target %s: %v\n", targetAddr, err)
		} else {
			fmt.Printf("[SOCKS5-UPSTREAM] Upstream selector succeeded for target %s\n", targetAddr)
		}
		return conn, err
	}

	// 默认直连
	timeout := s.config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	fmt.Printf("[SOCKS5-UPSTREAM] Using direct connection for target %s (timeout: %v)\n", targetAddr, timeout)
	conn, err := net.DialTimeout("tcp", targetAddr, timeout)
	if err != nil {
		fmt.Printf("[SOCKS5-UPSTREAM] Direct connection failed for target %s: %v\n", targetAddr, err)
	} else {
		fmt.Printf("[SOCKS5-UPSTREAM] Direct connection established for target %s\n", targetAddr)
	}
	return conn, err
}

// Close 关闭服务端
func (s *SOCKS5Server) Close() error {
	return s.Shutdown()
}

// Shutdown 优雅关闭服务端
func (s *SOCKS5Server) Shutdown() error {
	fmt.Printf("[SOCKS5-SERVER] Initiating graceful shutdown...\n")
	close(s.shutdown)
	if s.listener != nil {
		fmt.Printf("[SOCKS5-SERVER] Closing listener...\n")
		s.listener.Close()
	}
	fmt.Printf("[SOCKS5-SERVER] Waiting for active connections to close...\n")
	s.wg.Wait()
	fmt.Printf("[SOCKS5-SERVER] Server shutdown completed successfully\n")
	return nil
}

// 注册SOCKS5服务端创建函数
func init() {
	interfaces.RegisterServer("socks5", func(config interfaces.ServerConfig) (interfaces.ProxyServer, error) {
		server := NewSOCKS5Server(config)
		if server == nil {
			return nil, fmt.Errorf("failed to create SOCKS5 server")
		}
		return server, nil
	})
}
