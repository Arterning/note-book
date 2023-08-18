Spark 的安装部署
一笔带过。简单点来说就是下好安装包丢到服务器，解压下来去到conf文件夹 vim spark-env.sh，配一下 Java 的环境变量和 zookeeper，然后 vim slaves 去配置 worker 节点。然后修改 Spark 的环境变量并且分发到 worker 中，然后 source /etc/profile 即可。

