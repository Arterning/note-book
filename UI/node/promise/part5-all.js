

const result = Promise.all([
    Promise.resolve('OK'),
    Promise.resolve('OOK'),
    Promise.reject('Fail')]);

console.log(result);

result.catch((reason) => {
    console.log(reason);
})


const ok = Promise.all([
    Promise.resolve('OK'),
    Promise.resolve('OOK'),
    Promise.resolve('OK')]);
ok.then(value=> {
    console.log(value);
}, reason => {
    console.log(reason);
})