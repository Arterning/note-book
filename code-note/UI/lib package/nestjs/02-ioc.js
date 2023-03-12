/**
 * nestjs的IOC设计类似于guice框架
 * 肯定是参考了他的设计
 * 包括关键字module provide等等
 * 现在理解module应该比较容易了 就是一个提供了provider的IOC小容器
 * APP Module再把这些小容器打包起来 是最大的那个容器
 * 所以和Guice很相像
 * 
 * 其实这种细粒度的Bean控制 在springboot中也有体现
 * 也就是定义一个Configuration类 在这个Configuration类中定义返回Bean的方法
 * 也相当于一个小Module了
 * 
 * springboot的自动配置就是基于一个个小的module最后汇聚在一起
 */