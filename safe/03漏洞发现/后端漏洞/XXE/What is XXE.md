
xml外部实体注入，全称为XML external entity injection，某些应用程序允许XML格式的数据输入和解析，可以通过引入外部实体的方式进行攻击。



## 挖掘思路

关注可能解析xml格式数据的功能处，较容易发现的是请求包参数包含XML格式数据，不容易发现的是文件上传及数据解析功能处，通过改请求方式、请求头Content-Type等方式进行挖掘，思路一般分三步：

1. 检测XML是否会被成功解析以及是否支持DTD引用外部实体，有回显或者报错；；
2. 需注意没有回显则可以使用Blind XXE漏洞来构建一条带外信道提取数据
3. 最后可以尝试XInclude，某些应用程序接收客户端提交的数据，将其嵌入到服务器端的XML文档中，然后解析文档，尝试payload：

<fooxmlns：xi=“http://www.w3.org/2001/XInclude”>

<xi：include parse =“text”href =“file：/// etc / passwd”/> </ foo>






## 总结

XMl 可以引入外部实体，如果攻击者 在这里 引入一些特殊文件，比如/etc/passwd 那么就可以获得机器密码






XXE -“xml external entity injection”，即"xml外部实体注入漏洞"。

XML中对数据的引用称为实体，实体中有一类叫外部实体，用来引入外部资源，有SYSTEM和PUBLIC两个关键字，表示实体来自本地计算机还是公共计算机，外部实体的引用可以借助各种伪协议，比如如下三种，具体的示例可以看下面的xxe靶场实验：
file:///path/file.ext
http://url
php://filter/read=convert.base64-encode/resource=conf.php


内部申明DTD格式
<!DOCTYPE 根元素 [元素申明]>

外部引用DTD格式
<!DOCTYPE 根元素 SYSTEM "外部DTD的URI">

引用公共DTD格式
<!DOCTYPE 根元素 PUBLIC "DTD标识名" "公共DTD的URI">


第一部分：XML声明部分
<?xml version="1.0"?>

第二部分：文档类型定义 DTD
<!DOCTYPE note[ 
<!--定义此文档是note类型的文档-->
<!ENTITY entity-name SYSTEM "URI/URL">
<!--外部实体声明-->
]>

第三部分：文档元素
<note>
<to>Dave</to>
<from>Tom</from>
<head>Reminder</head>
<body>You are a good man</body>
</note>


