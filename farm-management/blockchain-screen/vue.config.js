console.log(
  'npm_lifecycle_event = ',
  process.env.npm_lifecycle_event,
  'NODE_ENV = ',
  process.env.NODE_ENV
)

const CompressionWebpackPlugin = require('compression-webpack-plugin')
const BundleAnalyzerPlugin =
  require('webpack-bundle-analyzer').BundleAnalyzerPlugin

const prodGzipExtensions = ['js', 'css']
const isProd = process.env.NODE_ENV === 'production'
const assetsCDN = {
  externals: {
    vue: 'Vue',
    'vue-router': 'VueRouter',
    axios: 'axios'
  },
  js: [
    '//cdn.jsdelivr.net/npm/vue@2.6.11/dist/vue.min.js',
    '//cdn.jsdelivr.net/npm/vue-router@3.5.2/dist/vue-router.min.js',
    '//cdn.jsdelivr.net/npm/axios@0.22.0/dist/axios.min.js'
  ]
}

module.exports = {
  publicPath: '/blockchain-screen/',
  assetsDir: 'assets', // 放置生成的静态资源 (js、css、img、fonts) 的 (相对于 outputDir 的) 目录。
  lintOnSave: process.env.NODE_ENV !== 'production',
  productionSourceMap: false,
  configureWebpack: config => {
    if (isProd) {
      config.optimization.splitChunks = {
        chunks: 'async',
        minSize: 30000,
        maxSize: 1024000,
        minChunks: 1,
        maxAsyncRequests: 5,
        maxInitialRequests: 3,
        automaticNameDelimiter: '~',
        name: true,
        cacheGroups: {
          vendors: {
            name: 'vendor',
            chunks: 'all',
            test: /[\\/]node_modules[\\/]/,
            priority: -10
          },
          default: {
            minChunks: 2,
            priority: -20,
            reuseExistingChunk: true
          }
        }
      }

      config.plugins.push(
        new CompressionWebpackPlugin({
          algorithm: 'gzip',
          test: new RegExp('\\.(' + prodGzipExtensions.join('|') + ')$'),
          threshold: 10240,
          minRatio: 0.8
        })
      )

      config.plugins.push(new BundleAnalyzerPlugin())

      config.externals = assetsCDN.externals
    }
  },
  chainWebpack: config => {
    if (isProd) {
      config.plugin('html').tap(args => {
        args[0].cdn = assetsCDN
        return args
      })
    }
  },
  css: {
    requireModuleExtension: false
  },
  devServer: {
    open: true,
    host: '0.0.0.0',
    port: 8000,
    https: true,
    hot: true
  }
}
