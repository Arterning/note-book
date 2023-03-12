let p = new Promise((resolve, reject) => {


    reject('error')
});


//catch内部的实现也是由then来实现的
//catch传入reject方法
//then同时传入resolve和reject方法
p.catch(reason => {
    console.log(reason);
})