
## install nvm
install nvm ..ok
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.2/install.sh | bash
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm


## install msf
install msf
curl https://raw.githubusercontent.com/rapid7/metasploit-omnibus/master/config/templates/metasploit-framework-wrappers/msfupdate.erb > msfinstall && \
  chmod 755 msfinstall && \
  ./msfinstall


## git提交后出现nano界面,解决方法
这个是使用nano进行编辑提交的页面，退出方法为：
Ctrl + X然后输入y再然后回车，就可以退出了
如果你想把默认编辑器换成别的：
方法一、在GIT配置中设置 core.editor: git config --global core.editor "vim"
方法二、编辑~/.gitconfig文件。在core中添加editor = vim。如此以后在使用git的时候就自动使用vim作为编辑器

smartIDE可以玩一下


## pgsql 配置
打开配置文件，postgresql.conf然后listen_addresses = '*'在CONNECTIONS AND AUTHENTICATION部分中添加。这指示服务器在所有网络接口上进行侦听。
sudo nano /etc/postgresql/11/main/postgresql.conf
修改如下内容：
listen_addresses = '*'     # what IP address(es) to listen on;
保存文件并重新启动PostgreSQL服务以使更改生效：


## pgsql 命令
psql命令有两种格式，分别是：
psql postgres://username:password@host:port/dbname  
psql -U username -h hostname -p port -d dbname
修改登录PostgreSQL密码
#ALTER USER postgres WITH PASSWORD 'xxxxxxxxxxx';
ALTER ROLE
注：
密码postgres要用引号引起来,命令最后有分号

登录PostgreSQL

#sudo -u postgres psql

\h：查看SQL命令的解释，比如\h select。

\?：查看psql命令列表。

\l：列出所有数据库。

\c [database_name]：连接其他数据库。

\d：列出当前数据库的所有表格。

\d [table_name]：列出某一张表格的结构。

\du：列出所有用户。

\e：打开文本编辑器。

\conninfo：列出当前数据库和连接的信息。


## Mysql连接报错：1130-host ... is not allowed to connect to this MySql server如何处理
　　这个问题是因为在数据库服务器中的mysql数据库中的user的表中没有权限(也可以说没有用户)，下面将记录我遇到问题的过程及解决的方法。

　　在搭建完LNMP环境后用Navicate连接出错

　　遇到这个问题首先到mysql所在的服务器上用连接进行处理

　　1、连接服务器: mysql -u root -p

　　2、看当前所有数据库：show databases;

　　3、进入mysql数据库：use mysql;

　　4、查看mysql数据库中所有的表：show tables;

　　5、查看user表中的数据：select Host, User,Password from user;

　　6、修改user表中的Host:update user set Host='%' where User='root';

　　7、最后刷新一下：flush privileges;

#一定要记得在写sql的时候要在语句完成后加上" ; "下面是图示说明

## PowerDesigner 是数据库建模软件
不好用


## Debian ufw防火墙工具
Debian 10安装了一个称为UFW的防火墙配置工具。它是用于管理iptables防火墙规则的用户友好型前端。它的主要目标是使防火墙的管理变得更容易，简单。

sudo ufw status verbose命令将会打印防火墙的状态。Status: inactive表示防火墙状态为非活动状态。Status: active表示UFW已激活。

sudo ufw status verbose
verbose表示打印详细信息
您可以运行命令sudo ufw app list列出服务器上所有可用的应用程序配置文件。sudo ufw app info命令可以查找指定配置文件包含的防火墙规则详细信息。

之前postgresql数据库一直无法连接 原来还是防火墙没有开放这端口
ufw allow 5432
ufw delete allow 5432

alter user postgres with password ''


## naviat破解
https://www.ykbkw.top/post/18.html



## ssh config配置文件
ssh config配置文件

Host 192.168.1.178
  HostName 192.168.1.178
  User kali

Host 192.168.56.10
  HostName 192.168.56.10
  User vagrant

Host vagrantos.com
  HostName vagrantos.com
  User vagrant

Host gitee.com
  Identityfile id_rsa.gitee
  HostName gitee.com
  User git

Host github.com
  Identityfile id_rsa.github
  HostName github.com
  User git


## mysql 开启远程访问权限
方式一：改表法
顾名思义,该方法就是直接修改更改"mysql"数据库里的"user"表里的"host"项，从"localhost"改为"%"
update user set host='%' where user='root';

