# SOCKS5-WebSocket 代理系统

一个使用 Go 语言开发的高性能网络代理系统，支持 SOCKS5 和 WebSocket
协议间的灵活转发和统一接口封装。

## 项目概述

本项目旨在提供一个高性能、可扩展的代理解决方案，支持多种上游连接方式和动态配置。系统采用模块化设计，支持协议间的灵活转发，适用于各种网络代理场景。

## 功能特性

### 🚀 核心功能

- **多协议支持**: 支持 SOCKS5、WebSocket 和 HTTP 协议
- **协议转换**: 支持不同协议间的数据转发
- **上游连接选择**: 支持多种上游连接方式和动态选择策略
- **认证机制**: 支持用户名密码认证
- **高并发**: 支持多客户端并发连接

### 🔧 技术特性

- **模块化设计**: 清晰的架构分层，易于扩展
- **配置管理**: 支持 JSON 配置文件和命令行参数
- **健康检查**: 内置连接健康检查机制
- **故障转移**: 支持自动故障转移和重连
- **IPv6 支持**: 完整支持 IPv4、IPv6 和域名解析

## 系统架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   SOCKS5 Client │    │ WebSocket Client│    │   Direct Client │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          └──────────────────────┼──────────────────────┘
                                 │
                    ┌─────────────▼─────────────┐
                    │   Upstream Selector       │
                    │  (动态选择上游连接方式)     │
                    └─────────────┬─────────────┘
                                 │
          ┌──────────────────────┼──────────────────────┐
          │                      │                      │
┌─────────▼───────┐    ┌─────────▼───────┐    ┌─────────▼───────┐
│   SOCKS5 Server │    │ WebSocket Server│    │   Direct Server │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 安装说明

### 环境要求

- Go 1.19+
- 操作系统: Windows/Linux/macOS

### 编译安装

```bash
# 克隆项目
git clone https://github.com/masx200/socks5-websocket-proxy-golang.git
cd socks5-websocket-proxy-golang

# 下载依赖
go mod tidy

# 编译
go build -o proxy-server.exe cmd/main.go

# 运行
./proxy-server.exe
```

## 快速开始

### 1. 服务端模式

#### 使用命令行参数启动

```bash
# 启动 SOCKS5 服务端
./proxy-server.exe -mode server -protocol socks5 -addr :1080 -username admin -password password123

# 启动 WebSocket 服务端
./proxy-server.exe -mode server -protocol websocket -addr :8080 -username admin -password password123
```

```bash
./proxy-server.exe -mode server -protocol socks5 -addr :1080 -upstream-type socks5 -upstream-address tcp://127.0.0.1:1081 -upstream-username user -upstream-password pass
```

#### 使用配置文件启动

```bash
./proxy-server.exe -mode server -config config/server-config.json
```

**开发规范**: 请查看 [项目开发规范](提示词.md)
了解项目的架构设计、技术实现规范和核心功能模块说明。

**任务进度**: 请查看 [任务清单](TASK_LIST.md)
了解项目的开发进度、已完成功能和待完成任务。

## 命令行参数说明

程序支持以下命令行参数，可以通过 `-h` 或 `--help` 查看帮助信息：

| 参数                 | 类型   | 默认值     | 说明                                                                              |
| -------------------- | ------ | ---------- | --------------------------------------------------------------------------------- |
| `-addr`              | string | `":1080"`  | 监听地址                                                                          |
| `-listen-host`       | string | `""`       | 监听主机地址，与 listen-port 合并后等同于 addr 参数                               |
| `-listen-port`       | string | `""`       | 监听端口号，与 listen-host 合并后等同于 addr 参数                                 |
| `-config`            | string | `""`       | 配置文件路径                                                                      |
| `-mode`              | string | `"server"` | 运行模式: server                                                                  |
| `-password`          | string | `""`       | 密码                                                                              |
| `-protocol`          | string | `"socks5"` | 协议类型: socks5 或 websocket                                                     |
| `-timeout`           | int    | `30`       | 超时时间(秒)                                                                      |
| `-upstream-address`  | string | `""`       | 上游代理地址，支持协议前缀: ws://, socks5://, http://, tcp://, tls://, socks5s:// |
| `-upstream-password` | string | `""`       | 上游代理密码                                                                      |
| `-upstream-type`     | string | `""`       | 上游连接类型: direct, socks5, websocket, http                                     |
| `-upstream-username` | string | `""`       | 上游代理用户名                                                                    |
| `-username`          | string | `""`       | 用户名                                                                            |

### 参数使用示例

#### 服务端模式示例

