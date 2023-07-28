（1）先生成一个公钥、密钥对

#生成一个公钥密钥对
openssl genpkey -algorithm rsa -out rsa_private.key 2048
（2）基于该私钥我们生成一个CSR（证书签名请求）
openssl req -new -key rsa_private.key -out rsa_private.csr
（3）将该CSR文件发给CA，“注册一下”，当然了这个过程是收费的，要钱的。
这里，我们把自己当作一个CA，自己给自己注册一下，当然了，产生的证书是没人认可的。
#生成一个证书：mycert.crt 证书的有效期 365天

`openssl x509 -req -days 365 -in rsa_private.csr -signkey rsa_private.key -out mycert.crt`


注意：centos版本如果是CentOS Linux release 8.0.1905 (Core)版本，私钥长度不能设置成1024位，必须2048位。不然再最后启动nginx时会出如下错误。
原因是使用的私钥长度太短了，需要高于 1024 位，这里我们再重新生成一个 2048 位的密钥和证书：

openssl genrsa -des3 -out server.key 2048
openssl req -new -key server.key -out server.csr

Remove Passphrase from key : server.key.org 、server.crt
cp server.key server.key.org
openssl rsa -in server.key.org -out server.key


openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

