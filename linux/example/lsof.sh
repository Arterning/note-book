# 列出所有的网络连接
###
 # @Author: ning.huang ning.huang
 # @Date: 2022-11-15 15:30:58
 # @LastEditors: ning.huang ning.huang
 # @LastEditTime: 2022-11-15 15:31:10
 # @FilePath: \code\linux\example\lsof.sh
 # @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
### 
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
