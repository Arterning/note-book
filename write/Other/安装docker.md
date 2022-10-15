curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
sudo gpasswd -a $USER docker
newgrp docker


因为需要搭建各种测试环境，mysql,redis,tomcat等等
我考虑了一下，仍然觉得还是使用docker最方便，一个镜像就解决了。


否则，你如果需要安装tomcat 你还得安装jdk 还要考虑jdk版本和tomcat版本的匹配问题。


vim /etc/docker/daemon.json
```json
{

"registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]

}
```


sudo service docker restart


docker pull tomcat:8
docker run --name tomcat -p 8080:8080 -v $PWD/test:/usr/local/tomcat/webapps/test -d tomcat


docker pull mysql:8.0.18
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql




install docker-compose

sudo curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose


The –L option tells the system to follow any redirects, in case the file has been moved
If you want a different version of Docker Compose, you may browse the list and substitute your preferred version for /1.24.0/
The –o option changes the filename, so it’s easier to type
The file will be saved in /usr/local/bin/

sudo chmod +x /usr/local/bin/docker-compose
You do not need to run an installation script for Docker Compose. Once downloaded, the software is ready to use.

Note: You can also install Docker Compose from the official Ubuntu repository. Simply run sudo apt-get install docker-compose. However, it is recommended to install the software package from Docker's official GitHub repository. That way, you are always installing the latest version.


docker–compose --version





How to Uninstall Docker Compose
To uninstall Docker Compose, simply delete the binary:

sudo rm /usr/local/bin/docker-compose
If you have installed Docker Compose using apt-get, use the following command to uninstall the package:

sudo apt-get remove docker-compose
Next, run a command to remove unnecessary software dependencies:

sudo apt-get autoremove




