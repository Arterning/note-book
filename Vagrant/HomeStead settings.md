
Install vbox

Homestead 脚本使用 Ruby 和 Shell 脚本编写而成。原理是对 Vagrantfile 文件做定制。将从 ~/Homestead/Homestead.yaml 读取的配置信息，在 provision 时，解析为 Vagrant 命令并进行对虚拟机的配置。Homestead 脚本的作用在于，提供了极其简单易用的接口，使我们只需要通过傻瓜化配置，即可完成复杂的任务。以下是几个常用的任务：

IP 配置，端口映射；
Nginx Site 创建；
数据库创建；
主机文件夹挂载到虚拟机等任务。

```bash
git clone https://git.coding.net/summerblue/homestead.git Homestead
```

```bash
bash init.sh
```

运行以上命令后，会在 ~/Homestead 目录下生成以下三个文件：

Homestead.yaml - 主要配置信息文件，我们可以在此文件中配置 Homestead 的站点和数据库等信息；
after.sh - 每一次 Homestead 盒子重置后（provision）会调用的 shell 脚本文件；
aliases - 每一次 Homestead 盒子重置后（provision），会被替换至虚拟机的 ~/.bash_aliases 文件中，aliases 里可以放一些快捷命令的定义。



Homestead.yaml 文件在 ~/Homestead 文件夹里，在 Git-Bash 里，你可以使用以下命令使用文件夹打开当前命令行所在目录：

Homestead.yaml 里的配置大致可以分为以下几种：

虚拟机设置；
SSH 秘钥登录配置；
共享文件夹配置；
站点配置；
数据库配置；
自定义变量；


```
cd ~/Homestead && vagrant up
```

If you vagrant up failure, try close the proxy settings in the terminal



Commands

```bash
vagrant init	初始化 vagrant
vagrant up	启动 vagrant
vagrant halt	关闭 vagrant
vagrant ssh	通过 SSH 登录 vagrant（需要先启动 vagrant）
vagrant provision	重新应用更改 vagrant 配置
vagrant reload --provision
vagrant destroy	删除 vagrant
```




If you can not ping the vbox host, try run this in the vbox 

```bash
sudo install net-tools
```



## Port mapping
We can see the default port mapping in the log

```
    homestead: 80 (guest) => 8000 (host) (adapter 1)
    homestead: 443 (guest) => 44300 (host) (adapter 1)
    homestead: 22 (guest) => 2222 (host) (adapter 1)
```

So we need to visit http://127.0.0.1:8000/ in the outside to visit the website

![[Pasted image 20240201171329.png]]
HomeStead的端口转发实际上基于这个virtualbox
可以发现，Docker和虚拟机都支持类似的功能 文件挂载 端口映射

## SSH key

SSH 秘钥登录配置
authorize 选项是指派登录虚拟机授权连接的公钥文件，此文件填写的是主机上的公钥文件地址，虚拟机初始化时，此文件里的内容会被复制存储到虚拟机的 /home/vagrant/.ssh/authorized_keys 文件中，从而实现 SSH 免密码登录。在这里我们默认填写即可。

`authorize: ~/.ssh/id_rsa.pub`

keys 是数组选项，填写的是本机的 SSH 私钥文件地址。虚拟机初始化时，会将此处填写的所有 SSH 私钥文件复制到虚拟机的 /home/vagrant/.ssh/ 文件夹中，从而使虚拟机能共享主机上的 SSH 私钥文件，使虚拟机具备等同于主机的身份认证。此功能为 SSH 授权提供了便利，例如在后面章节中，我们只需要在 GitHub 上配置一个 SSH 公钥，即可实现 GitHub 对虚拟机和主机共同认证。

此处我们将公钥和私钥一起同步到虚拟机中：

```bash
keys:
    - ~/.ssh/id_rsa
    - ~/.ssh/id_rsa.pub
```

