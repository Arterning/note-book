复杂sql

多表查询 select where /select join on都可以实现多表查询，但是是有区别的，注意。

子查询（查询嵌套) select ...xx=(select ..)或者seleect ..xx in(select ..)

group by

having

exist



select for update





type: system>const>eq_ref>**ref>range>index**>all

system和const,eq_ref一般达不到，我们一般达到ref和range就可以了。

system/const:结果只有一条数据

eq_ref:结果多条，但是每条数据唯一

ref:结果多条，每条数据可以是0条或者多条



range:检索指定范围的行，where后面是一个范围，in/between and.. 大于小于



index:查询全部索引中的数据 select 索引列 from table

all:查询全表数据 select 非索引列 from ..





索引长度

在utf8一个字符占用3个字节，如果索引字段可以为null，则会使用一个字节用于标志。

用2个字节标志可变长度。



复合索引数据结构是如何实现的？？



ref：指明当前表参照的字段，比如

select ..wher a.c=b.x;(参照了b表中的x,如果b.x是常量，那么是const)



Extra:

using filesort 性能消耗巨大，需要额外的一次排序。

避免：select哪些字段，就order哪些字段。

using temporary:性能损耗大，用到了临时表，一般出现于group by 语句中。

避免：select哪些字段，就按照哪些列group





复合索引相当于二级目录

create index index_name on table(c1,c2,c3);

什么样的查询会使用到复合索引，是这样的

select * from table where c1="xx" and c2="xx" and c3="xx";

where里面必须包括复合索引里面的所有列，才会起作用。

而且最好where的顺序就是复合索引的顺序。

![](D:\code\笔记\图片\55.png)







```sql
grant replication slave,reload,super ON *.* TO 'root'@'192.168.0.%' IDENTIFIED BY '1995';
flush privileges;
```



```sql
change master to
master_host='192.168.0.10',
master_user='root',
master_password='1995',
master_port=3306,
master_log_file='mysql-bin.000001',
master_log_pos=707;
```





从机

start slave;

show slave status;

