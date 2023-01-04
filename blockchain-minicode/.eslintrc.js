module.exports = {
  root: true,

  env: {
    browser: true,
    node: true,
    es6: true
  },
  extends: ['eslint:recommended', 'plugin:vue/essential', '@vue/prettier'],

  // add your custom rules here
  rules: {
    eqeqeq: 'error',
    'vue/multi-word-component-names': 'off',
    'no-debugger': 'warn'
  },

  globals: {
    uni: 'readonly',
    ROUTES: 'readonly'
  }
}
