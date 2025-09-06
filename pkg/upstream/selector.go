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

	// 返回一个模拟的net.Conn接口
	return &SOCKS5ConnWrapper{client: client}, nil
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

	// 返回一个模拟的net.Conn接口
	return &WebSocketConnWrapper{client: client}, nil
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
}

func (w *SOCKS5ConnWrapper) Read(b []byte) (n int, err error) {
	// 这里需要实现从SOCKS5客户端读取数据的逻辑
	// 由于SOCKS5客户端接口没有直接提供Read方法，我们需要重新设计
	return 0, fmt.Errorf("SOCKS5ConnWrapper.Read not implemented")
}

func (w *SOCKS5ConnWrapper) Write(b []byte) (n int, err error) {
	// 这里需要实现向SOCKS5客户端写入数据的逻辑
	return 0, fmt.Errorf("SOCKS5ConnWrapper.Write not implemented")
}

func (w *SOCKS5ConnWrapper) Close() error {
	return w.client.Close()
}

func (w *SOCKS5ConnWrapper) LocalAddr() net.Addr {
	return &dummyAddr{addr: "local"}
}

func (w *SOCKS5ConnWrapper) RemoteAddr() net.Addr {
	return &dummyAddr{addr: "remote"}
}

func (w *SOCKS5ConnWrapper) SetDeadline(t time.Time) error {
	return nil
}

func (w *SOCKS5ConnWrapper) SetReadDeadline(t time.Time) error {
	return nil
}

func (w *SOCKS5ConnWrapper) SetWriteDeadline(t time.Time) error {
	return nil
}

// WebSocketConnWrapper WebSocket客户端连接包装器，实现net.Conn接口
type WebSocketConnWrapper struct {
	client interfaces.ProxyClient
}

func (w *WebSocketConnWrapper) Read(b []byte) (n int, err error) {
	// 这里需要实现从WebSocket客户端读取数据的逻辑
	return 0, fmt.Errorf("WebSocketConnWrapper.Read not implemented")
}

func (w *WebSocketConnWrapper) Write(b []byte) (n int, err error) {
	// 这里需要实现向WebSocket客户端写入数据的逻辑
	return 0, fmt.Errorf("WebSocketConnWrapper.Write not implemented")
}

func (w *WebSocketConnWrapper) Close() error {
	return w.client.Close()
}

func (w *WebSocketConnWrapper) LocalAddr() net.Addr {
	return &dummyAddr{addr: "local"}
}

func (w *WebSocketConnWrapper) RemoteAddr() net.Addr {
	return &dummyAddr{addr: "remote"}
}

func (w *WebSocketConnWrapper) SetDeadline(t time.Time) error {
	return nil
}

func (w *WebSocketConnWrapper) SetReadDeadline(t time.Time) error {
	return nil
}

func (w *WebSocketConnWrapper) SetWriteDeadline(t time.Time) error {
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