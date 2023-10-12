# Layout 的概念

在SPA中，Layout表示页面布局，相当于是一个父组件，比如管理系统，需要有左侧导航栏，顶部面包屑，底部区域。
这些都是表示系统的整体布局和外观，可以抽取成一个Layout组件，同时在路由这样注册。

比如主页

``` javascript
{
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: (resolve) => require(['@/views/home'], resolve),
        name: 'Dashboard',
        meta: { title: '首页', icon: 'index', affix: true, noCache: true }
      }
    ]
  }
```

子路由的路径包括了它的所有父级路由的路径。在这个例子中，父路由的路径是根路径（/），而子路由的路径是dashboard，因此完整的访问路径是/dashboard。

此外，动态添加的路由也应该是挂在Layout下的。