```bash
# 基本SOCKS5服务端
./proxy-server.exe -mode server -protocol socks5 -addr :1080 -username admin -password password123

# 使用 listen-host 和 listen-port 参数
./proxy-server.exe -mode server -protocol socks5 -listen-host 0.0.0.0 -listen-port 1080 -username admin -password password123

# 只使用 listen-host 参数（使用默认端口）
./proxy-server.exe -mode server -protocol socks5 -listen-host 192.168.1.100 -username admin -password password123

# 只使用 listen-port 参数（使用默认主机）
./proxy-server.exe -mode server -protocol socks5 -listen-port 8080 -username admin -password password123

# 带上游代理的SOCKS5服务端
./proxy-server.exe -mode server -protocol socks5 -addr :1080 -upstream-type socks5 -upstream-address tcp://127.0.0.1:1081 -upstream-username user -upstream-password pass

# WebSocket服务端
./proxy-server.exe -mode server -protocol websocket -addr :8080 -username admin -password password123

# 使用配置文件
./proxy-server.exe -mode server -config config/server-config.json
```

## 配置说明

### 配置文件格式

配置文件使用 JSON 格式，支持以下配置项：

#### 基本配置

- `listen_addr`: 监听地址，格式为 `:端口` 或 `IP:端口`
- `protocol`: 协议类型，支持 `socks5` 或 `websocket`
- `timeout`: 超时时间，单位为纳秒（30 秒 = 30000000000）

#### 认证配置

- `auth_users`: 用户认证信息，格式为 `"用户名": "密码"`

#### 上游连接配置

- `enable_upstream`: 是否启用上游连接选择器
- `upstream_config`: 上游连接配置列表

##### 上游连接类型

- `type`: `"direct"` = TCP 直连，`"socks5"` = SOCKS5 代理，`"websocket"` =
  WebSocket 代理，`"http"` = HTTP 代理
- `proxy_address`: 代理服务器地址，支持以下协议前缀：
  - `tcp://` 或 `socks5://` - 普通 TCP 连接
  - `tls://` 或 `socks5s://` - TLS 加密连接
  - `ws://` - WebSocket 连接
  - `http://` - HTTP 代理连接
  - 无前缀 - 默认使用 TCP 连接
- `proxy_username`: 代理用户名
- `proxy_password`: 代理密码
- `timeout`: 代理连接超时时间

### 配置示例

#### 1. 基本 SOCKS5 服务端

```json
{
  "listen_addr": ":1080",
  "protocol": "socks5",
  "timeout": 30000000000,
  "auth_users": {
    "admin": "password123"
  },
  "enable_upstream": false,
  "upstream_config": [
    {
      "type": "direct",
      "proxy_address": "",
      "proxy_username": "",
      "proxy_password": "",
      "timeout": 30000000000
    }
  ]
}
```

#### 2. 启用上游连接选择器

```json
{
  "listen_addr": ":1080",
  "protocol": "socks5",
  "timeout": 30000000000,
  "auth_users": {
    "admin": "password123",
    "user1": "pass123"
  },
  "enable_upstream": true,
  "upstream_config": [
    {
      "type": "direct",
      "proxy_address": "",
      "proxy_username": "",
      "proxy_password": "",
      "timeout": 30000000000
    },
    {
      "type": "socks5",
      "proxy_address": "tcp://proxy.example.com:1080",
      "proxy_username": "proxy_user",
      "proxy_password": "proxy_pass",
      "timeout": 30000000000
    },
    {
      "type": "websocket",
      "proxy_address": "ws://websocket-proxy.example.com:8080/proxy",
      "proxy_username": "ws_user",
      "proxy_password": "ws_pass",
      "timeout": 30000000000
    },
    {
      "type": "http",
      "proxy_address": "http://http-proxy.example.com:8080",
      "proxy_username": "http_user",
      "proxy_password": "http_pass",
      "timeout": 30000000000
    }
  ]
}
```

### HTTP 代理支持

系统支持使用 HTTP 代理作为上游连接方式，提供完整的 HTTP 代理功能。

#### HTTP 代理配置

- **类型**: `http`
- **地址格式**: `http://proxy-server:port`
- **认证**: 支持基本认证（Basic Authentication）
- **协议支持**: 支持 HTTP/HTTPS 请求转发

#### HTTP 代理特性

- **CONNECT 方法**: 支持 HTTPS 连接的隧道模式
- **认证机制**: 支持 Proxy-Authorization 头部认证
- **连接池**: 实现连接复用提高性能
- **超时控制**: 支持连接和读写超时设置

#### 使用示例

