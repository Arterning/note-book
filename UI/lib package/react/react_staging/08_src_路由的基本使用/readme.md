react路由

SPA Single Page Application
单页面应用 是不是就是一个页面 那个页面就是index.html吧 我佛了
虽然是单页面 但是多组件！！

外面必须包一层BrowserRouter。。。
除了BrowserRouter
还有HashRouter 就是在路径前面有个#符号
#符号后面的内容不会发送给服务器

Route负责注册路由 但是居然不是在js代码中实现的 而是在jsx定义的
<Route path="/about" component={About}/>
<Route path="/home" component={Home}/>


Link负责生成超链接 是使用路由的地方
<Link className="list-group-item" to="/about">About</Link>

路由的原理应该就是history.push()..

路由组件和一般组件的区别
路由组件可以接受到路由器发送的非常重要的props信息
有History location match ...


node react远程前端工作比较多
不过还是要学好一点 
和别人有差异 那就英文
这个资料确实有用 注意一下react-stage下的readme
已经做了总结了 不懂的地方再看一下视频 基本可以快速复习。