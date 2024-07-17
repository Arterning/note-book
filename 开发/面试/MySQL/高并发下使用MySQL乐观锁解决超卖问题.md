
在秒杀系统设计中，超卖是一个经典、常见的问题，任何商品都会有数量上限，如何避免成功下订单买到商品的人数不超过商品数量的上限，这是每个抢购活动都要面临的难点。

## 超卖问题描述

在多个用户同时发起对同一个商品的下单请求时，先查询商品库存，再修改商品库存，会出现资源竞争问题，导致库存的最终结果出现异常。

问题：当商品A一共有库存15件，用户甲先下单10件，用户乙下单8件，这时候库存只能满足一个人下单成功，如果两个人同时提交，就出现了超卖的问题。

![[Pasted image 20240403125611.png]]


可以采用多种方式解决超卖问题。使用synchronized可以保证数据一致性，但是效率低，并且分布式环境下无用；使用数据库锁表会造成数据库性能低下。**单体条件下，采用乐观锁是比较合适的方式，集群可以考虑分布式锁**。


为什么会出现超卖问题呢?

因为如果你的代码没有加锁 而查询库存和修改库存是2个分开的操作

```sql
select num from stock where sku = ?

update stock set num = num - #{count} where sku = ? and num - #{count} > 0;
```

即使你加了事务,也只能保证这2个操作同时成功





## 乐观锁实现

```sql

select num,version from stock where sku = ?

update stock set num = num -1, version = new_version where version = ? and sku = ?
```


# 悲观锁实现

需要注意在sku字段上加索引, 这样就只会锁住一行记录. 否则会出现锁表的情况

```sql

select num from stock where sku = ? for udpate;
update stock set num = num -1 where sku = ?;
```




