const fs = require('fs')

let p = new Promise((resolve, reject) => {

    fs.readFile('../file/content.txt', (err, data) => {
        if (err) reject(err);
        else resolve(data);
    })
});

p.then(value => {
    console.log(value.toString());
}, reason => {
    console.log(reason);
})