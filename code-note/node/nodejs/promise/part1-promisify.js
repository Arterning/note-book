const util = require('util')

const fs = require('fs')

let mineReadFile = util.promisify(fs.readFile)


mineReadFile('../file/content.txt').then(value => {
    //console.log(value); buffer
    console.log(value.toString());
})

