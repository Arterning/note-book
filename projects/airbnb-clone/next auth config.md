
In middleware.ts

Here to set the protected route

```js
export { default } from "next-auth/middleware"

export const config = { 
  matcher: [
    "/trips",
    "/reservations",
    "/properties",
    "/favorites"
  ]
};

```