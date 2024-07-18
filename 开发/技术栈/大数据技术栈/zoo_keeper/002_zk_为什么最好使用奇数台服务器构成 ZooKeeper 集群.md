我们知道在Zookeeper中 Leader 选举算法采用了Zab协议。Zab核心思想是当多数 Server 写成功，则任务数据写成功。①如果有3个Server，则最多允许1个Server 挂掉。②如果有4个Server，则同样最多允许1个Server挂掉。既然3个或者4个Server，同样最多允许1个Server挂掉，那么它们的可靠性是一样的，所以选择奇数个ZooKeeper Server即可，这里选择3个Server。

作者：JavaGuide
链接：https://juejin.cn/post/6844903677367418893
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//TODO zab协议