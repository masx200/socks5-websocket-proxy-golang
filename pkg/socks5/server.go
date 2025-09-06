package socks5

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/upstream"
)

// SOCKS5Server SOCKS5服务端实现
type SOCKS5Server struct {
	config    interfaces.ServerConfig
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

	return &SOCKS5Server{
		config:    config,
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

// handleHandshake 处理SOCKS5握手
func (s *SOCKS5Server) handleHandshake(conn net.Conn) error {
	reader := bufio.NewReader(conn)

	// 读取客户端认证方法请求
	header := make([]byte, 2)
	if _, err := io.ReadFull(reader, header); err != nil {
		return fmt.Errorf("failed to read handshake header: %w", err)
	}

	if header[0] != 0x05 {
		return errors.New("invalid SOCKS version")
	}

	methodCount := int(header[1])
	methods := make([]byte, methodCount)
	if _, err := io.ReadFull(reader, methods); err != nil {
		return fmt.Errorf("failed to read auth methods: %w", err)
	}

	// 选择认证方法
	var selectedMethod byte = 0x00 // 无认证
	if s.authUsers != nil && len(s.authUsers) > 0 {
		// 检查是否支持用户名密码认证
		for _, method := range methods {
			if method == 0x02 {
				selectedMethod = 0x02
				break
			}
		}
	}

	// 发送选择的认证方法
	response := []byte{0x05, selectedMethod}
	if _, err := conn.Write(response); err != nil {
		return fmt.Errorf("failed to send auth method response: %w", err)
	}

	// 如果需要用户名密码认证
	if selectedMethod == 0x02 {
		return s.handleUsernamePasswordAuth(conn)
	}

	return nil
}

// handleUsernamePasswordAuth 处理用户名密码认证
func (s *SOCKS5Server) handleUsernamePasswordAuth(conn net.Conn) error {
	reader := bufio.NewReader(conn)

	// 读取认证请求
	authHeader := make([]byte, 2)
	if _, err := io.ReadFull(reader, authHeader); err != nil {
		return fmt.Errorf("failed to read auth header: %w", err)
	}

	if authHeader[0] != 0x01 {
		return errors.New("invalid auth subnegotiation version")
	}

	usernameLen := int(authHeader[1])
	username := make([]byte, usernameLen)
	if _, err := io.ReadFull(reader, username); err != nil {
		return fmt.Errorf("failed to read username: %w", err)
	}

	passwordLenByte, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read password length: %w", err)
	}

	passwordLen := int(passwordLenByte)
	password := make([]byte, passwordLen)
	if _, err := io.ReadFull(reader, password); err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}

	// 验证用户名密码
	usernameStr := string(username)
	passwordStr := string(password)

	var authResult byte = 0x00 // 成功
	if !s.Authenticate(usernameStr, passwordStr) {
		authResult = 0x01 // 失败
	}

	// 发送认证结果
	authResponse := []byte{0x01, authResult}
	if _, err := conn.Write(authResponse); err != nil {
		return fmt.Errorf("failed to send auth response: %w", err)
	}

	if authResult != 0x00 {
		return errors.New("authentication failed")
	}

	return nil
}

// handleConnectRequest 处理连接请求
func (s *SOCKS5Server) handleConnectRequest(conn net.Conn) (string, int, error) {
	reader := bufio.NewReader(conn)

	// 读取连接请求
	request := make([]byte, 4)
	if _, err := io.ReadFull(reader, request); err != nil {
		return "", 0, fmt.Errorf("failed to read connect request: %w", err)
	}

	if request[0] != 0x05 {
		return "", 0, errors.New("invalid SOCKS version")
	}

	if request[1] != 0x01 {
		return "", 0, errors.New("only CONNECT command supported")
	}

	addrType := request[3]
	var targetHost string
	var targetPort int

	switch addrType {
	case 0x01: // IPv4
		ipBytes := make([]byte, 4)
		if _, err := io.ReadFull(reader, ipBytes); err != nil {
			return "", 0, fmt.Errorf("failed to read IPv4 address: %w", err)
		}
		targetHost = net.IP(ipBytes).String()

	case 0x03: // 域名
		domainLenByte, err := reader.ReadByte()
		if err != nil {
			return "", 0, fmt.Errorf("failed to read domain length: %w", err)
		}
		domainLen := int(domainLenByte)
		domainBytes := make([]byte, domainLen)
		if _, err := io.ReadFull(reader, domainBytes); err != nil {
			return "", 0, fmt.Errorf("failed to read domain: %w", err)
		}
		targetHost = string(domainBytes)

	case 0x04: // IPv6
		ipBytes := make([]byte, 16)
		if _, err := io.ReadFull(reader, ipBytes); err != nil {
			return "", 0, fmt.Errorf("failed to read IPv6 address: %w", err)
		}
		targetHost = net.IP(ipBytes).String()

	default:
		return "", 0, errors.New("unsupported address type")
	}

	// 读取端口号
	portBytes := make([]byte, 2)
	if _, err := io.ReadFull(reader, portBytes); err != nil {
		return "", 0, fmt.Errorf("failed to read port: %w", err)
	}
	targetPort = int(binary.BigEndian.Uint16(portBytes))

	return targetHost, targetPort, nil
}

// sendConnectResponse 发送连接响应
func (s *SOCKS5Server) sendConnectResponse(conn net.Conn, isIP bool) error {
	var response []byte

	if isIP {
		// 如果是IP地址，返回绑定地址
		response = []byte{
			0x05,                   // SOCKS版本
			0x00,                   // 成功
			0x00,                   // RSV
			0x01,                   // 地址类型 (IPv4)
			0x00, 0x00, 0x00, 0x00, // 绑定地址 (0.0.0.0)
			0x00, 0x00, // 绑定端口 (0)
		}
	} else {
		// 如果是域名，返回域名格式
		response = []byte{
			0x05,       // SOCKS版本
			0x00,       // 成功
			0x00,       // RSV
			0x03,       // 地址类型 (域名)
			0x00,       // 域名长度 (0)
			0x00, 0x00, // 绑定端口 (0)
		}
	}

	_, err := conn.Write(response)
	return err
}

// forwardData 转发数据
func (s *SOCKS5Server) forwardData(clientConn, targetConn net.Conn) error {
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

// 注册SOCKS5服务端创建函数
func init() {
	interfaces.RegisterServer("socks5", func(config interfaces.ServerConfig) (interfaces.ProxyServer, error) {
		return NewSOCKS5Server(config), nil
	})
}
