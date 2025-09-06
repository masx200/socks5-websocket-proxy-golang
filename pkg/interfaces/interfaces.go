package interfaces

import (
	"fmt"
	"net"
	"time"
)

// UpstreamType 上游连接类型
type UpstreamType int

const (
	UpstreamDirect    UpstreamType = iota // TCP直连
	UpstreamSOCKS5                        // SOCKS5代理
	UpstreamWebSocket                     // WebSocket代理
)

// UpstreamConfig 上游连接配置
type UpstreamConfig struct {
	Type          UpstreamType  // 上游连接类型
	ProxyAddress  string        // 代理服务器地址
	ProxyUsername string        // 代理用户名
	ProxyPassword string        // 代理密码
	Timeout       time.Duration // 超时时间
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Username   string        // 用户名
	Password   string        // 密码
	ServerAddr string        // 服务器地址
	Protocol   string        // 协议类型
	Timeout    time.Duration // 超时时间
}

// ServerConfig 服务端配置
type ServerConfig struct {
	ListenAddr     string            // 监听地址
	AuthUsers      map[string]string // 认证用户映射
	Protocol       string            // 协议类型
	Timeout        time.Duration     // 超时时间
	UpstreamConfig []UpstreamConfig  // 上游连接配置列表
	EnableUpstream bool              // 是否启用上游连接
}

// ProxyClient 代理客户端接口
type ProxyClient interface {
	Connect(host string, port int) error
	Authenticate(username, password string) error
	ForwardData(conn net.Conn) error
	Close() error
}

// ProxyServer 代理服务端接口
type ProxyServer interface {
	Listen() error
	HandleConnection(conn net.Conn) error
	Authenticate(username, password string) bool
	SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error)
	Shutdown() error
}

// CreateClientFunc 客户端创建函数类型
type CreateClientFunc func(config ClientConfig) (ProxyClient, error)

// CreateServerFunc 服务端创建函数类型
type CreateServerFunc func(config ServerConfig) (ProxyServer, error)

// 客户端创建函数映射
var clientCreateFuncs = make(map[string]CreateClientFunc)

// 服务端创建函数映射
var serverCreateFuncs = make(map[string]CreateServerFunc)

// RegisterClient 注册客户端创建函数
func RegisterClient(protocol string, createFunc CreateClientFunc) {
	clientCreateFuncs[protocol] = createFunc
}

// RegisterServer 注册服务端创建函数
func RegisterServer(protocol string, createFunc CreateServerFunc) {
	serverCreateFuncs[protocol] = createFunc
}

// CreateClient 创建客户端实例
func CreateClient(protocol string, config ClientConfig) (ProxyClient, error) {
	createFunc, exists := clientCreateFuncs[protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}
	return createFunc(config)
}

// CreateServer 创建服务端实例
func CreateServer(protocol string, config ServerConfig) (ProxyServer, error) {
	createFunc, exists := serverCreateFuncs[protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}
	return createFunc(config)
}
