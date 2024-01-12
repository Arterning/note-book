
```
pnpm init
pnpm install express
```

在开始用 Express 改写上面的服务器之前，我们先介绍一下上面提到的**两大封装与改进**。

### 更强大的 Request 和 Response 对象

首先是 Request 请求对象，通常我们习惯用 `req` 变量来表示。下面列举一些 `req` 上比较重要的成员（如果不知道是什么也没关系哦）：

- `req.body`：客户端请求体的数据，可能是表单或 JSON 数据
- `req.params`：请求 URI 中的路径参数
- `req.query`：请求 URI 中的查询参数
- `req.cookies`：客户端的 cookies

[`req.params`](http://expressjs.com/en/api.html#req.params) contains route parameters (in the path portion of the URL), and [`req.query`](http://expressjs.com/en/api.html#req.query) contains the URL query parameters (after the `?` in the URL).

  
然后是 Response 响应对象，通常用 `res` 变量来表示，可以执行一系列响应操作，例如：

```js
// 发送一串 HTML 代码
res.send('HTML String');

// 发送一个文件
res.sendFile('file.zip');

// 渲染一个模板引擎并发送
res.render('index');
```

chain invoke response

```js
// 设置状态码为 404，并返回 Page Not Found 字符串
res.status(404).send('Page Not Found');
```