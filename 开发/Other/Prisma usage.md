
use `npx prisma migrate dev --name alterUser` to generate migration file and sync db
use `npx prisma db push ` to just sync db , but not generate migration file..

use `npx prisma migrate deploy` in production env



```bash
◭  Prisma is a modern DB toolkit to query, migrate and model your database (https://prisma.io)

Usage

  $ prisma [command]

Commands

            init   Set up Prisma for your app
        generate   Generate artifacts (e.g. Prisma Client)
              db   Manage your database schema and lifecycle
         migrate   Migrate your database
          studio   Browse your data with Prisma Studio
        validate   Validate your Prisma schema
          format   Format your Prisma schema

Flags

     --preview-feature   Run Preview Prisma commands

Examples

  Set up a new Prisma project
  $ prisma init

  Generate artifacts (e.g. Prisma Client)
  $ prisma generate

  Browse your data
  $ prisma studio

  Create migrations from your Prisma schema, apply them to the database, generate artifacts (e.g. Prisma Client)
  $ prisma migrate dev

  Pull the schema from an existing database, updating the Prisma schema
  $ prisma db pull

  Push the Prisma schema state to the database
  $ prisma db push

  Validate your Prisma schema
  $ prisma validate

  Format your Prisma schema
  $ prisma format
```

