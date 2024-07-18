

springboot maven 项目
子项目必须要放到父项目的目录下，否则会有问题



Notion Next 手册

[操作手册](https://tangly1024.com/article/notion-next-guide)

[更新记录](https://tangly1024.com/article/notion-next-changelogs)

[联系我们](https://tangly1024.com/about#9bcf6d8f32db4a3eb87612f06b0a0acb)


### **Authentication和Authorization的区别：**

Authentication：用户认证，指的是验证用户的身份，例如你希望以小A的身份登录，那么应用程序需要通过用户名和密码确认你真的是小A。

Authorization：授权，指的是确认你的身份之后提供给你权限，例如用户小A可以修改数据，而用户小B只能阅读数据。

由于http协议是无状态的，每一次请求都无状态。当一个用户通过用户名和密码登录了之后，他的下一个请求不会携带任何状态，应用程序无法知道他的身份，那就必须重新认证。因此我们希望用户登录成功之后的每一次http请求，都能够保存他的登录状态。

目前主流的用户认证方法有基于token和基于session两种方式。



## 需要不断打磨自己的求职案例

前端的案例可以部署到静态网站部署服务上，这个反正不需要自己维护服务器

后端的案例比较麻烦，需要自己搞一台服务器。还是可以整

使用进程内缓存和分布式缓存相结合的方式替换原有的单独进程内缓存；缓存加载方式从原来的全量加载改成按需加载；将账户模块的热点数据进行保存，从而缩短页面加载时间。

- 热练掌握分布式下的常见理论CAP、BASE。熟练掌挥 RPC(Feign)、分布式事务(Seata2PC、3PC、TCC)、配置中心(Apollo)、分布式链路追踪(SkyWalking)、分布式id(UUID、Snowflake)的使用及原理



## Resource
跟着Mic学架构
https://github.com/KrisCheng/500-interview-question-for-programmers/tree/master


## 关于笔记

对于有图片的情况 ，使用在线的编辑器编辑笔记是比较好的选择

总体使用感受下来，使用github托管仓库，使用vscode编辑笔记的**效果是不好的**

所以我现在基本要抛弃这种方案了

如果实在需要本地也保存 应该可以使用导出功能

## 云计算什么时候才是适合的

没有绝对适合的银弹

正如 Hansson 当时所说，云计算在两种极端情况下大有裨益：其一是应用程序极其简单且流量很低的情况，这时选择完全托管服务能摆脱大部分复杂性要素；其二是负载波动几乎毫无规律可言、大家不知道该部署多少服务器的情况，这时上云是最好的选择。但如今 37signals 已经不适用于上述两种情况。“Basecamp 多年的商业模式跟自有硬件都能良好协同，业务的增长轨迹也有很好的可预测性。”

事实上，虽然近年来云计算加速增长，但企业并没有放弃本地数据中心，很多企业继续依赖传统数据中心来处理其关键任务工作负载。


## 远程开机

**受控端主板需要支持网络唤醒，且在bios里开启了Wake-On-Lan唤醒功能（以下称wol唤醒）**

进入bios后，通常到“Power Managment（电源管理）”下寻找如下列选项：

"Boot on LAN";

"Wake on LAN";

"PME Event WakeUp";

"Resume by MAC LAN";

"Wake-Up by PCI card";

"Wake Up On PCI PME";

"Power On by PCI Card";

"WakeUp by PME of PCI";

"Power On By PCI Devices";

"WakeUp by Onborad LAN";

"Resume By PCI or PCI-E Device"或类似的功能，并可以启用它。

注：不同主板显示的功能名字不同，找到其中一个设置即可，笔记本基本不支持，具体情况可以咨询硬件生产商

报错信息 "Hardware assisted virtualization and data execution protection must be enabled in the BIOS" 意味着您的计算机上的虚拟化和数据执行保护功能需要在 BIOS 中启用才能够正常运行 Docker Desktop。这些功能对于 Docker 在 Windows 上的运行是必需的。

- **硬件支持限制**：某些较老的计算机可能不支持虚拟化功能，或者硬件本身不允许启用这些功能。在这种情况下，您可能无法在该计算机上运行 Docker Desktop。


## 网络请求的固定写法

新建一个http或者api的目录

新建一个http.ts或者request.ts 封装axios

对于一个资源的请求新建一个文件 比如todo.ts

导出函数或者类

```jsx
export function addMenu (parameter) {
  return request({
    url: api.addMenu,
    method: 'post',
    data: parameter
  })
}
```

```jsx
export class ArticleProvider {
  /**
   * 获取所有文章
   */
  static async getArticles(params): Promise<[IArticle[], number]> {
    return httpProvider.get('/article', { params });
  }
```

## nest 代码组织

```jsx
main.ts bootstrap() NestFactory.create(MainModule)
MainModule includes subModules
```

context 可以视为react官方给出的状态共享方法

context可以定义全局状态和修改全局状态的方法

全局状态和全局方法可以定义在app组件当中

无论是接单还是远程工作，准备自己的案例是非常重要的

面试，需要项目经验和算法题目

## projects

wipi next nest blog

springboot blog

django blog

## study

fastify study

快速生成html rest api endpoint

leetcode 算法

系统设计面试题

简历继续优化

## 数据库迁移和种子数据

通过typeorm 自动生成表结构

数据库迁移生成sql

通过生成随机字符串生成种子数据，命令的方式

## typeorm query builder

```jsx
const res = await this.articleRepository
      .createQueryBuilder('article')
      .where('article.title LIKE :keyword')
      .orWhere('article.summary LIKE :keyword')
      .orWhere('article.content LIKE :keyword')
      .setParameter('keyword', `%${keyword}%`)
      .getMany();
```

query 通过链式调用构建查询条件和底层sql 然后最后调用getMany或者getOne执行。

基本上和django和mybatis plus基本类似。

## typeorm entity

```jsx
@ApiProperty()
  @Column({ type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' })
  publishAt: Date; // 发布日期

  @ApiProperty()
  @CreateDateColumn({
    type: 'datetime',
    comment: '创建时间',
    name: 'create_at',
  })
  createAt: Date;

  @ApiProperty()
  @UpdateDateColumn({
    type: 'datetime',
    comment: '更新时间',
    name: 'update_at',
  })
  updateAt: Date;
```

## wipi entity 主要关系

```jsx
@ApiProperty()
  @OneToMany(
    () => Article,
    (article) => article.category
  )
  articles: Array<Article>;
```

```jsx
@ApiProperty()
  @ManyToMany(
    () => Article,
    (article) => article.tags
  )
  articles: Array<Article>;
```

```jsx
@ApiProperty()
  @ManyToOne(
    () => Category,
    (category) => category.articles,
    { cascade: true }
  )
  @JoinTable()
  category: Category;

  @ApiProperty()
  @ManyToMany(
    () => Tag,
    (tag) => tag.articles,
    { cascade: true }
  )
  @JoinTable()
  tags: Array<Tag>;
```

## 日志记录在不同的文件中

```jsx
import * as fs from 'fs-extra';
import * as log4js from 'log4js';
import { join } from 'path';

const LOG_DIR_NAME = '../../logs';

/**
 * 确保当前文件所在目录下的 'logs' 目录中包含了 'request'、'response' 和 'error' 三个子目录。
 * 如果这些子目录不存在，则会创建它们；如果已经存在，则不做任何操作。
 * 这样在后续的日志记录操作中，就可以直接往这些子目录中写入日志，而不用再关心目录的创建问题。
 */
fs.ensureDirSync(join(__dirname, LOG_DIR_NAME));
void ['request', 'response', 'error'].forEach((t) => {
  fs.ensureDirSync(join(__dirname, LOG_DIR_NAME, t));
});

const resolvePath = (dir, filename) => join(__dirname, LOG_DIR_NAME, dir, filename);

const commonCinfig = {
  type: 'dateFile',
  pattern: '-yyyy-MM-dd.log',
  alwaysIncludePattern: true,
};

log4js.configure({
  appenders: {
    request: {
      ...commonCinfig,
      filename: resolvePath('request', 'request.log'),
      category: 'request',
    },
    response: {
      ...commonCinfig,
      filename: resolvePath('response', 'response.log'),
      category: 'response',
    },
    error: {
      ...commonCinfig,
      filename: resolvePath('error', 'error.log'),
      category: 'error',
    },
  },
  categories: {
    default: { appenders: ['request'], level: 'info' },
    response: { appenders: ['response'], level: 'info' },
    error: { appenders: ['error'], level: 'info' },
  },
});

export const requestLogger = log4js.getLogger('request');
export const responseLogger = log4js.getLogger('response');
export const errorLogger = log4js.getLogger('error');
```

ensureDirSync 该函数确保该目录存在，如果目录结构不存在，它将由该函数创建

## 按照模块组织和测试代码

nest和django都是如此 模块划分 边界清晰 独立测试

## nest todo project

react 组件和页面代码

对于一个新的组件和页面代码，新建一个对应的目录，按照驼峰命名方式命名

页面文件统一命名为index.tsx作为export的对象