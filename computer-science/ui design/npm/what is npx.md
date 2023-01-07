# What is npx

npx：Node Package Execute 即 node 包执行器

npx 想要解决的主要问题，就是调用项目内部安装的模块。比如，项目内部安装了测试工具 Mocha。
一般来说，调用 Mocha ，只能在项目脚本和 package.json 的 scripts 字段里面， 如果想在命令行下调用，必须像下面这样。

```shell
# 项目的根目录下执行
$ node-modules/.bin/mocha --version
```

npx 就是想解决这个问题，让项目内部安装的模块用起来更方便，只要像下面这样调用就行了。

```shell
npx mocha --version

```

# 避免全局安装模块

除了调用项目内部模块，npx 还能避免全局安装的模块。比如，create-react-app 这个模块是全局安装，npx 可以运行它，而且不进行全局安装。

npx create-react-app my-react-app

上面代码运行时，npx 将 create-react-app 下载到一个临时目录，使用以后再删除。所以，以后再次执行上面的命令，会重新下载 create-react-app。
