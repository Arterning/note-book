
当元素的 `position` 属性值为 `static` 时，`top`、`right`、`bottom` 和 `left` 这些定位属性是不起作用的。这是因为 `static` 定位是元素的默认定位方式，元素在正常文档流中，并且不受这些定位属性的影响。

如果你想使用 `top`、`right`、`bottom` 和 `left` 来控制元素的位置，你需要将元素的 `position` 属性设置为其他值，比如 `relative`、`absolute` 或 `fixed`。


block盒子的默认posistion是什么

`block` 盒子（块级盒子）的默认 `position` 属性值是 `static`。在 CSS 中，`position: static;` 是元素的默认定位方式。

当元素的 `position` 属性设置为 `static` 时，元素遵循正常的文档流，并根据在 HTML 中的布局顺序进行排列。这意味着元素会按照其在 HTML 中的出现顺序一个接一个地放置，并且它们之间会有换行，不会受到 `top`、`right`、`bottom`、`left` 和 `z-index` 属性的影响。

以下是一个简单的例子，演示了 `position: static;` 的默认行为：

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <style>
    .block {
      width: 100px;
      height: 100px;
      background-color: lightblue;
      margin: 10px;
    }
  </style>
  <title>Static Position</title>
</head>
<body>
  <div class="block">1</div>
  <div class="block">2</div>
  <div class="block">3</div>
</body>
</html>
```

在这个例子中，三个块级盒子按照它们在 HTML 中的顺序依次排列，因为它们的 `position` 属性默认为 `static`。