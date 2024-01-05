
https://replicate.com/


Run and fine-tune open-source models. Deploy custom models at scale. All with one line of code.


DashBoard
https://replicate.com/


We use replicate to generate video and audio

Other we use OpenAI



Schema is simple
```js
generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "mysql"
  url = env("DATABASE_URL")
  relationMode = "prisma"
}

model UserApiLimit {
  id        String      @id @default(cuid())
  userId    String   @unique
  count     Int      @default(0)
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt
}

model UserSubscription {
  id        String      @id @default(cuid())
  userId    String   @unique
  stripeCustomerId       String?   @unique @map(name: "stripe_customer_id")
  stripeSubscriptionId   String?   @unique @map(name: "stripe_subscription_id")
  stripePriceId          String?   @map(name: "stripe_price_id")
  stripeCurrentPeriodEnd DateTime? @map(name: "stripe_current_period_end")
}

```