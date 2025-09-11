# 变更日志 (Changelog)

本文档记录了项目的所有重要变更。

## [2025-09-11]

### 新增功能 (Features)

#### HTTP代理支持

- **提交**: `http-upstream-support` - feat(upstream):
  添加HTTP代理支持并改进连接日志
- **描述**: 为上游代理选择器添加HTTP代理支持，包括创建HTTP连接和包装器实现
- **影响**: 现在支持使用HTTP代理作为上游连接，扩展了代理系统的兼容性和灵活性

### 优化改进 (Improvements)

#### HTTP代理客户端连接日志

- **提交**: `http-upstream-support` - feat(http):
  在HTTP代理客户端中添加连接日志以便调试
- **描述**:
  在HTTP代理客户端中添加详细的连接日志记录，包括连接建立、数据转发和连接关闭等关键节点
- **影响**: 提高了HTTP代理连接的可调试性，便于排查连接问题和性能优化

#### 上游配置优化

- **提交**: `http-upstream-support` - feat(upstream):
  当配置上游代理但未配置选择器时，直接使用上游配置创建连接
- **描述**:
  优化了上游连接处理逻辑，当启用上游代理但选择器未正确配置时，直接使用上游配置创建连接
- **影响**: 提高了系统的容错性和稳定性，避免了因选择器配置问题导致的连接失败

## [2025-09-07]

### Bug修复 (Bug Fixes)

#### SOCKS5服务器类型转换错误修复

- **提交**: `2ac44df` - fix(socks5): 修复panic: interface conversion: net.Addr
  is net.pipeAddr, not *net.TCPAddr
- **描述**:
  修复了在使用WebSocket上游连接时SOCKS5服务器发生的类型转换错误，解决了net.pipeAddr无法转换为*net.TCPAddr的问题
- **影响**:
  解决了服务运行时的panic崩溃问题，提高了WebSocket上游连接的稳定性和可靠性

#### SOCKS5客户端NetConn方法实现

- **提交**: `2ac44df` - fix(socks5): 实现SOCKS5客户端NetConn方法
- **描述**:
  实现了SOCKS5客户端中未完成的NetConn方法，替换了原来的panic("unimplemented")实现
- **影响**:
  使SOCKS5客户端完整实现了ProxyClient接口，避免了调用NetConn方法时的程序崩溃

## [2025-09-06]

### 新增功能 (Features)

#### 配置热重载功能

- **提交**: `7995bfc` - feat(配置): 实现配置热重载功能
- **提交**: `a1d26a1` - feat(配置): 添加配置文件动态重载功能
- **描述**: 实现了配置文件的动态重载功能，支持在不重启服务器的情况下更新配置
- **影响**: 用户可以在运行时修改配置文件，系统会自动检测并应用新的配置

#### 协议热切换功能

- **提交**: `f41b5b8` - feat(config): 支持协议热切换并增强配置监听功能
- **描述**: 实现了协议的热切换功能，支持在SOCKS5和WebSocket协议之间动态切换
- **影响**:
  用户可以通过修改配置文件中的protocol字段来切换服务器协议，无需手动重启

### 优化改进 (Improvements)

#### 配置处理逻辑优化

- **提交**: `fabc271` - refactor(config): 优化配置处理逻辑并添加日志
- **描述**: 优化了配置文件的处理逻辑，增加了详细的日志记录
- **影响**: 提高了配置处理的可靠性和可调试性

#### 上游选择器重构

- **提交**: `bdedc4b` - refactor(upstream):
  修改选择器类型为interface以支持多种实现
- **描述**: 将上游选择器的类型修改为interface，支持多种不同的实现方式
- **影响**: 提高了代码的可扩展性和灵活性

#### WebSocket协议安全性优化

- **提交**: `websocket-security` - refactor(websocket):
  将目标信息从URL查询参数移动到HTTP Headers
- **描述**:
  修改WebSocket协议实现，将targetHost和targetPort从URL查询参数移动到HTTP
  Headers中传输
- **影响**:
  提高了协议安全性，减少了URL中暴露敏感信息的风险，使WebSocket连接URL更加简洁

#### WebSocket服务端稳定性优化

- **提交**: `websocket-fix` - fix(websocket): 修复服务端空指针解引用错误
- **描述**:
  修复WebSocket服务端在SelectUpstreamConnection方法中的空指针解引用错误，优化NewWebSocketServer函数的selector初始化逻辑
- **影响**:
  解决了因selector初始化不当导致的panic问题，提高了WebSocket服务端的稳定性和可靠性

#### 客户端连接关闭callback功能

- **提交**: `connection-callback` - feat(client): 实现连接关闭callback功能
- **描述**:
  为ProxyClient接口添加SetConnectionClosedCallback方法，支持连接关闭时自动触发回调函数，替换原来的轮询检查机制
