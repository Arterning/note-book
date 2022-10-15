事务ACID特性，其中I代表隔离性(Isolation)。

 

什么是事务的隔离性？
隔离性是指，多个用户的并发事务访问同一个数据库时，一个用户的事务不应该被其他用户的事务干扰，多个并发事务之间要相互隔离。


InnoDB实现了哪几种事务的隔离级别？

按照SQL92标准，InnoDB实现了四种不同事务的隔离级别：

读未提交(Read Uncommitted)

读提交(Read Committed, RC)

可重复读(Repeated Read, RR)

串行化(Serializable)

 

不同事务的隔离级别，实际上是一致性与并发性的一个权衡与折衷。

 

InnoDB的四种事务的隔离级别，分别是怎么实现的？

InnoDB使用不同的锁策略(Locking Strategy)来实现不同的隔离级别。

 

一，读未提交(Read Uncommitted)

这种事务隔离级别下，select语句不加锁。

画外音：官方的说法是

SELECT statements are performed in a nonlocking fashion.

 

此时，可能读取到不一致的数据，即“读脏”。这是并发最高，一致性最差的隔离级别。

 

二，串行化(Serializable)

这种事务的隔离级别下，所有select语句都会被隐式的转化为select ... in share mode.

 

这可能导致，如果有未提交的事务正在修改某些行，所有读取这些行的select都会被阻塞住。

画外音：官方的说法是

To force a plain SELECT to block if other transactions have modified the selected rows.
这是一致性最好的，但并发性最差的隔离级别。
在互联网大数据量，高并发量的场景下，几乎不会使用上述两种隔离级别。



三，可重复读(Repeated Read, RR)
这是InnoDB默认的隔离级别，在RR下：
(1)普通的select使用快照读(snapshot read)，这是一种不加锁的一致性读(Consistent Nonlocking Read)，底层使用MVCC来实现，具体的原理在《InnoDB并发如此高，原因竟然在这？》中有详细的描述；
(2)加锁的select(select ... in share mode / select ... for update), update, delete等语句，
它们的锁，依赖于它们是否在唯一索引(unique index)上使用了唯一的查询条件(unique search condition)，
或者范围查询条件(range-type search condition)：
在唯一索引上使用唯一的查询条件，会使用记录锁(record lock)，而不会封锁记录之间的间隔，即不会使用间隙锁(gap lock)与临键锁(next-key lock)
范围查询条件，会使用间隙锁与临键锁，锁住索引记录之间的范围，避免范围间插入记录，以避免产生幻影行记录，以及避免不可重复的读
画外音：这一段有点绕，多读几遍。
关于记录锁，间隙锁，临键锁的更多说明，详见《InnoDB，select为啥会阻塞insert？》。

 

四，读提交(Read Committed, RC)
这是互联网最常用的隔离级别，在RC下：
(1)普通读是快照读；
(2)加锁的select, update, delete等语句，除了在外键约束检查(foreign-key constraint checking)以及重复键检查(duplicate-key checking)时会封锁区间，其他时刻都只使用记录锁；
此时，其他事务的插入依然可以执行，就可能导致，读取到幻影记录。
总结
并发事务之间相互干扰，可能导致事务出现读脏，不可重复度，幻读等问题
InnoDB实现了SQL92标准中的四种隔离级别
(1)读未提交：select不加锁，可能出现读脏；
(2)读提交(RC)：普通select快照读，锁select /update /delete 会使用记录锁，可能出现不可重复读；
(3)可重复读(RR)：普通select快照读，锁select /update /delete 根据查询条件情况，会选择记录锁，或者间隙锁/临键锁，以防止读取到幻影记录；
(4)串行化：select隐式转化为select ... in share mode，会被update与delete互斥；
InnoDB默认的隔离级别是RR，用得最多的隔离级别是RC
或许有朋友问，为啥没提到insert？可以查阅《InnoDB并发插入，居然使用意向锁？》