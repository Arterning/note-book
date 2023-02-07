# 中联智慧农场运营平台 blockchain-platform

## 简介
blockchain-platform，基于[vue-element-admin@4.4.4](https://panjiachen.github.io/vue-element-admin)搭建。

- [使用文档](https://panjiachen.github.io/vue-element-admin-site/zh/)

## 前序准备
你需要在本地安装 VSCode 插件`Vetur`， `ESLint`，`Prettier-Code formatter`及`KoroFileHeader`。编辑器配置文件 setting.json 参考：
```bash
{
  // 全局文件预览，default = true（单击文件呈现覆盖效果，双击即打开新文件）
  "workbench.editor.enablePreview": false,

  // koroFileHeader 插件配置（参考：https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE）
  "fileheader.configObj": {
    "createFileTime": false, // 更改为当前生成注释的时间
    "wideSame": true, // 设置注释字段等宽
    "wideNum": 13, // 足够宽的长度，每个字段都有空格
    "autoAdd": true, // 自动添加文件头部注释
    "autoAlready": true, // 只让支持的语言，自动添加头部注释
    "prohibitAutoAdd": ["json", "js"] // 自动添加头部注释黑名单
  },
  // koroFileHeader 头部注释
  "fileheader.customMade": {
    "Description": "",
    "Autor": "Your Name",
    "Date": "Do not edit",
    "LastEditors": "Your Name",
    "LastEditTime": "Do not edit",
    "FilePath": "Do not edit"
  },
  // koroFileHeader 函数注释
  "fileheader.cursorMode": {
    "description": "",
    "param": "",
    "return": "",
    "author": "Your Name",
    "example": ""
  },

  // 设置保存时自动修复（参考：https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint）
  "editor.codeActionsOnSave": {
    "source.fixAll": true, // 包含eslint在内的自动修复
  },
  "diffEditor.ignoreTrimWhitespace": false,
  "files.eol": "\n",
}
```

## 功能

```
- 待完善

- 待完善
  - 待完善
  - ...

- ...
```

## 开发

```bash
# 克隆项目
git clone http://gitlab.zoomlion.com/agriculture/blockchain-platform.git

# 进入项目目录
cd blockchain-platform

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动服务
npm run serve
```

浏览器访问 http://localhost:9527

## 发布

```bash
# 构建开发环境
npm run build:dev

# 构建测试环境
npm run build:test

# 构建生产环境
npm run build
```

## 其它

```bash
# 代码格式检查并自动修复
npm run lint

# 本地服务mock
npm run mock

# 本地服务调试测试环境
npm run serve:test

# 本地服务调试生产环境
npm run serve:prod

# 预览生产环境打包分析
npm run analyze
```
## Browsers support

Modern browsers and Internet Explorer 10+.

| [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/edge/edge_48x48.png" alt="IE / Edge" width="24px" height="24px" />](https://godban.github.io/browsers-support-badges/)</br>IE / Edge | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/firefox/firefox_48x48.png" alt="Firefox" width="24px" height="24px" />](https://godban.github.io/browsers-support-badges/)</br>Firefox | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/chrome/chrome_48x48.png" alt="Chrome" width="24px" height="24px" />](https://godban.github.io/browsers-support-badges/)</br>Chrome | [<img src="https://raw.githubusercontent.com/alrra/browser-logos/master/src/safari/safari_48x48.png" alt="Safari" width="24px" height="24px" />](https://godban.github.io/browsers-support-badges/)</br>Safari |
| --------- | --------- | --------- | --------- |
| IE10, IE11, Edge | last 2 versions | last 2 versions | last 2 versions |

## License

[MIT](https://github.com/Hemingway-AT/vue-element-admin-custom/blob/main/LICENSE)

Copyright (c) 2021-present Hemingway-AT