- **影响**:
  提高了连接状态检测的效率和可靠性，当客户端连接被关闭时程序能够自动退出，改善了用户体验

### 文档更新 (Documentation)

#### 文档清理和更新

- **提交**: `d720f3a` - docs: 清理文档中的多余空行并更新内容
- **提交**: `b46daac` - Update README.md
- **描述**: 清理了文档中的多余空行，更新了相关文档内容
- **影响**: 改善了文档的可读性和维护性

### 配置文件更新

#### 服务器配置更新

- **提交**: `e1f712f` - Update server-config.json
- **提交**: `fc91d54` - Update server-config.json
- **描述**: 更新了服务器配置文件，包含最新的协议和配置参数
- **影响**: 确保配置文件与最新功能保持一致

## 技术细节

### 配置热重载实现

- 使用 `fsnotify` 库监听配置文件变化
- 实现防抖机制避免频繁重载
- 支持配置验证和错误处理
- 添加详细的日志记录

### 协议热切换实现

- 检测配置文件中的protocol字段变更
- 实现服务器的优雅关闭和重启
- 支持SOCKS5和WebSocket协议之间的无缝切换
- 保持配置监听器在重启后的连续性

### WebSocket协议安全性优化实现

- 修改客户端buildWebSocketURL方法，移除URL查询参数中的host和port
- 更新客户端buildHeaders方法，添加X-Proxy-Target-Host和X-Proxy-Target-Port头部
- 修改服务端parseAuthInfo方法，从HTTP Headers中解析目标主机和端口信息
- 保持原有认证机制和功能完整性，提高了数据传输安全性

### WebSocket服务端稳定性优化实现

- 修复NewWebSocketServer函数中selector变量的初始化逻辑，将类型从*upstream.UpstreamSelector改为interface{}
- 在SelectUpstreamConnection方法中加强类型断言检查，同时检查类型和是否为nil
- 优化错误处理逻辑，在selector类型检查失败时立即返回错误并记录详细日志
- 确保当config.EnableUpstream为false时，selector被正确设置为nil，避免空指针解引用

### WebSocket服务端interface{}类型检查优化实现

- 添加isNilInterface函数，使用反射(reflect)正确检查interface{}类型是否真正为nil
- 修复SelectUpstreamConnection方法中的nil检查逻辑，从简单的`s.selector != nil`改为`s.selector != nil && !isNilInterface(s.selector)`
- 解决在Upstream selector enabled: false情况下仍进入选择器分支的问题
- 添加reflect包导入，支持更严格的interface{}内部值检查，避免包含nil指针的interface{}被误判为有效值

### 客户端连接关闭callback功能实现

- 在ProxyClient接口中添加SetConnectionClosedCallback(callback func())
  error方法声明
- 修改cmd/main.go中的startClient函数，使用callback机制替换原来的轮询检查连接状态方式
- 在SOCKS5Client和WebSocketClient结构体中添加connectionClosedCallback func()字段
- 实现SetConnectionClosedCallback方法用于设置回调函数，在Close方法中调用connectionClosedCallback
- 移除不再使用的isConnectionAlive函数和reflect包导入，清理代码提高可维护性
- 通过callback机制实现连接关闭时自动退出程序，提高效率和可靠性

### SOCKS5服务器类型转换错误修复实现

- 在pkg/upstream/selector.go中添加tcpAddrWrapper结构体，包装*net.TCPAddr以解决类型兼容性问题
- 实现newTCPAddrWrapper函数，根据不同地址情况返回兼容的TCP地址包装器
- 修改SOCKS5ConnWrapper和WebSocketConnWrapper的LocalAddr()和RemoteAddr()方法，将dummyAddr替换为tcpAddrWrapper
- 确保连接包装器返回的地址类型与SOCKS5库期望的*net.TCPAddr类型兼容
- 解决了net.Pipe()创建的连接返回net.pipeAddr类型导致的类型转换panic问题

### SOCKS5客户端NetConn方法实现

- 修改pkg/socks5/client.go中的NetConn方法，将panic("unimplemented")替换为return
  c.conn
- 使NetConn方法正确返回SOCKS5Client实例中的底层网络连接对象
- 确保SOCKS5客户端完整实现了interfaces.ProxyClient接口的所有方法
- 避免了调用NetConn方法时的程序崩溃，提高了接口的完整性和可用性

## 注意事项

- 配置文件修改后，系统会自动检测变化并应用新配置
- 协议切换会导致当前连接中断，但服务器会快速重启
- 建议在低峰期进行协议切换操作
- 所有配置变更都会记录在日志中，便于调试和监控

---

_最后更新时间: 2025-09-07_
