/**
size 属性
set()
get()
has()
delete()
clear()
 * 增删改查
 */
function base() {
    const m = new Map()

    m.set('edition', 6)        // 键是字符串
    m.set(262, 'standard')     // 键是数值
    m.set(undefined, 'nah')    // 键是 undefined
    m.set(1, 'a').set(2, 'b').set(3, 'c') // 链式操作

    const hello = function () { console.log('hello'); };
    m.set(hello, 'Hello ES6!') // 键是函数
    m.get(hello)  // Hello ES6!


    m.has('edition')     // true


    m.set(undefined, 'nah');
    m.has(undefined)     // true
    m.delete(undefined)
    m.has(undefined)       // false

    map.clear()

}


/**
 * 
 * map 遍历
 */
function iterator() {
    const map = new Map([
        ['F', 'no'],
        ['T',  'yes'],
      ]);
      
    for (let key of map.keys()) {
    console.log(key);
    }
    // "F"
    // "T"
    
    for (let value of map.values()) {
    console.log(value);
    }
    // "no"
    // "yes"
    
    for (let item of map.entries()) {
    console.log(item[0], item[1]);
    }
    // "F" "no"
    // "T" "yes"
    
    // 或者
    for (let [key, value] of map.entries()) {
    console.log(key, value);
    }
    // "F" "no"
    // "T" "yes"
    
    // 等同于使用map.entries()
    for (let [key, value] of map) {
    console.log(key, value);
    }
    // "F" "no"
    // "T" "yes"
    
    map.forEach(function(value, key, map) {
    console.log("Key: %s, Value: %s", key, value);
    });
}
