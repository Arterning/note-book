
## use swr to fetch data

define fetcher

```js
import axios from 'axios';

  

const fetcher = (url: string) => axios.get(url).then((res) => res.data);

  

export default fetcher;
```


fetcher just return the response object


## how to use

```js
import useSWR from 'swr';

  

import fetcher from '@/libs/fetcher';

  

const usePost = (postId: string) => {

  const { data, error, isLoading, mutate } = useSWR(postId ? `/api/posts/${postId}` : null, fetcher);

  

  return {

    data,

    error,

    isLoading,

    mutate

  }

};

  

export default usePost;
```


what is swr
https://swr.vercel.app/zh-CN

“SWR” 这个名字来自于 `stale-while-revalidate`：一种由 [HTTP RFC 5861(opens in a new tab)](https://tools.ietf.org/html/rfc5861) 推广的 HTTP 缓存失效策略。这种策略首先从缓存中返回数据（过期的），同时发送 fetch 请求（重新验证），最后得到最新数据。


使用 SWR，组件将会**不断地**、**自动**获得最新数据流。  
UI 也会一直保持**快速响应**。



SWR 提供了 `mutate` 和 `useSWRMutation` 两个 API 用于更改远程数据及相关缓存。


当你调用 `mutate(key)`（或者只是使用绑定数据更改 API `mutate()`）时没有传入任何数据，它会触发资源的ravalidation(将数据标记为已过期并触发重新请求)。





## schema define

```js
// This is your Prisma schema file,

// learn more about it in the docs: https://pris.ly/d/prisma-schema

  

generator client {

  provider = "prisma-client-js"

}

  

datasource db {

  provider = "mongodb"

  url      = env("DATABASE_URL")

}

  

model User {

  id              String @id @default(auto()) @map("_id") @db.ObjectId

  name            String?

  username        String?   @unique

  bio             String?

  email           String?   @unique

  emailVerified   DateTime?

  image           String?

  coverImage      String?

  profileImage    String?

  hashedPassword  String?

  createdAt       DateTime @default(now())

  updatedAt       DateTime @updatedAt

  followingIds    String[] @db.ObjectId

  hasNotification Boolean?

  

  posts         Post[]

  comments      Comment[]

  notifications Notification[]

}

  

model Post {

  id                 String @id @default(auto()) @map("_id") @db.ObjectId

  body               String

  createdAt          DateTime @default(now())

  updatedAt          DateTime @updatedAt

  userId             String @db.ObjectId

  likedIds           String[] @db.ObjectId

  image              String?

  

  user User @relation(fields: [userId], references: [id], onDelete: Cascade)

  

  comments          Comment[]

}

  

model Comment {

  id                 String @id @default(auto()) @map("_id") @db.ObjectId

  body               String

  createdAt          DateTime @default(now())

  updatedAt          DateTime @updatedAt

  userId             String @db.ObjectId

  postId             String @db.ObjectId

  

  user User @relation(fields: [userId], references: [id], onDelete: Cascade)

  post Post @relation(fields: [postId], references: [id], onDelete: Cascade)

}

  

model Notification {

  id                 String @id @default(auto()) @map("_id") @db.ObjectId

  body               String

  userId             String @db.ObjectId

  createdAt          DateTime @default(now())

  

  user User @relation(fields: [userId], references: [id], onDelete: Cascade)

}
```



