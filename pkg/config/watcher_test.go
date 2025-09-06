package config

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

// Connection 接口定义
type Connection interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

// MockServer 用于测试的模拟服务器
type MockServer struct {
	config      interfaces.ServerConfig
	reloadCount int
}

func (m *MockServer) Listen() error {
	return nil
}

func (m *MockServer) HandleConnection(conn net.Conn) error {
	return nil
}

func (m *MockServer) Authenticate(username, password string) bool {
	return true
}

func (m *MockServer) SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error) {
	return nil, nil
}

func (m *MockServer) Shutdown() error {
	return nil
}

func (m *MockServer) ReloadConfig(config interfaces.ServerConfig) error {
	m.config = config
	m.reloadCount++
	return nil
}

func TestConfigWatcher(t *testing.T) {
	// 创建临时目录和配置文件
	tempDir, err := ioutil.TempDir("", "config-test")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configFile := filepath.Join(tempDir, "test-config.json")
	
	// 创建初始配置
	initialConfig := interfaces.ServerConfig{
		ListenAddr: ":1080",
		AuthUsers: map[string]string{
			"user1": "pass1",
		},
		EnableUpstream: false,
	}
	
	configData, _ := json.MarshalIndent(initialConfig, "", "  ")
	if err := ioutil.WriteFile(configFile, configData, 0644); err != nil {
		t.Fatalf("写入配置文件失败: %v", err)
	}

	// 创建模拟服务器
	mockServer := &MockServer{}
	
	// 创建配置监听器
	watcher, err := NewConfigWatcher(configFile, mockServer)
	if err != nil {
		t.Fatalf("创建配置监听器失败: %v", err)
	}

	// 开始监听
	watcher.Start()
	defer watcher.Stop()

	// 等待监听器启动
	time.Sleep(100 * time.Millisecond)

	// 修改配置文件
	updatedConfig := interfaces.ServerConfig{
		ListenAddr: ":1081",
		AuthUsers: map[string]string{
			"user1": "pass1",
			"user2": "pass2",
		},
		EnableUpstream: true,
	}
	
	updatedData, _ := json.MarshalIndent(updatedConfig, "", "  ")
	if err := ioutil.WriteFile(configFile, updatedData, 0644); err != nil {
		t.Fatalf("更新配置文件失败: %v", err)
	}

	// 等待配置重载
	time.Sleep(2 * time.Second)

	// 验证配置是否已更新
	if mockServer.reloadCount < 1 {
		t.Errorf("配置重载失败，reloadCount = %d", mockServer.reloadCount)
	}

	if mockServer.config.ListenAddr != ":1081" {
		t.Errorf("配置未正确更新，期望 :1081，实际 %s", mockServer.config.ListenAddr)
	}

	if len(mockServer.config.AuthUsers) != 2 {
		t.Errorf("用户配置未正确更新，期望 2 个用户，实际 %d 个", len(mockServer.config.AuthUsers))
	}

	if !mockServer.config.EnableUpstream {
		t.Errorf("上游配置未正确更新，期望 true，实际 %t", mockServer.config.EnableUpstream)
	}
}

func TestHasFileChanged(t *testing.T) {
	// 创建临时文件
	tempFile, err := ioutil.TempFile("", "test-file")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// 获取初始修改时间
	info, _ := os.Stat(tempFile.Name())
	initialModTime := info.ModTime()

	// 创建配置监听器
	watcher := &ConfigWatcher{
		configFile: tempFile.Name(),
		lastMod:    initialModTime,
	}

	// 修改文件
	time.Sleep(100 * time.Millisecond) // 确保修改时间不同
	if _, err := tempFile.WriteString("test data"); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}
	tempFile.Close()

	// 检查文件是否变化
	if !watcher.hasFileChanged() {
		t.Error("文件变化检测失败")
	}
}

func TestValidateConfig(t *testing.T) {
	// 测试有效配置
	validConfig := interfaces.ServerConfig{
		ListenAddr: ":1080",
		Timeout:    30 * time.Second,
	}
	
	if err := ValidateConfig(validConfig); err != nil {
		t.Errorf("有效配置验证失败: %v", err)
	}
	
	// 测试无效配置
	invalidConfig := interfaces.ServerConfig{
		ListenAddr: "",
		Timeout:    30 * time.Second,
	}
	
	if err := ValidateConfig(invalidConfig); err == nil {
		t.Error("无效配置应该返回错误")
	}
}