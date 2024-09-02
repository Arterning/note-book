
### **Functional components are more reusable**

This one might be a bit controversial. But by removing function level state, we often make our components easier to use and more widely applicable. Let’s take a peek at two implementations of a custom checkbox.

**Functional Checkbox:**

```javascript
const Checkbox = ({ checked, label, handleClick }) => (
   <div
      className={checked ? 'Checkbox-container checked' : 'Checkbox-container'}
      onClick={handleClick}
      role="button"
      tabIndex={0}
      data-label={label}
   >
      <p className="label" data-label={label}>{label</p>
   </div>
);
```

Making the checkbox a functional component forces us to strip the component down to its most primitive features, which has the side effect of making it more generally applicable. Think of it as a forced best practice.

The checkbox component does not have to choose what its default state should be or what it should do when clicked. Instead, this is delegated to the higher level component. Could you write a class component in the exact same way?  Sure, but a functional component **forces** us to follow best practices which results in a cleaner, more reusable component.



## We can extract business logic via hooks

because we can split state into different part instead of put it together in class state, we can extract business logic like this, this is called `use customized hook`


For example, we can define a useCounter hook like this : 
Then we can use this hook wherever we want
```js
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
```


### **Functional components are easy to test**

It’s easier to test functional components because you don’t have to worry about hidden state or [side effects](https://en.wikipedia.org/wiki/Side_effect_(computer_science)). For every input (props), functional components have exactly one output.

Given certain props, I can assert exactly what the HTML output will be. This means you don’t have to rely on any mocking, state manipulation, or crazy testing libraries. It’s pretty awesome.

### **Functional components can potentially have a better performance**

Since functional components offer no state or lifecycle methods, you would think that this would allow the internals of React to avoid unnecessary overhead such as lifecycle events. Unfortunately, this is currently not the case.