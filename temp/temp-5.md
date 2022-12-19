
# git撤销修改
git checkout -- readme.txt
撤销所有修改
git checkout ..

# mysql 各个版本
1. MySQL Community Server 社区版本，开源免费，但不提供官方技术支持。
2. MySQL Enterprise Edition 企业版本，需付费，可以试用30天。
3. MySQL Cluster 集群版，开源免费。可将几个MySQL Server封装成一个Server。
4. MySQL Cluster CGE 高级集群版，需付费。
5. MySQL Workbench（GUI TOOL）一款专为MySQL设计的ER/数据库建模工具。它是著名的数据库设计工具DBDesigner4的继任者。MySQL Workbench又分为两个版本，分别是社区版（MySQL Workbench OSS）、商用版（MySQL Workbench SE）。


# docker基础命令
```shell
uname -r

yum install docker

#设置下国内源
sudo vim /etc/docker/daemon.json
#add
{
 "registry-mirrors": ["https://registry.docker-cn.com"]
}
sudo service docker restart


systemctl start docker

systemctl enable docker

systemctl stop docker

#检索
docker search tomcat 
docker search redis

#tag可选 tag就是镜像的版本
docker pull tomcat:tag

#查看本地镜像
docker images

#启动
-d：后台运行
-p: 将主机的端口映射到容器的一个端口，主机端口:容器内部的端口
--name:给容器起的名字
-v 宿主机目录:容器目录
docker run -p 8888:8080 --name mytomcat -d tomcat:latest
docker run -p 8888:8080 --name mytomcat -d tomcat:latest -v /root/docker/tomcat/webapps:/usr/local/tomcat/webapps




#查看运行中的容器
docker ps  
docker ps -a

#停止指定的容器
docker stop  容器的id

#删除
docker rm 容器id


```


# 其他
潜伏
人间正道是沧桑
长征
雍正王朝
楚汉

喜欢历史，读历史
我觉得一个人，尤其是一个青年人，满脑子想的都是升官发财，那是可悲的。
当然啦，对于现在的我来说，我还是要多赚钱的，不能降低对自己的要求，因为房子还没买呢。。媳妇还没娶呢。。
感觉，有点真香了。。
我有不奢望能够大富大贵，能保证我和我的家人平安富足的生活我就满足啦。

赚钱
利用信息不对称赚钱
技术人通过什么方式增加自己的价值
钻研某一个领域，可以记录博客，然后记录公众号。提高自己的影响力，可以出书。。

然后转型，做产品，做管理，或者其他的。
然后最终的目标，希望能够做一个自由职业者，写自己感兴趣的软件和代码。。

“打仗不是纸上谈兵，画一个箭头就可以达到目的地的


经常看到有些函数的参数用下划线开头，有什么特殊意义吗？
一般有两个作用：
1.不易被其他用户重新定义导致错误
2.程序内部调用的函数，一般用户不会使用

只是一种特殊的命名风格吧

helloGithub项目
介绍有趣的github项目