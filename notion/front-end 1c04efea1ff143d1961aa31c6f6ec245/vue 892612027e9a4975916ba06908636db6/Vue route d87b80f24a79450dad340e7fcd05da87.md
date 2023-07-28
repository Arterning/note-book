# Vue route

### hash 模式

vue-router 默认 hash 模式 —— 使用 URL 的 hash 来模拟一个完整的 URL，于是当 URL 改变时，页面不会重新加载，这样的好处是无需我们在服务端做任何配置就可以无缝接入我们的后端。# 是 url 的一个锚点，记载了网页中的位置，在实际的请求中，hash并不会被带到后端。所以说Hash模式通过锚点值的改变，根据不同的值，渲染指定DOM位置的不同数据。hash 模式的原理是 onhashchange 事件(监测hash值变化)，可以在 window 对象上监听这个事件。

### history模式

由于hash模式本身会在我们的url中带一串`hash`，而且还有 `#` 符号,所以我们的URL可能会有那么一丢丢不好看，这个我觉得`感知不强`。如果你想要比如 `user/id/1` 这种比较优雅的模式，那么可以使用history模式，不过这种模式尽管带来了比较优雅的url，但是需要我们单独设置服务端，比如Nginx，我们可能要加这样一行配置，才可以保证我们的VUE项目可以正常运行：

```jsx
location / {
  try_files $uri $uri/ /index.html;
}
```

```jsx
// 这块就是配置我们的路由信息了，一个是首页 home，一个是关于页面 about
// 其中 name: 'Home' 是我们的别名， path: '/',是我们的路由路径，component是我们路由对应的组件。

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue') // 官方默认的除了Home其他视图组件都这样导入，所以后续我们也是采取这样的方式。
  }
]

// mode 说明了我们使用的是history模式，如果不设置mode，则会使用默认的hash模式。
const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

```

```jsx
<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

```

## 声明方式

```jsx
<router-link :to="{name:xxx,params:{key:value}}">valueString</router-link>
```

接受参数

```jsx
$route.params.key
```

或者参数放在url中，有点类似Django

```jsx
{
    path:'/user/:id/',
    // 其他属性省略掉了
}
```

```jsx
<router-link to="/user/1">params</router-link>
```

## 编程式路由跳转

```jsx
// 不带参数
this.$router.push({name:'xxx'})

// 带参数
this.$router.push({name:'xxx',params:{key:value}})

this.$router.push({name:'xxx',query:{key:value}})

this.$router.go(-1)//跳转到上一次浏览的页面
this.$router.replace('/menu')//指定跳转的地址
this.$router.replace({name:'menuLink'})//指定跳转路由的名字下

```

## **$router和$route的区别**

```jsx
//$router：是路由操作对象，只写对象
///$route：路由信息对象，只读对象
//操作 路由跳转
this.$router.push({
    name: 'hello',
    params: {
        name: 'world',
        age: '100'
    }
})
//读取 路由参数接收 
this.name = this.$route.params.name;
this.age = this.$route.params.age;

```

## push 和 replace 的区别：

- 使用push方法的跳转会向 history 栈添加一个新的记录，当我们点击浏览器的返回按钮时可以看到之前的页面。也就是说是支持后退的。
- 使用replace方法不会向 history 添加新记录，而是替换掉当前的 history 记录，即当replace跳转到的网页后，'后退' 按钮不能查看之前的页面。

## **param 和 query 的区别**

query可以使用name或者path方式跳转

```jsx
//query传参，使用name跳转
this.$router.push({
    name:'second',
    query: {
        queryId:'20180822',
        queryName: 'query'
    }
})

//query传参，使用path跳转
this.$router.push({
    path:'second',
    query: {
        queryId:'20180822',
        queryName: 'query'
    }
})

//query传参接收
this.queryName = this.$route.query.queryName;
this.queryId = this.$route.query.queryId;
```

params只能使用name方式进行跳转引入

```jsx
//params传参 使用name
this.$router.push({
  name:'second',
  params: {
    id:'20180822',
     name: 'query'
  }
})

//params接收参数
this.id = this.$route.params.id ;
this.name = this.$route.params.name ;

```

## **全局前置守卫**

```jsx
router.beforeEach((to, from, next) => {
    let token = localStorage.getItem('Authorization')
	if(to.name == "Login" || to.name == "Reigster"){
		next()
	}
	else{
		if(token == null){
			next('/login')
		}else{
			next()
		}
	}
})

```

- to 即将要进入的目标 路由对象
- from 当前导航正要离开的路由
- next 要执行的操作

## 导航守卫：

导航守卫的应用场景非常明确，类似于Django中的中间件，以及Spring 中的拦截器，都是在请求执行之前的一个环节，或者说做一个权限校验，合法的请求我们就放行，不合法的请求我们就重定向，或者在每个请求执行之前，做额外的一些工作，比如添加响应头，等等等等。

作者：韩数链接：https://juejin.cn/post/6844904138262741000来源：稀土掘金著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

## 全局后置守卫

```jsx
router.afterEach((to, from) => {
  // ...
})
```

## 路由独享守卫

```jsx
const router = new VueRouter({
  routes: [
    {
      path: '/foo',
      component: Foo,
      beforeEnter: (to, from, next) => {
        // ...
      }
    }
  ]
})

```

## 带参数的路由匹配

```jsx
// 这些都会传递给 `createRouter`
const routes = [
  // 动态字段以冒号开始
  { path: '/users/:id', component: User },
]
```

```jsx
const User = {
  template: '<div>User {{ $route.params.id }}</div>',
}
```

## 路由过渡动画效果

只需要使用`transition`标签包裹住`router-view`标签即可，动画效果可以自己定义，参考`transition`组件的用法。也可以在父组件或者`app.js`中使用`watch`监听`$route`变化，根据不同路由替换`transition`组件的`name`属性，实现不同的动画效。

```jsx
<transition :name="transitionName">
  <router-view></router-view>
</transition>
```

监听

```jsx
watch: {
  '$route' (to, from) {
    const toD = to.path.split('/').length
    const fromD = from.path.split('/').length
    this.transitionName = toD < fromD ? 'slide-right' : 'slide-left'
  }
}
```