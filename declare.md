declare 关键字用来告诉编译器，某个类型是存在的，可以在当前文件中使用。

它的主要作用，就是让当前文件可以使用其他文件声明的类型。举例来说，自己的脚本使用外部库定义的函数，编译器会因为不知道外部函数的类型定义而报错，这时就可以在自己的脚本里面使用`declare`关键字，告诉编译器外部函数的类型。这样的话，编译单个脚本就不会因为使用了外部类型而报错。

declare 关键字可以描述以下类型。

- 变量（const、let、var 命令声明）
- type 或者 interface 命令声明的类型
- class
- enum
- 函数（function）
- 模块（module）
- 命名空间（namespace）

declare 关键字的重要特点是，它只是通知编译器某个类型是存在的，不用给出具体实现。比如，只描述函数的类型，不给出函数的实现，如果不使用`declare`，这是做不到的。

declare 只能用来描述已经存在的变量和数据结构，不能用来声明新的变量和数据结构。另外，所有 declare 语句都不会出现在编译后的文件里面。

## declare module，declare namespace[​](https://typescript.p6p.net/typescript-tutorial/declare.html#declare-module-declare-namespace)

如果想把变量、函数、类组织在一起，可以将 declare 与 module 或 namespace 一起使用。

typescript

```
declare namespace AnimalLib {
  class Animal {
    constructor(name: string);
    eat(): void;
    sleep(): void;
  }

  type Animals = "Fish" | "Dog";
}

// 或者
declare module AnimalLib {
  class Animal {
    constructor(name: string);
    eat(): void;
    sleep(): void;
  }

  type Animals = "Fish" | "Dog";
}
```

上面示例中，declare 关键字给出了 module 或 namespace 的类型描述。


[declare 关键字 | 阮一峰 TypeScript 教程 (p6p.net)](https://typescript.p6p.net/typescript-tutorial/declare.html#declare-module-declare-namespace)

