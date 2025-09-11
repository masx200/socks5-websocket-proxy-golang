package upstream

import (
	"testing"
	"time"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

// createTestUpstreamConfig 创建测试用的上游配置
func createTestUpstreamConfig(addr string, timeout time.Duration) interfaces.UpstreamConfig {
	return interfaces.UpstreamConfig{
		Type:         interfaces.UpstreamSOCKS5,
		ProxyAddress: addr,
		Timeout:      timeout,
	}
}

// TestNewDynamicUpstreamSelector 测试动态上游选择器的创建
func TestNewDynamicUpstreamSelector(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)
	if selector == nil {
		t.Fatal("Expected selector to be created, got nil")
	}

	if len(selector.configs) != 2 {
		t.Errorf("Expected 2 configs, got %d", len(selector.configs))
	}

	if selector.strategy != StrategyRoundRobin {
		t.Errorf("Expected StrategyRoundRobin, got %v", selector.strategy)
	}

	if len(selector.healthChecks) != 2 {
		t.Errorf("Expected 2 health checks, got %d", len(selector.healthChecks))
	}
}

// TestSelectRoundRobin 测试轮询选择策略
func TestSelectRoundRobin(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)

	// 测试轮询顺序
	selected := make([]string, 0, 6)
	for i := 0; i < 6; i++ {
		config := selector.selectRoundRobin(configs)
		selected = append(selected, config.ProxyAddress)
	}

	expected := []string{
		"127.0.0.1:1080",
		"127.0.0.1:1081",
		"127.0.0.1:1082",
		"127.0.0.1:1080",
		"127.0.0.1:1081",
		"127.0.0.1:1082",
	}

	for i, addr := range selected {
		if addr != expected[i] {
			t.Errorf("Expected %s at position %d, got %s", expected[i], i, addr)
		}
	}
}

// TestSelectRandom 测试随机选择策略
func TestSelectRandom(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRandom)

	// 测试随机选择多次，确保每次都能返回有效的配置
	for i := 0; i < 100; i++ {
		config := selector.selectRandom(configs)
		if config == nil {
			t.Fatalf("Expected non-nil config at iteration %d", i)
		}

		// 验证返回的配置在输入列表中
		found := false
		for _, expectedConfig := range configs {
			if config.ProxyAddress == expectedConfig.ProxyAddress {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Returned config %s not in input list", config.ProxyAddress)
		}
	}
}

// TestSelectWeighted 测试加权选择策略
func TestSelectWeighted(t *testing.T) {
	// 创建不同超时时间的配置，超时时间越短权重越高
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 10*time.Second), // 高权重
		createTestUpstreamConfig("127.0.0.1:1081", 20*time.Second), // 中权重
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second), // 低权重
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyWeighted)

	// 统计每个配置被选中的次数
	counts := make(map[string]int)
	iterations := 1000

	for i := 0; i < iterations; i++ {
		config := selector.selectWeighted(configs)
		counts[config.ProxyAddress]++
	}

	// 验证高权重配置被选中的次数更多
	// 127.0.0.1:1080 (10秒超时) 应该被选中最多
	// 127.0.0.1:1082 (30秒超时) 应该被选中最少
	if counts["127.0.0.1:1080"] <= counts["127.0.0.1:1082"] {
		t.Errorf("Expected 127.0.0.1:1080 (high weight) to be selected more than 127.0.0.1:1082 (low weight)")
		t.Errorf("Counts: 1080=%d, 1081=%d, 1082=%d", counts["127.0.0.1:1080"], counts["127.0.0.1:1081"], counts["127.0.0.1:1082"])
	}

	// 验证所有配置都被选中过
	for _, config := range configs {
		if counts[config.ProxyAddress] == 0 {
			t.Errorf("Config %s was never selected", config.ProxyAddress)
		}
	}
}

// TestSelectFailover 测试故障转移选择策略
func TestSelectFailover(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyFailover)

	// 故障转移策略应该总是返回第一个健康的配置
	for i := 0; i < 10; i++ {
		config := selector.selectFailover(configs)
		if config.ProxyAddress != "127.0.0.1:1080" {
			t.Errorf("Expected 127.0.0.1:1080 at iteration %d, got %s", i, config.ProxyAddress)
		}
	}
}

// TestSelectConfig 测试配置选择逻辑
func TestSelectConfig(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
	}

	// 测试空配置
	emptySelector := NewDynamicUpstreamSelector([]interfaces.UpstreamConfig{}, StrategyRoundRobin)
	config := emptySelector.selectConfig()
	if config != nil {
		t.Error("Expected nil config for empty configs, got non-nil")
	}

	// 测试健康检查
	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)

	// 设置第一个配置为不健康
	selector.UpdateHealthStatus("127.0.0.1:1080", false)

	// 应该只返回健康的配置
	config = selector.selectConfig()
	if config == nil {
		t.Error("Expected non-nil config when some configs are healthy")
	} else if config.ProxyAddress != "127.0.0.1:1081" {
		t.Errorf("Expected 127.0.0.1:1081, got %s", config.ProxyAddress)
	}

	// 设置所有配置为不健康
	selector.UpdateHealthStatus("127.0.0.1:1081", false)

	config = selector.selectConfig()
	if config != nil {
		t.Error("Expected nil config when all configs are unhealthy")
	}
}

