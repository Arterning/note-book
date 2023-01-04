# blockchain-minicode

## 简介
blockchain-minicode，中联智慧水稻 区块链溯源 微信小程序，一个 uni-app 项目。

- [参考文档](https://uniapp.dcloud.io/)

## 前序准备
该项目为 mp-weixin 平台而开发。开发工具：VSCode + 微信开发者工具。
## VSCode 环境
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
  // koroFileHeader 头部注释（快捷：ctrl + win + i）
  "fileheader.customMade": {
    "Description": "",
    "Autor": "Your Name",
    "Date": "Do not edit",
    "LastEditors": "Your Name",
    "LastEditTime": "Do not edit",
    "FilePath": "Do not edit"
  },
  // koroFileHeader 函数注释（快捷：ctrl + win + t）
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

## 开发

```bash
# 克隆项目
git clone http://gitlab.zoomlion.com/blockchain/blockchain-minicode.git

# 进入项目目录
cd blockchain-minicode

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动开发环境（编译到微信小程序平台）
npm run serve
```

## 发布

```bash
# 打包开发环境
npm run build:dev

# 打包测试环境
npm run build:test

# 打包生产环境
npm run build
```