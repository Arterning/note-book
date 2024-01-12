通常网站需要提供静态文件服务，例如图片、CSS 文件、JS 文件等等，而 Express 已经自带了静态文件服务中间件 `express.static`，使用起来非常方便。

例如，我们添加静态文件中间件如下，并指定静态资源根目录为 `public`：

```js
// ...

app.use(express.static('public'));

app.get('/', (req, res) => {
  res.render('index');
});

// ...

```




假设项目的 public 目录里面有这些静态文件：

```
public
├── css
│   └── style.css
└── img
    └── tuture-logo.png
```


就可以分别通过以下路径访问：

```
http://localhost:3000/css/style.css
http://localhost:3000/img/tuture-logo.png
```