/**
 * 
 * 
 * 学了大半天 主要是还是觉得约定大约配置 最好是开箱即用 不要让我们去配置依赖 设定版本
 * 我们对这种事情很烦！！！！！
 * 还有很多命令可以自动个我们生成各种class 很方便

```shell
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
**/
