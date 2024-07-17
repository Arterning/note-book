

# 大小类

```
w-{size}
h-{size}
max-w-{size}
max-h-{size}
min-w-{size}
min-h-{size}
```


# 背景类
```
bg-{color} bg-gray-300
bg-{image} 
bg-{position} 比如bg-center
bg-{size} 设置背景尺寸 比如bg-auto表示使用原始图片大小


白天黑夜切换
bg-white dark:bg-[#313338]

删除按钮用红色
bg-rose-500 text-white
```


# 文本类

```
text-{color}
text-{size}
font-{family}
font-{weight}
leading-{size} 行间距

灰色文字
text-zinc-500

粗体
font-semibold

白天黑夜切换
text-zinc-500 dark:text-zinc-400

骚气的紫色
text-indigo-500 dark:text-indigo-400
```


# 宽高
```
h-4 w-4


可爱的正方形
aspect-square rounded-xl
```


## 内边距和外边距
```
ml-auto
mr-auto
mx-auto

```





# 边框
```
borer-{color}
border-{size}
border-{slide} border-l 表示只再左侧添加边框
rounded-size 圆角
```



# 过渡
```
transition  hover:opacity-75 ease-in-out
```

## 动画
```
animate-spin 转圈圈
animate-ping 跳动
animate-pluse 脉冲
animate-bounce 弹跳
```



# 布局
```
横排 居中排列
flex items-center justify-center gap-x-2
flex flex-row items-center justify-between space-y-0 pb-2




竖着来
flex flex-col items-center justify-center h-full gap-x-2



如果有超出的内容 那么显示滚动条
overflow-x-auto
overflow-y-auto
```




# 响应式
```
hidden md:flex


breakpoint:
sm:
md:
lg:
```





# 其他
```
focus-visible:ring-0 focus-visible:ring-offset-0

object-cover

<Image className="fill"
```


`focus-visible:ring-0` 和 `focus-visible:ring-offset-0` 是指定在用户使用键盘导航时，焦点（focus）样式的一种方式，通常在网页开发中使用。这两个样式是针对键盘焦点的，而不是鼠标点击导致的焦点。

- `focus-visible:ring-0`: 这表示当用户使用键盘导航时，不应用焦点环（ring）。焦点环是一个可见的样式，通常是一个轻微的边框或阴影，用于表示当前具有焦点的元素。通过将 `ring-0` 应用于 `focus-visible`，你可以在键盘焦点时移除默认的焦点环。

- `focus-visible:ring-offset-0`: 这表示在键盘焦点时，不应用焦点环的同时，也不应用焦点环偏移（ring-offset）。`ring-offset` 是指焦点环相对于元素的偏移量。通过将 `ring-offset-0` 应用于 `focus-visible`，你可以在键盘焦点时同时移除焦点环和其偏移。

这些样式通常用于提高可访问性，以确保在用户使用键盘浏览网页时，焦点样式不会引起混淆或干扰。使用 `focus-visible` 可以确保只有在键盘导航时应用焦点样式，而不会在使用鼠标点击时产生。这有助于保持一致的用户体验，同时提供良好的可访问性。



`object-fit: cover` 是一个 CSS 属性，用于指定一个可替换元素（比如 `<img>` 或 `<video>`）的内容在其容器中如何适应。具体来说，`cover` 值是其中的一种，它的作用是让元素的内容等比例缩放，以填充整个容器，可能会导致元素的部分内容被裁剪。

例如，如果你有一个包含图片的容器，并且设置了 `object-fit: cover;`，那么图片会被缩放，直到宽度和高度都适应容器，这可能导致图片的一部分被裁剪，但保持了整个容器的填充。

```css
img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
```

在这个例子中，`object-fit: cover;` 使得图片在容器中被缩放并填充，确保容器完全被图片填满，同时保持了图片的宽高比。

这是一个用于处理图像的常见 CSS 技术，特别是在响应式设计中，你希望图像能够适应不同大小的容器而不失真。