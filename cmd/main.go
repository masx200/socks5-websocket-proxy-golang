package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/resolver"
	config_module "github.com/masx200/socks5-websocket-proxy-golang/pkg/config"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/proxy"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/socks5"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

type multiString []string

func (m *multiString) String() string {
	return "[" + strings.Join(*m, ", ") + "]"
}

func (m *multiString) Set(value string) error {
	*m = append(*m, value)
	return nil
}
func main() {

	var (
		dohurls  multiString
		dohips   multiString
		dohalpns multiString
	)
	// 注册可重复参数
	flag.Var(&dohurls, "dohurl", "DOH URL (可重复),支持http协议和https协议")
	flag.Var(&dohips, "dohip", "DOH IP (可重复),支持ipv4地址和ipv6地址")
	flag.Var(&dohalpns, "dohalpn", "DOH alpn (可重复),支持h2协议和h3协议")
	// 解析命令行参数
	mode := flag.String("mode", "server", "运行模式: server")
	protocol := flag.String("protocol", "socks5", "协议类型: socks5 或 websocket")
	addr := flag.String("addr", ":1080", "监听地址(服务端)或服务器地址(客户端)")
	listenHost := flag.String("listen-host", "", "监听主机地址")
	listenPort := flag.String("listen-port", "", "监听端口号")
	username := flag.String("username", "", "用户名")
	password := flag.String("password", "", "密码")
	timeout := flag.Int("timeout", 30, "超时时间(秒)")
	configFile := flag.String("config", "", "配置文件路径")
	upstreamType := flag.String("upstream-type", "", "上游连接类型: direct, socks5, websocket, http")
	upstreamAddress := flag.String("upstream-address", "", "上游代理地址")
	upstreamUsername := flag.String("upstream-username", "", "上游代理用户名")
	upstreamPassword := flag.String("upstream-password", "", "上游代理密码")

	flag.Parse()

	var proxyoptions = options.ProxyOptions{}
	for i, dohurl := range dohurls {

		var dohip string
		if len(dohips) > i {
			dohip = dohips[i]
		} else {
			dohip = ""
		}
		var dohalpn string
		if len(dohalpns) > i {
			dohalpn = dohalpns[i]
		} else {
			dohalpn = ""
		}

		proxyoptions = append(proxyoptions, options.ProxyOption{Dohurl: dohurl, Dohip: dohip, Dohalpn: dohalpn})
	}
	// 处理 listen-host 和 listen-port 参数
	if *listenHost != "" || *listenPort != "" {
		// 如果同时提供了 listen-host 和 listen-port
		if *listenHost != "" && *listenPort != "" {
			*addr = *listenHost + ":" + *listenPort
		} else if *listenHost != "" {
			// 只提供了 listen-host，需要从现有 addr 中提取端口
			if strings.HasPrefix(*addr, ":") {
				*addr = *listenHost + *addr
			} else {
				// 如果 addr 不是以 : 开头，尝试提取端口
				if lastColon := strings.LastIndex(*addr, ":"); lastColon != -1 {
					port := (*addr)[lastColon:]
					*addr = *listenHost + port
				} else {
					// 如果没有找到端口，使用默认端口
					*addr = *listenHost + ":1080"
				}
			}
		} else if *listenPort != "" {
			// 只提供了 listen-port，需要从现有 addr 中提取主机
			if strings.HasPrefix(*addr, ":") {
				// 如果 addr 以 : 开头，说明没有指定主机，使用默认主机
				*addr = ":" + *listenPort
			} else {
				// 如果 addr 不是以 : 开头，尝试提取主机
				if lastColon := strings.LastIndex(*addr, ":"); lastColon != -1 {
					host := (*addr)[:lastColon]
					*addr = host + ":" + *listenPort
				} else {
					// 如果没有找到端口，直接添加端口
					*addr = *addr + ":" + *listenPort
				}
			}
		}
		log.Printf("使用监听地址: %s (由 listen-host=%s, listen-port=%s 合并而成)", *addr, *listenHost, *listenPort)
	}

	// 设置日志
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 处理信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	switch *mode {
	case "server":
		startServer(*protocol, *addr, *username, *password, time.Duration(*timeout)*time.Second, *configFile, *upstreamType, *upstreamAddress, *upstreamUsername, *upstreamPassword, proxyoptions, sigChan)
	default:
		log.Fatal("不支持的运行模式: ", *mode)
	}
}

// startServer 启动服务端
func startServer(initialProtocol, addr, username, password string, timeout time.Duration, configFile, upstreamType, upstreamAddress, upstreamUsername, upstreamPassword string, proxyoptions options.ProxyOptions, sigChan chan os.Signal) {
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

		// 如果有DOH参数，创建解析器并使用带有解析器的服务器
		if len(proxyoptions) > 0 {
			log.Printf("检测到DOH配置参数，创建自定义DNS解析器")
			dohResolver := resolver.CreateHostsAndDohResolver(proxyoptions, nil)

			switch strings.ToLower(protocol) {
			case "socks5":
				return socks5.NewSOCKS5ServerWithResolver(config, dohResolver), nil
			case "websocket":
				return websocket.NewWebSocketServerWithResolver(config, dohResolver), nil
			default:
				return proxy.CreateServer(protocol, config)
			}
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
