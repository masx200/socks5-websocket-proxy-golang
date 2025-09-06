# SOCKS5-WebSocket 代理系统开发任务清单

## 项目概述

使用 Go 语言开发支持 SOCKS5 和 WebSocket 协议的网络代理系统

## 已完成任务

- [x] 项目初始化（go.mod 创建）
- [x] 配置文件解析功能实现
- [x] 配置文件模板和说明文档创建

## 待完成任务

### 1. 项目基础结构搭建

- [x] 创建项目目录结构
  - [x] 创建`config/`目录 - 配置文件
  - [x] 创建`internal/`目录 - 内部包
  - [x] 创建`pkg/`目录 - 可重用包
  - [x] 创建`cmd/`目录 - 命令行入口
  - [x] 创建`examples/`目录 - 示例配置
  - [x] 创建`tests/`目录 - 测试文件

### 2. 核心接口定义

- [x] 定义统一客户端接口`ProxyClient`
- [x] 定义统一服务端接口`ProxyServer`
- [x] 定义配置结构体`ClientConfig`和`ServerConfig`
- [x] 定义上游连接配置`UpstreamConfig`
- [x] 创建客户端和服务端工厂函数

### 3. SOCKS5 组件实现

- [x] 实现 SOCKS5 客户端
  - [x] 基础连接功能
  - [x] 用户名密码认证
  - [x] 数据转发
- [x] 实现 SOCKS5 服务端
  - [x] 监听和连接处理
  - [x] 认证验证
  - [x] 数据转发
  - [x] 集成 github.com/armon/go-socks5 库
  - [x] 删除后备方案代码

### 4. WebSocket 组件实现

- [x] 实现 WebSocket 客户端
  - [x] 基础连接功能
  - [x] HTTP Headers 认证
  - [x] 数据转发
- [x] 实现 WebSocket 服务端
  - [x] HTTP 服务端和 WebSocket 升级
  - [x] 认证信息解析
  - [x] 数据转发

### 5. 上游连接选择器

- [x] 实现 TCP 直连模式
- [x] 实现 WebSocket 代理模式
- [x] 实现 SOCKS5 代理模式
- [x] 实现动态选择机制
  - [x] 轮询策略
  - [x] 随机策略
  - [x] 加权策略
  - [x] 故障转移策略
  - [x] 健康检查机制

### 6. 配置管理

- [x] 实现配置文件解析
- [x] 实现配置热重载
- [x] 创建配置文件模板

### 7. 日志系统

- [x] 实现服务端连接成功和失败日志打印
  - [x] SOCKS5 服务端连接日志
  - [x] WebSocket 服务端连接日志

### 10. 文档和示例

- [x] API 文档
- [x] 使用示例
- [x] 部署说明
- [x] 配置文件示例

## 开发优先级

1. **高优先级**：核心接口定义、SOCKS5 组件、WebSocket 组件、上游连接选择器
2. **中优先级**：配置管理、日志系统、测试开发
3. **低优先级**：文档示例、构建部署、高级功能

## 当前进度

- 项目初始化：✅ 完成
- 核心开发：✅ 完成（上游连接选择器已完成）
- 配置管理：🔄 进行中（配置文件解析已完成）
- 测试验证：⏳ 待开始
- 文档完善：🔄 进行中（配置文件模板已创建）

## 依赖库

- `golang.org/x/net/websocket`
- `golang.org/x/net/proxy` (SOCKS5 支持)
- `github.com/gorilla/websocket` (WebSocket 支持)
- `github.com/armon/go-socks5`

## 更新记录

- 2025-06-17：创建任务清单，开始项目开发
- 2025-06-18：完成上游连接选择器实现，支持多种选择策略和健康检查
- 2025-06-20：完成配置文件解析功能，支持 JSON
  格式配置文件，创建配置文件模板和说明文档
- 2025-06-23：修改 UpstreamType 从数字类型改为字符串类型，更新配置文件中的 type
  字段为字符串格式
- 2025-06-23：编写完整的 README.md
  文档，包含项目概述、功能特性、安装使用、配置说明、开发指南等内容
- 2025-09-06：修改 SOCKS5 服务端使用 github.com/armon/go-socks5 库，修复编译错误
- 2025-09-06：删除 SOCKS5 服务端后备方案代码，完全依赖
  github.com/armon/go-socks5 库
- 2025-09-06：为服务端添加连接成功和失败日志打印功能，包括 SOCKS5 和 WebSocket
  服务端的详细连接状态日志
- 2025-09-06：修复 SOCKS5
  客户端地址解析问题，解决尾部斜杠导致端口解析错误的问题，在 parseServerAddr
  函数中添加 strings.TrimSuffix 处理
- 2025-09-06：修改 WebSocket 协议实现，将 targetHost 和 targetPort 从 URL
  查询参数移动到 HTTP Headers 中，提高安全性
- 2025-09-06：完善 SOCKS5 和 WebSocket 服务端的日志打印功能，统一日志前缀格式，添加连接状态、认证信息、上游选择、数据转发等详细日志信息
- 2025-09-06：修复 WebSocket 服务端空指针解引用错误，解决在 SelectUpstreamConnection 方法中因 selector 初始化不当导致的 panic 问题，优化 NewWebSocketServer 函数的 selector 初始化逻辑
