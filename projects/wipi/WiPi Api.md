This is CMS System


## Basic
we should set the database url in the `app.module.js`
```js
TypeOrmModule.forRootAsync({
      imports: [ConfigModule],
      inject: [ConfigService],
      useFactory: async (configService: ConfigService) => ({
        type: 'mysql',
        entities: [User, File, Knowledge, Article, Category, Tag, Comment, Setting, SMTP, Page, View, Search],
        host: configService.get('DB_HOST', 'localhost'),
        port: configService.get<number>('DB_PORT', 3306),
        username: configService.get('DB_USER', 'root'),
        password: configService.get('DB_PASSWD', '123456'),
        database: configService.get('DB_DATABASE', 'wipi'),
        charset: 'utf8mb4',
        timezone: '+08:00',
        synchronize: true,
      }),
    }),
```

main.js is jsut like the main method in java springboot, has bootstrap function
```js

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.enableCors();
  app.setGlobalPrefix(app.get('ConfigService').get('SERVER_API_PREFIX', '/api'));
  app.use(
    rateLimit({
      windowMs: 60 * 1000, // 1 minutes
      max: 10000, // limit each IP to 1000 requests per windowMs
    })
  );
  app.use(compression()); // 启用 gzip 压缩
  app.use(helmet());
  app.useGlobalInterceptors(new TransformInterceptor()); // 正常情况下，响应值统一
  app.useGlobalFilters(new HttpExceptionFilter()); // 异常情况下，响应值统一
  app.use(bodyParser.json({ limit: '10mb' })); // 修改请求的容量
  app.use(bodyParser.urlencoded({ limit: '10mb', extended: true }));
  const swaggerConfig = new DocumentBuilder()
    .setTitle('Wipi Open Api')
    .setDescription('Wipi Open Api Document')
    .setVersion('1.0')
    .build();
  const document = SwaggerModule.createDocument(app, swaggerConfig);
  SwaggerModule.setup('api', app, document);

  await app.listen(app.get('ConfigService').get('SERVER_PORT', 3003));
  console.log('[wipi] 服务启动成功');
}

bootstrap();
```

auth module

i can not understand it much....TODO
```js
import { forwardRef, Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { PassportModule } from '@nestjs/passport';

import { SettingModule } from '../setting/setting.module';
import { SMTPModule } from '../smtp/smtp.module';
import { UserModule } from '../user/user.module';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { JwtStrategy } from './jwt.strategy';

const passModule = PassportModule.register({ defaultStrategy: 'jwt' });
const jwtModule = JwtModule.register({
  secret: 'wipi',
  signOptions: { expiresIn: '4h' },
});

@Module({
  imports: [
    forwardRef(() => UserModule),
    passModule,
    jwtModule,
    forwardRef(() => SettingModule),
    forwardRef(() => SMTPModule),
  ],
  providers: [AuthService, JwtStrategy],
  controllers: [AuthController],
  exports: [passModule, jwtModule],
})
export class AuthModule {}

```

common exception handler
```jsx
import { ArgumentsHost, Catch, ExceptionFilter, HttpException, HttpStatus } from '@nestjs/common';

import { errorLogger } from '../logger';

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  // eslint-disable-next-line class-methods-use-this
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse();
    const request = ctx.getRequest();

    const url = request.originalUrl; // 请求路由
    const status = exception instanceof HttpException ? exception.getStatus() : HttpStatus.INTERNAL_SERVER_ERROR;
    const msg = exception.message.message || exception.message; // 错误信息
    const errorResponse = {
      statusCode: status,
      msg,
      success: false,
      data: null,
    };

    // 设置返回的状态码、请求头、发送错误信息
    response.status(status);
    response.header('Content-Type', 'application/json; charset=utf-8');
    response.send(errorResponse);
    errorLogger.error(url, errorResponse);
  }
}

```


forward ref is used to fix the cycle dependency

在NestJS中，`forwardRef`是一个用于处理循环依赖的工具函数。循环依赖是指两个或多个模块相互依赖，可能会导致循环引用的问题。使用`forwardRef`可以解决这种情况。

在NestJS中，通常通过`@Inject()` 装饰器来注入依赖项。但是，如果两个模块相互依赖，直接在构造函数中使用 `@Inject()` 可能会导致循环引用错误。

使用`forwardRef`，你可以延迟对模块或提供者的引用，从而避免循环依赖。以下是一个示例：

```typescript
import { Injectable, forwardRef, Inject } from '@nestjs/common';
import { SomeService } from './some.service';

@Injectable()
export class AnotherService {
  constructor(@Inject(forwardRef(() => SomeService)) private someService: SomeService) {
    // 在构造函数中使用SomeService
  }
}
```

在上面的例子中，`forwardRef(() => SomeService)` 告诉NestJS在解析依赖项时要等待 `SomeService` 的引用。这样可以避免循环引用问题。请注意，`@Inject()` 装饰器中使用 `forwardRef` 只是在解决循环依赖时的一种方式，具体的解决方案可能会根据应用程序的结构和需求而有所不同。




