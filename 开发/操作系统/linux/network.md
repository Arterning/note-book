lsof -i

列出所有tcp 网络连接信息
lsof -i tcp

列出所有udp网络连接信息
lsof -i udp

列出谁在使用某个端口
lsof -i :8080

列出谁在使用某个特定的udp端口
lsof -i udp:5500
或者：特定的tcp端口命令：
lsof -i tcp:8081