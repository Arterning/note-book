
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