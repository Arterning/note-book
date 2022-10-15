Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: 
Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": 
dial unix /var/run/docker.sock: connect: permission denied


sudo gpasswd -a ning docker   #将普通用户username加入到docker组
newgrp docker

# docker进程使用Unix Socket而不是TCP端口。而默认情况下，Unix socket属于root用户，需要root权限才能访问。
# 这个每次打开一个新的终端就得执行下这个。这个咋感觉是临时生效的，只在当前终端有用

sudo groupadd docker     #添加docker用户组
sudo gpasswd -a $USER docker     #将登陆用户加入到docker用户组中
usermod -aG docker $USER
newgrp docker     #更新用户组
docker ps    #测试docker命令是否可以使用sudo正常使用

# restart machine