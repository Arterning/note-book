const p = Promise.resolve('OK')

p.then((value) => {
    console.log('成功了');
})


p.then((value) => {
    console.log('value is', value);
})


p.then((value) => {
    console.log('第三次回调');
})