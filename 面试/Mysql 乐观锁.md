
**MySQL select…for update的Row Lock与Table Lock**

上面我们提到，使用select…for update会把数据给锁住，不过我们需要注意一些锁的级别，MySQL  InnoDB默认Row-Level Lock，所以只有「明确」地指定主键，MySQL 才会执行Row lock (只锁住被选取的数据)  ，否则MySQL 将会执行Table Lock (将整个数据表单给锁住)。





version 

```
update table set num = num + 1 where version = {version} and id = {id}
```


timstamp
```
update table set num = num + 1, update_time = now() where id = {id} and update_time = {update_time}
```