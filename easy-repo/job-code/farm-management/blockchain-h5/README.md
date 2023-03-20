# blockchain-h5

## 简介
blockchain-h5，智慧稻米区块链溯源平台-用户h5，基于[vue-cli](https://uniapp.dcloud.io/quickstart-cli)方式创建的 uni-app 项目。

- [参考文档](https://uniapp.dcloud.io/)

## 前序准备
该项目模板为 h5 平台而设计。开发工具：VSCode。

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
git clone https://github.com/Hemingway-AT/blockchain-h5.git

# 进入项目目录
cd blockchain-h5

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动开发环境
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

## 其它

```bash
# 代码格式检查并自动修复
npm run lint

# 启动测试环境
npm run serve:test

# 启动生产环境
npm run serve:prod

# 预览生产环境打包分析
npm run analyze
```

## 注意事项
1. `master`和`test`分支为保护分支，`develop分支`保存了开发者的最新代码；
2. `feature-*`，基于develop分支新建；开发者应在feature分支（自行命名）下进行开发，功能完成后将代码合并至develop分支；
3. `hotfix-*`分支，为线上修复分支，基于master分支新建，修复好后合并到master分支进行发布，同时也需将最新更改同步到develop分支，最后`删除分支`，不可持续使用；
4. 注重命名和书写的可读性，代码应保证健壮性、避免出现undefined等运行时异常，异步模式下尽量采用async/await同步写法......

## License

[MIT](https://github.com/Hemingway-AT/uni-app-template/blob/main/LICENSE)

Copyright (c) 2021-present Hemingway-AT