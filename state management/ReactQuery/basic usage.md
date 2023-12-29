Great video. Important addition for everyone watching this: Don't use state managers for fetching data, use React Query and you will find that probably 90% of your state is actually server state. And React Query is handling refetching, loading state, errors and caching for you out of the box.

Can you elaborate? What if I got an e-commerce app with a cart?

 [@belkocik](https://www.youtube.com/channel/UCNOp5LoqK45NKEwGg7D-32g)  When a user is logged in, you store their cart on the server anyway. When adding an item, use useMutation to send a request to add item on the server and automatically refetch all items in the cart. When the user visits the site again, use useQuery to get all the items in the cart.

​ [@GVal98](https://www.youtube.com/channel/UCFFY80BHAAIL73xT8qRgu6A)  What about if user is not logged in? How to handle the problem when a not logged in user added a few items to a cart and then logged in? The question is what if there are items on the server that he added previously when he was logged in? Fetch only the items from server and show them in cart or compare the items in local storage with the server's one and do what? :D How to handle the issue? Would love to see some code beacause I don't know how to do it :D

 [@belkocik](https://www.youtube.com/channel/UCNOp5LoqK45NKEwGg7D-32g)  Either when adding items, you can ask the user to login to add items. Otherwise, if some items have been added to the localstorage, as soon as he logs in, add those new items to the server as well (need not update the existing items in the server)

[@belkocik](https://www.youtube.com/channel/UCNOp5LoqK45NKEwGg7D-32g)  You can generate a guest id/session/token and use it in requests just like a regular user id. But the server must be ready for this and distinguish users from guests. As for the second question, you just do it like you did before with your state managment. I think it's up to you whether you want to combine items or just download from them the server.



[@GVal98](https://www.youtube.com/channel/UCFFY80BHAAIL73xT8qRgu6A)  I don't think it's a good idea to fetch an auth data of the current user (email, name, phone, etc.) that doesn't need to be revalidated and only needs to be reset when a user logs out and then refetched again when they log in. Also, you need some api to manipulate that data. For example, in mobx I create a class with a public field "user", public methods "patchName", "logOut" and etc where I simply make some required requests and then change the store data like so: async logOut() { await userApi.logOut(); this.user = null; } I don't see such a simplicity with react-query. Also, when changing one data, somethimes you need to change other. How am I going to make a connection between states in react-query? In MobX I can simply do dependecy injection and refer to a root store like so: this.global.someStore.someMethod();


