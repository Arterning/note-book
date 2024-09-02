OpenSSL 是一个开源项目，其组成主要包括一下三个组件：

openssl：多用途的命令行工具
libcrypto：加密算法库
libssl：加密模块应用库，实现了ssl及tls
openssl可以实现：秘钥证书管理、对称加密和非对称加密。

对称加密
对称加密需要使用的标准命令为 enc ，用法如下：
-in filename：指定要加密的文件存放路径

-out filename：指定加密后的文件存放路径

-salt：自动插入一个随机数作为文件内容加密，默认选项

-e：可以指明一种加密算法，若不指的话将使用默认加密算法

-d：解密，解密时也可以指定算法，若不指定则使用默认算法，但一定要与加密时的算法一致

-a/-base64：使用-base64位编码格式

```shell
openssl enc -ciphername [-in filename] [-out filename] [-pass arg] [-e] [-d] [-a/-base64]
       [-A] [-k password] [-kfile filename] [-K key] [-iv IV] [-S salt] [-salt] [-nosalt] [-z] [-md]
       [-p] [-P] [-bufsize number] [-nopad] [-debug] [-none] [-engine id]
```
