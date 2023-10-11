我忍不住想知道对于 React 功能组件使用普通函数和粗箭头之间是否有任何优势

```javascript
const MyMainComponent = () => (
  <main>
    <h1>Welcome to my app</h1>
  </main>
)

function MyMainComponent() {
  return (
    <main>
      <h1>Welcome to my app</h1>
    </main>
  )
}
```

当然，两者都工作得很好，但是有推荐的方法来编写它们吗？有支持其中之一的论据吗？

编辑：我知道普通 javascript 函数（即上下文、堆栈跟踪、返回关键字等）的差异可能会对函数的使用类型产生影响。然而我纯粹是根据 React 组件来问这个问题的。



主要是偏好问题。然而，存在一些（微小的、几乎微不足道的）差异：

- `return`如果您直接返回 JSX，而不需要任何先前的表达式，则粗箭头语法允许您省略花括号和关键字。使用 ES5 函数，您必须拥有`{ return ... }`.
    
- 粗箭头语法不会创建 的新上下文`this`，而 ES5 函数会创建新的上下文。`this`当您想要在函数内部引用 React 组件或想要跳过该`this.foo = this.foo.bind(this);`步骤时，这可能很有用。
    

它们之间有更多差异，但在 React 中编码时它们很少是相对的（例如使用`arguments`、`new`等）。

_就个人而言，我尽可能使用粗箭头语法，因为我更喜欢这种语法。_


原文：
https://stackoverflow.com/questions/54331084/function-or-fat-arrow-for-a-react-functional-component#:~:text=There%20is%20no%20practical%20difference,MyMainComponent()%20%7B%20return%20React.
