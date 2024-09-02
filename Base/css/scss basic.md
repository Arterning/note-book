# SCSS

## 多属性叠写

```scss
/**
如果我们需要一次性指定 border-width border-style
可以直接嵌套写入
**/

#test {
    border: {
        style: solid;
        width: 10px;
        color:pink;
    }
}

//compiled css
#test {
    border-style: solid;
    border-width: 10px;
    border-color:pink;
}
```

## 变量的声明

```scss
//变量的声明

$demo-color: pink !default;

$demo-border : 1px solid $demo-color;

//demo-color后续可以改变 覆盖 如果没有覆盖就使用默认值
//$demo-color:blue;

.nav-border {
    border: $demo-border;
}

//compiled css result

.nav-border {
    border: 1px solid pink;
}
```

## 更强大的选择器

可以在嵌套层内使用多种选择器 比如

- 相邻兄弟选择器

> 子元素选择器 ~ 后继兄弟选择器

```scss
#strick {
    > h1{
        border: 1px solid pink;
    }

    + h2 {

    }

    ~ h3 {
        
    }
}
```

## 嵌套

使用sass可以轻松的把多个内容嵌套在一起 而不用像css一样需要分开写 这应该是最有用的特性

```scss
#content {
    div {
        h1 {
            border: 1px solid pink;
        }

        p {
            width:100px;
        }
    }
}

//compiled css
#content div h1 {
    border: 1px solid pink;
}

#content div p {
    width: 100px;
}
```

## mixin

可以把他当成一个大型的结构体 使用@include使用

```scss
@mixin test {
    color: pink;
}

.wrapper {
    @include test;
    width: 10px;
}

//接受参数的minxin
@mixin test-function($color, $hover, $visited){
    color:$color;
    &:hover{
        color: $hover
    };
    &:visited{
        color:$visited
    };
}

a {
    @include test-function(blue, red, yellow)
}

//compiled css
a {
    color : blue;
}

a:hover {
    color:red;
}

a:visited {
    color:yellow;
}
```

## 特殊符号

& 该符号后面的内容和父元素直接拼接 没有空格 适用于各种伪类

```scss
article a {
    color: gray;
    &:hover {
        transform: translateX(120%)
    }

    &:visited {
        color:pink;
    }
}

//compiled css
article a {
    color: gray;
}

article a:hover {
    transform: translateX(120%)
}

article a:visited {
    color:pink;
}
```

## 群组嵌套

content 和 子元素 一一配对