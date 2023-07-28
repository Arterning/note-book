
## 关于nginx安装openssl自签名证书
申请证书比较麻烦 还需要拥有域名 所以先暂时创建一个自签名证书 
由于花费了不少时间 搞得挺头疼 在此记录一下过程

如果要提供一个有效的证书，服务器的证书必须从VeriSign这样的证书颁发机构签名，这样，浏览器就可以验证通过，否则，浏览器给出一个证书无效的警告。

申请一个证书签名的费用是一年几十~几百刀不等，所以如果只是出于管理目的，可以创建自签名证书，保证管理员通过浏览器安全连接到服务器。

下面简单介绍如何创建一个自签名的SSL证书。

创建自签名证书需要安装openssl，使用以下步骤：
1. 创建Key；
2. 创建签名请求；
3. 将Key的口令移除；为什么要移除口令？？
4. 用Key签名证书。



## 创建服务器证书密钥文件 server.key
openssl genrsa -des3 -out server.key 2048

这个时候会提示输入密码 这个密码要记住
注意des加密是对称加密 这里为什么要指定des加密？

## 创建服务器证书的申请文件 server.csr
openssl req -new -key server.key -out server.csr


...

CSR是Certificate Signing Request的英文缩写，即证书请求文件

CSR是以-----BEGIN CERTIFICATE REQUEST-----开头，-----END CERTIFICATE REQUEST-----为结尾的base64格式的编码。将其保存为文本文件，就是所谓的CSR文件。
匹配的KEY必须保存
有CSR必定有KEY，是成对的，CSR最终变成为证书，和私钥key配对使用。Key是以-----BEGIN RSA PRIVATE KEY-----开头的，-----END RSA PRIVATE KEY-----结尾的。Key必须保存好。

## 备份一份服务器密钥文件
cp server.key server.key.org

## 去除文件口令
openssl rsa -in server.key.org -out server.key

## 生成证书文件server.crt
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt


## 配置Nginx的证书 
ssl_certificate_key 配置server.key文件
ssl_certificate_crt 配置server.crt文件




```bash
#user  nobody;
worker_processes 1;
#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;
#pid        logs/nginx.pid;
events {
  worker_connections 1024;
}
http {
  include mime.types;
  default_type application/octet-stream;
  #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
  #                  '$status $body_bytes_sent "$http_referer" '
  #                  '"$http_user_agent" "$http_x_forwarded_for"';
  #access_log  logs/access.log  main;
  sendfile on;
  #tcp_nopush     on;
  #keepalive_timeout  0;
  keepalive_timeout 65;
  #gzip  on;
  server {
    listen 80;
    server_name localhost;
    rewrite ^(.*)$ https://$host$1; #将所有HTTP请求通过rewrite指令重定向到HTTPS
    #charset koi8-r;
    #access_log  logs/host.access.log  main;
    location / {
      root html;
      index index.html index.htm;
    }
    #error_page  404              /404.html;
    # redirect server error pages to the static page /50x.html
    #
    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
      root html;
    }
    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}
    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    #location ~ \.php$ {
    #    root           html;
    #    fastcgi_pass   127.0.0.1:9000;
    #    fastcgi_index  index.php;
    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
    #    include        fastcgi_params;
    #}
    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    #location ~ /\.ht {
    #    deny  all;
    #}
  }
  # another virtual host using mix of IP-, name-, and port-based configuration
  #
  #server {
  #    listen       8000;
  #    listen       somename:8080;
  #    server_name  somename  alias  another.alias;
  #    location / {
  #        root   html;
  #        index  index.html index.htm;
  #    }
  #}
  # HTTPS server
  #
  server {
      listen       443 ssl;
      server_name  localhost;
      ssl_certificate      server.crt;
      ssl_certificate_key  server.key;
    
ssl_session_timeout 5m;
ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
#表示使用的加密套件的类型。
ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3; #表示使用的TLS协议的类型。
ssl_prefer_server_ciphers on;
      location / {
          root   html;
          index  index.html index.htm;
      }
  }
}

```

## CSR在线生成工具
https://myssl.com/csr_create.html