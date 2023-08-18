SET key value
含义：
将字符串值 value 关联到 key 。
如果 key 已经持有其他值， SET 就覆写旧值，无视类型。

SETEX mykey 10 "Hello"
键 key 的值设置为 value ， 并将键 key 的生存时间设置为10秒钟。


SETNX key value
将 key 的值设为 value ，当且仅当 key 不存在。
若给定的 key 已经存在，则 SETNX 不做任何动作。
SETNX 是『SET if Not eXists』(如果不存在，则 SET)的简写。


HMSET runoobkey name "redis tutorial" description "redis basic commands for caching" likes 20 visitors 23000
OK

HGETALL runoobkey
1) "name"----------------key
2) "redis tutorial"------value

3) "description"
4) "redis basic commands for caching"

5) "likes"
6) "20"

7) "visitors"
8) "23000"


runoobkey就相当于java里面的hashmap的变量名




Redis 列表(List)
Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）
LPUSH runoobkey redis
LPUSH runoobkey mongodb
LPUSH runoobkey mysql



Redis 集合(Set)
redis 127.0.0.1:6379> SADD runoobkey redis
(integer) 1
redis 127.0.0.1:6379> SADD runoobkey mongodb
(integer) 1
redis 127.0.0.1:6379> SADD runoobkey mysql
(integer) 1
redis 127.0.0.1:6379> SADD runoobkey mysql


Redis 有序集合(sorted set)
Redis 有序集合和集合一样也是 string 类型元素的集合,且不允许重复的成员。
redis 127.0.0.1:6379> ZADD runoobkey 1 redis
(integer) 1
redis 127.0.0.1:6379> ZADD runoobkey 2 mongodb
(integer) 1
redis 127.0.0.1:6379> ZADD runoobkey 3 mysql
(integer) 1
redis 127.0.0.1:6379> ZADD runoobkey 3 mysql
(integer) 0
redis 127.0.0.1:6379> ZADD runoobkey 4 mysql



Redis 发布订阅 是订阅 runoobChat 频道。
redis 127.0.0.1:6379> SUBSCRIBE runoobChat


往 runoobChat 频道发送消息
redis 127.0.0.1:6379> PUBLISH runoobChat "Redis PUBLISH test"
redis 127.0.0.1:6379> PUBLISH runoobChat "Learn redis by runoob.com"




MULTI 开始一个事务， 然后将多个命令入队到事务中， 最后由 EXEC 命令触发事务


Rdis PipeLine
管道（pipeline）可以一次性发送多条命令给服务端，服务端依次处理完完毕后，通过一条响应一次性将结果返回
pipeline是非原子性
事务是原子性的
https://juejin.cn/post/6844904127001001991