方式二：授权法
通过GRANT命令可以授予主机远程访问权限
grant all privileges on db.* to 'user'
--赋予任何主机访问权限：
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'password' WITH GRANT OPTION;
--允许指定主机(IP地址)访问权限：
GRANT ALL PRIVILEGES ON *.* TO 'myuser'@'192.168.1.3' IDENTIFIED BY 'root' WITH GRANT OPTION;
通过GRANT命令赋权后,需要通过FLUSH PRIVILEGES刷新权限表使修改生效：
FLUSH PRIVILEGES;

再次查看MySQL远程访问权限配置
select  User,authentication_string,Host from user;

出于安全性考虑，尤其是正式环境下

1.不推荐直接给root开启远程访问权限。
本案例仅以root用户作为例子做开启远程访问权限的配置,此为演示环境!

2.建议做权限细分和限制
正式环境中，推荐通过创建Mysql用户并给对应的用户赋权的形式来开放远程服务权限，并指定IP地址，赋权时根据用户需求，在GRANT命令中只开放slect、update等权限，做到权限粒度最小化。


## 微软app无法正常打开的解决方法
安装windows terminal 和 wsl ubuntu linux
在某博客看到的方法，供参考，可以尝试一下，我的也是这么解决的
1、打开“运行”输入 inetcpl.cpl （“WINDOWS”+“R”键，输入 inetcpl.cpl亦可）
2、点开高级往下拉，勾上"使用TLS 1.2"选项，或者点还原高级设置。



## install mosh
https://ideawrights.com/mosh-windows-wsl/


mosh的安装和配置
https://www.cnblogs.com/journeyonmyway/p/8908483.html
install mosh client
sudo apt-get install python-software-properties
sudo add-apt-repository ppa:keithw/mosh
sudo apt-get update
sudo apt-get install mosh
https://www.linode.com/docs/guides/install-mosh-server-as-ssh-alternative-on-linux/





## 翻墙机场推荐
https://clashnode.xyz/
这个好想是clash开发者维护的网站？



## 计划
这倒是激励了我。
在明年三月份之前，找到一个远程工作，月薪突破20K
然后突破40K 最终达到60K
然后打造个人的独立开发产品 实现工资外的另外收入印钞机产品
fabric区块链框架 学习


## mosh-server启动服务后，需要打开端口
sudo ufw allow 60000:61000/udp
如果你只打开一个端口
就会出现Address used的问题



## nvm install on win10
nvm use latest 使用最新版本
nvm use lts 使用长期支持版本
nvm use version-number 使用你已安装的任何其他版本
nvm use报错exit status 1解决方法
nvm use 版本报错，出现exit status 1后面跟一堆乱码

此时是因为无权限，打开c盘C:\Windows\System32找到cmd以管理员身份打开，再次使用nvm use 版本报错命令，成功切换



## 奇怪的cpu 100%问题？？？？
postgres always lead to 100% cpu , which is weird...
why??
I used kill -9 pid to solve it temp, but i feel it will appear again....


## 远程工作，信任是开始远程的基础
https://github.com/greatghoul/remote-working
https://www.yuque.com/greatghoul/remote/teams#g3sCG
远程工作，信任是开始远程的基础，如果对员工不信任，远程监控没有意义。
我们团队已经全员远程6年了，员工的工作都是自觉自主的，我对团队成员也是充分信任，我们没有监控、没有日报、也没有每天的视频立会。我们通过每周视频例会沟通各项工作计划和进度，平常任何成员认为有沟通必要直接预约相关人员通过企业微信语音或视频沟通。作为团队带头人，我会不定期的与团队成员进行沟通。
我们从2008年开始至今一直专注于打造企业数字化平台一款产品，团队成员也相对很稳定，目前入职时间最短的伙伴已经来公司3年半了。我们努力打造一个小而美的远程团队，我们不融资、不盲目扩张、全员远程，没有地域限制、965。
我们这里没有996、没有办公室政治、也没有互联网公司的高额收入，虽然伙伴们生活在三线、四线甚至是五线城市，我们为每位员工在北京上了工资全额的社保和公积金，能够拿到高于所在城市同行从业者的收入，是我们远程团队稳定持续的基础。


