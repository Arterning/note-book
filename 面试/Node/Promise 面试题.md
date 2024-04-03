

总结：

- promise.all: 只有参数里面的promise全部成功，才会触发成功（返回结果按照参数顺序），一旦有一个异常，就返回失败
- promise.allSettled: 等待参数里面的promise执行结束，一起返回，不管成功还是失败
- promise.any:谁先成功就先返回，我觉得改名成promise.anyFulfilled 更好。
- promise.race:谁先结束就先返回

  
```js
const delay = n => new Promise(resolve => setTimeout(resolve, n));

const promises = [
  delay(100).then(() => 1),
  delay(200).then(() => 2),
  ]

Promise.all(promises).then(values=>console.log(values))
// 最终输出： [1, 2]
```


可是，是一旦有一个promise出现了异常，被reject了，情况就会变的麻烦。

```js
const promises = [
  delay(100).then(() => 1),
  delay(200).then(() => 2),
  Promise.reject(3)
  ]

Promise.all(promises).then(values=>console.log(values))
// 最终输出： Uncaught (in promise) 3

Promise.all(promises)
.then(values=>console.log(values))
.catch(err=>console.log(err))
// 加入catch语句后，最终输出：3
```


尽管能用catch捕获其中的异常，但你会发现其他执行成功的Promise的消息都丢失了，仿佛石沉大海一般。

要么全部成功，要么全部重来，这是Promise.all本身的强硬逻辑，也是痛点的来源，不能说它错，但这的确给Promise.allSettled留下了立足的空间。

假如使用Promise.allSettled来处理这段逻辑会怎样呢?



```js
const promises = [
  delay(100).then(() => 1),
  delay(200).then(() => 2),
  Promise.reject(3)
  ]

Promise.allSettled(promises).then(values=>console.log(values))
// 最终输出： 
//    [
//      {status: "fulfilled", value: 1},
//      {status: "fulfilled", value: 2},
//      {status: "rejected", value: 3},
//    ]
```

可以看到所有promise的数据都被包含在then语句中，且每个promise的返回值多了一个status字段，表示当前promise的状态，没有任何一个promise的信息被丢失。

因此，当用Promise.allSettled时，我们只需专注在then语句里，当有promise被异常打断时，我们依然能妥善处理那些已经成功了的promise，不必全部重来。
