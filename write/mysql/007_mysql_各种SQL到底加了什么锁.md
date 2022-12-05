一、普通select
(1)在读未提交(Read Uncommitted)，读提交(Read Committed, RC)，可重复读(Repeated Read, RR)这三种事务隔离级别下，
普通select使用快照读(snpashot read)，不加锁，并发非常高；

二、加锁select
加锁select主要是指：
select ... for update
select ... in share mode

(1)如果，在唯一索引(unique index)上使用唯一的查询条件(unique search condition)，
会使用记录锁(record lock)，而不会封锁记录之间的间隔，即不会使用间隙锁(gap lock)与临键锁(next-key lock)；

select * from t where id=1 for update;
只会封锁记录，而不会封锁区间。
update t set name=xxx where id=1;
也只加记录锁

同样是写操作，insert和update与delete不同，它会用排它锁封锁被插入的索引记录，而不会封锁记录之前的范围。
同时，会在插入区间加插入意向锁(insert intention lock)，但这个并不会真正封锁区间，也不会阻止相同区间的不同KEY插入。

总结
record lock
gap lock
netxt-key lock