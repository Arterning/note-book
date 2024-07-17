

BFC 是**「块级格式化上下文」**，是一个独立的渲染区域，让处于 BFC 内部的元素与外部的元素相互隔离，使内外元素的定位不会相互影响。


- BFC 中子元素的 margin box 的左边， 与包含块 (BFC) border box的左边相接触 (子元素 absolute 除外)
    
- BFC 的区域不会与 float 的元素区域重叠
    
- `BFC` 就是一个块级元素，块级元素会在垂直方向一个接一个的排列
    
- `BFC` 就是页面中的一个隔离的独立容器，容器里的标签不会影响到外部标签
    
- 垂直方向距离由 `margin` 决定，属于同一个 `BFC` 的两个相邻标签外边距会发生重叠，属于不同 `BFC` 的不会重叠
    
- 计算 `BFC` 的高度时，浮动元素也参与计算 （可以用于清除浮动）



## 如何让元素变成BFC

最简单： 把overflow设置成hidden



## BFC 特性及应用

**「1. 阻止`margin`重叠」**

对于同一个 BFC 内的两个相邻元素，如果分别设置了上下边距，那么会发生折叠（取大的那个）

```html
<head>
div{
    width: 50px;
    height: 50px;
    background: pink;
    margin: 50px;
}
</head>
<body>
    <div></div>
    <div></div>
</body>
```




```html
<style>
.container {  
    overflow: hidden;  
}  
  
.child {  
    width: 50px;  
    height: 50px;  
    background: pink;  
    margin: 50px;  
}
</style>

<div class="container">
    <div class="child"></div>
</div>
<div class="container">
    <div class="child"></div>
</div>
```




**「2. 清除浮动」**

我们都知道，浮动的元素会脱离普通文档流，来看下面一个例子

```html
<style>
  .container {
    width: 500px;
    border: 1px solid #000;
  }
  
  .child {
    float: left;
    width: 100px;
    height: 100px;
    background: pink;
  }
</style>

<div class="container">
  <div class="child"></div>
</div>
```

解释：由于容器内元素浮动，脱离了文档流，所以就只有外面容器的 2px 边距高度。

**「即方块没有如期将容器撑开」**


对于上面的例子，如果想内容在设置了浮动的条件下，还能撑开容器高度，那么可以通过触发容器的 BFC，那么容器将会包裹着浮动元素。

```html
<style>
  .container {
    width: 500px;
    border: 1px solid #000;
    overflow: hidden;
  }
  
  .child {
    float: left;
    width: 100px;
    height: 100px;
    background: pink;
  }
</style>

<div class="container">
  <div class="child"></div>
</div>
```


**「3. BFC 可以阻止元素被浮动元素覆盖」**

```html
<style>
  .left {
    height: 100px;
    width: 100px;
    float: left;
    background: green;
  }

  .none {
    width: 200px;
    height: 200px;
    background: grey;
  }
</style>
    
<div class="left">左浮动元素</div>
<div class="none">
  没设置浮动的元素，我的文字由于环绕没被覆盖，但实际方块被覆盖了
</div>
```


这时候其实第二个元素有部分被浮动元素所覆盖，(但是文本信息不会被浮动元素所覆盖) 如果想避免元素被覆盖，可以触发第二个元素的 BFC 特性，在第二个元素中加入 **「overflow: hidden」**，就会变成：


```html
<style>
  .left {
    height: 100px;
    width: 100px;
    float: left;
    background: green;
  }

  .none {
    width: 200px;
    height: 200px;
    background: grey;
    overflow: hidden;
  }
</style>
    
<div class="left">左浮动元素</div>
<div class="none">
  没设置浮动的元素，我的文字由于环绕没被覆盖，但实际方块被覆盖了
</div>
```





本文讲了BFC的概念是什么； BFC的约束规则；咋样才能触发生成新的BFC；BFC在布局中的应用：防止margin重叠(塌陷,以最大的为准)； 清除内部浮动；自适应两（多）栏布局。

### 1. BFC是什么？

`Block fomatting context` = `block-level box` + `Formatting Context`

#### Box: 

　　Box即盒子模型；

+ block-level box即块级元素

> display属性为block, list-item, table的元素，会生成block-level box；并且参与 `block fomatting context`；


+ inline-level box即行内元素

> display 属性为 inline, inline-block, inline-table的元素，会生成inline-level box。并且参与 `inline formatting context`；



#### Formatting context

　　Formatting context是W3C CSS2.1规范中的一个概念。它是页面中的一块渲染区域，并且有一套渲染规则，它决定了其子元素将如何定位，以及和其他元素的关系、相互作用。最常见的 Formatting context 有 Block fomatting context (简称BFC)和 Inline formatting context(简称IFC)。

　　CSS2.1 中只有BFC和IFC, CSS3中还增加了G（grid）FC和F(flex)FC。
　　
#### BFC 定义

　　BFC(Block formatting context)直译为"块级格式化上下文"。它是一个独立的渲染区域，只有Block-level box参与， 它规定了内部的Block-level Box如何布局，并且与这个区域外部毫不相干。
　　　

### 2.BFC的生成

