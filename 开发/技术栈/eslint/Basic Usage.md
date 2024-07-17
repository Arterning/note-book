# eslint basic

* eslint和插件
eslint本身是一个命令行工具，可以通过eslint index.js 检查文件，而安装vscode插件可以在vscode中直接检查错误。

* 配置文件

```javascript
{
    "env": {
        "browser": true,
        "es2021": true
    },
    "extends": ["eslint:recommended", "plugin:react/recommended", "plugin:react/jsx-runtime"],
    "overrides": [
    ],
    "parserOptions": {
        "ecmaVersion": "latest"
    },
    "rules": {
        "prefer-destructuring": ["error", {
            "array": false,
            "object": true
        }]
    }
}

```

