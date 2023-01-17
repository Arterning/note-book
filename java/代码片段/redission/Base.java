/**
 * Redisson 是一个在 Redis 的基础上实现的 Java 驻内存数据网格客户端（In-Memory Data Grid）。它不仅提供了一系列的 redis 常用数据结构命令服务，还提供了许多分布式服务，例如分布式锁、分布式对象、分布式集合、分布式远程服务、分布式调度任务服务等等。
相比于 Jedis、Lettuce 等基于 redis 命令封装的客户端，Redisson 提供的功能更加高端和抽象，逼格高！

其实我感觉Redission相比于Jedis没那么好用 不是很直观
而且第一次使用的时候出现了很多很奇怪的问题 比如在图形化工具中设置了一个key
但是使用Redission却怎么也读取不出来 必须要在value的两端加上"" redission才认为是个字符串吗？

refer https://zhuanlan.zhihu.com/p/596334390
 */
public class Base {
    
}
