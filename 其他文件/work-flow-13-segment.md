# Gson、FastJson、Jackson、json-lib 对比总结

# java 下的 redis 各种 client 工具

# 关于 python 虚拟环境的问题

python 安装一个包，在本地仓库中并没有区分版本
所以 venv 相当于 nodejs 中的 node_module
虽然 package.json 中有定义版本
requirement.txt 中也有定义版本
但是下载到本地之后，是没有版本信息的
所以不同的项目需要把各自的依赖分开
在 nodejs 中使用的方法是每个项目都有自己的 node_module
在 python 中就是使用的 venv 虚拟环境

# springboot spirngcloud springcloud alibaba 的版本需要注意

https://github.com/alibaba/spring-cloud-alibaba/wiki/%E7%89%88%E6%9C%AC%E8%AF%B4%E6%98%8E

# 使用云原生脚手架创建项目

云原生应用脚手架
https://start.aliyun.com/

# Spring Cloud Alibaba 中包含的组件

阿里开源组件
Nacos：一个更易于构建云原生应用的动态服务发现、配置管理和服务管理平台。

Sentinel：把流量作为切入点，从流量控制、熔断降级、系统负载保护等多个维度保护服务的稳定性。

RocketMQ：开源的分布式消息系统，基于高可用分布式集群技术，提供低延时的、高可靠的消息发布与订阅服务。

Dubbo：这个就不用多说了，在国内应用非常广泛的一款高性能 Java RPC 框架。

Seata：阿里巴巴开源产品，一个易于使用的高性能微服务分布式事务解决方案。

Arthas：开源的 Java 动态追踪工具，基于字节码增强技术，功能非常强大。

# 升级 Gradle 之后报错

java.lang.NoClassDefFoundError: javax/xml/bind/JAXBException

add dependency

```gradle
dependencies {
    // JAX-B dependencies for JDK 9+
    implementation "jakarta.xml.bind:jakarta.xml.bind-api:2.3.2"
    implementation "org.glassfish.jaxb:jaxb-runtime:2.3.2"
}
```

[stack overflow](https://stackoverflow.com/questions/43574426/how-to-resolve-java-lang-noclassdeffounderror-javax-xml-bind-jaxbexception)
