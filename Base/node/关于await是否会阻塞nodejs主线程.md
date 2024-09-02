  

答案是不会
  

像这样的代码：

  

```javascript

await foo();

console.log(“hello”);

```

  

类似于

  

```javascript

foo().then(() => {

    console.log(“hello”);

});

```

  

因此,等待只是将该范围中的以下代码放入一个不可见的.then()处理程序中,其他所有内容的工作方式与使用.then()处理程序实际编写的内容完全相同.

  
  

## Refer

- [article](https://www.zhihu.com/tardis/zm/art/61807318?source_id=1003)