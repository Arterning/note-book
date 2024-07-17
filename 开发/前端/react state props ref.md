
state and props change will render the UI

UI = render(state, props)

so , we often use state and props in the jsx 

```jsx
state.users.map(user => {
	<div>user.id</div>
	....
})

props.todos.map(todo => {
	<div>todo.name</div>
	...
})
```


what is ref ?

`useRef` is a React Hook that lets you reference a value that’s not needed for rendering


ref is an object like this:
```js
{
	current: value
}
```


we can use ref to hold the value like interval, dom element ...

React expects that the body of your component [behaves like a pure function](https://react.dev/learn/keeping-components-pure):

- If the inputs ([props](https://react.dev/learn/passing-props-to-a-component), [state](https://react.dev/learn/state-a-components-memory), and [context](https://react.dev/learn/passing-data-deeply-with-context)) are the same, it should return exactly the same JSX.
- Calling it in a different order or with different arguments should not affect the results of other calls.

Reading or writing a ref **during rendering** breaks these expectations.





















