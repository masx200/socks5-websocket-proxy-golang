### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:36:24

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 19804

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :1080 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 42484

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试1进程PID: 50664

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:36:17 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10775363422310857891
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:36:17 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10775363422310857891
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试2进程PID: 29972

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:36:17 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10921881889941542796
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:36:17 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10921881889941542796
```

### 📋 所有进程PID记录

所有进程PID: 19804, 42484, 50664, 29972

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:36:24

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 19804

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :1080 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 42484

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试1进程PID: 50664

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:36:17 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10775363422310857891
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:36:17 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10775363422310857891
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试2进程PID: 29972

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:36:17 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10921881889941542796
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:36:17 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10921881889941542796
```

### 📋 所有进程PID记录

所有进程PID: 19804, 42484, 50664, 29972

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程... ✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程... ✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程... ✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件

### WebSocket服务器日志输出

```
2025/09/11 23:36:25 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/11 23:36:25 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/11 23:36:25 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:36:25 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/11 23:36:25 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/11 23:36:25 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/11 23:36:26 server.go:90: url /
2025/09/11 23:36:26 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[BvLEnTazHQiK404OmdgdMw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:36:26 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:45159 at 2025-09-11 23:36:26
2025/09/11 23:36:26 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:45159
2025/09/11 23:36:26 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:36:26 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:36:26 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:45159
2025/09/11 23:36:26 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:45159
2025/09/11 23:36:26 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:36:26 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:36:26 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:36:29 server.go:90: url /
2025/09/11 23:36:29 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[YWZiLRaP5Ji3NkrJe8Xm0Q==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:36:29 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:45211 at 2025-09-11 23:36:29
2025/09/11 23:36:29 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:45211
2025/09/11 23:36:29 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:36:29 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:36:29 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:45211
2025/09/11 23:36:29 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:45211
2025/09/11 23:36:29 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:36:29 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:36:29 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:36:29 server.go:90: url /
2025/09/11 23:36:29 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[BOv/SIeEjxBNvgH0CWZQ1Q==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:36:29 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:45216 at 2025-09-11 23:36:29
2025/09/11 23:36:29 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:45216
2025/09/11 23:36:29 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/11 23:36:29 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:36:29 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:45216
2025/09/11 23:36:29 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:45216
2025/09/11 23:36:29 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/11 23:36:29 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/11 23:36:29 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/11 23:36:26 main.go:71: 启动socks5服务端，监听地址: :1080
2025/09/11 23:36:26 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/11 23:36:26 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :1080
2025/09/11 23:36:26 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:36:26 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/11 23:36:26 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/11 23:36:26 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:45150 (total: 1)
2025/09/11 23:36:26 server.go:164: [SOCKS5-CONN] New connection from [::1]:45150 at 2025-09-11 23:36:26
2025/09/11 23:36:26 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:45150
2025/09/11 23:36:26 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:45150
2025/09/11 23:36:26 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:1080, Remote: [::1]:45150
2025/09/11 23:36:26 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:36:26 client.go:98: url: ws://localhost:8080
2025/09/11 23:36:26 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:36:26 client.go:110: url: http://localhost:8080
2025/09/11 23:36:26 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[um4aszTncLwlFwZqJFsWtds6i8o=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:36:26 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:36:29 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:45208 (total: 2)
2025/09/11 23:36:29 server.go:164: [SOCKS5-CONN] New connection from [::1]:45208 at 2025-09-11 23:36:29
2025/09/11 23:36:29 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:45208
2025/09/11 23:36:29 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:45208
2025/09/11 23:36:29 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:1080, Remote: [::1]:45208
2025/09/11 23:36:29 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:36:29 client.go:98: url: ws://localhost:8080
2025/09/11 23:36:29 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:36:29 client.go:110: url: http://localhost:8080
2025/09/11 23:36:29 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[eGehmGlpFxP+vh/Mafr+rfI0eRI=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:36:29 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:36:29 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:45213 (total: 3)
2025/09/11 23:36:29 server.go:164: [SOCKS5-CONN] New connection from [::1]:45213 at 2025-09-11 23:36:29
2025/09/11 23:36:29 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:45213
2025/09/11 23:36:29 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:45213
2025/09/11 23:36:29 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:1080, Remote: [::1]:45213
2025/09/11 23:36:29 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/11 23:36:29 client.go:98: url: ws://localhost:8080
2025/09/11 23:36:29 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:36:29 client.go:110: url: http://localhost:8080
2025/09/11 23:36:29 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[Q84AlmKTW4cduTodadLT1gc2jSs=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:36:29 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/11 23:36:29 selector.go:270: WebSocket forward data error: read tcp [::1]:45211->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:36:29 selector.go:270: WebSocket forward data error: read tcp [::1]:45216->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:36:29 selector.go:270: WebSocket forward data error: read tcp [::1]:45159->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:36:29 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:45213, duration: 84.2759ms
2025/09/11 23:36:29 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:45213
2025/09/11 23:36:29 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:45208, duration: 145.0847ms
2025/09/11 23:36:29 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:45208
2025/09/11 23:36:29 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:45150, duration: 2.3087024s
2025/09/11 23:36:29 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:45150
```

