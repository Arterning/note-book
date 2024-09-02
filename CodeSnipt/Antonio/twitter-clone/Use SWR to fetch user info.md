
what is swr
https://swr.vercel.app/zh-CN

“SWR” 这个名字来自于 `stale-while-revalidate`：一种由 [HTTP RFC 5861(opens in a new tab)](https://tools.ietf.org/html/rfc5861) 推广的 HTTP 缓存失效策略。这种策略首先从缓存中返回数据（过期的），同时发送 fetch 请求（重新验证），最后得到最新数据。


使用 SWR，组件将会**不断地**、**自动**获得最新数据流。  
UI 也会一直保持**快速响应**。

SWR 提供了 `mutate` 和 `useSWRMutation` 两个 API 用于更改远程数据及相关缓存。


当你调用 `mutate(key)`（或者只是使用绑定数据更改 API `mutate()`）时没有传入任何数据，它会触发资源的ravalidation(将数据标记为已过期并触发重新请求)。



in the useUsers.js hook, we use swr hook to fetch user info

```js
import useSWR from 'swr';

import fetcher from '@/libs/fetcher';

const useUsers = () => {
  const { data, error, isLoading, mutate } = useSWR('/api/users', fetcher);

  return {
    data,
    error,
    isLoading,
    mutate
  }
};

export default useUsers;

```

## use swr to fetch data

define fetcher

```js
import axios from 'axios';

const fetcher = (url: string) => axios.get(url).then((res) => res.data);

export default fetcher;
```


fetcher just return the response object