// TestUpdateHealthStatus 测试健康状态更新
func TestUpdateHealthStatus(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)

	// 初始状态应该是健康的
	if !selector.isHealthy(&configs[0]) {
		t.Error("Expected config to be healthy initially")
	}

	// 更新为不健康
	selector.UpdateHealthStatus("127.0.0.1:1080", false)
	if selector.isHealthy(&configs[0]) {
		t.Error("Expected config to be unhealthy after update")
	}

	// 更新为健康
	selector.UpdateHealthStatus("127.0.0.1:1080", true)
	if !selector.isHealthy(&configs[0]) {
		t.Error("Expected config to be healthy after update")
	}
}

// TestCreateDirectConnection 测试直连创建
func TestCreateDirectConnection(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)

	// 测试连接到一个不存在的地址，应该返回错误
	conn, err := selector.createDirectConnection("127.0.0.1", 12345)
	if err == nil {
		conn.Close()
		t.Error("Expected error when connecting to non-existent address")
	}
}

// TestSelectConnection 测试连接选择
func TestSelectConnection(t *testing.T) {
	// 测试空配置情况
	emptySelector := NewDynamicUpstreamSelector([]interfaces.UpstreamConfig{}, StrategyRoundRobin)
	conn, err := emptySelector.SelectConnection("127.0.0.1", 80)
	if err == nil {
		conn.Close()
		t.Error("Expected error when no configs available")
	}

	// 测试直连模式
	directConfig := []interfaces.UpstreamConfig{
		{
			Type:    interfaces.UpstreamDirect,
			Timeout: 30 * time.Second,
		},
	}
	directSelector := NewDynamicUpstreamSelector(directConfig, StrategyRoundRobin)
	conn, err = directSelector.SelectConnection("127.0.0.1", 80)
	if err == nil {
		conn.Close()
		t.Error("Expected error when connecting to non-existent address")
	}
}

// TestStrategySelection 测试策略选择
func TestStrategySelection(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
	}

	// 测试轮询策略
	roundRobinSelector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)
	if roundRobinSelector.strategy != StrategyRoundRobin {
		t.Error("Expected StrategyRoundRobin")
	}

	// 测试随机策略
	randomSelector := NewDynamicUpstreamSelector(configs, StrategyRandom)
	if randomSelector.strategy != StrategyRandom {
		t.Error("Expected StrategyRandom")
	}

	// 测试加权策略
	weightedSelector := NewDynamicUpstreamSelector(configs, StrategyWeighted)
	if weightedSelector.strategy != StrategyWeighted {
		t.Error("Expected StrategyWeighted")
	}

	// 测试故障转移策略
	failoverSelector := NewDynamicUpstreamSelector(configs, StrategyFailover)
	if failoverSelector.strategy != StrategyFailover {
		t.Error("Expected StrategyFailover")
	}
}

// TestConcurrentAccess 测试并发访问
func TestConcurrentAccess(t *testing.T) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)

	// 并发测试
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				// 简化操作，避免复杂的锁竞争
				if j%2 == 0 {
					selector.UpdateHealthStatus("127.0.0.1:1080", id%2 == 0)
				} else {
					selector.UpdateHealthStatus("127.0.0.1:1081", id%2 == 1)
				}
			}
			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 验证选择器仍然正常工作
	config := selector.selectConfig()
	if config == nil {
		t.Error("Selector should still work after concurrent access")
	}
}

// BenchmarkSelectRoundRobin 基准测试：轮询选择
func BenchmarkSelectRoundRobin(b *testing.B) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRoundRobin)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selector.selectRoundRobin(configs)
	}
}

// BenchmarkSelectRandom 基准测试：随机选择
func BenchmarkSelectRandom(b *testing.B) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyRandom)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selector.selectRandom(configs)
	}
}

// BenchmarkSelectWeighted 基准测试：加权选择
func BenchmarkSelectWeighted(b *testing.B) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyWeighted)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selector.selectWeighted(configs)
	}
}

// BenchmarkSelectFailover 基准测试：故障转移选择
func BenchmarkSelectFailover(b *testing.B) {
	configs := []interfaces.UpstreamConfig{
		createTestUpstreamConfig("127.0.0.1:1080", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1081", 30*time.Second),
		createTestUpstreamConfig("127.0.0.1:1082", 30*time.Second),
	}

	selector := NewDynamicUpstreamSelector(configs, StrategyFailover)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selector.selectFailover(configs)
	}
}
