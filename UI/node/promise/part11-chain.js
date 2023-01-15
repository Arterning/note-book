/**
 * proimse串联多个任务 这样就没有回调地狱的问题了
 */

let p = new Promise((resolve, reject) => {
    setTimeout(() => {
        console.log('api invoke ...');
        resolve('200 ok');
    }, 500);
});

p.then(value => {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log('api invoke2');
        }, 100);
        resolve('200 ok');
    })
}).then(value => {
    console.log(value);
}).then(value => {
    console.log(value);
})