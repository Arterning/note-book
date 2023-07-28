//Why is async required to call await inside a JavaScript function body? 
//https://stackoverflow.com/questions/66113393/why-is-async-required-to-call-await-inside-a-javascript-function-body?noredirect=1&lq=1
//https://stackoverflow.com/questions/44894691/why-await-requires-async-in-function-definition

/**
 * 
 * In ECMAScript language versions prior to 2015, await was not a keyword. Marking a function async provides a syntactic "bailout" to indicate a breaking change in the language grammar within the body of the function.

This is the most important reason. Without the async keyword, all programs written in ECMAScript 5 or older would no longer work if they used the await keyword as a variable (in fact this was done intentionally in some cases as a polyfill before async/await was standardized), since that would cause a breaking change without the addition of async to the specification. Because of this, async is syntactically necessary to avoid breaking changes to the language.

简单来说就是ECMA5之前是没有await这个关键字的 这次加了这个关键字
那么之前用await作为变量名的程序或者代码就都无法工作了
所以必须增加限制 只有在async function内 await才是关键字。。。好像是这个意思？？

由此可以看出 给语言增加关键字是没那么容易的。


 * 
 */

/**
 *  "asynchronous" means "does not block the calling thread".
    "synchronous" means "blocks the calling thread".
 */

/**
 * 
 * async function task() {
 *      const result = await fetch(url);
 *      console.log('in');
 *      return result;
 * }
 * 
 * task();
 * console.log('out');
 */

/**
 * 上面这段代码先打印out 
 * 然后打印in
 * 从这一点可以看出
 * async 声明了一个 异步函数，
 * 这个 异步函数 体内有一行 await 语句，
 * 它告示了该行为同步执行，并且与上下相邻的代码是依次逐行执行的。就是说：

    async 函数执行，返回了一个 promise 对象, 不会阻塞后面的代码, 所以可以理解 task 是一个异步方法
    await 所在的那一行语句是同步的
    既然返回promise, 上面task就可以这样

    task 整体是一个异步函数，但是内部的await是同步, 简称“外异内同”

 */

/**
 * 如果我们修改一下代码 
 * 这种情况下 就是先打印in  然后打印out
 */
/**
 * 
 * async function task() {
 *      const result = await fetch(url);
 *      console.log('in');
 *      return result;
 * }
 * 
 * const tt = await task();
 * console.log('out');
 */


/**
 * 在redis client 工具lettuce中
 * 也是这么表示的
 * sync代表同步获取命令的返回结果 也就是会等待
 * async则是异步 通过回调函数获取结果 不会等待
 */

/**
 * 
 * JS中的 async 本质上是JS Promise的语法糖 他到底做了上面 
 * 可以看这一段代码
 * 
 * Consider the following two functions:

function foo() {
  if (Math.random() < 0.5) {
    return 'return';
  } else {
    throw 'throw';
  }
}

async function bar() {
  if (Math.random() < 0.5) {
    return 'return';
  } else {
    throw 'throw';
  }
}

async performs the following transformation of function bar():

function bar() {
  return new Promise((resolve, reject) => {
    try {
      if (Math.random() < 0.5) {
            str = 'return';
          } else {
            str = 'throw';
          }
      }
      resolve(str);
      } catch (reason) {
        reject(reason);
      }
    });
  }

可以看到实际上就是用Promsise把结果包装了一下
然后await相当于then 把结果获取出来
 */

/**
 * 
 * 之所以当时对为什么叫async有疑惑
 * 因为await代表会等待 应该叫做同步 sync才对啊？
 * 
 * 有了上面的解释就好说了
 * 只有await那一行是同步
 * 加了async的整个函数依然是异步的，只是new了一个Promise对象
 * 只有在调用方调用task的时候前面加了await 才会等待Promise里面的结果
 */

/**
 * 
 * 同步编程模型 同步就是多个任务一个一个执行 必须第一个运行完了才能运行第二个
 * 
 * 
 * 
 * 
 *  异步编程模型 异步就是多个任务都在执行
 * 
 * 我们看一下JS是如何实现的 比如下面这段代码
 * new Promise(请求urla)
 * new Promise(请求urlb)
 * new Promise(请求urlc)
 * 上面这段代码是一段伪代码
 * 代表三个任务 请求urla urlb urlc
 * 虽然三个url请求是一个一个请求的 (你可以想象一下相当于顺序开启了三个线程 虽然不是很准确)
 *  但是url请求的实际工作是各自分开进行的 互不影响 
 * （实际上底层应该只有一个CPU Core 
 * 三个任务的代码确实是都在执行了
 * 但是是基于划分CPU时间片原理的
 * 就是CPU一会运行任务A 一会运行任务B 一会运行任务C 确实是同时执行了
 * 相当于我今天一会学习java 一会学习go 一会学习rust
 * 三个学习任务他就是异步运行的
 * 但是需要注意的是 在任意时刻 核心只能执行其中的一个任务)
 * b请求和c请求 并不需要等待a的返回
 * 需要说明的是，JS在浏览器是单线程，他的异步编程是基于事件循环的。
 * 所以nodejs相当于用一个线程就实现了异步编程模型
 * 这样的好处是没有CPU上下文切换的开销
 * 也不需要处理共享资源竞争的问题(这是个很复杂的问题而且很棘手!!)
 * 
 * 我们再来看一下java是如何实现异步编程模型的
 * 很明显就是多线程
 * new Thread(请求urla)
 * new Thread(请求urlb)
 * new Thread(请求urlc)
 * 这是一种完全截然不同的方式 涉及到共享资源竞争的问题
 * 比如三个线程在执行成之后需要对count加一 
 * 你就必须考虑对count的操作加锁 
 * 也就是在异步里面设置同步 对于共享资源的访问 异步任务需要同步访问 也就是按照顺序一个一个访问
 * 
 * 所以说了这么多 可以总结一下异步和同步在编程语境下的含义了
 * 异步就是多个任务都在执行
 * 同步就是多个任务一个一个执行 必须第一个运行完了才能运行第二个
 * 但是JS中你不需要考虑这种问题！！
 */