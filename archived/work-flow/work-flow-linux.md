# linux 系统磁盘空间不足 如何清理？ 如何查看哪个目录占用空间大

清理 Linux 系统的磁盘空间不足的方法有很多，这里介绍一些常见的方法：

删除不必要的文件：可以通过命令行删除临时文件、备份文件和日志文件等。

清除缓存：使用命令 sudo sync; echo 3 | sudo tee /proc/sys/vm/drop_caches 可以清除系统的缓存。

卸载不使用的软件：如果有安装了不再使用的软件，可以考虑卸载它们以释放空间。

清理包管理器：如果系统使用了包管理器，可以考虑清理已经卸载的包。

查看磁盘使用情况：

可以使用 df 命令查看磁盘使用情况，例如：
df 是 disk free 的缩写

```
$ df -h
Filesystem Size Used Avail Use% Mounted on
/dev/sda1 20G 5.2G 14G 28% /

```

可以使用 du 命令查看目录占用空间，例如：
disk usage

```
$ du -sh \*
4.0K bin
4.0K boot
40K dev
```

...
使用 -h 选项可以显示以人类可读的方式查看磁盘使用情况。

更多的查看磁盘使用情况的方法请参考 Linux 文档。
