# 变更日志 (Changelog)

本文档记录了项目的所有重要变更。

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

## 注意事项

- 配置文件修改后，系统会自动检测变化并应用新配置
- 协议切换会导致当前连接中断，但服务器会快速重启
- 建议在低峰期进行协议切换操作
- 所有配置变更都会记录在日志中，便于调试和监控

---

_最后更新时间: 2025-09-06_
