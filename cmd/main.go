package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

config_module	"github.com/masx200/socks5-websocket-proxy-golang/pkg/config"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/proxy"
)

func main() {
	// 解析命令行参数
	mode := flag.String("mode", "server", "运行模式: server 或 client")
	protocol := flag.String("protocol", "socks5", "协议类型: socks5 或 websocket")
	addr := flag.String("addr", ":1080", "监听地址(服务端)或服务器地址(客户端)")
	username := flag.String("username", "", "用户名")
	password := flag.String("password", "", "密码")
	timeout := flag.Int("timeout", 30, "超时时间(秒)")
	targetHost := flag.String("host", "", "目标主机(客户端模式)")
	targetPort := flag.Int("port", 0, "目标端口(客户端模式)")
	configFile := flag.String("config", "", "配置文件路径")

	flag.Parse()

	// 设置日志
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 处理信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	switch *mode {
	case "server":
		startServer(*protocol, *addr, *username, *password, time.Duration(*timeout)*time.Second, *configFile, sigChan)
	case "client":
		if *targetHost == "" || *targetPort == 0 {
			log.Fatal("客户端模式需要指定目标主机和端口 (-host 和 -port)")
		}
		startClient(*protocol, *addr, *username, *password, time.Duration(*timeout)*time.Second, *targetHost, *targetPort, sigChan)
	default:
		log.Fatal("不支持的运行模式: ", *mode)
	}
}

// startServer 启动服务端
func startServer(protocol, addr, username, password string, timeout time.Duration, configFile string, sigChan chan os.Signal) {
	log.Printf("启动%s服务端，监听地址: %s\n", protocol, addr)

	// 创建服务端配置
	config := proxy.ServerConfig{
		ListenAddr: addr,
		Timeout:    timeout,
	}

	// 如果有配置文件，从配置文件加载
	if configFile != "" {
		fileConfig, err := loadServerConfig(configFile)
		if err != nil {
			log.Printf("加载配置文件失败: %v，使用命令行配置\n", err)
		} else {
			config = fileConfig
		}
	}

	// 设置认证用户
	if username != "" && password != "" {
		config.AuthUsers = map[string]string{username: password}
	}

	// 创建服务端
	server, err := proxy.CreateServer(protocol, config)
	if err != nil {
		log.Fatal("创建服务端失败:", err)
	}

	// 启动服务端
	if err := server.Listen(); err != nil {
		log.Fatal("启动服务端失败:", err)
	}

	log.Printf("%s服务端已启动，按Ctrl+C停止\n", protocol)

	// 创建配置监听器（如果指定了配置文件）
	var configWatcher *config_module.ConfigWatcher
	if configFile != "" {
		configWatcher, err = config_module.NewConfigWatcher(configFile, server)
		if err != nil {
			log.Printf("创建配置监听器失败: %v，配置热重载功能不可用\n", err)
		} else {
			configWatcher.Start()
			log.Printf("配置热重载已启用，监听文件: %s", configFile)
		}
	}

	// 等待停止信号
	<-sigChan
	log.Println("正在停止服务端...")

	// 停止配置监听器
	if configWatcher != nil {
		configWatcher.Stop()
	}

	// 优雅关闭
	if err := server.Shutdown(); err != nil {
		log.Printf("停止服务端时出错: %v\n", err)
	}

	log.Println("服务端已停止")
}

// startClient 启动客户端
func startClient(protocol, serverAddr, username, password string, timeout time.Duration, targetHost string, targetPort int, sigChan chan os.Signal) {
	log.Printf("启动%s客户端，服务器地址: %s，目标: %s:%d\n", protocol, serverAddr, targetHost, targetPort)

	// 创建客户端配置
	config := proxy.ClientConfig{
		Username:   username,
		Password:   password,
		ServerAddr: serverAddr,
		Protocol:   protocol,
		Timeout:    timeout,
	}

	// 创建客户端
	client, err := proxy.CreateClient(protocol, config)
	if err != nil {
		log.Fatal("创建客户端失败:", err)
	}

	// 连接到目标主机
	if err := client.Connect(targetHost, targetPort); err != nil {
		log.Fatal("连接目标主机失败:", err)
	}

	// 进行认证
	if username != "" && password != "" {
		if err := client.Authenticate(username, password); err != nil {
			log.Fatal("认证失败:", err)
		}
	}

	log.Printf("已连接到目标主机 %s:%d\n", targetHost, targetPort)

	// 创建一个模拟的连接来测试数据转发
	// 在实际应用中，这里应该是一个真实的网络连接
	// 这里只是演示，实际使用时需要根据具体需求调整
	log.Println("客户端已启动，按Ctrl+C停止")

	// 等待停止信号
	<-sigChan
	log.Println("正在停止客户端...")

	// 关闭客户端
	if err := client.Close(); err != nil {
		log.Printf("关闭客户端时出错: %v\n", err)
	}

	log.Println("客户端已停止")
}

// loadServerConfig 从配置文件加载服务端配置
func loadServerConfig(configFile string) (interfaces.ServerConfig, error) {
	// 读取配置文件
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return interfaces.ServerConfig{}, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析JSON配置文件
	var config interfaces.ServerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return interfaces.ServerConfig{}, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 验证必要字段
	if config.ListenAddr == "" {
		return interfaces.ServerConfig{}, fmt.Errorf("配置文件缺少必要字段: listen_addr")
	}

	// 设置默认值
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return config, nil
}
