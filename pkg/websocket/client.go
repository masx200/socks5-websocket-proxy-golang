package websocket

import (
	"encoding/base64"
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
	interfaces.RegisterClient("websocket", func(config interfaces.ClientConfig) interfaces.ProxyClient {
		return NewWebSocketClient(config)
	})
}

// WebSocketClient WebSocket客户端实现
type WebSocketClient struct {
	config        interfaces.ClientConfig
	conn          *websocket.Conn
	httpClient    *http.Client
	authenticated bool
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
	authenticated := true
	return nil
}

// ForwardData 转发数据
func (c *WebSocketClient) ForwardData(conn net.Conn) error {
	if c.conn == nil {
		return errors.New("connection not established")
	}

	if !authenticated {
		return errors.New("not authenticated")
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

	// 构建查询参数
	query := url.Values{}
	query.Set("host", targetHost)
	query.Set("port", fmt.Sprintf("%d", targetPort))

	// 如果有用户名密码，添加到查询参数
	if c.config.Username != "" {
		query.Set("username", c.config.Username)
	}
	if c.config.Password != "" {
		query.Set("password", c.config.Password)
	}

	// 构建WebSocket URL
	wsScheme := "ws"
	if serverURL.Scheme == "https" {
		wsScheme = "wss"
	}

	wsURL := fmt.Sprintf("%s://%s%s?%s",
		wsScheme,
		serverURL.Host,
		serverURL.Path,
		query.Encode())

	return wsURL, nil
}

// buildHeaders 构建HTTP请求头
func (c *WebSocketClient) buildHeaders(targetHost string, targetPort int) http.Header {
	headers := make(http.Header)

	// 设置基本的WebSocket头
	headers.Set("Upgrade", "websocket")
	headers.Set("Connection", "Upgrade")
	headers.Set("Sec-WebSocket-Version", "13")
	headers.Set("Sec-WebSocket-Key", generateWebSocketKey())

	// 将认证信息序列化并添加到HTTP Headers
	if c.config.Username != "" || c.config.Password != "" {
		authInfo := c.serializeAuthInfo(targetHost, targetPort)
		headers.Set("X-Proxy-Auth", authInfo)
	}

	return headers
}

// serializeAuthInfo 序列化认证信息
func (c *WebSocketClient) serializeAuthInfo(targetHost string, targetPort int) string {
	// 构建认证信息字符串
	authInfo := fmt.Sprintf("username=%s&password=%s&host=%s&port=%d",
		url.QueryEscape(c.config.Username),
		url.QueryEscape(c.config.Password),
		url.QueryEscape(targetHost),
		targetPort)

	// Base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(authInfo))
	return encoded
}

// generateWebSocketKey 生成WebSocket握手密钥
func generateWebSocketKey() string {
	// 生成16字节的随机密钥
	key := make([]byte, 16)
	// 在实际应用中应该使用crypto/rand，这里简化处理
	for i := range key {
		key[i] = byte(i % 256)
	}
	return base64.StdEncoding.EncodeToString(key)
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
