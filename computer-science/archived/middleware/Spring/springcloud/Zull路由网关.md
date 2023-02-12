### 什么是Zull

比方说，已经有三个Order服务，订单服务（8001）订单服务（8002）订单服务（8003）

如果我们直接访问对应的服务，就是<http://localhost:8001/provider/order/get/1>

但是如果我们配置了Zull,比如zuul-gateway（6001）那么我们可以通过Zull访问

<http://localhost:6001/microservice-order/provider/order/get/1> 

microservice-order是微服务名称。

那么我们自然会想到，如果我把三个订单服务都启动起来，因为它们的服务名称都是 microservice-order，zuul 到底会将请求转发给哪个服务呢？zuul 中默认集成了轮询的规则，三个服务轮流调用。



```yaml
# 配置路由规则
zuul:
  ignored-services: microservice-order # 不允许用微服务名访问了，如果禁用所有的，可以使用 "*"
  routes:
    prefix: /zuul # 给路由加一个统一的前缀
    # 如下指定新的映射
    order:
      serviceId: microservice-order
      path: /order/**
```

这样的话，就可以使用：<http://localhost:6001/zuul/order/provider/order/get/1> 来访问订单服务了。

