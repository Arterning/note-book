csrf get
login
username lucy
password 123456

csrf post
参考xss漏洞post的利用场景

csrf token
每次都返回一个token给客户端
客户端请求，需要携带token 服务端验证token是否正确。


感觉有点像验证码
不过目的不一样
token的目的是防止csrf 因为攻击者无法获取token 也就无法构造带有正确token的url
验证码的目的是区分是人在请求还是机器人在爆破