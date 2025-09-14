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
开始运行命令...
[2025-09-14 15:38:17.100] [执行命令] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
开始运行命令...
[2025-09-14 15:38:17.100] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
```
执行结果: 成功
进程PID: 30492
执行时间: 1.546875s
---
```
开始运行命令...
[2025-09-14 15:38:18.443] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:38:18.443] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 12088
执行时间: 921.875ms
---
```
开始运行命令...
[2025-09-14 15:38:18.960] [执行命令] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
开始运行命令...
[2025-09-14 15:38:18.960] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
```
执行结果: 成功
进程PID: 31224
---
```
开始运行命令...
[2025-09-14 15:38:19.966] [执行命令] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
开始运行命令...
[2025-09-14 15:38:19.981] [HTTP] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
```
执行结果: 成功
进程PID: 15572
---
```
开始运行命令...
[2025-09-14 15:38:23.114] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:23.114] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 2512
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:06 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11073579306158350275
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:06 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11073579306158350275
---
```
开始运行命令...
[2025-09-14 15:38:23.219] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:23.219] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 5764
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
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
< Date: Sun, 14 Sep 2025 07:38:06 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11997290771477372611
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:06 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11997290771477372611
---
```
开始运行命令...
[2025-09-14 15:38:27.720] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:38:27.720] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 31448
执行时间: 1.328125s
---
```
开始运行命令...
[2025-09-14 15:38:28.284] [执行命令] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:38:28.290] [SERVER] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:38:30.472] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:30.472] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 32472
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:13 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11649120965597767094
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:13 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11649120965597767094
---
```
开始运行命令...
[2025-09-14 15:38:30.578] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:30.578] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 2404
执行时间: 46.875ms
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
< Date: Sun, 14 Sep 2025 07:38:13 GMT
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
< Date: Sun, 14 Sep 2025 07:38:13 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=6ocvb74lbhjsbbji0pp867bl66; expires=Sun, 14-Sep-2025 07:48:13 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=E9CD7692CB0FBC33D53E81D9207A851C.1757835493919; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:38:13 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:38:13 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=6ocvb74lbhjsbbji0pp867bl66; expires=Sun, 14-Sep-2025 07:48:13 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=E9CD7692CB0FBC33D53E81D9207A851C.1757835493919; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:38:30.853] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:30.854] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 29812
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9449282112370809853
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9449282112370809853
---
```
开始运行命令...
[2025-09-14 15:38:37.947] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:38:37.947] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 13360
执行时间: 1.21875s
---
```
开始运行命令...
[2025-09-14 15:38:38.578] [执行命令] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:38:38.584] [SERVER] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:38:41.717] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:41.717] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 33176
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:24 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11054057142219905570
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:24 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11054057142219905570
---
```
开始运行命令...
[2025-09-14 15:38:41.866] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:41.866] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 13560
执行时间: 15.625ms
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
< Date: Sun, 14 Sep 2025 07:38:25 GMT
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
< Date: Sun, 14 Sep 2025 07:38:25 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=a6nrfi92k92veichv7sltjfcg4; expires=Sun, 14-Sep-2025 07:48:25 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=27DDEFA74E4BC35C8128F199B3BC2480.1757835505261; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:38:25 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:38:25 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=a6nrfi92k92veichv7sltjfcg4; expires=Sun, 14-Sep-2025 07:48:25 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=27DDEFA74E4BC35C8128F199B3BC2480.1757835505261; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:38:42.186] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:42.186] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 18536
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0> HEAD / HTTP/1.1
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
< Date: Sun, 14 Sep 2025 07:38:25 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11194744388054521339
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:25 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11194744388054521339
---
```
开始运行命令...
[2025-09-14 15:38:42.403] [执行命令] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:38:42.403] [SYSTEM] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:38:47.505] [执行命令] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
开始运行命令...
[2025-09-14 15:38:47.505] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
```
执行结果: 成功
进程PID: 33060
执行时间: 1.40625s
---
```
开始运行命令...
[2025-09-14 15:38:48.372] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:38:48.372] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 27456
执行时间: 953.125ms
---
```
开始运行命令...
[2025-09-14 15:38:48.913] [执行命令] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
开始运行命令...
[2025-09-14 15:38:48.913] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
```
执行结果: 成功
进程PID: 11928
---
```
开始运行命令...
[2025-09-14 15:38:49.919] [执行命令] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
开始运行命令...
[2025-09-14 15:38:49.919] [HTTP] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
```
执行结果: 成功
进程PID: 2116
---
```
开始运行命令...
[2025-09-14 15:38:52.040] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:52.040] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 872
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:36 GMT
< Content-Type: text/plain; charset=utf-8
< 
  0     0    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0  0     0    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 302 Found
