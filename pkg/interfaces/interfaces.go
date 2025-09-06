package interfaces

import (
	"net"
	"time"
)

// UpstreamType 定义上游连接类型
type UpstreamType string

const (
	UpstreamDirect    UpstreamType = "direct"    // TCP直连
	UpstreamWebSocket UpstreamType = "websocket" // WebSocket代理
	UpstreamSOCKS5    UpstreamType = "socks5"    // SOCKS5代理
)

// UpstreamConfig 上游连接配置
type UpstreamConfig struct {
	Type          UpstreamType  // 连接类型
	ProxyAddress  string        // 代理地址（当Type为websocket或socks5时需要）
	ProxyUsername string        // 代理用户名（可选）
	ProxyPassword string        // 代理密码（可选）
	Timeout       time.Duration // 连接超时时间
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Username   string
	Password   string
	ServerAddr string
	Protocol   string // "socks5" or "websocket"
	Timeout    time.Duration
}

// ServerConfig 服务端配置
type ServerConfig struct {
	ListenAddr     string            // 监听地址
	AuthUsers      map[string]string // username:password
	Protocol       string            // "socks5" or "websocket"
	Timeout        time.Duration     // 连接超时时间
	UpstreamConfig *UpstreamConfig   // 上游连接配置（可选，未配置时使用直连）
	EnableUpstream bool              // 是否启用上游连接选择器
}

// ProxyClient 统一客户端接口
type ProxyClient interface {
	Connect(targetHost string, targetPort int) error
	Authenticate(username, password string) error
	ForwardData(conn net.Conn) error
	Close() error
}

// ProxyServer 统一服务端接口
type ProxyServer interface {
	Listen(address string) error
	HandleConnection(conn net.Conn) error
	Authenticate(username, password string) bool
	// 上游连接选择器：根据配置选择连接上游TCP连接的方式
	SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error)
	Shutdown() error
}

// CreateClient 客户端创建工厂函数
type CreateClientFunc func(protocol string, config ClientConfig) (ProxyClient, error)

// CreateServer 服务端创建工厂函数
type CreateServerFunc func(protocol string, config ServerConfig) (ProxyServer, error)
