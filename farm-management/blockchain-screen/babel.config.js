// const isProd = process.env.NODE_ENV === 'production'

const plugins = [
  [
    'component',
    {
      libraryName: 'element-ui',
      styleLibraryName: 'theme-chalk'
    }
  ]
]
// if (isProd) {
//   plugins.push('transform-remove-console')
// }

module.exports = {
  presets: ['@vue/cli-plugin-babel/preset'],
  plugins
}
