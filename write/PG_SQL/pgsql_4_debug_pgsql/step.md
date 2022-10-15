sudo apt-get install build-essential libreadline-dev zlib1g-dev flex bison libxml2-dev libxslt-dev libssl-dev libxml2-utils xsltproc
sudo apt-get install -y libsystemd-dev

./configure --enable-depend --enable-cassert --enable-debug --with-systemd
//允许调试，开启postgresql的service服务

make

//because we need to install to /usr/local
sudo make install

PATH="/usr/local/pgsql/bin:$PATH"

mkdir <datadir> // 目录的存放位置随意，最好在自己的用户目录下

initdb -D <datadir>

/usr/local/pgsql/bin/initdb -D <datadir>


用于调试postgresql8.4源码的镜像，搭配《PostgreSQL数据库内核分析》学习。 此镜像基于centos镜像构建，安装了gcc与gdb，编译和安装了postgresql的源代码。 
https://hub.docker.com/r/lishizhen/postgresql84_debug


download website
http://ftp.postgresql.org/pub/source/


install from source code
http://postgres.cn/docs/14/install-procedure.html

debug postgresql by vscode 
http://www.postgres.cn/v2/news/viewone/1/459