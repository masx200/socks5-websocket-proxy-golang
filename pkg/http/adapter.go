package http

import (
	"context"
	"net"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

// HttpProxyAdapter 适配HTTP代理客户端到系统接口
type HttpProxyAdapter struct {
	client *HttpProxyClient
}

// NewHttpProxyAdapter 创建HTTP代理适配器
func NewHttpProxyAdapter(config interfaces.ClientConfig) (interfaces.ProxyClient, error) {
	// 转换配置
	httpConfig := interfaces.ClientConfig{
		Username:   config.Username,
		Password:   config.Password,
		ServerAddr: config.ServerAddr,
		Protocol:   config.Protocol,
		Timeout:    config.Timeout,
	}

	// 创建HTTP代理客户端
	client, err := NewHttpProxyClient(httpConfig)
	if err != nil {
		return nil, err
	}

	return &HttpProxyAdapter{
		client: client.(*HttpProxyClient),
	}, nil
}

// Connect 连接到目标主机
func (a *HttpProxyAdapter) Connect(host string, port int) error {
	return a.client.Connect(host, port)
}

// ForwardData 转发数据
func (a *HttpProxyAdapter) ForwardData(conn net.Conn) error {
	return a.client.ForwardData(conn)
}

// Close 关闭连接
func (a *HttpProxyAdapter) Close() error {
	return a.client.Close()
}

// SetConnectionClosedCallback 设置连接关闭回调
func (a *HttpProxyAdapter) SetConnectionClosedCallback(callback func()) error {
	return a.client.SetConnectionClosedCallback(callback)
}

// NetConn 获取底层网络连接
func (a *HttpProxyAdapter) NetConn() net.Conn {
	return a.client.NetConn()
}

// DialContext 实现DialContext方法
func (a *HttpProxyAdapter) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	return a.client.DialContext(ctx, network, addr)
}

// 注册HTTP代理客户端创建函数
func init() {
	interfaces.RegisterClient("http", NewHttpProxyAdapter)
}
