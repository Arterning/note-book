/**
 * race返回一个新的promise 第一个完成的promise的结果就是最终的结果
 * p1先完成 所以最终的结果是成功的
 */
const p1 = new Promise((resolve, reject) => {
    setTimeout(() => {
        resolve('OKK');
    }, 500);
});
const p2 = new Promise((resolve, reject) => {
    setTimeout(() => {
        reject('Fuck Faild');
    }, 600);
})

const result = Promise.race([p1, p2]);

result.then((value) => {
    console.log(value);
},(reason) => {
    console.log(reason);
})