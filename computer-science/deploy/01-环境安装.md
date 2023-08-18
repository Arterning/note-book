# php 面板 小皮系统

安装 mysql8 php nginx redis

# psotgres 数据库

这里使用 EnterpriseDB 来下载安装，EnterpriseDB 是全球唯一一家提供基于 PostgreSQL 企业级产品与服务的厂商。

下载地址：https://www.enterprisedb.com/downloads/postgres-postgresql-downloads。

官网下载可执行程序 直接执行

# 编程语言

## java 和 scala

使用 idea 下载

## go

```bash
# Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz


# Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

# Add /usr/local/go/bin to the PATH environment variable.
# You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

export PATH=$PATH:/usr/local/go/bin

# Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

# Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version

# Confirm that the command prints the installed version of Go.
```

## python

执行安装包.exe

## rust
