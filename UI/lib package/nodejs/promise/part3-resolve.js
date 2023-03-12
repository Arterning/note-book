
//返回的p是一个成功的promise对象
let p = Promise.resolve(520);

console.log(p);



//如果传入的对象是一个Promise对象 那么这个传入的promise结果决定了pp的结果
let pp = Promise.resolve(new Promise((resolve, reject) => {
    resolve('ok')
})) 
console.log(pp);


let fail = Promise.resolve(new Promise((resolve, reject) => {
    reject('fail')
})) 
console.log(fail);

fail.catch((reason) => {
    console.log(reason);
})