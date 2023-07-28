# typescript 装饰器

乍一看非常类似java注解

装饰器类型及其执行优先级为

1. 类装饰器——优先级 4 （对象实例化，静态）
2. 方法装饰器——优先级 2 （对象实例化，静态）
3. 访问器或属性装饰器——优先级 3 （对象实例化，静态）
4. 参数装饰器——优先级 1 （对象实例化，静态）

方法装饰器函数有三个参数：

1. **target** —— 当前对象的原型，也就是说，假设 Employee 是对象，那么 target 就是 `Employee.prototype`
2. **propertyKey** —— 方法的名称
3. **descriptor** —— 方法的属性描述符，即 `Object.getOwnPropertyDescriptor(Employee.prototype, propertyKey)`

```jsx
Object.getOwnPropertyDescriptor([], "length")
Object { value: 0, writable: true, enumerable: false, configurable: false }

configurable: false

enumerable: false

value: 0

writable: true
```