
user can put  their own house to the system, and other people rent it.

use clerk to authenicate



![[Pasted image 20231219203511.png]]


## NavBar

```html
<div className="fixed w-full bg-white z-10 shadow-sm">
```

## Category
```html
<div
        className="
          pt-4
          flex 
          flex-row 
          items-center 
          justify-between
          overflow-x-auto
        "
      >
        {categories.map((item) => (
          <CategoryBox 
            key={item.label}
            label={item.label}
            icon={item.icon}
            selected={category === item.label}
          />
        ))}
      </div>
```


## List
This component is used to display all the houses
```jsx
<div 
          className="
            pt-24
            grid 
            grid-cols-1 
            sm:grid-cols-2 
            md:grid-cols-3 
            lg:grid-cols-4
            xl:grid-cols-5
            2xl:grid-cols-6
            gap-8
          "
        >
          {listings.map((listing: any) => (
            <ListingCard
              currentUser={currentUser}
              key={listing.id}
              data={listing}
            />
          ))}
        </div>
```


## prisma schema

```js
// This is your Prisma schema file,
// learn more about it in the docs: https://pris.ly/d/prisma-schema

generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "mongodb"
  url      = env("DATABASE_URL")
}

model User {
  id              String @id @default(auto()) @map("_id") @db.ObjectId
  name            String?
  email           String?   @unique
  emailVerified   DateTime?
  image           String?
  hashedPassword  String?
  createdAt       DateTime @default(now())
  updatedAt       DateTime @updatedAt
  favoriteIds     String[] @db.ObjectId

  //why user has may accouns??
  accounts Account[]

  //user have many listing(house)
  listings Listing[]

  //user have many reservation
  reservations Reservation[]
}

model Account {
  id String @id @default(auto()) @map("_id") @db.ObjectId
  userId             String   @db.ObjectId
  type               String
  provider           String
  providerAccountId  String
  refresh_token      String?  @db.String
  access_token       String?  @db.String
  expires_at         Int?
  token_type         String?
  scope              String?
  id_token           String?  @db.String
  session_state      String?

  user User @relation(fields: [userId], references: [id], onDelete: Cascade)

  @@unique([provider, providerAccountId])
}

model Listing {
  id             String @id @default(auto()) @map("_id") @db.ObjectId
  title String
  description String
  imageSrc String
  createdAt DateTime @default(now())
  category  String
  roomCount Int
  bathroomCount Int
  guestCount Int
  locationValue String
  userId String @db.ObjectId
  price Int

  user User @relation(fields: [userId], references: [id], onDelete: Cascade)
  reservations Reservation[]
}

model Reservation {
  id String @id @default(auto()) @map("_id") @db.ObjectId
  userId String @db.ObjectId
  listingId String @db.ObjectId  
  startDate DateTime
  endDate DateTime
  totalPrice Int
  createdAt DateTime @default(now())

  user User @relation(fields: [userId], references: [id], onDelete: Cascade)
  listing Listing @relation(fields: [listingId], references: [id], onDelete: Cascade)
}
```


## db query

```js
const favorites = await prisma.listing.findMany({
      where: {
        id: {
          in: [...(currentUser.favoriteIds || [])]
        }
      }
    });
    
const listing = await prisma.listing.findUnique({
      where: {
        id: listingId,
      },
      include: {
        user: true
      }
    });
```


## ClientOnly component
```jsx

const ClientOnly: React.FC<ClientOnlyProps> = ({ 
  children
}) => {
  const [hasMounted, setHasMounted] = useState(false);

  useEffect(() => {
      setHasMounted(true);
  }, [])

  if (!hasMounted) return null;

  return (
    <>
      {children}
    </>
  );
};

```

we use client only component to wrap all the client component, to avoid hydration error

```jsx
<ClientOnly>
          <ToasterProvider />
          <LoginModal />
          <RegisterModal />
          <SearchModal />
          <RentModal />
          <Navbar currentUser={currentUser} />
</ClientOnly>
```