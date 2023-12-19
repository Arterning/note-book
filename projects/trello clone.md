
## create safe action
combine with the zod , abstract the common logic
```js
import { z } from "zod";

  

export type FieldErrors<T> = {

  [K in keyof T]?: string[];

};

  

export type ActionState<TInput, TOutput> = {

  fieldErrors?: FieldErrors<TInput>;

  error?: string | null;

  data?: TOutput;

};

  

export const createSafeAction = <TInput, TOutput>(

  schema: z.Schema<TInput>,

  handler: (validatedData: TInput) => Promise<ActionState<TInput, TOutput>>

) => {

  return async (data: TInput): Promise<ActionState<TInput, TOutput>> => {

  

    // Validate the data

    const validationResult = schema.safeParse(data);

    if (!validationResult.success) {

      return {

        fieldErrors: validationResult.error.flatten().fieldErrors as FieldErrors<TInput>,

      };

    }

  

    // Call the handler

    return handler(validationResult.data);

  

  };

};
```


## schema
```js
datasource db {

  provider     = "mysql"

  url          = env("DATABASE_URL")

  relationMode = "prisma"

}

  

generator client {

  provider = "prisma-client-js"

}

  

model Board {

  id            String @id @default(uuid())

  orgId         String

  title         String

  imageId       String

  imageThumbUrl String @db.Text

  imageFullUrl  String @db.Text

  imageUserName String @db.Text

  imageLinkHTML String @db.Text

  

  lists         List[]

  

  createdAt     DateTime @default(now())

  updatedAt     DateTime @updatedAt

}

  

model List {

  id        String @id @default(uuid())

  title     String

  order     Int

  

  boardId   String

  board     Board @relation(fields: [boardId], references: [id], onDelete: Cascade)

  

  cards     Card[]

  

  createdAt DateTime @default(now())

  updatedAt DateTime @updatedAt

  

  @@index([boardId])

}

  

model Card {

  id          String @id @default(uuid())

  title       String

  order       Int

  description String? @db.Text

  

  listId      String

  list        List  @relation(fields: [listId], references: [id], onDelete: Cascade)

  

  createdAt   DateTime @default(now())

  updatedAt   DateTime @updatedAt

  

  @@index([listId])

}

  

enum ACTION {

  CREATE

  UPDATE

  DELETE

}

  

enum ENTITY_TYPE {

  BOARD

  LIST

  CARD

}

  

model AuditLog {

  id          String  @id @default(uuid())

  orgId       String

  action      ACTION

  entityId    String

  entityType  ENTITY_TYPE

  entityTitle String

  userId      String

  userImage   String @db.Text

  userName    String @db.Text

  

  createdAt   DateTime @default(now())

  updatedAt   DateTime  @updatedAt

}

  

model OrgLimit {

  id          String @id @default(uuid())

  orgId       String @unique

  count       Int @default(0)

  

  createdAt   DateTime @default(now())

  updatedAt   DateTime  @updatedAt

}

  

model OrgSubscription {

  id                String @id @default(uuid())

  orgId             String @unique

  

  stripeCustomerId  String? @unique @map(name: "stripe_customer_id")

  stripeSubscriptionId String? @unique @map(name: "stripe_subscription_id")

  stripePriceId        String? @map(name: "stripe_price_id")

  stripeCurrentPeriodEnd DateTime? @map(name: "stripe_current_period_end")

}
```
