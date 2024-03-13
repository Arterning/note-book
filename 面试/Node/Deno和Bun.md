
Deno含有以下功能亮点：

- 默认安全。外部代码没有文件系统、网络、环境的访问权限，除非显式开启。
    
- 支持开箱即用的TypeScript的环境，尽管它仍需要编译为JavaScript才能运行，但这是在内部完成的，因此对用户来说TypeScript的行为就好像它是原生支持的一样。
    
- 只分发一个独立的可执行文件（deno）。
    
- 有着内建的工具箱，比如一个依赖信息查看器（deno info）和一个代码格式化工具（deno fmt）。
    
- 有一组经过审计的标准模块，保证能在Deno上工作。
    
- 脚本代码能被打包为一个单独的JavaScript文件。

- 它基于 Rust 语言构建（对比与 C++ 之于 Node）。



事实上，Bun 除了作为 JavaScript 的执行环境，同时也能跑测试 (test runner)，也可以作为打包工具 (bundler)，像是 Webpack 或 Vite，同时能做为套件管理工具 (package manager) 像是 npm 或 Yan。简言之，是个一体成形，让你可以直接在上面用绝大多数开发 JavaScript 时要用的工具。


Bun 的一大卖点是速度快。举例来说，Deno 每秒能处理的 React SSR 请求量是 Node.js 的两倍多，但是 Bun 更进一步做到能处理 Deno 两倍多的请求量 (也就是 Node.js 四倍多的请求量)。在 WebSocket 的讯息传递量，Bun 也是 Deno 的两倍多。在跟其他套件管理工具 (例如 npm 与 Yarn) 比，Bun 的套件安装速度超过十倍快。而在跑测试上，也比目前社群最快的 Vitest 还快不少。

除此之外，Bun 原生支援 TypeScript，让现在社群中多数人选择的 TypeScript，在使用上可以省去转译 (transpile)。在过去 Node.js，需要先把 TypeScript 转译成 JavaScript 后才能跑，Bun 原生支持就不需要额外转译，速度也就比较快。



















