

```js
interface Animal {
  name: string;
  speak(): void;
}

interface Dog extends Animal {
  breed: string;
}

class Labrador implements Dog {
  name: string;
  breed: string;

  constructor(name: string, breed: string) {
    this.name = name;
    this.breed = breed;
  }

  speak() {
    console.log(`${this.name} says woof!`);
  }
}

```


`class` 是一种定义类型和实现的方式。它既包含类型信息，也包含实际的属性和方法实现。与 `type` 和 `interface` 不同，`class` 定义的类型信息会保留在编译后的代码中，因为它们在运行时是必需的。

  

`class` 可以通过关键字 `extends` 实现类继承，还可以通过关键字 `implements` 实现接口实现。这使得 `class` 成为创建具有多层次结构和行为的对象的理想选择。






总结：


在 TypeScript 中，`type`、`interface` 和 `class` 分别具有自己的用途和特点。


- `type` 适用于定义类型别名、联合类型、交叉类型等，并且不需要运行时信息。
    
- `interface` 主要用于定义对象的类型和形状，支持继承和实现。
    
- `class` 既包含类型信息，也包含实际的属性和方法实现。在实际开发中，我们应根据需求选择合适的类型声明方式。
    
  

虽然 `type` 和 `interface` 在很多场景下可以互换使用，但它们在某些特定场景下有着各自的优势。`type` 更适用于组合不同类型，如联合类型、交叉类型等，而 `interface` 更适用于定义对象的形状，特别是在面向对象编程中。`class` 则提供了完整的类型定义和实现，可以在运行时进行实例化和操作。

  

在实践中，我们应该根据实际需求和场景选择合适的类型声明方式。例如，在定义一个复杂的对象类型时，可以使用 `interface`；在组合不同类型时，可以使用 `type`；在创建具有行为的对象时，可以使用 `class`。