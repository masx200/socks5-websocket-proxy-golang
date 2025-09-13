package upstream

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/masx200/http-proxy-go-server/resolver"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

// SelectionStrategy 选择策略
type SelectionStrategy int

const (
	StrategyRoundRobin SelectionStrategy = iota // 轮询
	StrategyRandom                              // 随机
	StrategyWeighted                            // 加权
	StrategyFailover                            // 故障转移
)

// DynamicUpstreamSelector 动态上游连接选择器
type DynamicUpstreamSelector struct {
	configs      []interfaces.UpstreamConfig
	strategy     SelectionStrategy
	currentIndex int
	mu           sync.RWMutex
	healthChecks map[string]bool
	resolver     resolver.NameResolver
}

// NewDynamicUpstreamSelector 创建动态上游连接选择器
func NewDynamicUpstreamSelector(configs []interfaces.UpstreamConfig, strategy SelectionStrategy) *DynamicUpstreamSelector {
	return NewDynamicUpstreamSelectorWithResolver(configs, strategy, nil)
}

// NewDynamicUpstreamSelectorWithResolver 创建带有解析器的动态上游连接选择器
func NewDynamicUpstreamSelectorWithResolver(configs []interfaces.UpstreamConfig, strategy SelectionStrategy, resolver resolver.NameResolver) *DynamicUpstreamSelector {
	selector := &DynamicUpstreamSelector{
		configs:      configs,
		strategy:     strategy,
		currentIndex: 0,
		healthChecks: make(map[string]bool),
		resolver:     resolver,
	}

	// 初始化健康检查
	for _, config := range configs {
		selector.healthChecks[config.ProxyAddress] = true
	}

	return selector
}

// SelectConnection 动态选择上游连接
func (s *DynamicUpstreamSelector) SelectConnection(targetHost string, targetPort int) (net.Conn, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.configs) == 0 {
		// 默认直连
		return s.createDirectConnection(targetHost, targetPort)
	}

	// 根据策略选择上游配置
	selectedConfig := s.selectConfig()
	if selectedConfig == nil {
		return s.createDirectConnection(targetHost, targetPort)
	}

	// 创建对应类型的连接
	switch selectedConfig.Type {
	case interfaces.UpstreamDirect:
		return s.createDirectConnection(targetHost, targetPort)
	case interfaces.UpstreamSOCKS5:
		return s.createSOCKS5Connection(selectedConfig, targetHost, targetPort)
	case interfaces.UpstreamWebSocket:
		return s.createWebSocketConnection(selectedConfig, targetHost, targetPort)
	case interfaces.UpstreamHTTP:
		return s.createHTTPConnection(selectedConfig, targetHost, targetPort)
	default:
		return s.createDirectConnection(targetHost, targetPort)
	}
}

// selectConfig 根据策略选择配置
func (s *DynamicUpstreamSelector) selectConfig() *interfaces.UpstreamConfig {
	if len(s.configs) == 0 {
		return nil
	}

	// 过滤健康的配置
	healthyConfigs := make([]interfaces.UpstreamConfig, 0)
	for _, config := range s.configs {
		if s.isHealthy(&config) {
			healthyConfigs = append(healthyConfigs, config)
		}
	}

	if len(healthyConfigs) == 0 {
		return nil
	}

	switch s.strategy {
	case StrategyRoundRobin:
		return s.selectRoundRobin(healthyConfigs)
	case StrategyRandom:
		return s.selectRandom(healthyConfigs)
	case StrategyWeighted:
		return s.selectWeighted(healthyConfigs)
	case StrategyFailover:
		return s.selectFailover(healthyConfigs)
	default:
		return &healthyConfigs[0]
	}
}

// selectRoundRobin 轮询选择
func (s *DynamicUpstreamSelector) selectRoundRobin(configs []interfaces.UpstreamConfig) *interfaces.UpstreamConfig {
	config := &configs[s.currentIndex%len(configs)]
	s.currentIndex++
	return config
}

