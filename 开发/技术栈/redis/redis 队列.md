```css
> LPUSH queue Java 码哥字节 Go
(integer) 3

> RPOP queue
"Java"
> RPOP queue
"码哥字节"
> RPOP queue
"Go"
```

`LPUSH、RPOP` 存在一个性能风险，生产者向队列插入数据的时候，List 并不会主动通知消费者及时消费。我们需要写一个 `while(true)` 不停地调用 `RPOP` 指令，当有新消息就会返回消息，否则返回空。

> 要如何避免循环调用导致的 CPU 性能损耗呢？

Redis 提供了 `BLPOP、BRPOP` 阻塞读取的命令，**消费者在在读取队列没有数据的时候自动阻塞，直到有新的消息写入队列，才会继续读取新消息执行业务逻辑。**

`BRPOP queue 0`

参数 0 表示阻塞等待时间无无限制

注意 queue只是一个名字 ，我们可以改成其他

有个想法，用两个服务器，node处理请求，把数据全部push到 redis缓存队列中，另一个php服务器不断的pop这个队列里的数据然后与mysql交互做持久化。大家觉得这么做怎么样？

## 重复消费的问题

重复消费 消息队列为每一条消息生成一个「全局 ID」； 生产者为每一条消息创建一条「全局 ID」，消费者把一件处理过的消息 ID 记录下来判断是否重复。


## **消息可靠性**

> 消费者从 List 中读取一条在消息处理过程中宕机了就会导致消息没有处理完成，可是数据已经没有保存在 List 中了咋办？

本质就是消费者在处理消息的时候崩溃了，就无法再还原消息，缺乏一个消息确认机制。

Redis 提供了 `RPOPLPUSH、BRPOPLPUSH(阻塞)`两个指令，含义是从 List 从读取消息的同时把这条消息复制到另一个 List 中（备份），并且是原子操作。

我们就可以在业务流程正确处理完成后再删除队列消息实现消息确认机制。如果在处理消息的时候宕机了，重启后再从备份 List 中读取消息处理。

```bash
LPUSH redisMQ key value
BRPOPLPUSH redisMQ redisMQBackup #从redisMQ中读取数据并且同时备份到redisMQBackup 
```

如果消费成功则把「redisMQBackup 」的消息删除即可，异常的话可以继续从 「redisMQBackup 」再次读取消息处理。

