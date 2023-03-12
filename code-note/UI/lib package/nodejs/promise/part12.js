/**
 * promise异常穿透
 * 当使用promise链式调用的时候 可以在最后指定失败的回调 
 * 前面任何操作出现异常 都可以传到最后失败的回调中处理
 */


/**
 * promise链的中断
 * 
 * 只有一种方式 就是返回一个pending状态的promise
 * return new Promise(() => {})
 */

const p = new Promise((resolve, reject) => {
    setTimeout(() => {
        //reject('Failed');
        resolve('OK')
    }, 2000);
})

const result = p.then(value => {
    console.log(111);
}).then(value => {
    console.log(222);
}).then(value => {
    console.log(333);
    return new Promise(() => {});
}).then(value => {
    console.log(444);
}).catch(reason => { 
    console.log(reason);
});

