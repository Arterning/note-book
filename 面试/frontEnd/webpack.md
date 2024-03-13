

## Loader和Plugin的区别

loader一般是将某个语法统一处理为统一的语法  它是一个转换器，将A文件进行编译成B文件，比如：将A.sss转换为A.css，单纯的文件转换过程。

loader:让webpack能够处理非js文件(自身职能理解js)，然后你就可以利用 webpack 的打包能力，对它们进行处理。  
例如：css-loader、style-loader、postcss-loader、sass-loader



plugin一般是在打包前或打包后对结果进行再次操作，它并不直接操作文件，而是基于事件机制工作，会监听webpack打包过程中的某些节点，执行广泛的任务



## Loader

Loader 本质就是一个函数，在该函数中对接收到的内容进行转换，返回转换后的结果。

因为 Webpack 只认识 JavaScript，所以 Loader 就成了翻译官，对其他类型的资源进行转译的预处理工作。

Loader 在 module.rules 中配置，作为模块的解析规则，类型为数组。每一项都是一个 Object，内部包含了 test(类型文件)、loader、options (参数)等属性。



总结：  loader是翻译官，plugin是干活滴


## Plugin



Plugin 在 plugins 中单独配置，类型为数组，每一项是一个 Plugin 的实例，参数都通过构造函数传入。


比如 HtmlWebpackPlugin
该插件将为你生成一个 HTML5 文件， 在 body 中使用 `script` 标签引入你所有 webpack 生成的 bundle。 只需添加该插件到你的 webpack 配置中，如下所示：

```js
  plugins: [
    new HtmlWebpackPlugin({
      title: "博客列表",
    }),
    new BundleAnalyzerPlugin(),
  ],
```