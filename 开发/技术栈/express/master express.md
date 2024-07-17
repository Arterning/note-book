
- Express 框架的两大核心概念：路由和中间件
- 用 Nodemon 加速开发迭代
- 使用模板引擎渲染页面，并接入 Express 框架中
- 使用 Express 的静态文件服务
- 编写自定义的错误处理函数
- 实现一个简单的 JSON API 端口
- 通过子路由拆分逻辑，实现模块化

  
The raw ways to create http server

```js
const http = require('http');

const hostname = 'localhost';
const port = 3000;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/html');
  res.end('Hello World\n');
});

server.listen(port, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});

```


可以发现，直接用内置的 http 模块去开发服务器有以下明显的弊端：

需要写很多底层代码——例如手动指定 HTTP 状态码和头部字段，最终返回内容。如果我们需要开发更复杂的功能，涉及到多种状态码和头部信息（例如用户鉴权），这样的手动管理模式非常不方便
没有专门的路由机制——路由是服务器最重要的功能之一，通过路由才能根据客户端的不同请求 URL 及 HTTP 方法来返回相应内容。但是上面这段代码只能在 http.createServer 的回调函数中通过判断请求 req 的内容才能实现路由功能，搭建大型应用时力不从心

由此就引出了 Express 对内置 http 的两大封装和改进：

更强大的请求（Request）和响应（Response）对象，添加了很多实用方法
灵活方便的路由的定义与解析，能够很方便地进行代码拆分

接下来，我们将开始用 Express 来开发 Web 服务器！



https://juejin.cn/post/6844904023380721678#heading-22


