## IDEA安装spring boot插件

在settings -> Plugins 里边搜Spring Assistant，安装完后重启idea



![](../../images/搜狗截图20190228194041.png)



![](../../images/搜狗截图20190228194118.png)

这样就很方便了，如果需要加入sql,mq,web,直接在这里选中就可以了。

## 注解意思总结

```java
@PropertySource(value = {"classpath:person.properties"}) //加载配置文件
@ConfigurationProperties(prefix = "person") //加载配置文件的那个配置项目
//上面两个注解用于定义一个配置Bean，这个Bean就只负责保存各种配置数据。
//比如HttpEncodingProperties

@ImportResource(locations = {"classpath:beans.xml"})//加载spring文件

@Configuration //指明当前类是一个配置类；就是来替代之前的Spring配置文件beans.xml，和@Bean配合使用
@Bean //定义在方法上，将方法的返回值添加到容器中；容器中这个组件默认的id就是方法名
//所以可以看出，@Configuration就相当于xml文件，@Bean相当于<bean>标签。

@EnableConfigurationProperties(HttpEncodingProperties.class)//用于加载配置Bean
//比如HttpEncodingAutoConfiguration就用这个注解来加载配置Bean来完成自动配置
//xxxxAutoConfiguration：帮我们给容器中自动配置组件；
//xxxxProperties:配置类来封装配置文件的内容；
```



