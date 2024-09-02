We allow non-loged in user can view some twitter, so we don't need protect all page and fuctions, so we dont't need protecte all page by sin-up page

so in next-auth config file , we don't config the sing-up pages


In sidebarItem.ts, we can check the curentUser, if it is null, then pop up login modal

```js
const handleClick = useCallback(() => {
    if (onClick) {
      return onClick();
    }

    if (auth && !currentUser) {
      loginModal.onOpen();
    } else if (href) {
      router.push(href);
    }
  }, [router, href, auth, loginModal, onClick, currentUser]);
```