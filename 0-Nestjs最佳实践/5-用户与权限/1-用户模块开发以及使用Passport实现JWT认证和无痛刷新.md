用户模块开发以及使用Passport实现JWT认证和无痛刷新
## 预准备
在学习本节课前请先阅读以下文档

- [JSON Web Token 入门教程](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)
- [laravel jwt无痛刷新](https://learnku.com/articles/7264/using-jwt-auth-to-implement-api-user-authentication-and-painless-refresh-access-token)（原symfony无痛刷新这边写得很好，在思否上找不到了，随便找了一篇差不多的可以看一下）
- [nestjs auth文档](https://docs.nestjs.com/security/authentication)
## 流程图
这张流程图即为本节课的用户登录与注册流程，请仔细查看并理解
![](https://img.pincman.com/media/202303230408799.png#id=Xzxvz&originHeight=686&originWidth=1280&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
## 预装类库

- passport.js是一个专门用于Auth认证的node.js库，@nestjs/passport是一个整合passport的nestjs模块
- passport-local用于构建用户认证的本地策略，即对发送登录请求的用户名密码验证并返回jwt token给前端
- passport-jwt用于构建用户认证的jwt策略，即对需要用户使用jwt登录后才能进行操作的api来判断当前token是否有效，@nestjs/jwt用于生成access token
- jsonwebtoken用于生成refresh token，本来是不需要装这个库的，但是@nestjs/jwt里面把jsonwebtoken封装成了一个提供者单例，所以为了能生成刷新使用的refresh token而不与原来@nestjs/jwt生成的access token冲突，我们单独装这个库，并再次创建一个它的实例来生成用于刷新的令牌
- uuid用于生成一个uuid字符串
- bcrypt用于加密用户密码并对比判断密码是否正确
- fastify不用多说，哈哈
```shell
~ pnpm add @nestjs/passport passport passport-local @nestjs/jwt passport-jwt bcrypt jsonwebtoken uuid fastify
~ pnpm add @types/passport-local @types/passport-jwt @types/bcrypt @types/jsonwebtoken @types/uuid -D
```
## 用户系统
下面我们来构建我们的用户与认证模块
### 配置
首先我们需要编写一个用户模块专用的配置
类型
`UserConfig`类型的属性如下

- `hash`: 代表密码加密的散列值
- `jwt`: 是`jwt`一些相关的配置

`JwtConfig`类型的属性如下

- `secret`: 访问令牌(`access_token`)的jwt加密密匙
- `token_expired`: 访问令牌的过期时间
- `refresh_secret`: 刷新令牌(`refresh_token`)的jwt加密密匙
- `refresh_token_expired`: 刷新令牌的jwt的过期时间
```typescript
// src/modules/user/types.ts
/**
 * 用户模块配置
 */
export interface UserConfig {
    hash: number;
    jwt: JwtConfig;
}

/**
 * JWT配置
 */
export interface JwtConfig {
    secret: string;
    token_expired: number;
    refresh_secret: string;
    refresh_token_expired: number;
}
```
#### 默认配置
这些事默认配置，我们需要确保`refresh_token_expired`的时间必须长于`token_expired`，否则刷新令牌功能就无效了
```typescript
// src/modules/user/helpers.ts
export const defaultUserConfig = (configure: Configure): UserConfig => {
    return {
        hash: 10,
        jwt: {
            secret: configure.env('USER_TOKEN_SECRET', 'my-refresh-secret'),
            token_expired: configure.env('USER_TOKEN_EXPIRED', (v) => toNumber(v), 3600),
            refresh_secret: configure.env('USER_REFRESH_TOKEN_SECRET', 'my-refresh-secret'),
            refresh_token_expired: configure.env(
                'USER_REFRESH_TOKEN_EXPIRED',
                (v) => toNumber(v),
                3600 * 30,
            ),
        },
    };
};
```
#### 配置创建
编写一个配置生产函数`createUserConfig`，并在配置目录中添加用户配置
> 别忘了，需要在`config/index.ts`中导出

```typescript
// src/modules/user/helpers.ts
export const createUserConfig: (
    register: ConfigureRegister<RePartial<UserConfig>>,
) => ConfigureFactory<UserConfig> = (register) => ({
    register,
    defaultRegister: defaultUserConfig,
});

// src/config/user.config.ts
export const user = createUserConfig((configure) => ({}));
```
#### 获取配置
可以编写一个快捷方法`getUserConfig`来快速地获取用户配置中的某一项配置
```typescript
// src/modules/user/helpers.ts
export async function getUserConfig<T>(key?: string): Promise<T> {
    return App.configure.get<T>(isNil(key) ? 'user' : `user.${key}`);
}
```
### 创建模块
编写用户模块引导类
```typescript
// src/modules/user/user.module.ts
@ModuleBuilder(async (configure) => ({
    imports: [PassportModule],
    providers: [],
    exports: [],
}))
export class UserModule {}
```
### 密码函数
编写两个工具函数，`encrypt`用于加密密码，`decrypt`用于对比明文和加密后的密码，即验证密码正确性
```typescript
// src/modules/user/helpers.ts
export const encrypt = async (password: string) => {
    return bcrypt.hashSync(password, (await getUserConfig<number>('hash')) || 10);
};

export const decrypt = (password: string, hashed: string) => {
    return bcrypt.compareSync(password, hashed);
};
```
### 模型
本节课的数据结构比较简单直白，就三个表，用户表，访问令牌表，刷新令牌表，它们的关系如下

- 用户表和访问令牌表形成一对多的关系，即一个用户可以在多个端点使用多个令牌登录
- 刷新令牌表与访问令牌表形成一对一的关系

其余字段都有注释，请自行查看模型代码
#### 用户模型
```typescript
// src/modules/user/entities/user.entity.ts
@Exclude()
@Entity('users')
export class UserEntity extends BaseEntity {
    ...
    @OneToMany(() => AccessTokenEntity, (accessToken) => accessToken.user, {
        cascade: true,
    })
    accessTokens!: AccessTokenEntity[];
}
```
#### Token基类
```typescript
// src/modules/user/entities/base.token.ts
@Exclude()
export abstract class BaseToken extends BaseEntity {
    @Column({ length: 500 })
    value!: string;

    @Column({
        comment: '令牌过期时间',
    })
    expired_at!: Date;

    @CreateDateColumn({
        comment: '令牌创建时间',
    })
    createdAt!: Date;
}
```
#### 访问Token
```typescript
// src/modules/user/entities/access-token.entity.ts
@Entity('user_access_tokens')
export class AccessTokenEntity extends BaseToken {
    @OneToOne(() => RefreshTokenEntity, (refreshToken) => refreshToken.accessToken, {
        cascade: true,
    })
    refreshToken!: RefreshTokenEntity;

    @ManyToOne((type) => UserEntity, (user) => user.accessTokens, {
        onDelete: 'CASCADE',
    })
    user!: UserEntity;
}
```
#### 刷新Token
```typescript
// src/modules/user/entities/refresh-token.entity.ts
@Entity('user_refresh_tokens')
export class RefreshTokenEntity extends BaseToken {
    @OneToOne(() => AccessTokenEntity, (accessToken) => accessToken.refreshToken, {
        onDelete: 'CASCADE',
    })
    @JoinColumn()
    accessToken!: AccessTokenEntity;
}
```
### 存储类
当前用户模块的Repository非常简单，只是写了个默认按用户创建时间排序
```typescript
// src/modules/user/repositories/user.repository.ts
@CustomRepository(UserEntity)
export class UserRepository extends BaseRepository<UserEntity> {
    protected _qbName = 'user';

    buildBaseQuery() {
        return this.createQueryBuilder(this.qbName).orderBy(`${this.qbName}.createdAt`, 'DESC');
    }
}
```
### 订阅者
`UserSubscriber`类中有几个比较重要的方法

- `generateUserName`方法用于自动生成一个随机用户名（本节课用不到），后续通过短信或邮件注册会用到
- `beforeInsert`用于在插入用户前判断是否有用户名或密码，没有则自动生成（本节课用不到），并对明文密码加密
- `beforeUpdate`用于在更新密码时判断密码是否被更新，如果密码更新则加密新密码
```typescript
// src/modules/user/subscribers/user.subscriber.ts
@EventSubscriber()
export class UserSubscriber extends BaseSubscriber<UserEntity> {
    protected entity = UserEntity;
    protected async generateUserName(event: InsertEvent<UserEntity>): Promise<string> {
        const username = `user_${randomBytes(4).toString('hex').slice(0, 8)}`;
        const user = await event.manager.findOne(UserEntity, {
            where: { username },
        });
        return !user ? username : this.generateUserName(event);
    }

    async beforeInsert(event: InsertEvent<UserEntity>) {
        if (!event.entity.username) {
            event.entity.username = await this.generateUserName(event);
        }
        if (!event.entity.password) {
            event.entity.password = randomBytes(11).toString('hex').slice(0, 22);
        }
        event.entity.password = await encrypt(event.entity.password);
    }

    async beforeUpdate(event: UpdateEvent<UserEntity>) {
        if (this.isUpdated('password', event)) {
            event.entity.password = encrypt(event.entity.password);
        }
    }
}
```
### DTO验证
此处的DTO验证类用于针对用户数据的CRUD请求的验证
#### 常量
添加两个枚举常量

- `UserOrderType`用于设置获取的用户列表的排序
- `UserValidateGroups`用于设置DTO属性的不同验证组
```typescript
// src/modules/user/constants.ts
export enum UserOrderType {
    CREATED = 'createdAt',
    UPDATED = 'updatedAt',
}

export enum UserValidateGroups {
    CREATE = 'user-create',
    UPDATE = 'user-update',
}
```
#### 公共DTO
首先我们编写一个公共的DTO类，在这里设置一些使用率比较高的验证属性，其它的DTO可以通过选取或排除的方法来继承这个类中的某些属性，关于属性的具体验证实现请自行查看代码，这些内容前面章节已经基本讲过，不再赘述
```typescript
// src/modules/user/dtos/common.dto.ts
@Injectable()
export class UserCommonDto {
    ...
}
```
#### 创建请求
创建用户时包含

- 用户名(必填)
- 昵称(可选)
- 密码(必填)
- 手机号(可选)
- 邮箱(可选)
```typescript
// src/modules/user/dtos/user.dto.ts
@DtoValidation({ groups: [UserValidateGroups.CREATE] })
export class CreateUserDto extends PickType(UserCommonDto, [
    'username',
    'nickname',
    'password',
    'phone',
    'email',
]) {}
```
#### 更新请求
```typescript
// src/modules/user/dtos/user.dto.ts
@DtoValidation({ groups: [UserValidateGroups.UPDATE] })
export class UpdateUserDto extends PartialType(CreateUserDto) {
    @ApiProperty({
        description: '待更新的用户ID',
    })
    @IsUUID(undefined, { groups: [UserValidateGroups.UPDATE], message: '用户ID格式不正确' })
    @IsDefined({ groups: ['update'], message: '用户ID必须指定' })
    id!: string;
}
```
#### 查询请求
查询用户列表时，相对于基础的列表查询验证，多了个用于指定排序的`orderBy`属性
```typescript
// src/modules/user/dtos/user.dto.ts
@DtoValidation({ type: 'query' })
export class QueryUserDto extends ListWithTrashedQueryDto {
    @ApiPropertyOptional({
        description: '排序规则:可指定用户列表的排序规则,默认为按创建时间降序排序',
        enum: UserOrderType,
    })
    @IsEnum(UserOrderType)
    orderBy?: UserOrderType;
}
```
### 服务
用户CRUD的服务类有几个重要的方法
> 其它的方法请自行查看代码

- `findOneByCredential`: 通过凭证来查找一个用户，凭证可以是用户名，邮箱或者手机号
- `findOneByCondition`: 通过任意键值对对象来查找一个用户
```typescript
// src/modules/user/services/user.service.ts
/**
 * 用户管理服务
 */
@Injectable()
export class UserService extends BaseService<UserEntity, UserRepository> {
   ...

    /**
     * 根据用户用户凭证查询用户
     * @param credential
     * @param callback
     */
    async findOneByCredential(credential: string, callback?: QueryHook<UserEntity>) {
        let query = this.userRepository.buildBaseQuery();
        if (callback) {
            query = await callback(query);
        }
        return query
            .where('user.username = :credential', { credential })
            .orWhere('user.email = :credential', { credential })
            .orWhere('user.phone = :credential', { credential })
            .getOne();
    }


    async findOneByCondition(condition: { [key: string]: any }, callback?: QueryHook<UserEntity>) {
        let query = this.userRepository.buildBaseQuery();
        if (callback) {
            query = await callback(query);
        }
        const wheres = Object.fromEntries(
            Object.entries(condition).map(([key, value]) => [key, value]),
        );
        const user = query.where(wheres).getOne();
        if (!user) {
            throw new EntityNotFoundError(UserEntity, Object.keys(condition).join(','));
        }
        return user;
    }

}
```
### 控制器
控制器与前面课程的内容一样写，不再赘述
```typescript
// src/modules/user/controllers/user.controller.ts
@ApiTags('用户管理')
@Depends(UserModule)
@Crud(() => ({
    id: 'user',
    enabled: [
        { name: 'list', option: createHookOption('用户查询,以分页模式展示') },
        { name: 'detail', option: createHookOption('用户详情') },
        { name: 'store', option: createHookOption('新增用户') },
        { name: 'update', option: createHookOption('修改用户信息') },
        { name: 'delete', option: createHookOption('删除用户') },
        { name: 'restore', option: createHookOption('恢复用户') },
    ],
    dtos: {
        list: QueryUserDto,
        store: CreateUserDto,
        update: UpdateUserDto,
    },
}))
@Controller('users')
export class UserController extends BaseControllerWithTrash<UserService> {
    constructor(protected userService: UserService) {
        super(userService);
    }
}
```
## 账户与认证
第一部分我们介绍了用户表的CRUD，基本与前面课程的内容是一致的。这节的重点是这部分
### DTO验证
添加几个新的DTO
#### 更改密码请求
更改密码，需要填入旧密码，以及新密码，并且重复填写新密码确认无误才能进行更改
> 后续我们会添加一个在忘记旧密码的情况下通过邮件或短信来更改密码的功能

```typescript
// src/modules/user/dtos/auth.dto.ts
export class UpdatePasswordDto extends PickType(UserCommonDto, ['password', 'plainPassword']) {
    @ApiProperty({
        description: '旧密码:用户在更改密码时需要输入的原密码',
        minLength: 8,
        maxLength: 50,
    })
    @IsPassword(5, {
        message: '密码必须由小写字母,大写字母,数字以及特殊字符组成',
        always: true,
    })
    @Length(8, 50, {
        message: '密码长度不得少于$constraint1',
        always: true,
    })
    oldPassword!: string;
}
```
#### 登录凭证
在公共DTO类中添加一个登录凭证属性，登录请求时填入的登录凭证不一定是用户名，也可以是邮箱地址或者手机号，所以做字符串验证即可
```typescript
// src/modules/user/dtos/common.dto.ts
export class UserCommonDto {
    @ApiProperty({
        description: '登录凭证:可以是用户名,手机号,邮箱地址',
        minLength: 4,
        maxLength: 20,
    })
    @Length(4, 30, {
        message: '登录凭证长度必须为$constraint1到$constraint2',
        always: true,
    })
    @IsNotEmpty({ message: '登录凭证不得为空', always: true })
    readonly credential!: string;
    ...
}
```
#### 登录请求
添加一个`CredentialDto`类，用于对通过"凭证+密码"的方式登录进行验证
```typescript
// src/modules/user/dtos/auth.dto.ts
export class CredentialDto extends PickType(UserCommonDto, ['credential', 'password']) {}
```
#### 注册请求
新增一个注册请求验证组
```typescript
// src/modules/user/constants.ts
export enum UserValidateGroups {
    CREATE = 'user-create',
    UPDATE = 'user-update',
    REGISTER = 'user-register',
}
```
添加一个注册验证DTO类
```typescript
// src/modules/user/dtos/auth.dto.ts
@DtoValidation({ groups: [UserValidateGroups.REGISTER] })
export class RegisterDto extends PickType(UserCommonDto, [
    'username',
    'nickname',
    'password',
    'plainPassword',
] as const) {}
```
### 服务
用户认证需要我们添加两个新的服务类对数据进行操作
#### 类型
`JwtPayload`类型用于设置JWT的荷载对象
```typescript
export interface JwtPayload {
    sub: string;
    iat: number;
}
```
#### 令牌服务
该服务用于操作令牌

- `refreshToken`方法用于属性出新的访问令牌并存储，然后把新的访问令牌放置到响应的`header`里返回给前端存储
- `generateAccessToken`方法用于生成一个新的访问令牌以及对应的刷新令牌存储到数据库
- `generateRefreshToken`方法用于生成一个新的刷新令牌
- `checkAccessToken`方法用于检查一个访问令牌是否存在于数据库
- `removeAccessToken`方法用于移除访问令牌且自动移除关联的刷新令牌
- `removeRefreshToken`方法用于移除访问令牌
- `verifyAccessToken`方法用于验证令牌是否有效
```typescript
// src/modules/user/services/token.service.ts
/**
 * 令牌服务
 */
@Injectable()
export class TokenService {
    constructor(protected readonly jwtService: JwtService) {}

    async refreshToken(accessToken: AccessTokenEntity, response: Response) {
        const { user, refreshToken } = accessToken;
        if (refreshToken) {
            const now = await getTime();
            // 判断refreshToken是否过期
            if (now.isAfter(refreshToken.expired_at)) return null;
            // 如果没过期则生成新的access_token和refresh_token
            const token = await this.generateAccessToken(user, now);
            await accessToken.remove();
            response.header('token', token.accessToken.value);
            return token;
        }
        return null;
    }

    async generateAccessToken(user: UserEntity, now: dayjs.Dayjs) {
        const config = await getUserConfig<JwtConfig>('jwt');
        const accessTokenPayload: JwtPayload = {
            sub: user.id,
            iat: now.unix(),
        };

        const signed = this.jwtService.sign(accessTokenPayload);
        const accessToken = new AccessTokenEntity();
        accessToken.value = signed;
        accessToken.user = user;
        accessToken.expired_at = now.add(config.token_expired, 'second').toDate();
        await accessToken.save();
        const refreshToken = await this.generateRefreshToken(accessToken, await getTime());
        return { accessToken, refreshToken };
    }

    async generateRefreshToken(
        accessToken: AccessTokenEntity,
        now: dayjs.Dayjs,
    ): Promise<RefreshTokenEntity> {
        const config = await getUserConfig<JwtConfig>('jwt');
        const refreshTokenPayload = {
            uuid: uuid(),
        };
        const refreshToken = new RefreshTokenEntity();
        refreshToken.value = jwt.sign(refreshTokenPayload, config.refresh_secret);
        refreshToken.expired_at = now.add(config.refresh_token_expired, 'second').toDate();
        refreshToken.accessToken = accessToken;
        await refreshToken.save();
        return refreshToken;
    }

    async checkAccessToken(value: string) {
        return AccessTokenEntity.findOne({
            where: { value },
            relations: ['user', 'refreshToken'],
        });
    }

    async removeAccessToken(value: string) {
        const accessToken = await AccessTokenEntity.findOne({
            where: { value },
        });
        if (accessToken) await accessToken.remove();
    }
  
    async removeRefreshToken(value: string) {
        const refreshToken = await RefreshTokenEntity.findOne({
            where: { value },
            relations: ['accessToken'],
        });
        if (refreshToken) {
            if (refreshToken.accessToken) await refreshToken.accessToken.remove();
            await refreshToken.remove();
        }
    }

    async verifyAccessToken(token: AccessTokenEntity) {
        const config = await getUserConfig<JwtConfig>('jwt');
        const result = jwt.verify(token.value, config.secret);
        if (!result) return false;
        return token.user;
    }
}
```
#### Auth服务
该类包含以下方法

- `validateUser`方法用于验证凭证(可以是用户名/邮箱/密码)以及密码是否正确
- `login`方法用于登录用户，即生成访问和刷新两个令牌
- `logout`方法用于登出用户，即把当前登录的访问令牌（包含刷新令牌）从数据库中移除
- `createToken`方法用于给某个用户生成新的令牌
- `register`方法用于注册新用户
- `updatePassword`方法用于更新密码
```typescript
// src/modules/user/services/auth.service.ts
@Injectable()
export class AuthService {
    constructor(
        private readonly userService: UserService,
        private readonly tokenService: TokenService,
    ) {}

    async validateUser(credential: string, password: string) {
        const user = await this.userService.findOneByCredential(credential, async (query) =>
            query.addSelect('user.password'),
        );
        if (user && decrypt(password, user.password)) {
            return user;
        }
        return false;
    }

    async login(user: UserEntity) {
        const now = await getTime();
        const { accessToken } = await this.tokenService.generateAccessToken(user, now);
        return accessToken.value;
    }

    async logout(req: Request) {
        const accessToken = ExtractJwt.fromAuthHeaderAsBearerToken()(req as any);
        if (accessToken) {
            await this.tokenService.removeAccessToken(accessToken);
        }

        return {
            msg: 'logout_success',
        };
    }
 
    async createToken(id: string) {
        const now = await getTime();
        let user: UserEntity;
        try {
            user = await this.userService.detail(id);
        } catch (error) {
            throw new ForbiddenException();
        }
        const { accessToken } = await this.tokenService.generateAccessToken(user, now);
        return accessToken.value;
    }

    async register(data: RegisterDto) {
        const { username, nickname, password } = data;
        const user = await this.userService.create({
            username,
            nickname,
            password,
            actived: true,
        } as any);
        return this.userService.findOneByCondition({ id: user.id });
    }
  
    async updatePassword(user: UserEntity, { password, oldPassword }: UpdatePasswordDto) {
        const item = await this.userRepository.findOneOrFail({
            select: ['password'],
            where: { id: user.id },
        });
        if (decrypt(item.password, oldPassword))
            throw new ForbiddenException('old password not matched');
        item.password = password;
        await this.userRepository.save(item);
        return this.userService.findOneByCondition({ id: item.id });
    }
}
```
### 本地策略
本地策略的作用在于根据传入的用户名密码来判断该登录是否有效
#### 策略
我们需要再构造函数里更改一下验证字段的映射，然后验证凭证及密码是否正确，正确的话就返回该用户
```typescript
// src/modules/user/strategies/local.strategy.ts
export class LocalStrategy extends PassportStrategy(Strategy) {
    constructor(private readonly authService: AuthService) {
        super({
            usernameField: 'credential',
            passwordField: 'password',
        });
    }

    async validate(credential: string, password: string): Promise<any> {
        const user = await this.authService.validateUser(credential, password);
        if (!user) {
            throw new UnauthorizedException();
        }
        return user;
    }
}
```
#### 守卫
该本地策略守卫通过获取获取请求`body`中的凭证和密码然后通过`CredentialDto`进行格式验证，如果验证通过，则继续把请求数据传入上面的本地策略进行正确性验证
```typescript
// src/modules/user/guards/local-auth.guard.ts
@Injectable()
export class LocalAuthGuard extends AuthGuard('local') {
    async canActivate(context: ExecutionContext) {
        const request = context.switchToHttp().getRequest();
        try {
            await validateOrReject(plainToClass(CredentialDto, request.body), {
                validationError: { target: false },
            });
        } catch (errors) {
            const messages = (errors as any[])
                .map((e) => e.constraints ?? {})
                .reduce((o, n) => ({ ...o, ...n }), {});
            throw new BadRequestException(Object.values(messages));
        }
        return super.canActivate(context) as boolean;
    }
}
```
#### 请求用户
在本地策略验证成功后，passport会自动把用户对象追加到请求中，这时我们可以加一个装饰器`ReqUser`来便捷的获取登录的用户
```typescript
// src/modules/user/decorators/user-request.decorator.ts
export const ReqUser = createParamDecorator(async (_data: unknown, ctx: ExecutionContext) => {
    const request = ctx.switchToHttp().getRequest();
    return request.user as ClassToPlain<UserEntity>;
});
```
#### 控制器
`login`方法的步骤如下

1. 通过本地守卫`@UseGuards(LocalAuthGuard)`来验证请求格式，该守卫会自动调用本地策略对当前用户的正确性进行验证，验证后会把用户对象添加到请求中
2. 通过本地策略验证后的用户对象已经添加到`request`中，我们可以通过`@ReqUser`装饰器获取该用户，并为该该用户创建一个jwt令牌返回给前端
```typescript
// src/modules/user/controllers/auth.controller.ts
@ApiTags('账户操作')
@Depends(UserModule)
@Controller('auth')
export class AuthController {
    constructor(protected readonly authService: AuthService) {}

    @Post('login')
    @ApiOperation({ summary: '用户通过凭证(可以是用户名,邮箱,手机号等)+密码登录' })
    @UseGuards(LocalAuthGuard)
    async login(@ReqUser() user: ClassToPlain<UserEntity>, @Body() _data: CredentialDto) {
        return { token: await this.authService.createToken(user.id) };
    }
}
```
### JWT策略
通过本地策略获取jwt的key后，前端需要存储起来（比如放到`localstorage`里面），然后下次访问时，在`header`上带上这个令牌就可以访问需要登录后才能访问的资源了
#### 常量
该常量用于设置一个API端点是否允许匿名访问
```typescript
// src/modules/restful/constants.ts
export const ALLOW_GUEST = 'allowGuest';
```
#### 策略
JWT策略通过从前端传入的令牌中自动解析出来的荷载来验证并查询用户，把查询出来的用户对象进行`plain`处理为普通对象，以方便在请求中传输。该对象会像本地策略一样追加到请求头中，所以我们一样可以通过`@ReqUser`装饰器拿到该登录用户
```typescript
// src/modules/user/strategies/jwt.strategy.ts
@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
    constructor(private readonly userRepository: UserRepository) {
        super({
            jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
            ignoreExpiration: false,
            secretOrKey: getUserConfig('jwt.secret'),
        });
    }

    async validate(payload: JwtPayload) {
        const user = await this.userRepository.findOneOrFail({ where: { id: payload.sub } });
        return instanceToPlain(user);
    }
}
```
#### 更改CRUD
为了使前面编写的CRUD框架构建的控制器中继承`BaseController`的方法可以自如的设置登录限制，我们为`CrudMethodOption`接口添加一个新的属性`allowGuest`，然后更改一下`registerCrud`让它为需要允许匿名访问的方法添加一个`ALLOW_GUEST`常量并设置成`true`
```typescript
// src/modules/restful/types.ts
export interface CrudMethodOption {
    /**
     * 该方法是否允许匿名访问
     */
    allowGuest?: boolean;
    ...
}
    

// src/modules/restful/crud.ts
export const registerCrud = async <T extends BaseController<any> | BaseControllerWithTrash<any>>(
    Target: Type<T>,
    options: CrudOptions,
) => {
        ...
        if (option.allowGuest) Reflect.defineMetadata(ALLOW_GUEST, true, Target.prototype, name);
        if (!isNil(option.hook)) option.hook(Target, name);
    }
};
```
#### Guest装饰器
同时为了支持非`BaseController`中的方法也能控制登录访问权限，我们添加一个`@Guest`装饰器
```typescript
// src/modules/user/decorators/guest.decorator.ts
export const Guest = () => SetMetadata(ALLOW_GUEST, true);
```
#### Jwt模块
还需要导入和注册`PassportModule`以及`JwtModule`才能使用passport策略。我们为开发环境设置一下"令牌永不过期"
```typescript
// src/modules/user/user.module.ts
const jwtModuleRegister = (configure: Configure) => async (): Promise<JwtModuleOptions> => {
    const config = await configure.get<UserConfig>('user');
    const isProd = configure.getRunEnv() === EnvironmentType.PRODUCTION;
    const option: JwtModuleOptions = {
        secret: config.jwt.secret,
        verifyOptions: {
            ignoreExpiration: !isProd,
        },
    };
    if (isProd) option.signOptions.expiresIn = `${config.jwt.token_expired}s`;
    return option;
};

@ModuleBuilder(async (configure) => ({
    imports: [
        PassportModule,
        JwtModule.registerAsync({
            useFactory: jwtModuleRegister(configure),
        }),
    ],
}))
export class UserModule {}
```
#### 守卫
添加一个JWT守卫用于检查用户登录状态，其它的不再赘述。最重要的讲一下无痛刷新，即在访问令牌过期的情况下，判断一下刷新令牌是否过期，如果也过期就直接`401`了，如果没过期，则用刷新令牌刷新出新的令牌，并把新的访问令牌放入到请求头中再次验证（此次一般是必定通过了的），同时把新的请求令牌放在响应头中，以便前端能替换掉旧的访问令牌
```typescript
// src/modules/user/guards/jwt-auth.guard.ts
@Injectable()
export class JwtAuthGuard extends AuthGuard('jwt') {
    constructor(protected reflector: Reflector, protected tokenService: TokenService) {
        super();
    }

    /**
     * 守卫方法
     * @param context
     */
    async canActivate(context: ExecutionContext) {
        const crudGuest = Reflect.getMetadata(
            ALLOW_GUEST,
            context.getClass().prototype,
            context.getHandler().name,
        );
        const defaultGuest = this.reflector.getAllAndOverride<boolean>(ALLOW_GUEST, [
            context.getHandler(),
            context.getClass(),
        ]);
        const allowGuest = crudGuest ?? defaultGuest;
        const request = this.getRequest(context);
        const response = this.getResponse(context);
        // if (!request.headers.authorization) return false;
        // 从请求头中获取token
        // 如果请求头不含有authorization字段则认证失败
        const requestToken = ExtractJwt.fromAuthHeaderAsBearerToken()(request);
        if (isNil(requestToken) && !allowGuest) return false;
        // 判断token是否存在,如果不存在则认证失败
        const accessToken = isNil(requestToken)
            ? undefined
            : await this.tokenService.checkAccessToken(requestToken!);
        if (isNil(accessToken) && !allowGuest) throw new UnauthorizedException();
        try {
            // 检测token是否为损坏或过期的无效状态,如果无效则尝试刷新token
            const result = await super.canActivate(context);
            if (allowGuest) return true;
            return result as boolean;
        } catch (e) {
            // 尝试通过refreshToken刷新token
            // 刷新成功则给请求头更换新的token
            // 并给响应头添加新的token和refreshtoken
            if (!isNil(accessToken)) {
                const token = await this.tokenService.refreshToken(accessToken, response);
                if (isNil(token) && !allowGuest) return false;
                if (token.accessToken) {
                    request.headers.authorization = `Bearer ${token.accessToken.value}`;
                }
                // 刷新失败则再次抛出认证失败的异常
                const result = await super.canActivate(context);
                if (allowGuest) return true;
                return result as boolean;
            }

            return allowGuest;
        }
    }

    /**
     * 自动请求处理
     * 如果请求中有错误则抛出错误
     * 如果请求中没有用户信息则抛出401异常
     * @param err
     * @param user
     * @param _info
     */
    handleRequest(err: any, user: any, _info: Error) {
        if (err || !user) {
            throw err || new UnauthorizedException();
        }
        return user;
    }

    protected getRequest(context: ExecutionContext) {
        return context.switchToHttp().getRequest();
    }

    protected getResponse(context: ExecutionContext) {
        return context.switchToHttp().getResponse();
    }
}
```
#### 全局守卫
最后我们需要更改一下应用创建器，以便能接收全局的守卫。我们在创建应用时把全局守卫换成`JwtAuthGuard`，这样就能让所有的控制器API端点都必须登录才能操作了（除非加了`@Guest`装饰器，或者CRUD框架方法的话，在`option`中设置了`allowGuest`为`true`）
```typescript
// src/modules/core/types.ts
export interface CreateOptions {
    globals?: {
       ...
        guard?: Type<IAuthGuard>;
    };
  
// src/modules/core/helpers/app.ts
export async function createBootModule(
    params: AppParams,
    options: Pick<Partial<CreateOptions>, 'meta' | 'modules' | 'globals'>,
): Promise<{ BootModule: Type<any>; modules: ModuleBuildMap }> {
    ...
    if (globals.filter !== null) {
        providers.push({
            provide: APP_FILTER,
            useClass: AppFilter,
        });
    }
    if (!isNil(globals.guard)) {
        providers.push({
            provide: APP_GUARD,
            useClass: globals.guard,
        });
    }
    ...
}
    
// src/creator.ts
export const creator = createApp({
    configs,
    configure: { storage: true },
    modules: [ContentModule],
    globals: { guard: JwtAuthGuard },
    ...
});
```
### 控制器
有了以上的准备，接下来我们就能编写端点啦！
#### 更改Hook
首先我们更改一下`createHookOption`这个方法，以便给那些不允许匿名登录的端点的swagger文档显示时添加上一把锁，即`ApiBearerAuth`方法
```typescript
// src/modules/restful/helpers.ts

export function createHookOption(
    option: { guest?: boolean; summary?: string } | string = {},
): CrudMethodOption {
    const params = typeof option === 'string' ? { summary: option } : option;
    const { guest: allowGuest, summary } = params;
    return {
        allowGuest,
        hook: (target, method) => {
            if (!isNil(summary))
                ApiOperation({ summary })(
                    target,
                    method,
                    Object.getOwnPropertyDescriptor(target.prototype, method),
                );
            if (!allowGuest) {
                ApiBearerAuth()(
                    target,
                    method,
                    Object.getOwnPropertyDescriptor(target.prototype, method),
                );
            }
        },
    };
}
```
#### 账户控制器
账户控制器主要是一些用户自身的操作，比如登录，登出等
```typescript
// src/modules/user/controllers/auth.controller.ts

export class AuthController {
    constructor(protected readonly authService: AuthService) {}

    @Post('login')
    @ApiOperation({ summary: '用户通过凭证(可以是用户名,邮箱,手机号等)+密码登录' })
    @Guest()
    @UseGuards(LocalAuthGuard)
    async login(@ReqUser() user: ClassToPlain<UserEntity>, @Body() _data: CredentialDto) {
        return { token: await this.authService.createToken(user.id) };
    }

    /**
     * 注销登录
     * @param req
     */
    @Post('logout')
    @ApiOperation({ summary: '用户登出账户' })
    @ApiBearerAuth()
    async logout(@Request() req: any) {
        return this.authService.logout(req);
    }

    /**
     * 使用用户名密码注册
     * @param data
     */
    @Post('register')
    @ApiOperation({ summary: '通过用户名+密码注册账户' })
    @Guest()
    async register(
        @Body()
        data: RegisterDto,
    ) {
        return this.authService.register(data);
    }
    /**
     * 更改密码
     * @param user
     * @param data
     */
    @Patch('reset-passowrd')
    @ApiOperation({ summary: '重置密码' })
    @SerializeOptions({
        groups: ['user-detail'],
    })
    async resetPassword(
        @ReqUser() user: ClassToPlain<UserEntity>,
        @Body() data: UpdatePasswordDto,
    ): Promise<UserEntity> {
        return this.authService.updatePassword(user, data);
    }
}
```
#### 用户控制器
用户控制器则用于用户管理的CRUD操作
```typescript
// src/modules/user/controllers/user.controller.ts

@ApiTags('用户管理')
@Depends(UserModule)
@Crud(() => ({
    id: 'user',
    enabled: [
        {
            name: 'list',
            option: createHookOption({ summary: '用户查询,以分页模式展示', guest: true }),
        },
        { name: 'detail', option: createHookOption({ summary: '用户详情', guest: true }) },
      ...
```
#### 内容控制器
我们更改一下前面的内容模块中的分类，文章和评论控制器，因为一些操作是允许匿名的，比如读取分类，文章，评论等
```typescript
// src/modules/content/controllers

@ApiTags('分类')
@Depends(ContentModule)
@Crud(async () => ({
    id: 'category',
    enabled: [
        {
            name: 'list',
            option: createHookOption({ summary: '分类查询,以分页模式展示', guest: true }),
        },
        {
            name: 'detail',
            option: createHookOption({ summary: '分类详情', guest: true }),
        },
       ...
     
@ApiTags('文章')
@Depends(ContentModule)
@Crud(async () => ({
    id: 'post',
    enabled: [
        {
            name: 'list',
            option: createHookOption({ summary: '文章查询,以分页模式展示', guest: true }),
        },
        {
            name: 'detail',
            option: createHookOption({ summary: '文章查询,以分页模式展示', guest: true }),
        },
       ...
      
@ApiTags('评论')
@Depends(ContentModule)
@Crud(async () => ({
    id: 'comment',
    enabled: [
        {
            name: 'list',
            option: createHookOption({ summary: '评论查询,以分页模式展示', guest: true }),
        },
        {
            name: 'detail',
            option: createHookOption({ summary: '评论详情', guest: true }),
        },
```
## 收尾工作
代码编写工作接近尾声，我们来编写一下用户模块并注册
### 编写模块
```typescript
// src/modules/user/user.module.ts
@ModuleBuilder(async (configure) => ({
    imports: [
        PassportModule,
        JwtModule.registerAsync({
            useFactory: jwtModuleRegister(configure),
        }),
        await addEntities(configure, Object.values(entities)),
        DatabaseModule.forRepository(Object.values(repositories)),
    ],
    providers: [
        ...Object.values(services),
        ...(await addSubscribers(configure, Object.values(subscribers))),
        ...Object.values(strategies),
        ...Object.values(guards),
    ],
    exports: [
        ...Object.values(services),
        DatabaseModule.forRepository(Object.values(repositories)),
    ],
}))
export class UserModule {}
```
### 注册模块
```typescript
// src/creator.ts
export const creator = createApp({
    configs,
    configure: { storage: true },
    modules: [UserModule, ContentModule],
    ...
});
```
### 注册控制器
```typescript
// src/routes/v1.ts
import * as contentControllers from '@/modules/content/controllers';
import { Configure } from '@/modules/core/configure';
import { ApiVersionOption } from '@/modules/restful/types';
import * as userControllers from '@/modules/user/controllers';

export const v1 = async (configure: Configure): Promise<ApiVersionOption> => ({
    routes: [
        {
            name: 'app',
            path: '/',
            controllers: [],
            doc: {
                title: '应用接口',
                description: 'CMS系统的应用接口',
                tags: [
                    { name: '分类', description: '分类的增删查改操作' },
                    { name: '文章', description: '文章的增删查改操作' },
                    { name: '评论', description: '评论的增删查操作' },
                    { name: '账户操作', description: 'Auth操作' },
                    { name: '用户管理', description: '用户的增删查改操作' },
                ],
            },
            children: [
                {
                    name: 'content',
                    path: 'content',
                    controllers: Object.values(contentControllers),
                },
                {
                    name: 'user',
                    path: 'user',
                    controllers: Object.values(userControllers),
                },
            ],
        },
    ],
});
```
## 命令
为了有一些初始的用户数据，我们来生成一些模拟的用户数据
### 数据工厂
```typescript
// src/database/factories/user.factory.ts
export type IUserFactoryOptions = Partial<{
    [key in keyof UserEntity]: UserEntity[key];
}>;
export const UserFactory = defineFactory(
    UserEntity,
    async (faker: Faker, settings: IUserFactoryOptions = {}) => {
        faker.setLocale('zh_CN');
        const user = new UserEntity();
        const optionals: (keyof IUserFactoryOptions)[] = ['username', 'password', 'email', 'phone'];
        optionals.forEach((key) => {
            if (settings[key] !== undefined) {
                user[key] = settings[key] as never;
            }
        });
        user.nickname = settings.nickname ?? faker.name.fullName();
        return user;
    },
);
```
### 数据填充
```typescript
// src/database/seeders/user.seeder.ts
export default class UserSeeder extends BaseSeeder {
    protected truncates = [AccessTokenEntity, RefreshTokenEntity, UserEntity];

    protected factorier!: DbFactory;

    async run(_factorier: DbFactory, _dataSource: DataSource, _em: EntityManager): Promise<any> {
        this.factorier = _factorier;
        await this.loadUsers();
    }

    private async loadUsers() {
        const repository = getCustomRepository(this.dataSource, UserRepository);
        const userFactory = this.factorier(UserEntity);
        const creator = await repository.findOneBy({ username: 'pincman' });
        if (isNil(creator)) {
            await userFactory<IUserFactoryOptions>({
                username: 'pincman',
                nickname: 'pincman',
                password: '123456aA$',
            }).create({}, 'username');
        }
        const count = await repository.count();
        if (count < 2) {
            await userFactory<IUserFactoryOptions>({
                username: 'xiaoming',
                nickname: '小明',
                phone: '+86.18605853847',
                password: '123456aA$',
            }).create({}, 'username');

            await userFactory<IUserFactoryOptions>({
                username: 'lige',
                nickname: '李哥',
                phone: '+86.15955959999',
                password: '123456aA$',
            }).create({}, 'username');
            await userFactory<IUserFactoryOptions>().createMany(15, {}, 'username');
        }
    }
}
```
### 运行命令
```bash
~ pnpm dbmg
~ pnpm dbmr -s
```
