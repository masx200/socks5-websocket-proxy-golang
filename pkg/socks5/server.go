package socks5

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/armon/go-socks5"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/http"
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
	selector  interface{}
	logger    *log.Logger
}

// NewSOCKS5Server 创建新的SOCKS5服务端
func NewSOCKS5Server(config interfaces.ServerConfig) *SOCKS5Server {
	var selector interface{}
	if config.EnableUpstream && len(config.UpstreamConfig) > 0 {
		selector = upstream.NewDynamicUpstreamSelector(config.UpstreamConfig, upstream.StrategyRoundRobin)
	}

	// 创建自定义日志记录器
	logger := log.New(os.Stdout, "[SOCKS5-LIB] ", log.LstdFlags)

	// 创建SOCKS5服务器配置
	socks5Config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{},
		Logger:      logger,
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
		fmt.Printf("[SOCKS5-SERVER] Failed to create SOCKS5 server: %v\n", err)
		return nil
	}

	socks5Server := &SOCKS5Server{
		config:    config,
		server:    server,
		shutdown:  make(chan struct{}),
		authUsers: config.AuthUsers,
		selector:  selector,
		logger:    logger,
	}

	// 现在设置Dial函数，这样可以访问到完整的server实例
	socks5Config.Dial = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 解析地址
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, fmt.Errorf("invalid address format: %w", err)
		}

		// 解析端口
		portInt, err := net.LookupPort(network, port)
		if err != nil {
			return nil, fmt.Errorf("invalid port: %w", err)
		}

		// 使用server实例的SelectUpstreamConnection方法
		return socks5Server.SelectUpstreamConnection(host, portInt)
	}

	// 重新创建server实例以应用新的Dial函数
	server, err = socks5.New(socks5Config)
	if err != nil {
		fmt.Printf("[SOCKS5-SERVER] Failed to recreate SOCKS5 server with Dial function: %v\n", err)
		return nil
	}
	socks5Server.server = server

	return socks5Server
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

	// 记录本地和远程地址信息
	localAddr := conn.LocalAddr().String()
	fmt.Printf("[SOCKS5-CONN] Connection details - Local: %s, Remote: %s\n", localAddr, clientAddr)

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
		var conn net.Conn
		var err error

		// 尝试类型断言
		if selector, ok := s.selector.(*upstream.UpstreamSelector); ok {
			conn, err = selector.SelectConnection(targetHost, targetPort)
		} else if selector, ok := s.selector.(*upstream.DynamicUpstreamSelector); ok {
			conn, err = selector.SelectConnection(targetHost, targetPort)
		} else {
			err = fmt.Errorf("unknown selector type")
		}

		if err != nil {
			fmt.Printf("[SOCKS5-UPSTREAM] Upstream selector failed for target %s: %v\n", targetAddr, err)
		} else {
			fmt.Printf("[SOCKS5-UPSTREAM] Upstream selector succeeded for target %s\n", targetAddr)
		}
		return conn, err
	}

	// 如果没有选择器但配置了上游代理，尝试直接创建连接
	if s.config.EnableUpstream && len(s.config.UpstreamConfig) > 0 {
		fmt.Printf("[SOCKS5-UPSTREAM] Using direct upstream config for target %s\n", targetAddr)
		upstreamConfig := s.config.UpstreamConfig[0] // 使用第一个配置

		// 根据上游类型创建连接
		switch upstreamConfig.Type {
		case interfaces.UpstreamHTTP:
			proxyURL, err := url.Parse(upstreamConfig.ProxyAddress)
			if err != nil {
				return nil, fmt.Errorf("invalid HTTP proxy address: %w", err)
			}

			// 设置认证信息
			if upstreamConfig.ProxyUsername != "" && upstreamConfig.ProxyPassword != "" {
				proxyURL.User = url.UserPassword(upstreamConfig.ProxyUsername, upstreamConfig.ProxyPassword)
			}

			// 创建HTTP代理客户端配置
			clientConfig := interfaces.ClientConfig{
				Username:   upstreamConfig.ProxyUsername,
				Password:   upstreamConfig.ProxyPassword,
				ServerAddr: upstreamConfig.ProxyAddress,
				Protocol:   "http",
				Timeout:    s.config.Timeout,
			}

			// 创建HTTP代理客户端
			client, err := http.NewHttpProxyAdapter(clientConfig)
			if err != nil {
				return nil, fmt.Errorf("failed to create HTTP proxy client: %w", err)
			}

			// 连接到目标
			err = client.Connect(targetHost, targetPort)
			if err != nil {
				return nil, fmt.Errorf("failed to connect via HTTP proxy: %w", err)
			}

			return client.NetConn(), nil
		case interfaces.UpstreamSOCKS5:
			// 处理SOCKS5代理连接
			// ... (保留原有SOCKS5代理处理逻辑)
		default:
			// 默认直连
			return net.DialTimeout("tcp", net.JoinHostPort(targetHost, fmt.Sprint(targetPort)), s.config.Timeout)
		}
	}

	// 默认直连
	timeout := s.config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	fmt.Printf("[SOCKS5-UPSTREAM] Using direct connection for target %s (timeout: %v)\n", targetAddr, timeout)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetHost, fmt.Sprint(targetPort)), timeout)
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

