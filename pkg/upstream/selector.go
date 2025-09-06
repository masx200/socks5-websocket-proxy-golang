package upstream

import (
	"fmt"
	"net"
	"time"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

// UpstreamSelector 上游连接选择器
type UpstreamSelector struct {
	config *interfaces.UpstreamConfig
}

// NewUpstreamSelector 创建新的上游连接选择器
func NewUpstreamSelector(config *interfaces.UpstreamConfig) *UpstreamSelector {
	return &UpstreamSelector{
		config: config,
	}
}

// SelectConnection 选择上游连接
func (s *UpstreamSelector) SelectConnection(targetHost string, targetPort int) (net.Conn, error) {
	if s.config == nil {
		// 默认直连
		return s.createDirectConnection(targetHost, targetPort)
	}

	switch s.config.Type {
	case interfaces.UpstreamDirect:
		return s.createDirectConnection(targetHost, targetPort)
	case interfaces.UpstreamSOCKS5:
		return s.createSOCKS5Connection(targetHost, targetPort)
	case interfaces.UpstreamWebSocket:
		return s.createWebSocketConnection(targetHost, targetPort)
	default:
		return s.createDirectConnection(targetHost, targetPort)
	}
}

// createDirectConnection 创建TCP直连
func (s *UpstreamSelector) createDirectConnection(targetHost string, targetPort int) (net.Conn, error) {
	timeout := s.getTimeout()
	return net.DialTimeout("tcp", fmt.Sprintf("%s:%d", targetHost, targetPort), timeout)
}

// createSOCKS5Connection 创建SOCKS5代理连接
func (s *UpstreamSelector) createSOCKS5Connection(targetHost string, targetPort int) (net.Conn, error) {
	if s.config.ProxyAddress == "" {
		return nil, fmt.Errorf("SOCKS5 proxy address not configured")
	}

	// 使用工厂函数创建SOCKS5客户端
	client, err := interfaces.CreateClient("socks5", interfaces.ClientConfig{
		Username:   s.config.ProxyUsername,
		Password:   s.config.ProxyPassword,
		ServerAddr: s.config.ProxyAddress,
		Protocol:   "socks5",
		Timeout:    s.getTimeout(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create SOCKS5 client: %w", err)
	}

	// 连接到目标
	err = client.Connect(targetHost, targetPort)
	if err != nil {
		return nil, fmt.Errorf("failed to connect via SOCKS5 proxy: %w", err)
	}

	// 创建一个管道来模拟net.Conn接口
	clientConn, serverConn := net.Pipe()

	// 启动goroutine来处理数据转发
	go func() {
		defer clientConn.Close()
		defer serverConn.Close()

		// 使用客户端的ForwardData方法进行数据转发
		if socks5Client, ok := client.(interface {
			ForwardData(net.Conn) error
		}); ok {
			err := socks5Client.ForwardData(serverConn)
			if err != nil {
				fmt.Printf("SOCKS5 forward data error: %v\n", err)
			}
		}
	}()

	// 返回客户端连接
	return &SOCKS5ConnWrapper{client: client, conn: clientConn}, nil
}

// createWebSocketConnection 创建WebSocket代理连接
func (s *UpstreamSelector) createWebSocketConnection(targetHost string, targetPort int) (net.Conn, error) {
	if s.config.ProxyAddress == "" {
		return nil, fmt.Errorf("WebSocket proxy address not configured")
	}

	// 使用工厂函数创建WebSocket客户端
	client, err := interfaces.CreateClient("websocket", interfaces.ClientConfig{
		Username:   s.config.ProxyUsername,
		Password:   s.config.ProxyPassword,
		ServerAddr: s.config.ProxyAddress,
		Protocol:   "websocket",
		Timeout:    s.getTimeout(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create WebSocket client: %w", err)
	}

	// 连接到目标
	err = client.Connect(targetHost, targetPort)
	if err != nil {
		return nil, fmt.Errorf("failed to connect via WebSocket proxy: %w", err)
	}

	// 创建一个管道来模拟net.Conn接口
	clientConn, serverConn := net.Pipe()

	// 启动goroutine来处理数据转发
	go func() {
		defer clientConn.Close()
		defer serverConn.Close()

		// 使用客户端的ForwardData方法进行数据转发
		if wsClient, ok := client.(interface {
			ForwardData(net.Conn) error
		}); ok {
			err := wsClient.ForwardData(serverConn)
			if err != nil {
				fmt.Printf("WebSocket forward data error: %v\n", err)
			}
		}
	}()

	// 返回客户端连接
	return &WebSocketConnWrapper{client: client, conn: clientConn}, nil
}

// getTimeout 获取超时时间
func (s *UpstreamSelector) getTimeout() time.Duration {
	if s.config.Timeout > 0 {
		return s.config.Timeout
	}
	return 30 * time.Second // 默认30秒超时
}

// SOCKS5ConnWrapper SOCKS5客户端连接包装器，实现net.Conn接口
type SOCKS5ConnWrapper struct {
	client interfaces.ProxyClient
	conn   net.Conn
	closed bool
}

func (w *SOCKS5ConnWrapper) Read(b []byte) (n int, err error) {
	if w.closed {
		return 0, net.ErrClosed
	}
	if w.conn == nil {
		return 0, fmt.Errorf("connection not established")
	}
	return w.conn.Read(b)
}

func (w *SOCKS5ConnWrapper) Write(b []byte) (n int, err error) {
	if w.closed {
		return 0, net.ErrClosed
	}
	if w.conn == nil {
		return 0, fmt.Errorf("connection not established")
	}
	return w.conn.Write(b)
}

func (w *SOCKS5ConnWrapper) Close() error {
	if w.closed {
		return nil
	}
	w.closed = true
	if w.conn != nil {
		return w.conn.Close()
	}
	return w.client.Close()
}

func (w *SOCKS5ConnWrapper) LocalAddr() net.Addr {
	if w.conn != nil {
		return w.conn.LocalAddr()
	}
	return &dummyAddr{addr: "local"}
}

func (w *SOCKS5ConnWrapper) RemoteAddr() net.Addr {
	if w.conn != nil {
		return w.conn.RemoteAddr()
	}
	return &dummyAddr{addr: "remote"}
}

func (w *SOCKS5ConnWrapper) SetDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetDeadline(t)
	}
	return nil
}

func (w *SOCKS5ConnWrapper) SetReadDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetReadDeadline(t)
	}
	return nil
}

func (w *SOCKS5ConnWrapper) SetWriteDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetWriteDeadline(t)
	}
	return nil
}

