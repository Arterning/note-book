拿到源码之后，第一时间寻找 web.config

接下来就寻找 asmx 结尾的文件
asmx 是什么？asmx 是 WEB 服务文件

soap 最容易出现什么漏洞
注入 上传 各种信息泄露等等
最简单的办法，发现 soap 直接丢到 sqlmap 百分之八十出注入

