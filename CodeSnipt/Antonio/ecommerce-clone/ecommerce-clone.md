

## nav bar
```jsx
<div className="border-b">
      <div className="flex h-16 items-center px-4">
        <StoreSwitcher items={stores} />
        <MainNav className="mx-6" />
        <div className="ml-auto flex items-center space-x-4">
          <ThemeToggle />
          <UserButton afterSignOutUrl="/" />
        </div>
      </div>
    </div>
```

- `border-b` used to split the navbar and the main area
- `ml-auto` used to align right and leaft left space as much as possible, we also can use `justify-between`

let's dive into `StoreSwitcher`

we store the store in the url, when user changed the store, we changed the url
```jsx
  const onStoreSelect = (store: { value: string, label: string }) => {
    setOpen(false);//after user selected, close the selected popup
    router.push(`/${store.value}`);
  };
  

{formattedItems.map((store) => (
                <CommandItem
                  key={store.value}
                  onSelect={() => onStoreSelect(store)}
                  className="text-sm"
                >
                  <Store className="mr-2 h-4 w-4" />
                  {store.label}
                  <Check
                    className={cn(
                      "ml-auto h-4 w-4",
                      currentStore?.value === store.value
                        ? "opacity-100"
                        : "opacity-0"
                    )}
                  />
                </CommandItem>
))}
```

## Page
![[Pasted image 20231219205810.png]]
dashboard is a async function, directly fetch the metric data and display
```jsx
const totalRevenue = await getTotalRevenue(params.storeId);
  const graphRevenue = await getGraphRevenue(params.storeId);
  const salesCount = await getSalesCount(params.storeId);
  const stockCount = await getStockCount(params.storeId);
```


billboard(广告牌) page

we must crate billboard, and we must uplod images, so we should set image cloue correctly
![[Pasted image 20231219210037.png]]

we don't need to create new dir and put page.jsx in it 
we just need to put `page.jsx` in `billboards/[billboardId]` dir, when 
we visit `billboards/create` or `billboards/xxy`, we just locaed in the 
`page.jsx`

if we visit `billbards/create`, the `params.billboardId` is create, and we can not 
fid any billboard that id is `create`, so the inialData is null.

in this way, the edit and create page can commonly use just one template, that is pretty good.

```jsx

const BillboardPage = async ({
  params
}: {
  params: { billboardId: string }
}) => {
  const billboard = await prismadb.billboard.findUnique({
    where: {
      id: params.billboardId
    }
  });

  return ( 
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <BillboardForm initialData={billboard} />
      </div>
    </div>
  );
}

export default BillboardPage;
```


## Api list
This is very cool, we can copy it and visit the api
```jsx
interface ApiListProps {
  entityName: string;
  entityIdName: string;
}

export const ApiList: React.FC<ApiListProps> = ({
  entityName,
  entityIdName,
}) => {
  const params = useParams();
  const origin = useOrigin();

  const baseUrl = `${origin}/api/${params.storeId}`;

  return (
    <>
      <ApiAlert title="GET" variant="public" description={`${baseUrl}/${entityName}`} />
      <ApiAlert title="GET" variant="public" description={`${baseUrl}/${entityName}/{${entityIdName}}`} />
      <ApiAlert title="POST" variant="admin" description={`${baseUrl}/${entityName}`} />
      <ApiAlert title="PATCH" variant="admin" description={`${baseUrl}/${entityName}/{${entityIdName}}`} />
      <ApiAlert title="DELETE" variant="admin" description={`${baseUrl}/${entityName}/{${entityIdName}}`} />
    </>
  );
};
```
