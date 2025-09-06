package socks5

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"golang.org/x/net/proxy"
)

// SOCKS5Client SOCKS5客户端实现
type SOCKS5Client struct {
	config        interfaces.ClientConfig
	conn          net.Conn
	authenticated bool
}

// NewSOCKS5Client 创建新的SOCKS5客户端
func NewSOCKS5Client(config interfaces.ClientConfig) *SOCKS5Client {
	return &SOCKS5Client{
		config: config,
	}
}

// Connect 连接到目标主机
func (c *SOCKS5Client) Connect(targetHost string, targetPort int) error {
	if c.config.ServerAddr == "" {
		return errors.New("server address not configured")
	}

	// 创建SOCKS5拨号器
	dialer, err := c.createSOCKS5Dialer()
	if err != nil {
		return fmt.Errorf("failed to create SOCKS5 dialer: %w", err)
	}

	// 连接到目标主机
	targetAddr := fmt.Sprintf("%s:%d", targetHost, targetPort)
	conn, err := dialer.Dial("tcp", targetAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to target %s: %w", targetAddr, err)
	}

	c.conn = conn
	return nil
}

// Authenticate 进行身份认证
func (c *SOCKS5Client) Authenticate(username, password string) error {
	if c.conn == nil {
		return errors.New("connection not established")
	}

	// SOCKS5认证握手
	// 发送认证方法选择
	authMethods := []byte{0x05, 0x01, 0x02} // 版本5，1种方法，用户名密码认证
	if _, err := c.conn.Write(authMethods); err != nil {
		return fmt.Errorf("failed to send auth methods: %w", err)
	}

	// 读取服务器选择的认证方法
	response := make([]byte, 2)
	if _, err := io.ReadFull(c.conn, response); err != nil {
		return fmt.Errorf("failed to read auth response: %w", err)
	}

	if response[0] != 0x05 {
		return errors.New("invalid SOCKS version")
	}

	if response[1] == 0x00 {
		// 无需认证
		c.authenticated = true
		return nil
	}

	if response[1] != 0x02 {
		return errors.New("server does not support username/password authentication")
	}

	// 发送用户名密码认证
	authRequest := c.buildAuthRequest(username, password)
	if _, err := c.conn.Write(authRequest); err != nil {
		return fmt.Errorf("failed to send auth request: %w", err)
	}

	// 读取认证响应
	authResponse := make([]byte, 2)
	if _, err := io.ReadFull(c.conn, authResponse); err != nil {
		return fmt.Errorf("failed to read auth response: %w", err)
	}

	if authResponse[1] != 0x00 {
		return errors.New("authentication failed")
	}

	c.authenticated = true
	return nil
}

// ForwardData 转发数据
func (c *SOCKS5Client) ForwardData(conn net.Conn) error {
	if c.conn == nil {
		return errors.New("connection not established")
	}

	if !c.authenticated {
		return errors.New("not authenticated")
	}

	// 双向数据转发
	done := make(chan error, 2)

	go func() {
		_, err := io.Copy(c.conn, conn)
		done <- err
	}()

	go func() {
		_, err := io.Copy(conn, c.conn)
		done <- err
	}()

	// 等待任意一个方向的复制完成
	return <-done
}

// Close 关闭连接
func (c *SOCKS5Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// createSOCKS5Dialer 创建SOCKS5拨号器
func (c *SOCKS5Client) createSOCKS5Dialer() (proxy.Dialer, error) {
	// 解析服务器地址，支持协议前缀
	serverAddr, useTLS, err := c.parseServerAddr(c.config.ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse server address: %w", err)
	}

	if c.config.Username != "" && c.config.Password != "" {
		// 带认证的SOCKS5拨号器
		auth := &proxy.Auth{
			User:     c.config.Username,
			Password: c.config.Password,
		}
		dialer, err := proxy.SOCKS5("tcp", serverAddr, auth, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to create authenticated SOCKS5 dialer: %w", err)
		}

		// 如果需要TLS，包装拨号器
		if useTLS {
			return c.createTLSDialer(dialer)
		}

		return dialer, nil
	}

	// 无认证的SOCKS5拨号器
	dialer, err := proxy.SOCKS5("tcp", serverAddr, nil, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOCKS5 dialer: %w", err)
	}

	// 如果需要TLS，包装拨号器
	if useTLS {
		return c.createTLSDialer(dialer)
	}

	return dialer, nil
}

// parseServerAddr 解析服务器地址，支持协议前缀
func (c *SOCKS5Client) parseServerAddr(addr string) (string, bool, error) {
	if strings.HasPrefix(addr, "tls://") {
		// 移除tls://前缀
		return strings.TrimPrefix(addr, "tls://"), true, nil
	} else if strings.HasPrefix(addr, "tcp://") {
		// 移除tcp://前缀
		return strings.TrimPrefix(addr, "tcp://"), false, nil
	}
	// 默认使用TCP
	return addr, false, nil
}

// createTLSDialer 创建TLS拨号器包装器
func (c *SOCKS5Client) createTLSDialer(baseDialer proxy.Dialer) (proxy.Dialer, error) {
	// 创建TLS配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证，生产环境应该设置为false
	}

	// 创建TLS拨号器包装器
	return &tlsDialerWrapper{
		baseDialer: baseDialer,
		tlsConfig:  tlsConfig,
	}, nil
}

// tlsDialerWrapper TLS拨号器包装器
type tlsDialerWrapper struct {
	baseDialer proxy.Dialer
	tlsConfig  *tls.Config
}

// Dial 实现proxy.Dialer接口
func (w *tlsDialerWrapper) Dial(network, addr string) (net.Conn, error) {
	// 使用基础拨号器建立连接
	conn, err := w.baseDialer.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	// 包装为TLS连接
	tlsConn := tls.Client(conn, w.tlsConfig)

	// 执行TLS握手
	if err := tlsConn.Handshake(); err != nil {
		conn.Close()
		return nil, err
	}

	return tlsConn, nil
}

// buildAuthRequest 构建认证请求
func (c *SOCKS5Client) buildAuthRequest(username, password string) []byte {
	var buf bytes.Buffer

	// 认证子协商版本
	buf.WriteByte(0x01)

	// 用户名长度和用户名
	buf.WriteByte(byte(len(username)))
	buf.WriteString(username)

	// 密码长度和密码
	buf.WriteByte(byte(len(password)))
	buf.WriteString(password)

	return buf.Bytes()
}

// 注册SOCKS5客户端创建函数
func init() {
	interfaces.RegisterClient("socks5", func(config interfaces.ClientConfig) (interfaces.ProxyClient, error) {
		return NewSOCKS5Client(config), nil
	})
}
