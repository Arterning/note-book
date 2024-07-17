
## **flex是flex-grow，flex-shrink，flex-basis的集合体**

上面这三个属性，其实我们在实际开发中并不常见。这是因为我们用flex这个属性来表达了上面三个属性。**flex**是flex-grow, flex-shrink, flex-basis的组合体，若上面三个值都取默认，则flex的值为0 1 auto。flex有两个快捷值：auto(1 1 auto)和none(0 0 auto)。前者代表等比例伸缩，后者代表不伸缩，这也是我们实际开发中常用的两个值。




```
flex: flex-grow flex-shirnk flex-basis
flex: 0 1 auto
flex: 1 1 auto
flex: 0 0 auto
```


flex-basis默认是auto，也就是容器原本的大小。当 flex-basis 值为 0% 时，是把该容器视为零尺寸的。

如果子容器同时设置了width与flex-basis属性，那么flex-basis优先