Location: https://www.baidu.com/error.html
Server: bfe
Date: Sun, 14 Sep 2025 07:38:36 GMT
Content-Type: text/plain; charset=utf-8
---
```
开始运行命令...
[2025-09-14 15:38:53.193] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:38:53.193] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 30356
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
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
< Date: Sun, 14 Sep 2025 07:38:36 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9482658131381401268
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:36 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9482658131381401268
---
```
开始运行命令...
[2025-09-14 15:38:57.586] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:38:57.587] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 20988
执行时间: 1.25s
---
```
开始运行命令...
[2025-09-14 15:38:58.142] [执行命令] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:38:58.148] [SERVER] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:39:01.290] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:01.290] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 28600
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:44 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11730532276966059372
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:44 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11730532276966059372
---
```
开始运行命令...
[2025-09-14 15:39:01.386] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:01.386] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 29532
执行时间: 62.5ms
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
< Date: Sun, 14 Sep 2025 07:38:44 GMT
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
< Date: Sun, 14 Sep 2025 07:38:44 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=jk4435nucaigfilr0tn14dr1e5; expires=Sun, 14-Sep-2025 07:48:44 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=E1C3EF514318D0868F43B15F176E0B92.1757835524648; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:38:44 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:38:44 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=jk4435nucaigfilr0tn14dr1e5; expires=Sun, 14-Sep-2025 07:48:44 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=E1C3EF514318D0868F43B15F176E0B92.1757835524648; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:39:01.596] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:01.596] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 24944
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:44 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11016964773909684571
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:44 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11016964773909684571
---
```
开始运行命令...
[2025-09-14 15:39:07.927] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:39:07.927] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 8740
执行时间: 1.25s
---
```
开始运行命令...
[2025-09-14 15:39:08.493] [执行命令] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:39:08.498] [SERVER] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:39:11.622] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:11.622] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 20380
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:54 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11703807732045468870
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:54 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11703807732045468870
---
```
开始运行命令...
[2025-09-14 15:39:11.743] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:11.743] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 21876
执行时间: 15.625ms
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
< Date: Sun, 14 Sep 2025 07:38:54 GMT
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
< Date: Sun, 14 Sep 2025 07:38:55 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=rjigp9udahbiqe2bu2vghkuoi6; expires=Sun, 14-Sep-2025 07:48:55 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=2DB3E9452A95AD168143D147413F3B7F.1757835535029; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:38:54 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:38:55 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=rjigp9udahbiqe2bu2vghkuoi6; expires=Sun, 14-Sep-2025 07:48:55 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=2DB3E9452A95AD168143D147413F3B7F.1757835535029; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:39:11.955] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:11.955] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 16920
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:38:55 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9646116919234277085
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:38:55 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9646116919234277085
---
```
开始运行命令...
[2025-09-14 15:39:12.177] [执行命令] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:39:12.177] [SYSTEM] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:39:29.030] [执行命令] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
开始运行命令...
[2025-09-14 15:39:29.030] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
```
执行结果: 成功
进程PID: 32128
执行时间: 1.390625s
---
```
开始运行命令...
[2025-09-14 15:39:29.892] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:39:29.892] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 18932
执行时间: 1.296875s
---
```
开始运行命令...
[2025-09-14 15:39:30.498] [执行命令] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
开始运行命令...
[2025-09-14 15:39:30.498] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
```
执行结果: 成功
进程PID: 33088
---
```
开始运行命令...
[2025-09-14 15:39:31.503] [执行命令] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
开始运行命令...
[2025-09-14 15:39:31.503] [HTTP] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
```
执行结果: 成功
进程PID: 15472
---
```
开始运行命令...
[2025-09-14 15:39:34.638] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:34.638] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 27152
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:80 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:39:17 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11114213725559066475
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:39:17 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11114213725559066475
---
```
开始运行命令...
[2025-09-14 15:39:34.738] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:34.738] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 33348
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.49.44, 180.101.51.73
* SOCKS5 connect to 180.101.49.44:443 (locally resolved)
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
< Date: Sun, 14 Sep 2025 07:39:17 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12053972361341042284
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:39:17 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12053972361341042284
---
```
开始运行命令...
[2025-09-14 15:39:39.259] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:39:39.259] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 7724
执行时间: 1.203125s
---
```
开始运行命令...
[2025-09-14 15:39:39.833] [执行命令] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:39:39.838] [SERVER] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:39:42.986] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:42.987] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 16992
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:39:26 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12099489011486908215
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:39:26 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12099489011486908215
---
```
开始运行命令...
[2025-09-14 15:39:43.121] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:43.122] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 28736
执行时间: 46.875ms
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
< Date: Sun, 14 Sep 2025 07:39:26 GMT
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
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
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
< Date: Sun, 14 Sep 2025 07:39:26 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=da6gg43sfsuphjo81e9a2vir54; expires=Sun, 14-Sep-2025 07:49:26 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=B7AAE668122BA8D78BBEB1D5A4335CD1.1757835566427; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:39:26 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:39:26 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=da6gg43sfsuphjo81e9a2vir54; expires=Sun, 14-Sep-2025 07:49:26 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=B7AAE668122BA8D78BBEB1D5A4335CD1.1757835566427; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:39:43.352] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:43.352] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 3616
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:39:26 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11578191710999550333
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:39:26 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11578191710999550333
---
```
开始运行命令...
[2025-09-14 15:39:49.646] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:39:49.646] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 26976
执行时间: 921.875ms
---
```
开始运行命令...
[2025-09-14 15:39:50.194] [执行命令] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:39:50.199] [SERVER] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:39:53.309] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:53.309] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 3656
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:39:36 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10604512394036338802
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:39:36 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10604512394036338802
---
```
开始运行命令...
[2025-09-14 15:39:53.396] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:53.396] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 4924
执行时间: 62.5ms
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
< Date: Sun, 14 Sep 2025 07:39:36 GMT
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
< Date: Sun, 14 Sep 2025 07:39:36 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=r5adc9btepii4r4ghj6rvjgpc1; expires=Sun, 14-Sep-2025 07:49:36 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=0851339E834AD57C7B91710E35EBBAE6.1757835576699; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:39:36 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:39:36 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=r5adc9btepii4r4ghj6rvjgpc1; expires=Sun, 14-Sep-2025 07:49:36 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=0851339E834AD57C7B91710E35EBBAE6.1757835576699; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:39:53.619] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:39:53.619] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 16652
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:39:37 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11690380719257571337
< 
  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:39:37 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11690380719257571337
