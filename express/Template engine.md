最后，我们的网站要开始展示一些实际内容了。Express 对当今主流的模板引擎（例如 Pug、Handlebars、EJS 等等）提供了很好的支持，可以做到两行代码接入。


```js
// 指定模板存放目录
app.set('views', '/path/to/templates');

// 指定模板引擎为 Handlebars
app.set('view engine', 'hbs');
```


在使用模板时，只需在路由函数中调用 `res.render` 方法即可：


```js
// 渲染名称为 hello.hbs 的模板
res.render('hello');
```