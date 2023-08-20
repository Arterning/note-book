多线程DAG调度
DAG即Directed Acyclic Graph,有向无环图的意思，DAG调度的目的就是把一个作业分成不同阶段，根据依赖构建一张DAG，并进入到阶段内部，把阶段划分为可以并行计算的任务，最后再把一个阶段内的所有任务交付给任务调度器来收尾。

多核和多CPU编程——有向无环图DAG应用
https://blog.csdn.net/fpcc/article/details/124208576

https://juejin.cn/post/6911518103402872846
