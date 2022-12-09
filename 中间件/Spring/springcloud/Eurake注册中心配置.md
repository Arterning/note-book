

1.添加依赖

```xml
        <!--eureka-server服务端-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-server</artifactId>
        </dependency>
```



2.使用@EnableEurekaServer

```java
@SpringBootApplication
@EnableEurekaServer
public class EurekaServer01 {

    public static void main(String[] args) {
        SpringApplication.run(EurekaServer01.class, args);
    }
}
```



3.配置Yml,SpringBoot启动自动配置，大圣轮回，启动！

```yml
server:
  port: 7001


eureka:
  instance:
    #eureka服务端的实例名称
    hostname: eureka01
  client:
    # false表示不向注册中心注册自己
    register-with-eureka: false
    # false表示自己端就是注册中心，我的职责就是维护服务实例，并不需要去检索服务
    fetch-registry: false
    service-url:
      #设置与Eureka Server交互的地址查询服务和注册服务都需要依赖这个地址（单机）。
#      defaultZone: http://${eureka.instance.hostname}:${server.port}/eureka/
      defaultZone: http://eureka02:7002/eureka/,http://eureka03:7003/eureka/

  server:
    eviction-interval-timer-in-ms: 3000
    enable-self-preservation: false
```



当然咯，要在hosts文件中配置 127.0.0.1 和 eureka01 的映射。

127.0.0.1  eureka01
127.0.0.1  eureka02
127.0.0.1  eureka03