// selectRandom 随机选择
func (s *DynamicUpstreamSelector) selectRandom(configs []interfaces.UpstreamConfig) *interfaces.UpstreamConfig {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(configs))
	return &configs[index]
}

// selectWeighted 加权选择（简化版，基于超时时间作为权重）
func (s *DynamicUpstreamSelector) selectWeighted(configs []interfaces.UpstreamConfig) *interfaces.UpstreamConfig {
	// 计算总权重（超时时间越短权重越高）
	totalWeight := 0
	weights := make([]int, len(configs))
	for i, config := range configs {
		weight := int(30 / config.Timeout.Seconds()) // 默认30秒，超时越短权重越高
		if weight <= 0 {
			weight = 1
		}
		weights[i] = weight
		totalWeight += weight
	}

	// 随机选择
	rand.Seed(time.Now().UnixNano())
	randValue := rand.Intn(totalWeight)

	// 根据权重选择
	currentWeight := 0
	for i, weight := range weights {
		currentWeight += weight
		if randValue < currentWeight {
			return &configs[i]
		}
	}

	return &configs[0]
}

// selectFailover 故障转移选择（总是选择第一个健康的）
func (s *DynamicUpstreamSelector) selectFailover(configs []interfaces.UpstreamConfig) *interfaces.UpstreamConfig {
	return &configs[0]
}

// isHealthy 检查配置是否健康
func (s *DynamicUpstreamSelector) isHealthy(config *interfaces.UpstreamConfig) bool {
	if config == nil {
		return false
	}
	return s.healthChecks[config.ProxyAddress]
}

// UpdateHealthStatus 更新健康状态
func (s *DynamicUpstreamSelector) UpdateHealthStatus(address string, healthy bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.healthChecks[address] = healthy
}

// createDirectConnection 创建TCP直连
func (s *DynamicUpstreamSelector) createDirectConnection(targetHost string, targetPort int) (net.Conn, error) {
	timeout := 30 * time.Second // 默认30秒超时
	portStr := fmt.Sprint(targetPort)

	// 如果目标主机是IP地址，直接连接
	if net.ParseIP(targetHost) != nil {
		return net.DialTimeout("tcp", net.JoinHostPort(targetHost, portStr), timeout)
	}

	// 如果有解析器，优先使用解析器的 LookupIP 方法获取所有IP地址
	if s.resolver != nil {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ips, err := s.resolver.LookupIP(ctx, "tcp", targetHost)
		if err == nil && len(ips) > 0 {
			log.Printf("[UPSTREAM] Resolved '%s' to %d IPs using custom resolver: %v", targetHost, len(ips), ips)

			// 按照解析到的IP地址列表依次尝试连接
			var lastErr error
			for _, ip := range ips {
				targetAddr := net.JoinHostPort(ip.String(), portStr)
				log.Printf("[UPSTREAM] Attempting connection to '%s'", targetAddr)

				conn, err := net.DialTimeout("tcp", targetAddr, timeout)
				if err == nil {
					log.Printf("[UPSTREAM] Successfully connected to '%s'", targetAddr)
					return conn, nil
				}

				log.Printf("[UPSTREAM] Failed to connect to '%s': %v", targetAddr, err)
				lastErr = err
			}

			// 所有IP地址都连接失败
			return nil, fmt.Errorf("all IP addresses failed for '%s', last error: %w", targetHost, lastErr)
		} else {
			log.Printf("[UPSTREAM] Custom resolver LookupIP failed for '%s': %v, falling back to system DNS", targetHost, err)
		}
	}

	// 回退到系统DNS解析
	log.Printf("[UPSTREAM] Using system DNS for '%s'", targetHost)
	return net.DialTimeout("tcp", net.JoinHostPort(targetHost, portStr), timeout)
}