## File Upload 
```js
/**
   * 上传文件
   * @param file
   */
  async uploadFile(file, unique): Promise<File> {
    const { originalname, mimetype, size, buffer } = file;
    const filename =
      +unique === 1
        ? `/${dateFormat(new Date(), 'yyyy-MM-dd')}/${uniqueid()}/${originalname}`
        : `/${dateFormat(new Date(), 'yyyy-MM-dd')}/${originalname}`;
    const url = await this.oss.putFile(filename, buffer);
    const newFile = await this.fileRepository.create({
      originalname,
      filename,
      url,
      type: mimetype,
      size,
    });
    await this.fileRepository.save(newFile);
    return newFile;
  }
```

first upload file to oss , then save the file info to db


```js
import { HttpException, HttpStatus } from '@nestjs/common';

import { SettingService } from '../modules/setting/setting.service';
import { AliyunOssClient } from './oss/aliyun-oss-client';
import { MiniOssClient } from './oss/minio-client';
import { OssClient } from './oss/oss-client';

export class Oss {
  settingService: SettingService;
  config: Record<string, unknown>;
  ossClient: OssClient;

  constructor(settingService: SettingService) {
    this.settingService = settingService;
  }

  private async getConfig() {
    const data = await this.settingService.findAll(true);
    const config = JSON.parse(data.oss);
    if (!config) {
      throw new HttpException('OSS 配置不完善，无法进行操作', HttpStatus.BAD_REQUEST);
    }
    return config as Record<string, unknown>;
  }

  private async getOssClient() {
    const config = await this.getConfig();
    const type = String(config.type).toLowerCase();

    switch (type) {
      case 'minio':
        return new MiniOssClient(config);
      case 'aliyun':
      default:
        return new AliyunOssClient(config);
    }
  }

  async putFile(filepath: string, buffer: ReadableStream) {
    const client = await this.getOssClient();
    const url = await client.putFile(filepath, buffer);
    return url;
  }

  async deleteFile(filepath: string, url: string) {
    const client = await this.getOssClient();
    await client.deleteFile(filepath, url);
  }
}

```

I use the minio client, so 
```js
import * as Minio from 'minio';

import minioConfig from './minio.config';
import { OssClient } from './oss-client';

export class MiniOssClient extends OssClient {
  /**
   * 删除文件
   * @param filepath
   * @param url
   */
  async deleteFile(filepath: string, url: string): Promise<void> {
    const client = this.buildClient();
    const bucketName = (this.config.bucket as string) || 'xiaohui';
    const objectName = url.split('/')[url.split('/').length - 1];
    await client.removeObject(bucketName, objectName, {});
    return Promise.resolve(undefined);
  }

  /**
   * 上传文件
   * @param filepath
   * @param buffer
   */
  async putFile(filepath: string, buffer: ReadableStream): Promise<string> {
    const client = this.buildClient();
    try {
      const ext = filepath.split('.')[1];
      const bucketName = (this.config.bucket as string) || 'xiaohui';
      const objectName = bucketName + '_' + new Date().getTime() + '.' + ext;
      await client.putObject(bucketName, objectName, buffer);
      return `http://${minioConfig.endPoint}:${minioConfig.port}/${bucketName}/${objectName}`;
    } catch (error) {
      throw new Error(`Minio upload error: ${error}`);
    }
  }

  private buildClient() {
    const config = this.config;
    const client = new Minio.Client({
      endPoint: 'localhost',
      port: 9000,
      useSSL: false,
      accessKey: (config.accessKeyId as string) || 'minioadmin',
      secretKey: (config.accessKeySecret as string) || 'minioadmin',
    });
    return client;
  }
}

```


## Setting Service
```js

@Injectable()
export class SettingService {
  constructor(
    @InjectRepository(Setting)
    private readonly settingRepository: Repository<Setting>
  ) {
//after system inital, it will read the josn files and save to db
    this.initI18n();
  }

  /**
   * 初始化时加载 i18n 配置
   */
  async initI18n() {
    const items = await this.settingRepository.find();
    const target = (items && items[0]) || ({} as Setting);
    let data = {};
    try {
      data = JSON.parse(target.i18n);
    } catch (e) {
      data = {};
    }
    target.i18n = JSON.stringify(merge({}, i18n, data));
    await this.settingRepository.save(target);
  }
..
}
```
where is the i18n
is in the config modules
and in the config module's index.js, we can find
```js
export * from './env';
export * from './i18n';

```
the i18n file have zh and en json object
```js
{
  "zh": "Chinese",
  "en": "English",
  "serverNotAvaliable": "The server is going to launch a rocket for Musk now 🚀",
  "pageMissing": "This page went to space with Bezos",
  "archives": "Archives",
  "total": "Total",
  ...
```

```js
{
  "zh": "汉语",
  "en": "英文",
  "serverNotAvaliable": "服务器暂时去给马斯克发射火箭去了🚀",
  "pageMissing": "页面和贝佐斯去太空旅行了~~",
  "archives": "归档",
  "total": "共计",
```


