## Web
Web is a client side rendering, deine route in app.tsx
```tsx
          <Switch>
            <PrivateRoute exact path="/">
              <Todo />
            </PrivateRoute>
            <Route path="/login">
              <Login />
            </Route>
            <AdminRoute path="/admin">
              <Admin />
            </AdminRoute>
          </Switch>
```

We use the AuthContext to store the token and login logout method to mutate token var

```js
import {createContext} from "react"

export interface Auth {
  token: string | null;
  setToken: Function;
  login: Function;
  logout: Function;
  isAdmin: boolean,
}

const AuthContext = createContext<Auth>({
  login: async () => {},
  logout: async () => {},
  token: null,
  isAdmin: false,
  setToken: () => {},
});

export default AuthContext;


```



create nest project command

```bash

npx create-nest-app@latest {app_name}  
%% not the above... %%

$ npm i -g @nestjs/cli
$ nest new project-name

```

## API

Let's say the todo entity
```js
import { ApiProperty } from '@nestjs/swagger';
import {
  Column,
  CreateDateColumn,
  Entity,
  ManyToOne,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';
import { User } from '../../user/entities/user.entity';
import { TodoType } from './todo.type.entity';

export enum TodoStatus {
  TODO = 0, // 待完成
  DONE = 1, // 未完成
}

@Entity()
export class Todo {
  @ApiProperty()
  @PrimaryGeneratedColumn()
  id: number; // 自增 id

  @ApiProperty()
  @Column({ length: 500 })
  title: string; // 标题

  @ApiProperty()
  @Column('text')
  description?: string; // 具体内容

  @ApiProperty()
  @Column('int', { default: TodoStatus.TODO })
  status: TodoStatus; // 状态

  @ApiProperty({ required: false })
  @Column('text')
  media?: string;

  @ManyToOne(() => User, (user) => user.todos)
  author: User;

  @ManyToOne(() => TodoType, (todoType) => todoType.todos)
  todoType: TodoType;

  @CreateDateColumn({
    comment: '创建时间',
  })
  createdAt: Date;

  @UpdateDateColumn({
    comment: '更新时间',
  })
  updatedAt: Date;
}

```

- ManyToOne has two params, 1st is the function that return the target Type, 2nd is the reverse functions
- OneToMany has two params, 1st is the function that return the target Type, 2nd is the reverse functions
Let's say Use Entity
```js
import { ApiProperty } from '@nestjs/swagger';
import * as bcrypt from 'bcrypt';
import { Exclude } from 'class-transformer';
import {
  BeforeInsert,
  Column,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { Todo } from '../../todo/entities/todo.entity';

@Entity()
export class User {
  @ApiProperty({ description: '自增 id' })
  @PrimaryGeneratedColumn()
  id: number;

  @ApiProperty({ description: '标题' })
  @Column({ length: 500 })
  username: string;

  @ApiProperty({ description: '密码' })
  @Exclude()
  @Column({ length: 500 })
  password: string;

  @ApiProperty({ description: '邮箱' })
  @Column({ length: 500 })
  email: string;

  @ApiProperty({ description: '是否为管理员' })
  @Column('int', { default: 1 })
  is_admin?: number;

  @OneToMany(() => Todo, (todo) => todo.author, { cascade: true })
  todos: Todo[];

  @BeforeInsert()
  private async hashPassword() {
    const salt = await bcrypt.genSalt();
    this.password = await bcrypt.hash(this.password, salt);
  }
}

```