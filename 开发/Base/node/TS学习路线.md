---
aliases: [3R手册]
banner: "![[working.png]]"
cssclass: fullwidth,noyaml,noscroll,myhome
obsidianUIMode: preview
---
# 3R手册

## 🐙加入社区
首先你需要加入微信群方便获得日常通知
> 站长会邀请你

有两个微信群，其中
- "3RCD交流(一/二/N)"为平时闲聊所用，可以选择性加入，加入后请把群备注换成"3R-名称"
- "3R教室"为日常通知所用，请务必加入，加入后请把群备注换成"地区-技术栈-名称"

然后加入[discord社区](https://discord.3rcd.com)，加入后，通知站长升级为`member`权限

![[Pasted image 20230526151925.png]]
## 📒 资料获取
### 🗂文档和代码
#### 安装git
首先下载并安装git
- macos下使用`brew install git`即可安装（确保已安装`brew`）
- windows下直接去git官网下载后安装
- linux下根据发行版的不同使用各自的包管理工具安装，比如`sudo apt-get install git`
#### 添加公钥
接着使用以下命令生成你的`ssh_key`

```shell
ssh-keygen -C "pincman1988@gmail.com"  # 请把邮箱地址换成你自己的
```

生成后使用`cat`命令显示key字符串，并复制
> 根据操作系统不同，你的`.ssh`目录的路径可能有有所不同，但是都是位于用户目录下

```shell
cat .ssh/id_rsa.pub
```

接下来登录`git.3rcd.com`，然后点头像并选择"设置"，点"SSH/GPG秘钥"，或者直接访问`https://git.3rcd.com/user/settings/keys`，然后点击"添加秘钥"来添加你的key
> 秘钥名称可以随意填写

![[Pasted image 20230527000149.png]]

#### 克隆和更新
打开命令窗口（windows下可使用`powershell`，macos下可使用`terminal`或者`iterm2`，linux下根据你的UI框架的不同可以选择使用`konsole`等）
使用`git clone git@git.3rcd.com:class/{仓库名}.git`来进行克隆（即下载）
> 尽量不要使用`https`克隆，请使用`ssh`
比如

```shell
git clone git@git.3rcd.com:class/frontend.git
```

如果要更新仓库则在当前仓库的目录下使用`pull`命令
例如

```shell
cd frontend # 在把仓库克隆到本地之后打开该仓库目录
git pull origin main
```

如果本地更改了仓库内容导致无法更新怎么办？
只要提交即可，如下操作

```shell
git add .
git commit -m 'changed'
```

#### 仓库列表
目前有以下几个仓库
> 仅列出课程相关仓库，其它，比如项目，官网等仓库请自行联系`@cloneable`添加
- 前端课程: https://git.3rcd.com/class/frontend
- 后端课程: https://git.3rcd.com/class/backend
- 课程文档: https://git.3rcd.com/class/docs

#### 仓库使用

```ad-tip
建议的学习方式是边学习文档边对着文档查看代码
```

课程代码请使用vscode打开学习
安装vscode后通过如下下方式打开（以前端代码为例）：

```shell
cd frontend
code ./app.code-workspace
```

文档仓库请使用obsidian打开
安装obsidian选择"file->Open Vault"打开你克隆的`courses`仓库，打开后如图
![[Pasted image 20230527024721.png]]

### 🍋视频和录制下载
所有视频或者工具通过阿里网盘或百度网盘下载
> 同时包含TS全栈课和淘金课的视频及工具

阿里网盘[点此下载](https://www.aliyundrive.com/s/D191GiTKndx "https://www.aliyundrive.com/s/D191GiTKndx") 访问密码: 05gs

![image-20230621094333114](https://img.pincman.com/media/202306210943809.png)

百度网盘[点击下载](https://pan.baidu.com/s/1rHV9qXfoShwxmZMHhQ9u1g?pwd=7zgh) 访问密码:  7zgh 

![image-20230621094815939](https://img.pincman.com/media/202306210948288.png)


## 🦘学习建议

> 仅提供《TS全栈课》的学习建议，《掘金课》请自行跟着视频实操即可

目前以技术学习目的（部分为学习掘金营销课的同学）加入3R教室的同学大致有以下几类

- A类：零基础小白或转行者（包括移动开发，大数据，运维等）
- B类：Java,PHP等其它后端语言转技术栈
- C类：Vue,Angular等其它前端开发者转技术栈
- D类：React/Nextjs开发者
- E类：Node.js（学习过Nestjs）开发者

### 🔍学习方向

**对于A类同学:** 

严重建议请先学习Nestjs，这是由以下原因决定的

- 因为前端整体上比较卷，无论远程还是坐班工作竞争力都非常大，而后端虽然学习成本高，难度比较大，所以不是特别卷，求职比较轻松
- 另一方面是因为从难到易简单，从易到难很麻烦，在学会nestjs后基本掌握了TS的方方面面，随时可以轻松学习React
- 并且我们的课程规划方面重点也在Node.js/Nestjs的后端开发上，更新速度会比前端方面快很多
- 另外，内推的远程工作大部分都是TS全栈，但是课程都是以Node后端为核心，React为辅助型的岗位

所以可以先学习后端课，然后再学习前端课

**对于B类同学：**

因为已经拥有良好的后端开发基础，学习我们的TS后端课会非常轻松，所以同样建议先学习Nestjs课，而后再学习React

**对于C类同学：**

我们也是建议先学习Nestjs课程，因为已经会一种前端，而快速掌握一种后端技术栈可以帮助提升你的求职竞争力以及自由职业全栈开发能力

**对于D类同学：**

不用再多说，直接学习Nestjs课程

**对于E类同学：**

也不用多说，直接学习React课程

### ✏️学习流程

我们建议的学习流程是这样

1. 在选择好[学习方向][#学习方向]后按自身的类别选择[提前准备][#提前准备]的知识学习
2. 准备学习完成后按[学习方向][#学习方向]学习我们的实战课
3. 在学习的时候如果遇到问题可以去[discord](https://discord.3rcd.com)的**"问答"**频道提问，请在发问题贴的时候`@cloneable`或者`@愧怍`
4. 没门实战课学习到一半左右，可以同时学习我们商业项目源码，请`@玛卡巴卡`，让他给你开通一个我们项目组的成熟商业项目源码的权限
5. 学习完一个商业项目和整个课程后尝试自己做一个开源项目
6. 如果项目组有新的项目进来（查看[discord](https://discord.3rcd.com)中的**"内推与外包"**频道），可以用你的开源项目证明你的开发能力，加入项目组参与开发一个项目
7. 建立一个个人网站，发表一些技术文章，并附上你的开源项目和你参与开发的商业项目
8. 在[discord](https://discord.3rcd.com)中的**"内推与外包"**频道发现适合你的内推岗位并投递简历或者可以去[电鸭](https://eleduck.com/)，[领英](https://www.linkedin.com/)等求职远程岗位
9. 也可以去[v2ex](https://v2ex.com/)发帖子接单做
10. 或者作为独立开发者学习我们的《掘金课》来销售你自己开发应用，从而实现被动收入

![[Pasted image 20230528114756.png]]

### 📎提前准备
虽然3R教室的TS全栈课程也有讲解部分基础的内容，但是建议在学习3R课程前先做一下前置的预准备，这是因为整体来说3R教室的课程还是偏向于实战，因为这是为直接求职及独立开发者准备的课程
站长为每种类型的同学都提供了不同的准备方案

|                             文档                             |          匹配           |          掌握程度          |
| :----------------------------------------------------------: | :---------------------: | :------------------------: |
| [菜鸟教程的ES6教程](https://www.runoob.com/w3cnote/es6-tutorial.html) |        A类，B类         |          完全掌握          |
| [Typescript手册](https://bosens-china.github.io/Typescript-manual/) |        A类，B类         |          完全掌握          |
| [`this`,`call`,`apply`和`bind`的使用](https://juejin.cn/post/6844903972646420488) |        A类，B类         |          完全掌握          |
|    [TS装饰器与反射详解](https://3rcd.com/wiki/decorator)     | A类，B类，C类，D类，E类 |   可学习Nestjs时慢慢理解   |
| [html5与css3的基本使用](https://www.w3cschool.cn/html5css3/) |      A类，B类，E类      |          完全掌握          |
|       [React快速入门](https://zh-hans.react.dev/learn)       |   A类，B类，C类，E类    |          完全掌握          |
|          [阮一峰的ES6](https://es6.ruanyifeng.com/)          |        A类，B类         |          用于查询          |
| [Mysql基础教程](https://www.runoob.com/mysql/mysql-tutorial.html) |      A类，C类，D类      | 初步掌握，跟着课程慢慢深入 |
| Redis,Rabbitmq,Websocket等工具的基础概念，这些在[菜鸟教程](https://www.runoob.com/)上直接搜索就行 |      A类，C类，D类      | 在学习中用到的时候再去查阅 |
| Jwt授权及OAuth2认证流程，可通过学习[阮一峰的JWT文章](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)以及[理解OAuth2.0](https://www.ruanyifeng.com/blog/2014/05/oauth_2_0.html)快速掌握 |           A类           | 在学习中用到的时候再去查阅 |
| [OS运维的教程](https://www.osyunwei.com/archives/11344.html)学习Linux |      A类，C类，D类      |      可选的，非必须学      |

### 📚课程表
#### Nestjs最佳实践

> 核心技术栈版本   Node.js: 20   Nestjs: v9

1. 『基础入门』[[1-Typescript+Eslint+Prettier搭建Nestjs工程及断点调试 |Typescript+Eslint+Prettier搭建Nestjs工程及断点调试]]
2. 『基础入门』[[2-Nestjs核心概念详解 |Nestjs核心概念详解]]
3. 『内容模块』[[1-Nestjs整合Typeorm实现基本的CRUD操作及分页数据查询 |Nestjs整合Typeorm实现基本的CRUD操作及分页数据查询]]
4. 『内容模块』[[2-请求数据的验证和响应数据的序列化|请求数据的验证和响应数据的序列化]]
5. 『内容模块』[[3-数据关联与树形嵌套结构的分类和评论的实现 |数据关联与树形嵌套结构的分类和评论的实现]]
6. 『内容模块』[[4-自定义全局的验证管道，拦截器和过滤器 |自定义全局的验证管道，拦截器和过滤器]]
7. 『内容模块』[[5-自定义数据验证约束及约束中的依赖注入 |自定义数据验证约束及约束中的依赖注入]]
8. 『内容模块』[[6-批量操作及软删除(回收站)功能使用 |批量操作及软删除(回收站)功能使用]]
9. 『内容模块』[[7-使用ElasticSearch及Mysql两种方式实现全文搜索 |使用ElasticSearch及Mysql两种方式实现全文搜索]]
10. 『核心框架』[[1-实现一个CRUD框架以抽象化代码 |实现一个CRUD框架以抽象化代码]]
11. 『核心框架』[[2-自建动态配置系统，动态模块构造器及解构化应用实例 |自建动态配置系统，动态模块构造器及解构化应用实例]]
12. 『核心框架』[[3-构建配置式路由与Open API文档 |构建配置式路由与Open API文档]]
13. 『CLI工具』[[1-Yargs构建命令行工具以及数据迁移的实现|Yargs命令行系统构建以及数据库迁移命令实现]]
14. 『CLI工具』[[2-数据填充命令及数据工厂的实现 |数据填充命令及数据工厂的实现]]
15. 『用户与权限』[[1-用户模块开发以及使用Passport实现JWT认证和无痛刷新 |用户模块开发以及使用Passport实现JWT认证和无痛刷新]]
16. 『用户与权限』数据表动态关联及内容作者
17. 『用户与权限』使用OAuth2实现Github等第三方登录
18. 『用户与权限』用户注册,登录,找回密码绑定邮箱和手机号等验证功能实现
19. 『用户与权限』使用Redis+BullMQ实现基于消息队列的异步短信及邮件验证
20. 『用户与权限』websocket实现即时聊天及消息离线存储功能
21. 『用户与权限』基于CASL的RBAC动态角色及权限系统实现
22. 『文件模块』Fastify驱动下的文件上传下载导出及图片流式加载实现等功能的实现
23. 『文件模块』图片的自动剪裁及压缩实现实现
24. 『文件模块』整合腾讯云SDK实现文件的云存储
25. 『运维与测试』整合Log4j2实现日志功能
26. 『运维与测试』Nestjs应用的缓存与性能优化
27. 『性能与运维』使用Gitea+Drone实现自动化CI/CD
28. 『性能与运维』Cluster均衡负载及Fork进程详解及LNMP+PM2的服务器生产环境配置
29. 『性能与运维』使用Jest编写TDD测试以及E2E测试编写
30. 『其它知识』SWC.js+Yargs构建极速TS编译工具
#### Nextjs最佳实践

> 核心技术栈版本   React: 18   Nextjs:v13

1. 『基础入门』[[1-Vite+TS+React+Tailwind+Antd应用构建|使用Vite构建React应用]]
2. 『基础入门』[[2-TailWindCSS的使用详解 |TailWindCSS的使用详解]]
3. 『基础入门』[[3-常用Hooks详解及实践|React的常用Hooks详解及实践]](以Antd与Tailwind的暗黑模式动态切换及Antd多国语言切换)
4. 『基础入门』高级API及Hooks使用方法(防抖，性能优化等)
5. 『基础入门』Nextjs+TS+React+Tailwind+Antd应用构建
6. 『基础入门』Nextjs核心功能详解（一）- 路由、样式、项目组织与数据操作
7. 『基础入门』Nextjs核心功能详解（二）- 服务端组件，
8. 『状态及存储』Zustand和Immer的使用与持久化存储封装
9. 『状态及存储』使用Zustand实现动态暗黑主题，动态皮肤与多国语言等配置组件
10. 『数据操作』整合Primsa直接操作数据（一）
11. 『数据操作』整合Primsa直接操作数据（二）
12. 『数据操作』使用Axios操作后端API及其封装
13. 『数据操作』Swr.js的使用以及与Axios的整合
14. 『后台管理』Antd响应式布局实现
15. 『后台管理』路由懒加载及Loadding的实现
16. 『后台管理』调用Nestjs的后端API接口
17. 『后台管理』权限路由与动态菜单实现
18. 『后台管理』Svg组件与基于Ionify的图标组件的封装
19. 『后台管理』登录页面编写及JWT登录实现
20. 『后台管理』Github等社会化平台的OAuth登录实现
21. 『后台管理』websockets和消息广播的实现
22. 『后台管理』Pro components的表单与表格使用详解
23. 『后台管理』云存储、备份等系统设置实现
24. 『后台管理』MDX的使用详解
25. 『后台管理』文章分类等内容管理与用户权限管理的实现
26. 『后台管理』基于echarts的可视化组件的封装与仪表盘实现
27. 『后台管理』文件管理其它后台功能的实现
28. 『后台管理』React-Spring动画库的使用详解
29. 『后台管理』React-DND拖动库的使用详解
30. 『后台管理』使用React-Spring与React-DND实现支持Keep Alive的多标签功能
31. 『网站前台』radix-ui与shadcn/ui使用
32. 『网站前台』网站布局与基本样式编写
33. 『网站前台』首页开发与Swiper轮播库使用
34. 『网站前台』文章发布，列表页面及文章详情页编写
35. 『网站前台』无限级分类、标签、热门等挂件实现
36. 『网站前台』网址导航页面实现
37. 『网站前台』项目案例页面实现
38. 『网站前台』复刻管理后台的功能(配置组件，图标组件，JWT与OAuth2,Websockets等)
39. 『部署与运维』生成静态页面与CDN部署
40. 『部署与运维』使用Vercel部署
41. 『部署与运维』使用Gitea+Drone+PM2+Nginx自建服务器部署并实现反向代理
42. 『其它知识』整合Nestjs实现Monorepo与前后端同构
43. 『其它知识』React中如何编写TDD及E2E测试
44. 『其它知识』SWC.js编译器与Turopack介绍

#### 远程淘金课

> 本课无文档，请自行在网盘中获取

1. 自由职业与数字游民概念详解
2. 应用软件的订阅制销售渠道和被动收入实践
3. 网课的制作与推广方法详解
4. 海外接单渠道发掘与客户谈判技巧
5. 经营性社区与平台盈利实现
6. 远程工作求职方法与准备
7. AI与PS图片设计入门
8. Final cut pro/Motion视频剪辑的简单使用
9. 使用Screenflow与Camtasia制作网课
10. 『你的服务器价值百万』挂载硬盘以及搭建SSH开发服务器
11. 『你的服务器价值百万』搭建LNMP+wordpress搭建个人网站
12. 『你的服务器价值百万』使用wordpress搭建资源售卖站
13. 『你的服务器价值百万』使用wordpress搭建企业或工作室网站
14. 『你的服务器价值百万』使用nextcloud搭建个人网盘
15. 『你的服务器价值百万』使用gitea搭建自己的个人GIT仓库系统
16. 『你的服务器价值百万』使用drone构建ci/cd系统
17. 『你的服务器价值百万』使用react+docusuarus搭建静态网站
18. 『你的服务器价值百万』使用crmeb搭建商城小程序
19. 『你的服务器价值百万』Nginx+PM2+Node.js实现自己的API服务器
20. 『你的服务器价值百万』架设7x24运行的小爬虫