只要元素满足下面任一条件即可触发 BFC 特性：

- body 根元素就是一个 BFC
    
- 浮动元素：float 除 none 以外的值
    
- 绝对定位元素：position (absolute、fixed)
    
- display 为 inline-block、table-cell、flex
    
- overflow 除了 visible 以外的值 (hidden、auto、scroll)




或者说


+ 根元素
+ float的值不为none
+ overflow的值不为visible
+ display的值为inline-block、table-cell、table-caption

> display：table也认为可以生成BFC，其实这里的主要原因在于Table会默认生成一个匿名的table-cell，正是这个匿名的table-cell生成了BFC

+ position的值为absolute或fixed


### 3. BFC的约束规则

+ 内部的Box会在垂直方向上一个接一个的放置
+ 垂直方向上的距离由margin决定。（完整的说法是：属于同一个BFC的两个相邻Box的margin会发生重叠（塌陷），与方向无关。）
+ 每个元素的左外边距与包含块的左边界相接触（从左向右），即使浮动元素也是如此。（这说明BFC中子元素不会超出他的包含块，而position为absolute的元素可以超出他的包含块边界）
+ BFC的区域不会与float的元素区域重叠
+ 计算BFC的高度时，浮动子元素也参与计算
+ BFC就是页面上的一个隔离的独立容器，容器里面的子元素不会影响到外面元素，反之亦然


看到以上的几条约束，想想我们学习css时的几条规则

+ Block元素会扩展到与父元素同宽，所以block元素会垂直排列
+ 垂直方向上的两个相邻DIV的margin会重叠，而水平方向不会(此规则并不完全正确)
+ 浮动元素会尽量接近往左上方（或右上方）
+ 为父元素设置overflow：hidden或浮动父元素，则会包含浮动元素


### 4. BFC在布局中的应用

#### 4.1 　防止margin重叠（塌陷）：


##### 两个相邻Box垂直方向margin重叠

```
<style>
    p {
        color: #f55;
        background: #fcc;
        width: 200px;
        line-height: 100px;
        text-align:center;
        margin: 100px;
    }
</style>
<body>
    <p>Haha</p>
    <p>Hehe</p>
</body>

```
　　
![图片](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/1.png)


两个p之间的距离为100px，发送了margin重叠（塌陷），以最大的为准，如果第一个P的margin为80的话，两个P之间的距离还是100，以最大的为准。

根据BFC布局规则第二条：

```

Box垂直方向的距离由margin决定。属于同一个BFC(上例中是body根元素的BFC)的两个相邻Box的margin会发生重叠

```

我们可以在p外面包裹一层容器，并触发该容器生成一个新BFC。那么两个P便不属于同一个BFC，就不会发生margin重叠了。

代码：

```
<style>
    .wrap {
        overflow: hidden;// 新的BFC
    }
    p {
        color: #f55;
        background: #fcc;
        width: 200px;
        line-height: 100px;
        text-align:center;
        margin: 100px;
    }
</style>
<body>
    <p>Haha</p>
    <div class="wrap">
        <p>Hehe</p>
    </div>
</body>

```


![图片](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/2.png)


##### 相邻Box水平方向margin重叠


```
<!doctype HTML>
<html>
<head>
<style type="text/css">

    #green {
        margin:10px 10px 10px 10px
    }
    #blue {
        margin:10px 10px 10px 10px
    }
    #red {
        margin:10px 10px 10px 10px
    }
    body {
        writing-mode:tb-rl;
    }

</style>
</head>
<body>

<div id="green" style="background:lightgreen;height:100px;width:100px;"></div>
<div id="blue" style="background:lightblue;height:100px;width:100px;"></div>
<div id="red" style="background:pink;height:100px;width:100px;"></div>

</body>
</html>
```

可以看到水平方向的margin发生了重叠。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/3.jpg)

我们可以给div加个`display:inline-block`，触每个div容器生成一个BFC。那么三个DIV便不属于同一个BFC（这个只body根元素形成的BFC），就不会发生margin重叠了。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/4.png)


##### 嵌套元素的margin重叠

```
<!DOCTYPE html>
<html>  
<head> 
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <!--The viewport meta tag is used to improve the presentation and behavior of the samples 
    on iOS devices-->
  <meta name="viewport" content="initial-scale=1, maximum-scale=1,user-scalable=no"/>
  <title></title>

  <style> 
    html, body { height: 100%; width: 100%; margin: 0; padding: 0; }
    #map{
      padding:0;
    }
    .first{
      margin:20px;
      background:lightgreen;
      width:100px;
      height:100px;
    }
    ul{
      /*display:inline-block;*/
      margin:10px;
      background:lightblue;
    }
    li{
      margin:25px;
    }
  </style> 
  
  
</head> 

<body class="claro"> 
  <div class="first"></div>
  <ul>
    <li>1</li>
    <li>2</li>
    <li>3</li>
  </ul>
</body> 

</html>
```

此时div与ul之间的垂直距离，取div、ul、li三者之间的最大外边距。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/5.png)

