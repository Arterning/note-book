async function ning () {
    //1. await右侧是其他类型
    let res = await 10;
    console.log(res);


    //2. await右侧是promise
    let p = new Promise((resolve, reject) => {
        resolve('OK')
    })
    const pp = await p;
    console.log(pp);//return OK


    //3. try catch
    let er = new Promise((resolve, reject) => {
        reject('Error')
    })
    try {
        const es = await er;
    } catch (error) {
        console.log('oh error ',error);
    }


}


ning();
