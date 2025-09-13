=== 命令执行日志文件 === 创建时间: 2025-09-14 01:41:42

## [2025-09-14 01:41:42.850] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go 执行结果: 成功 进程PID: 2500 执行时间: 1.078125s

[2025-09-14 01:41:45.980] [SERVER] ./main.exe -mode server -protocol websocket
-addr :8080 [2025-09-14 01:41:49.987] [SERVER] ./main.exe -mode server -protocol
socks5 -addr :10810 -upstream-type websocket -upstream-address
ws://localhost:8080 [2025-09-14 01:41:51.989] [TEST] C:\Program
Files\Git\mingw64\bin\curl.exe curl -v -I http://www.baidu.com -x
socks5://localhost:10810 执行结果: 成功 进程PID: 35136 执行时间: 31.25ms 输出: *
Host localhost:10810 was resolved.

- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:10810...
- Host www.baidu.com:80 was resolved.
- IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
- IPv4: 180.101.51.73, 180.101.49.44
- SOCKS5 connect to 180.101.51.73:80 (locally resolved)
- SOCKS5 request granted.
- Connected to localhost () port 10810
- using HTTP/1.x
- Connected to localhost (::1) port 10810
- using HTTP/1.x

> HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.14.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sat, 13 Sep 2025 17:41:36 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_12241266023760466573

<

0 277 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0HTTP/1.1 200 OK Accept-Ranges:
bytes Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive Content-Length: 277 Content-Type: text/html Date: Sat, 13
Sep 2025 17:41:36 GMT Etag: "575e1f60-115" Last-Modified: Mon, 13 Jun 2016
02:50:08 GMT Pragma: no-cache Server: bfe/1.0.8.18 Tr_id:
bfe_12241266023760466573

- Connection #0 to host localhost left intact

---
[2025-09-14 01:41:52.131] [TEST] C:\Program Files\Git\mingw64\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10810
执行结果: 成功
进程PID: 33680
执行时间: 62.5ms
输出: * Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
* IPv4: 180.101.51.73, 180.101.49.44
* SOCKS5 connect to 180.101.51.73:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 10810
* using HTTP/1.x
* schannel: disabled automatic use of client certificate
* ALPN: curl offers http/1.1
* ALPN: server accepted http/1.1
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.14.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Sat, 13 Sep 2025 17:41:36 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11531902842713116404
<
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:41:36 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11531902842713116404
* Connection #0 to host localhost left intact
---

## [2025-09-14 01:42:47.844] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go 执行结果: 成功 进程PID: 30256 执行时间: 1.359375s

[2025-09-14 01:42:48.849] [SERVER] ./main.exe -dohurl
https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl
https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 [2025-09-14
01:42:51.995] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v
-I http://www.baidu.com -x socks5://localhost:1080 执行结果: 成功 进程PID: 41748
执行时间: 0s 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:1080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:1080...
- Host www.baidu.com:80 was resolved.
- IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
- IPv4: 180.101.51.73, 180.101.49.44
- SOCKS5 connect to 180.101.51.73:80 (locally resolved)
- SOCKS5 request granted.
- Connected to localhost () port 1080
- using HTTP/1.x
- Connected to localhost (::1) port 1080
- using HTTP/1.x

> HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sat, 13 Sep 2025 17:42:35 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_12308250123216297774

<

0 277 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Accept-Ranges:
  bytes Cache-Control: private, no-cache, no-store, proxy-revalidate,
  no-transform Connection: keep-alive Content-Length: 277 Content-Type:
  text/html Date: Sat, 13 Sep 2025 17:42:35 GMT Etag: "575e1f60-115"
  Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT Pragma: no-cache Server:
  bfe/1.0.8.18 Tr_id: bfe_12308250123216297774

---
[2025-09-14 01:42:52.096] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:1080
执行结果: 成功
进程PID: 12176
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.so.com:80 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 302 Moved Temporarily
< Server: openresty
< Date: Sat, 13 Sep 2025 17:42:36 GMT
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
*   Trying [::1]:1080...
* Host www.so.com:443 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
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
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:42:36 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=vu4s59c6d754romadpuuva7507; expires=Sat, 13-Sep-2025 17:52:36 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=31A9FE582F86D99741C1316FB4045967.1757785356136; Max-Age=63072000; Domain=so.com; Path=/
<
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sat, 13 Sep 2025 17:42:36 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/
HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:42:36 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=vu4s59c6d754romadpuuva7507; expires=Sat, 13-Sep-2025 17:52:36 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=31A9FE582F86D99741C1316FB4045967.1757785356136; Max-Age=63072000; Domain=so.com; Path=/
---

[2025-09-14 01:42:52.283] [TEST]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x socks5://localhost:1080 执行结果: 成功 进程PID: 8264
执行时间: 46.875ms 输出: Note: Using embedded CA bundle, for proxies (233263
bytes)

