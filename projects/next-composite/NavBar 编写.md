Create Next Projects

```bash

%% create app %%
npx create-next-app@latest my-app --typescript --tailwind --eslint


%% init components %%
npx shadcn-ui@latest init


%% add component %%
npx shadcn-ui@latest add button

```



RootLayout
```jsx
<html lang="en">
      <body className={inter.className}>
        <NavBar/>
        {/* because navbar is in fixed position, we should add padding top on every page ,
        we can add common style here for every page*/}

        <Toaster />
        <main className="md:pl-20 pt-16 h-full">
          {children}
        </main>
      </body>
</html>
```



```jsx
import Link from "next/link"


export const NavBar = () => {
    return (
        <div className="fixed w-full flex justify-between items-center py-2 px-4 h-16 border-b">
            <Link href="/">
                Home
            </Link>
            <Link href="/test">
                Test
            </Link>
            <Link href="/book">
                Book
            </Link>
        </div>
    )
}

```