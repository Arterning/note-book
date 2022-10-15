

1.添加依赖

```xml
        <!--eureka-client客户端-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <!-- spring boot actuator 监控信息 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

```



### 2.配置yml

```yml
#配置服务名称是microservice-order
spring:
  application:
    name: microservice-order # 对外暴露的服务名称
    
# 客户端注册进eureka服务列表里，因为有三个eureka，所以订单服务要注册到三个eureka 中。
eureka:
  client:
    service-url:
      defaultZone: http://eureka01:7001/eureka/,http://eureka02:7002/eureka/,http://eureka03:7003/eureka/,
    healthcheck:
      enabled: true
  instance:
    instance-id: 书籍订单服务-8001  # 人性化显示出服务的信息
    prefer-ip-address: true    # 访问路径可显示ip地址
    lease-renewal-interval-in-seconds: 2
    lease-expiration-duration-in-seconds: 5
    
# 使用actuator来展示项目的基本信息
info:
  author.name: shengwu ni
  app.name: microservice
  server.port: ${server.port}
  application.name: ${spring.application.name}    
```



### 3.使用@EnableEurekaClient

```java
@SpringBootApplication
@EnableEurekaClient
@MapperScan("com.itcodai.springcloud.dao")
public class OrderProvider01 {

    public static void main(String[] args) {
        SpringApplication.run(OrderProvider01.class, args);
    }
}
```



注意，如果Eureka Client配置好了，但是Eureka Server没有启动，那么Client启动会报错

```java
com.sun.jersey.api.client.ClientHandlerException: java.net.ConnectException: Connection refused: connect
```





### 服务端集群配置

- 对外暴露的服务名称必须要相同，因为都是同一个服务，只不过有多个而已，因为接下来Ribbon是通过服务名来调用服务的；
- 每个服务连接了不同的数据库，这样用来区分不同的服务，便于测试，实际中也可能是便于维护；
- 每个服务的个性化名称展示可以区分一下，这样在eureka里可以很好的辨别出来







