
> `any`: 动态的变量类型（失去了类型检查的作用）。  
> `never`: 永不存在的值的类型。例如：never 类型是那些总是会抛出异常或根本就不会有返回值的函数表达式或箭头函数表达式的返回值类型。  
> `unknown`: 任何类型的值都可以赋给 unknown 类型，但是 unknown 类型的值只能赋给 unknown 本身和 any 类型。  
> `null & undefined`: 默认情况下 null 和 undefined 是所有类型的子类型。 就是说你可以把 null 和 undefined 赋值给 number 类型的变量。当你指定了 --strictNullChecks 标记，null 和 undefined 只能赋值给 void 和它们各自。  
> `void`: 没有任何类型。例如：一个函数如果没有返回值，那么返回值可以定义为void。


### tuple

元祖类型，允许表示一个已知元素数量和类型的数组，各元素的类型不必相同

```
let tupleArr:[number, string, boolean];
tupleArr = [12, '34', true]; //ok
typleArr = [12, '34'] // no ok
```


赋值的类型、位置、个数需要和定义（生明）的类型、位置、个数一致


## enum

`enum`类型是对JavaScript标准数据类型的一个补充，使用枚举类型可以为一组数值赋予友好的名字

```
enum Color {Red, Green, Blue}
let c: Color = Color.Green;
```


### null 和 和 undefined
在`JavaScript` 中 `null`表示 "什么都没有"，是一个只有一个值的特殊类型，表示一个空对象引用，而`undefined`表示一个没有设置值的变量

默认情况下`null`和`undefined`是所有类型的子类型， 就是说你可以把 `null`和 `undefined`赋值给 `number`类型的变量

但是`ts`配置了`--strictNullChecks`标记，`null`和`undefined`只能赋值给`void`和它们各自




