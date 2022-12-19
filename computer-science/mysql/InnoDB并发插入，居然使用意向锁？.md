//TODO MYSQL
InnoDB并发插入，居然使用意向锁？
原创 58沈剑 架构师之路 2018-08-19
《插入InnoDB自增列，居然是表级别锁？》介绍了InnoDB所使用的七种锁中的一种，自增锁。



今天，将要介绍InnoDB另外三种：共享/排他锁，意向锁，插入意向锁。



一，共享/排它锁(Shared and Exclusive Locks)

《InnoDB并发为何这么高？》一文介绍了通用的共享/排它锁，在InnoDB里当然也实现了标准的行级锁(row-level locking)，共享/排它锁：

(1)事务拿到某一行记录的共享S锁，才可以读取这一行；

(2)事务拿到某一行记录的排它X锁，才可以修改或者删除这一行；

 

其兼容互斥表如下：

          S          X

S      兼容      互斥

X      互斥      互斥



即：

(1)多个事务可以拿到一把S锁，读读可以并行；

(2)而只有一个事务可以拿到X锁，写写/读写必须互斥；

 

共享/排它锁的潜在问题是，不能充分的并行，解决思路是数据多版本，具体思路在《InnoDB并发为何这么高？》里介绍过，这里不再深入展开。

 

二，意向锁(Intention Locks)

InnoDB支持多粒度锁(multiple granularity locking)，它允许行级锁与表级锁共存，实际应用中，InnoDB使用的是意向锁。

 

意向锁是指，未来的某个时刻，事务可能要加共享/排它锁了，先提前声明一个意向。

 

意向锁有这样一些特点：

(1)首先，意向锁，是一个表级别的锁(table-level locking)；

(2)意向锁分为：

意向共享锁(intention shared lock, IS)，它预示着，事务有意向对表中的某些行加共享S锁

意向排它锁(intention exclusive lock, IX)，它预示着，事务有意向对表中的某些行加排它X锁

 

举个例子：

select ... lock in share mode，要设置IS锁；

select ... for update，要设置IX锁；

 

(3)意向锁协议(intention locking protocol)并不复杂：

事务要获得某些行的S锁，必须先获得表的IS锁

事务要获得某些行的X锁，必须先获得表的IX锁

 

(4)由于意向锁仅仅表明意向，它其实是比较弱的锁，意向锁之间并不相互互斥，而是可以并行，其兼容互斥表如下：

          IS          IX

IS      兼容      兼容

IX      兼容      兼容

 

(5)额，既然意向锁之间都相互兼容，那其意义在哪里呢？它会与共享锁/排它锁互斥，其兼容互斥表如下：

          S          X

IS      兼容      互斥

IX      互斥      互斥

画外音：排它锁是很强的锁，不与其他类型的锁兼容。这也很好理解，修改和删除某一行的时候，必须获得强锁，禁止这一行上的其他并发，以保障数据的一致性。

 

三，插入意向锁(Insert Intention Locks)

对已有数据行的修改与删除，必须加强互斥锁X锁，那对于数据的插入，是否还需要加这么强的锁，来实施互斥呢？插入意向锁，孕育而生。

 

插入意向锁，是间隙锁(Gap Locks)的一种（所以，也是实施在索引上的），它是专门针对insert操作的。

画外音：有点尴尬，间隙锁下一篇文章才会介绍，暂且理解为，它是一种实施在索引上，锁定索引某个区间范围的锁。

 

它的玩法是：

多个事务，在同一个索引，同一个范围区间插入记录时，如果插入的位置不冲突，不会阻塞彼此。

画外音：官网的说法是

Insert Intention Lock signals the intent to insert in such a way that multiple transactions inserting into the same index gap need not wait for each other if they are not inserting at the same position within the gap.

 

这样，之前挖坑的例子，就能够解答了。

 

在MySQL，InnoDB，RR下：

t(id unique PK, name);

 

数据表中有数据：

10, shenjian

20, zhangsan

30, lisi

 

事务A先执行，在10与20两条记录中插入了一行，还未提交：

insert into t values(11, xxx);

 

事务B后执行，也在10与20两条记录中插入了一行：

insert into t values(12, ooo);

 

(1)会使用什么锁？

(2)事务B会不会被阻塞呢？

 

回答：虽然事务隔离级别是RR，虽然是同一个索引，虽然是同一个区间，但插入的记录并不冲突，故这里：

使用的是插入意向锁

并不会阻塞事务B

 

思路总结

(1)InnoDB使用共享锁，可以提高读读并发；

(2)为了保证数据强一致，InnoDB使用强互斥锁，保证同一行记录修改与删除的串行性；

(3)InnoDB使用插入意向锁，可以提高插入并发；

 

结尾

假设不是插入并发，而是读写并发，又会是什么样的结果呢？

 

MySQL，InnoDB，默认的隔离级别(RR)。

t(id unique PK, name);

 

数据表中有数据：

10, shenjian

20, zhangsan

30, lisi

 

事务A先执行，查询了一些记录，还未提交：

select * from t where id>10;

 

事务B后执行，在10与20两条记录中插入了一行：

insert into t values(11, xxx);

 

这里：

(1)会使用什么锁？

(2)事务B会不会被阻塞呢？

 

下回分解。



相关文章：

《InnoDB并发为何这么高？》

《插入InnoDB自增列，居然是表级别锁？》

《带团队，要不要言传身教？》



知识，即使一小点，也是好的，转。

阅读原文

微信扫一扫
关注该公众号