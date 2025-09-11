package socks5

import (
	// "bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"golang.org/x/net/proxy"
)

// SOCKS5Client SOCKS5客户端实现
type SOCKS5Client struct {
	config                   interfaces.ClientConfig
	conn                     net.Conn
	authenticated            bool
	connectionClosedCallback func()
}

// DialContext implements interfaces.ProxyClient.
func (c *SOCKS5Client) DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
	// 解析目标地址
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address %s: %w", addr, err)
	}

	// 将端口转换为整数
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("failed to parse port %s: %w", port, err)
	}

	// 创建管道连接
	clientConn, serverConn := net.Pipe()

	// 使用Connect方法建立连接
	go func() {
		if err := c.Connect(host, portInt); err != nil {
			serverConn.Close()
			return
		}

		// 使用ForwardData进行数据转发
		if err := c.ForwardData(serverConn); err != nil {
			serverConn.Close()
			return
		}
	}()

	// 返回客户端连接
	return clientConn, nil
}

// NetConn implements interfaces.ProxyClient.
func (c *SOCKS5Client) NetConn() net.Conn {
	return c.conn
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
	c.authenticated = true
	return nil
}

// Authenticate 进行身份认证
// func (c *SOCKS5Client) Authenticate(username, password string) error {
// 	// 使用golang.org/x/net/proxy.SOCKS5处理认证
// 	// 由于Connect方法已经使用了proxy.SOCKS5拨号器，认证已经在拨号过程中完成
// 	// 这里只需要标记认证状态即可
// 	c.authenticated = true
// 	return nil
// }

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
		if c.connectionClosedCallback != nil {
			c.connectionClosedCallback()
		}
		done <- err
	}()

	go func() {
		_, err := io.Copy(conn, c.conn)
		if c.connectionClosedCallback != nil {
			c.connectionClosedCallback()
		}
		done <- err
	}()

	// 等待任意一个方向的复制完成
	return <-done
}

// Close 关闭连接
func (c *SOCKS5Client) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		// 如果设置了连接关闭回调，则调用它
		if c.connectionClosedCallback != nil {
			c.connectionClosedCallback()
		}
		return err
	}
	return nil
}

// SetConnectionClosedCallback 设置连接关闭回调
func (c *SOCKS5Client) SetConnectionClosedCallback(callback func()) error {
	c.connectionClosedCallback = callback
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
		// 移除tls://前缀，并去除末尾的斜杠
		parsedAddr := strings.TrimSuffix(strings.TrimPrefix(addr, "tls://"), "/")
		return parsedAddr, true, nil
	} else if strings.HasPrefix(addr, "tcp://") {
		// 移除tcp://前缀，并去除末尾的斜杠
		parsedAddr := strings.TrimSuffix(strings.TrimPrefix(addr, "tcp://"), "/")
		return parsedAddr, false, nil
	}
	// 默认使用TCP，去除末尾的斜杠
	parsedAddr := strings.TrimSuffix(addr, "/")
	return parsedAddr, false, nil
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
// func (c *SOCKS5Client) buildAuthRequest(username, password string) []byte {
// 	var buf bytes.Buffer

// 	// 认证子协商版本
// 	buf.WriteByte(0x01)

// 	// 用户名长度和用户名
// 	buf.WriteByte(byte(len(username)))
// 	buf.WriteString(username)

// 	// 密码长度和密码
// 	buf.WriteByte(byte(len(password)))
// 	buf.WriteString(password)

// 	return buf.Bytes()
// }

// 注册SOCKS5客户端创建函数
func init() {
	interfaces.RegisterClient("socks5", func(config interfaces.ClientConfig) (interfaces.ProxyClient, error) {
		return NewSOCKS5Client(config), nil
	})
}
