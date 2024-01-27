[Vagrant](http://vagrantup.com/) 是一款用来构建虚拟开发环境的工具，非常适合 php/python/ruby/java 这类语言开发 web 应用，“代码在我机子上运行没有问题”这种说辞将成为历史。

我们可以通过 Vagrant 封装一个 Linux 的开发环境，分发给团队成员。成员可以在自己喜欢的桌面系统（Mac/Windows/Linux）上开发程序，代码却能统一在封装好的环境里运行，非常霸气。


## **初始化开发环境**

这个就是我要的64位 ubuntu14.04 系统。先新建一个新的目录，然后进入目录，到终端里执行初始化命令：

```console
vagrant init ubuntu/trusty64
```

接下来执行开机命令，就会进入安装：

```console
vagrant up
```

成功之后可以输入一下命令可以连接进虚拟机：

```bash
vagrant ssh
```


Vagrant 初始化成功后，会在初始化的目录里生成一个 Vagrantfile 的配置文件，可以修改配置文件进行个性化的定制。

添加下面几行到 Vagrantfile 文件。用于调整内存大小

```text
config.vm.provider "virtualbox" do |v|
  v.memory = 2048
end
```

然后执行：

```text
vagrant reload
```

来加载设置就可以了。




## Add images to vagrant

假设我们下载的镜像存放路径是 `~/box/precise64.box`，在终端里输入：

`$ vagrant box add hahaha ~/box/precise64.box`

`hahaha` 是我们给这个 box 命的名字，`~/box/precise64.box` 是 box 所在路径。

可以使用下面的命令查看所有的镜像

`$ vagrant box list`


这时候，必须介绍一下 Boxes 这个概念了。Boxes 是一个后缀为 `box` 的文件，实际上它就是一个包含了虚拟机配置、虚拟机硬盘镜像和 Vagrant 配置的压缩包。之前在虚拟机里安装 Linux，首先需要下载 iso 安装镜像并新建虚拟机，然后修改虚拟机的 cpu、内存、光驱、网络等等硬件配置，再从光驱启动并安装虚拟机，完成后还有装虚拟机增强工具、配置系统、映射端口、映射磁盘目录等好多事要做。有了封装好的 Boxes，我们就可以快速运行所需要的操作系统了。不同 Provider 之间，Boxes 不能混用。以 VirtualBox 为例，下面这样就可以运行一个 `Ubuntu Server 14.04` 虚拟机了：

```
$ vagrant init ubuntu/trusty64
$ vagrant up
```

一个 Boxes 通常有几百 MB，执行上述命令后可以留意一下屏幕输出，查看 Boxes 的下载速度。如果速度很慢，可以复制出具体的 `.box` 文件 URL 并终止这次命令，改用其他工具下载。有了下载好的 Boxes 文件之后，打开 `VagrantFile`，修改 `config.vm.box` 配置为本地文件：

```bash
# config.vm.box = "ubuntu/trusty64"
# config.vm.box = "http://example.com/virtualbox.box"
config.vm.box = "~/Downloads/virtualbox.box"
```


虚拟机环境配好后，我们希望尽可能在宿主机操作，web 环境代码、配置最好都放在外面。Vagrant 已经考虑到这一点，主机的 Vagrant 工作目录（例如前面的 vagrant_project）会被自动映射为虚拟机的 `/vagrant` 目录。所以我们只要把虚拟机里的 Nginx 配置文件和 vhost 主目录都放在 `/vagrant` 里即可。



## Install something 

Write following code in Vagrantfile

```
  config.vm.provision "shell", inline: <<-SHELL
    apt-get update

    # Install nvm
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash

    # Activate nvm
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm

    # Install Node.js 18 using nvm
    nvm install 18
    nvm use 18

    # Install pnpm
    npm install -g pnpm
  SHELL
end
```




## port mapping
Vagrant 默认是使用端口映射方式将虚拟机的端口映射本地从而实现类似 `http://localhost:80` 这种访问方式，这种方式比较麻烦，新开和修改端口的时候都得编辑相比较而言，host-only 模式显得方便多了。打开`Vagrantfile`，将下面这行的注释去掉（移除 `#`）并保存：

```
config.vm.network :private_network, ip: "192.168.33.10"
```

重启虚拟机，这样我们就能用 `192.168.33.10` 访问这台机器了，你可以把 IP 改成其他地址，只要不产生冲突就行。



### 打包分发

当你配置好开发环境后，退出并关闭虚拟机。在终端里对开发环境进行打包(vagrant package –output NAME –vagrantfile FILE)：

$ vagrant package

打包完成后会在当前目录生成一个 `package.box` 的文件，将这个文件传给其他用户，其他用户只要添加这个 box 并用其初始化自己的开发目录就能得到一个一模一样的开发环境了。



## 常用命令

```bash
vagrant init  # 初始化
vagrant up  # 启动虚拟机
vagrant halt  # 挂起虚拟机
vagrant reload  # 重启虚拟机
vagrant ssh  # SSH 登录至虚拟机
vagrant status  # 查看虚拟机运行状态
vagrant destroy  # 销毁当前虚拟机
```