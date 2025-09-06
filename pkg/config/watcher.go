package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
)

// ConfigWatcher 配置文件监听器
type ConfigWatcher struct {
	watcher    *fsnotify.Watcher
	configFile string
	server     interfaces.ProxyServer
	lastMod    time.Time
	debounce   time.Duration
}

// NewConfigWatcher 创建新的配置监听器
func NewConfigWatcher(configFile string, server interfaces.ProxyServer) (*ConfigWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}

	cw := &ConfigWatcher{
		watcher:    watcher,
		configFile: configFile,
		server:     server,
		debounce:   1 * time.Second, // 防抖时间，避免频繁重载
	}

	// 获取初始文件信息
	if info, err := os.Stat(configFile); err == nil {
		cw.lastMod = info.ModTime()
	}

	// 添加文件监听
	absPath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// 监听配置文件所在的目录
	configDir := filepath.Dir(absPath)
	if err := watcher.Add(configDir); err != nil {
		return nil, fmt.Errorf("failed to watch config directory: %w", err)
	}

	return cw, nil
}

// Start 开始监听配置文件变化
func (cw *ConfigWatcher) Start() {
	go cw.watchLoop()
	log.Printf("[CONFIG-WATCHER] Started watching config file: %s", cw.configFile)
}

// Stop 停止监听
func (cw *ConfigWatcher) Stop() {
	if cw.watcher != nil {
		cw.watcher.Close()
		log.Printf("[CONFIG-WATCHER] Stopped watching config file")
	}
}

// watchLoop 监听循环
func (cw *ConfigWatcher) watchLoop() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[CONFIG-WATCHER] Recovered from panic: %v", r)
		}
	}()

	for {
		select {
		case event, ok := <-cw.watcher.Events:
			if !ok {
				return
			}

			// 检查是否是配置文件的变化
			if filepath.Base(event.Name) == filepath.Base(cw.configFile) {
				if event.Op&fsnotify.Write == fsnotify.Write ||
					event.Op&fsnotify.Create == fsnotify.Create ||
					event.Op&fsnotify.Rename == fsnotify.Rename {

					// 防抖处理
					time.Sleep(cw.debounce)

					// 检查文件是否真的发生了变化
					if cw.hasFileChanged() {
						cw.handleConfigChange()
					}
				}
			}

		case err, ok := <-cw.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("[CONFIG-WATCHER] Watch error: %v", err)
		}
	}
}

// hasFileChanged 检查文件是否真的有变化
func (cw *ConfigWatcher) hasFileChanged() bool {
	info, err := os.Stat(cw.configFile)
	if err != nil {
		log.Printf("[CONFIG-WATCHER] Failed to stat config file: %v", err)
		return false
	}

	if info.ModTime().After(cw.lastMod) {
		cw.lastMod = info.ModTime()
		return true
	}

	return false
}

// handleConfigChange 处理配置文件变化
func (cw *ConfigWatcher) handleConfigChange() {
	log.Printf("[CONFIG-WATCHER] Detected config file change, reloading...")

	// 重新加载配置
	config, err := LoadConfig(cw.configFile)
	if err != nil {
		log.Printf("[CONFIG-WATCHER] Failed to load new config: %v", err)
		return
	}

	// 验证配置
	if err := ValidateConfig(config); err != nil {
		log.Printf("[CONFIG-WATCHER] Invalid config: %v", err)
		return
	}

	// 重新加载配置到服务器
	if err := cw.server.ReloadConfig(config); err != nil {
		log.Printf("[CONFIG-WATCHER] Failed to reload config: %v", err)
		return
	}

	log.Printf("[CONFIG-WATCHER] Configuration reloaded successfully")
}

// LoadConfig 加载配置文件
func LoadConfig(configFile string) (interfaces.ServerConfig, error) {
	return interfaces.LoadConfig(configFile)
}

// ValidateConfig 验证配置
func ValidateConfig(config interfaces.ServerConfig) error {
	return interfaces.ValidateConfig(config)
}