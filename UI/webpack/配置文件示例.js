const path = require("path");

/**
 * 插件
 */
const HtmlWebpackPlugin = require("html-webpack-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const BundleAnalyzerPlugin =
  require("webpack-bundle-analyzer").BundleAnalyzerPlugin;

/**
 * path.resolve() 方法将路径或路径片段的序列解析为绝对路径。

path.resolve('/foo/bar', './baz');
// 返回: '/foo/bar/baz'

path.resolve('/foo/bar', '/tmp/file/');
// 返回: '/tmp/file'

path.resolve('wwwroot', 'static_files/png/', '../gif/image.gif');
// 如果当前工作目录是 /home/myself/node，
// 则返回 '/home/myself/node/wwwroot/static_files/gif/image.gif'

 */
module.exports = {
  mode: "development",
  devtool: "inline-source-map",
  entry: "./src/index.js",
  output: {
    filename: "dist.js",
    path: path.resolve(__dirname, "dist"),
  },
  resolve: {
    alias: {
      utils: path.resolve(__dirname, "src/utils"),
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
    new HtmlWebpackPlugin({
      title: "博客列表",
    }),
    new BundleAnalyzerPlugin(),
  ],
  module: {
    rules: [
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"],
      },
      {
        test: /\.(png|svg|jpg|jpeg|gif)$/i,
        type: "asset/resource",
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env"],
          },
        },
      },
    ],
  },
};

/**
 * 有了rules index.js 就可以直接使用import 导入css和图片了
 */

/**
 * 
 * 
 * 
import { getBlogPosts } from "./data";
import "./style.css";
import HeroImage from "./assets/images/hero.jpeg";
import "./test/date/printDate";

const blogs = getBlogPosts();
const ul = document.createElement("ul");
blogs.forEach((blog) => {
  const li = document.createElement("li");
  li.innerText = blog;
  ul.appendChild(li);
});
document.body.appendChild(ul);

const image = document.createElement("img");
image.src = HeroImage;
document.body.prepend(image);

const h1 = document.createElement("h1");
h1.innerText = "博客列表";
document.body.prepend(h1);

 */