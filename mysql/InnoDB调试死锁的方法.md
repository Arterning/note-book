超赞，InnoDB调试死锁的方法！
原创 58沈剑 架构师之路 2018-09-05
近期写了不少InnoDB锁相关的文章，不少小伙伴问，如何在MySQL终端模拟并发事务，如何复现之前文章中的案例。今天，咱们一起动起手来，模拟并发事务的互斥与死锁。

 

【事前准备】

安装MySQL服务端

安装MySQL客户端

安装能够模拟多个并发事务的终端

画外音：楼主使用的是MySQL5.6，官方客户端mysql，模拟并发终端用的SecureCRT。

 

【配置的确认与修改】

要测试InnoDB的锁互斥，以及死锁，有几个配置务必要提前确认：

区间锁是否关闭

事务自动提交(auto commit)是否关闭

事务的隔离级别(isolation level)

这几个参数，会影响实验结果。



【事务的隔离级别】辅助材料：

《事务的隔离级别，InnoDB如何实现？》



间隙锁是否关闭

区间锁(间隙锁，临键锁)是InnoDB特有施加在索引记录区间的锁，MySQL5.6可以手动关闭区间锁，它由innodb_locks_unsafe_for_binlog参数控制：

设置为ON，表示关闭区间锁，此时一致性会被破坏（所以是unsafe）

设置为OFF，表示开启区间锁

 

可以这么查询该参数：

show global variables like "innodb_locks%";



【间隙锁，临键锁】辅助材料：

《InnoDB，索引记录上的三种锁》



事务自动提交

MySQL默认把每一个单独的SQL语句作为一个事务，自动提交。

 

可以这么查询事务自动提交的参数：

show global variables like "autocommit";

 

事务的隔离级别

不同事务的隔离级别，InnoDB的锁实现是不一样。



可以这么查询事务的隔离级别：

show global variables like "tx_isolation";

 

可以这么设置事务的隔离级别：

set session transaction isolation level X;

X取：

read uncommitted

read committed

repeatable read

serializable 



这三个参数，MySQL5.6的默认值如上：

OFF，表示使用区间锁

On，表示事务自动提交

RR，事务隔离级别为可重复读

 

要模拟并发事务，需要修改事务自动提交这个选项，每个session要改为手动提交。



任何连上MySQL的session，都要手动执行：

set session autocommit=0;

以手动控制事务的提交。



如上图，需要把session的autocommit设置为OFF。

可以看到，修改session变量，并不影响global变量，全局其他的session仍然是ON。

画外音：session变量默认继承global变量，也可以单独修改。



【数据准备】

InnoDB的行锁都是实现在索引上的，实验可以使用主键，建表时设定为innodb引擎：

create table t (

id int(10) primary key

)engine=innodb;

 

插入一些实验数据：

start transaction;

insert into t values(1);

insert into t values(3);

insert into t values(10);

commit;

 

这是实验的初始状态，不同实验开始之初，都默认回到初始状态。

 

【实验一，间隙锁互斥】

开启区间锁，RR的隔离级别下，上例会有：

(-infinity, 1)

(1, 3)

(3, 10)

(10, infinity)

这四个区间。

 

事务A删除某个区间内的一条不存在记录，获取到共享间隙锁，会阻止其他事务B在相应的区间插入数据，因为插入需要获取排他间隙锁。

 

session A：

set session autocommit=0;

start transaction;

delete from t where id=5;



session B：

set session autocommit=0;

start transaction;

insert into t values(0);

insert into t values(2);

insert into t values(12);

insert into t values(7);

 

事务B插入的值：0, 2, 12都不在(3, 10)区间内，能够成功插入，而7在(3, 10)这个区间内，会阻塞。

 

可以使用：

show engine innodb status;

来查看锁的情况。 



如上图，可以看到（请把图放大）：

insert into t values(7);

正在等待共享间隙锁的释放。

 

如果事务A提交或者回滚，事务B就能够获得相应的锁，以继续执行。



如果事务A一直不提交，事务B会一直等待，直到超时，超时后会显示：

ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction

 

【实验二，共享排他锁死锁】

回到数据的初始状态，这次需要三个并发的session。

画外音：SecureCRT得开三个窗口了。

 

session A先执行：

set session autocommit=0;

start transaction;

insert into t values(7);

 

session B后执行：

set session autocommit=0;

start transaction;

insert into t values(7);

 

session C最后执行：

set session autocommit=0;

start transaction;

insert into t values(7);

 

三个事务都试图往表中插入一条为7的记录：

(1)A先执行，插入成功，并获取id=7的排他锁；

(2)B后执行，需要进行PK校验，故需要先获取id=7的共享锁，阻塞；

(3)C后执行，也需要进行PK校验，也要先获取id=7的共享锁，也阻塞；

 

如果此时，session A执行：

rollback;

id=7排他锁释放。

 

则B，C会继续进行主键校验：

(1)B会获取到id=7共享锁，主键未互斥；

(2)C也会获取到id=7共享锁，主键未互斥；

 

B和C要想插入成功，必须获得id=7的排他锁，但由于双方都已经获取到id=7的共享锁，它们都无法获取到彼此的排他锁，死锁就出现了。

 

当然，InnoDB有死锁检测机制，B和C中的一个事务会插入成功，另一个事务会自动放弃：

ERROR 1213 (40001): Deadlock found when trying to get lock; try restarting transaction

 

【实验三，并发间隙锁的死锁】

共享排他锁，在并发量插入相同记录的情况下会出现，相应的案例比较容易分析。而并发的间隙锁死锁，是比较难定位的。

 

回到数据的初始状态，这次需要两个并发的session，其SQL执行序列如下：

 

A：set session autocommit=0;

A：start transaction;

A：delete from t where id=6;

         B：set session autocommit=0;

         B：start transaction;

         B：delete from t where id=7;

A：insert into t values(5);

         B：insert into t values(8);

 

A执行delete后，会获得(3, 10)的共享间隙锁。

B执行delete后，也会获得(3, 10)的共享间隙锁。

A执行insert后，希望获得(3, 10)的排他间隙锁，于是会阻塞。

B执行insert后，也希望获得(3, 10)的排他间隙锁，于是死锁出现。

 

仍然使用：

show engine innodb status;

来查看死锁的情况。



事务1占有什么锁，请求什么锁；事务2占有什么锁，请求什么锁，一清二楚（请把图放大）。

 

另外，检测到死锁后，事务2自动回滚了：

WE ROLL BACK TRANSACTION (2)

事务1将会执行成功。

 

【总结】

说了很多，希望大家能起手来，这样对InnoDB锁的机制，以及锁的调试印象会更加深刻：

并发事务，间隙锁可能互斥

(1)A删除不存在的记录，获取共享间隙锁；

(2)B插入，必须获得排他间隙锁，故互斥；

并发插入相同记录，可能死锁(某一个回滚)

并发插入，可能出现间隙锁死锁(难排查)

show engine innodb status; 可以查看InnoDB的锁情况，也可以调试死锁

 



架构师之路-分享可落地的架构文章



相关推荐：

《InnoDB，并发如此之高的原因》转发1000+

《InnoDB，巧妙实现四种隔离级别》转发1000+

《InnoDB，索引记录上的三种锁》转发1000+

《InnoDB，RR和RC的快照读有何不同》

《聚集索引与普通索引的差异》转发1000+

《索引，底层是如何实现的？》转发1000+



希望大家有收获，求转。

阅读原文

微信扫一扫
关注该公众号