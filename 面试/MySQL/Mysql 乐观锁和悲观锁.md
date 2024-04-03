
## Row Lock与Table Lock

上面我们提到，使用select…for update会把数据给锁住，不过我们需要注意一些锁的级别，MySQL  InnoDB默认Row-Level Lock，所以只有「明确」地指定主键，MySQL 才会执行Row lock (只锁住被选取的数据)  ，否则MySQL 将会执行Table Lock (将整个数据表单给锁住)。


## 乐观锁


version 

```
version = select version from table ;
update table set num = num + 1 where version = {version} and id = {id}

```


timstamp
```
update_time = select update_time from table ;
update table set num = num + 1, update_time = now() where id = {id} and update_time = {update_time}
```


如果更新失败，那么affected_row就肯定是0，我们可以判断affected_row来判断是否操作成功。


## 悲观锁

xxl-job 调度中心集群就是通过悲观锁来防止重复调度的。


```
MySQL select…for update

select id from lock_table where id = '12234' for update;
```

