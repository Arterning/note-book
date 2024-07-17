## 坑爹的版本问题

如何解决 “0308010c:digital envelope routines::unsupported” 的错误
你至少有 3 种方法可以解决这个错误。我们将一个一个地看。任何一个都应该对你有用。

将 --openssl-legacy-provider 传递给 Webpack 或 CLI 工具
例如，在 React 应用程序中，你可以将 --openssl-legacy-provider 传递给启动脚本，如 "react-scripts --openssl-legacy-provider start"。

这应该就可以了。但是，如果这不能修复错误，那么就进行下一个修复。在许多情况下，它是有效的。

使用 Node JS  的 LTS 版本
考虑将你的 Node 版本降级到 16.16.0 或其他 LTS 版本。

目前，Node 的最新 LTS 版本是 18.12.1。你可以从 Node JS 官方网站下载它，或者使用 NVM 来安装它。


## 参考资料

- ****[微任务（Microtask）](https://zh.javascript.info/microtask-queue)****
- ****[事件循环：微任务和宏任务](https://zh.javascript.info/event-loop)****