---
```
开始运行命令...
[2025-09-14 15:39:54.866] [执行命令] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:39:54.866] [SYSTEM] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:40:49.974] [执行命令] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
开始运行命令...
[2025-09-14 15:40:49.974] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
```
执行结果: 成功
进程PID: 33260
执行时间: 1.953125s
---
```
开始运行命令...
[2025-09-14 15:40:50.835] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:40:50.835] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 13204
执行时间: 1.0625s
---
```
开始运行命令...
[2025-09-14 15:40:51.384] [执行命令] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
开始运行命令...
[2025-09-14 15:40:51.384] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
```
执行结果: 成功
进程PID: 33004
---
```
开始运行命令...
[2025-09-14 15:40:52.388] [执行命令] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
开始运行命令...
[2025-09-14 15:40:52.388] [HTTP] ./main.exe -listen-port 10800 -upstream-type websocket -upstream-address ws://localhost:38800
```
执行结果: 成功
进程PID: 33592
---
```
开始运行命令...
[2025-09-14 15:40:55.507] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:40:55.507] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 20140
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:40:38 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9089392805058455881
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:40:38 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9089392805058455881
---
```
开始运行命令...
[2025-09-14 15:40:55.613] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:40:55.613] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 32484
执行时间: 78.125ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:40:38 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11814693453895016186
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:40:38 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11814693453895016186
---
```
开始运行命令...
[2025-09-14 15:41:00.018] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:41:00.018] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 20924
执行时间: 890.625ms
---
```
开始运行命令...
[2025-09-14 15:41:00.595] [执行命令] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:41:00.600] [SERVER] ./main.exe -listen-port 10800 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
开始运行命令...
[2025-09-14 15:41:03.772] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:41:03.772] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 32052
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:40:47 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11850497976356770840
< 
  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:40:47 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11850497976356770840
---
```
开始运行命令...
[2025-09-14 15:41:04.894] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:41:04.894] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 33092
执行时间: 31.25ms
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
< Date: Sun, 14 Sep 2025 07:40:48 GMT
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
< Date: Sun, 14 Sep 2025 07:40:48 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=pns6bpjq2fhip50a82qnlgbce1; expires=Sun, 14-Sep-2025 07:50:48 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=75A4D036FDC2007FD33111BC61387860.1757835648174; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:40:48 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:40:48 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=pns6bpjq2fhip50a82qnlgbce1; expires=Sun, 14-Sep-2025 07:50:48 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=75A4D036FDC2007FD33111BC61387860.1757835648174; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:41:05.114] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:41:05.114] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 26152
执行时间: 62.5ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:40:48 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11700703724719699381
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:40:48 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11700703724719699381
---
```
开始运行命令...
[2025-09-14 15:41:11.536] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
开始运行命令...
[2025-09-14 15:41:11.536] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```
执行结果: 成功
进程PID: 22704
执行时间: 1.1875s
---
```
开始运行命令...
[2025-09-14 15:41:12.088] [执行命令] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:41:12.093] [SERVER] ./main.exe -listen-port 10800
开始运行命令...
[2025-09-14 15:41:15.184] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:41:15.184] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 17880
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:80 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Sun, 14 Sep 2025 07:40:59 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10099163501024157018
< 
  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:40:59 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10099163501024157018
---
```
开始运行命令...
[2025-09-14 15:41:16.364] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:41:16.364] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 1920
执行时间: 46.875ms
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
< Date: Sun, 14 Sep 2025 07:40:59 GMT
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
< Date: Sun, 14 Sep 2025 07:40:59 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=ninbt3mvlhbbrtdgh4b1nngid3; expires=Sun, 14-Sep-2025 07:50:59 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=D104688DEE8EBCF31A42C117313B9CB1.1757835659649; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sun, 14 Sep 2025 07:40:59 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:40:59 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=ninbt3mvlhbbrtdgh4b1nngid3; expires=Sun, 14-Sep-2025 07:50:59 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=D104688DEE8EBCF31A42C117313B9CB1.1757835659649; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:41:16.572] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
开始运行命令...
[2025-09-14 15:41:16.572] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10800
```
执行结果: 成功
进程PID: 21380
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10800 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10800...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sun, 14 Sep 2025 07:40:59 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10582056024353338807
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:40:59 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10582056024353338807
---
```
开始运行命令...
[2025-09-14 15:41:16.745] [执行命令] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
开始运行命令...
[2025-09-14 15:41:16.745] [SYSTEM] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
