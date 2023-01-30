const s = new Set();

//当添加实例中已经存在的元素，set不会进行处理添加

s.add(1).add(2).add(2); // 2只被添加了一次
s.delete(1)
s.has(2)//true
s.clear()

/**
 * 
keys()：返回键名的遍历器
values()：返回键值的遍历器
entries()：返回键值对的遍历器
forEach()：使用回调函数遍历每个成员
 */

let set = new Set(['red', 'green', 'blue']);

for (let item of set.keys()) {
  console.log(item);
}
// red
// green
// blue

for (let item of set.values()) {
  console.log(item);
}
// red
// green
// blue

for (let item of set.entries()) {
  console.log(item);
}
// ["red", "red"]
// ["green", "green"]
// ["blue", "blue"]


let ss = new Set([1, 4, 9]);
ss.forEach((value, key) => console.log(key + ' : ' + value))
// 1 : 1
// 4 : 4
// 9 : 9

/**
 * 
 * 扩展运算符和Set 结构相结合实现数组或字符串去重
 */
 //数组
 let arr = [3, 5, 2, 2, 5, 5];
 let unique_arr = [...new Set(arr)]; // [3, 5, 2]
 
 // 字符串
 let str = "352255";
 let unique_str = [...new Set(str)].join(""); // "352"


/**
 * 
 * 实现并集、交集、和差集
 */
let a = new Set([1, 2, 3]);
let b = new Set([4, 3, 2]);

// 并集
let union = new Set([...a, ...b]);
// Set {1, 2, 3, 4}

// 交集
let intersect = new Set([...a].filter(x => b.has(x)));
// set {2, 3}

// （a 相对于 b 的）差集
let difference = new Set([...a].filter(x => !b.has(x)));