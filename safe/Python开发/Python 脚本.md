1. 查询Whois
Python 的 whois 模块的 whois 函数能够获取目标的 whois 信息；需要导入 python-whois 模块 

2. 判断 CDN
原理 Python 通过 os 库的 system 函数调用执行系统的 nslookup 命令来解析目标地址，如果解析 目标地址的 IP 地址过多，那么说明使用了 CDN 服务器。 

3. 子域名查询 
实现功能 查询目标的子域名 
原理 如输入 www.baidu.com；先正则去掉 www;然后加载字典，如内容为 aa;与处理后的 url 进行 拼接，即 aa.baidu.com；然后调用 socket 模块的 gethostbyname 函数来判断该域名是否能够 解析 IP，如果能说明该域名存在，不能则说明不存在。 

4. 端口扫描 


#glassfish 任意文件读取漏洞
从fofa中搜索到的可能存在漏洞的目标
使用编程批量操作 一次性返回所有有漏洞的目标

#python ftp/mysql/oracle密码爆破 其实也可以用现成的工具 这里只是介绍一下

#利用php代码混淆生成免杀木马

#SqlmapAPI 调用实现自动化 SQL 注入安全检测

#pocsuite3是一个漏洞检测框架 我们可以基于他做二次开发
