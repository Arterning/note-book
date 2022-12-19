# 解决SSH连接Vultr VPS主机经常出现断开连接失败
UseDNS修改配置
修改服务器配置：

/etc/ssh/sshd_config

将UseDNS yes 改为 no，保存。

最后，我们重启SSH然后检查是否稳定一些。

当然最后这些方法都没有什么用 我是用mosh来解决这个问题的