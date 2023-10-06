## HashRouter和BrowserRouter的区别

HashRouter 只会修改URL中的哈希值部分；而 BrowserRouter 修改的是URL本身
HashRouter 是纯前端路由，可以通过输入URL直接访问；使用时 BrowserRouter 直接输入URL会显示404，除非配置Nginx将请求指向对应的HTML文件。初次进入 / 路径时或点击 Link 组件跳转时不会发送请求
