# 组件优化

## Component的2个问题 

> 1. 只要执行setState(),即使不改变状态数据(旧的state可能和新的state一样), 组件也会重新render()
>
> 2. 只当前组件重新render(), 就会自动重新render子组件 ==> 效率低

## 效率高的做法

>  只有当组件的state或props数据发生改变时才重新render()

## 原因

>  Component中的shouldComponentUpdate()总是返回true

## 解决

        办法1: 
            重写shouldComponentUpdate()方法
            比较新旧state或props数据, 如果有变化才返回true, 如果没有返回false
        办法2:  
            使用PureComponent
            PureComponent重写了shouldComponentUpdate(), 只有state或props数据有变化才返回true
            注意: 
                只是进行state和props数据的浅比较, 如果只是数据对象内部数据变了, 返回false  
                不要直接修改state数据, 而是要产生新数据
        项目中一般使用PureComponent来优化


## 函数组件如何实现

函数组件如何实现类似PureComponent的效果？

在函数式组件中，无法像类组件一样继承PureComponent。因为函数式组件没有实例，无法在组件内部维护shouldComponentUpdate生命周期函数，也就无法像类组件一样自动进行浅比较来决定是否更新组件。

但是可以通过React.memo()方法来实现PureComponent类似的优化。React.memo()是一个高阶组件，可以将函数式组件包装起来，返回一个新的组件。React.memo()通过比较前后两次组件接收到的props是否有变化来判断是否需要重新渲染组件。

例如：

``` javascript
import React from 'react';

function MyComponent(props) {
  // ...
}

export default React.memo(MyComponent);

```
这样，当父组件重新渲染并传递给MyComponent的props没有变化时，MyComponent就不会重新渲染，从而达到类似PureComponent的效果。