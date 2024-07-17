## hox 介绍

hox 是完全拥抱 React Hooks 的状态管理器，model 层也是用 custom Hook 来定义的，它有以下几个特性：

- 只有一个 API，简单高效，几乎无需学习成本
- 使用 custom Hooks 来定义 model，完美拥抱 React Hooks
- 完美的 TypeScript 支持
- 支持多数据源，随用随取

下面我们进入正题，hox 怎么用？

### 定义 Model

任意一个 custom Hook ，用 `createModel` 包装后，就变成了持久化，且全局共享的数据。

```js
import { createModel } from 'hox';

/* 任意一个 custom Hook */
function useCounter() {
  const [count, setCount] = useState(0);
  const decrement = () => setCount(count - 1);
  const increment = () => setCount(count + 1);
  return {
    count,
    decrement,
    increment
  };
}

export default createModel(useCounter)
```


### 使用 Model

`createModel` 返回值是个 Hook，你可以按 React Hooks 的用法正常使用它。

```jsx
import { useCounterModel } from "../models/useCounterModel";

function App(props) {
  const counter = useCounterModel();
  return (
    <div>
      <p>{counter.count}</p>
      <button onClick={counter.increment}>Increment</button>
    </div>
  );
}
```

`useCounterModel` 是一个真正的 Hook，会订阅数据的更新。也就是说，当点击 "Increment" 按钮时，会触发 counter model 的更新，并且最终通知所有使用 `useCounterModel` 的组件或 Hook。


### 其它

- 基于上面的用法，你肯定已经知道了在 model 之间互相依赖怎么写了，就是单纯的 Hooks 互相依赖，自然而然咯。
```jsx
import { useCounterModel } from "./useCounterModel";

export function useCounterDouble() {
  const counter = useCounterModel();
  return {
    ...counter,
    count: counter.count * 2
  };
}
```

- 只读不订阅更新就更简单了。

```js
import { useCounterModel } from "./useCounterModel";

export function useCounterDouble() {
  const counter = useCounterModel.data;
  return {
    ...counter,
    count: counter.count * 2
  };
}
```


## Refer
- https://segmentfault.com/a/1190000020811761
