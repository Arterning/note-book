![](../../images/搜狗截图20190301215129.png)

```xml
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-config-server</artifactId>
        </dependency>
```



```java
@SpringBootApplication
@EnableConfigServer
public class MicroserviceConfig {
    public static void main(String[] args) {
        SpringApplication.run(MicroserviceConfig.class, args);
    }
}

```



```java
server:
  port: 5555

spring:
  application:
    name: microservice-config
  cloud:
    config:
      server:
        git:
          uri: https://github.com/eson15/microservice-config.git
```

