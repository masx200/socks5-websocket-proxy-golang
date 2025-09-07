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

	config_module "github.com/masx200/socks5-websocket-proxy-golang/pkg/config"
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
	upstreamType := flag.String("upstream-type", "", "上游连接类型: direct, socks5, websocket")
	upstreamAddress := flag.String("upstream-address", "", "上游代理地址")
	upstreamUsername := flag.String("upstream-username", "", "上游代理用户名")
	upstreamPassword := flag.String("upstream-password", "", "上游代理密码")

	flag.Parse()

	// 设置日志
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 处理信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	switch *mode {
	case "server":
		startServer(*protocol, *addr, *username, *password, time.Duration(*timeout)*time.Second, *configFile, *upstreamType, *upstreamAddress, *upstreamUsername, *upstreamPassword, sigChan)
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
func startServer(initialProtocol, addr, username, password string, timeout time.Duration, configFile, upstreamType, upstreamAddress, upstreamUsername, upstreamPassword string, sigChan chan os.Signal) {
	var (
		currentProtocol = initialProtocol
		server          interfaces.ProxyServer
		configWatcher   *config_module.ConfigWatcher
		err             error
	)

	// 创建服务器重启通道
	restartChan := make(chan string, 1)

	// 服务器启动函数
	startServerInstance := func(protocol string) (interfaces.ProxyServer, error) {
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

		// 设置上游配置（如果通过命令行参数指定）
		if upstreamType != "" {
			upstreamConfig := interfaces.UpstreamConfig{
				Type:          interfaces.UpstreamType(upstreamType),
				ProxyAddress:  upstreamAddress,
				ProxyUsername: upstreamUsername,
				ProxyPassword: upstreamPassword,
				Timeout:       timeout,
			}
			config.UpstreamConfig = []interfaces.UpstreamConfig{upstreamConfig}
			config.EnableUpstream = true
			log.Printf("使用命令行参数设置上游配置: 类型=%s, 地址=%s", upstreamType, upstreamAddress)
		}

		// 创建服务端
		return proxy.CreateServer(protocol, config)
	}

	// 协议变更回调函数
	onProtocolChange := func(newProtocol string) {
		log.Printf("[MAIN] 检测到协议变更，准备重启服务器: %s -> %s", currentProtocol, newProtocol)
		restartChan <- newProtocol
	}

	// 启动初始服务器
	server, err = startServerInstance(currentProtocol)
	if err != nil {
		log.Fatal("创建服务端失败:", err)
	}

	// 启动服务端
	if err := server.Listen(); err != nil {
		log.Fatal("启动服务端失败:", err)
	}

	log.Printf("%s服务端已启动，按Ctrl+C停止\n", currentProtocol)

	// 创建配置监听器（如果指定了配置文件）
	if configFile != "" {
		configWatcher, err = config_module.NewConfigWatcherWithCallback(configFile, server, onProtocolChange)
		if err != nil {
			log.Printf("创建配置监听器失败: %v，配置热重载功能不可用\n", err)
		} else {
			configWatcher.Start()
			log.Printf("配置热重载已启用，监听文件: %s", configFile)
		}
	}

	// 主循环，处理信号和协议变更
	for {
		select {
		case <-sigChan:
			log.Println("正在停止服务端...")
			goto shutdown
		case newProtocol := <-restartChan:
			log.Printf("[MAIN] 收到协议变更信号，正在重启服务器: %s -> %s", currentProtocol, newProtocol)

			// 停止当前服务器
			if err := server.Shutdown(); err != nil {
				log.Printf("[MAIN] 停止当前服务器时出错: %v", err)
			}

			// 停止配置监听器
			if configWatcher != nil {
				configWatcher.Stop()
			}

			// 更新协议
			currentProtocol = newProtocol

			// 启动新服务器
			server, err = startServerInstance(currentProtocol)
			if err != nil {
				log.Fatalf("[MAIN] 创建新服务器失败: %v", err)
			}

			// 启动服务端
			if err := server.Listen(); err != nil {
				log.Fatalf("[MAIN] 启动新服务器失败: %v", err)
			}

			log.Printf("[MAIN] %s服务器已启动", currentProtocol)

			// 重新创建配置监听器
			if configFile != "" {
				configWatcher, err = config_module.NewConfigWatcherWithCallback(configFile, server, onProtocolChange)
				if err != nil {
					log.Printf("[MAIN] 重新创建配置监听器失败: %v，配置热重载功能不可用\n", err)
				} else {
					configWatcher.Start()
					log.Printf("[MAIN] 配置热重载已重新启用，监听文件: %s", configFile)
				}
			}
		}
	}

shutdown:
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
		log.Println("正在认证...", username, password)
		// if err := client.Authenticate(username, password); err != nil {
		// 	log.Fatal("认证失败:", err)
		// }
	}

	log.Printf("已连接到目标主机 %s:%d\n", targetHost, targetPort)

	// 创建一个模拟的连接来测试数据转发
	// 在实际应用中，这里应该是一个真实的网络连接
	// 这里只是演示，实际使用时需要根据具体需求调整
	log.Println("客户端已启动，按Ctrl+C停止")

	// 设置连接关闭回调
	connectionClosed := make(chan bool, 1)
	err = client.SetConnectionClosedCallback(func() {
		log.Println("连接已关闭，正在退出程序...")
		connectionClosed <- true
	})
	if err != nil {
		log.Printf("设置连接关闭回调失败: %v\n", err)
	}

	// 等待停止信号或连接关闭
	var connectionWasClosed bool
	select {
	case <-sigChan:
		log.Println("正在停止客户端...")
	case <-connectionClosed:
		log.Println("连接已关闭，正在退出程序...")
		connectionWasClosed = true
	}

	// 关闭客户端
	if err := client.Close(); err != nil {
		log.Printf("关闭客户端时出错: %v\n", err)
	}

	log.Println("客户端已停止")
	// 如果是连接关闭导致的退出，直接退出程序
	if connectionWasClosed {
		os.Exit(1)
	}
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