要阻止嵌套元素的margin重叠，只需让ul生成BFC即可（将上例中的注释去掉），这样div、ul、li之间便不会发生重叠现象。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/6.png)


而li位于同一BFC内所以仍然存在重叠现象。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/7.png)

给li设置line-block重新生成一个bfc就不存在重叠现象了。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/8.png)

需要注意的是：

如果为ul设置了border或padding，那元素的margin便会被包含在父元素的盒式模型内，不会与外部div重叠。

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/9.png)

《CSS权威指南》中提到块级正常流元素的高度设置为auto，而且只有块级子元素，其默认高度将是从最高块级子元素的外边框边界到最低块级子元素外边框边界之间的距离。如果块级元素右上内边距或下内边距，或者有上边框或下边框，其高度是从其最高子元素的上外边距边界到其最低子元素的下外边距边界之间的距离。


#### 4.2 清除内部浮动


```
<style>
    .par {
        border: 5px solid #fcc;
        width: 300px;
    }
 
    .child {
        border: 5px solid #f66;
        width:100px;
        height: 100px;
        float: left;
    }
</style>
<body>
    <div class="par">
        <div class="child"></div>
        <div class="child"></div>
    </div>
</body>
```

页面：

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/10.png)

根据BFC布局规则第六条：

```
计算BFC的高度时，浮动元素也参与计算

```
为达到清除内部浮动，我们可以触发par生成BFC，那么par在计算高度时，par内部的浮动元素child也会参与计算。

```
.par {
    overflow: hidden;
}
```

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/11.png)


#### 4.3 自适应多栏布局的


##### 4.3.1 自适应两栏布局

```
<style>
    body {
        width: 300px;
        position: relative;
    }
 
    .aside {
        width: 100px;
        height: 150px;
        float: left;
        background: #f66;
    }
 
    .main {
        height: 200px;
        background: #fcc;
    }
</style>
<body>
    <div class="aside"></div>
    <div class="main"></div>
</body>
```

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/13.png)

根据BFC布局规则第3条：

```
每个元素的margin box的左边， 与包含块border box的左边相接触(对于从左往右的格式化，否则相反)。即使存在浮动也是如此。
```
　　
　　
因此，虽然存在浮动的元素aslide，但main的左边依然会与包含块的左边相接触。

根据BFC布局规则第四条：

```
BFC的区域不会与float box重叠。

```

我们可以通过通过触发main生成BFC， 来实现自适应两栏布局。

```
.main {
    overflow: hidden;
}

```
当触发main生成BFC后，这个新的BFC不会与浮动的aside重叠。因此会根据包含块的宽度，和aside的宽度，自动变窄。效果如下：

![](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/14.png)



##### 4.3.2 自适应两栏布局

```
<!DOCTYPE html>
<html>  
<head> 
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <!--The viewport meta tag is used to improve the presentation and behavior of the samples 
    on iOS devices-->
  <meta name="viewport" content="initial-scale=1, maximum-scale=1,user-scalable=no"/>
  <title></title>

  <style> 
    html, body { height: 100%; width: 100%; margin: 0; padding: 0; }
    .left{
      background:pink;
      float: left;
      width:180px;
    }
    .center{
      background:lightyellow;
      overflow:hidden;
      
    }
    .right{
      background: lightblue;
      width:180px;
      float:right;
    }
  </style> 
  
  
</head> 

<body class="claro"> 
  <div class="container">
    <div class="left">
      <pre>
  .left{
    background:pink;
    float: left;
    width:180px;
  }
      </pre>
    </div>
    <div class="right">
       <pre>
  .right{
    background:lightblue;
    width:180px;
    float:right;
  }
      </pre>
    </div>
    <div class="center">
    <pre>
  .center{
    background:lightyellow;
    overflow:hidden;
    height:116px;
  }
      </pre>
    </div>
  </div>

</html>
```
这种布局的特点在于左右两栏宽度固定，中间栏可以根据浏览器宽度自适应。

![图片](https://github.com/zuopf769/notebook/blob/master/fe/BFC%E5%8E%9F%E7%90%86%E5%89%96%E6%9E%90/12.png)




### 4. 总结

其实以上的几个例子都体现了BFC布局规则第五条：

```
BFC就是页面上的一个隔离的独立容器，容器里面的子元素不会影响到外面的元素。反之也如此。
```

因为BFC内部的元素和外部的元素绝对不会互相影响，因此， 

当BFC外部存在浮动时，它不应该影响BFC内部Box的布局，BFC会通过变窄，而不与浮动有重叠。
同样的，当BFC内部有浮动时，为了不影响外部元素的布局，BFC计算高度时会包括浮动的高度。
避免margin重叠也是这样的一个道理。
　　
### 5. 延申阅读

+ [CSS清浮动处理（Clear与BFC）](http://www.cnblogs.com/dolphinX/p/3508869.html)

+ [A new micro clearfix hack](http://nicolasgallagher.com/micro-clearfix-hack/)

+ [清除浮动和BFC](https://github.com/zuopf769/notebook/blob/master/fe/%E6%B8%85%E9%99%A4%E6%B5%AE%E5%8A%A8%E5%92%8CBFC/README.md)