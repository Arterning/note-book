```java
Socket socket = server.accept();//这里是阻塞的，这里等待客户端连接
InputStream.read(byte[] in);//这里也阻塞的。
```

所以，我们就是来一个请求就用线程池里面的线程来运行。

不然只有一个线程的话，当一个用户访问的时候，其他用户都只能阻塞，那就完蛋了。



但是还有问题，一个线程只能为一个消费者服务，相当于一个服务生只能服务一个客人，这样非常消耗性能，所以还有更合理的解决方案。一个线程为多个客户端处理任务。



老的tomcat是socket实现的



ServerSocketChannel->ServerSocket

SocketChannel->Socket

Selector负责监听

javaNIO可以实现单线程多个客户端服务。

一个线程有select就可以为多个客户端服务。

![](../images/搜狗截图20190225093315.png)



## java NIO几个要素

1. Channel 相当于Strem
2. Buffer
3. Selector



```java
 selector.select();//是阻塞的，为什么说是nio
 selector.select(1000);//可以不阻塞
 selector.wakeup();//可以唤醒selector
// SelectionKey.OP_READ表示底层缓冲区是否有空间
```



## netty的应用领域

1. 分布式的进程通信
2. 游戏服务器开发，页游，手游







