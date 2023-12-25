
## Use React Query to manage remote data
```jsx
import { QueryClient, QueryClientProvider } from 'react-query';
import '../styles/globals.css';

const queryClient = new QueryClient();

function App({ Component, pageProps }) {
  return (
    <QueryClientProvider client={queryClient}>
      <Component {...pageProps} />
    </QueryClientProvider>
  );
}

export default App;

```



## Page

index page is a grid view to show all the products
```jsx
import Page from '../components/Page';
import ProductCard from '../components/ProductCard';
import { getProducts } from '../lib/products';

export async function getStaticProps() {
  console.log('[HomePage] getStaticProps()');
  const products = await getProducts();
  return {
    props: { products },
    revalidate: parseInt(process.env.REVALIDATE_SECONDS),
  };
}

function HomePage({ products }) {
  console.log('[HomePage] render:', products);
  return (
    <Page title="Indoor Plants">
      <ul className="grid grid-cols-1 lg:grid-cols-3 gap-4">
        {products.map((product) => (
          <li key={product.id}>
            <ProductCard product={product} />
          </li>
        ))}
      </ul>
    </Page>
  );
}

export default HomePage;

```


Every product has it's own page use SSR
```jsx
import Image from 'next/image';
import AddToCartWidget from '../../components/AddToCartWidget';
import Page from '../../components/Page';
import { useUser } from '../../hooks/user';
import { ApiError } from '../../lib/api';
import { getProduct, getProducts } from '../../lib/products';

export async function getStaticPaths() {
  const products = await getProducts();
  const paths = products.map((product) => ({
    params: { id: product.id.toString() },
  }));
  console.log('getStaticPaths', paths);
  return {
    paths,
    fallback: 'blocking',
  };
}

export async function getStaticProps({ params: { id } }) {
  try {
    const product = await getProduct(id);
    console.log('[ProductPage] getStaticProps()', product);
    return {
      props: { product },
      revalidate: parseInt(process.env.REVALIDATE_SECONDS),
    };
  } catch (err) {
    if (err instanceof ApiError && err.status === 404) {
      return { notFound: true };
    }
    throw err;
  }
}

function ProductPage({ product }) {
  const user = useUser();

  console.log('[ProductPage] render:', product);
  return (
    <Page title={product.title}>
      <div className="flex flex-col lg:flex-row">
        <div>
          <Image src={product.pictureUrl} alt=""
            width={640} height={480} priority
          />
        </div>
        <div className="flex-1 lg:ml-4">
          <p className="text-sm">
            {product.description}
          </p>
          <p className="text-lg font-bold mt-2">
            {product.price}
          </p>
          {user && <AddToCartWidget productId={product.id} />}
        </div>
      </div>
    </Page>
  );
}

export default ProductPage;

```


Cart Page 

```jsx
function CartPage() {
  const query = useQuery('cartItems', () => fetchJson('/api/cart'));
  const cartItems = query.data;

  console.log('[CartPage] cartItems:', cartItems);
  return (
    <Page title="Cart">
      {cartItems && <CartTable cartItems={cartItems} />}
    </Page>
  );
}
```

Add to cart button is defined in `AddToCartWidget.js`

```jsx
const { useState } = require('react');
import { useRouter } from 'next/router';
import { useMutation } from 'react-query';
import { fetchJson } from '../lib/api';
import Button from './Button';

function AddToCartWidget({ productId }) {
  const router = useRouter();
  const [quantity, setQuantity] = useState(1);
  const mutation = useMutation(() =>
    fetchJson('/api/cart', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ productId, quantity })
    }));

  const handleClick = async () => {
    await mutation.mutateAsync();
    router.push('/cart');
  };

  return (
    <div className="py-2">
      <input type="number" min="1"
        className="border rounded px-3 py-1 mr-2 w-16 text-right"
        value={quantity.toString()}
        onChange={(event) => setQuantity(parseInt(event.target.value))}
      />
      {mutation.isLoading ? (
        <p>Loading...</p>
      ) : (
        <Button onClick={handleClick}>
          Add to cart
        </Button>
      )}
    </div>
  );
}

export default AddToCartWidget;

```



## BackEnd with Srapi

wrap hook to fetch user info
```js
export function useUser() {
  const query = useQuery(USER_QUERY_KEY, async () => {
    try {
      return await fetchJson('/api/user');
    } catch (err) {
      return undefined;
    }
  }, {
    cacheTime: Infinity,
    staleTime: 30_000, // ms
  });
  return query.data;
}

```






