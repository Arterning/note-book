# useContext

[详细笔记地址](https://github.com/puxiao/react-hook-tutorial/blob/master/03%20useState%E9%AB%98%E7%BA%A7%E7%94%A8%E6%B3%95.md)

```jsx
import React,{ useContext } from 'react'

const UserContext = React.createContext();
const NewsContext = React.createContext();

function AppComponent() {
  return (
    <UserContext.Provider value={{name:'puxiao'}}>
        <NewsContext.Provider value={{title:'Hello React Hook.'}}>
            <ChildComponent />
        </NewsContext.Provider>
    </UserContext.Provider>
  )
}

function ChildComponent(){
  const user = useContext(UserContext);
  const news = useContext(NewsContext);
  return <div>
    {user.name} - {news.title}
  </div>
}

export default AppComponent;
```

useReducer

```jsx
import React, { useReducer } from 'react';

function reducer(state,action){
  switch(action){
    case 'add':
        return state + 1;
    case 'sub':
        return state - 1;
    case 'mul':
        return state * 2;
    default:
        console.log('what?');
        return state;
  }
}

function CountComponent() {
  const [count, dispatch] = useReducer(reducer,0);

  return <div>
    {count}
    <button onClick={() => {dispatch('add')}} >add</button>
    <button onClick={() => {dispatch('sub')}} >sub</button>
    <button onClick={() => {dispatch('mul')}} >mul</button>
  </div>;
}

export default CountComponent;
```

dispatch 传递对象

```jsx
import React, { useReducer } from 'react';

function reducer(state,action){
  //根据action.type来判断该执行哪种修改
  switch(action.type){
    case 'add':
      //count 最终加多少，取决于 action.param 的值
      return state + action.param;
    case 'sub':
      return state - action.param;
    case 'mul':
      return state * action.param;
    default:
      console.log('what?');
      return state;
  }
}

function getRandom(){
  return Math.floor(Math.random()*10);
}

function CountComponent() {
  const [count, dispatch] = useReducer(reducer,0);

  return <div>
    {count}
    <button onClick={() => {dispatch({type:'add',param:getRandom()})}} >add</button>
    <button onClick={() => {dispatch({type:'sub',param:getRandom()})}} >sub</button>
    <button onClick={() => {dispatch({type:'mul',param:getRandom()})}} >mul</button>
  </div>;
}

export default CountComponent;
```