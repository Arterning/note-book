**简介：** React 状态管理工具五花八门，dva、mobx、recoil、zustand。换做是你，你会怎么选呢？选择一个合适的状态管理工具，对项目研发是至关重要的，来看看我的选择方案吧


在状态共享这方面，不像 `Vuex`，`React` 的官方并没有强力推荐某种封装方案，所以 `React` 的状态管理工具五花八门，百花齐放。其中就有：

- 做什么都要 `dispatch` 的 `redux` 流派。包括：`react-redux` 、`dva`、新星代表 `zustand`
- 响应式流派 `mobx`。以及新星代表 `valtio` ，以及一个很有特点的库 `resso`
- 原子状态流派。来自 `facebook` 开源的 `recoil` ，以及新星代表 `jotai`
- 完全体 `hooks` 流派。`hox`、`reto`、`umijs@4 内置数据流`，包括 `Vue` 官方推荐的新状态管理工具 `pinia` 也是这个流派。

# 我们需要什么样的状态管理工具

## 可能不需要？

我们阅读一些状态管理工具的文档时候，可能就会先被这样一篇文章甩在脸上：[《你可能不需要状态管理工具》](https://medium.com/@dan_abramov/you-might-not-need-redux-be46360cf367?spm=a2c6h.12873639.article-detail.10.7dff7db6ItEaFI)。


## 对状态管理工具的要求？

随着项目经验的积累，我总结出了状态管理工具应该满足的几个特性：

1. 共享状态（基础），能够满足上面列举的几个场景；
2. 共享业务逻辑，比如修改个人密码后需要退出并跳转至登录页（在菜单栏和个人中心都要调用，相同的逻辑不应该写多次）；
3. 共享状态模块化，即按不同业务逻辑，分开不同的文件创建共享状态。
4. 再复杂一些的，涉及到共享状态之间的依赖，比如当我修改当前登录人的角色之后（比如从”项目经理“切换至”系统管理员“），记录菜单权限的状态也需要更新。
5. 使用时，有清晰的依赖来源（`import from`）。
6. 对 TypeScript 支持良好，易编写。




## zustand

拥有 22K stars 的 [`zustand`](https://github.com/pmndrs/zustand) 则是非常值得尝试的传统派替代品。

- 它面向 `hooks` API，一个 `store` 就是一个 `hooks`。  
- 它更简洁，直接将 `state` / `reducers` / `effects` 平铺开来，`function` 即 `effects` / `reducers` ，其它都是 `state`。
- `import` 来源清晰，对 `TypeScript` 的支持也更友好
- 提供了 `subscribe` 接口，可在组件外监听状态变化，以实现状态依赖。

```js
import { create } from 'zustand'

const useStore = create((set) => ({
  count: 1,
  inc: () => set((state) => ({ count: state.count + 1 })),
}))

function Counter() {
  const { count, inc } = useStore()
  return (
    <div>
      <span>{count}</span>
      <button onClick={inc}>one up</button>
    </div>
  )
}
```

set 相当于redux中的dispatch


## Mobx


```jsx
// TodoStore.js
import { makeObservable, observable, action } from 'mobx';

class TodoStore {
  todos = [];

  constructor() {
    makeObservable(this, {
      todos: observable,
      addTodo: action,
      toggleTodo: action,
      removeTodo: action,
    });
  }

  addTodo = (text) => {
    this.todos.push({ text, completed: false });
  };

  toggleTodo = (index) => {
    this.todos[index].completed = !this.todos[index].completed;
  };

  removeTodo = (index) => {
    this.todos.splice(index, 1);
  };
}

const todoStore = new TodoStore();
export default todoStore;

```

```jsx
// TodoList.js
import React, { useState } from 'react';
import { observer } from 'mobx-react';
import todoStore from './TodoStore';

const TodoList = observer(() => {
  const [newTodo, setNewTodo] = useState('');

  const handleAddTodo = () => {
    todoStore.addTodo(newTodo);
    setNewTodo('');
  };

  const handleToggleTodo = (index) => {
    todoStore.toggleTodo(index);
  };

  const handleRemoveTodo = (index) => {
    todoStore.removeTodo(index);
  };

  return (
    <div>
      <h2>Todo List</h2>
      <ul>
        {todoStore.todos.map((todo, index) => (
          <li key={index}>
            <input
              type="checkbox"
              checked={todo.completed}
              onChange={() => handleToggleTodo(index)}
            />
            {todo.text}
            <button onClick={() => handleRemoveTodo(index)}>Remove</button>
          </li>
        ))}
      </ul>
      <div>
        <input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
        />
        <button onClick={handleAddTodo}>Add Todo</button>
      </div>
    </div>
  );
});

export default TodoList;

```

```jsx
// App.js
import React from 'react';
import TodoList from './TodoList';

function App() {
  return (
    <div>
      <h1>MobX Todo App</h1>
      <TodoList />
    </div>
  );
}

export default App;

```


mobx 的出现给 redux 带来了很大的冲击。

通过一个装饰器（observer / observable），就能使普通的组件能够监听变量的变化而渲染，完全抛弃掉 state 。不仅如此，mobx 提供了 computed 等一些好东西，使 React 也能使用到 Vue 组件的特性。而从 Vue 转过来学习 React 的同学，都会对 mobx 拍手称赞。

在体验过 mobx 的爽快之后，当时团队中有部分声音，是希望以后抛弃掉 state，转而全面使用 mobx 作为组件状态。是啊，一套方案，解决内部状态和共享状态两个问题，何乐而不为呢。


### 缺憾

很快，这种 `mobx` 为王的声音很快又消失了，因为...

- 响应式 API 和 `React` 水土不服，`React` 就是需要 `setState` 来修改状态，现在你 `state.value++` 就改了，是要造反么？新来的同学学完 `React` 基础，结果和他说，那些用不上，再学学 `mobx` 响应式吧，新同学是否会心中充满问号？
- 基础组件如果也使用 `mobx` 则违反了高内聚的原则，不使用，两边风格又不统一。你见过哪个组件库需要附带一套状态管理工具的？
- 响应式其实就是基于 `Proxy` 实现的，我明明希望传递的是一个数组，但拿到的却是一个 `Proxy`强迫症实在受不了。
- 会有一些隐蔽的坑，比如往 `observable` 添加属性时，不能直接添加，而要通过 `extendObservable` ，我们有很多同学踩了这坑。

所以，`mobx` 很优秀，但我实在爱不来。


## Recoil/Jotai
一个 **atom** 代表一个**状态**。Atom 可在任意组件中进行读写。读取 atom 值的组件隐式订阅了该 atom，因此任何 atom 的更新都将致使使用对应 atom 的组件重新渲染：

```js
function CharacterCounter() {
  return (
    <div>
      <TextInput />
      <CharacterCount />
    </div>
  );
}

function TextInput() {
  const [text, setText] = useRecoilState(textState);

  const onChange = (event) => {
    setText(event.target.value);
  };

  return (
    <div>
      <input type="text" value={text} onChange={onChange} />
      <br />
      Echo: {text}
    </div>
  );
}
```

**selector** 代表一个**派生状态**，派生状态是状态的**转换**。你可以将派生状态视为将状态传递给以某种方式修改给定状态的纯函数的输出：
```js
const charCountState = selector({
  key: 'charCountState', // unique ID (with respect to other atoms/selectors)
  get: ({get}) => {
    const text = get(textState);

    return text.length;
  },
});
```
我们可以使用 `useRecoilValue()` 的 hook，来读取 `charCountState` 的值：
```js
function CharacterCount() {
  const count = useRecoilValue(charCountState);

  return <>Character Count: {count}</>;
}
```
为了解决共享状态依赖的需求，recoil 还很贴心地提供了 selector API，用于实现共享状态的拆分和依赖，你把它当作 useMemo 或者计算属性来看待就可以了。（当然 selector 还有支持写入(set)以及异步处理，但我还没找到必须要用它的场景）

不足
recoil 理念真的很简单，就是以 useState 的习惯实现状态共享。所以在业务逻辑共享这一块，它似乎没有给出很好的方案。但是既然已经是面向 hooks API 了，自定义 hooks 本身就可以实现业务逻辑的复用了


## hooks 完全体  hox
`hox` 刚出来不久，我就关注到了，并觉得其思想非常棒。但翻阅了一下源码后发现，它依赖了一个实验性的渲染器 `react-reconciler`，以至于我不敢将它用于生产环境。直到 `react@18.x` 带来了一个新的 `hooks`: [`useSyncExternalStore`](https://reactjs.org/docs/hooks-reference.html#usesyncexternalstore) ，以及基于它实现的 `hox@2.x` 。

我们来看它的介绍：

> 在 hox 中，任意的自定义 Hook，经过  `createStore` 包装后，就变成了持久化，可以在组件间进行共享的状态。

我的天，真的太神奇了，你只要用 `createStore` 包裹你写的某个 `hooks` ，它里面的状态就变成了可共享的了。


### 实现原理

举个简单的例子，我写了一个自定义 `hooks` useMyHook:

```
export const useMyHook = () => {
  const [value, setValue] = useState(1);

  return [value, setValue];
}
```

我在 `组件 A` 和 `组件 B` 中都使用了它。正常情况下，A 和 B 中的 value 当然是不同的。

但是假如我“偷偷地”将 `hooks` 放在最外面执行，比如 `App` 下，然后再用 `Context` 传递下去：

```
// App.tsx
export const Context = React.createContext({});

export default function App() {

  // 在最外层执行 hooks
  const myHookResult = useMyHook();

  // 通过 Context 向下传递
  return (
    <Context.Provider value={myHookResult}>
      {children}
    </ContextProvider>
  )
}
```

我再给你一个封装后的 Hook:

```
// 在 组件A 和 组件B 中使用这个 hooks
export const useMyHookWrapped = () => {
  // 从外部获取到 useMyHook 的内容
  const myHookResult = useContext(Context);
  return myHookResult;
}
```

综上，你就会发现，相当于 `useMyHook` 只被使用了一次，其它需要用到的地方，都是使用 `useContext` 获取的。那么自然就实现了状态共享了。而这些，都是 `createStore` 实现的。

并且，和状态相关的业务逻辑，也写在了同一个 `hooks`，还可以不受限制地获得完整的 `hooks` 体验，使用第三方 `hooks` 库。**最令人惊叹的是**，由于都是 `hooks` API，你可以先在组件中编写业务逻辑，当发现逻辑需要共享时，直接复制抽离出去；或者是当你需要迁移部分功能到另外的项目，不需要共享了，只需要去掉 `createStore` ，它就又变成了普通 `hooks` 了。



综上，你可以发现，`hox` 的方案算是所有工具中最好的，能够满足所有的场景，并获得不错的评分。我也会在今后的工作中更深入地使用 `hox` 方案，发掘其是否还有更优秀的用法或隐藏的问题。

最后也欢迎各位一同讨论你们在状态管理工具上的取舍。


# PS：为什么不选

## zustand

`zustand` 很优秀，但使用它仍然会让我联想到在 `effects` 中使用 `dispatch` 更新状态的痛苦时光。是啊，在 `zustand` 你还是要使用 `set` 来修改状态。虽然它也提供了额外的 `setState` 方法让你直接更新状态，但是风格和 `useState` 差得远啊，让我觉得不伦不类。

## mobx / valtio

这两个很明显了，响应式的风格，跟 `React` 对着干，我当然是不选的。

## recoil / jotai

我原本是比较倾向于 `recoil` 的，奈何 `selector` 的用法有一定的上手成本，让新手学完 `useMemo` 就能运用上不好么。

> 细心点你会发现，zustand / valtio / jotai 是同一个开源团队的作品，而且他们的名字分别是 德语、芬兰语、日语 中的“状态”。所以，不选择某个工具，不是说这些工具不好，他们之间并没有孰优孰劣，只是面向风格不一样罢了。