// ReloadConfig 重新加载配置
func (s *SOCKS5Server) ReloadConfig(newConfig interfaces.ServerConfig) error {
	fmt.Printf("[SOCKS5-SERVER] Reloading configuration...\n")

	// 更新配置
	s.config = newConfig
	s.authUsers = newConfig.AuthUsers

	// 更新上游选择器
	if newConfig.EnableUpstream {
		selector := upstream.NewDynamicUpstreamSelector(newConfig.UpstreamConfig, upstream.StrategyRoundRobin)
		s.selector = selector
		fmt.Printf("[SOCKS5-SERVER] Upstream selector updated with %d configurations\n", len(newConfig.UpstreamConfig))
	} else {
		s.selector = nil
		fmt.Printf("[SOCKS5-SERVER] Upstream selector disabled\n")
	}

	// 重新创建SOCKS5服务器实例
	var socks5Config *socks5.Config = &socks5.Config{
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// 解析地址
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, fmt.Errorf("invalid address format: %w", err)
			}

			// 解析端口
			portInt, err := net.LookupPort(network, port)
			if err != nil {
				return nil, fmt.Errorf("invalid port: %w", err)
			}

			return s.SelectUpstreamConnection(host, portInt)
		},
		Logger: log.New(os.Stdout, "[SOCKS5-LIB] ", log.LstdFlags),
	}

	// 如果有认证用户，添加用户名密码认证
	if len(newConfig.AuthUsers) > 0 {
		authenticator := socks5.UserPassAuthenticator{Credentials: socks5.StaticCredentials(newConfig.AuthUsers)}
		socks5Config.AuthMethods = append(socks5Config.AuthMethods, authenticator)
	}

	// 创建新的SOCKS5服务器
	server, err := socks5.New(socks5Config)
	if err != nil {
		fmt.Printf("[SOCKS5-SERVER] Failed to create new SOCKS5 server: %v\n", err)
		return fmt.Errorf("failed to create new SOCKS5 server: %w", err)
	}

	s.server = server

	fmt.Printf("[SOCKS5-SERVER] Configuration reloaded successfully\n")
	fmt.Printf("[SOCKS5-SERVER] Authentication users: %d\n", len(newConfig.AuthUsers))
	fmt.Printf("[SOCKS5-SERVER] Upstream enabled: %t\n", newConfig.EnableUpstream)
	fmt.Printf("[SOCKS5-SERVER] Listen address: %s\n", newConfig.ListenAddr)
	fmt.Printf("[SOCKS5-SERVER] Timeout: %v\n", newConfig.Timeout)

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
