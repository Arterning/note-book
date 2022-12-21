PTH
pass the hash 利用LM hash或NTLM hash的值进行渗透测试

PTT 
pass the ticket 利用ticket票据进行。。

PTK 
pass the key 利用ekeys aes256..


LM和NTLM是windows密码的加密算法
#PTH在内网渗透种是一种很经典的攻击方式,原理就是攻击者可以直接通过LM Hash和NTLM Hash访问远程主机或服务,而不用提供明文密码。

如果禁用lentlm认证,PsExec无法利用获得的ntml hash进行远程连接,但是使用mimikatz还是可以攻击成功。对于8.1/2012r2,安装补丁kb2871997的win7/2008r/8/2012等,可以利用AES keys代替NT 哈市来实现ptk攻击



sekurlsa::pth



国产Ladon内网杀器测试