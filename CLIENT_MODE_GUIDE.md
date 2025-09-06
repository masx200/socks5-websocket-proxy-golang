# SOCKS5-WebSocket 代理系统 - 客户端模式使用指南

## 概述

本指南详细介绍如何使用 SOCKS5-WebSocket
代理系统的客户端模式。客户端模式允许您通过 SOCKS5 或 WebSocket
协议连接到代理服务器，并将流量转发到目标主机。

## 客户端模式架构

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

## 基本使用方法

### 命令行参数

客户端模式支持以下命令行参数：

```bash
./proxy-server.exe -mode client [其他参数]
```

#### 必需参数

- `-mode client`: 指定运行模式为客户端
- `-protocol`: 协议类型，支持 `socks5` 或 `websocket`
- `-addr`: 代理服务器地址
- `-host`: 目标主机地址
- `-port`: 目标主机端口

#### 可选参数

- `-username`: 代理服务器用户名
- `-password`: 代理服务器密码
- `-timeout`: 连接超时时间（秒）
- `-config`: 配置文件路径

### SOCKS5 客户端模式

#### 基本连接

```bash
# 连接到 SOCKS5 代理服务器并访问目标主机
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -host www.example.com -port 80
```

#### 带认证的连接

```bash
# 使用用户名密码认证连接到 SOCKS5 代理服务器
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -username admin -password password123 -host www.example.com -port 443
```

#### 自定义超时时间

```bash
# 设置连接超时时间为 60 秒
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -host www.example.com -port 80 -timeout 60
```

### WebSocket 客户端模式

#### 基本连接

```bash
# 连接到 WebSocket 代理服务器并访问目标主机
./proxy-server.exe -mode client -protocol websocket -addr ws://proxy.example.com:8080/proxy -host www.example.com -port 80
```

#### 带认证的连接

```bash
# 使用用户名密码认证连接到 WebSocket 代理服务器
./proxy-server.exe -mode client -protocol websocket -addr ws://proxy.example.com:8080/proxy -username admin -password password123 -host www.example.com -port 443
```

#### 自定义超时时间

```bash
# 设置连接超时时间为 60 秒
./proxy-server.exe -mode client -protocol websocket -addr ws://proxy.example.com:8080/proxy -host www.example.com -port 80 -timeout 60
```

## 实际应用场景

### 1. 网络访问代理

#### 通过 SOCKS5 代理访问受限网站

```bash
# 通过 SOCKS5 代理访问 Google
./proxy-server.exe -mode client -protocol socks5 -addr tcp://your-socks5-server.com:1080 -username user -password pass -host www.google.com -port 443
```

#### 通过 WebSocket 代理绕过防火墙

```bash
# 通过 WebSocket 代理访问被防火墙阻止的服务
./proxy-server.exe -mode client -protocol websocket -addr ws://websocket-proxy.com:8080/proxy -username user -password pass -host blocked-site.com -port 443
```

### 2. 数据库连接代理

#### MySQL 数据库连接

```bash
# 通过 SOCKS5 代理连接远程 MySQL 数据库
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.company.com:1080 -username dbuser -password dbpass -host mysql.internal.company.com -port 3306
```

#### PostgreSQL 数据库连接

```bash
# 通过 WebSocket 代理连接远程 PostgreSQL 数据库
./proxy-server.exe -mode client -protocol websocket -addr ws://proxy.company.com:8080/proxy -username dbuser -password dbpass -host postgres.internal.company.com -port 5432
```

### 3. API 服务访问

#### REST API 访问

```bash
# 通过 SOCKS5 代理访问内部 API 服务
./proxy-server.exe -mode client -protocol socks5 -addr tcp://api-gateway.company.com:1080 -username apiuser -password apipass -host api.internal.company.com -port 8080
```

#### GraphQL API 访问

```bash
# 通过 WebSocket 代理访问 GraphQL 服务
./proxy-server.exe -mode client -protocol websocket -addr ws://api-gateway.company.com:8080/proxy -username apiuser -password apipass -host graphql.internal.company.com -port 4000
```

### 4. 开发和测试环境

#### 本地开发连接远程服务

```bash
# 本地开发时通过 SOCKS5 代理连接远程开发环境
./proxy-server.exe -mode client -protocol socks5 -addr tcp://dev-proxy.company.com:1080 -username devuser -password devpass -host dev-service.internal.company.com -port 3000
```

