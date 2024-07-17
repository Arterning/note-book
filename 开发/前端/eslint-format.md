

## webstorm
设置eslint 格式化

![run eslint on save.png](run%20eslint%20on%20save.png)

设置Nodejs路径为wsl ubuntu的路径

![node interpreter.png](node%20interpreter.png)

这样点击package.json中的运行图标即可运行项目，不用敲击命令

我喜欢这样，越简单越好。

## vscode

配置setting.json

```javascript
{
    "editor.formatOnSave": false,
    "editor.codeActionsOnSave": {
      "source.fixAll.eslint": true,
      "source.fixAll.stylelint": true
    },
    "eslint.workingDirectories": ["./*"]
}
```

## 工程配置文件是否应该上传到git仓库

是的，通常情况下，你应该将 `.vscode` 文件夹以及其中的配置文件上传到 Git 仓库中，特别是其中包含了与工程设置相关的配置文件，比如 `settings.json`。

这么做的好处包括：

1. **共享配置：** 将 `.vscode` 文件夹上传到 Git 仓库中可以确保团队中的所有成员都能够使用相同的工程设置，从而保持代码风格和一致的开发体验。

2. **避免重复配置：** 在不同的开发环境中（比如不同的开发者电脑或不同的团队成员），使用相同的配置可以减少手动配置的时间和工作量。

3. **维护历史记录：** 将配置文件纳入版本控制可以确保你能够追溯配置的变更历史，这在解决问题或回退到以前的配置时非常有用。

然而，有一些例外情况需要考虑：

1. **安全性：** 确保你没有将包含敏感信息的配置（如认证密钥）上传到仓库中。这类敏感信息应该通过环境变量或其他安全方式进行管理，而不是直接写入配置文件。

2. **团队协作：** 在某些情况下，团队可能会有不同的配置偏好。如果你的团队中有不同的配置需求，可以考虑提供一个默认的 `.vscode/settings.json` 文件，并在工程的 README 文件中说明如何根据个人需求进行定制。

总的来说，将 `.vscode` 文件夹中的配置文件上传到 Git 仓库中是一个很好的做法，但你也需要注意敏感信息和团队需求等方面的因素。


## 总结

如果项目存在比较复杂的工程配置，最好将工程配置文件.idea或者.vscode上传到github
给团队的其他成员使用。不要信其他人的，说什么工程配置文件不应该上传github。

不然初学者还得费劲巴拉的找vscode配置，这不是扯淡吗。

