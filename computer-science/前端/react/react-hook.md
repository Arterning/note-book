
## useEffect


`useEffect` 是 React 提供的一个 Hook，用于在组件渲染后执行副作用操作，例如发送网络请求、订阅事件、更新 DOM 等。下面是 `useEffect` 的基本用法：

javascript

```javascript
import React, { useEffect } from 'react';

function MyComponent() {
  useEffect(() => {
    // 执行副作用操作
    console.log('Component is mounted or updated');
    // 清理副作用操作
    return () => console.log('Component is unmounted');
  });
  // 组件渲染
  return <div>Hello, World!</div>;
}
```

在上面的代码中，`useEffect` 接受一个回调函数，这个回调函数会在组件渲染后执行。如果有返回函数，则会在组件卸载时执行这个函数，用于清理副作用操作。

如果您需要在 `useEffect` 中访问组件中的 state 或 props，可以通过将它们作为依赖项传递给 `useEffect` 来实现：

javascript

```javascript
import React, { useEffect, useState } from 'react';

function MyComponent({ count }) {
  const [message, setMessage] = useState('');

  useEffect(() => {
    // 执行副作用操作
    setMessage(`The current count is ${count}`);
  }, [count]);

  return <div>{message}</div>;
}
```

在上面的代码中，我们将组件的 `count` prop 作为依赖项传递给 `useEffect`，这意味着每次 `count` prop 发生变化时，都会重新执行 `useEffect` 中的回调函数，从而更新 `message` 状态。

需要注意的是，`useEffect` 中的副作用操作可能会对组件性能产生影响，例如在每次组件更新时都发送网络请求可能会导致频繁的网络请求，降低页面响应速度。因此，建议在使用 `useEffect` 时，根据实际情况对副作用操作进行优化，例如通过添加依赖项或使用 cleanup 函数等方式。

---

> useEffect第二个参数传递\[\]是什么意思？

---

在 `useEffect` 中，第二个参数是一个依赖项数组，用于控制何时执行 `useEffect` 中的副作用操作。这个依赖项数组是一个可选参数，可以传递一个空数组 `[]`，表示只在组件挂载和卸载时执行副作用操作，不会在组件更新时执行副作用操作。

下面是一个例子，演示了如何使用空数组作为 `useEffect` 的第二个参数：

javascript

```javascript
import React, { useEffect, useState } from 'react';

function MyComponent() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    console.log('Component is mounted');
    return () => console.log('Component is unmounted');
  }, []);

  return (
    <div>
      <div>The current count is {count}</div>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
  );
}
```

在上面的代码中，我们使用 `[]` 作为 `useEffect` 的第二个参数，这意味着 `useEffect` 中的回调函数只在组件挂载和卸载时执行，不会在组件更新时执行。在这个例子中，我们只需要在组件挂载和卸载时执行一些副作用操作，因此使用空数组作为第二个参数可以有效避免多余的副作用操作，提高组件性能。

需要注意的是，如果第二个参数传递了一个非空数组，那么这个数组中的每个元素都会被作为依赖项，只有当其中某个依赖项发生变化时，才会重新执行 `useEffect` 中的回调函数。如果传递了空数组，那么 `useEffect` 中的回调函数只会在组件挂载和卸载时执行一次。如果不传递第二个参数，那么 `useEffect` 中的回调函数会在每次组件更新时都执行。




## useMemo

useMemo 是 React 提供的一个 Hook，用于在组件渲染过程中缓存计算结果。使用 useMemo 可以避免在每次渲染时都重新计算某些值，从而提高组件的性能。

useMemo 接受两个参数：

一个函数，该函数会在组件渲染过程中被调用，并返回要缓存的值。

一个依赖数组，如果依赖数组中的任何一个值发生变化，useMemo 将重新计算缓存的值。

下面是一个使用 useMemo 的示例：

