# gastby

[参考文档](https://www.gatsbyjs.cn/tutorial/part-zero/)

要安装 Gatsby 和 Node.js，建议使用 Homebrew。 开始做一些设置可以免去以后很多麻烦！

Gatsby CLI 可通过 npm 安装，命令行运行 `npm install -g gatsby-cli` 进行全局安装。

要查看可用的命令，请运行 `gatsby --help`。

## 开始

```jsx
gatsby new hello-world https://github.com/gatsbyjs/gatsby-starter-hello-world
cd hello-world
gatsby develop

//解释
//gatsby new [SITE_DIRECTORY_NAME] [URL_OF_STARTER_GITHUB_REPO]
```

每个定义在 `src/pages/*.js` 中的 React 组件，在浏览器里都是一个页面。让我们动手来看看。

新建一个文件，路径是 `src/pages/about.js`

访问这个页面 [http://localhost:8000/about/](http://localhost:8000/about/%E3%80%82) 即可

## 连接不同页面

来看看 Gatsby 网站的路由功能。

```jsx
<Link to="/contact/">Contact</Link>
```

## 构建和部署

```jsx
gatsby build
```

## 使用CSS

全局样式 global.css

```jsx
├── package.json
├── src
│   └── pages
│       └── index.js
│   └── styles
│       └── global.css
```

新建文件 gatsby-browser.js

```jsx
├── package.json
├── src
│   └── pages
│       └── index.js
│   └── styles
│       └── global.css
├── gatsby-browser.js
```

导入样式

gatsby-browser.js

```jsx

import "./src/styles/global.css"
// 或者:
// require('./src/styles/global.css')

```

注意：CommonJS 格式的（require）和 ES Module 格式的（import）语法都可以在这里使用。 如果你不确定选择哪个，通常最好使用 import。 但是，当处理仅在 Node.js 环境中运行的文件时（如 gatsby-node.js），则需要使用 require。

## ****组件范围内 CSS****

文件名以 .module.css 结尾，而不是通常的 .css 结尾。 这是告诉 Gatsby 该 CSS 文件应作为 CSS 模块而不是纯 CSS 处理的方式。

```jsx
import React from "react"
import containerStyles from "./container.module.css"

export default ({ children }) => (
  <div className={containerStyles.container}>{children}</div>
)
```

```jsx
.container {
  margin: 3rem auto;
  max-width: 600px;
}
```

[布局组件](gastby%2051d92d9b6e86404d8d7c5ca43ca41fbc/%E5%B8%83%E5%B1%80%E7%BB%84%E4%BB%B6%200d7e58c645b94f7383fd2c1c6fef63b4.md)

[gastby中的数据](gastby%2051d92d9b6e86404d8d7c5ca43ca41fbc/gastby%E4%B8%AD%E7%9A%84%E6%95%B0%E6%8D%AE%205d1a495963b349dc99e8fe3c92690f4a.md)

[数据源插件](gastby%2051d92d9b6e86404d8d7c5ca43ca41fbc/%E6%95%B0%E6%8D%AE%E6%BA%90%E6%8F%92%E4%BB%B6%20d610335c76b449f5a53345c4e1e21a99.md)

[数据转换插件](gastby%2051d92d9b6e86404d8d7c5ca43ca41fbc/%E6%95%B0%E6%8D%AE%E8%BD%AC%E6%8D%A2%E6%8F%92%E4%BB%B6%2045d771776dd0477e84c2153824b177c9.md)

[展示markdown](gastby%2051d92d9b6e86404d8d7c5ca43ca41fbc/%E5%B1%95%E7%A4%BAmarkdown%2016c3f05bc5904864bf83a1b9900e2200.md)