
## Spring Boot Security 基本流程

Spring Security的整体原理为：

1. 当http请求进来时，使用severlet的Filter来拦截。
2. 提取http请求中的认证信息，例如username和password，或者Token。
3. 从数据库（或者其他地方，例如Redis）中查询用户注册时的信息，然后进行比对，相同则认证成功，反之失败。

其实spring security整体也是这样的，只是流程化后，兼顾扩展导致搞的很复杂。

`springboot 2.7.0`以后弃用了`WebSecurityConfigurerAdapter`，所以我们直接采用最新的写法，网上教程很多都是老的写法。


## proxyBeanMethods

`@Configuration(proxyBeanMethods = false)` 里的proxyBeanMethods = false 什么作用

当proxyBeanMethods = false 时，`@Bean`不使用代理，不进入容器，每次获取到的对象都是不同的对象

这样可以加快启动速度。

当proxyBeanMethods = true 时，多次获取bean，都是同一个对象。