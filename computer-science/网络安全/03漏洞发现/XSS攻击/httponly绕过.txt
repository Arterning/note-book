如果HTTP响应头中包含HttpOnly标志，只要浏览器支持HttpOnly标志，客户端脚本就无法访问cookie。因此，即使存在跨站点脚本（XSS）缺陷，且用户意外访问利用此漏洞的链接，浏览器也不会向第三方透露cookie。如果浏览器不支持HttpOnly并且网站尝试设置HttpOnly cookie，浏览器会忽略HttpOnly标志，从而创建一个传统的，脚本可访问的cookie。

目前sun公司还没有公布相关的http only API，但PHP、C#均有实现。搞javaEE的兄弟们比较郁闷了，
如何在Java中设置cookie是HttpOnly呢？
Servlet 2.5 API 不支持 cookie设置HttpOnly 
http://docs.oracle.com/cd/E17802_01/products/products/servlet/2.5/docs/servlet-2_5-mr2/
建议升级Tomcat7.0，它已经实现了Servlet3.0
http://tomcat.apache.org/tomcat-7.0-doc/servletapi/javax/servlet/http/Cookie.html
但是苦逼的是现实是，老板是不会让你升级的。



javaee
response.setHeader("Set-Cookie", "cookiename=value; Path=/;Domain=domainvalue;Max-Age=seconds;HTTPOnly");

PHP4
header("Set-Cookie: hidden=value; httpOnly");

PHP5
setcookie("abc", "test", NULL, NULL, NULL, NULL, TRUE)


将cookie设置成HttpOnly是为了防止XSS攻击，窃取cookie内容，这样就增加了cookie的安全性，即便是这样，也不要将重要信息存入cookie。
服务端可以通过在它创建的cookie上设置HttpOnly标志来缓解这个问题，指出不应在客户端上访问cookie。客户端脚本代码尝试读取包含HttpOnly标志的cookie，如果浏览器支持HttpOnly，则返回一个空字符串作为结果。这样能够阻止恶意代码（通常是XSS攻击）将cookie数据发到攻击者网站。

支持HttpOnly的主流浏览器有哪些呢？谷歌了一下，常见的浏览器都支持。

对方开启HttpOnly你在盗取cookie失败的情况下可以采用其他方案
登陆后台权限方式
1.以cookie形式
2.直接账号密码登录：
保存账号密码读取：通过读取他保存在本地的数据
（需要xss产生于登录地址，利用表单劫持）