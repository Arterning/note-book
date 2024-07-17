


```css
img {
    height:15px;
    width:15px;
}

img:hover{
    height: 450px;
    width: 450px;
}


img{
    transition: 1sv ease;
}
```


除了ease以外，其他模式还包括

> （1）linear：匀速
> 
> （2）ease-in：加速
> 
> （3）ease-out：减速
> 
> （4）cubic-bezier函数：自定义速度模式


### transition的局限

transition的优点在于简单易用，但是它有几个很大的局限。

（1）transition需要事件触发，所以没法在网页加载时自动发生。

（2）transition是一次性的，不能重复发生，除非一再触发。

（3）transition只能定义开始状态和结束状态，不能定义中间状态，也就是说只有两个状态。

（4）一条transition规则，只能定义一个属性的变化，不能涉及多个属性。

CSS Animation就是为了解决这些问题而提出的

## CSS Animation

```css
div:hover {
// 动画时间 动画对应的关键帧 动画无线循环展示
  animation: 1s rainbow infinite;
}
```


```css
@keyframes rainbow {
  0% { background: #c00; }
  50% { background: orange; }
  100% { background: yellowgreen; }
}
```

