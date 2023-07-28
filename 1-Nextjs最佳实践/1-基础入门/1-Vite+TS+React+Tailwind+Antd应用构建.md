---
aliases: [1-Vite+TS+React+Tailwind+Antd应用构建, 1-Vite_Nextjs+TS+React+Tailwind+Antd应用构建]
tags: []
cssclass:
source:
date created: 星期三, 五月 24日 2023, 4:43:16 凌晨
date modified: 星期五, 六月 2日 2023, 2:26:17 下午
---
## 创建应用
创建vite+react-ts的项目非常简单，因为我使用的是pnpm,所以直接使用pnpm创建即可。对于node.js,pnpm这些环境的配置，请参看[Nestjs教程](https://www.yuque.com/pincman/3r/chapter1)就可以了。
使用以下命令来创建一个`vite`+`react`的项目

```bash
~ pnpm create vite
~ reactplus # 这里随便取一个应用名称，另外可能出现卡主的现象，但是你按下键盘上任意一个字符就会出现"Project name"，输入你要创建的项目文件夹名称即可，然后按下回车键
~ # 这一步选择React，按下回车键
~ # 这一步选择Typescript, 注意不要选择"Typescript-swc"，实测使用swc，后面会有问题，按下回车键
~ cd reactplus #进入应用目录
~ pnpm i # 安装依赖
~ pnpm dev # 启动应用，启动后会提示访问地址，点击打开地址链接
```

打开应用地址后是这样的
![](https://img.pincman.com/media/202304040307260.png#id=NZJAz&originHeight=1888&originWidth=2942&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
## 环境配置
### 使用Commonjs
默认vite使用了`ESM`模块，为了支持我们的eslint,stylint等配置文件，我们需要把它改成Commonjs模块
> 或者可以把`eslintrc.js`等文件改成`.cjs`后缀

只要把`package.json`中的`"type": "module",`删除即可
### Typescript
设置`allowJs`和`esModuleInterop`为`true`,并添加一些其它规则后`tsconfig.json`的配置如下

```json
// tsconfig.json
{
  "compilerOptions": {
    "target": "ESNext",
    "useDefineForClassFields": true,
    "lib": ["DOM", "DOM.Iterable", "ESNext"],
    "allowJs": true,
    "skipLibCheck": true,
    "esModuleInterop": true,
    "allowSyntheticDefaultImports": true,
    "strict": true,
    "forceConsistentCasingInFileNames": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx",
    "declaration": true,
    "removeComments": true,
    "experimentalDecorators": true,
    "alwaysStrict": true,
    "sourceMap": true,
    "incremental": true,
    "noUnusedLocals": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "pretty": true,
    "noImplicitAny": true,
    "importsNotUsedAsValues": "remove"
  },
  "include": ["src"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

### Prettier

```bash
~ pnpm add prettier -D
```

新增`.pretterrc.js`以添加prettier配置

```javascript
// .prettierrc.js
/** @format */
module.exports = {
    singleQuote: true,
    trailingComma: 'all',
    printWidth: 100,
    proseWrap: 'never',
    endOfLine: 'auto',
    semi: true,
    tabWidth: 4,
    vueIndentScriptAndStyle: true,
    htmlWhitespaceSensitivity: 'strict',
    overrides: [
        {
            files: '.prettierrc',
            options: {
                parser: 'json',
            },
        },
        {
            files: 'document.ejs',
            options: {
                parser: 'html',
            },
        },
    ],
};
```

复制代码中的`.prettierignore`到根目录来配置需要让prettier忽略的文件
### Eslint
我们使用airbnb的编码风格来编写代码，为了支持eslint+prettier来统一代码风格，需要安装以下依赖
其中`eslint-plugin-prettier`与`eslint-config-prettier`用于整合prettier，以防止规则冲突

```bash
~ pnpm add jest \
          eslint \
          eslint-config-airbnb \
          eslint-config-airbnb-typescript \
          eslint-config-prettier \
          eslint-plugin-import \
          eslint-plugin-jest \
          eslint-plugin-jsx-a11y \
          eslint-plugin-prettier \
          eslint-plugin-react \
          eslint-plugin-react-hooks \
          eslint-plugin-unused-imports \
          @typescript-eslint/eslint-plugin \
          @typescript-eslint/parser -D
```

一个`tsconfig.eslint.json`文件用于设置在eslint在格式化代码时需要额外包含的文件

```json
// tsconfig.eslint.json
{
  "extends": "./tsconfig.json",
  "include": [
    "./src",
    "./test",
    "./typings",
    "./build",
    "./mock",
    "**.js",
    "**.ts"
  ],
  "exclude": ["node_modules"]
}
```

新增`.eslintrc.js`用于设置eslint的配置。与Nestjs的Eslint配置大体相同，一些不同的地方是

- 在`parserOptions`中启用了`jsx`
- 在`env`中增加了`es6`和`browser`
- `extends`中把airbnb的`airbnb-base`和`airbnb-typescript/base`的`base`给去掉了，因为`base`是不支持react的不完整版
- 增加了`airbnb/hooks`，以支持`react`
- 同时在`ruls`中添加了一些常用的`react`与`jsx-a11y`的规则配置
> 其余规则的解释可自行参考[nestjs课程](https://www.yuque.com/pincman/3r/chapter1)

```javascript
// .eslintrc.js
module.exports = {
    parser: '@typescript-eslint/parser',
    parserOptions: {
        // 指定ESLint可以解析JSX语法
        ecmaVersion: 'latest',
        sourceType: 'module',
        project: './tsconfig.eslint.json',
        // React启用jsx
        ecmaFeatures: {
            jsx: true,
        },
    },
    env: {
        es6: true,
        browser: true,
        jest: true,
        node: true,
    },
    plugins: ['@typescript-eslint', 'jest', 'prettier', 'import', 'unused-imports'],
    extends: [
        'airbnb',
        'airbnb-typescript',
        'airbnb/hooks',
        'plugin:@typescript-eslint/recommended',
        'plugin:@typescript-eslint/recommended-requiring-type-checking',
        'plugin:jest/recommended',
        'prettier',
        'plugin:prettier/recommended',
    ],
    rules: {
      ...
```

复制代码中的`.eslintignore`到根目录来配置需要让eslint忽略的文件
### Stylelint
sylelint用于定制和统一css代码的风格
安装以下依赖

```bash
~ pnpm add stylelint \
          stylelint-config-css-modules \
          stylelint-config-recess-order \
          stylelint-config-standard \
          stylelint-prettier -D
```

新增`stylint.config.js`以配置stylelint
> 具体规则的设置，请自行查看代码

```javascript
// stylelint.config.js
module.exports = {
    extends: [
        'stylelint-config-standard',
        'stylelint-config-css-modules',
        'stylelint-config-recess-order',
        'stylelint-prettier/recommended',
    ],
    rules: {...
```

复制代码中的`.stylelintignore`到根目录来配置需要让stylint忽略的文件
### Vite
为了后续方便地编写其它vite配置，我们先建立一个专门用于vite构建配置的目录:`build`，然后在里面写vite配置
我们需要配置一下`tsconfig.node.json`，在`include`中添加`build`，以便ts能解析到`build`目录中的文件

```json
// tsconfig.node.json
{
    ...
    "include": ["vite.config.ts", "./build"]
}
```

添加一下`deepmerge`这个库用于深度合并对象

```bash
~ pnpm add deepmerge
```

编写一个用于获取某个目录的绝对路径的函数

```typescript
// build/utils.ts
export const pathResolve = (dir: string) => resolve(__dirname, '../', dir);
```

编写一个vite自定义配置生成函数的类型，`isBuild`为是否处于构建环境中（即生产环境）

```typescript
// build/types.ts
export type Configure = (params: ConfigEnv, isBuild: boolean) => UserConfig;
```

编写配置构建函数

- 该函数接收两个参数，`params`参数用于接收vite默认的环境参数
- `configure`参数用于接收一个自定义的配置生成函数，是可选的

其逻辑为先设置定一些预定义的配置，比如**路径别名**，然后如果有自定义的配置生成函数传入，则生成自定义的配置，然后没有，把自定义配置设置成控对象，最后合并并覆盖预定义配置，然后返回最终的配置
在此处我们添加了以下配置

- 路径别名，作用是可以使用`@/`为前缀来便捷的导入模块，所以需要分别配置vite和typescript
- 使用`camelCaseOnly`(即驼峰命名法：如`containerMain`)来定义CSS MODULES的类名
- 加载默认的`react`插件

```typescript
// build/index.ts
export const createConfig = (params: ConfigEnv, configure?: Configure): UserConfig => {
    const isBuild = params.command === 'build';
    return merge<UserConfig>(
        {
            resolve: {
                alias: {
                    '@': pathResolve('src'),
                },
            },
            css: {
                modules: {
                    localsConvention: 'camelCaseOnly',
                },
            },
            plugins: createPlugins(isBuild),
        },
        typeof configure === 'function' ? configure(params, isBuild) : {},
        {
            arrayMerge: (_d, s, _o) => Array.from(new Set([..._d, ...s])),
        },
    );
};
```

编写一个插件创建函数，用于放置所有的vite插件

```typescript
// build/plugins.ts
export function createPlugins(isBuild: boolean) {
    const vitePlugins: (PluginOption | PluginOption[])[] = [react()];
    return vitePlugins;
}
```

最后在`vite.config.ts`中构建配置

```typescript
// vite.config.ts
import { createConfig } from './build';

export default defineConfig((params: ConfigEnv): UserConfig => {
    const config = createConfig(params);
    return config;
});
```

最后我们为`tsconfig.json`添加上别名，根路径以及编译路径的配置

```json
// tsconfig.json
{
    "compilerOptions": {
        ...
        "outDir": "./dist",
        "baseUrl": "./",
        "paths": {
            "@/*": ["src/*"]
        }
    },
    "include": ["src"],
    "references": [{ "path": "./tsconfig.node.json" }]
}
```

### VSCode
> 请自行按照，eslint,prettier,stylint以及postcss插件

使用以下配置可自动对不满足elsint，stylint以及prettier预设规则的部分代码进行自动修复

- `editor.formatOnSave`关闭vscode的默认的自动代码格式化，`editor.codeActionsOnSave`把`ts/js`与`stylint`自动代码格式化分别托管给`eslint`和`stylint`以及通过它们调用`prettier`进行修复
- `xxx..validate: false`用于禁用`less`等所有样式语言的语法验证，把它们全部托管给`postcss`
- 使用`emmet.includeLanguages`把`postcss`的识别改为`css`，这样我们就可以在写`postcss`时获取`css`的提示了
- 把`less`,`scss`,`css`,`postcss`的所有样式语言的语法语法识别和验证托管给`stylelint`
- 其余的配置相信你自己能看懂，我就不赘述了

```typescript
// .vscode/settings.json
{
    "editor.formatOnSave": false,
    "editor.codeActionsOnSave": {
        "source.fixAll.eslint": true,
        "source.fixAll.stylelint": true
    },
    "css.validate": false,
    "less.validate": false,
    "scss.validate": false,
    "postcss.validate": false,
    "emmet.includeLanguages": {
        "postcss": "css"
    },
    "stylelint.snippet": ["css", "scss", "less", "postcss"],
    "stylelint.validate": ["css", "scss", "less", "postcss"],
    "javascript.preferences.importModuleSpecifier": "project-relative",
    "typescript.suggest.jsdoc.generateReturns": false,
    "eslint.packageManager": "pnpm",
    "stylelint.packageManager": "pnpm",
    "npm.packageManager": "pnpm"
}
```

## 样式与组件库
### 支持Tailwind
首先安装一下`tailwindCSS`的VSCode插件
安装依赖

```
~ pnpm add tailwindcss postcss autoprefixer postcss-import postcss-mixins postcss-nested postcss-nesting -D
```

生成配置

```bash
~ pnpx tailwindcss init -p
```

可以看到生成了两个配置文件，分别为`postcss.config.js`与`tailwind.config.js`
首先我们修改一下`postcss.config.js`，增加一些插件，它们的作用如下

- `postcss-import`: 可以让一个css文件导入其它css文件
- `postcss-nesting`: 可以编写`postcss-nesting`规范的嵌套css
- `tailwindcss/nesting`: 可以编写`scss`规范的嵌套css
- `autoprefixer`: 自动为css样式添加浏览器适配前缀
- `postcss-mixins`: 编写css样板代码，使一段css代码供多个地方使用

```javascript
// postcss.config.js
module.exports = {
    plugins: {
        'postcss-import': {},
        'postcss-nesting': {},
        'tailwindcss/nesting': {},
        tailwindcss: {},
        autoprefixer: {},
        'postcss-mixins': {},
    },
};
```

接下来修改一下`tailwind.config.js`

- 为所有tailwind的类添加`tw-`标识
- 手动切换暗黑模式
- 在`index.html`以及通过`glob`匹配的`./src/**/*.{js,ts,jsx,tsx}`这些文件中,tailwind才会生效
- 并通过`theme.screens`自定义一下响应式屏幕界点(为了通用性，我们与[bootstrap](https://getbootstrap.com/)一致)

```javascript
// tailwind.config.js
/** @type {import('tailwindcss').Config} */
module.exports = {
    prefix: 'tw-',
    darkMode: 'class',
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
        extend: {},
    },
    plugins: [],
};
```

接下来我们创建一下css文件
先创建一个用于定义css变量的文件`vars.css`，我们在里面定义一下标准的字体大小，以及标准的字体和一些特殊字体

```css
/* src/styles/vars.css */
:root {
    --font-size-base: 0.875rem;
    --font-family-standard: 'Source Sans Pro', 'Hiragino Sans GB', 'Microsoft Yahei', simsun,
        helvetica, arial, sans-serif, monospace;
    --font-family-kaiti: '楷体', 'STKaiti', 'Source Sans Pro', 'Hiragino Sans GB', 'Microsoft Yahei',
        simsun, helvetica, arial, sans-serif, monospace;
    --font-family-firacode: 'Fira Code', 'Source Sans Pro', 'Hiragino Sans GB', 'Microsoft Yahei',
        simsun, helvetica, arial, sans-serif, monospace;
}
```

在`tailwind.config.js`定义这些字体，后面就可以用`tw-字体名`来局部更改字体了

```javascript
// tailwind.config.js
module.exports = {
   ...
        extend: {
            fontFamily: {
                standard: 'var(--font-family-standard)',
                firacode: 'var(--font-family-firacode)',
                kaiti: 'var(--font-family-kaiti)',
            },
        },
    },
    plugins: [],
};
```

然后创建一个`styles/tailwind`目录，用于放置tailwind的自定义扩展样式
三个文件`base.css`,`components.css`,`utilities.css`分别用于

- 添加自定义tailwind基础层样式,一般用于覆盖一些tailwind中默认的基础样式
- 添加自定义tailwind组件层样式,一般无特殊需求可以用react组件抽象而不是在这里定义css类
- 添加自定义tailwindg工具层样式,可以在这里添加一些tailwind中不存在的一些样式类

需要注意的是这些文件中

- 如果要引用tailwind自带的值或`tailwind.config.js`的theme中配置的值,可以通过 `@apply`指令或`theme`函数获取
- 在`@layer`中添加的样式如果在程序中没有用到会在编译后被清除,如果需要强制存在于编译后的样式表,请在`@layer`外定义

它们的内容如下

```css
/* src/styles/tailwind/base.css */
@layer base {
    html {
        font-family: var(--font-family-base);
        font-size: var(--font-size-base);
    }
}
/* src/styles/tailwind/components.css */
@layer components {
}
/* src/styles/tailwind/utilities.css */
@layer utilities {
}
```

新建一个`app.css`文件用于放置全局样式，并在这个文件里我们测试一下tailwind的引用

```css
/* src/styles/app.css */
html,
body,
#root {
    @apply tw-bg-white tw-h-[100vh] tw-w-full tw-flex;
}
```

然后创建一个入口样式文件`index.css`来引用这些样式文件
> 注意引用顺序

```css
/* src/styles/index.css */
@import 'tailwindcss/base';
@import './tailwind/base.css';
@import 'tailwindcss/components';
@import './tailwind/components.css';
@import 'tailwindcss/utilities';
@import './tailwind/utilities.css';
@import './app.css';
```

现在我们不需要`src`目录下的`index.css`以及`App.css`了，直接删除掉，然后在`main.tsx`中导入`@/styles/index.css`

```tsx
// src/main.tsx
import '@/styles/index.css';

import App from './App';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <App />
    </React.StrictMode>,
);
```

然后更改一下`App.tsx`的代码

```tsx
// src/App.tsx
const App = () => {
    return (
        <div className="tw-w-full tw-bg-yellow-300 tw-flex tw-items-center tw-justify-center">
            <div className="tw-shadow-md tw-p-5 tw-bg-black tw-text-center tw-text-white tw-text-lg">
                欢迎来到3R教室，这是<span>React课程第一节</span>
            </div>
        </div>
    );
};

export default App;
```

启动一下应用查看样式是否生效
### 使用CSS MODULES
创建一个`App.moudule.css`，把`App.tsx`的样式放进去

```css
/* src/App.module.css */
.app {
    @apply tw-w-full tw-bg-yellow-300 tw-flex tw-items-center tw-justify-center;

    & > .container {
        @apply tw-shadow-md tw-p-5 tw-bg-black tw-text-center tw-text-white tw-text-lg;
    }
}
```

修改`App.tsx`

```tsx
// src/App.tsx
import $styles from './App.module.css';

const App = () => {
    return (
        <div className={$styles.app}>
            <div className={$styles.container}>
                欢迎来到3R教室，这是<span>React课程第一节</span>
            </div>
        </div>
    );
};

export default App;
```

因为css样式抽象到模块了，控制台会出现“No utility classes were detected in your source files. If this is unexpected, double-check the `content` option in your Tailwind CSS configuration”的警告，不用去理它，后面在`tsx`文件中某些地方直接使用tw的样式则自然会消失！
### 整合Antd
安装依赖
> 因为antd的时间组件依赖于dayjs，所以需要安装dayjs

```bash
~ pnpm add antd @ant-design/cssinjs dayjs
```

引入Antd的样式

```tsx
// src/main.tsx
import 'antd/dist/reset.css';
import '@/styles/index.css';
...
```

为了防止tailwind与antd产生样式冲突，需要修改一下tailwind的配置

```javascript
// tailwind.config.js
...
    corePlugins: {
        preflight: false,
    },
    plugins: [],
};
```

包装应用的时候需要使用`StyleProvider`取消Antd的降权（同样是为了防止tailwind与antd产生样式冲突），并且在`ConfigProvider`中把背景取消，然后换个紧凑皮肤`theme.defaultAlgorithm`，代码变成这样

```tsx
// src/App.tsx
import { Button, ConfigProvider, theme } from 'antd';
import { StyleProvider } from '@ant-design/cssinjs';
import 'dayjs/locale/zh-cn';
import zhCN from 'antd/locale/zh_CN';

import $styles from './App.module.css';

const App = () => {
    return (
        <ConfigProvider
            locale={zhCN}
            theme={{
                algorithm: theme.defaultAlgorithm,
                token: {
                    colorPrimary: '#00B96B',
                },
                components: {
                    Layout: {
                        colorBgBody: '',
                    },
                },
            }}
        >
            <StyleProvider hashPriority="high">
                <div className={$styles.app}>
                    <div className={$styles.container}>
                        欢迎来到3R教室，这是<span>React课程第一节</span>
                        <Button
                            type="primary"
                            className="!tw-bg-lime-400 !tw-text-emerald-900"
                            href="https://pincman.com/3r"
                            target="_blank"
                        >
                            点此打开
                        </Button>
                    </div>
                </div>
            </StyleProvider>
        </ConfigProvider>
    );
};
```

启动应用看一下是否正常，本节搞定！
![](https://img.pincman.com/media/202304040913399.png#id=Hu5Ej&originHeight=1870&originWidth=2934&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