// createSOCKS5Connection 创建SOCKS5代理连接
func (s *DynamicUpstreamSelector) createSOCKS5Connection(config *interfaces.UpstreamConfig, targetHost string, targetPort int) (net.Conn, error) {
	if config.ProxyAddress == "" {
		return nil, fmt.Errorf("SOCKS5 proxy address not configured")
	}

	// 使用工厂函数创建SOCKS5客户端
	client, err := interfaces.CreateClient("socks5", interfaces.ClientConfig{
		Username:   config.ProxyUsername,
		Password:   config.ProxyPassword,
		ServerAddr: config.ProxyAddress,
		Protocol:   "socks5",
		Timeout:    config.Timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create SOCKS5 client: %w", err)
	}

	// 连接到目标
	err = client.Connect(targetHost, targetPort)
	if err != nil {
		s.UpdateHealthStatus(config.ProxyAddress, false)
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
				log.Printf("SOCKS5 forward data error: %v\n", err)
			}
		}
	}()

	// 返回客户端连接
	return &SOCKS5ConnWrapper{client: client, conn: clientConn}, nil
}

// createWebSocketConnection 创建WebSocket代理连接
func (s *DynamicUpstreamSelector) createWebSocketConnection(config *interfaces.UpstreamConfig, targetHost string, targetPort int) (net.Conn, error) {
	if config.ProxyAddress == "" {
		return nil, fmt.Errorf("WebSocket proxy address not configured")
	}

	// 使用工厂函数创建WebSocket客户端
	client, err := interfaces.CreateClient("websocket", interfaces.ClientConfig{
		Username:   config.ProxyUsername,
		Password:   config.ProxyPassword,
		ServerAddr: config.ProxyAddress,
		Protocol:   "websocket",
		Timeout:    config.Timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create WebSocket client: %w", err)
	}

	// 连接到目标
	err = client.Connect(targetHost, targetPort)
	if err != nil {
		s.UpdateHealthStatus(config.ProxyAddress, false)
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
				log.Printf("WebSocket forward data error: %v\n", err)
			}
		}
	}()

	// 返回客户端连接
	return &WebSocketConnWrapper{client: client, conn: clientConn}, nil
}

// createHTTPConnection 创建HTTP代理连接
func (s *DynamicUpstreamSelector) createHTTPConnection(config *interfaces.UpstreamConfig, targetHost string, targetPort int) (net.Conn, error) {
	if config.ProxyAddress == "" {
		return nil, fmt.Errorf("HTTP proxy address not configured")
	}

	// 使用工厂函数创建HTTP客户端
	client, err := interfaces.CreateClient("http", interfaces.ClientConfig{
		Username:   config.ProxyUsername,
		Password:   config.ProxyPassword,
		ServerAddr: config.ProxyAddress,
		Protocol:   "http",
		Timeout:    config.Timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	// 连接到目标
	err = client.Connect(targetHost, targetPort)
	if err != nil {
		s.UpdateHealthStatus(config.ProxyAddress, false)
		return nil, fmt.Errorf("failed to connect via HTTP proxy: %w", err)
	}

	// 创建一个管道来模拟net.Conn接口
	clientConn, serverConn := net.Pipe()

	// 启动goroutine来处理数据转发
	go func() {
		defer clientConn.Close()
		defer serverConn.Close()

		// 使用客户端的ForwardData方法进行数据转发
		if httpClient, ok := client.(interface {
			ForwardData(net.Conn) error
		}); ok {
			err := httpClient.ForwardData(serverConn)
			if err != nil {
				log.Printf("HTTP forward data error: %v\n", err)
			}
		}
	}()

	// 返回客户端连接
	return &HTTPConnWrapper{client: client, conn: clientConn}, nil
}

// HTTPConnWrapper HTTP客户端连接包装器，实现net.Conn接口
type HTTPConnWrapper struct {
	client interfaces.ProxyClient
	conn   net.Conn
	closed bool
}

func (w *HTTPConnWrapper) Read(b []byte) (n int, err error) {
	if w.closed {
		return 0, net.ErrClosed
	}
	if w.conn == nil {
		return 0, fmt.Errorf("connection not established")
	}
	return w.conn.Read(b)
}

func (w *HTTPConnWrapper) Write(b []byte) (n int, err error) {
	if w.closed {
		return 0, net.ErrClosed
	}
	if w.conn == nil {
		return 0, fmt.Errorf("connection not established")
	}
	return w.conn.Write(b)
}

func (w *HTTPConnWrapper) Close() error {
	if w.closed {
		return nil
	}
	w.closed = true
	if w.conn != nil {
		return w.conn.Close()
	}
	return w.client.Close()
}

func (w *HTTPConnWrapper) LocalAddr() net.Addr {
	if w.client.NetConn() != nil {
		return w.client.NetConn().LocalAddr()
	}
	return newTCPAddrWrapper("local")
}

func (w *HTTPConnWrapper) RemoteAddr() net.Addr {
	if w.client.NetConn() != nil {
		return w.client.NetConn().RemoteAddr()
	}
	return newTCPAddrWrapper("remote")
}

func (w *HTTPConnWrapper) SetDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetDeadline(t)
	}
	return nil
}

