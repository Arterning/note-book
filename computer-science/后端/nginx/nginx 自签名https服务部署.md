---  
authors: endi  
tags: [nginx]  
title: nginx和https  
---  
  
# 生成自签名证书  
  
使用 OpenSSL 工具生成自签名证书。按照以下步骤操作：  
  
1.  安装 OpenSSL 工具：如果您使用的是 Linux 操作系统，可以使用您的包管理器安装 OpenSSL 工具。如果您使用的是 Windows 操作系统，则可以从 OpenSSL 网站下载预编译的二进制文件并安装。  
    2.  生成私钥：使用以下命令生成私钥文件：  
  
```bash  
openssl genrsa -out mykey.pem 2048```  
  
这将生成一个 2048 位的 RSA 私钥，并将其保存在名为“mykey.pem”的文件中。  
  
3.  生成证书签名请求：使用以下命令生成证书签名请求（CSR）：  
  
```bash  
openssl req -new -key mykey.pem -out mycsr.csr```  
  
在生成过程中，您需要输入一些证书详细信息，例如国家/地区名称、州/省份、城市、组织名称、组织单位名称等。  
  
4.  生成自签名证书：使用以下命令生成自签名证书：  
  
csharp  
  
```bash  
openssl x509 -req -days 365 -in mycsr.csr -signkey mykey.pem -out mycert.pem```  
  
  
# nginx 配置文件  
```bash  
worker_processes  1;  
events {    worker_connections  1024;}  
  
http {    include       mime.types;    default_type  application/octet-stream;  
    sendfile        on;    keepalive_timeout  65;  
    server {        listen 80;        server_name your_ip; #如果您没有域名，可以指定您服务器的 IP 地址。  
  
        return 301 https://$server_name$request_uri; # 80重定向到403  
    }  
    server {        listen 443 ssl;        server_name your_ip;  
        ssl_certificate /etc/nginx/ssl/mycert.pem; # 证书文件  
        ssl_certificate_key /etc/nginx/ssl/mykey.pem; # 私钥文件  
  
        location / {            root /usr/share/nginx/html;            index index.html;        }    }}  
  
```  
  
# 准备 Docker compose file  
  
```docker-compose  
services:  
  nginx:    image: nginx    ports:      - "80:80"      - "443:443"    volumes:      - ./nginx-data/nginx.conf:/etc/nginx/nginx.conf      - ./nginx-data/conf.d:/etc/nginx/conf.d      - ./nginx-data/ssl:/etc/nginx/ssl```  
挂载的文件必须要先准备好 也就是nginx.conf。  
挂载的目录可以不用创建，docker会自动帮我们创建。