// WebSocketConnWrapper WebSocket客户端连接包装器，实现net.Conn接口
type WebSocketConnWrapper struct {
	client interfaces.ProxyClient
	conn   net.Conn
	closed bool
}

func (w *WebSocketConnWrapper) Read(b []byte) (n int, err error) {
	if w.closed {
		return 0, net.ErrClosed
	}
	if w.conn == nil {
		return 0, fmt.Errorf("connection not established")
	}
	return w.conn.Read(b)
}

func (w *WebSocketConnWrapper) Write(b []byte) (n int, err error) {
	if w.closed {
		return 0, net.ErrClosed
	}
	if w.conn == nil {
		return 0, fmt.Errorf("connection not established")
	}
	return w.conn.Write(b)
}

func (w *WebSocketConnWrapper) Close() error {
	if w.closed {
		return nil
	}
	w.closed = true
	if w.conn != nil {
		return w.conn.Close()
	}
	return w.client.Close()
}

func (w *WebSocketConnWrapper) LocalAddr() net.Addr {
	if w.conn != nil {
		return w.conn.LocalAddr()
	}
	return &dummyAddr{addr: "local"}
}

func (w *WebSocketConnWrapper) RemoteAddr() net.Addr {
	if w.conn != nil {
		return w.conn.RemoteAddr()
	}
	return &dummyAddr{addr: "remote"}
}

func (w *WebSocketConnWrapper) SetDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetDeadline(t)
	}
	return nil
}

func (w *WebSocketConnWrapper) SetReadDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetReadDeadline(t)
	}
	return nil
}

func (w *WebSocketConnWrapper) SetWriteDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetWriteDeadline(t)
	}
	return nil
}

// dummyAddr 虚拟地址实现
type dummyAddr struct {
	addr string
}

func (d *dummyAddr) Network() string {
	return "tcp"
}

func (d *dummyAddr) String() string {
	return d.addr
}
