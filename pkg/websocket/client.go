package websocket

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

func init() {
	// 注册WebSocket客户端创建函数
	interfaces.RegisterClient("websocket", func(config interfaces.ClientConfig) (interfaces.ProxyClient, error) {
		return NewWebSocketClient(config), nil
	})
}

// WebSocketClient WebSocket客户端实现
type WebSocketClient struct {
	config     interfaces.ClientConfig
	conn       *websocket.Conn
	httpClient *http.Client
}

// NewWebSocketClient 创建新的WebSocket客户端
func NewWebSocketClient(config interfaces.ClientConfig) *WebSocketClient {
	return &WebSocketClient{
		config:     config,
		httpClient: &http.Client{Timeout: config.Timeout},
	}
}

// Connect 连接到目标主机
func (c *WebSocketClient) Connect(targetHost string, targetPort int) error {
	if c.config.ServerAddr == "" {
		return errors.New("server address not configured")
	}

	// 构建WebSocket URL
	wsURL, err := c.buildWebSocketURL(targetHost, targetPort)
	if err != nil {
		return fmt.Errorf("failed to build WebSocket URL: %w", err)
	}

	// 创建请求头
	headers := c.buildHeaders(targetHost, targetPort)

	// 建立WebSocket连接
	dialer := websocket.Dialer{
		HandshakeTimeout: c.config.Timeout,
	}

	conn, _, err := dialer.Dial(wsURL, headers)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket server: %w", err)
	}

	c.conn = conn
	return nil
}

// Authenticate 进行身份认证
func (c *WebSocketClient) Authenticate(username, password string) error {
	if c.conn == nil {
		return errors.New("connection not established")
	}

	// WebSocket认证通过HTTP Headers在连接时已经完成
	// 这里只需要验证连接是否成功建立
	return nil
}

// ForwardData 转发数据
func (c *WebSocketClient) ForwardData(conn net.Conn) error {
	if c.conn == nil {
		return errors.New("connection not established")
	}

	// 双向数据转发
	done := make(chan error, 2)

	// TCP -> WebSocket
	go func() {
		buffer := make([]byte, 32*1024) // 32KB buffer
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				done <- err
				return
			}

			if err := c.conn.WriteMessage(websocket.BinaryMessage, buffer[:n]); err != nil {
				done <- err
				return
			}
		}
	}()

	// WebSocket -> TCP
	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				done <- err
				return
			}

			if _, err := conn.Write(message); err != nil {
				done <- err
				return
			}
		}
	}()

	return <-done
}

// Close 关闭连接
func (c *WebSocketClient) Close() error {
	if c.conn != nil {
		// 发送关闭消息
		return c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}
	return nil
}

// buildWebSocketURL 构建WebSocket URL
func (c *WebSocketClient) buildWebSocketURL(targetHost string, targetPort int) (string, error) {
	// 解析服务器地址
	serverURL, err := url.Parse(c.config.ServerAddr)
	if err != nil {
		return "", fmt.Errorf("invalid server address: %w", err)
	}

	// 构建WebSocket URL (不再包含查询参数)
	wsScheme := "ws"
	if serverURL.Scheme == "https" {
		wsScheme = "wss"
	}

	wsURL := fmt.Sprintf("%s://%s%s",
		wsScheme,
		serverURL.Host,
		serverURL.Path)

	return wsURL, nil
}

// buildHeaders 构建HTTP请求头
func (c *WebSocketClient) buildHeaders(targetHost string, targetPort int) http.Header {
	headers := make(http.Header)

	// 注意：不手动设置标准的WebSocket握手头（Upgrade、Connection、Sec-WebSocket-Version、Sec-WebSocket-Key）
	// 这些header由gorilla/websocket库自动设置，避免重复header错误

	// 直接将用户名和密码添加到HTTP Headers
	if c.config.Username != "" {
		headers.Set("X-Proxy-Username", c.config.Username)
	}
	if c.config.Password != "" {
		headers.Set("X-Proxy-Password", c.config.Password)
	}

	// 将目标主机和端口添加到HTTP Headers
	headers.Set("X-Proxy-Target-Host", targetHost)
	headers.Set("X-Proxy-Target-Port", fmt.Sprintf("%d", targetPort))

	return headers
}

// NetConn 将WebSocket连接包装为net.Conn接口
func (c *WebSocketClient) NetConn() net.Conn {
	if c.conn == nil {
		return nil
	}
	return &websocketConn{conn: c.conn}
}

// websocketConn 将websocket.Conn包装为net.Conn
type websocketConn struct {
	conn       *websocket.Conn
	readBuffer []byte
}

func (w *websocketConn) Read(b []byte) (n int, err error) {
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

func (w *websocketConn) Write(b []byte) (n int, err error) {
	err = w.conn.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (w *websocketConn) Close() error {
	return w.conn.Close()
}

func (w *websocketConn) LocalAddr() net.Addr {
	return w.conn.LocalAddr()
}

func (w *websocketConn) RemoteAddr() net.Addr {
	return w.conn.RemoteAddr()
}

func (w *websocketConn) SetDeadline(t time.Time) error {
	return w.conn.SetReadDeadline(t)
}

func (w *websocketConn) SetReadDeadline(t time.Time) error {
	return w.conn.SetReadDeadline(t)
}

func (w *websocketConn) SetWriteDeadline(t time.Time) error {
	return w.conn.SetWriteDeadline(t)
}
