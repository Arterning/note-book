Spring Boot 的核心注解是哪个？它主要由哪几个注解组成的？
启动类上面的注解是@SpringBootApplication，它也是 Spring Boot 的核心注解，主要组合包含了以下 3 个注解： 

@SpringBootApplication
@SpringBootConfifiguration：组合了 @Confifiguration 注解，实现配置文件的功能。 
@EnableAutoConfifiguration：打开自动配置的功能，也可以关闭某个自动配置的选项， 例 如： java 如关闭数据源自动配置功能： @SpringBootApplication(exclude = { DataSourceAutoConfiguration.class })。 


SpringBoot Starter的工作原理
我个人理解SpringBoot就是由各种Starter组合起来的，我们自己也可以开发Starter 
在sprinBoot启动时由@SpringBootApplication注解会自动去maven中读取每个starter中的spring.factories文件,该文件里配置了所有需要被创建spring容器中的bean，并且进行自动配置把bean注入SpringContext中 //（SpringContext是Spring的配置文件）

Spring Boot 有哪几种读取配置的方式？
Spring Boot 可以通过 @PropertySource,@Value,@Environment, @ConfifigurationPropertie注解来绑定变量


Spring Boot 配置加载顺序？
1. properties文件
2. yaml文件
3. 系统环境变量
4. java opts参数

YAML 配置的优势在哪里 ?
简洁

spring boot 核心配置文件是什么？bootstrap.properties 和 application.properties 有何区别 ?


什么是 Spring Profiles？


比较一下 Spring Security 和 Shiro 各自的优缺点 ?


Spring Boot 中如何解决跨域问题 ?
在可以通过实现WebMvcConfifigurer接口然后重写addCorsMappings方法解决跨域问题。 


SpringBoot性能如何优化
如果项目比较大，类比较多，不使用@SpringBootApplication，采用@Compoment指定扫包范围
在项目启动时设置JVM初始内存和最大内存相同 
将springboot内置服务器由tomcat设置为undertow


Spring Boot 中的 starter 到底是什么 ?
可以理解为对依赖的一种合成  starter会把一个或者一套功能相关的依赖全部包含进来 避免自己去手动找依赖 设置版本
避免了各种冲突问题 提高了开发效率
就是基于自动化配置类的自动装配 而且相关的自动配置都提供了默认值
比如redis starter就是基于 RedisAutoConfiguration 通过条件注解决定配置是否生效
比如是否引入了某些jar包 是否定义了属性


spring-boot-starter-parent 有什么用 ?
定义了版本号 其他starter就不需要定义版本号了
执行打包操作


SpringBoot异常处理相关注解?
@ControllerAdvice 
@ExceptionHandler






而对于 Spring Bean 的生命周期来说：
实例化 Instantiation   构造函数
属性赋值 Populate      getter/setter
初始化 Initialization  init-method 所指定的方法
销毁 Destruction       destory-method 所指定的方法
https://segmentfault.com/a/1190000040365130





