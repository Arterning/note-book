# nest cli

crud 生成器

```jsx
nest g resource users
```

```bash
npm i -g @nestjs/cli

# create a new nest project
nest new

# run it!
npm run start

# in dev mode
npm run start:dev

# generate controller
nest generate controller
nest g co

# generate service
nest generate service name
nest g s name
nest g module name

#generate dto class
nest g class name/dto/create-name.dto --no-spec

# validate
npm i class-validator class-transformer
npm i @nestjs/mapped-types

# type orm
npm i @nestjs/typeorm typeorm pg
```