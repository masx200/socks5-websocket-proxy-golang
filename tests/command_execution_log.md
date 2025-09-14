=== 命令执行日志文件 === 创建时间: 2025-09-14 15:35:43

开始运行命令... [2025-09-14 15:35:43.481] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd 开始运行命令... [2025-09-14
15:35:43.481] [BUILD] C:\Program Files\Go\bin\go.exe go build -o
socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 3196
执行时间: 1.40625s
---
```

开始运行命令... [2025-09-14 15:35:44.382] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go 开始运行命令...
[2025-09-14 15:35:44.383] [BUILD] C:\Program Files\Go\bin\go.exe go build -o
main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 12568
执行时间: 1.265625s
---
```

开始运行命令... [2025-09-14 15:35:44.980] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:38800 开始运行命令... [2025-09-14 15:35:44.980] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:38800

```
执行结果: 成功
进程PID: 9116
---
```

开始运行命令... [2025-09-14 15:35:45.985] [执行命令] ./main.exe -listen-port
10800 -upstream-type websocket -upstream-address ws://localhost:38800
开始运行命令... [2025-09-14 15:35:45.985] [HTTP] ./main.exe -listen-port 10800
-upstream-type websocket -upstream-address ws://localhost:38800

```
执行结果: 成功
进程PID: 4088
---
```

开始运行命令... [2025-09-14 15:35:49.123] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:35:49.123] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I http://www.baidu.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 18640
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* Connected to localhost (::1) port 10800
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
< Date: Sun, 14 Sep 2025 07:35:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11365344504782341765
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:35:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11365344504782341765
---
```

开始运行命令... [2025-09-14 15:35:49.244] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:35:49.245] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I https://www.baidu.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 33760
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:443 (locally resolved)
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* SOCKS5 request granted.
* Connected to localhost () port 10800
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
* Connected to localhost (::1) port 10800
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
< Date: Sun, 14 Sep 2025 07:35:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9099951637799389841
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:35:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9099951637799389841
---
```

开始运行命令... [2025-09-14 15:35:53.789] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go 开始运行命令...
[2025-09-14 15:35:53.789] [BUILD] C:\Program Files\Go\bin\go.exe go build -o
main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 12088
执行时间: 1.296875s
---
```

开始运行命令... [2025-09-14 15:35:54.394] [执行命令] ./main.exe -listen-port
10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6
-dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 开始运行命令...
[2025-09-14 15:35:54.397] [SERVER] ./main.exe -listen-port 10800 -dohurl
https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl
https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 开始运行命令...
[2025-09-14 15:35:57.566] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:35:57.566] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I http://www.baidu.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 16652
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* Connected to localhost (::1) port 10800
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
< Date: Sun, 14 Sep 2025 07:35:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9909554112435744668
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:35:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9909554112435744668
---
```

