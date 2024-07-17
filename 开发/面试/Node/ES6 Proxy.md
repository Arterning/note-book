
```
var person = {
  name: "张三"
};

var proxy = new Proxy(person, {
  get: function(target, propKey) {
    return Reflect.get(target,propKey)
  }
});

proxy.name // "张三"
```


```
关于handler拦截属性，有如下：

get(target,propKey,receiver)：拦截对象属性的读取
set(target,propKey,value,receiver)：拦截对象属性的设置
has(target,propKey)：拦截propKey in proxy的操作，返回一个布尔值
deleteProperty(target,propKey)：拦截delete proxy[propKey]的操作，返回一个布尔值
ownKeys(target)：拦截Object.keys(proxy)、for...in等循环，返回一个数组
getOwnPropertyDescriptor(target, propKey)：拦截Object.getOwnPropertyDescriptor(proxy, propKey)，返回属性的描述对象
defineProperty(target, propKey, propDesc)：拦截Object.defineProperty(proxy, propKey, propDesc），返回一个布尔值
preventExtensions(target)：拦截Object.preventExtensions(proxy)，返回一个布尔值
getPrototypeOf(target)：拦截Object.getPrototypeOf(proxy)，返回一个对象
isExtensible(target)：拦截Object.isExtensible(proxy)，返回一个布尔值
setPrototypeOf(target, proto)：拦截Object.setPrototypeOf(proxy, proto)，返回一个布尔值
apply(target, object, args)：拦截 Proxy 实例作为函数调用的操作
construct(target, args)：拦截 Proxy 实例作为构造函数调用的操作
```