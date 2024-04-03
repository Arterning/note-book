
在面试场景中，只要和 React 相关，面试官一定舔着脸问你：“ **宝子，setState 是同步还是异步的呀** ” 。面对这样的无耻刁难，我们需要先明确，从 API 层面上说，它就是普通的调用执行的函数，自然是同步 API 。

因此，这里所说的同步和异步指的是 API 调用后更新 DOM 是同步还是异步的。

  
  
先说结论，首先，**同步和异步主要取决于它被调用的环境**。

- **如果 setState 在 React 能够控制的范围被调用，它就是异步的。**

比如合成事件处理函数, 生命周期函数, 此时会进行批量更新, 也就是将状态合并后再进行 DOM 更新。

- **如果 setState 在原生 JavaScript 控制的范围被调用，它就是同步的。**

比如原生事件处理函数中, 定时器[回调函数](https://www.zhihu.com/search?q=%E5%9B%9E%E8%B0%83%E5%87%BD%E6%95%B0&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A2710857100%7D)中, Ajax 回调函数中, 此时 setState 被调用后会立即更新 DOM 。

  
  
其实，**我们看到的所谓的 “异步”，是开启了 “批量更新” 模式的。**

批量更新模式可以减少真实 DOM 渲染的次数，所以只要是 React 能够控制的范围，出于性能因素考虑，一定是批量更新模式。批量更新会先合并状态，再一次性做 DOM 更新。



setState 异步（或者说是批量更新）的一个重要动机就是**避免频繁的 re-render**。




在实际的 React 运行时中，setState 异步的实现方式有点类似于浏览器里的 Event-Loop：

每来一个setState，就把它塞进一个队列里。等时机成熟，再把队列里的 state 结果做合并，最后只针对最新的 state 值走一次更新流程。

这个过程，叫作“批量更新”，批量更新的过程正如下面代码中的箭头[流程图](https://www.zhihu.com/search?q=%E6%B5%81%E7%A8%8B%E5%9B%BE&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A2710857100%7D)所示：

![[Pasted image 20240323100044.png]]



只要我们的同步代码还在执行，“进队列” 这个动作就不会停止。因此就算我们在React 中写了一个 N 次的 setState 循环，也只是会增加 state 任务入队的次数，并不会带来频繁的 re-render。当 N 次调用结束后，仅仅是 state 的任务队列内容发生了变化， state 本身并不会立刻改变。

  
如果为非批量更新模式，调用多少次 setState 就会渲染多少次真实 DOM，性能较低。


