## MQ

消息重复消费？

> 生产端为了保证消息发送成功，可能会重复推送（直到收到成功ACK），会产生重复消息。但是一个成熟的MQ Server框架一般会想办法解决，避免存储重复消息（比如：空间换时间，存储已处理过的message_id），给生产端提供一个幂等性的发送消息接口。
> 
> 但是消费端却无法根本解决这个问题，在高并发标准要求下，拉取消息+业务处理+提交消费位移需要做事务处理，另外消费端服务可能宕机，很可能会拉取到重复消息。
> 
> 所以，只能业务端自己做控制，对于已经消费成功的消息，本地数据库表或Redis缓存业务标识，每次处理前先进行校验，保证幂等。

保证消息有序？

TODO

事务消息的实现？

事务消息可以实现分布式事务

TODO
