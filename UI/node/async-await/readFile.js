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