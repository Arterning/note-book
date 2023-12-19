![[Pasted image 20231219144029.png]]

we can use grid layout
"grid grid-cols-4 h-full"


## Sidbar

define menu items array and map the items, display all the menu items.
```js
const items = [
 {
      icon: BsHouseFill,
      label: 'Home',
      href: '/',
    },...
]

items.map(item=><SideBarItem/>)
```

```
flex flex-col items-end
```


## Post List

```js
import usePosts from '@/hooks/usePosts';
import PostItem from './PostItem';

interface PostFeedProps {

  userId?: string;

}
const PostFeed: React.FC<PostFeedProps> = ({ userId }) => {

  const { data: posts = [] } = usePosts(userId);
  return (
    <>

      {posts.map((post: Record<string, any>,) => (
        <PostItem userId={userId} key={post.id} data={post} />
      ))}
    </>
  );

};
export default PostFeed;
```



## FollowBar


```js
{
	users.map(user: Record<string, any>) => (
		...
	)
}
```


use flex layout
```css
flex flex-col gap-6 mt-4
```


## force user to login
if user is not log in , we should popup login model to force user to login

```js
  const { data: currentUser } = useCurrentUser();

  const onClick = useCallback(() => {
    if (!currentUser) {
      return loginModal.onOpen();
    }
    router.push('/');

  }, [...]);
```

model state is store in zustand
```js
import { create } from 'zustand';


interface LoginModalStore {

  isOpen: boolean;

  onOpen: () => void;

  onClose: () => void;

}

  

const useLoginModal = create<LoginModalStore>((set) => ({

  isOpen: false,

  onOpen: () => set({ isOpen: true }),

  onClose: () => set({ isOpen: false })

}));

  
  

export default useLoginModal;
```


LoginModel component is placed in the App Component
```js
export default function App({ Component, pageProps }: AppProps) {

  return (
    <SessionProvider session={pageProps.session}>
      <Toaster />
      <RegisterModal />
      <LoginModal />
      <EditModal />
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </SessionProvider>

  )

}
```
