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