func (w *HTTPConnWrapper) SetReadDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetReadDeadline(t)
	}
	return nil
}

func (w *HTTPConnWrapper) SetWriteDeadline(t time.Time) error {
	if w.conn != nil {
		return w.conn.SetWriteDeadline(t)
	}
	return nil
}

// 保持原有的连接包装器实现
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
	return newTCPAddrWrapper("local")
}

func (w *SOCKS5ConnWrapper) RemoteAddr() net.Addr {
	if w.conn != nil {
		return w.conn.RemoteAddr()
	}
	return newTCPAddrWrapper("remote")
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
	if w.client.NetConn() != nil {
		return w.client.NetConn().LocalAddr()
	}
	return newTCPAddrWrapper("local")
}

func (w *WebSocketConnWrapper) RemoteAddr() net.Addr {
	if w.client.NetConn() != nil {
		return w.client.NetConn().RemoteAddr()
	}
	return newTCPAddrWrapper("remote")
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

// tcpAddrWrapper TCP地址包装器，用于解决socks5库的类型转换问题
type tcpAddrWrapper struct {
	*net.TCPAddr
}

// newTCPAddrWrapper 创建一个新的TCP地址包装器
func newTCPAddrWrapper(addr string) *tcpAddrWrapper {
	// 解析地址
	if addr == "local" {
		// 返回本地地址
		return &tcpAddrWrapper{
			TCPAddr: &net.TCPAddr{
				IP:   net.ParseIP("127.0.0.1"),
				Port: 0,
				Zone: "",
			},
		}
	}

	if addr == "remote" {
		// 返回远程地址
		return &tcpAddrWrapper{
			TCPAddr: &net.TCPAddr{
				IP:   net.ParseIP("0.0.0.0"),
				Port: 0,
				Zone: "",
			},
		}
	}

	// 尝试解析为TCP地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		// 如果解析失败，返回默认地址
		return &tcpAddrWrapper{
			TCPAddr: &net.TCPAddr{
				IP:   net.ParseIP("127.0.0.1"),
				Port: 0,
				Zone: "",
			},
		}
	}

	return &tcpAddrWrapper{TCPAddr: tcpAddr}
}

// 保持原有的UpstreamSelector以向后兼容
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
	if net.ParseIP(targetHost) != nil && net.ParseIP(targetHost).To4() == nil {
		return net.DialTimeout("tcp", net.JoinHostPort(targetHost, fmt.Sprint(targetPort)), timeout)
	}
	return net.DialTimeout("tcp", net.JoinHostPort(targetHost, fmt.Sprint(targetPort)), timeout)
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
				log.Printf("SOCKS5 forward data error: %v\n", err)
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
				log.Printf("WebSocket forward data error: %v\n", err)
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
