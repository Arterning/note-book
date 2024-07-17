
https://juejin.cn/post/7208440329652650043


## 受控组件和非受控组件

通过通过自己维护一个 state 来获取或更新 input 输入的值，以这种方式控制取值的 **表单输入元素** 就叫做 **受控组件**。

非受控组件是指表单元素的值由 DOM 节点直接处理，而不是由 React 组件来管理。当然，这意味着当用户输入数据时，React 无法追踪表单元素的值。因此，当您需要访问表单元素的值时，您需要使用DOM API来获取它们。



### 为什么需要非受控组件

在 React 中，通常使用受控组件来处理表单。受控组件表单元素的值由 React 组件来管理，当表单数据发生变化时，React 会自动更新组件状态，并重新渲染组件。这种方式可以使得表单处理更加可靠和方便，也可以使得表单数据和应用状态之间保持一致。

但在实际的开发中，表单往往是最复杂的场景，有的表单有数十个字段，如果使用受控组件去构建表单，那么我们就需要维护大量 state，且 React 又不像 Vue 可以通过双向绑定直接修改 state 的值，每一个表单字段还需要定义一下 `onChange` 方法。因此在维护复杂表单时，使用受控组件会有很大的额外代码量。

  




## 基本使用

我们不需要定义任何 state 即可在 submit 时获取到表单中的数据

- 用`register` 注册表单Field，并且指定表单验证规则，不需要维护大量的state
- `onsubmit`可以拿到提交的表单数据


```js
import React from "react";
import { useForm } from "react-hook-form";

function MyForm() {
  const { register, handleSubmit } = useForm();
  const onSubmit = (data) => console.log(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register("firstName")} />
      
      <input {...register("lastName")} />
      
      <button type="submit">Submit</button>
    </form>
  );
}

```