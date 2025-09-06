# 配置文件说明

## 配置文件格式

配置文件使用 JSON 格式，支持以下配置项：

### 基本配置

- `listen_addr`: 监听地址，格式为 `:端口` 或 `IP:端口`
- `protocol`: 协议类型，支持 `socks5` 或 `websocket`
- `timeout`: 超时时间，单位为纳秒（30 秒 = 30000000000）

### 认证配置

- `auth_users`: 用户认证信息，格式为 `"用户名": "密码"`

### 上游连接配置

- `enable_upstream`: 是否启用上游连接选择器
- `upstream_config`: 上游连接配置列表

#### 上游连接类型

- `type`: 0 = TCP 直连，1 = SOCKS5 代理，2 = WebSocket 代理
- `proxy_address`: 代理服务器地址
- `proxy_username`: 代理用户名
- `proxy_password`: 代理密码
- `timeout`: 代理连接超时时间

## 使用示例

### 1. 基本 SOCKS5 服务端

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
      "type": 0,
      "proxy_address": "",
      "proxy_username": "",
      "proxy_password": "",
      "timeout": 30000000000
    }
  ]
}
```

### 2. 启用上游连接选择器

```json
{
  "listen_addr": ":1080",
  "protocol": "socks5",
  "timeout": 30000000000,
  "auth_users": {
    "admin": "password123"
  },
  "enable_upstream": true,
  "upstream_config": [
    {
      "type": 0,
      "proxy_address": "",
      "proxy_username": "",
      "proxy_password": "",
      "timeout": 30000000000
    },
    {
      "type": 1,
      "proxy_address": "proxy.example.com:1080",
      "proxy_username": "proxy_user",
      "proxy_password": "proxy_pass",
      "timeout": 30000000000
    }
  ]
}
```

## 运行方式

使用配置文件启动服务端：

```bash
./proxy-server.exe -mode server -config config/server-config.json
```

配置文件中的参数会覆盖命令行参数。
