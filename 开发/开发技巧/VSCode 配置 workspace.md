

案例

```json
{
  "folders": [
    {
      "path": "chapter1-hello"
    },
    {
      "path": "chapter2-theme"
    },
    {
      "path": "chapter3-context"
    },
    {
      "path": "chapter4-layout"
    }
  ],
  "settings": {
    "editor.formatOnSave": false,
    "editor.codeActionsOnSave": {
      "source.fixAll.eslint": true,
      "source.fixAll.stylelint": true
    },
    "eslint.workingDirectories": ["./*/"],
    "css.validate": false,
    "less.validate": false,
    "scss.validate": false,
    "postcss.validate": false,
    "emmet.includeLanguages": {
      "postcss": "css"
    },
    "stylelint.snippet": ["css", "scss", "less", "postcss"],
    "stylelint.validate": ["css", "scss", "less", "postcss"],
    "javascript.preferences.importModuleSpecifier": "project-relative",
    "typescript.suggest.jsdoc.generateReturns": false,
    "typescript.tsserver.experimental.enableProjectDiagnostics": true,
    "typescript.tsdk": "node_modules/typescript/lib",
    "eslint.packageManager": "pnpm",
    "stylelint.packageManager": "pnpm",
    "npm.packageManager": "pnpm"
  }
}

```