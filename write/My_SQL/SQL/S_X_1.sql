use test;
set session autocommit=0;
select  * from t;

start transaction;

insert into t values(7);


rollback ;