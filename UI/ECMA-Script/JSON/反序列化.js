/**
 * 拿到一个JSON格式的字符串，我们直接用JSON.parse()把它变成一个JavaScript对象：
 */
JSON.parse("[1,2,3,true]"); // [1, 2, 3, true]
JSON.parse('{"name":"小明","age":14}'); // Object {name: '小明', age: 14}
JSON.parse("true"); // true
JSON.parse("123.45"); // 123.45
