****在React项目中使用css-in-js方案****

```jsx
const Widget = () => {
  <div style={{
      color: 'white',
      fontSize: '12px'
  }} onClick={() => doSometing()}>
    text  
  </div>
}

```

但是这种写法的弊端在于，react中的style仅仅是简单的Object，不支持复杂的嵌套、选择器等特性，使用起来很不方便。 因此，便出现了大量的三方库来进行拓展，这些库统称为css-in-js。它们一般都支持样式嵌套、选择器、主题配置等特性。

> 注意 style 是双括号 className是单括号



有人专门统计了现有的css-in-js三方库，轮子不要太多： [css in js 三方库一览](https://link.juejin.cn/?target=http%3A%2F%2Fmichelebertoli.github.io%2Fcss-in-js%2F) 。比较流行的主要有: styled-components, emotion, glamorous。

```jsx
// Create a Title component that'll render an <h1> tag with some styles
const Title = styled.h1`
  font-size: 1.5em;
  text-align: center;
  color: #BF4F74;
`;

// Create a Wrapper component that'll render a <section> tag with some styles
const Wrapper = styled.section`
  padding: 4em;
  background: papayawhip;
`;

// Use Title and Wrapper like any other React component – except they're styled!
render(
  <Wrapper>
    <Title>
      Hello World!
    </Title>
  </Wrapper>
);
```

****[在React项目中使用css-in-js方案](https://juejin.cn/post/6844903993047531533)****