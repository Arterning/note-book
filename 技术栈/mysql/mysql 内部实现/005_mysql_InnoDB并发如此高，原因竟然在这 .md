通过并发控制保证数据一致性的常见手段有：

锁（Locking）
数据多版本（Multi Versioning）

共享锁（Share Locks，记为S锁），读取数据时加S锁
排他锁（eXclusive Locks，记为X锁），修改数据时加X锁
共享锁与排他锁的玩法是：
共享锁之间不互斥，简记为：读读可以并行
排他锁与任何锁互斥，简记为：写读，写写不可以并行


数据多版本是一种能够进一步提高并发的方法，它的核心原理是：
（1）写任务发生时，将数据克隆一份，以版本号区分；
（2）写任务操作新克隆的数据，直至提交；
（3）并发读任务可以继续读取旧版本的数据，不至于阻塞；

1. 最开始数据的版本是V0；
2. T1时刻发起了一个写任务，这是把数据clone了一份，进行修改，版本变为V1，但任务还未完成；
3. T2时刻并发了一个读任务，依然可以读V0版本的数据；
4. T3时刻又并发了一个读任务，依然不会阻塞；
可以看到，数据多版本，通过“读取旧版本数据”能够极大提高任务的并发度。
提高并发的演进思路，就在如此：
普通锁，本质是串行执行
读写锁，可以实现读读并发
数据多版本，可以实现读写并发

如果T4时刻来了一个写任务，会怎样？？？
如果T1还没有完成，也可以写，写入版本数据V2
这个时候读任务仍然读取V0版本数据

现在T1写入完毕，V1版本替换V0版本
T2也写入完毕，V2版本也替换V0版本
好像非常完美。。




为什么要有redo日志？
数据库事务提交后，必须将更新后的数据刷到磁盘上，以保证ACID特性。磁盘随机写性能较低，如果每次都刷盘，会极大影响数据库的吞吐量。
优化方式是，将修改行为先写到redo日志里（此时变成了顺序写），再定期将数据刷到磁盘上，这样能极大提高性能。
画外音：这里的架构设计方法是，随机写优化为顺序写，思路更重要。
redo log其实就是do something log...
log buffer--->log file



为什么要有undo日志？
数据库事务未提交时，会将事务修改数据的镜像（即修改前的旧版本）存放到undo日志里，
当事务回滚时，或者数据库奔溃时，可以利用undo日志，即旧版本数据，撤销未提交事务对数据库产生的影响。
对于insert操作，undo日志记录新数据的PK(ROW_ID)，回滚时直接删除；
对于delete/update操作，undo日志记录旧数据row，回滚时直接恢复；
存储undo日志的地方，是回滚段。


InnoDB是基于多版本并发控制的存储引擎
《大数据量，高并发量的互联网业务，一定要使用InnoDB》提到，InnoDB是高并发互联网场景最为推荐的存储引擎，
根本原因，就是其多版本并发控制（Multi Version Concurrency Control, MVCC）。行锁，并发，事务回滚等多种特性都和MVCC相关。
MVCC就是通过“读取旧版本数据”来降低并发事务的锁冲突，提高任务的并发度。

InnoDB的内核，会对所有row数据增加三个内部属性：
(1)DB_TRX_ID，6字节，记录每一行最近一次修改它的事务ID；
(2)DB_ROLL_PTR，7字节，记录指向回滚段undo日志的指针；
(3)DB_ROW_ID，6字节，单调递增的行ID；


总结
InnoDB为何能够做到这么高的并发？
回滚段(undo log)里的数据，其实是历史数据的快照（snapshot），这些数据是不会被修改，select可以肆无忌惮的并发读取他们。
快照读（Snapshot Read），这种一致性不加锁的读（Consistent Nonlocking Read），就是InnoDB并发如此之高的核心原因之一。
所谓的快照读，就是读取前文所说的V0数据
(1)常见并发控制保证数据一致性的方法有锁，数据多版本；
(2)普通锁串行，读写锁读读并行，数据多版本读写并行；
(3)redo日志保证已提交事务的ACID特性，设计思路是，通过顺序写替代随机写，提高并发；
(4)undo日志用来回滚未提交的事务，它存储在回滚段里；
(5)InnoDB是基于MVCC的存储引擎，它利用了存储在回滚段里的undo日志，即数据的旧版本，提高并发；
(6)InnoDB之所以并发高，快照读不加锁；
(7)InnoDB所有普通select都是快照读；

非快照读是指：
select * from t where id>2 lock in share mode;
select * from t where id>2 for update;
问题来了，这些显示加锁的读，是什么读？会加什么锁？和事务的隔离级别又有什么关系？





