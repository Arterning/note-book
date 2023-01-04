'use strict'

console.log(
  'npm_lifecycle_event = ',
  process.env.npm_lifecycle_event,
  'NODE_ENV = ',
  process.env.NODE_ENV
)

const CompressionWebpackPlugin = require('compression-webpack-plugin')
const prodGzipExtensions = ['js', 'css']
const isProd = process.env.NODE_ENV === 'production'

const TransformPages = require('uni-read-pages')
const { webpack } = new TransformPages()
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer')
  .BundleAnalyzerPlugin

const isAnalyze = process.env.npm_lifecycle_event === 'analyze'

module.exports = {
  lintOnSave: process.env.NODE_ENV === 'development', // 是否开启eslint保存检测，有效值：ture | false | 'error'
  runtimeCompiler: true, // 运行时版本是否需要编译
  productionSourceMap: false,
  transpileDependencies: ['@dcloudio/uni-ui', 'uni-simple-router'], // 通过babel显式转义指定依赖（参考：https://cli.vuejs.org/zh/config/#transpiledependencies）
  chainWebpack: config => {
    // when there are many pages, it will cause too many meaningless requests
    config.plugins.delete('prefetch')

    // set preserveWhitespace
    config.module
      .rule('vue')
      .use('vue-loader')
      .loader('vue-loader')
      .tap(options => {
        options.compilerOptions.preserveWhitespace = true
        return options
      })
      .end()
  },
  configureWebpack: config => {
    config.plugins.push(
      new webpack.DefinePlugin({
        ROUTES: webpack.DefinePlugin.runtimeValue(() => {
          const tfPages = new TransformPages({
            includes: ['path', 'aliasPath', 'name', 'meta']
          })
          return JSON.stringify(tfPages.routes)
        }, true) // 注入全局Routes对象
      })
    )

    if (isProd) {
      config.plugins.push(
        new CompressionWebpackPlugin({
          algorithm: 'gzip',
          test: new RegExp('\\.(' + prodGzipExtensions.join('|') + ')$'),
          threshold: 10240,
          minRatio: 0.8
        })
      )
    }

    if (isAnalyze) {
      config.plugins.push(new BundleAnalyzerPlugin())
    }
  }
}
