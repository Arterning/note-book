`git clone https://github.com/mysql/mysql-server`

`sudo apt install cmake libssl-dev libzstd-dev libncurses5-dev libreadline-dev bison pkg-config`

`cd mysql-server`

`
mkdir build`

`
cd build`

`
cmake .. -DWITH_DEBUG=1 -DDOWNLOAD_BOOST=1 -DWITH_BOOST=~/boost_1_69_0`

`
make`

./build/bin/mysqld --initialize #would generate random password here
./build/bin/mysqld --debug

./build/bin/mysql -u root -p 
ALTER USER 'root'@'localhost' IDENTIFIED BY "root";
exit
./build/bin/mysqladmin -u root -p shutdown



refer
https://www.eversql.com/why-and-how-build-debug-mysql-from-source-code/
