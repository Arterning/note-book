中间件并不是 Express 独有的概念。相反，它是一种广为使用的软件工程概念（甚至已经延伸到了其他行业），是指**将具体的业务逻辑和底层逻辑解耦的组件**（可查看这个[讨论](https://link.juejin.cn/?target=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F19730582 "https://www.zhihu.com/question/19730582")）。换句话说，中间件就是能够适用多个应用场景、可复用性良好的代码。

![[Pasted image 20240110152840.png]]



在 Express 中，中间件就是一个函数：

```js
function someMiddleware(req, res, next) {
  // 自定义逻辑
  next();
}
```

> **注意**
> 
> 如果忘记在中间件中调用 `next` 函数，并且又不直接返回响应时，服务器会直接卡在这个中间件不会继续执行下去哦！

在 Express 使用中间件有两种方式：**全局中间件**和**路由中间件**。


#### 全局中间件

通过 `app.use` 函数就可以注册中间件，并且此中间件会在用户发起**任何请求**都可能会执行，例如：


`app.use(someMiddleware);`


#### 路由中间件

通过在路由定义时注册中间件，此中间件只会在用户访问该路由对应的 URI 时执行，例如：


`app.get('/middleware', someMiddleware, (req, res) => {   res.send('Hello World'); });`

  
  接下来我们就开始实现第一个 Express 中间件。功能很简单，就是在终端打印客户端的访问时间、 HTTP 请求方法和 URI，名为 `loggingMiddleware`


```js
const app = express();

function loggingMiddleware(req, res, next) {
  const time = new Date();
  console.log(`[${time.toLocaleString()}] ${req.method} ${req.url}`);
  next();
}

app.use(loggingMiddleware);

app.get('/', (req, res) => {
  res.send('Hello World');
});

```