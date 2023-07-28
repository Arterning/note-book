---
aliases: [1-Yargs构建命令行工具以及数据迁移的实现]
tags: 
cssclass:
source:
date created: 星期四, 五月 25日 2023, 4:46:30 凌晨
date modified: 星期六, 五月 27日 2023, 12:51:06 下午
---
## 预准备
在学习本节课之前你需要阅读以下文档

- [yargs官网文档](http://yargs.js.org/docs/)或[3R教室的yargs翻译](https://pincman.com/docs/yargs/usaged)(由**诗侬**同学加紧翻译中)
- [typeorm的迁移使用文档](https://typeorm.io/migrations)
- [ora v5的文档](https://github.com/sindresorhus/ora/tree/v5.4.1)
## 预装库

- `yargs`是一个构建命令行的工具
- `ora`是一个在命令行中输出雪碧动画的库，这可以用来示意一些耗时运行
- `@sqltools/formatter`这个库我们无需关注
- `cross-env`用于跨平台指定环境变量

```bash
~ pnpm add yargs @sqltools/formatter ora@5
~ pnpm add cross-env @types/yargs -D
```

## 核心模块
在编写命令行工具之前我们需要对核心模块就行一些修改
### 类型
新增及修改以下类型

- `ReRequired`: 该类型用于定义深度嵌套对象，且每个键值对必须是必选的
- `CommandItem`: 该类型用于定义一个命令构造器函数，它接收一个`AppParams`类型的参数，并返回`yargs`的`CommandModule`类型的命令模块
- `CommandCollection`: 类型构造器函数的数组集合
- 在原来的`CreateOptions`类型中新增一个`CommandCollection`类型的属性`commands`。这使我们不仅可以在模块上添加命令，还可以在构建应用时直接在`createApp`中传入一些额外的命令，比如我们可以新增一个`commands`目录用于存放模块之外的一些当前应用的特殊命令
- 在原来的`creatorData`类型上添加`CommandCollection`类型的`commands`属性。这使得`createApp`函数的返回结果中可以带有所有命令模块（包括自定义传入的和模块中的命令）
- 在原来的`ModuleBuilderMeta`类型上添加`CommandCollection`类型的`commands`属性。这样就可以在模块构造器的选项参数中添加每个模块专属的命令了

```typescript
// typings/global.d.ts
declare type ReRequired<T> = {
    [P in keyof T]-?: T[P] extends (infer U)[] | undefined
        ? ReRequired<U>[]
        : T[P] extends object | undefined
        ? T[P] extends ((...args: any[]) => any) | ClassType<T[P]> | undefined
            ? T[P]
            : ReRequired<T[P]>
        : T[P];
};

// src/modules/database/types.ts
export type CommandItem<T = Record<string, any>, U = Record<string, any>> = (
    params: Required<AppParams>,
) => CommandModule<T, U>;

export type CommandCollection = Array<CommandItem<any, any>>;

export interface CreateOptions {
    ...
    commands?: CommandCollection;
}

export interface CreatorData extends Required<AppParams> {
    modules: ModuleBuildMap;
    commands: CommandCollection;
}

export type ModuleBuilderMeta = ModuleMetadata & {
    global?: boolean;
    commands?: CommandCollection;
};
  
@ModuleBuilder(async (configure) => {
    const imports: ModuleMetadata['imports'] = [];

    if (!configure.has('database')) {
        panic({ message: 'Database config not exists or not right!' });
    }
    ...
    };
})
```

### 读取命令
`createImportModules`函数中构造模块时排除`commands`，防止`commands`属性被附加到模块的metadata参数中

```typescript
// src/modules/core/helpers/app.ts
async function createImportModules(
    configure: Configure,
    modules: ModuleItem[],
): Promise<ModuleBuildMap> {
    for (const m of modules) {
        ...
        Module(omit(metadata, ['global', 'commands']))(option.module);
    }
    return maps;
}
```

为`createApp`函数（即`App`类的`create`方法）参数添加自定义的`commands`的接收以及返回值中添加最终返回的`commands`值（包含前面的传入参数中的自定义的`commands`以及每个模块中的`commands`）

```typescript
// src/modules/core/app.ts
export class App {
    static async create(options: CreateOptions): Promise<CreatorData> {
        const { builder, configs, configure, commands = [] } = options;
        ...
        return { configure: this._configure, app: this._app, modules, commands };
    }
}
```

### panic函数
此函数用于标识命令运行失败，其参数类型如下

```typescript
// src/modules/core/types.ts
export interface PanicOption {
    /**
     * 报错消息
     */
    message: string;
    /**
     * ora对象
     */
    spinner?: Ora;
    /**
     * 抛出的异常信息
     */
    error?: any;
    /**
     * 是否退出进程
     */
    exit?: boolean;
}
```

逻辑代码如下
> 这个函数的逻辑代码非常简单，不再赘述，请自行理解

```typescript
// src/modules/core/helpers/utils.ts
export function panic(option: PanicOption | string) {
    console.log();
    if (typeof option === 'string') {
        console.log(chalk.red(`\n❌ ${option}`));
        process.exit(1);
    }
    const { error, spinner, message, exit = true } = option;
    if (error) console.log(chalk.red(error));
    spinner ? spinner.fail(chalk.red(`\n❌${message}`)) : console.log(chalk.red(`\n❌ ${message}`));
    if (exit) process.exit(1);
}
```

更改应用，使一些错误可以直接使用`panic`函数处理

```typescript
// src/modules/core/app.ts
export class App {    
    static async create(options: CreateOptions): Promise<CreatorData> {
        try {
            ...
        } catch (error) {
            panic({ message: 'Create app failed!', error });
            exit(0);
        }

        return { configure: this._configure, app: this._app, modules, commands };
    }
}
 
// src/modules/database/database.module.ts
@ModuleBuilder(async (configure) => {
    const imports: ModuleMetadata['imports'] = [];

    if (!configure.has('database')) {
        panic({ message: 'Database config not exists or not right!' });
    }
    ...
})
export class DatabaseModule {
}
```

### 构建CLI
在构建应用之前，我们需要编写一个函数读取所有自定义和模块中的命令
需要注意的是，我们对每个命令的处理器函数需要进行二次封装，以便在执行命令后关闭应用并退出进程

```typescript
// src/modules/core/helpers/app.ts
export async function createCommands(params: CreatorData): Promise<CommandModule<any, any>[]> {
    const { app, modules } = params;
    const moduleCommands: Array<CommandItem<any, any>> = Object.values(modules)
        .map((m) => m.meta.commands ?? [])
        .reduce((o, n) => [...o, ...n], []);
    const commands = [...params.commands, ...moduleCommands].map((item) => {
        const command = item(params);
        return {
            ...command,
            handler: async (args: yargs.Arguments<any>) => {
                const handler = command.handler as (
                    ...argvs: yargs.Arguments<any>
                ) => Promise<void>;
                await handler({ ...params, ...args });
                await app.close();
                process.exit();
            },
        };
    });
    return commands;
}
```

然后使用yargs编写一个命令构建函数（把原来的`creator`函数从`main.ts`中抽出来放置到一个单独的文件中，以免重复编写）
`buildCli`函数的实现逻辑如下

1. 根据传入的应用构建函数构建应用，并得到所有命令
2. 遍历每个命令并绑定到`yargs`对象
3. 使用`yargs`的API构建CLI

```typescript
// src/creator.ts
export const creator = createApp({
    ...
});

// src/main.ts
bootApp(creator, ({ app, configure }) => async () => {
    const restful = app.get(Restful);
    echoApi(configure, restful);
});

// src/modules/core/helpers/app.ts
async function buildCli(builder: () => Promise<CreatorData>) {
    const params = await builder();
    const commands = await createCommands(params);
    console.log();
    commands.forEach((command) => yargs.command(command));
    yargs
        .usage('Usage: $0 <command> [options]')
        .scriptName('cli')
        .demandCommand(1, '')
        .fail((msg, err, y) => {
            if (!msg && !err) {
                yargs.showHelp();
                process.exit();
            }
            if (msg) console.error(chalk.red(msg));
            if (err) console.error(chalk.red(err.message));
            process.exit();
        })
        .strict()
        .alias('v', 'version')
        .help('h')
        .alias('h', 'help').argv;
}
buildCli(creator);
```

### 运行命令
要运行CLI命令，我们需要明白两点

- 在开发环境下我们需要再TS环境下运行命令，因为我们不可能编译一次再运行一次，这样就非常麻烦了，在生产环境下是没有源码的，所以命令必须是支持加载JS模块的
- 一些命令只能在某个环境下运行，比如生成迁移文件只能在开发环境下运行，而运行迁移则在开发和生产环境下都能运行

为了解决以上两个问题，我们需要用以下的方式来编写命令运行工具
首先编写一个js文件来调用命令
> 如果感兴趣，可以把tsnode换成swc或者esbuild试试，可以提高性能哦

1. 此js文件中必须判断上层有没有用于编译TS代码的`tsconfig.json`文件（本项目中为`tsconfig.build.json`）
2. 如果有，证明是在源码环境中，那么我们要运行命令，必须先使用tsnode或者swc这类工具转义ts代码然后运行
3. 如果没有，则在已编译环境中，直接运行即可

```javascript
// src/cli.js
#!/usr/bin/env node
const { existsSync } = require('fs');
const { join } = require('path');

const projectPath = join(__dirname, '../tsconfig.build.json');
if (existsSync(projectPath)) {
    require('ts-node').register({
        files: true,
        transpileOnly: true,
        project: projectPath,
    });
    require('tsconfig-paths/register');
}

const { creator } = require('./creator');
const { buildCli } = require('./modules/core/helpers/app');

buildCli(creator);
```

然后我们就可以来尝试运行该命令了
在源码环境下使用`./src/cli.js -h`或者`node ./src/cli.js -h`查看命令列表（加不加`-h`都可以），在编译文件中（比如生成环境下只部署了`dist`目录），可以直接`./cli.js -h`或`node ./cli.js -h`运行命令
如果不想加`node`前缀，请使用`chmod +x ./src/cli.js`来设置一下可执行权限
为了在源码环境下可以通过`pnpm`来运行cli命令，同时也为了解决命令中的运行环境问题， 需要使用`cross-env`在`scripts`中设置一下环境变量
> 因为在`Configure`类中我们把默认环境设置为`production`，所以在生产环境下运行`./cli.js xxx`时不需要手动设置运行环境了

```json
// package.json
{
    ...
    "scripts": {
        "cli": "cross-env NODE_ENV=development ./src/cli.js",
        "prebuild": "rimraf dist",
        "build": "rimraf dist && nest build",
        "format": "prettier --write \"src/**/*.ts\" \"test/**/*.ts\"",
        "start": "cross-env NODE_ENV=development nest start",
        "start:dev": "cross-env NODE_ENV=development nest start --watch",
        "start:debug": "cross-env NODE_ENV=development nest start --debug --watch",
        "start:prod": "node dist/main",
        "lint": "eslint \"{src,apps,libs,test}/**/*.ts\" --fix",
        "test": "cross-env NODE_ENV=test jest",
        "test:watch": "cross-env NODE_ENV=test jest --watch",
        "test:cov": "cross-env NODE_ENV=test jest --coverage",
        "test:debug": "cross-env NODE_ENV=development node --inspect-brk -r tsconfig-paths/register -r ts-node/register node_modules/.bin/jest --runInBand",
        "test:e2e": "cross-env NODE_ENV=test jest --config ./test/jest-e2e.json"
    },
}
```

后续再开发环境中就可以直接使用`pnpm cli xxx`来运行命令了
## 表结构迁移
前面我的应用是使用数据库同步的方式来更改数据库表结构，也就是每次更改模型后一启动应用就会自动同步最新的模型生成的表结构到数据库。这种做法其实是非常不安全的，尤其在生产环境下，很容易导致我们的数据丢失。所以我们需要使用一种较为安全的方式来修改表结构-即数据结构迁移！
默认的Typeorm虽然支持迁移，但是没有同Nestjs进行整合的，包括动态配置，命令整合等都无。不过有了我们前面章节的基础之后，现在我们就可以直接通过修改Typeorm的默认迁移命令来深度整合我们的课程框架里的Nestjs应用啦，接下来我们学习一下如何实现！
### 永久关闭同步
在实现迁移命令之前，我们首先需要永久地关闭数据库同步功能，以防迁移出错

```typescript
// src/modules/database/helpers.ts
export const createDbOptions = (configure: Configure, options: DbConfigOptions) => {
        ...
        return deepMerge(
            newOptions.common,
            {
                ...newOption,
                synchronize: false,
                autoLoadEntities: true,
            } as any,
           ...
};
```

### CLI中加载模型和订阅者
由于

- 迁移用到的数据库配置中无法自动加载模型和订阅者等 
- 后续我们需要对模型可以进行更多的定制，比如添加动态关联等 

所以在编写迁移之前，我们先来编写添加模型和订阅者的函数

- `addEntities`函数会根据传入的`dataSource`来判断把模型追加到哪个连接池，并使用`TypeOrmModule.forFeature`为Nestjs中的TypeormModule自动注册这些模型
- `addSubscribers`函数会根据传入的`dataSource`来判断把订阅者追加到哪个连接池

```typescript
// src/modules/database/helpers.ts
export const addEntities = async (
    configure: Configure,
    entities: EntityClassOrSchema[] = [],
    dataSource = 'default',
) => {
    const database = await configure.get<DbConfig>('database');
    if (isNil(database)) throw new Error(`Typeorm have not any config!`);
    const dbConfig = database.connections.find(({ name }) => name === dataSource);
    // eslint-disable-next-line prettier/prettier, prefer-template
    if (isNil(dbConfig)) throw new Error(`Database connection named ${dataSource} not exists!`);
    const oldEntities = (dbConfig.entities ?? []) as ObjectLiteral[];
    /**
     * 更新数据库配置,添加上新的模型
     */
    configure.set(
        'database.connections',
        database.connections.map((connection) =>
            connection.name === dataSource
                ? {
                      ...connection,
                      entities: [...entities, ...oldEntities],
                  }
                : connection,
        ),
    );
    return TypeOrmModule.forFeature(entities, dataSource);
};

export const addSubscribers = async (
    configure: Configure,
    subscribers: Type<any>[] = [],
    dataSource = 'default',
) => {
    // 写法与模型类似
    return subscribers;
};
```

### 修改订阅者
因为在迁移时，数据库配置中无法直接使用依赖注入，比如订阅者中注入数据连接池或者其它服务等，所以需要修改一下订阅者
首先我们修改一下订阅者基类，使得`dataSource`变成可选依赖，这样在CLI加载这个订阅者（Subscriber会被自动实例化）的时候就不会出现因为丢失`dataSource`依赖了

```typescript
// src/modules/database/base/subscriber.ts
@EventSubscriber()
export abstract class BaseSubscriber<E extends ObjectLiteral>
    implements EntitySubscriberInterface<E>
{
  ...
   constructor(@Optional() protected dataSource?: DataSource) {
       if (!isNil(this.dataSource)) this.dataSource.subscribers.push(this);
   }
}
```

接下来修改`PostSubscriber`，使`SanitizeService`通过`app`实例获取，而不是直接注入

```typescript
// src/modules/content/subscribers/post.subscriber.ts
@EventSubscriber()
export class PostSubscriber extends BaseSubscriber<PostEntity> {
    protected entity = PostEntity;

    listenTo() {
        return PostEntity;
    }

    /**
     * 加载文章数据的处理
     * @param entity
     */
    async afterLoad(entity: PostEntity) {
        const sanitizeService = App.app.get(SanitizeService, { strict: false });
        if (entity.type === PostBodyType.HTML && !isNil(sanitizeService)) {
            entity.body = sanitizeService.sanitize(entity.body);
        }
    }
}
```

然后修改一下模块的加载方式

```typescript
// src/modules/content/content.module.ts
@ModuleBuilder(async (configure) => {
    const searchType = await configure.get<SearchType>('content.searchType', 'against');
    const providers: ModuleMetadata['providers'] = [
        ...Object.values(services),
        ...(await addSubscribers(configure, [PostSubscriber])),
        ...
    return {
        imports: [
            await addEntities(configure, Object.values(entities)),
            DatabaseModule.forRepository(Object.values(repositories)),
        ],
        ...
    };
})
export class ContentModule {}
```

### 公共类型

1. 在实现迁移之前我们先为TypeORM的数据库配置添加上一个`path`类型的属性，用于设置`migrations`以及其它文件的路径。然后把这个`path`放在一个新的类型`DbAdditionalOption`里面，再为`TypeormOption`，`DbConfigOptions`以及`DbConfig`的`common`属性添加上该类型
2. 所有命令的参数选项都继承自`yargs.Arguments`，而数据库命令都是需要设置链接的，所以在此基础上添加上`connection`选项，得到一个公共的数据库命令参数类型`TypeOrmArguments`

```typescript
// src/modules/database/types.ts
/**
 * 自定义数据库配置
 */
export type DbConfigOptions = {
    common: Record<string, any> & ReRequired<DbAdditionalOption>;
    connections: Array<TypeOrmModuleOptions>;
};

/**
 * 最终数据库配置
 */
export type DbConfig = Record<string, any> & {
    common: Record<string, any> & ReRequired<DbAdditionalOption>;
    connections: TypeormOption[];
};

/**
 * Typeorm连接配置
 */
export type TypeormOption = Omit<TypeOrmModuleOptions, 'name' | 'migrations'> & {
    name: string;
} & Required<DbAdditionalOption>;

/**
 * 额外数据库选项,用于CLI工具
 */
type DbAdditionalOption = {
    paths?: {
        /**
         * 迁移文件路径
         */
        migration?: string;
    };
};

/**
 * 基础数据库命令参数类型
 */
export type TypeOrmArguments = yargs.Arguments<{
    connection?: string;
}>;
```

### 生成迁移
首先来完成一个自动生成迁移文件的命令
#### 辅助函数
添加一个辅助函数用于自动生成迁移文件名和类名

```typescript
// src/modules/core/helpers/utils.ts
/**
 * 生成只包含字母的固定长度的字符串
 * @param length
 */
export const getRandomCharString = (length: number) => {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
    const charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
};
```

#### 类型
命令的参数类型为`TypeOrmArguments`这个数据库命令的公共参数类型结合该命令特有的`MigrationGenerateOptions`类型的联合类型
`MigrationGenerateOptions`代表的参数选项如下

- `connection`: 数据库连接名称，默认为`default` ，默认: `default`
- `name`: 手动指定迁移文件的类名（即文件名）, 默认: 自动生成
- `pretty`: 是否打印生成的迁移文件所要执行的SQL，默认：`false`
- `dryrun`: 是否只打印生成的迁移文件的内容而不是直接生成迁移文件，默认: `false`
- `check`: 是否只验证数据库是最新的而不是直接生成迁移，默认: `false`

```typescript
// src/modules/database/types.ts
export type MigrationGenerateArguments = TypeOrmArguments & MigrationGenerateOptions;

export interface MigrationGenerateOptions {
    name?: string;
    run?: boolean;
    pretty?: boolean;
    outputJs?: boolean;
    dryrun?: boolean;
    check?: boolean;
}
```

#### 执行逻辑
所有执行逻辑参考TypeOrm的生成迁移的源码修改，请参考[此处](https://github.com/typeorm/typeorm/blob/master/src/commands/MigrationGenerateCommand.ts)以及课程代码

```typescript
// src/modules/database/commands/typeorm-migration-generate.ts
export class TypeormMigrationGenerate {
    async handler(args: HandlerOptions) {
      ...
    }
}
```

#### 处理器
命令处理器的执行步骤如下

1. 判断当前运行环境是否为生产环境（生产环境下只能运行迁移，而不需要生成迁移文件），如果是则抛出异常并结束命令进程
2. 运行一次迁移（运行迁移命令请看下一部分）以便把前面没有运行的迁移先同步到数据库，这是为了防止重复生成迁移文件
3. 在指定目录中生成迁移文件
4. 如果是`run`参数为`true`的话，再次运行迁移以便直接把生成的最新迁移同步到数据库

```typescript
// src/modules/database/commands/migration-generate.handler.ts
export const MigrationGenerateHandler = async (
    configure: Configure,
    args: yargs.Arguments<MigrationGenerateArguments>,
) => {
    if (configure.getRunEnv() === EnvironmentType.PRODUCTION) {
        panic('Migration generate command can not run in production environment!');
    }
    await MigrationRunHandler(configure, { connection: args.connection } as any);
    console.log();
    const spinner = ora('Start to generate migration');
    const cname = args.connection ?? 'default';
    try {
        spinner.start();
        console.log();
        const { connections = [] }: DbConfig = await configure.get<DbConfig>('database');
        const dbConfig = connections.find(({ name }) => name === cname);
        if (isNil(dbConfig)) panic(`Database connection named ${cname} not exists!`);
        console.log();
        const dir = dbConfig.paths.migration ?? resolve(__dirname, '../../../database/migrations');
        const runner = new TypeormMigrationGenerate();
        // console.log(((dbConfig.entities ?? []) as ClassType<any>[]).map((e) => e.name));
        // process.exit();
        const dataSource = new DataSource({ ...dbConfig } as DataSourceOptions);
        console.log();
        await runner.handler({
            name: args.name ?? getRandomCharString(6),
            dir,
            dataSource,
            ...pick(args, ['pretty', 'outputJs', 'dryrun', 'check']),
        });
        if (dataSource.isInitialized) await dataSource.destroy();
        spinner.succeed(chalk.greenBright.underline('\n 👍 Finished generate migration'));
        if (args.run) {
            console.log();
            await MigrationRunHandler(configure, { connection: args.connection } as any);
        }
    } catch (error) {
        panic({ spinner, message: 'Generate migration failed!', error });
    }
};
```

#### 命令编写
要使用上面的处理器则必须编写一个命令构造器，其代码如下

```typescript
// src/modules/database/commands/migration-generate.command.ts
export const GenerateMigrationCommand: CommandItem<any, MigrationGenerateArguments> = ({
    configure,
}) => ({
    source: true,
    command: ['db:migration:generate', 'dbmg'],
    describe: 'Generates a new migration file with sql needs to be executed to update schema.',
    builder: {
        connection: ...,
        name: ...,
        dir: ...,
        pretty: ...,
        dryrun: ...,
        check: ...,
    } as const,

    handler: async (args: yargs.Arguments<MigrationGenerateArguments>) =>
        MigrationGenerateHandler(configure, args),
});

// src/modules/database/commands/index.ts
export * from './migration-generate.command';
```

#### 命令注册
最后我们在数据库模块的`@ModuleBuilder`中添加所有的命令

```typescript
// src/modules/database/database.module.ts
import * as commands from './commands';

@ModuleBuilder(async (configure) => {
    ...
    return {
        global: true,
        commands: Object.values(commands),
        imports,
        providers,
    };
})
```

### 运行迁移
运行迁移命令可以把生成的迁移文件的内容同步到数据库中，这样我们每次在更改模型时生成一个新的迁移文件，然后运行一下迁移来使数据库永远保持最新状态
#### 类型
运行迁移命令的参数类型除了公共类型`TypeOrmArguments`外，还继承了回滚迁移命令参数的`MigrationRevertOptions`类型，具体如下

- `connection`: 链接名称，默认为`default`
- `transaction`: 运行迁移的事务名称，默认为`default`
- `fake`：如果数据库表结构已经被更改，是否需要模拟运行迁移，默认为`false`
- `refresh`: 是否刷新数据库，即删除所有表结构后重新运行（生产环境下不可用），默认为`false`
- `onlydrop`：只删除数据库表结构而不运行（生产环境下不可用），默认为`false`

```typescript
// src/modules/database/types.ts
/**
 * 运行迁移的命令参数
 */
export type MigrationRunArguments = TypeOrmArguments & MigrationRunOptions;

/**
 * 运行迁移处理器选项
 */
export interface MigrationRunOptions extends MigrationRevertOptions {
    refresh?: boolean;
    onlydrop?: boolean;
    clear?: boolean;
}

/**
 * 恢复迁移处理器选项
 */
export interface MigrationRevertOptions {
    transaction?: string;
    fake?: boolean;
}
```

#### 执行逻辑
所有执行逻辑参考TypeOrm的运行迁移的源码修改，请参考[此处](https://github.com/typeorm/typeorm/blob/master/src/commands/MigrationRunCommand.ts)以及课程代码

```typescript
// src/modules/database/commands/tyeporm-migration-run.ts
type HandlerOptions = MigrationRunOptions & { dataSource: DataSource };
export class TypeormMigrationRun {
    async handler({ transaction, fake, dataSource }: HandlerOptions) {
        const options = {
            transaction:
                dataSource.options.migrationsTransactionMode ?? ('all' as 'all' | 'none' | 'each'),
            fake,
        };
        switch (transaction) {
            case 'all':
                options.transaction = 'all';
                break;
            case 'none':
            case 'false':
                options.transaction = 'none';
                break;
            case 'each':
                options.transaction = 'each';
                break;
            default:
            // noop
        }

        await dataSource.runMigrations(options);
    }
}
```

#### 处理器
运行迁移的处理器逻辑如下

1. 判断是否需要先删除数据库表结构（包括刷新或者只删除操作），如果需要删除数据库则必须确保在生产环境下运行，否则抛出异常并终止进程，如果在其他环境下则正常删除数据库表结构
2. 连接数据库并运行迁移

```typescript
// src/modules/database/commands/migration-run.handler.ts
export const MigrationRunHandler = async (
    configure: Configure,
    args: yargs.Arguments<MigrationRunArguments>,
) => {
    const isProd = configure.getRunEnv() === EnvironmentType.PRODUCTION;
    const spinner = ora('Start to run migrations');
    const cname = args.connection ?? 'default';
    let dataSource: DataSource | undefined;
    try {
        spinner.start();
        const { connections = [] }: DbConfig = await configure.get<DbConfig>('database');
        const dbConfig = connections.find(({ name }) => name === cname);
        if (isNil(dbConfig)) panic(`Database connection named ${cname} not exists!`);
        let dropSchema = false;
        if (isProd && (args.refresh || args.onlydrop)) {
            panic('Migration refresh database schema can not run in production environment!');
        }
        dropSchema = args.refresh || args.onlydrop;
        console.log();
        const dir = dbConfig.paths.migration ?? resolve(__dirname, '../../../database/migrations');
        const runner = new TypeormMigrationRun();
        dataSource = new DataSource({ ...dbConfig } as DataSourceOptions);
        if (dataSource && dataSource.isInitialized) await dataSource.destroy();
        const options = {
            subscribers: [],
            synchronize: false,
            migrationsRun: false,
            dropSchema,
            logging: ['error'],
            migrations: [join(dir, isProd ? '**/*.js' : '**/*.ts')],
        } as any;
        if (dropSchema) {
            dataSource.setOptions(options);
            await dataSource.initialize();
            await dataSource.destroy();
            spinner.succeed(chalk.greenBright.underline('\n 👍 Finished drop database schema'));
            if (args.onlydrop) process.exit();
        }
        dataSource.setOptions({ ...options, dropSchema: false });
        await dataSource.initialize();
        console.log();
        await runner.handler({
            dataSource,
            transaction: args.transaction,
            fake: args.fake,
        });
        spinner.succeed(chalk.greenBright.underline('\n 👍 Finished run migrations'));
    } catch (error) {
        if (dataSource && dataSource.isInitialized) await dataSource.destroy();
        panic({ spinner, message: 'Run migrations failed!', error });
    } finally {
        if (dataSource && dataSource.isInitialized) await dataSource.destroy();
    }
};
```

#### 命令
最后编写一下命令并注册

```typescript
// src/modules/database/commands/migration-run.command.ts
export const RunMigrationCommand: CommandItem<any, MigrationRunArguments> = ({ configure }) => ({
    source: true,
    command: ['db:migration:run', 'dbmr'],
    describe: 'Runs all pending migrations.',
    builder: {
        connection: {
            type: 'string',
            alias: 'c',
            describe: 'Connection name of typeorm to connect database.',
        },
        transaction: {
            type: 'string',
            alias: 't',
            describe:
                'Indicates if transaction should be used or not for migration run/revert/reflash. Enabled by default.',
            default: 'default',
        },
        fake: {
            type: 'boolean',
            alias: 'f',
            describe:
                'Fakes running the migrations if table schema has already been changed manually or externally ' +
                '(e.g. through another project)',
        },
        refresh: {
            type: 'boolean',
            alias: 'r',
            describe: 'drop database schema and run migration',
            default: false,
        },

        onlydrop: {
            type: 'boolean',
            alias: 'o',
            describe: 'only drop database schema',
            default: false,
        },
    } as const,

    handler: async (args: yargs.Arguments<MigrationRunArguments>) =>
        MigrationRunHandler(configure, args),
});
```

### 创建迁移和回滚迁移
这两个命令与生成以前和运行迁移大同小异，此处不再赘述，具体代码请自行参考课程源码
需要注意的是

- 手动创建迁移命令`dbmc`一般用于对比当前迁移文件得到当前数据库表结构是最新版本时，你又想新增一些表结构修改操作，这时候可以手动创建一个迁移文件再`up`以及`down`方法里做一些调整，一般情况下用不到，因为手动创建迁移是不会判断数据库状态的
- 回滚迁移一般用于运行当前最新的迁移文件导致数据库错误或者崩溃时把表结构回滚到上一次的迁移状态
### 迁移命令列表

- `db:migration:create`或`dbmc`：用于手动创建一个迁移文件（不会判断数据库状态）
- `db:migration:generate`或`dbmg`: 用于自动生成一个迁移文件
- `db:migration:run`或`dbmr`: 用于运行迁移或加上`-r`参数刷新数据库（即先删除数据库再运行迁移）或加上`-o`参数删除数据库
- `db:migration:revert`：用于把数据库回滚到最后一个迁移的状态，这是为了修复运行新的迁移可能导致的错误

详细命令及参数可以通过`pnpm cli -h`自行查看
每个命令的所有参数的解释可以运行`pnpm cli xxx -h`，比如`pnpm cli dbmg -h`
![](https://img.pincman.com/media/202303100542928.png#id=bfPDm&originHeight=1034&originWidth=1300&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
