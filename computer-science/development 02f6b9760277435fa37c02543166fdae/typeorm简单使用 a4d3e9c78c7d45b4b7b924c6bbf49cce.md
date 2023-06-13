# typeorm简单使用

Created time: May 12, 2023 8:59 AM

[文档地址](https://typeorm.bootcss.com/many-to-many-relations)

orm框架

如果项目单纯使用 js，推荐使用 [sequelize](https://link.zhihu.com/?target=https%3A//www.sequelize.com.cn/)，适配主流数据库，各种查询模式相对稳定，主流开源 node + js 框架（例如：[egg](https://link.zhihu.com/?target=https%3A//eggjs.org/zh-cn/tutorials/sequelize.html)）的数据库默认 orm 框架。如果项目还使用了 ts，推荐使用 [typeorm](https://link.zhihu.com/?target=https%3A//typeorm.bootcss.com/)，同样适配主流数据库，另外有更好的类型推导，主流开源 node + ts 框架（例如：[midway](https://link.zhihu.com/?target=https%3A//www.midwayjs.org/doc/component/typeorm)）的数据库默认 orm 框架。可以关注下 [prisma](https://link.zhihu.com/?target=https%3A//prisma.bootcss.com/)，相比 typeorm，类型推导上更加出色（属性查询的实体等），简洁的实体声明语法（语法高亮提示用 vscode 插件），还免费带有 [可视化桌面端应用](https://link.zhihu.com/?target=https%3A//www.prisma.io/studio)，整个生态相对比较完备。

注解

```jsx

//创建一个主列，它可以获取任何类型的任何值。你也可以指定列类型。 
//如果未指定列类型，则将从属性类型自动推断。
@PrimaryColumn() 

/**
@PrimaryGeneratedColumn("uuid") 创建一个主列，该值将使用uuid自动生成。
 Uuid 是一个独特的字符串 id。 你不必在保存之前手动分配其值，该值将自动生成。
**/
@PrimaryGeneratedColumn()
@PrimaryGeneratedColumn("uuid")

```

一对多

```jsx
import { Entity, PrimaryGeneratedColumn, Column, ManyToOne } from "typeorm";
import { User } from "./User";

@Entity()
export class Photo {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  url: string;

/**
第一个参数：表示返回类型是User
第二个参数：表示user如何映射到photos
可以在@ManyToOne / @OneToMany关系中省略@JoinColumn，
除非你需要自定义关联列在数据库中的名称
**/
  @ManyToOne(() => User, user => user.photos)
  user: User;
}
```

```jsx
import { Entity, PrimaryGeneratedColumn, Column, OneToMany } from "typeorm";
import { Photo } from "./Photo";

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @OneToMany(() => Photo, photo => photo.user)
  photos: Photo[];
}
```

多对多的关系

```jsx
import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";

@Entity()
export class Category {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;
}
```

```jsx
import { Entity, PrimaryGeneratedColumn, Column, ManyToMany, JoinTable } from "typeorm";
import { Category } from "./Category";

@Entity()
export class Question {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  title: string;

  @Column()
  text: string;

  @ManyToMany(() => Category)
  @JoinTable()
  categories: Category[];
}
```

如果需要双向，修改为

```jsx
import { Entity, PrimaryGeneratedColumn, Column, ManyToMany } from "typeorm";
import { Question } from "./Question";

@Entity()
export class Category {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;

  @ManyToMany(() => Question, question => question.categories)
  questions: Question[];
}
```