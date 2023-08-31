[TOC]

# Object.keys，values，entries

对于普通对象，下列这些方法是可用的：

- [Object.keys(obj)]
- [Object.values(obj)]
- [Object.entries(obj)]

```jsx
let user = {
	name: "John",
	age: 30
};

/**
Object.keys(user) = ["name", "age"]
Object.values(user) = ["John", 30]
Object.entries(user) = [ ["name","John"], ["age",30] ]
**/
```

例如，我们有一个带有价格的对象，并想将它们加倍：

```jsx
let prices = {
  banana: 1,
  orange: 2,
  meat: 4,
};

let doublePrices = Object.fromEntries(
  // 将价格转换为数组，将每个键/值对映射为另一对
  // 然后通过 fromEntries 再将结果转换为对象
  Object.entries(prices).map(entry => [entry[0], entry[1] * 2])
);
```


## Set
Set 是一个特殊的类型集合 —— “值的集合”（没有键），它的每一个值只能出现一次。

它的主要方法如下：

new Set(iterable) —— 创建一个 set，如果提供了一个 iterable 对象（通常是数组），将会从数组里面复制值到 set 中。
set.add(value) —— 添加一个值，返回 set 本身
set.delete(value) —— 删除值，如果 value 在这个方法调用的时候存在则返回 true ，否则返回 false。
set.has(value) —— 如果 value 在 set 中，返回 true，否则返回 false。
set.clear() —— 清空 set。
set.size —— 返回元素个数。

Set 迭代（iteration）
我们可以使用 for..of 或 forEach 来遍历 Set：

```jsx
let set = new Set(["oranges", "apples", "bananas"]);

for (let value of set) alert(value);

// 与 forEach 相同：
set.forEach((value, valueAgain, set) => {
alert(value);
});
```

注意一件有趣的事儿。forEach 的回调函数有三个参数：一个 value，然后是 同一个值 valueAgain，最后是目标对象。没错，同一个值在参数里出现了两次。

forEach 的回调函数有三个参数，是为了与 Map 兼容。当然，这看起来确实有些奇怪。但是这对在特定情况下轻松地用 Set 代替 Map 很有帮助，反之亦然。


## Map

Map 是怎么比较键的？
Map 使用 SameValueZero 算法来比较键是否相等。它和严格等于 === 差不多，但区别是 NaN 被看成是等于 NaN。所以 NaN 也可以被用作键。

这个算法不能被改变或者自定义。

如果要在 map 里使用循环，可以使用以下三个方法：

map.keys() —— 遍历并返回一个包含所有键的可迭代对象，
map.values() —— 遍历并返回一个包含所有值的可迭代对象，
map.entries() —— 遍历并返回一个包含所有实体 [key, value] 的可迭代对象，for..of 在默认情况下使用的就是这个。

```jsx
let recipeMap = new Map([
  ['cucumber', 500],
  ['tomatoes', 350],
  ['onion',    50]
]);

// 遍历所有的键（vegetables）
for (let vegetable of recipeMap.keys()) {
  alert(vegetable); // cucumber, tomatoes, onion
}

// 遍历所有的值（amounts）
for (let amount of recipeMap.values()) {
  alert(amount); // 500, 350, 50
}

// 遍历所有的实体 [key, value]
for (let entry of recipeMap) { // 与 recipeMap.entries() 相同
  alert(entry); // cucumber,500 (and so on)
}
```

除此之外，`Map` 有内建的 `forEach` 方法，与 `Array` 类似：

```jsx
// 对每个键值对 (key, value) 运行 forEach 函数
recipeMap.forEach( (value, key, map) => {
  alert(`${key}: ${value}`); // cucumber: 500 etc
});
```

每一次 map.set 调用都会返回 map 本身，所以我们可以进行“链式”调用：

```jsx

map.set('1', 'str1')
  .set(1, 'num1')
  .set(true, 'bool1');

```

map初始化

```jsx
// 键值对 [key, value] 数组
let map = new Map([
  ['1',  'str1'],
  [1,    'num1'],
  [true, 'bool1']
]);
```

## 从对象创建map

如果我们想从一个已有的普通对象（plain object）来创建一个 `Map`，那么我们可以使用内建方法 [Object.entries(obj)](https://developer.mozilla.org/zh/docs/Web/JavaScript/Reference/Global_Objects/Object/entries)，该方法返回对象的键/值对数组，该数组格式完全按照 `Map` 所需的格式。

```jsx
let obj = {
  name: "John",
  age: 30
};

let map = new Map(Object.entries(obj));

alert( map.get('name') ); // John
```

## 从map创建对象

```jsx
let prices = Object.fromEntries([
  ['banana', 1],
  ['orange', 2],
  ['meat', 4]
]);

// 现在 prices = { banana: 1, orange: 2, meat: 4 }
alert(prices.orange); // 2
```


