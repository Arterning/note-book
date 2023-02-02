
/**
 * 相当于return Promise.resolve(123)
 * @returns 
 */
async function main() {

    return 123;
}

let result = main()
console.log(result);


async function mmin() {
    return new Promise((resolve, reject) => {
        resolve('OK')
    })
}

const rr = mmin();
setTimeout(() => {
    console.log(rr);    
}, 3000);


async function masc() {
    throw 'on fuck you'
}
const rc = masc();
console.log(rc);