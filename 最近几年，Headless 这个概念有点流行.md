最近几年，Headless 这个概念有点流行，很多东西都在被 Headless 化，Headless 本意是“无头”，是指一个应用程序在后台运行，没有 UI 界面，但现在这个概念被运用到了更多的领域，如 Headless CMS、Headless Commerce、Headless 组件库等。

这里就简单梳理一下那些 Headless 化的东西：

- **Headless Browser**，也就是无头浏览器，最常见的有[Puppeteer](https://link.juejin.cn?target=https%3A%2F%2Fgithub.com%2Fpuppeteer%2Fpuppeteer "https://github.com/puppeteer/puppeteer")、[Playwright](https://link.juejin.cn?target=https%3A%2F%2Fplaywright.dev "https://playwright.dev")。它们都是一个没有 UI 界面的浏览器，可以通过特定的协议或 API 来操作页面（如打开 URL、点击页面上的按钮等），一般用于网页截图或生成 PDF 文件、爬取 SPA 网页、UI 自动化测试等。
    
- **Headless CMS**，是指仅有管理后台的内容管理系统，常见的有 [Strapi](https://link.juejin.cn?target=https%3A%2F%2Fstrapi.io "https://strapi.io")、[Directus](https://link.juejin.cn?target=https%3A%2F%2Fdirectus.io "https://directus.io") 等。传统的 CMS 如 Wordpress 等是包含管理后台、前台及一些默认主题的，而 Headless CMS 仅包含管理后台，对于前台页面、内容模型没有任何定义，这样用户可以灵活地定制自己所需的内容模型，并且前台页面无论是样式还是技术栈也都可以更灵活。
    
- **Headless Commerce**，是指无头电商系统，可以把它理解为特定领域的 Headless CMS，常见的有 [Commerce.js](https://link.juejin.cn?target=https%3A%2F%2Fcommercejs.com "https://commercejs.com")，Shopify 也有相关解决方案。这种电商系统仅提供管理后台，包含商品管理、订单管理、支付服务等基础功能，而面向用户的界面则可以灵活实现，无论是 Web 页面还是 App 都可以，样式也可以更个性化。
    
- **Headless Components**，是指无样式的组件或组件库，常见的有 [Radix UI](https://link.juejin.cn?target=https%3A%2F%2Fwww.radix-ui.com "https://www.radix-ui.com")、[Headless UI](https://link.juejin.cn?target=https%3A%2F%2Fheadlessui.dev "https://headlessui.dev")、[TanStack Table](https://link.juejin.cn?target=https%3A%2F%2Ftanstack.com%2Ftable "https://tanstack.com/table") 等。这种组件库的好处是样式可以完全自定义，不会让你的产品看上去千篇一律，而常规的组件库如 Ant Design 等无论怎么自定义主题，都能看到默认主题的影子，产品不够个性化。
    

Shadcn UI 实际上并不是组件库或 UI 框架。相反，它是可以根据文档“让我们复制并粘贴到应用程序中的可复用组件的集合”。它是由 vercel 的工程师[Shadcn](https://link.juejin.cn?target=https%3A%2F%2Ftwitter.com%2Fshadcn "https://twitter.com/shadcn")创建的，他还创建了一些知名的开源项目，如 [Taxonomy](https://link.juejin.cn?target=https%3A%2F%2Fvercel.com%2Ftemplates%2Fnext.js%2Ftaxonomy "https://vercel.com/templates/next.js/taxonomy")，[Next.js for Drupal](https://link.juejin.cn?target=https%3A%2F%2Fnext-drupal.org%2F "https://next-drupal.org/")和[Reflexjs](https://link.juejin.cn?target=https%3A%2F%2Freflexjs.org%2F "https://reflexjs.org/")。

Radix UI - 是一个无头 UI 库。也就是说，它有组件 API，但没有样式。Shadcn UI 建立在 Tailwind CSS 和 Radix UI 之上，目前支持 Next.js、Gatsby、Remix、Astro、Laravel 和 Vite，并且拥有与其他项目快速集成的能力——[安装指南](https://link.juejin.cn?target=https%3A%2F%2Fui.shadcn.com%2Fdocs%2Finstallation%2Fmanual "https://ui.shadcn.com/docs/installation/manual")。

  