#### 测试环境服务访问

```bash
# 通过 WebSocket 代理访问测试环境服务
./proxy-server.exe -mode client -protocol websocket -addr ws://test-proxy.company.com:8080/proxy -username testuser -password testpass -host test-service.internal.company.com -port 5000
```

## 高级配置

### 1. 多重代理链

您可以通过组合多个代理实例来创建代理链：

```bash
# 第一层：WebSocket 代理
./proxy-server.exe -mode client -protocol websocket -addr ws://first-proxy.com:8080/proxy -username user1 -password pass1 -host second-proxy.com -port 1080 &

# 第二层：SOCKS5 代理
./proxy-server.exe -mode client -protocol socks5 -addr tcp://localhost:1080 -username user2 -password pass2 -host target-site.com -port 443
```

### 2. 端口转发

将本地端口转发到远程服务：

```bash
# 将本地 8080 端口转发到远程服务
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -username user -password pass -host remote-service.com -port 80
```

### 3. 动态端口分配

使用系统自动分配的端口：

```bash
# 让系统自动分配本地端口
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -username user -password pass -host remote-service.com -port 80
```

## 故障排除

### 常见问题及解决方案

#### 1. 连接被拒绝

**问题**: `dial tcp tcp://proxy.example.com:1080: connect: connection refused`

**解决方案**:

- 检查代理服务器地址是否正确
- 确认代理服务器正在运行
- 检查网络连接和防火墙设置
- 验证端口是否正确

#### 2. 认证失败

**问题**: `authentication failed`

**解决方案**:

- 检查用户名和密码是否正确
- 确认代理服务器支持指定的认证方式
- 验证认证信息格式是否正确

#### 3. 超时错误

**问题**: `i/o timeout`

**解决方案**:

- 增加超时时间设置
- 检查网络延迟和稳定性
- 确认目标主机可达性

#### 4. 协议不支持

**问题**: `unsupported protocol: xxx`

**解决方案**:

- 检查协议名称是否正确（支持 `socks5` 和 `websocket`）
- 确认代理服务器支持指定的协议
- 检查协议名称拼写

### 调试技巧



#### 2. 测试连接连通性

```bash
# 测试代理服务器连通性
telnet proxy.example.com 1080

# 测试目标主机连通性
telnet www.example.com 80
```

## 性能优化

### 1. 连接池配置

客户端模式支持连接复用，可以通过以下方式优化性能：

```bash
# 设置较长的超时时间以复用连接
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -host target.com -port 80 -timeout 300
```

### 2. 并发连接

客户端模式支持多个并发连接：

```bash
# 启动多个客户端实例
for i in {1..5}; do
    ./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -host target.com -port 80 &
done
```

### 3. 资源管理

合理设置超时时间以避免资源浪费：

```bash
# 根据网络状况设置合适的超时时间
./proxy-server.exe -mode client -protocol socks5 -addr tcp://proxy.example.com:1080 -host target.com -port 80 -timeout 60
```

## 安全注意事项

### 1. 认证安全

- 使用强密码
- 定期更换密码
- 不要在命令行中明文传递密码（建议使用配置文件）
- 限制认证失败次数

### 2. 网络安全

- 使用加密连接（如 HTTPS、WSS）
- 验证服务器证书
- 使用可信的代理服务器
- 避免在不安全的网络中使用

### 3. 数据安全

- 敏感数据使用加密传输
- 定期审计访问日志
- 限制访问权限
- 监控异常活动

## 最佳实践

### 1. 配置管理

- 使用配置文件管理敏感信息
- 将配置文件权限设置为仅当前用户可读
- 定期备份配置文件
- 使用版本控制管理配置变更

### 2. 监控和日志

- 启用日志记录
- 定期检查日志文件
- 设置监控告警
- 记录关键操作和错误

### 3. 维护和更新

- 定期更新代理软件
- 监控安全漏洞
- 测试新版本兼容性
- 制定回滚计划

## 总结

SOCKS5-WebSocket 代理系统的客户端模式提供了灵活的网络代理解决方案，支持 SOCKS5
和 WebSocket 两种协议。通过本指南，您应该能够：

1. 理解客户端模式的基本概念和架构
2. 掌握命令行参数和配置文件的使用方法
3. 应对各种实际应用场景
4. 解决常见问题和故障
5. 优化性能和确保安全性

如需更多帮助，请参考项目文档或提交问题到项目仓库。
