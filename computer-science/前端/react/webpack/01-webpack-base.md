# webpack basic

## 前言
在webpack中，任何文件都可以通过import导入，前提是设置了对应的Loader
例如：
```javascript
import { getBlogPosts } from "./data";
import "./style.css";
import HeroImage from "./assets/images/hero.jpeg";
import "./test/date/printDate";

```

在打包过程中，可以通过插件干预打包过程
```javascript
  plugins: [
    new HtmlWebpackPlugin({
      title: "博客列表",
    }),
    new BundleAnalyzerPlugin(),
  ]
```

* 安装依赖

```bash
yarn add webpack webpack-cli --dev

# or
yarn install
```

因为只需要在开发的时候用到webpack，所以安装到Devdependencies即可

* 安装loader

```bash
yarn add css-loader style-loader
```


* 执行打包

```bash
npx webpack
```


# 示例配置文件
```javascript
const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const BundleAnalyzerPlugin =
  require("webpack-bundle-analyzer").BundleAnalyzerPlugin;

module.exports = {
  mode: "development",
  devtool: "inline-source-map",
  entry: "./src/index.js", //entry file name
  output: {
    filename: "[name].[content-hash].js", //output file name
    path: path.resolve(__dirname, "dist"), //output file directory
  },
  resolve: {
    alias: {
      utils: path.resolve(__dirname, "src/utils"),//指定路径别名
    },
  },
  optimization: {
    minimize: true,
    minimizer: [new TerserPlugin()],
  },
  devServer: {
    static: "./dist",
  },
  plugins: [
    //HtmlWebpackPlugin will auto generate index html file
    new HtmlWebpackPlugin({
      title: "博客列表",
    }),
    new BundleAnalyzerPlugin(),
  ],
  module: {
    rules: [
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"], //use css-loader to support import css file
      },
      {
        test: /\.(png|svg|jpg|jpeg|gif)$/i, //webpack native support import picture file, so don't need to config loader
        type: "asset/resource",
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",//use babel-loader to convert js file
          options: {
            presets: ["@babel/preset-env"],
          },
        },
      },
    ],
  },
};

```