- Host localhost:1080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:1080...
- Host www.baidu.com:443 was resolved.
- IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
- IPv4: 180.101.51.73, 180.101.49.44
- SOCKS5 connect to 180.101.51.73:443 (locally resolved)
- SOCKS5 request granted.
- Connected to localhost () port 1080
- using HTTP/1.x
- ALPN: curl offers h2,http/1.1
- TLSv1.3 (OUT), TLS handshake, Client hello (1): } [308 bytes data]
- CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
- CApath: none
- TLSv1.3 (IN), TLS handshake, Server hello (2): { [102 bytes data]
- TLSv1.2 (IN), TLS handshake, Certificate (11): { [4771 bytes data]
- TLSv1.2 (IN), TLS handshake, Server key exchange (12): { [333 bytes data]
- TLSv1.2 (IN), TLS handshake, Server finished (14): { [4 bytes data]
- TLSv1.2 (OUT), TLS handshake, Client key exchange (16): } [70 bytes data]
- TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1): } [1 bytes data]
- TLSv1.2 (OUT), TLS handshake, Finished (20): } [16 bytes data]
- TLSv1.2 (IN), TLS change cipher, Change cipher spec (1): { [1 bytes data]
- TLSv1.2 (IN), TLS handshake, Finished (20): { [16 bytes data]
- SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
- ALPN: server accepted http/1.1
- Server certificate:
- subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science
  Technology Co., Ltd; CN=baidu.com
- start date: Jul 9 07:01:02 2025 GMT
- expire date: Aug 10 07:01:01 2026 GMT
- subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
- issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
- SSL certificate verify ok.
- Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Connected to localhost (::1) port 1080
- using HTTP/1.x

> HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sat, 13 Sep 2025 17:42:36 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_11744985404428428057

<

0 277 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Accept-Ranges:
  bytes Cache-Control: private, no-cache, no-store, proxy-revalidate,
  no-transform Connection: keep-alive Content-Length: 277 Content-Type:
  text/html Date: Sat, 13 Sep 2025 17:42:36 GMT Etag: "575e1f60-115"
  Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT Pragma: no-cache Server:
  bfe/1.0.8.18 Tr_id: bfe_11744985404428428057

---
[2025-09-14 01:42:57.793] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 1852
执行时间: 1.34375s
---

[2025-09-14 01:42:58.960] [SERVER] ./main.exe [2025-09-14 01:43:01.053] [TEST]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x socks5://localhost:1080 执行结果: 成功 进程PID: 39564
执行时间: 46.875ms 输出: Note: Using embedded CA bundle, for proxies (233263
bytes)

- Host localhost:1080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:1080...
- Host www.baidu.com:80 was resolved.
- IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
- IPv4: 180.101.51.73, 180.101.49.44
- SOCKS5 connect to 180.101.51.73:80 (locally resolved)
- SOCKS5 request granted.
- Connected to localhost () port 1080
- using HTTP/1.x
- Connected to localhost (::1) port 1080
- using HTTP/1.x

> HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sat, 13 Sep 2025 17:42:45 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_9571653954901737043

<

0 277 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Accept-Ranges:
  bytes Cache-Control: private, no-cache, no-store, proxy-revalidate,
  no-transform Connection: keep-alive Content-Length: 277 Content-Type:
  text/html Date: Sat, 13 Sep 2025 17:42:45 GMT Etag: "575e1f60-115"
  Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT Pragma: no-cache Server:
  bfe/1.0.8.18 Tr_id: bfe_9571653954901737043

---
[2025-09-14 01:43:01.133] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x socks5://localhost:1080
执行结果: 成功
进程PID: 33172
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:1080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:1080...
* Host www.so.com:80 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:80 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
* using HTTP/1.x
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 302 Moved Temporarily
< Server: openresty
< Date: Sat, 13 Sep 2025 17:42:45 GMT
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
*   Trying [::1]:1080...
* Host www.so.com:443 was resolved.
* IPv6: (none)
* IPv4: 180.163.237.15
* SOCKS5 connect to 180.163.237.15:443 (locally resolved)
* SOCKS5 request granted.
* Connected to localhost () port 1080
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
* Connected to localhost (::1) port 1080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:42:45 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=ul07r9l8nvaad053htc6te90v0; expires=Sat, 13-Sep-2025 17:52:45 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=80B2AAAD11A01A2C6F3C7122E7FA81E4.1757785365163; Max-Age=63072000; Domain=so.com; Path=/
<
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Moved Temporarily
Server: openresty
Date: Sat, 13 Sep 2025 17:42:45 GMT
Content-Type: text/html
Connection: keep-alive
Location: https://www.so.com/
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/
HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:42:45 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=ul07r9l8nvaad053htc6te90v0; expires=Sat, 13-Sep-2025 17:52:45 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=80B2AAAD11A01A2C6F3C7122E7FA81E4.1757785365163; Max-Age=63072000; Domain=so.com; Path=/
---

[2025-09-14 01:43:01.306] [TEST]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x socks5://localhost:1080 执行结果: 成功 进程PID: 37820
执行时间: 31.25ms 输出: Note: Using embedded CA bundle, for proxies (233263
bytes)

- Host localhost:1080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:1080...
- Host www.baidu.com:443 was resolved.
- IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
- IPv4: 180.101.51.73, 180.101.49.44
- SOCKS5 connect to 180.101.51.73:443 (locally resolved)
- SOCKS5 request granted.
- Connected to localhost () port 1080
- using HTTP/1.x
- ALPN: curl offers h2,http/1.1
- TLSv1.3 (OUT), TLS handshake, Client hello (1): } [308 bytes data]
- CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
- CApath: none
- TLSv1.3 (IN), TLS handshake, Server hello (2): { [102 bytes data]
- TLSv1.2 (IN), TLS handshake, Certificate (11): { [4771 bytes data]
- TLSv1.2 (IN), TLS handshake, Server key exchange (12): { [333 bytes data]
- TLSv1.2 (IN), TLS handshake, Server finished (14): { [4 bytes data]
- TLSv1.2 (OUT), TLS handshake, Client key exchange (16): } [70 bytes data]
- TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1): } [1 bytes data]
- TLSv1.2 (OUT), TLS handshake, Finished (20): } [16 bytes data]
- TLSv1.2 (IN), TLS change cipher, Change cipher spec (1): { [1 bytes data]
- TLSv1.2 (IN), TLS handshake, Finished (20): { [16 bytes data]
- SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
- ALPN: server accepted http/1.1
- Server certificate:
- subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science
  Technology Co., Ltd; CN=baidu.com
- start date: Jul 9 07:01:02 2025 GMT
- expire date: Aug 10 07:01:01 2026 GMT
- subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
- issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
- SSL certificate verify ok.
- Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Connected to localhost (::1) port 1080
- using HTTP/1.x

> HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sat, 13 Sep 2025 17:42:45 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_11513690725326516743

<

0 277 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Accept-Ranges:
  bytes Cache-Control: private, no-cache, no-store, proxy-revalidate,
  no-transform Connection: keep-alive Content-Length: 277 Content-Type:
  text/html Date: Sat, 13 Sep 2025 17:42:45 GMT Etag: "575e1f60-115"
  Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT Pragma: no-cache Server:
  bfe/1.0.8.18 Tr_id: bfe_11513690725326516743

---
[2025-09-14 01:43:01.492] [SYSTEM] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
[2025-09-14 01:43:07.067] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 11024
执行时间: 1.53125s
---

[2025-09-14 01:43:10.176] [SERVER] ./main.exe -mode server -protocol websocket
-addr :8080 [2025-09-14 01:43:14.179] [SERVER] ./main.exe -mode server -protocol
socks5 -addr :10810 -upstream-type websocket -upstream-address
ws://localhost:8080 [2025-09-14 01:43:16.189] [TEST]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x socks5://localhost:10810 执行结果: 成功 进程PID: 31524
执行时间: 46.875ms 输出: Note: Using embedded CA bundle, for proxies (233263
bytes)

- Host localhost:10810 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:10810...
- Host www.baidu.com:80 was resolved.
- IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
- IPv4: 180.101.51.73, 180.101.49.44
- SOCKS5 connect to 180.101.51.73:80 (locally resolved)
- SOCKS5 request granted.
- Connected to localhost () port 10810
- using HTTP/1.x
- Connected to localhost (::1) port 10810
- using HTTP/1.x

> HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sat, 13 Sep 2025 17:43:00 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_11301888076704696750

<

0 277 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Accept-Ranges:
  bytes Cache-Control: private, no-cache, no-store, proxy-revalidate,
  no-transform Connection: keep-alive Content-Length: 277 Content-Type:
  text/html Date: Sat, 13 Sep 2025 17:43:00 GMT Etag: "575e1f60-115"
  Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT Pragma: no-cache Server:
  bfe/1.0.8.18 Tr_id: bfe_11301888076704696750

---
[2025-09-14 01:43:16.285] [TEST] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x socks5://localhost:10810
执行结果: 成功
进程PID: 36860
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Host www.baidu.com:443 was resolved.
* IPv6: 240e:e9:6002:1fd:0:ff:b0e1:fe69, 240e:e9:6002:1ac:0:ff:b07e:36c5
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
< Date: Sat, 13 Sep 2025 17:43:00 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11603825484550360596
<
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:43:00 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11603825484550360596
---
