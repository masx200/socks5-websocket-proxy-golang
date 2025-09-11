

###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:05:06

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 50184

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 55208

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 53088

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
< Date: Thu, 11 Sep 2025 16:04:59 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11083680905935525707
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:04:59 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11083680905935525707


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 43424

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
< Date: Thu, 11 Sep 2025 16:04:59 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11340813506796900831
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:04:59 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11340813506796900831


```

### 📋 所有进程PID记录

所有进程PID: 50184, 55208, 53088, 43424



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:05:06

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 50184

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 55208

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 53088

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
< Date: Thu, 11 Sep 2025 16:04:59 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11083680905935525707
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:04:59 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11083680905935525707


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 43424

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
< Date: Thu, 11 Sep 2025 16:04:59 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11340813506796900831
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:04:59 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11340813506796900831


```

### 📋 所有进程PID记录

所有进程PID: 50184, 55208, 53088, 43424

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:05:06 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/12 00:05:06 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/12 00:05:06 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:05:06 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:05:06 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:05:06 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:05:09 server.go:90: url /
2025/09/12 00:05:09 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[9BCwR5LPOE1j8dazdf3m/g==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/12 00:05:09 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:62509 at 2025-09-12 00:05:09
2025/09/12 00:05:09 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:62509
2025/09/12 00:05:09 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.51.73', targetPort: 80
2025/09/12 00:05:09 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:05:09 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:62509
2025/09/12 00:05:09 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.51.73:80 from [::1]:62509
2025/09/12 00:05:09 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.51.73:80 (timeout: 30s)
2025/09/12 00:05:09 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.51.73:80
2025/09/12 00:05:09 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:05:11 server.go:90: url /
2025/09/12 00:05:11 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[FEhvLENRbokkeHoweCtIcA==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/12 00:05:11 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:62554 at 2025-09-12 00:05:11
2025/09/12 00:05:11 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:62554
2025/09/12 00:05:11 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.51.73', targetPort: 80
2025/09/12 00:05:11 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:05:11 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:62554
2025/09/12 00:05:11 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.51.73:80 from [::1]:62554
2025/09/12 00:05:11 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.51.73:80 (timeout: 30s)
2025/09/12 00:05:11 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.51.73:80
2025/09/12 00:05:11 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:05:11 server.go:90: url /
2025/09/12 00:05:11 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[gBcbrhupjHau4DhEXn0LIQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[443]]
2025/09/12 00:05:11 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:62559 at 2025-09-12 00:05:11
2025/09/12 00:05:11 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:62559
2025/09/12 00:05:11 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.51.73', targetPort: 443
2025/09/12 00:05:11 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:05:11 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:62559
2025/09/12 00:05:11 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.51.73:443 from [::1]:62559
2025/09/12 00:05:11 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.51.73:443 (timeout: 30s)
2025/09/12 00:05:11 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.51.73:443
2025/09/12 00:05:11 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:05:07 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/12 00:05:07 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/12 00:05:07 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/12 00:05:07 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:05:07 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/12 00:05:07 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/12 00:05:08 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:62508 (total: 1)
2025/09/12 00:05:08 server.go:164: [SOCKS5-CONN] New connection from [::1]:62508 at 2025-09-12 00:05:08
2025/09/12 00:05:08 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:62508
2025/09/12 00:05:08 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:62508
2025/09/12 00:05:08 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:62508
2025/09/12 00:05:09 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.51.73:80
2025/09/12 00:05:09 client.go:98: url: ws://localhost:8080
2025/09/12 00:05:09 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/12 00:05:09 client.go:110: url: http://localhost:8080
2025/09/12 00:05:09 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[gaw6Z5TFUjmqgFf9Yra4545YorE=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:05:09 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.51.73:80
2025/09/12 00:05:11 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:62551 (total: 2)
2025/09/12 00:05:11 server.go:164: [SOCKS5-CONN] New connection from [::1]:62551 at 2025-09-12 00:05:11
2025/09/12 00:05:11 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:62551
2025/09/12 00:05:11 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:62551
2025/09/12 00:05:11 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:62551
2025/09/12 00:05:11 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.51.73:80
2025/09/12 00:05:11 client.go:98: url: ws://localhost:8080
2025/09/12 00:05:11 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[80]]
2025/09/12 00:05:11 client.go:110: url: http://localhost:8080
2025/09/12 00:05:11 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[8qESUsyt7LDAoboEijr8cWxadjk=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:05:11 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.51.73:80
2025/09/12 00:05:11 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:62556 (total: 3)
2025/09/12 00:05:11 server.go:164: [SOCKS5-CONN] New connection from [::1]:62556 at 2025-09-12 00:05:11
2025/09/12 00:05:11 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:62556
2025/09/12 00:05:11 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:62556
2025/09/12 00:05:11 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:62556
2025/09/12 00:05:11 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.51.73:443
2025/09/12 00:05:11 client.go:98: url: ws://localhost:8080
2025/09/12 00:05:11 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.51.73] X-Proxy-Target-Port:[443]]
2025/09/12 00:05:11 client.go:110: url: http://localhost:8080
2025/09/12 00:05:11 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[D7M4PahU1W2Cquk86bUrR8DQdL0=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:05:11 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.51.73:443
2025/09/12 00:05:11 selector.go:270: WebSocket forward data error: read tcp [::1]:62509->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:05:11 selector.go:270: WebSocket forward data error: read tcp [::1]:62554->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:05:11 selector.go:270: WebSocket forward data error: read tcp [::1]:62559->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:05:11 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:62551, duration: 146.1329ms
2025/09/12 00:05:11 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:62551
2025/09/12 00:05:11 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:62556, duration: 80.8832ms
2025/09/12 00:05:11 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:62508, duration: 2.2240168s
2025/09/12 00:05:11 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:62556
2025/09/12 00:05:11 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:62508
```

✅ 端口8080已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:09:16

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 34444

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 28160

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 6168, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:09 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11577489015499807356
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:09 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11577489015499807356


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 12924, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:10 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10715536817220349730
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:10 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10715536817220349730


```

### 📋 所有进程PID记录

所有进程PID: 34444, 28160, 6168, 12924



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:09:16

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 34444

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 28160

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 6168, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:09 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11577489015499807356
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:09 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11577489015499807356


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 12924, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:10 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10715536817220349730
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:10 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10715536817220349730


```

### 📋 所有进程PID记录

所有进程PID: 34444, 28160, 6168, 12924

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:09:17 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/12 00:09:17 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/12 00:09:17 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:09:17 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:09:17 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:09:17 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:09:19 server.go:90: url /
2025/09/12 00:09:19 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[fjL/Qz/EBRptH2bnOszDpw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:19 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65047 at 2025-09-12 00:09:19
2025/09/12 00:09:19 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65047
2025/09/12 00:09:19 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/12 00:09:19 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:09:19 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65047
2025/09/12 00:09:19 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:65047
2025/09/12 00:09:19 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/12 00:09:19 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/12 00:09:19 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:09:21 server.go:90: url /
2025/09/12 00:09:21 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[ni9taqf2rVpSEhRyUEadkw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:21 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65069 at 2025-09-12 00:09:21
2025/09/12 00:09:21 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65069
2025/09/12 00:09:21 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/12 00:09:21 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:09:21 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65069
2025/09/12 00:09:21 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:65069
2025/09/12 00:09:21 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/12 00:09:21 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/12 00:09:21 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:09:21 server.go:90: url /
2025/09/12 00:09:21 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[37v2hMWSHfLuGNfKmeDGcQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/12 00:09:21 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65075 at 2025-09-12 00:09:21
2025/09/12 00:09:21 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65075
2025/09/12 00:09:21 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/12 00:09:21 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:09:21 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65075
2025/09/12 00:09:21 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:65075
2025/09/12 00:09:21 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/12 00:09:21 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/12 00:09:21 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:09:18 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/12 00:09:18 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/12 00:09:18 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/12 00:09:18 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:09:18 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/12 00:09:18 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/12 00:09:19 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:65046 (total: 1)
2025/09/12 00:09:19 server.go:164: [SOCKS5-CONN] New connection from [::1]:65046 at 2025-09-12 00:09:19
2025/09/12 00:09:19 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:65046
2025/09/12 00:09:19 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:65046
2025/09/12 00:09:19 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:65046
2025/09/12 00:09:19 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/12 00:09:19 client.go:98: url: ws://localhost:8080
2025/09/12 00:09:19 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:19 client.go:110: url: http://localhost:8080
2025/09/12 00:09:19 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[Xl4qZX7wexBlIVZnI7tr0Ek8ibw=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:09:19 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/12 00:09:21 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:65066 (total: 2)
2025/09/12 00:09:21 server.go:164: [SOCKS5-CONN] New connection from [::1]:65066 at 2025-09-12 00:09:21
2025/09/12 00:09:21 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:65066
2025/09/12 00:09:21 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:65066
2025/09/12 00:09:21 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:65066
2025/09/12 00:09:21 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/12 00:09:21 client.go:98: url: ws://localhost:8080
2025/09/12 00:09:21 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:21 client.go:110: url: http://localhost:8080
2025/09/12 00:09:21 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[2xBopePMMmeyTQVAghIgInOTxjM=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:09:21 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/12 00:09:21 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:65072 (total: 3)
2025/09/12 00:09:21 server.go:164: [SOCKS5-CONN] New connection from [::1]:65072 at 2025-09-12 00:09:21
2025/09/12 00:09:21 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:65072
2025/09/12 00:09:21 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:65072
2025/09/12 00:09:21 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:65072
2025/09/12 00:09:21 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/12 00:09:21 client.go:98: url: ws://localhost:8080
2025/09/12 00:09:21 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/12 00:09:21 client.go:110: url: http://localhost:8080
2025/09/12 00:09:21 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[76YwdOQiKQIv/OZcYe18J0dZwtE=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:09:21 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/12 00:09:22 selector.go:270: WebSocket forward data error: read tcp [::1]:65047->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:09:22 selector.go:270: WebSocket forward data error: read tcp [::1]:65069->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:09:22 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:65046, duration: 2.3549947s
2025/09/12 00:09:22 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:65046
2025/09/12 00:09:22 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:65066, duration: 241.9348ms
2025/09/12 00:09:22 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:65066
2025/09/12 00:09:22 selector.go:270: WebSocket forward data error: read tcp [::1]:65075->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:09:22 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:65072, duration: 156.6334ms
2025/09/12 00:09:22 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:65072
```

✅ 端口8080已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:09:39

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 3844

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 32872

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 45212, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10700276317694402791
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10700276317694402791


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 55236, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11347669773719220544
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11347669773719220544


```

### 📋 所有进程PID记录

所有进程PID: 3844, 32872, 45212, 55236



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:09:39

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 3844

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 32872

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 45212, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10700276317694402791
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10700276317694402791


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 55236, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:09:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11347669773719220544
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:09:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11347669773719220544


```

### 📋 所有进程PID记录

所有进程PID: 3844, 32872, 45212, 55236

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:09:40 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/12 00:09:40 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/12 00:09:40 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:09:40 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:09:40 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:09:40 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:09:42 server.go:90: url /
2025/09/12 00:09:42 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[PAgl4z54vEw5slxa+Dxolw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:42 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65334 at 2025-09-12 00:09:42
2025/09/12 00:09:42 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65334
2025/09/12 00:09:42 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/12 00:09:42 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:09:42 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65334
2025/09/12 00:09:42 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:65334
2025/09/12 00:09:42 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/12 00:09:42 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/12 00:09:42 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:09:44 server.go:90: url /
2025/09/12 00:09:44 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[GCNbdUkoz3+W5WHIVxTpiQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:44 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65340 at 2025-09-12 00:09:44
2025/09/12 00:09:44 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65340
2025/09/12 00:09:44 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/12 00:09:44 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:09:44 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65340
2025/09/12 00:09:44 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:65340
2025/09/12 00:09:44 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/12 00:09:44 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/12 00:09:44 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:09:44 server.go:90: url /
2025/09/12 00:09:44 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[XMaxLQwqPigbj32UH0557Q==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/12 00:09:44 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65346 at 2025-09-12 00:09:44
2025/09/12 00:09:44 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65346
2025/09/12 00:09:44 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/12 00:09:44 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:09:44 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65346
2025/09/12 00:09:44 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:65346
2025/09/12 00:09:44 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/12 00:09:44 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/12 00:09:44 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:09:41 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/12 00:09:41 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/12 00:09:41 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/12 00:09:41 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:09:41 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/12 00:09:41 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/12 00:09:42 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:65333 (total: 1)
2025/09/12 00:09:42 server.go:164: [SOCKS5-CONN] New connection from [::1]:65333 at 2025-09-12 00:09:42
2025/09/12 00:09:42 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:65333
2025/09/12 00:09:42 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:65333
2025/09/12 00:09:42 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:65333
2025/09/12 00:09:42 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/12 00:09:42 client.go:98: url: ws://localhost:8080
2025/09/12 00:09:42 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:42 client.go:110: url: http://localhost:8080
2025/09/12 00:09:42 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[itGPRPBMyq5smm8jJU57udBDPBc=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:09:42 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/12 00:09:44 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:65337 (total: 2)
2025/09/12 00:09:44 server.go:164: [SOCKS5-CONN] New connection from [::1]:65337 at 2025-09-12 00:09:44
2025/09/12 00:09:44 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:65337
2025/09/12 00:09:44 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:65337
2025/09/12 00:09:44 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:65337
2025/09/12 00:09:44 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/12 00:09:44 client.go:98: url: ws://localhost:8080
2025/09/12 00:09:44 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:09:44 client.go:110: url: http://localhost:8080
2025/09/12 00:09:44 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[W7h19Sq3VBNBKxI78qRSFYl6s7Q=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:09:44 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/12 00:09:44 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:65343 (total: 3)
2025/09/12 00:09:44 server.go:164: [SOCKS5-CONN] New connection from [::1]:65343 at 2025-09-12 00:09:44
2025/09/12 00:09:44 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:65343
2025/09/12 00:09:44 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:65343
2025/09/12 00:09:44 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:65343
2025/09/12 00:09:44 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/12 00:09:44 client.go:98: url: ws://localhost:8080
2025/09/12 00:09:44 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/12 00:09:44 client.go:110: url: http://localhost:8080
2025/09/12 00:09:44 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[zOdQTzyABiqdnYHUJVXfKU+1mL0=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:09:44 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/12 00:09:44 selector.go:270: WebSocket forward data error: read tcp [::1]:65334->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:09:44 selector.go:270: WebSocket forward data error: read tcp [::1]:65340->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:09:44 selector.go:270: WebSocket forward data error: read tcp [::1]:65346->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:09:44 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:65333, duration: 2.2491441s
2025/09/12 00:09:44 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:65333
2025/09/12 00:09:44 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:65337, duration: 171.7239ms
2025/09/12 00:09:44 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:65337
2025/09/12 00:09:44 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:65343, duration: 96.3115ms
2025/09/12 00:09:44 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:65343
```

✅ 端口8080已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:11:21

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 45164

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 54728

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 49848, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:11:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11221686880848611684
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:11:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11221686880848611684


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 41324, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:11:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11424520220313549222
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:11:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11424520220313549222


```

### 📋 所有进程PID记录

所有进程PID: 45164, 54728, 49848, 41324



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:11:21

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./main.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 45164

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe -mode server -protocol socks5 -addr :10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 54728

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试1进程PID: 49848, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:11:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11221686880848611684
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:11:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11221686880848611684


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x socks5://localhost:10810`

📋 Curl测试2进程PID: 41324, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 16:11:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11424520220313549222
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:11:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11424520220313549222


```

### 📋 所有进程PID记录

所有进程PID: 45164, 54728, 49848, 41324

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:11:22 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/12 00:11:22 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/12 00:11:22 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:11:22 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:11:22 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:11:22 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:11:24 server.go:90: url /
2025/09/12 00:11:24 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[SU+6UoFYRKhHRgAjgNxhuw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:11:24 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:1457 at 2025-09-12 00:11:24
2025/09/12 00:11:24 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:1457
2025/09/12 00:11:24 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/12 00:11:24 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:11:24 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:1457
2025/09/12 00:11:24 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:1457
2025/09/12 00:11:24 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/12 00:11:24 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/12 00:11:24 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:11:26 server.go:90: url /
2025/09/12 00:11:26 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[InSDjMl8kiDde4CE9WBS4A==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:11:26 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:1486 at 2025-09-12 00:11:26
2025/09/12 00:11:26 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:1486
2025/09/12 00:11:26 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 80
2025/09/12 00:11:26 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:11:26 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:1486
2025/09/12 00:11:26 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:80 from [::1]:1486
2025/09/12 00:11:26 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:80 (timeout: 30s)
2025/09/12 00:11:26 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:80
2025/09/12 00:11:26 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:11:26 server.go:90: url /
2025/09/12 00:11:26 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[KUYsUclm+39KbMlqlbrvKQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/12 00:11:26 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:1492 at 2025-09-12 00:11:26
2025/09/12 00:11:26 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:1492
2025/09/12 00:11:26 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '180.101.49.44', targetPort: 443
2025/09/12 00:11:26 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:11:26 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:1492
2025/09/12 00:11:26 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 180.101.49.44:443 from [::1]:1492
2025/09/12 00:11:26 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 180.101.49.44:443 (timeout: 30s)
2025/09/12 00:11:26 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 180.101.49.44:443
2025/09/12 00:11:26 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:11:23 main.go:71: 启动socks5服务端，监听地址: :10810
2025/09/12 00:11:23 main.go:105: 使用命令行参数设置上游配置: 类型=websocket, 地址=ws://localhost:8080
2025/09/12 00:11:23 server.go:108: [SOCKS5-SERVER] Server started successfully, listening on :10810
2025/09/12 00:11:23 server.go:109: [SOCKS5-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:11:23 server.go:111: [SOCKS5-SERVER] Upstream selector enabled: true
2025/09/12 00:11:23 main.go:129: socks5服务端已启动，按Ctrl+C停止
2025/09/12 00:11:24 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:1456 (total: 1)
2025/09/12 00:11:24 server.go:164: [SOCKS5-CONN] New connection from [::1]:1456 at 2025-09-12 00:11:24
2025/09/12 00:11:24 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:1456
2025/09/12 00:11:24 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:1456
2025/09/12 00:11:24 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:1456
2025/09/12 00:11:24 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/12 00:11:24 client.go:98: url: ws://localhost:8080
2025/09/12 00:11:24 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:11:24 client.go:110: url: http://localhost:8080
2025/09/12 00:11:24 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[bRdeU2Z+6b2XuMXK5UnWGqhJgw8=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:11:24 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/12 00:11:26 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:1483 (total: 2)
2025/09/12 00:11:26 server.go:164: [SOCKS5-CONN] New connection from [::1]:1483 at 2025-09-12 00:11:26
2025/09/12 00:11:26 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:1483
2025/09/12 00:11:26 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:1483
2025/09/12 00:11:26 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:1483
2025/09/12 00:11:26 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:80
2025/09/12 00:11:26 client.go:98: url: ws://localhost:8080
2025/09/12 00:11:26 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[80]]
2025/09/12 00:11:26 client.go:110: url: http://localhost:8080
2025/09/12 00:11:26 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[kJC+N1PCmPkesL6IITXyY3TkknM=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:11:26 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:80
2025/09/12 00:11:26 server.go:143: [SOCKS5-SERVER] New connection accepted from [::1]:1489 (total: 3)
2025/09/12 00:11:26 server.go:164: [SOCKS5-CONN] New connection from [::1]:1489 at 2025-09-12 00:11:26
2025/09/12 00:11:26 server.go:170: [SOCKS5-CONN] Connection timeout set to 30s for client [::1]:1489
2025/09/12 00:11:26 server.go:177: [SOCKS5-CONN] No authentication required for client [::1]:1489
2025/09/12 00:11:26 server.go:182: [SOCKS5-CONN] Connection details - Local: [::1]:10810, Remote: [::1]:1489
2025/09/12 00:11:26 server.go:223: [SOCKS5-UPSTREAM] Using upstream selector for target 180.101.49.44:443
2025/09/12 00:11:26 client.go:98: url: ws://localhost:8080
2025/09/12 00:11:26 client.go:99: headers: map[X-Proxy-Target-Host:[180.101.49.44] X-Proxy-Target-Port:[443]]
2025/09/12 00:11:26 client.go:110: url: http://localhost:8080
2025/09/12 00:11:26 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[glodA2YvafhErcoo2Xk0OzqVR0U=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:11:26 server.go:239: [SOCKS5-UPSTREAM] Upstream selector succeeded for target 180.101.49.44:443
2025/09/12 00:11:26 selector.go:270: WebSocket forward data error: read tcp [::1]:1486->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:11:26 selector.go:270: WebSocket forward data error: read tcp [::1]:1457->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:11:26 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:1483, duration: 143.7822ms
2025/09/12 00:11:26 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:1483
2025/09/12 00:11:26 selector.go:270: WebSocket forward data error: read tcp [::1]:1492->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:11:26 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:1456, duration: 2.2149127s
2025/09/12 00:11:26 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:1456
2025/09/12 00:11:26 server.go:191: [SOCKS5-CONN] Connection completed successfully for client [::1]:1489, duration: 77.4462ms
2025/09/12 00:11:26 server.go:159: [SOCKS5-CONN] Connection closed for client [::1]:1489
```

✅ 端口8080已成功释放
✅ 端口10810已成功释放
