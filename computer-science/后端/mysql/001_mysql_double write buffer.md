double write buffer，你居然没听过？
https://cloud.tencent.com/developer/article/1603076


MySQL的buffer一页的大小是16K，文件系统一页的大小是4K，也就是说，MySQL将buffer中一页数据刷入磁盘，要写4个文件系统里的页
如果执行到一半断电，会不会出现问题呢？
会，这就是所谓的“页数据损坏”。

如上图所示，MySQL内page=1的页准备刷入磁盘，才刷了3个文件系统里的页，掉电了，
则会出现：重启后，page=1的页，物理上对应磁盘上的1+2+3+4四个格，数据完整性被破坏。


Double Write Buffer，但它与传统的buffer又不同，它分为内存和磁盘的两层架构。
当有页数据要刷盘时：
第一步：页数据先memcopy到DWB的内存里；
第二步：DWB的内存里，会先刷到DWB的磁盘上；写第一次
第三步：DWB的内存里，再刷到数据磁盘存储上；写第二次（写了2次，所以是double write）

第一步失败，没关系，可以从redlog恢复
第二步失败，没关系，可以从redlog恢复
第三步失败，DWB磁盘上还是有完整的数据
写了2次，总有一个地方的数据是OK的。

结尾
MySQL有很强的数据安全性机制：
（1）在异常崩溃时，如果不出现“页数据损坏”，能够通过redo恢复数据；
（2）在出现“页数据损坏”时，能够通过double write buffer恢复页数据；
double write buffer：
（1）不是一个内存buffer，是一个内存/磁盘两层的结构，是InnoDB里On-Disk架构里很重要的一部分；
（2）是一个通过写两次，保证页完整性的机制；
