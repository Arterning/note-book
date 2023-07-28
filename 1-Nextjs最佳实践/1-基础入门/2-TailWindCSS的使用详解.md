---
aliases: [2-TailWindCSS的使用详解]
tags: []
cssclass:
source:
date created: 星期四, 五月 25日 2023, 4:54:35 凌晨
date modified: 星期五, 六月 2日 2023, 2:24:13 下午
---
## 安装与配置

关于简单的安装与配置我们[[1-Typescript+Eslint+Prettier搭建Nestjs工程及断点调试|前一节课]]已经有所讲解，重复的内容不再多讲，可自行参考[[1-Typescript+Eslint+Prettier搭建Nestjs工程及断点调试|前一节课]]前一节课

### 内容扫描

在配置文件的`content`中指定需要扫描的文件,不在`content`中的文件将无法使用`tailwind`.编译时框架会自动跟`content`中的文件使用了哪些`tailwind`中的类(包括`@layer`中自定义的)来排除其它没有使用的.其中的路径使用[glob规则](https://en.wikipedia.org/wiki/Glob_(programming))进行匹配.

**注意: 千万不要去扫描任何css文件甚至是包含tailwind类的css文件,否则系统直接爆**

一般情况下针对`react`使用如下配置即可(如果是`vue`,请加上`.vue`)

```js
// tailwind.config.js
content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
```

为了能扫描到,在使用动态类名时请使用完整类名,如下

```jsx
<div class="{{ error ? 'text-red-600' : 'text-green-600' }}"></div>

// 下面这样是错误的
<div class="text-{{ error ? 'red' : 'green' }}-600"></div>
```

如果在某些情况下实在需要使用不完整类名来拼装(比如用户动态发表的内容),请把它加入安全列表

```js
// tailwind.config.js
module.exports = {
  safelist: [
    'text-red',
    'text-green',
    {
      // 使用正则匹配多个类
      pattern: /bg-(red|green|blue)-(100|200|300)/,
      // 同时把匹配到的类的修饰符类也加上
      variants: ['hover', 'focus'],
    },
  ]
  // ...
}
```

另外关于**转换源文件**(比如markdown需要转成html才能读取其中的类)以及**自定义抽取逻辑**这部分的内容一般用不到,如有需要自行看[官网文档](https://tailwindcss.com/docs/content-configuration#transforming-source-files)即可

### 主题配置

在主题配置中可以自定义色系、类型比例、字体、断点、边框半径值等

主题配置有两个概念组成

- 全局配置包含`screens`,`colors`,`spacing`三个对象
- 核心插件指所有的工具类

全局配置的属性如下

- `screens`用于定义全局屏幕断点
- `colors`用于定于全局色系,默认情况下它由`backgroundColor`,`textColor`这些核心插件决定
- `spaces`用于定义全局间距,默认情况下它由 `padding`, `margin`这些核心插件决定

有两种方式类配置主题

- 直接通过在`theme`对象中定义一个属性的对象,这样会完全替代掉这个属性的默认配置
- 通过 `extends`来新增或修改一个值,而不是完全替代整个属性

对于通过读取一个属性的值来配置另一个属性有两种方法

- 使用`theme`函数来读取当前配置的值
- 读取默认配置的值

注意点

- 所有核心插件均支持`DEFAULT`属性,用于没有后缀的默认值,比如`className="border-radius"`
- 不要在`theme`中配置空对象,如果要禁用某个核心插件,在`corePlugins`把它配置成`false`即可
- 所有默认配置[在此查看](https://github.com/tailwindlabs/tailwindcss/blob/v1/stubs/defaultConfig.stub.js#L5)

例

> 注意：以下为演示而用，不在本节代码中

```js
// tailwind.config.js
const defaultTheme = require('tailwindcss/defaultTheme')
module.exports = {
  // ...
  // 禁用的核心插件
  corePlugins: {
    opacity: false,
  },
  theme: {
    // 基础配置
    // 替代默认的spacing
    spacing: { ... }
    // 核心插件
    // 替代默认的borderRadius
    borderRadius: { ... },
    // 通过theme函数读取最新的当前配置中的spacing
    backgroundSize: ({ theme }) => ({
      auto: 'auto',
      cover: 'cover',
      contain: 'contain',
      ...theme('spacing')
    }),
    extends: {
       // 新增一个screen
      screens: {
        '3xl': '1600px',
      },
      // 通过默认配置的值来修改默认的sans
      fontFamily: {
        sans: [
          'Lato',
          ...defaultTheme.fontFamily.sans,
        ]
      }
    }
  }
}
```

本节课的配置如下

```javascript
// tailwind.config.js
/** @type {import('tailwindcss').Config} */
module.exports = {
    prefix: 'tw-',
    darkMode: ['class', '[data-theme="dark"]'],
    content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
    theme: {
        screens: {
            xs: '480px',
            sm: '576px',
            md: '768px',
            lg: '992px',
            xl: '1200px',
            '2xl': '1400px',
        },
        extend: {
            fontFamily: {
                standard: 'var(--font-family-standard)',
                firacode: 'var(--font-family-firacode)',
                kaiti: 'var(--font-family-kaiti)',
            },
        },
    },
    corePlugins: { preflight: false },
    plugins: [],
};
```

### 配置文件

使用`pnpx tailwindcss init`默认会创建一个名为`tailwind.config.js`的配置文件,同时会生成一个`postcss.config.js`.

如果需要自定义配置文件名称,比如使用`tailwind.js`作为配置文件,需要到`postcss.config.js`中修改一下

```js
module.exports = {
  plugins: {
    tailwindcss: { config: './tailwindcss.js' },
  },
}
```

如果想生成包含默认配置的一个配置文件,使用命令`pnpx tailwindcss init --full`初始化

### 扩展样式

在使用TW时，我们往往需要定义一些全局的样式来覆盖和新增一些全局的Tailwind类，以方便再全局状态下随处可用

官方文档默认使用`@tailwind`指令来导入其核心文件,这样做就无法按顺序来导入自定义的css文件,因为`@import`必须总是在css文件最前面,像下面这样是错误的

```css
@tailwind base;
@import "./custom-base-styles.css";

@tailwind components;
@import "./custom-components.css";

@tailwind utilities;
@import "./custom-utilities.css";
```

但是如果像下这样定义,则会使用`custom-base-styles.css`中在`@layer`外定义的样式同时会覆盖掉框架内部的`components`和`utilities`中的样式

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
@import "./custom-base-styles.css";
```

所以我们应该像下面这样定义

```css
@import 'tailwindcss/base';
@import './tailwind/base';
@import 'tailwindcss/components';
@import './tailwind/components';
@import 'tailwindcss/utilities';
@import './tailwind/utilities';
```

添加各个层的css文件

```css
/* src/styles/tailwind/base.css */
@layer base {
}
/* src/styles/tailwind/components.css */
@layer components {
}
/* src/styles/tailwind/components.css */
@layer utilities {
}
```

在`@layer`中定义的样式如果在应用中没有用到则编译后被删除,如果要一直存在可以被组件动态使用的样式,请定义在`@layer`指令外部.

同时安装上面的导入顺序,现在的样式覆盖顺序如下

```
utilities.css[外部]`->`utilities.css[@layer]`->`tailwindcss/utilities`->`components.css[外部]`->`components.css[@layer]`->`tailwindcss/components`->`base.css[外部]`->`base.css[@layer]`->`tailwindcss/base
```

这就是我们想要的!

其中

- `@layer base`: 添加或覆盖基础样式
- `@layer components`: 添加或覆盖组件类样式
- `@layer utilities`: 添加或覆盖工具类样式

例如，用于在`base`中修改一些框架内部类

```css
/* src/styles/tailwind/base.css */
@layer base {
    html {
        font-family: var(--font-family-standard);
        font-size: var(--font-size-base);
    }

    h1 {
        @apply tw-text-3xl;
    }

    h2 {
        @apply tw-text-2xl;
    }

    h3 {
        @apply tw-text-xl;
    }

    h4 {
        @apply tw-text-lg;
    }

    h5 {
        @apply tw-text-sm;
    }

    h6 {
        @apply tw-text-xs;
    }
}

```

基于以上，所以本节课的CSS样例如下

```css
/* src/styles/index.css */
@import './vars.css';
@import 'tailwindcss/base';
@import './tailwind/base.css';
@import 'tailwindcss/components';
@import './tailwind/components.css';
@import 'tailwindcss/utilities';
@import './tailwind/utilities.css';
@import './app.css';
```

## 整合React

### 组件中直接使用

要在组件中直接使用，请可以把TW的类写在`className`，例如

```tsx
// src/App.tsx
const App: FC = () => {
    return (
        <Container>
            <div className="tw-shadow-md tw-p-5 tw-bg-black tw-text-center tw-text-white tw-text-lg">
                <h1>
                    欢迎来到3R教室，这是
                    <span className="tw-text-white tw-cursor-pointer">React课程第一节</span>
                </h1>
            </div>
        </Container>
    );
};
const Container: FC<{ children?: ReactNode }> = ({ children }) => {
    return (
        <ConfigProvider
            ...
        >
            <StyleProvider hashPriority="high">
                <div
                    className='tw-w-full
     tw-bg-[image:url("https://img.pincman.com/media/202305242046398.png")]
     tw-bg-cover
     tw-bg-no-repeat
     tw-bg-center
     tw-flex 
     tw-items-center 
     tw-justify-center
     tw-flex-col'
                >
                    {children}
                </div>
            </StyleProvider>
        </ConfigProvider>
    );
};

export default App;
```

### 样式复用

要在CSS Modules中使用使用TW，我们需要用到`@apply`指令。该指令用于在自定义样式中组合一些已经存在的类并关联形成一个行的类,不仅可以在`@layer`中使用,可以单独拿到外部类使用，比如在**"css modules"**中

```css
/* src/App.module.css */
.app {
    @apply tw-w-full
     tw-bg-[image:url("https://img.pincman.com/media/202305242046398.png")]
     tw-bg-cover
     tw-bg-no-repeat
     tw-bg-center
     tw-flex 
     tw-items-center 
     tw-justify-center
     tw-flex-col;

    & > .container {
        @apply tw-shadow-md tw-p-5 tw-bg-black tw-text-center tw-text-white tw-text-lg;
    }

    & > .test {
        @apply tw-w-[20rem] tw-h-[10rem];
    }
}
```

在组件中导入`App.module.css`以使用
> 其中`.module.css`的后缀是vite对css modules的约定命名方式

```tsx
// src/App.tsx
import $styles from './App.module.css';

const App: FC = () => {
    return (
        <Container>
            <div className={$styles.container}>
                <h1>
                    欢迎来到3R教室，这是
                </h1>
            </div>
        </Container>
    );
};
const Container: FC<{ children?: ReactNode }> = ({ children }) => {
    return (
        <ConfigProvider
            ...
        >
            <StyleProvider hashPriority="high">
                <div className={$styles.app}>{children}</div>
            </StyleProvider>
        </ConfigProvider>
    );
};

export default App;
```

### 添加扩展类

尝试通过`@layer `来扩展TW的类，这些类可以在`TSX`组件中使用

> 需要注意的是 我们不能在CSS Modules或者全局等其它CSS文件中使用`@apply`来应用扩展的类，因为这会导致无限嵌套

```css
/* src/styles/tailwind/components.css */
@layer components {
    .tw-panel {
        /* 一个毛玻璃块 */
        @apply tw-mt-6
                    tw-bg-zinc-300/25
                    tw-shadow-slate-700/20
                    tw-shadow-md
                    hover:tw-shadow-lg
                    hover:tw-shadow-slate-700/20
                    tw-transition-shadow
                    tw-rounded-lg
                    tw-backdrop-blur-sm
                    tw-flex
                    tw-items-center
                    tw-justify-center
                    tw-text-black/75;
    }
}

/* src/styles/tailwind/utilities.css */
@layer utilities {
    /* 鼠标移动到文字出现渐进式下划线 */
    .tw-animate-decoration {
        background: linear-gradient(red, blue) 0% 100% / 0% 1px no-repeat;
        transition: background-size ease-out 200ms;

        &:not(:focus):hover {
            background-size: 100% 1px;
        }
    }
}
```

在`App.module.css`添加一个`test`类，用于设置这个panel的宽度和高度

```css
.app {
    & > .test {
        @apply tw-w-[20rem] tw-h-[10rem];
    }
}
```

现在，我们就可以在TSX中使用`tw-panel`了

```tsx
const App: FC = () => {
    return (
        <Container>
             <div className={clsx($styles.container)}>
                <h1>
                    欢迎来到3R教室，这是
                    <span className="tw-animate-decoration tw-cursor-pointer">React课程第一节</span>
                </h1>
            </div>
            <div className={clsx($styles.test, 'tw-panel')}>测试</div>
        </Container>
    );
};
```

### 暗黑模式

你可以使用`class`来手动切换，但是我们也可以使用其它标签来切换，比如`[data-theme="dark"]`
在`tailwind.config.js`中添加`[data-theme="dark"]`切换

```javascript
module.exports = {
    prefix: 'tw-',
    darkMode: ['class', '[data-theme="dark"]'],
    ...
};
```

添加之后，每次当`<html>`标签中有`data-theme="dark"`时就会变成暗黑模式

接着，我们在React中添加一个切换暗黑模式的按钮

> 本节只关注tailwind，所以此处没有切换Antd的暗黑模式，这个在下一节课实现

```tsx
const Container: FC<{ children?: ReactNode }> = ({ children }) => {
    const [dark, setDark] = useState(true);
    const toggleDark = useCallback(() => setDark((state) => !state), []);
    useEffect(() => {
        const html = document.getElementsByTagName('html')[0];
        if (dark) {
            html.setAttribute('data-theme', 'tw-dark');
        } else {
            html.setAttribute('data-theme', 'tw-light');
        }
    }, [dark]);
    return (
        <ConfigProvider
           ...
        >
            <StyleProvider hashPriority="high">
                <div className={clsx($styles.app)}>
                    <Button onClick={toggleDark}>切换暗黑模式</Button>
                    {children}
                </div>
            </StyleProvider>
        </ConfigProvider>
    );
};

export default App;
```

下面我们为需要有暗黑模式的样式加上暗黑模式

有两种方法

第一种，是应用`html[*data-theme*='tw-dark']`来添加暗黑模式

```css
.app {
    ...
}

html[data-theme='tw-dark'] {
    .app {
        @apply tw-bg-[image:url("https://img.pincman.com/media/202305242029933.png")];
    }

    .test {
        @apply tw-bg-black/20 
            tw-text-white/75   
             tw-shadow-slate-200/20
             hover:tw-shadow-slate-200/20;
    }
}

```

第二种，是在tailwind的类前添加`dark`前缀

```css
@layer components {
    .tw-panel {
        @apply tw-mt-6
                    tw-bg-zinc-300/25
                    tw-shadow-slate-700/20
                    tw-shadow-md
                    hover:tw-shadow-lg
                    hover:tw-shadow-slate-700/20
                    tw-transition-shadow
                    tw-rounded-lg
                    tw-backdrop-blur-sm
                    tw-flex
                    tw-items-center
                    tw-justify-center
                    tw-text-black/75
                    dark:tw-bg-black/20 
           dark:tw-text-white/75   
             dark:tw-shadow-slate-200/20
             dark:hover:tw-shadow-slate-200/20;
    }
}
```