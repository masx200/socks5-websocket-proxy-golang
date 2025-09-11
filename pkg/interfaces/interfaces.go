package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"

	// "os"
	"time"
)

// UpstreamType 上游连接类型
type UpstreamType string

const (
	UpstreamDirect    UpstreamType = "direct"    // TCP直连
	UpstreamSOCKS5    UpstreamType = "socks5"    // SOCKS5代理
	UpstreamWebSocket UpstreamType = "websocket" // WebSocket代理
	UpstreamHTTP      UpstreamType = "http"      // HTTP代理
)

// UpstreamConfig 上游连接配置
type UpstreamConfig struct {
	Type          UpstreamType  `json:"type"`           // 上游连接类型
	ProxyAddress  string        `json:"proxy_address"`  // 代理服务器地址
	ProxyUsername string        `json:"proxy_username"` // 代理用户名
	ProxyPassword string        `json:"proxy_password"` // 代理密码
	Timeout       time.Duration `json:"timeout"`        // 超时时间
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Username   string        `json:"username"`    // 用户名
	Password   string        `json:"password"`    // 密码
	ServerAddr string        `json:"server_addr"` // 服务器地址
	Protocol   string        `json:"protocol"`    // 协议类型
	Timeout    time.Duration `json:"timeout"`     // 超时时间
}

// ServerConfig 服务端配置
type ServerConfig struct {
	ListenAddr     string            `json:"listen_addr"`     // 监听地址
	AuthUsers      map[string]string `json:"auth_users"`      // 认证用户映射
	Protocol       string            `json:"protocol"`        // 协议类型
	Timeout        time.Duration     `json:"timeout"`         // 超时时间
	UpstreamConfig []UpstreamConfig  `json:"upstream_config"` // 上游连接配置列表
	EnableUpstream bool              `json:"enable_upstream"` // 是否启用上游连接
}

// ProxyClient 代理客户端接口
type ProxyClient interface {
	Connect(host string, port int) error
	// Authenticate(username, password string) error
	ForwardData(conn net.Conn) error
	Close() error
	SetConnectionClosedCallback(callback func()) error
	NetConn() net.Conn

	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}

// ProxyServer 代理服务端接口
type ProxyServer interface {
	Listen() error
	HandleConnection(conn net.Conn) error
	Authenticate(username, password string) bool
	SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error)
	Shutdown() error
	ReloadConfig(config ServerConfig) error
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

// LoadConfig 从文件加载配置
func LoadConfig(configFile string) (ServerConfig, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return ServerConfig{}, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config ServerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return ServerConfig{}, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return config, nil
}

// ValidateConfig 验证配置
func ValidateConfig(config ServerConfig) error {
	if config.ListenAddr == "" {
		return fmt.Errorf("监听地址不能为空")
	}

	if config.Timeout <= 0 {
		return fmt.Errorf("超时时间必须大于0")
	}

	// 验证上游配置
	if config.EnableUpstream {
		for i, upstream := range config.UpstreamConfig {
			if upstream.Type == "" {
				return fmt.Errorf("上游配置[%d]: 类型不能为空", i)
			}
			if upstream.ProxyAddress == "" && upstream.Type != "direct" {
				return fmt.Errorf("上游配置[%d]: 代理地址不能为空", i)
			}
			if upstream.Timeout <= 0 {
				return fmt.Errorf("上游配置[%d]: 超时时间必须大于0", i)
			}
		}
	}

	return nil
}
