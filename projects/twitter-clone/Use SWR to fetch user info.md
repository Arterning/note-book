
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