开始运行命令... [2025-09-14 15:35:57.691] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:35:57.691] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I -L http://www.so.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 19092
执行时间: 78.125ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.so.com:80 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* Connected to localhost (::1) port 10800
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 302 Moved Temporarily
< Server: openresty
< Date: Sun, 14 Sep 2025 07:35:40 GMT
< Content-Type: text/html
< Connection: keep-alive
< Location: https://www.so.com/
< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/
* Ignoring the response-body
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
* Clear auth, redirects to port from 80 to 443
* Issue another request to this URL: 'https://www.so.com/'
* Hostname localhost was found in DNS cache
*   Trying [::1]:10800...
* Host www.so.com:443 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [305 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [93 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [5077 bytes data]
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
* ALPN: server did not agree on a protocol. Uses default.
* Server certificate:
*  subject: CN=*.so.com
*  start date: Aug 28 00:00:00 2025 GMT
*  expire date: Sep 28 23:59:59 2026 GMT
*  subjectAltName: host "www.so.com" matched cert's "*.so.com"
*  issuer: C=CN; O=WoTrus CA Limited; CN=WoTrus DV Server CA  [Run by the Issuer]
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha384WithRSAEncryption
*   Certificate level 2: Public key type ? (4096/128 Bits/secBits), signed using sha384WithRSAEncryption
* Connected to localhost (::1) port 10800
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 07:35:40 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=dbnnaali9g2aejb5383ip359o2; expires=Sun, 14-Sep-2025 07:45:40 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=A004C70D44916A1B70EFA1534685E535.1757835340930; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:35:40 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:35:40 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=dbnnaali9g2aejb5383ip359o2; expires=Sun, 14-Sep-2025 07:45:40 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=A004C70D44916A1B70EFA1534685E535.1757835340930; Max-Age=63072000; Domain=so.com; Path=/
---
```

开始运行命令... [2025-09-14 15:35:57.893] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:35:57.893] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I https://www.baidu.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 5452
执行时间: 62.5ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
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
* Connected to localhost (::1) port 10800
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
< Date: Sun, 14 Sep 2025 07:35:42 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10808763743021232895
< 
  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:35:42 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10808763743021232895
---
```

开始运行命令... [2025-09-14 15:36:05.626] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go 开始运行命令...
[2025-09-14 15:36:05.626] [BUILD] C:\Program Files\Go\bin\go.exe go build -o
main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 31700
执行时间: 1.21875s
---
```

开始运行命令... [2025-09-14 15:36:06.205] [执行命令] ./main.exe -listen-port
10800 开始运行命令... [2025-09-14 15:36:06.210] [SERVER] ./main.exe -listen-port
10800 开始运行命令... [2025-09-14 15:36:08.298] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:36:08.298] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I http://www.baidu.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 30756
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* Connected to localhost (::1) port 10800
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 302 Found
< Location: https://www.baidu.com/error.html
< Server: bfe
< Date: Sun, 14 Sep 2025 07:35:51 GMT
< Content-Type: text/plain; charset=utf-8
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 302 Found
Location: https://www.baidu.com/error.html
Server: bfe
Date: Sun, 14 Sep 2025 07:35:51 GMT
Content-Type: text/plain; charset=utf-8
---
```

开始运行命令... [2025-09-14 15:36:08.384] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:36:08.384] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I -L http://www.so.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 20672
执行时间: 78.125ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.so.com:80 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* Connected to localhost (::1) port 10800
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 302 Moved Temporarily
< Server: openresty
< Date: Sun, 14 Sep 2025 07:35:51 GMT
< Content-Type: text/html
< Connection: keep-alive
< Location: https://www.so.com/
< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/
* Ignoring the response-body
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
* Clear auth, redirects to port from 80 to 443
* Issue another request to this URL: 'https://www.so.com/'
* Hostname localhost was found in DNS cache
*   Trying [::1]:10800...
* Host www.so.com:443 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
* using HTTP/1.x
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [305 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [93 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [5077 bytes data]
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
* ALPN: server did not agree on a protocol. Uses default.
* Server certificate:
*  subject: CN=*.so.com
*  start date: Aug 28 00:00:00 2025 GMT
*  expire date: Sep 28 23:59:59 2026 GMT
*  subjectAltName: host "www.so.com" matched cert's "*.so.com"
*  issuer: C=CN; O=WoTrus CA Limited; CN=WoTrus DV Server CA  [Run by the Issuer]
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha384WithRSAEncryption
*   Certificate level 2: Public key type ? (4096/128 Bits/secBits), signed using sha384WithRSAEncryption
* Connected to localhost (::1) port 10800
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 07:35:51 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=thirk9air48nho2r3s5a9ahf56; expires=Sun, 14-Sep-2025 07:45:51 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=577175C03B9CA17063BA812AB06FD47F.1757835351641; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:35:51 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:35:51 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=thirk9air48nho2r3s5a9ahf56; expires=Sun, 14-Sep-2025 07:45:51 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=577175C03B9CA17063BA812AB06FD47F.1757835351641; Max-Age=63072000; Domain=so.com; Path=/
---
```

开始运行命令... [2025-09-14 15:36:08.584] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x socks5://localhost:10800 开始运行命令... [2025-09-14
15:36:08.584] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I https://www.baidu.com -x socks5://localhost:10800

```
执行结果: 成功
进程PID: 21284
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1ac:0:ff:b07e:36c5, 240e:e9:6002:1fd:0:ff:b0e1:fe69
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10800
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
* Connected to localhost (::1) port 10800
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
< Date: Sun, 14 Sep 2025 07:35:51 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11049126764577581673
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:35:51 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11049126764577581673
---
```

开始运行命令... [2025-09-14 15:36:08.816] [执行命令]
C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe 开始运行命令...
[2025-09-14 15:36:08.817] [SYSTEM] C:\Windows\System32\taskkill.exe taskkill /F
/IM go.exe
