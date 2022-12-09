

### 导入依赖

```xml
<!-- feign的支持 -->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-openfeign</artifactId>
</dependency>
```

### 启动类注解配置

```java
@SpringBootApplication
@EnableEurekaClient
@EnableFeignClients
public class OrderConsumerFeign {

    public static void main(String[] args) {
        SpringApplication.run(OrderConsumerFeign.class, args);
    }
}

```



### Feign客户端

```java
//Feign 客户端要添加 @FeignClient 注解，value 属性表示作用到哪个微服务上
//方法内定义了两个接口，接口上即和普通的 SpringMVC 没什么区别
@FeignClient(value = "MICROSERVICE-ORDER")
public interface OrderClientService {

    @GetMapping("/provider/order/get/{id}")
    TOrder getOrder(@PathVariable(value = "id") Long id);

    @GetMapping("/provider/order/get/list")
    List<TOrder> getAll();
}

```



### Feign 客户端的使用

```java
@RestController
@RequestMapping("/consumer/order")
public class OrderConsumerController
{
    /**
     * 定义在microservice-common模块的feinClient
     */
    @Resource
    private OrderClientService orderClientService;

    @GetMapping("/get/{id}")
    public TOrder getOrder(@PathVariable Long id)
    {
        return orderClientService.getOrder(id);
    }

    @GetMapping("/get/list")
    public List<TOrder> getAll()
    {
        return orderClientService.getAll();
    }
}
```



对比下没有Feign的RestTemplate方式调用

1. 需要引入服务名称，耦合大，如果服务名称修改，那么都要修改
2. 需要引入请求API路径，比如"/provider/order/get/"这种东西。

```java

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





### 总结

可以看到,Feign的好处就是完全隐藏了RestTemplate的细节,调用方不必关心任何东西，直接注入OrderClientService，然后调用就完事了。

我们看一下官方的解释：Feign 是一个声明式 WebService 客户端。使用 Feign 能让编写的 WebService 客户端更加简洁，它的使用方法式定义一个接口，然后在上面添加注解。Spring Cloud 对 Feign 进行了封装，使其支持了 Spring MVC 标准注解和 HttpMessageConverters。Feign 可以与 Eureka 和 Ribbon 组合使用以支持负载均衡。

前面在使用 Ribbon + RestTemplate 时，利用 RestTemplate 对 http 请求的封装处理，形成了一套模板化的调用方法。但是在实际开发中，由于对服务依赖的调用可能不止一处，往往一个接口会被多处调用，所以通常都会针对每个微服务自行封装一些客户端类来包装这些依赖服务的调用。

所以，Feign 在此基础上做了进一步的封装，由它来帮助我们定义和实现依赖服务接口的定义。使用 Feign 只需要创建一个接口并使用一个注解来配置它即可。这就类似于我们在 dao 层的接口上标注 @Mapper 注解一样。这样的话，即完成了对服务提供方的接口绑定，简化了使用 Spring Cloud Ribbon 时的开发量。

