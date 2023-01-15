## Spring Cloud中使用Ribbon实现负载均衡详解

### 导入依赖

```xml
        <!--eureka Client-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
```

### 配置 application.yml

```yaml
server:
  port: 9001

eureka:
  client:
    register-with-eureka: true
    service-url:
      defaultZone: http://eureka01:7001/eureka/,http://eureka02:7002/eureka/,http://eureka03:7003/eureka/

  instance:
    instance-id: 消费方服务-9001  # 人性化显示出服务的信息
    prefer-ip-address: true    # 访问路径可显示ip地址
    lease-renewal-interval-in-seconds: 2
    lease-expiration-duration-in-seconds: 5

spring:
  cloud:
    loadbalancer:
      retry:
        enabled: true
  application:
    name: ribbon-client
```



### 向 http 中植入 Ribbon

```java
//修改下 RestTemplate 
@Configuration
public class RestTemplateConfig {
    @Bean
    @LoadBalanced //配置：@LoadBalanced'注解表示使用Ribbon实现客户端负载均衡
    public RestTemplate getRestTemplate() {
        return new RestTemplate();
    }

}

//主启动类上添加 @EnableEurekaClient，因为这个消费方也是一个 Eureka Client
@SpringBootApplication
@EnableEurekaClient
@RibbonClient(name = "MICROSERVICE-ORDER")
public class OrderConsumer {

    public static void main(String[] args) {
        SpringApplication.run(OrderConsumer.class, args);
    }
}

//将ip改成服务名称
@RestController
@RequestMapping("/consumer/order")
public class OrderConsumerController {

    private static final String ORDER_PROVIDER_URL_PREFIX = "http://MICROSERVICE-ORDER";

    @Resource
    private RestTemplate restTemplate;

    @GetMapping("/get/{id}")
    public TOrder getOrder(@PathVariable Long id) {
        return restTemplate.getForObject(ORDER_PROVIDER_URL_PREFIX + "/provider/order/get/" + id, TOrder.class);
    }

    @SuppressWarnings("unchecked")
    @GetMapping("/get/list")
    public List<TOrder> getAll() {
        return restTemplate.getForObject(ORDER_PROVIDER_URL_PREFIX + "/provider/order/get/list", List.class);
    }

    @GetMapping("/discovery")
    public Object discovery() {
        return restTemplate.getForObject(ORDER_PROVIDER_URL_PREFIX + "/provider/order/discovery", Object.class);
    }
}


```

