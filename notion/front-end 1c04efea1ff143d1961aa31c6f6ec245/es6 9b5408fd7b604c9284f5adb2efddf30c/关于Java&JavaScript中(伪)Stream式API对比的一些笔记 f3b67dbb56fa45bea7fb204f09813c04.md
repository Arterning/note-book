# 关于Java&JavaScript中(伪)Stream式API对比的一些笔记

```jsx
// 顺序进行
List<Apple> listStream = list.stream()
        .filter((Apple a) -> a.getWeight() >20 || "green".equals(a.getColor()))
        .collect(Collectors.toList());
//并行进行
List<Apple> listStreamc = list.parallelStream()
        .filter((Apple a) -> a.getWeight() >20 || "green".equals(a.getColor()))
        .collect(Collectors.toList());
```

```jsx
let colors = ["red", "blue", "grey"];

colors.forEach((item, index, arr) ==> {
    if(item === "red") {
        arr.splice(index, 1);
    }
});
```

## Filter

```jsx
List<Integer> list  = Arrays.asList(12, 3, 4, 5, 4);
list.stream().filter( i -> i % 2 == 0)
             .forEach(System.out::print);// 12 4 4
```

```jsx
let users = [{ name: "毋意", value: "202201" }, { name: "毋必", value: "202202" }, 
             { name: "毋固", value: "202203" },{ name: "毋我", value: "202204" }]

users.filter(o => o.value === 202201 ).forEach(o =>console.log('out :%s',o))
//out :{ name: '毋意', value: '202201' }
```

## map

```jsx
List<Integer> list  = Arrays.asList(12, 3, 4, 5, 4);
list.stream().map(o -> o+1 ).forEach(System.out::println);
```

```jsx
let users = [{ name: "毋意", value: "202201" }, { name: "毋必", value: "202202" }, 
             { name: "毋固", value: "202203" },{ name: "毋我", value: "202204" }]             
users.map( o => o.name ).forEach(o =>console.log('out :%s',o))
```

## ****flatMap 扁平化****

```jsx
List<String> strings = Arrays.asList("Hello","World");
strings.stream().map(o -> o.split(""))
        .flatMap(Arrays::stream)
        .forEach(System.out::println);
```

```jsx
let string = ["Hello","World"]
string.flatMap( o => o.split("")).forEach(o =>console.log('out :%s',o))
```

`JS`提供了`flat`方法可以默认展开数组,flat() 方法会按照一个`可指定的深度`递归遍历数组，并将所有元素与遍历到的子数组中的元素合并为一个新数组返回。

```jsx
[1, 2, [3, [4, 5]]].flat()
// [1, 2, 3, [4, 5]]

[1, 2, [3, [4, 5]]].flat(2)
// [1, 2, 3, 4, 5]
```

## 截断

通过`截断流`我们可以看到`Java的JavaScript在Stream上本质的不同`，Java通过Stream 对象本身`OP_MASK`属性来截断,而JS没有实际意义上的Stream对象， 但是可以通过`filter结合index`来完成，或者使用`slice`

```jsx
List<Integer> list  = Arrays.asList(12, 3, 4, 5, 4);
list.stream().limit(2).forEach(System.out::println);
```

```jsx
let users = [{ name: "毋意", value: "202201" }, { name: "毋必", value: "202202" }, 
             { name: "毋固", value: "202203" },{ name: "毋我", value: "202204" }]   
users.slice(2).forEach(o =>console.log('out :%s',o))

//user.slice(2) 表示只保留前2个元素
======================================
out :{ name: '毋意', value: '202201' }
out :{ name: '毋必', value: '202202' }

users.filter((_, i) => i < 2).forEach(o => console.log('out :%s', o))
============
out :{ name: '毋意', value: '202201' }
out :{ name: '毋必', value: '202202' }
```

## 排序

```jsx
List<Integer> list  = Arrays.asList(12, 3, 4, 5, 4);
list.stream()
        .sorted( (o1,o2) -> o1 > o2 ? 1 : (o1 < o2 ? -1 : 0 ))
        .forEach(System.out::println);
```

```jsx
let users = [{ name: "毋意", value: "202201" }, { name: "毋必", value: "202202" }, 
             { name: "毋固", value: "202203" },{ name: "毋我", value: "202204" }]  
users.map(o => { return { name: o.name, value: +o.value } })
     .sort((o1, o2) => o1.value > o2.value ? -1 : (o1.value < o2.value ? 1 : 0))
     .forEach(o => console.log(o))
==================================
{ name: '毋我', value: 202204 }
{ name: '毋固', value: 202203 }
{ name: '毋必', value: 202202 }
{ name: '毋意', value: 202201 }
```

## 去重

```jsx
List<Integer> list  = Arrays.asList(12, 3, 4, 5, 4);
list.stream().distinct().forEach(System.out::println);
```

```jsx
let numbers = [2,3,4,3,5,2]
Array.from(new Set(numbers)).forEach(o => console.log(o))
```

## 跳过

`跳过元素`：返回一个扔掉了前n个元素的流。如果流中元素不足n个，则返回一个空流。

```jsx
List<Integer> list  = Arrays.asList(12, 3, 4, 5, 4);
list.stream().skip(2).forEach(System.out::println);
==================
4
5
4
```

```jsx
let users = [{ name: "毋意", value: "202201" }, { name: "毋必", value: "202202" }, 
             { name: "毋固", value: "202203" },{ name: "毋我", value: "202204" }] 
users.slice(0, 2).forEach(o => console.log(o))             
=========
{ name: '毋固', value: '202203' }
{ name: '毋我', value: '202204' }
```

## group

分组操作的结果是一个`Map`，把`分组函数返回的值作为映射的键`，把流中所有具有这个分类值的项目的列表作为对应的`映射值`

```jsx
List<String> lists  = Arrays.asList("123", "1234", "4564", "789");
lists.stream().collect(Collectors.groupingBy( o -> o.length()))
              .forEach((o1,o2) -> System.out.printf("%s:%s\n",o1,o2));
=========
3:[123, 789]
4:[1234, 4564]
```

```jsx
const array = [1, 2, 3, 4, 5];

array.group((num, index, array) => {
  return num % 2 === 0 ? 'even': 'odd';
});
// { odd: [1, 3, 5], even: [2, 4] }
```

groupToMap()的作用和用法与group()完全一致，唯一的区别是返回值是一个`Map 结构`，而不是`对象`

```jsx
const array = [1, 2, 3, 4, 5];

const odd  = { odd: true };
const even = { even: true };
array.groupToMap((num, index, array) => {
  return num % 2 === 0 ? even: odd;
});
//  Map { {odd: true}: [1, 3, 5], {even: true}: [2, 4] }
```

## 查找

```jsx
List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6);
System.out.println(numbers.stream().findAny().get()); //1
System.out.println(numbers.stream().findFirst().get()); //1
```

```jsx
let users = [{ name: "毋意", value: "202201" }, { name: "毋必", value: "202202" },
{ name: "毋固", value: "202203" }, { name: "毋我", value: "202204" }]

let user = users.find(o => o.name === "毋固")
console.log(user) //{ name: '毋固', value: '202203' }
let useri = users.findIndex(o => o.name === "毋固")
console.log(useri) //2
```

- ****[关于Java&JavaScript中(伪)Stream式API对比的一些笔记](https://bbs.huaweicloud.com/blogs/364943)****