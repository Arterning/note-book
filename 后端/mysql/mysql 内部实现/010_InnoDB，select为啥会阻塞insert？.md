InnoDB，select为啥会阻塞insert？
MySQL的InnoDB的细粒度行锁，是它最吸引人的特性之一。


临键锁，是记录锁与间隙锁的组合，它的封锁范围，既包含索引记录，又包含索引区间。
更具体的，临键锁会封锁索引记录本身，以及索引记录之前的区间。
如果一个会话占有了索引记录R的共享/排他锁，其他会话不能立刻在R之前的区间插入新的索引记录。
画外音：原文是说
If one session has a shared or exclusive lock on record R in an index, 
another session cannot insert a new index record in the gap immediately before R in the index order.
临键锁的主要目的，也是为了避免幻读(Phantom Read)。如果把事务的隔离级别降级为RC，临键锁则也会失效。

五、总结
(1)InnoDB的索引与行记录存储在一起，这一点和MyISAM不一样；
(2)InnoDB的聚集索引存储行记录，普通索引存储PK，所以普通索引要查询两次；
(3)记录锁锁定索引记录；
(4)间隙锁锁定间隔，防止间隔中被其他事务插入；
(5)临键锁锁定索引记录+间隔，防止幻读；