```javascript
import React, { useMemo } from 'react';

function MyComponent(props) {
  const expensiveValue = useMemo(() => {
    // 这里可以进行一些昂贵的计算，例如大量的数据处理或者复杂的函数运算
    return computeExpensiveValue(props.value);
  }, [props.value]);

  return (
    <div>
      <p>The expensive value is: {expensiveValue}</p>
      <p>Some other content...</p>
    </div>
  );
}

```
在上面的代码中，我们使用 useMemo 缓存了一个名为 expensiveValue 的值。每次组件重新渲染时，useMemo 将会检查 props.value 是否发生了变化。如果 props.value 没有变化，useMemo 将返回缓存的值，从而避免了进行昂贵的计算。如果 props.value 发生了变化，useMemo 将重新计算 expensiveValue 并返回新的值。

需要注意的是，useMemo 的返回值仅仅是缓存的值，并不是组件的状态。如果您需要在组件中存储一些状态，应该使用 useState 或者 useReducer 等 Hook。


## useReducer


> useReducer 是什么？

---

`useReducer` 是 React 提供的一个 Hook，用于管理组件中的状态。与 `useState` 不同，`useReducer` 可以管理更复杂的状态逻辑，例如需要进行多次状态更新的场景。

`useReducer` 接受两个参数：

1.  一个 reducer 函数，用于根据当前状态和动作(action)计算下一个状态。
2.  一个初始状态(initialState)。

`useReducer` 返回一个包含两个值的数组：

1.  当前状态(state)。
2.  一个 `dispatch` 函数，用于向 reducer 发送动作(action)，以触发状态更新。

下面是一个使用 `useReducer` 的示例：

javascript

```javascript
import React, { useReducer } from 'react';

function counterReducer(state, action) {
  switch (action.type) {
    case 'increment':
      return { count: state.count + 1 };
    case 'decrement':
      return { count: state.count - 1 };
    default:
      throw new Error();
  }
}

function Counter() {
  const [state, dispatch] = useReducer(counterReducer, { count: 0 });

  return (
    <div>
      <p>Count: {state.count}</p>
      <button onClick={() => dispatch({ type: 'increment' })}>+</button>
      <button onClick={() => dispatch({ type: 'decrement' })}>-</button>
    </div>
  );
}
```

在上面的代码中，我们使用 `useReducer` 来管理一个计数器的状态。`counterReducer` 函数根据动作(action)计算下一个状态，而 `Counter` 组件中的 `dispatch` 函数用于向 `counterReducer` 发送动作(action)，从而触发状态更新。

需要注意的是，`useReducer` 的好处是可以将状态逻辑分离到 reducer 函数中，使得组件更加简洁、易于维护。但是，如果您只需要管理简单的状态，例如单个数字或者字符串，`useState` 可能更加适合您的需求。

---

> useReducer 和redux

---

`useReducer` 是 React 自带的一个 Hook，与 Redux 的 reducer 概念类似，但两者有一些区别。

`useReducer` 与 `useState` 相似，都是用于管理组件状态的 Hooks，但是 `useReducer` 的优势在于处理复杂的状态逻辑，例如需要进行多次状态更新的场景。

与 Redux 的 reducer 概念类似，`useReducer` 接受一个 reducer 函数和一个初始状态(initialState)，并返回一个包含当前状态(state)和 dispatch 函数的数组。使用 dispatch 函数可以发送 action，reducer 函数根据当前状态和 action 计算下一个状态。与 Redux 不同的是，`useReducer` 只能用于组件内部的状态管理，而 Redux 可以用于全局状态管理。

Redux 可以提供更强大的状态管理能力，例如可预测的状态变化、中间件、时间旅行等，但是在一些简单的场景下使用 Redux 可能会显得过于繁琐。因此，如果您只需要在组件内部管理一些简单的状态，可以优先考虑使用 `useState` 或 `useReducer`。如果需要管理更复杂的状态逻辑，可以考虑使用 Redux 等状态管理工具。
