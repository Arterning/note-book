## 域名配置中的record type[​](https://arterning.github.io/blogs/docs/notes/DNS%E7%9A%84%E8%AE%B0%E5%BD%95%E7%B1%BB%E5%9E%8B#%E5%9F%9F%E5%90%8D%E9%85%8D%E7%BD%AE%E4%B8%AD%E7%9A%84record-type "Direct link to 域名配置中的record type")

比如: A | AAAA | CNAME | MX | TXT/SPF | SRV | CAA，这些是什么玩意？

这些是域名系统（Domain Name System，DNS）中的不同记录类型。DNS记录类型定义了一个特定的记录应该包含哪些信息，以及如何使用该信息。

以下是常见的 DNS 记录类型及其作用：

A 记录：将域名映射到 IPv4 地址。

AAAA 记录：将域名映射到 IPv6 地址。

CNAME 记录：将一个域名映射到另一个域名。例如，可以将一个子域名的 CNAME 记录指向主域名，以便访问子域名时自动重定向到主域名。

MX 记录：指定邮件服务器，以便将电子邮件发送到该域名的收件人。

TXT/SPF 记录：用于验证邮件服务器发送电子邮件的授权。SPF（Sender Policy Framework）记录是 TXT 记录的一种特殊类型。

SRV 记录：指定提供特定服务的服务器。例如，可以使用 SRV 记录指定 VoIP 服务器的位置。

CAA 记录：指定哪些证书颁发机构（CA）可以签发该域名的 SSL/TLS 证书。

这些记录类型可以在 DNS 服务器的配置文件中设置。在使用任何这些记录类型之前，需要先确保您有权访问并配置 DNS 服务器或 DNS 记录。



netfly 部署静态网站，然后用域名指向它
![[Pasted image 20231019155517.png]]

注意使用CNAME类型


# 同时存在A记录和CNAME记录，哪个生效

当同时存在 A 记录和 CNAME 记录时，A 记录优先于 CNAME 记录。这是因为 A 记录直接将域名映射到一个 IP 地址，而 CNAME 记录则将域名映射到另一个域名上。

例如，如果存在以下记录：

```
example.com.     IN     A      192.0.2.1www.example.com. IN     CNAME  example.com.
```

则 A 记录将优先于 CNAME 记录。这意味着，当您访问 "[www.example.com"](http://www.example.com%22/) 时，DNS 服务器将首先查找 A 记录以确定其 IP 地址，然后将您重定向到该 IP 地址。即使 CNAME 记录指向相同的主机名，也不会使用该记录，因为 CNAME 记录被视为"别名"，不能直接映射到 IP 地址。

但是，当只有 CNAME 记录存在时，它将直接映射到另一个域名。例如：

```
www.example.com. IN     CNAME  example.com.
```

在这种情况下，当访问 "[www.example.com"](http://www.example.com%22/) 时，DNS 服务器将首先查找 CNAME 记录以确定其目标域名，然后将您重定向到该域名。如果目标域名也有 A 记录，则该记录将优先于 CNAME 记录。