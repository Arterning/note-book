# tailwind-css

install

```
npm install -D tailwindcss
npx tailwindcssinit
```

tailwind 实际上是一个 postcss 插件

在入口 css 里加上这三行代码：

```css
@tailwind base
@tailwind components
@tailwind utilities
```

css 没有模块作用域，所以可能你在这里加了一个样式，结果别的地方样式错乱了。

而用原子 class 就没这种问题，因为样式是只是作用在某个 html 标签的。

我觉得光这三点好处就能够说服我用它了，特别是不用起 class 名字这点。

**但是还要每次去查文档哪些 class 对应什么样式呀**

这个可以用 tailwind css 提供的 vscode 插件来解决

### 自定义样式

由于 tailwind 的是低层次框架，现有的 tailwindcss 默认的 utility，base，component 不足以满足所有的场景；

而使用 @layer 指令，Tailwind 将自动将这些样式移动到 `@tailwind base`， `@tailwind utility`，`@tailwind component` 的位置

比如我们要做一个按钮组件，提炼对应的样式：

```css
@layer components {
  .btn-blue {
    @apply bg-blue-500 text-white font-bold py-2 px-4 rounded;
  }
  .btn-blue:hover {
    @apply bg-blue-700;
  }
}

```

[flex-basis](tailwind-css%20686b9a2bce0a4b4c896aa8419a969c6f/flex-basis%206837323c4dca408f91a9d853f83b48e1.md)

[grid-basis](tailwind-css%20686b9a2bce0a4b4c896aa8419a969c6f/grid-basis%2043f4a3da7262478188eeb656274965ee.md)

[background](tailwind-css%20686b9a2bce0a4b4c896aa8419a969c6f/background%205da183faebf04e0ca378e8dedc0afa57.md)

[Spacing](tailwind-css%20686b9a2bce0a4b4c896aa8419a969c6f/Spacing%2027be355eef9b42c4a937263f9219ad6f.md)