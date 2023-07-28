set session autocommit=0;

start transaction;

use test;
delete from t where id=5;


commit;
rollback ;