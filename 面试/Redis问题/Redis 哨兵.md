
上一节实现了Redis集群的主从模式，这一节来实现一下哨兵模式。

有了主从模式，为什么还会需要哨兵模式呢？

上一节我们了解到，为了缓解服务器压力，我们可以对redis的系统实现一个主从分离，一个master主要实现写功能，一个或多个 slave 主要负责读功能。

那么这种模式有个什么问题呢，

万一我们的master 突然宕机了，不能提供写服务了，只能依靠slave实现读的功能，这种情况是很糟糕的。

所以为了解决这种困境，前辈们又提出了一种解决方式，哨兵模式。



## 概念

Redis 官方提供一个高可用方案——**哨兵（Sentinel）**

首先介绍一下哨兵模式的大致架构，还是一主多从模式，不过我们还需要加上一个或多个能提供监测服务的哨兵（sentinel）。

这个哨兵，或者说几个哨兵，能够实时检测我们的master的状态，当master处于正常运行状态时，一切都OK，平安夜。

一旦master出现故障，服务器宕机，内存满了等问题导致master不能再提供服务了。

哨兵就会马上监测到这个情况，在默认的等待时间30s后，如果master还是不能使用，那这些哨兵就会自动地从slave中间挑选一个作为 master，重新形成一个 一主多从的局面。



**Sentinel配置**

Redis源码中包含一个名为sentinel.conf的文件，是一个你可以用来配置Sentinel的示例配置文件。一个典型的最小配置文件像下面这样：

```
sentinel monitor mymaster 127.0.0.1 6379 2
sentinel down-after-milliseconds mymaster 60000
sentinel failover-timeout mymaster 180000
sentinel parallel-syncs mymaster 1

sentinel monitor resque 192.168.1.3 6380 4
sentinel down-after-milliseconds resque 10000
sentinel failover-timeout resque 180000
sentinel parallel-syncs resque 5
```

sentinel monitor参数的意思在下面

```
sentinel monitor <master-group-name> <ip> <port> <quorum>
```

为了更加清晰明了，让我们一行一行来检查配置选项的意思：

第一行用来告诉Redis监控一个叫做_mymaster_的主节点，地址是 127.0.0.1 端口号是6379，并且有2个仲裁机器。所有的意思都很明显，但是除了这个**quorum** 参数：

- **quorum** 是 需要同意主节点不可用的Sentinels的数量
- 然而**quorum** 仅仅只是用来检测失败。为了实际的执行故障转移，Sentinels中的一个需要被选定为leader并且被授权进行操作，这仅仅发生在大多数Sentinels进行投票的时候。


比如如果你有五个Sentinel进程，对于一个主节点quorum被设置为2，下面是发生的事情：

同时有两个Sentinels同意主节点不可用，其中的一个将会尝试开始故障转移。
如果至少有三个Sentinels是可用的，故障转移将会被授权并且开始
实际中，这意味着在失败时，如果大多数的Sentinel进程没有同意，Sentinel永远不会开始故障转移。



其他的Sentinels选项

其他的选项几乎都是如下形式：

```
sentinel <option_name> <master_name> <option_value>
```
用途如下：

down-after-milliseconds：当一个实例失去联系（要么不回复我们的请求，要么回复一个错误）超过了这个时间（毫秒为单位），Sentinel就开始认为这个实例挂掉了。

parallel-syncs：设置的从节点的数量，这些从节点在一次故障转移过后可以使用新的主节点进行重新配置。数量越少，完成故障转移过程将花费更多的时间，如果从节点为旧的数据提供服务，你或许不想所有的从节点使用主节点进行重新同步。复制进程对于从节点来说大部分是非阻塞的，还是有一个时刻它会停下来去从主节点加载数据。你或许想确保一次只有一个从节点是不可达的，可以通过设置这个选项的值为1来完成。

别的选项在文章的其他部分进行描述。





**运行Sentinel**

如果你使用redis-sentinel可执行文件，你可以使用下面的命令来运行Sentinel：

```
redis-sentinel /path/to/sentinel.conf
```

另外，你可以直接使用redis-server并以Sentinel模式来启动：

```
redis-server /path/to/sentinel.conf --sentinel
```

两种方式是一样的。











