### 服务熔断

多个微服务之间调用的时候，假设微服务A调用微服务B和微服务C，微服务B和微服务C又在调用其他的微服务，这就是所谓的“扇出”。如果扇出的链路上某个微服务的调用响应时间过长或者不可用，那么对微服务A的调用就会占用越来越多的系统资源，进而引起系统崩溃，这就是所谓的“**雪崩效应**”。

服务熔断机制是应对雪崩效应的一种微服务链路保护机制。当扇出链路的某个微服务不可用或者响应时间太长，就会进行服务的降级，快速熔断该节点微服务的调用，返回默认的响应信息。当检测到该节点微服务调用响应正常后即可恢复。



### 服务降级

服务降级是在客户端完成的，不是服务端，与服务端是没有关系的。就像银行某个窗口挂了“暂停服务”，那客户会自然去别的窗口。



### 依赖导入

```xml
 <!-- hystrix -->
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-starter-netflix-hystrix</artifactId>
    </dependency>
```





### 添加 `@EnableCircuitBreaker` 注解

```java
@SpringBootApplication
@EnableEurekaClient
@MapperScan("com.itcodai.springcloud.dao")
@EnableCircuitBreaker
public class OrderProvider01Hystrix {

    public static void main(String[] args) {
        SpringApplication.run(OrderProvider01Hystrix.class, args);
    }
}
```



### 对Controller接口的改动

关键部分

```java
@HystrixCommand(fallbackMethod = "processGetOrderHystrix")
```



```java
@RestController
@RequestMapping("/provider/order")
public class OrderProviderController {

    @Resource
    private OrderService orderService;
    @Resource
    private EurekaClient client;

    private static final Logger LOGGER = LoggerFactory.getLogger(OrderProviderController.class);

    /**
     * HystrixCommond注解中的fallbackMethod指示的是：当该方法出异常时，调用processGetOrderHystrix方法
     * @param id id
     * @return 订单信息
     */
    @GetMapping("/get/{id}")
    @HystrixCommand(fallbackMethod = "processGetOrderHystrix")
    public TOrder getOrder(@PathVariable Long id) {
        TOrder order = orderService.findById(id);
        if (order == null) {
            throw new RuntimeException("数据库没有对应的信息");
        }
        return order;
    }

    /**
     * 上面getOrder()方法出异常后的熔断处理方法
     * @param id id
     * @return 订单信息
     */
    public TOrder processGetOrderHystrix(@PathVariable Long id) {
        return new TOrder().setId(id)
                .setName("未找到该ID的结果")
                .setPrice(0d)
                .setDbSource("No this datasource");
    }

}

```



### 思考

 @HystrixCommand 注解是加在 Controller 层的接口方法上的，这会导致两个问题：

第一：如果接口方法很多，那么我是不是要在每个方法上都得加上该注解，而且，针对每个方法，我都要指定一个处理函数，这样会导致 Controller 变得越来越臃肿。

第二：这也不符合设计规范，理论上来说，Controller 层就是 Controller 层，我只管写接口即可。就像上一节介绍的 Feign，也是面向接口的，做均衡处理，我自己定义一个接口专门用来做均衡处理，在 Controller 层将该接口注入即可。那么 hystrix 是否也可以有类似的处理呢？

答案是肯定的，这跟面向切面编程一个道理，Cotroller 你只管处理接口逻辑，当出了问题，OK，交给我 hystrix ，我 hystrix 不在你 Controller 这捣蛋，我去其他地方呆着，你有问题了，我再来处理。这才是正确的、合理的设计方式。

**所以解决方案是，把histrix配置在Feign中**

- 在microservice-common中定义OrderClientServiceFallbackFactory

```java
@Component
public class OrderClientServiceFallbackFactory implements FallbackFactory<OrderClientService> {

    @Override
    public OrderClientService create(Throwable throwable) {
        return new OrderClientService() {
            @Override
            public TOrder getOrder(Long id) {
                return new TOrder().setId(id)
                        .setName("未找到该ID的结果")
                        .setPrice(0d)
                        .setDbSource("No this datasource");
            }

            @Override
            public List<TOrder> getAll() {
                return null;
            }
        };
    }
}
```



- 给 Feign 指定 hystrix

```java
@FeignClient(value = "MICROSERVICE-ORDER", fallbackFactory = OrderClientServiceFallbackFactory.class)
public interface OrderClientService {

    @GetMapping("/provider/order/get/{id}")
    TOrder getOrder(@PathVariable(value = "id") Long id);

    @GetMapping("/provider/order/get/list")
    List<TOrder> getAll();
}

```

这样的话，Controller 中的所有方法都可以在 hystrix 里有个默认实现了。

- 开启熔断

