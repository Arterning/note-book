spark是一个计算引擎
但是不是给予单台机器的，而是基于分布式机器的
可以在运行在计算集群上，而且计算是在内存中进行
这是他独特的地方


比如我有一个计算量很大的任务，而且我手里有多台服务器
而且这个计算任务是可以拆分的，（如果计算单元之间互相有依赖关系怎么办？？）
那么这个时候我就可以考虑使用spark....
