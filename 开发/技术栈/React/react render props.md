  
## 前言  
为什么需要做render props?  
  
* 原始的做法：  
  
```javascript  
<A>  
    <B>xxxx</B></A>  
```  
存在的问题在于: 如果B组件需要A组件内的数据是做不到的  
  
* 解决方式：  
  
向组件内部动态传入带内容的结构(标签)  
  
   Vue中:   
使用slot技术, 也就是通过组件标签体传入结构  <AA><BB/></AA>   React中:  
      使用children props: 通过组件标签体传入结构  
      使用render props: 通过组件标签属性传入结构, 一般用render函数属性  
  
## 概念  
Render props 模式是一种 React 组件设计模式，用于将组件之间的代码复用。该模式通过将一个函数作为组件的 props，使得这个函数可以被组件内部调用，并且将组件的状态和逻辑传递给这个函数。这个函数就成为了一个“渲染 prop”，因为它负责渲染组件的一部分内容。  
  
在 React 中，常见的使用 render props 模式的组件有很多，例如 React Router 的 Route 组件，React 的 Context 组件，以及像 Formik 这样的表单处理库。这些组件在使用时都需要传递一个函数作为 props，这个函数接收组件的状态和逻辑，然后渲染组件的一部分内容。  
  
Render props 模式可以使组件的代码更加可复用和灵活，使得组件可以更加专注于处理数据和逻辑，而不是关注于如何渲染。同时，这种模式也可以帮助组件避免继承和 mixin 等混乱的代码复用方式，使得组件更加清晰和易于维护。  
  
  
  
## 例子  
  
``` javascript

import React, { Component } from 'react';  
  
class Tooltip extends Component {  
  state = { text: '' };  
  showTooltip = text => {    this.setState({ text });  };  
  hideTooltip = () => {    this.setState({ text: '' });  };  
  render() {    const { text } = this.state;  
    return this.props.children({ showTooltip: this.showTooltip, hideTooltip: this.hideTooltip, text });  }}  
  
class App extends Component {  
  render() {    return (      <div>        <Tooltip>          {({ showTooltip, hideTooltip, text }) => (            <div>              <button onMouseEnter={() => showTooltip('Hello World')} onMouseLeave={hideTooltip}>                Hover me              </button>              {text && <div>{text}</div>}            </div>          )}        </Tooltip>      </div>    );  }}  
  
export default App;
  
```