## 相关公司
凡泰极客公司
https://www.finogeeks.com/
远程办公的公司-极狐GitLab


前端相关要求
3 年以上前端开发工作经验；
精通 HTML5 、ES6+、CSS3 等 Web 开发技术；
精通 React 技术栈，至少精通一种现代前端开发框架；
具体基础知识。您深刻掌握了某种技术的原理，例如 HTTP 协议、CSS/SCSS/LESS 样式的各类玩法、React/Vue/Angular2 或其他 MVC/MVVM 框架、ES5/ES6+的常用语法糖、Nodejs 的异步机制等等，并对这些前端生态圈的发展动态了如指掌。其他如面向对象设计、函数语言的基本原理等等基础技术“功底”，我们认为也非常重要；
加分项
那些你做过很酷的事情。例如熟悉 Electron 及 Nwjs 等跨平台客户端框架、看过 Chromium 源代码、开发过一些浏览器扩展（ Extension ）程序、写过一些关于 JavaScript 技术的博客文章… 具体不限，我们愿闻其详；
知识面。如果您尝试过 Web Assembly 、开发过 PWA 或者 windows 客户端、掌握 JS 以外的其他语言例如 Rust/Golang/Dart 等等、自学过 Flutter 之类框架，我们非常迫切希望跟您交流。




后端相关要求
熟练掌握 Go 语言，对代码基本规范有清晰的认识；
2 年以上 Go 语言项目开发经验，熟悉 Java/Python/Nodejs 优先；
具有 RESTFul 应用的相关设计开发经验，了解后端系统的核心思想和基础架构者优先；
熟悉 docker 工具,以及 kubernetes 服务编排工具；
熟悉 Postgres/Mongo/Redis 数据库的使用和性能调优等工作；
具有丰富的 linux 下高性能程序的调试工作，熟悉性能提升的技巧；
能阅读和理解优秀的开源系统代码，有参与开源代码或者开源框架贡献者优先；
对计算机网络系统有深刻理解，有分布式、点对点网络协议或应用开发经验者优先。


熟练掌握 Go 语言，对代码基本规范有清晰的认识；
2 年以上 Go 语言项目开发经验，熟悉 Java/Python/Nodejs 优先；
具有 RESTFul 应用的相关设计开发经验，了解后端系统的核心思想和基础架构者优先；
熟悉 docker 工具,以及 kubernetes 服务编排工具；
熟悉 Postgres/Mongo/Redis 数据库的使用和性能调优等工作；
具有丰富的 linux 下高性能程序的调试工作，熟悉性能提升的技巧；
能阅读和理解优秀的开源系统代码，有参与开源代码或者开源框架贡献者优先；
对计算机网络系统有深刻理解，有分布式、点对点网络协议或应用开发经验者优先。


功能研发：金融级 IM 服务后端开发，兼容开源去中心化实时通讯标准 Matrix ；
单元测试：为后端服务模块编写充分的单元测试用例，用工程化的方法保障产品质量；
性能优化：包括服务的高性能且并发安全、水平扩展、高可用、数据可靠性传递和存储，需要处理异步系统中的分布式事务问题；
稳定性保证：构建服务端的链路追踪系统，监控服务端的运行情况，确保服务端的稳定性；
开源社区贡献：以开源方式参与项目研发工作，贡献产品所依赖的开源社区。



## 关于区块链Transaction的解释 不用看视频了
https://learnblockchain.cn/article/619

我们为什么用GO语言来做区块链？
https://www.8btc.com/article/541146

主流区块链开发语言大比拼
https://learnblockchain.cn/article/3451


## fabric教程文档
https://hyperledger-fabric.readthedocs.io/en/latest/install.html
我真是服了，用中文搜索出来的搜索文档狗屁不通 而且文档以及过时了，对应的事老版本，坑的一逼
看来学习技术，最好还是用英文搜索，然后看英文文档，不然很容易浪费时间。我他妈无语。

Linux挖矿病毒的清除与分析


## Go path目录
$HOME/go
GOPATH目录一般为： $HOME/go --bin # 存放编译后的可执行文件 --pkg # 依赖包编译后的*.a文件 --src # 存放源码文件，以代码包为组织形式




## 回去做一下镜像
一个是hystreia镜像
一个是ngin防火墙镜像

