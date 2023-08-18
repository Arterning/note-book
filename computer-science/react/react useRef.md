# useRef

修改State和Props都会导致组件重新渲染，但是修改Ref不会导致组件重新渲染

useRef，他的作用是“勾住”某些组件挂载完成或重新渲染完成后才拥有的某些对象，并返回该对象的引用。该引用在组件整个生命周期中都固定不变，该引用并不会随着组件重新渲染而失效。

这句话中的“某些对象”主要分为3种：JSX组件转换后对应的真实DOM对象、在useEffect中创建的变量、子组件内自定义的函数(方法)。

勾住渲染结束后的dom元素

```javascript
import React,{useEffect,useRef} from 'react'

function Component() {
  //先定义一个inputRef引用变量，用于“勾住”挂载网页后的输入框
  const inputRef = useRef(null);

  useEffect(() => {
    //inputRef.current就是挂载到网页后的那个输入框，一个真实DOM，因此可以调用html中的方法focus()
    inputRef.current.focus();
  },[]);

  return <div>
      {/* 通过 ref 属性将 inputRef与该输入框进行“挂钩” */}
      <input type='text' ref={inputRef} />
    </div>
}
export default Component
```


```javascript
import React,{useState,useEffect,useRef} from 'react'

function Component() {
  const [count,setCount] =  useState(0);
  const timerRef = useRef(null);//先定义一个timerRef引用变量，用于“勾住”useEffect中通过setIntervale创建的计时器

  useEffect(() => {
    //将timerRef.current与setIntervale创建的计时器进行“挂钩”
    timerRef.current = setInterval(() => {
        setCount((prevData) => { return prevData +1});
    }, 1000);
    return () => {
        //通过timerRef.current，清除掉计时器
        clearInterval(timerRef.current);
    }
  },[]);

  const clickHandler = () => {
    //通过timerRef.current，清除掉计时器
    clearInterval(timerRef.current);
  };

  return (
    <div>
        {count}
        <button onClick={clickHandler} >stop</button>
    </div>
  )
}

export default Component
```