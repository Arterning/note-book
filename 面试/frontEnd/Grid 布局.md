

Flex 可以指定 纵向或者横向排列 但是不能指定 几行 几列  Grid 就可以



设置3列 每一列的宽度是设置的数字
```css

.container {
  display: grid;
  grid-template-columns: 400px 400px 400px;
}
```


```css
.container {
  display: grid;
  grid-template-columns: 400px 400px 400px;
  grid-template-rows: 100px 200px;
}
```


还可以用名字 而不是px定义大小
```css
.container{
    grid-template-areas: "header header header header"
                         "main main . sidebar"
                         "footer footer footer footer";
}

```