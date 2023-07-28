
## 间隙锁是否关闭
show global variables like '%innodb_locks%';


## 事务隔离级别
show global variables like 'tx_isolation';

##X取：
## read uncommitted,
# read committed
# repeatable read
# serializable
set session transaction isolation level read committed ;



## 事务自动提交的参数
show global variables like 'autocommit';
set session autocommit=0;

create database test;
use test;
create table t (

    id int(10) primary key

)engine=innodb;


start transaction;
insert into t values(1);
insert into t values(3);
insert into t values(10);
commit;

select * from t;


