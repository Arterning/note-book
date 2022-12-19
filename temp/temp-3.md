# drone教程
Drone CI for Github 的部署到此就真的结束了，一路走来，踩了不少坑，Drone 的文档你慢慢看了以后就会发现有多烂，烂到心累想哭，许多在 Demo 中出现的变量在参考手册中找不到，完全不知道什么含义，只能靠瞎猜和摸索。

有的东西文档烂 真是不好学啊 包括开源代码也是 如果文档烂
根本不好学

- https://learnku.com/articles/50293
- https://www.drone.io/


# 打造自己的jenkis ci
有的项目 本身技术特点比较负责 你要部署起来就比较麻烦
一般来说 
1. 你要下载依赖 maven jar包 node js包 经常出现下载缓慢 下载失败的情况
2. 你要编译二进制文件
3. 你要配置数据库。。
所以把这一整套流程用Dockerfile描述 用镜像承载这一切是非常好的办法
同时结合CI 当代码推送到github的时候自动触发构建 更新环境。

# vaultwarden项目学习
1. vaultwarden 前端Docker仓库 https://hub.docker.com/r/vaultwarden/web-vault/tags

2. 对应的移动端 使用的是C# Xamarin构建的 https://github.com/bitwarden/mobile

3. xamarin官网 https://dotnet.microsoft.com/zh-cn/apps/xamarin

4. 用xamarin开发App的体验如何？ https://www.zhihu.com/question/31159902


# web dav
WebDAV 是 web-based distributed authoring and versioning 的简称。它是 HTTP(S) 协议的一种延伸，可让 Web 服务器变成一般标准的网络驱动器。在启用 WebDAV 服务后，你就能透过 HTTP 或 HTTPS 协议联机至 NAS。

WebDAV协议为用户在服务器上创建、更改和移动文档提供了一个框架。WebDAV协议最重要的功能包括维护作者或修改日期的属性、名字空间管理、集合和覆盖保护。维护属性包括创建、删除和查询文件信息等。名字空间管理处理在服务器名称空间内复制和移动网页的能力。集合（Collections）处理各种资源的创建、删除和列举。覆盖保护处理与锁定文件相关的方面。

根据描述有点类似FTP 但是应该比FTP更先进吧 毕竟FTP协议这么旧了

WebDAV 基于 HTTP 协议的通信协议，在GET、POST、HEAD等几个HTTP标准方法以外添加了一些新的方法，使应用程序可对Web Server直接读写，并支持写文件锁定(Locking)及解锁(Unlock)，还可以支持文件的版本控制。


## 几种文件共享方式的比较
常用的文件共享有三种：FTP、Samba、WebDAV，它们各有优缺点，了解后才能更好地根据自己的需求选择方案。

- FTP属于古老的文件共享方式了，因为安全性，现代浏览器最新已默认不能打开FTP协议。SFTP在FTP基础上增加了加密，在Linux上安装OpenSSH后可以直接用SFTP协议传输。使用SFTP临时传送文件还可以，但做文件共享，性能不高，速度较慢。

- Samba是Linux下CIFS协议的实现，优势在于对于小白使用简章，和Windows系统文件共享访问一样，不需要安装第三方软件，而且移动端也有大量APP支持。苹果手机文件APP中添加网络存储用的就是这种方式。Windows下文件共享使用445端口，且不能更改。445端口常常受黑客关照，在广域网上大多运营封掉了访端口，所以这种文件共享只适合在内网使用。

- WebDAV 基于 HTTP 协议的通信协议，在GET、POST、HEAD等几个HTTP标准方法以外添加了一些新的方法，使应用程序可对Web Server直接读写，并支持写文件锁定(Locking)及解锁(Unlock)，还可以支持文件的版本控制。因为基于HTTP，在广域网上共享文件有天然的优势，移动端文件管理APP也大多支持WebDAV协议。使用HTTPS还能保安全性。Apache和Nginx支持WebDAV，可作为WebDAV文件共享服务器软件。也可以使用专门的WebDAV软件部署。

## 开源的webdev实现
- https://github.com/hacdias/webdav



# 密码管理器 bitwarden会跑路吗？
1. 免费的，盈利模式是啥？帮用户保管密码服务器不要钱吗？难道用爱发电？
2. （线上和线下的）安全做的怎么样？
3. 密码是比较隐私的东西，是否会收集和利用用户信息？比如做个密码词典。
4. 如何正确使用bitwarden？

回答

1. 如果你担心他们无法盈利你可以可以升级高级版，可以存文件。
2. 你的密码被你的主密码加密，只要你的主密码不泄露是安全的。反之，如果你忘记了主密码，你是无法通过邮箱或者任何东西找回密码的。
3. 和上一个问题一样，如果不知道你的主密码，没有人可以查看你的密码。
4. 推荐1Password或者KeePassP.S. 这个东西是开源的，即使他们跑路，你也可以用源码自建服务器。

作者：面铺棉门帘
链接：https://www.zhihu.com/question/419975229/answer/1459780865
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

# other
1. bitwarden官网 https://bitwarden.com/

2. 为什么国内的uniapp一直没人讨论呢？ https://www.zhihu.com/question/330607558

3. 用了六年 Keepass 了，发现大家都在推荐 Bitwarden https://www.v2ex.com/t/894022

