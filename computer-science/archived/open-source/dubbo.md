分布式系统
dubbo是rpc框架
rpc
remote process call

rpc的原理是socket连接
a将需要调用的方法名和方法参数传递给b
b完成处理之后，将结果返回给a


所有的服务都注册到注册中心
注册中心注册和维护了那个服务在那个位置。
a-->注册中心-->b


dubbo支持多种注册中心，推荐使用zk


下载，解压zk， 然后在相应的配置文件中配置好参数
1. datadir
2. port


除了注册中心之外 还有一个可选的监控中心。
dubbo-admin
