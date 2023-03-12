//返回的p是一个失败的promise对象
//let p = Promise.reject(404);

let p = Promise.reject(new Promise((resolve, reject) => {
    //调用resolve 表示直接返回一个成功的promise对象
    resolve("OK")
}));

console.log(p);


p.catch((reason) => {
    console.log(reason);
})