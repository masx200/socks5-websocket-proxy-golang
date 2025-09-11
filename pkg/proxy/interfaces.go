package proxy

import (
	"fmt"
	"strings"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/http"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/socks5"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

// 重新导出接口类型以保持向后兼容性
type (
	UpstreamType   = interfaces.UpstreamType
	UpstreamConfig = interfaces.UpstreamConfig
	ClientConfig   = interfaces.ClientConfig
	ServerConfig   = interfaces.ServerConfig
	ProxyClient    = interfaces.ProxyClient
	ProxyServer    = interfaces.ProxyServer
)

// 重新导出常量
const (
	UpstreamDirect    = interfaces.UpstreamDirect
	UpstreamWebSocket = interfaces.UpstreamWebSocket
	UpstreamSOCKS5    = interfaces.UpstreamSOCKS5
	UpstreamHTTP      = interfaces.UpstreamHTTP
)

// CreateClient 客户端创建工厂
func CreateClient(protocol string, config ClientConfig) (ProxyClient, error) {
	switch strings.ToLower(protocol) {
	case "socks5":
		return socks5.NewSOCKS5Client(config), nil
	case "websocket":
		return websocket.NewWebSocketClient(config), nil
	case "http":
		return http.NewHttpProxyAdapter(config)
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}
}

// CreateServer 服务端创建工厂
func CreateServer(protocol string, config ServerConfig) (ProxyServer, error) {
	switch strings.ToLower(protocol) {
	case "socks5":
		return socks5.NewSOCKS5Server(config), nil
	case "websocket":
		return websocket.NewWebSocketServer(config), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}
}
