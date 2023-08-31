## **都可以描述一个对象或者函数**

```tsx
interface User {
  name: string
  age: number
}

interface SetUser {
  (name: string, age: number): void;
}
```

```tsx
type User = {
  name: string
  age: number
};

type SetUser = (name: string, age: number)=> void;
```

### 都允许拓展（extends）

interface 和 type 都可以拓展，并且两者并不是相互独立的，也就是说 interface 可以 extends type, type 也可以 extends interface 。 **虽然效果差不多，但是两者语法不同**。

```tsx
interface Name { 
  name: string; 
}
interface User extends Name { 
  age: number; 
}
```

```tsx
type Name = { 
  name: string; 
}
type User = Name & { age: number  };
```

```tsx
type Name = { 
  name: string; 
}
interface User extends Name { 
  age: number; 
}
```

```tsx
interface Name { 
  name: string; 
}
type User = Name & { 
  age: number; 
}
```

## **type 可以而 interface 不行**

type 可以声明基本类型别名，联合类型，元组等类型

```tsx
// 基本类型别名
type Name = string

// 联合类型
interface Dog {
    wong();
}
interface Cat {
    miao();
}

type Pet = Dog | Cat

// 具体定义数组每个位置的类型
type PetList = [Dog, Pet]

```

```tsx
// 当你想获取一个变量的类型时，使用 typeof
let div = document.createElement('div');
type B = typeof div
```

```tsx
type StringOrNumber = string | number;  
type Text = string | { text: string };  
type NameLookup = Dictionary<string, Person>;  
type Callback<T> = (data: T) => void;  
type Pair<T> = [T, T];  
type Coordinates = Pair<number>;  
type Tree<T> = T | { left: Tree<T>, right: Tree<T> };

```

### interface 可以而 type 不行

interface 能够声明合并

```tsx
interface User {
  name: string
  age: number
}

interface User {
  sex: string
}

/*
User 接口为 {
  name: string
  age: number
  sex: string 
}
*/

```

## 总结

用interface描述**数据结构**，用type描述**类型关系**

`interface Point { x: number, y: number }`Point包含 x、y 两个数字属性，就是结构。

`type Point3D =Point & { z: number }` Point3D 是 Point 与{ z: number } 组成（并集关系）

精辟，但这个场景，`interface Point3D extends Point`似乎更好？

## 参考

- [link](https://juejin.cn/post/6844903749501059085)