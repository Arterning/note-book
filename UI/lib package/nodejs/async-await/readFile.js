const fs = require('fs')
const util = require('util')

const mineReadFile = util.promisify(fs.readFile);

/**
 * 回调地狱的写法
 */
fs.readFile('../file/1.html',(err, data1) => {
    if(err) throw err;
    fs.readFile('../file/2.html',(err, data2) => {
        if(err) throw err;
        fs.readFile('../file/3.html', (err, data3) => {
            if(err) throw err;
            console.log(data1 + data2 + data3);
        })
    })
})

function readFileByPromise() {

    try {
        let data1;
        let data2;
        let data3;
        mineReadFile('../file/1.html')
        .then(value1 => {
            data1 = value1;
            return mineReadFile('../file/2.html')//return new promise
        })
        .then(value2 => {
            data2 = value2;
            return mineReadFile('../file/3.html')//return new promise
        })
        .then(value3 => {
            data3 = value3;
        });
        //const data4 = await mineReadFile('../file/4.html'); 
        console.log(data1 + ',' + data2 + ',' + data3);
    } catch (error) {
        console.log(error);
    }
   
}

/**
 * async await的写法
 */

async function readFile() {

    try {
        const data1 = await mineReadFile('../file/1.html');
        const data2 = await mineReadFile('../file/2.html');
        const data3 = await mineReadFile('../file/3.html'); 
        //const data4 = await mineReadFile('../file/4.html'); 
        console.log(data1 + ',' + data2 + ',' + data3);
    } catch (error) {
        console.log(error);
    }
   
}

readFile();