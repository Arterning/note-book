在Node.js中，事件循环是处理异步操作的核心机制。事件循环包括两种主要类型的任务：宏任务（Macro Tasks）和微任务（Micro Tasks）。

1. 宏任务（Macro Tasks）：
   - 宏任务通常代表一些较大的、离散的工作单元，比如I/O操作、定时器回调、网络请求等。
   - 宏任务由事件循环的不同阶段处理，例如定时器阶段、I/O阶段等。
   - 宏任务的执行是异步的，它们会在事件循环的不同轮次中执行。

2. 微任务（Micro Tasks）：
   - 微任务通常代表一些较小的、即时的任务，比如Promise的回调函数、process.nextTick等。
   - 微任务在事件循环的每个阶段之间执行，确保它们在宏任务之前执行。
   - 微任务通常在同一个事件循环轮次内按顺序执行，直到队列为空。

在Node.js中，事件循环的顺序大致如下：

1. 执行当前宏任务队列中的一个宏任务。
2. 执行所有当前微任务队列中的微任务，直到微任务队列为空。
3. 检查是否有定时器超时，执行相关的宏任务。
4. 执行I/O操作对应的宏任务。
5. 重复步骤1-4，直到没有宏任务和微任务。

示例：

```javascript
console.log('Start');

// 宏任务
setTimeout(() => {
  console.log('Timeout (Macro Task)');
}, 0);

// 微任务
Promise.resolve().then(() => {
  console.log('Promise (Micro Task)');
});

console.log('End');

// 输出顺序：
// Start
// End
// Promise (Micro Task)
// Timeout (Macro Task)
```

在上面的示例中，"Start" 和 "End" 是宏任务，而Promise的回调函数和setTimeout的回调函数是微任务和宏任务，它们的执行顺序符合事件循环规则。