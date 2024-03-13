
module和module.exports

module.exports是本体

module是指向module.exports的引用


所以我们可以这样修改

```js
module.a = "sdf"
module.b = "xx"
```

但是不能这样修改
```js
module = {
	"a": "sdkf",
	"b":'sdfk"
}
```

因为这就相当于module重新指向了另外一个对象，这样module.exports实际上是没有变化的
这是js的语言特性决定的

所以我们推荐使用module.exports导出
```js
module.exports = {
	"a": "sdkf",
	"b":'sdfk"
}
```


https://www.liaoxuefeng.com/wiki/1022910821149312/1023027697415616