```json
{
  "type": "http",
  "proxy_address": "http://proxy.example.com:8080",
  "proxy_username": "user",
  "proxy_password": "password",
  "timeout": 30000000000
}
```

#### 命令行启动示例

```bash
# 使用 HTTP 代理作为上游连接
./proxy-server.exe -mode server -protocol socks5 -addr :1080 -upstream-type http -upstream-address http://127.0.0.1:8080 -upstream-username user -upstream-password pass
```

## 上游连接选择策略

系统支持多种上游连接选择策略：

### 1. 轮询策略 (Round Robin)

按顺序轮流选择可用的上游连接

### 2. 随机策略 (Random)

随机选择可用的上游连接

### 3. 加权策略 (Weighted)

根据权重选择上游连接（基于超时时间计算权重）

### 4. 故障转移策略 (Failover)

优先选择第一个健康的上游连接

## 协议支持

### SOCKS5 协议

- 完整支持 SOCKS5 协议 (RFC 1928)
- 支持用户名密码认证 (RFC 1929)
- 支持域名解析和 IPv4/IPv6 连接
- 支持 CONNECT 命令

### WebSocket 协议

- 支持 WebSocket 协议完整实现
- 通过 HTTP Headers 传递认证信息
- 支持连接重试和错误恢复
- 支持二进制数据传输

## 开发指南

### 项目结构

```
socks5-websocket-proxy-golang/
├── cmd/                    # 命令行入口
│   └── main.go            # 主程序入口
├── config/                 # 配置文件
│   ├── server-config.json # 服务端配置示例

├── internal/              # 内部包
├── pkg/                   # 可重用包
│   ├── interfaces/        # 接口定义
│   ├── proxy/            # 代理实现
│   ├── socks5/           # SOCKS5 协议实现
│   ├── upstream/         # 上游连接选择器
│   └── websocket/        # WebSocket 协议实现
├── tests/                # 测试文件
├── examples/             # 示例配置
├── go.mod                # Go 模块文件
└── README.md             # 项目说明
```

### 核心接口

#### ProxyClient 接口

```go
type ProxyClient interface {
    Connect(host string, port int) error
    Authenticate(username, password string) error
    ForwardData(conn net.Conn) error
    Close() error
}
```

#### ProxyServer 接口

```go
type ProxyServer interface {
    Listen() error
    HandleConnection(conn net.Conn) error
    Authenticate(username, password string) bool
    SelectUpstreamConnection(targetHost string, targetPort int) (net.Conn, error)
    Shutdown() error
}
```

## 性能优化

### 并发处理

- 使用 goroutine 池管理并发连接
- 实现连接数限制和超时控制
- 支持优雅关闭和资源清理
- 实现背压机制防止过载

### 内存优化

- 使用缓冲池减少内存分配
- 实现零拷贝数据传输
- 支持流式处理大文件

## 故障排除

### 常见问题

#### 1. 连接被拒绝

- 检查监听地址是否正确
- 确认端口未被占用
- 验证防火墙设置

#### 2. 认证失败

- 检查用户名密码是否正确
- 确认配置文件格式正确
- 验证认证信息是否正确加载

#### 3. 上游连接失败

- 检查上游代理服务器地址
- 验证上游代理认证信息
- 确认网络连接正常

### 日志系统

系统提供完整的日志记录功能，包括连接成功和失败状态的详细日志输出：

#### 服务端连接日志

- **SOCKS5 服务端**: 记录客户端连接地址、认证状态、连接建立和处理结果
- **WebSocket 服务端**: 记录连接尝试、认证成功/失败、升级失败及连接处理结果

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 联系方式

- 项目地址: https://github.com/masx200/socks5-websocket-proxy-golang
- 问题反馈: https://github.com/masx200/socks5-websocket-proxy-golang/issues

## 更新日志

### v1.0.0 (2025-09-06)

- ✅ 完成核心功能开发
- ✅ 支持 SOCKS5 和 WebSocket 协议
- ✅ 实现上游连接选择器
- ✅ 支持多种选择策略和健康检查
- ✅ 完善配置管理和文档
- ✅ 修改 UpstreamType 为字符串类型
- ✅ 实现 SOCKS5 服务端连接日志功能
- ✅ 实现 WebSocket 服务端连接日志功能
- ✅ 完善客户端模式使用指南
- ✅ 添加详细的连接状态监控和日志输出

# socks5-websocket-proxy-golang

A high-performance SOCKS5 proxy server with WebSocket support written in Go

## 使用 curl 测试

```bash
curl -v  -I   https://www.baidu.com  -x socks5://localhost:9999
```

## 测试项目

```bash
go test -v -p 1 -parallel 1 ./...
```
