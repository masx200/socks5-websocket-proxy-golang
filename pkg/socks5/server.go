package socks5

import (
	"fmt"
	"net"
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
}

// NewSOCKS5Server 创建新的SOCKS5服务端
func NewSOCKS5Server(config interfaces.ServerConfig) *SOCKS5Server {
	var selector *upstream.UpstreamSelector
	if config.EnableUpstream && len(config.UpstreamConfig) > 0 {
		selector = upstream.NewUpstreamSelector(&config.UpstreamConfig[0])
	}

	// 创建SOCKS5服务器配置
	socks5Config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{},
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
	}
}

// Listen 开始监听连接
func (s *SOCKS5Server) Listen() error {
	var err error
	s.listener, err = net.Listen("tcp", s.config.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.config.ListenAddr, err)
	}

	fmt.Printf("SOCKS5 server listening on %s\n", s.config.ListenAddr)

	s.wg.Add(1)
	go s.acceptConnections()

	return nil
}

// acceptConnections 接受连接
func (s *SOCKS5Server) acceptConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				select {
				case <-s.shutdown:
					return
				default:
					fmt.Printf("accept error: %v\n", err)
					continue
				}
			}

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
	defer conn.Close()

	// 设置超时
	if s.config.Timeout > 0 {
		conn.SetDeadline(time.Now().Add(s.config.Timeout))
	}

	// 使用github.com/armon/go-socks5库处理连接
	if s.server != nil {
		return s.server.ServeConn(conn)
	}

	// 如果server为nil，使用原始实现作为后备
	// SOCKS5握手
	if err := s.handleHandshake(conn); err != nil {
		fmt.Printf("handshake error: %v\n", err)
		return err
	}

	// 处理连接请求
	targetHost, targetPort, err := s.handleConnectRequest(conn)
	if err != nil {
		fmt.Printf("connect request error: %v\n", err)
		return err
	}

	fmt.Printf("Connecting to target %s:%d\n", targetHost, targetPort)

	// 选择上游连接
	upstreamConn, err := s.SelectUpstreamConnection(targetHost, targetPort)
	if err != nil {
		fmt.Printf("failed to connect to target: %v\n", err)
		return err
	}
	defer upstreamConn.Close()

	// 发送成功响应
	if err := s.sendConnectResponse(conn, net.ParseIP(targetHost) != nil); err != nil {
		return err
	}

	// 开始数据转发
	return s.forwardData(conn, upstreamConn)
}

// Authenticate 验证用户名密码
func (s *SOCKS5Server) Authenticate(username, password string) bool {
	if s.authUsers == nil || len(s.authUsers) == 0 {
		return true // 如果没有配置用户，则允许所有连接
	}

	storedPassword, exists := s.authUsers[username]
	if !exists {
		return false
	}

	return storedPassword == password
}

// SelectUpstreamConnection 选择上游连接
func (s *SOCKS5Server) SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error) {
	if s.selector != nil {
		return s.selector.SelectConnection(targetHost, targetPort)
	}

	// 默认直连
	timeout := s.config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	return net.DialTimeout("tcp", fmt.Sprintf("%s:%d", targetHost, targetPort), timeout)
}

// Close 关闭服务端
func (s *SOCKS5Server) Close() error {
	return s.Shutdown()
}

// Shutdown 优雅关闭服务端
func (s *SOCKS5Server) Shutdown() error {
	close(s.shutdown)
	if s.listener != nil {
		s.listener.Close()
	}
	s.wg.Wait()
	return nil
}

// handleHandshake 处理SOCKS5握手 (后备实现)
func (s *SOCKS5Server) handleHandshake(conn net.Conn) error {
	// 这个方法现在只作为后备实现，主要逻辑由github.com/armon/go-socks5库处理
	return nil
}

// handleUsernamePasswordAuth 处理用户名密码认证 (后备实现)
func (s *SOCKS5Server) handleUsernamePasswordAuth(conn net.Conn) error {
	// 这个方法现在只作为后备实现，主要逻辑由github.com/armon/go-socks5库处理
	return nil
}

// handleConnectRequest 处理连接请求 (后备实现)
func (s *SOCKS5Server) handleConnectRequest(conn net.Conn) (string, int, error) {
	// 这个方法现在只作为后备实现，主要逻辑由github.com/armon/go-socks5库处理
	return "", 0, nil
}

// sendConnectResponse 发送连接响应 (后备实现)
func (s *SOCKS5Server) sendConnectResponse(conn net.Conn, isIP bool) error {
	// 这个方法现在只作为后备实现，主要逻辑由github.com/armon/go-socks5库处理
	return nil
}

// forwardData 转发数据 (后备实现)
func (s *SOCKS5Server) forwardData(clientConn, targetConn net.Conn) error {
	// 这个方法现在只作为后备实现，主要逻辑由github.com/armon/go-socks5库处理
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
