传统Hadoop计算过程中，MapReduce任务需要跑很多次，需要多次迭代，每次迭代计算的结果都需要存下来，
存到HDFS，而HDFS本身就是一些硬盘，所以本质上就是把每次计算的结果存到硬盘上。而且存到硬盘上还需要考虑备份，
一般是三次备份。于是计算总时间中一大部分将花到硬盘存储上。之前我们提到程序运行时间，知道它包括四个因素：
计算时间，数据传输时间，任务调度时间，和并行度。在传统MapReduce计算当中，存储占用了大部分时间。

而Spark不同，它是将中间计算的结果放在内存当中，然后在内存中进行迭代计算，速度自然更快。
另外，Spark还存下了计算结果从何而来，即Lineage。如果内存数据丢失，通过Lineage再找父母要，再计算一遍。
虽然重复计算丢失的数据将花费较多时间，但是数据丢失的概率很低，所有Spark整体计算的速度将提升10到100倍。

总结
Spark是一个支持任务调度，监控及分布式部署的通用计算引擎，它通过内存内运算技术和计算关系的血统来提升运算速度。