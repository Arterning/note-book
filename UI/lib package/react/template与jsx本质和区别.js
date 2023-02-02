/**
 * 
 * template与jsx本质和区别
 * 
template
模版语法（html的扩展）
数据绑定使用Mustache语法（双大括号）
<span> hello: {{message}} </span>

优点 学习成本低 大量内置指令开发 组件作用域CSS

jsx
javaScript的语法扩展
数据绑定使用单大括号
<span> hello: {message}</span>

优点 灵活 灵活 还是他妈的灵活


更抽象一点来看 我们可以把组件分成两大类
1. 偏视图表现的 比较适合使用模板
2. 偏逻辑的  适合使用jsx
这两类组件的比例会根据应用类型的不同有所变化 但是整体来说表现类型的组件多于逻辑类组件
 */