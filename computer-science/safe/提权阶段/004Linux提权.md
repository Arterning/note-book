#自动信息收集

主要使用三个用于枚举机器的脚本,它们在脚本之间有些区别,但是它们输出的内容很多相同。
LinEnum
Linuxprivchecker(是一个python文件,前期要收集服务器是否能够运行python文件)
chmod +x LinEnum


两个漏洞探针:linux-exploit-suggester ,linux-exploit-suggester2(是一个pl脚本)

Linux提权SUID配合脚本演示-Aliyun(基于web权限)
漏洞成因:chmod u+s给予了suid u-s删除了suid
执行过chmod u+s后,如果用户调用了这个程序会以root权限运行

https://pentestlab.blog/2017/09/25/suid-executables/