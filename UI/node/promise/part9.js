/**
 * 改变promise状态和指定回调函数 谁先谁后
 * 
 * 都有可能
 * 
 * 正常情况是先指定回调 然后 改变状态
 * const p = new Promise(里面一个异步任务 可能需要等待2ms)
 * p.then(..)
 * 这样就是then先执行
 * 
 * 
 * 也可以这样先改变状态 再指定回调
 * 先改变状态 再调用then
 * 
 * 
 */

//先改变状态 再调用then
const p = new Promise((resolve, reject) => {
    console.log('直接执行同步任务 先改变状态');
    resolve('OK');
})

//什么时候拿到数据就是什么时候回调函数执行
p.then((value) => {
    console.log(value);
})


