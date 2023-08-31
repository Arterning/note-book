# ch04-require和import的区别

require/exports	2009	CommonJS
import/export	2015	ECMAScript2015（ES6）


CommonJS 模块化方案 require/exports 是为服务器端开发设计的。服务器模块系统同步读取模块文件内容，编译执行后得到模块接口。（Node.js 是 CommonJS 规范的实现）。
在浏览器端，因为其异步加载脚本文件的特性，CommonJS 规范无法正常加载。所以出现了 RequireJS、SeaJS 等（兼容 CommonJS ）为浏览器设计的模块化方案。直到 ES6 规范 出现，浏览器才拥有了自己的模块化方案 import/export。



# import/export写法
```javascript
import fs from 'fs'
import {readFile} from 'fs' //从 fs 导入 readFile 模块
import {default as fs} from 'fs' //从 fs 中导入使用 export default 导出的模块
import * as fileSystem from 'fs' //从 fs 导入所有模块，引用对象名为 fileSystem
import {readFile as read} from 'fs' //从 fs 导入 readFile 模块，引用对象名为 read

export default fs
export const fs
export function readFile
export {readFile, read}
export * from 'fs'
```

1. 引入 export default 导出的模块不用加 {},引入非 export default 导出的模块需要加 {}
2. 一个文件只能导出一个 default 模块。

# 总结
require是CommonJS的规范 是为服务器开发设计的
Import是ES6的规范


# refer
https://zhuanlan.zhihu.com/p/121770261