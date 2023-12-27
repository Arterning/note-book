
发现一个特点，许多的npm包的核心功能用JS写
只在types目录中负责声明所有的类型，提供类型支持

比如这个recharts库， 核心功能在lib目录中
types目录定义所有类型

我们再看下package.json 定义

```js
{
  "name": "recharts",
  "version": "2.9.0",
  "description": "React charts",
  "main": "lib/index",
  "module": "es6/index",
  "jsnext:main": "es6/index",
  "types": "types/index.d.ts",
  ...
```


![[1703642856811.png]]