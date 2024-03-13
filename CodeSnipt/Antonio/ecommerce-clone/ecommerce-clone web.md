![[Pasted image 20231219211523.png]]

## Integrate to admin

we should set the env as below:
```env
NEXT_PUBLIC_API_URL=http://localhost:3000/api/14a12048-95c2-414f-9ffb-7303156975cf

STORE_BILLBOARD="c44b835a-18d8-43e0-86b7-68e2e31fe0e0"
```

`14a12048-95c2-414f-9ffb-7303156975cf` is the store id
`c44b835a-18d8-43e0-86b7-68e2e31fe0e0` is the first billboard



## NavBar 
NavBar display all the billboards(Categorys)
```jsx

const Navbar = async () => {
  const categories = await getCategories();

  return ( 
    <div className="border-b">
      <Container>
        <div className="relative px-4 sm:px-6 lg:px-8 flex h-16 items-center">
          <Link href="/" className="ml-4 flex lg:ml-0 gap-x-2">
            <p className="font-bold text-xl">STORE</p>
          </Link>
          <MainNav data={categories} />
          <NavbarActions />
        </div>
      </Container>
    </div>
  );
};
```

- `NavbarActions` has set ml-auto inside, so it align right
- `NavbarActions` click , lead ser to cart page

## Product Display 
we can filter the products by size and color
![[Pasted image 20231219212553.png]]

the color and size we put it in the url, so in the server component, we can use the searchParam to find
the products that correspond to the filter conditions
```jsx
const products = await getProducts({ 
    categoryId: params.categoryId,
    colorId: searchParams.colorId,
    sizeId: searchParams.sizeId,
  });
```


