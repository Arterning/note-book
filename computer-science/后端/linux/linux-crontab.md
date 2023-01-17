# crontab

# 设置crontab
若某个前面是 /x (x为数字)的形式,就是每x时间执行一次
1,3 1-3 分别指 1和3 与 1到3
对应 
分钟 小时 一个月中的第几日 月份 一个星期的第几天
```
* * * * * [command] >> [path/xx.log] > 2&1
```
比如每五分钟执行一次
```shell
crontab -e
*/5 * * * * /root/git-repo/schedule.sh
crontab -l
```

# 手动开启crontab日志
在root的状态下，编辑/etc/rsyslog.conf 把#cron.* /var/log/cron.log的注释#删去

```shell
# 查看crontab的执行情况
tail -f /var/log/cron.log
# 重启日志服务
/etc/init.d/rsyslog restart
# 查看日志信息
tail -f /var/log/cron.log
```

我们看一下这个日志配置文件的主要内容
```shell
auth,authpriv.*			/var/log/auth.log
*.*;auth,authpriv.none		-/var/log/syslog
cron.*				/var/log/cron.log
daemon.*			-/var/log/daemon.log
kern.*				-/var/log/kern.log
lpr.*				-/var/log/lpr.log
mail.*				-/var/log/mail.log
user.*				-/var/log/user.log
```