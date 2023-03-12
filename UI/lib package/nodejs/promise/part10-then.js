/**
 * then的返回结果是promise
 * 那么then返回的Promise的结果如何决定？
 * 答案是由回调函数的执行结果决定的
 */

const p = Promise.resolve('ok')


const pp = p.then((value) => {
    console.log(value);
    //throw 'error'   

    //return 123

    return Promise.resolve('pp result')
}, reason => {
    console.log(reason);
})


pp.then((value) => {
    console.log(value);
}, (reason) => {
    console.warn(reason);
})

console.log(pp);