✅ 端口8080已成功释放 ✅ 端口1080已成功释放

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:48:14

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 20772

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :1080 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 54224

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试1进程PID: 3804

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:48:07 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11611241440836188865
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:48:07 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11611241440836188865
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试2进程PID: 54504

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:48:07 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11436605390823430906
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:48:07 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11436605390823430906
```

### 📋 所有进程PID记录

所有进程PID: 20772, 54224, 3804, 54504

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:48:14

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 20772

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :1080 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 54224

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试1进程PID: 3804

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:48:07 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11611241440836188865
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:48:07 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11611241440836188865
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:1080`

📋 Curl测试2进程PID: 54504

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:48:07 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11436605390823430906
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:48:07 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11436605390823430906
```

### 📋 所有进程PID记录

所有进程PID: 20772, 54224, 3804, 54504

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程... ✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程... ✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程... ✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件

### WebSocket服务器日志输出

```
2025/09/11 23:48:15 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/11 23:48:15 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/11 23:48:15 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:48:15 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/11 23:48:15 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/11 23:48:15 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/11 23:48:17 server.go:90: url /
2025/09/11 23:48:17 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[nAckR1R5Y+EcvkGUJ/xRGw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:48:17 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:52263 at 2025-09-11 23:48:17
2025/09/11 23:48:17 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:52263
2025/09/11 23:48:17 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:48:17 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:48:17 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:52263
2025/09/11 23:48:17 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:52263
2025/09/11 23:48:17 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:48:17 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:48:17 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:48:19 server.go:90: url /
2025/09/11 23:48:19 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[UoMtcd6A42WCbXHkhppWHw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:48:19 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:52268 at 2025-09-11 23:48:19
2025/09/11 23:48:19 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:52268
2025/09/11 23:48:19 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:48:19 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:48:19 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:52268
2025/09/11 23:48:19 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:52268
2025/09/11 23:48:19 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:48:19 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:48:19 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:48:19 server.go:90: url /
2025/09/11 23:48:19 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[pZpirmpw0Q6hIwpVs7vjxQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:48:19 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:52273 at 2025-09-11 23:48:19
2025/09/11 23:48:19 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:52273
2025/09/11 23:48:19 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/11 23:48:19 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:48:19 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:52273
2025/09/11 23:48:19 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:52273
2025/09/11 23:48:19 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/11 23:48:19 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/11 23:48:19 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/11 23:48:16 main.go:71: 启动socks5服务端，监听地址: :1080
2025/09/11 23:48:16 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/11 23:48:16 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :1080
2025/09/11 23:48:16 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:48:16 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/11 23:48:16 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/11 23:48:17 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:52262 (total: 1)
2025/09/11 23:48:17 server.go:164: [SOCKS5-CONN] New connection from [::1]:52262 at 2025-09-11 23:48:17
2025/09/11 23:48:17 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:52262
2025/09/11 23:48:17 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:52262
2025/09/11 23:48:17 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:1080, Remote: [::1]:52262
2025/09/11 23:48:17 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:48:17 client.go:98: url: ws://localhost:8080
2025/09/11 23:48:17 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:48:17 client.go:110: url: http://localhost:8080
2025/09/11 23:48:17 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[dnLjVYKu1oWycKKsjBuAssVZu5w=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:48:17 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:48:19 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:52265 (total: 2)
2025/09/11 23:48:19 server.go:164: [SOCKS5-CONN] New connection from [::1]:52265 at 2025-09-11 23:48:19
2025/09/11 23:48:19 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:52265
2025/09/11 23:48:19 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:52265
2025/09/11 23:48:19 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:1080, Remote: [::1]:52265
2025/09/11 23:48:19 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:48:19 client.go:98: url: ws://localhost:8080
2025/09/11 23:48:19 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:48:19 client.go:110: url: http://localhost:8080
2025/09/11 23:48:19 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[nSxwzWwlrA6TrW1+iHQroGohGIs=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:48:19 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:48:19 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:52270 (total: 3)
2025/09/11 23:48:19 server.go:164: [SOCKS5-CONN] New connection from [::1]:52270 at 2025-09-11 23:48:19
2025/09/11 23:48:19 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:52270
2025/09/11 23:48:19 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:52270
2025/09/11 23:48:19 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:1080, Remote: [::1]:52270
2025/09/11 23:48:19 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/11 23:48:19 client.go:98: url: ws://localhost:8080
2025/09/11 23:48:19 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:48:19 client.go:110: url: http://localhost:8080
2025/09/11 23:48:19 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[2OqJhTu/2WTC2oeIC3sb+6hbZfY=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:48:19 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/11 23:48:19 selector.go:270: WebSocket forward data error: read tcp [::1]:52263->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:48:19 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:52262, duration: 2.2173146s
2025/09/11 23:48:19 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:52262
2025/09/11 23:48:19 selector.go:270: WebSocket forward data error: read tcp [::1]:52268->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:48:19 selector.go:270: WebSocket forward data error: read tcp [::1]:52273->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:48:19 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:52265, duration: 145.5968ms
2025/09/11 23:48:19 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:52265
2025/09/11 23:48:19 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:52270, duration: 79.5878ms
2025/09/11 23:48:19 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:52270
```

✅ 端口8080已成功释放 ✅ 端口1080已成功释放

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:54:52

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 50692

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 50192

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 33328

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:45 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10761801990532464050
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:45 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10761801990532464050
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 24436

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:45 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11050621022157571654
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:45 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11050621022157571654
```

### 📋 所有进程PID记录

所有进程PID: 50692, 50192, 33328, 24436

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:54:52

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 50692

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 50192

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 33328

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:45 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10761801990532464050
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:45 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10761801990532464050
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 24436

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:45 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11050621022157571654
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:45 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11050621022157571654
```

### 📋 所有进程PID记录

所有进程PID: 50692, 50192, 33328, 24436

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程... ✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程... ✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程... ✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件

### WebSocket服务器日志输出

```
2025/09/11 23:54:53 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/11 23:54:53 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/11 23:54:53 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:54:53 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/11 23:54:53 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/11 23:54:53 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/11 23:54:55 server.go:90: url /
2025/09/11 23:54:55 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[X/aNu1HwErDRk/+2swAcgQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:54:55 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:55437 at 2025-09-11 23:54:55
2025/09/11 23:54:55 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:55437
2025/09/11 23:54:55 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:54:55 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:54:55 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:55437
2025/09/11 23:54:55 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:55437
2025/09/11 23:54:55 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:54:55 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:54:55 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:54:57 server.go:90: url /
2025/09/11 23:54:57 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[eYkWP4y/fxzmNNxVB6ZLWw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:54:57 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:55508 at 2025-09-11 23:54:57
2025/09/11 23:54:57 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:55508
2025/09/11 23:54:57 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:54:57 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:54:57 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:55508
2025/09/11 23:54:57 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:55508
2025/09/11 23:54:57 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:54:57 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:54:57 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:54:57 server.go:90: url /
2025/09/11 23:54:57 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[U2qOxVj7YHLGQVMX1exUMg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:54:57 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:55514 at 2025-09-11 23:54:57
2025/09/11 23:54:57 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:55514
2025/09/11 23:54:57 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/11 23:54:57 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:54:57 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:55514
2025/09/11 23:54:57 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:55514
2025/09/11 23:54:57 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/11 23:54:57 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/11 23:54:57 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/11 23:54:54 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/11 23:54:54 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/11 23:54:54 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/11 23:54:54 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:54:54 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/11 23:54:54 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/11 23:54:55 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:55436 (total: 1)
2025/09/11 23:54:55 server.go:164: [SOCKS5-CONN] New connection from [::1]:55436 at 2025-09-11 23:54:55
2025/09/11 23:54:55 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:55436
2025/09/11 23:54:55 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:55436
2025/09/11 23:54:55 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:55436
2025/09/11 23:54:55 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:54:55 client.go:98: url: ws://localhost:8080
2025/09/11 23:54:55 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:54:55 client.go:110: url: http://localhost:8080
2025/09/11 23:54:55 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[o0Yjq+iDOS2xwhRS1X/QcmgO44k=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:54:55 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:54:57 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:55505 (total: 2)
2025/09/11 23:54:57 server.go:164: [SOCKS5-CONN] New connection from [::1]:55505 at 2025-09-11 23:54:57
2025/09/11 23:54:57 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:55505
2025/09/11 23:54:57 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:55505
2025/09/11 23:54:57 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:55505
2025/09/11 23:54:57 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:54:57 client.go:98: url: ws://localhost:8080
2025/09/11 23:54:57 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:54:57 client.go:110: url: http://localhost:8080
2025/09/11 23:54:57 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[8qf/SMmm0Yat6TdA/NnoxXEzC6w=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:54:57 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:54:57 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:55511 (total: 3)
2025/09/11 23:54:57 server.go:164: [SOCKS5-CONN] New connection from [::1]:55511 at 2025-09-11 23:54:57
2025/09/11 23:54:57 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:55511
2025/09/11 23:54:57 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:55511
2025/09/11 23:54:57 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:55511
2025/09/11 23:54:57 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/11 23:54:57 client.go:98: url: ws://localhost:8080
2025/09/11 23:54:57 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:54:57 client.go:110: url: http://localhost:8080
2025/09/11 23:54:57 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[ZRHN+wSlmB4E2AiO9lYadkzUnbU=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:54:57 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/11 23:54:57 selector.go:270: WebSocket forward data error: read tcp [::1]:55437->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:54:57 selector.go:270: WebSocket forward data error: read tcp [::1]:55514->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:54:57 selector.go:270: WebSocket forward data error: read tcp [::1]:55508->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:54:57 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:55511, duration: 152.4798ms
2025/09/11 23:54:57 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:55511
2025/09/11 23:54:57 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:55505, duration: 234.8067ms
2025/09/11 23:54:57 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:55505
2025/09/11 23:54:57 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:55436, duration: 2.3166689s
2025/09/11 23:54:57 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:55436
```

✅ 端口8080已成功释放 ✅ 端口10810已成功释放

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:55:01

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 51728

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 44664

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 51660

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:54 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11458504935299255454
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:54 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11458504935299255454
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 49936

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:54 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11937286947190974553
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:54 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11937286947190974553
```

### 📋 所有进程PID记录

所有进程PID: 51728, 44664, 51660, 49936

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:55:01

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 51728

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 44664

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 51660

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:54 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11458504935299255454
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:54 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11458504935299255454
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 49936

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:54:54 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11937286947190974553
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:54:54 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11937286947190974553
```

### 📋 所有进程PID记录

所有进程PID: 51728, 44664, 51660, 49936

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程... ✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程... ✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程... ✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件

### WebSocket服务器日志输出

```
2025/09/11 23:55:02 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/11 23:55:02 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/11 23:55:02 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:55:02 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/11 23:55:02 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/11 23:55:02 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/11 23:55:04 server.go:90: url /
2025/09/11 23:55:04 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[memLivL/1zCa76uzsWRSgA==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:55:04 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:55588 at 2025-09-11 23:55:04
2025/09/11 23:55:04 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:55588
2025/09/11 23:55:04 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:55:04 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:55:04 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:55588
2025/09/11 23:55:04 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:55588
2025/09/11 23:55:04 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:55:04 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:55:04 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:55:06 server.go:90: url /
2025/09/11 23:55:06 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[q11dgvbF46gJUb+sN42Rog==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:55:06 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:55595 at 2025-09-11 23:55:06
2025/09/11 23:55:06 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:55595
2025/09/11 23:55:06 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/11 23:55:06 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:55:06 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:55595
2025/09/11 23:55:06 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:55595
2025/09/11 23:55:06 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/11 23:55:06 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/11 23:55:06 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:55:06 server.go:90: url /
2025/09/11 23:55:06 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[4nrH96kCsSJ4P3/cZFFw/Q==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:55:06 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:55603 at 2025-09-11 23:55:06
2025/09/11 23:55:06 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:55603
2025/09/11 23:55:06 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/11 23:55:06 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:55:06 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:55603
2025/09/11 23:55:06 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:55603
2025/09/11 23:55:06 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/11 23:55:06 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/11 23:55:06 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/11 23:55:03 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/11 23:55:03 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/11 23:55:03 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/11 23:55:03 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:55:03 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/11 23:55:03 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/11 23:55:04 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:55587 (total: 1)
2025/09/11 23:55:04 server.go:164: [SOCKS5-CONN] New connection from [::1]:55587 at 2025-09-11 23:55:04
2025/09/11 23:55:04 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:55587
2025/09/11 23:55:04 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:55587
2025/09/11 23:55:04 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:55587
2025/09/11 23:55:04 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:55:04 client.go:98: url: ws://localhost:8080
2025/09/11 23:55:04 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:55:04 client.go:110: url: http://localhost:8080
2025/09/11 23:55:04 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[qNXQcWNdgKTdWmGlendSNdviX7Q=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:55:04 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:55:06 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:55592 (total: 2)
2025/09/11 23:55:06 server.go:164: [SOCKS5-CONN] New connection from [::1]:55592 at 2025-09-11 23:55:06
2025/09/11 23:55:06 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:55592
2025/09/11 23:55:06 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:55592
2025/09/11 23:55:06 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:55592
2025/09/11 23:55:06 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/11 23:55:06 client.go:98: url: ws://localhost:8080
2025/09/11 23:55:06 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/11 23:55:06 client.go:110: url: http://localhost:8080
2025/09/11 23:55:06 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[InOmZ9Xpz2w1gXLUvx24utPf7ds=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:55:06 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/11 23:55:06 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:55600 (total: 3)
2025/09/11 23:55:06 server.go:164: [SOCKS5-CONN] New connection from [::1]:55600 at 2025-09-11 23:55:06
2025/09/11 23:55:06 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:55600
2025/09/11 23:55:06 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:55600
2025/09/11 23:55:06 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:55600
2025/09/11 23:55:06 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/11 23:55:06 client.go:98: url: ws://localhost:8080
2025/09/11 23:55:06 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/11 23:55:06 client.go:110: url: http://localhost:8080
2025/09/11 23:55:06 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[6adKfbE7KsHhvtm8Cje6wb1aU+A=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:55:06 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/11 23:55:06 selector.go:270: WebSocket forward data error: read tcp [::1]:55595->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:55:06 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:55592, duration: 159.2939ms
2025/09/11 23:55:06 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:55592
2025/09/11 23:55:06 selector.go:270: WebSocket forward data error: read tcp [::1]:55588->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:55:06 selector.go:270: WebSocket forward data error: read tcp [::1]:55603->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:55:06 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:55587, duration: 2.242579s
2025/09/11 23:55:06 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:55587
2025/09/11 23:55:06 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:55600, duration: 88.9791ms
2025/09/11 23:55:06 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:55600
```

✅ 端口8080已成功释放 ✅ 端口10810已成功释放

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:58:03

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 43852

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 30284

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 54904

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:57:56 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11488254279561136173
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:57:56 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11488254279561136173
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 47424

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:57:57 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_7607454513343109295
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:57:57 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_7607454513343109295
```

### 📋 所有进程PID记录

所有进程PID: 43852, 30284, 54904, 47424

### 

# WebSocket和SOCKS5级联代理测试记录

## 测试时间

2025-09-11 23:58:03

## 1. 编译代理服务器

执行命令: `go build -o main.exe ./cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 43852

等待WebSocket服务器启动... ✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令:
`./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 30284

等待SOCKS5服务器启动... ✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 54904

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:57:56 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11488254279561136173
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:57:56 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11488254279561136173
```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 47424

✅ 测试成功

输出结果:

```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 15:57:57 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_7607454513343109295
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 15:57:57 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_7607454513343109295
```

### 📋 所有进程PID记录

所有进程PID: 43852, 30284, 54904, 47424

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程... ✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程... ✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程... ✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件

### WebSocket服务器日志输出

```
2025/09/11 23:58:04 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/11 23:58:04 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/11 23:58:04 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:58:04 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/11 23:58:04 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/11 23:58:04 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/11 23:58:06 server.go:90: url /
2025/09/11 23:58:06 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[WuFkTeNk3L4QlYiecs4kaA==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/11 23:58:06 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:58663 at 2025-09-11 23:58:06
2025/09/11 23:58:06 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:58663
2025/09/11 23:58:06 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.51.73', targetPort: 80
2025/09/11 23:58:06 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:58:06 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:58663
2025/09/11 23:58:06 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.51.73:80 from [::1]:58663
2025/09/11 23:58:06 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.51.73:80 (timeout: 30s)
2025/09/11 23:58:06 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.51.73:80
2025/09/11 23:58:06 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:58:08 server.go:90: url /
2025/09/11 23:58:08 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[SF704R+AowNiqsLOLXfiHg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/11 23:58:08 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:58681 at 2025-09-11 23:58:08
2025/09/11 23:58:08 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:58681
2025/09/11 23:58:08 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.51.73', targetPort: 80
2025/09/11 23:58:08 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:58:08 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:58681
2025/09/11 23:58:08 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.51.73:80 from [::1]:58681
2025/09/11 23:58:08 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.51.73:80 (timeout: 30s)
2025/09/11 23:58:08 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.51.73:80
2025/09/11 23:58:08 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/11 23:58:08 server.go:90: url /
2025/09/11 23:58:08 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[TVtRf9i7AZyVCwz4fgMkDg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[443]]
2025/09/11 23:58:08 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:58686 at 2025-09-11 23:58:08
2025/09/11 23:58:08 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:58686
2025/09/11 23:58:08 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.51.73', targetPort: 443
2025/09/11 23:58:08 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/11 23:58:08 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:58686
2025/09/11 23:58:08 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.51.73:443 from [::1]:58686
2025/09/11 23:58:08 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.51.73:443 (timeout: 30s)
2025/09/11 23:58:08 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.51.73:443
2025/09/11 23:58:08 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/11 23:58:05 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/11 23:58:05 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/11 23:58:05 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/11 23:58:05 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/11 23:58:05 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/11 23:58:05 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/11 23:58:06 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:58661 (total: 1)
2025/09/11 23:58:06 server.go:164: [SOCKS5-CONN] New connection from [::1]:58661 at 2025-09-11 23:58:06
2025/09/11 23:58:06 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:58661
2025/09/11 23:58:06 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:58661
2025/09/11 23:58:06 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:58661
2025/09/11 23:58:06 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.51.73:80
2025/09/11 23:58:06 client.go:98: url: ws://localhost:8080
2025/09/11 23:58:06 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/11 23:58:06 client.go:110: url: http://localhost:8080
2025/09/11 23:58:06 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[BYcqnj6JYkqtT59qGDb7NhTn8bw=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:58:06 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.51.73:80
2025/09/11 23:58:08 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:58678 (total: 2)
2025/09/11 23:58:08 server.go:164: [SOCKS5-CONN] New connection from [::1]:58678 at 2025-09-11 23:58:08
2025/09/11 23:58:08 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:58678
2025/09/11 23:58:08 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:58678
2025/09/11 23:58:08 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:58678
2025/09/11 23:58:08 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.51.73:80
2025/09/11 23:58:08 client.go:98: url: ws://localhost:8080
2025/09/11 23:58:08 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/11 23:58:08 client.go:110: url: http://localhost:8080
2025/09/11 23:58:08 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[xp2HXR6lzqr5IDgWFOtF20E9NCw=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:58:08 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.51.73:80
2025/09/11 23:58:08 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:58683 (total: 3)
2025/09/11 23:58:08 server.go:164: [SOCKS5-CONN] New connection from [::1]:58683 at 2025-09-11 23:58:08
2025/09/11 23:58:08 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:58683
2025/09/11 23:58:08 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:58683
2025/09/11 23:58:08 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:58683
2025/09/11 23:58:08 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.51.73:443
2025/09/11 23:58:08 client.go:98: url: ws://localhost:8080
2025/09/11 23:58:08 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[443]]
2025/09/11 23:58:08 client.go:110: url: http://localhost:8080
2025/09/11 23:58:08 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[9w513JV2no7jktFsssoKw2bZZFw=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/11 23:58:08 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.51.73:443
2025/09/11 23:58:09 selector.go:270: WebSocket forward data error: read tcp [::1]:58663->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:58:09 selector.go:270: WebSocket forward data error: read tcp [::1]:58681->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:58:09 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:58678, duration: 144.5892ms
2025/09/11 23:58:09 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:58661, duration: 2.2525952s
2025/09/11 23:58:09 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:58678
2025/09/11 23:58:09 selector.go:270: WebSocket forward data error: read tcp [::1]:58686->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/11 23:58:09 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:58661
2025/09/11 23:58:09 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:58683, duration: 80.3034ms
2025/09/11 23:58:09 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:58683
```

✅ 端口8080已成功释放 ✅ 端口10810已成功释放
