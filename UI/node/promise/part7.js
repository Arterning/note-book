/**
 * 如何改变promise状态
 */

const p = new Promise((resolve, reject) => {

    //1. resovle pending=>fulfilled

    //2. reject pending=>rejected

    //3. throw error
    throw '丢雷楼某 失败了';
})

//这时候p是pending状态 因为resolve 和 reject都没有被调用
//如果调用了resolve 那么p就是成功状态 如果调用的是reject 那么p就是失败状态
console.log(p);

p.catch((reason) => {
    console.log(reason);
})