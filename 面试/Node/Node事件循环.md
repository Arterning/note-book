
![[Pasted image 20240309105635.png]]

总结上述过程：

- 第一阶段：每一次**异步 I/O 的调用**，首先在 nodejs 底层设置请求参数和回调函 callback，形成**请求对象**。
    
- 第二阶段：形成的请求对象，会被放入**线程池**，如果线程池有空闲的 I/O 线程，会执行此次 I/O 任务，得到结果。
    
- 第三阶段：**事件循环**中 **I/O 观察者**，会从请求对象中找到已经得到结果的 I/O 请求对象，取出结果和回调函数，将回调函数放入事件循环中，执行回调，完成整个异步 I/O 任务。
    
- 对于如何感知异步 I/O 任务执行完毕的？以及如何获取完成的任务的呢？ libuv 作为中间层， 在不同平台上，采用手段不同，在 unix 下通过 epoll 轮询，在 Windows 下通过内核（ IOCP ）来实现 ，FreeBSD 下通过 kqueue 实现。
    


观察者实际上就是队列嘛






圈圈外面是宏任务
里面是微任务

可以看到微任务会在执行宏任务的时候穿插执行

![[Pasted image 20240309110421.png]]


#### 一次循环要经过六个阶段：

1. **timers：计时器（setTimeout、setInterval等的回调函数存放在里边）**
2. pending callback 执行延迟到下一个循环的IO回调
3. idle prepare 系统内部使用 
4. poll：检索新的IO事件，执行IO相关的回调 。轮询队列（除timers、check之外的回调存放在这里）
5. **check：检查阶段（使用 setImmediate 的回调会直接进入这个队列）**
6. close callbacks


接下来总结一下 Nodejs 事件循环。

- Nodejs 的事件循环分为 6 大阶段。分别为 timer 阶段，pending 阶段，prepare 阶段，poll 阶段， check 阶段，close 阶段。
    
- nextTick 队列和 Microtasks 队列执行特点，在每一阶段完成后执行， nextTick 优先级大于 Microtasks （ Promise ）。
    
- poll 阶段主要处理 I/O，如果没有其他任务，会处于轮询阻塞阶段。
    
- timer 阶段主要处理定时器/延时器，它们并非准确的，而且创建需要额外的性能浪费，它们的执行还收到 poll 阶段的影响。
    
- pending 阶段处理 I/O 过期的回调任务。
    
- check 阶段处理 setImmediate。 setImmediate 和 setTimeout 执行时机和区别。
    




接下来总结一下 poll 阶段的本质：

- poll 阶段就是通过 timeout 来判断，是否阻塞事件循环。poll 也是一种轮询，轮询的是 i/o 任务，事件循环倾向于 poll 阶段的持续进行，其目的就是更快的执行 I/O 任务。如果没有其他任务，那么将一直处于 poll 阶段。
- 如果有其他阶段更紧急待执行的任务，比如 timer ，close ，那么 poll 阶段将不阻塞，会进行下一个 tick 阶段。


  