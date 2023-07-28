构建配置式路由与Open API文档
## 效果和使用
控制台
![](https://img.pincman.com/media/202303050839752.png#id=zhS1Y&originHeight=1034&originWidth=1300&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
Open API
![](https://img.pincman.com/media/202303050839288.png#id=IXPwO&originHeight=1796&originWidth=2922&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
学完这一节后我们在测试api的时候（没学到E2E测试之前），可以很方便的使用Insomnia导入测试了
只要在Insomnia的中Create一个新的Design，然后导入你的API文档地址+json（比如`http://127.0.0.1:3100/api-docs/v1-json`或`http://127.0.0.1:3100/api-docs-json`）即可获取所有的API进行测试，而且API变动后重新导入即可覆盖更新
![](https://img.pincman.com/media/202303050843637.png#id=dOi4t&originHeight=1888&originWidth=2942&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)
## 预先阅读
请在学习本节课前预习阅读以下文档

- [Nestjs与Open API的结合使用的所有文档](https://docs.nestjs.com/openapi/introduction)
- [Nestjs的Router Module](https://docs.nestjs.com/recipes/router-module#router-module)
## 预装类库

- chalk用于对命令行中的文字进行颜色标识，此处我们安装老版本，因为新版本是esm写的不支持commonjs
- `@fastify/static`库使用用于fastify加载静态文件的，因为Open API(Swagger)文档展示静态的html和css文件，所以需要安装此库
```bash
~ pnpm add chalk@4 @fastify/static
```
## 核心模块
在编码之前我们需要对原来的核心模块进行一些修改，以便能通过配置来实现APP监听的地址以及端口等
### 类型
新增一个应用配置的类型，每个属性的作用请参考以下代码注释
```typescript
// src/modules/core/types.ts
export interface AppConfig {
    /**
     * 主机地址,默认为127.0.0.1
     */
    host: string;
    /**
     * 监听端口,默认3100
     */
    port: number;
    /**
     * 是否开启https,默认false
     */
    https: boolean;
    /**
     * 时区,默认Asia/Shanghai
     */
    timezone: string;
    /**
     * 语言,默认zh-cn
     */
    locale: string;
    /**
     * 控制台打印的url,默认自动生成
     */
    url?: string;
    /**
     * 由url+api前缀生成的基础api url
     */
    api?: string;
}
```
### 配置工厂
添加一个应用配置的配置工厂，在工厂中我们增加一些默认配置
```typescript
// src/modules/core/helpers/options.ts
export const createAppConfig: (
    register: ConfigureRegister<RePartial<AppConfig>>,
) => ConfigureFactory<AppConfig> = (register) => ({
    register,
    defaultRegister: (configure) => ({
        host: configure.env('APP_HOST', '127.0.0.1'),
        port: configure.env('APP_PORT', (v) => toNumber(v), 3000),
        https: configure.env('APP_PORT', (v) => toBoolean(v), false),
        timezone: configure.env('APP_TIMEZONE', 'Asia/Shanghai'),
        locale: configure.env('APP_LOCALE', 'zh-cn'),
    }),
});
```
### 添加配置
添加配置函数并导出
```typescript
// src/config/app.config.ts
export const app = createAppConfig((configure) => ({}));
// src/config/index.ts
export * from './app.config';
```
### 环境变量
修改环境变量的样板文件以及环境变量文件
```shell
# env.example
#应用
APP_HOST=127.0.0.1
APP_PORT=3000
APP_SSL=false
APP_TIMEZONE=Asia/Shanghai
APP_LOCAL=zh-cn
#数据库
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=3r

# .env
APP_PORT=3100
DB_PASSWORD=12345678
DB_NAME=3r
```
## Restful模块
### CRUD注册函数
CRUD注册函数用于定义一个生成CRUD选项的并通过metadata存储的函数。我们前面的课是直接通过对象的方式来存储的，这就导致无法灵活的使用一些其他功能，比如传入`configure`等，可以修改为通过函数存储CRUD信息，并在生成路由时自动执行此函数获得CRUD配置
首先定义一个注册函数的类型
```typescript
// src/modules/restful/types.ts
export type CrudOptionsRegister = (configure: Configure) => CrudOptions | Promise<CrudOptions>;
```
定义一个用于被Metadata存储函数的常量
```typescript
// src/modules/restful/constants.ts
export const CRUD_OPTIONS_REGISTER = 'crud_options_register';
```
修改原CRUD装饰器，使之只用于存储生成函数所用
```typescript
// src/modules/restful/decorators/crud.decorator.ts
export const Crud =
    (factory: CrudOptionsRegister) =>
    <T extends BaseController<any> | BaseControllerWithTrash<any>>(Target: Type<T>) => {
        Reflect.defineMetadata(CRUD_OPTIONS_REGISTER, factory, Target);
    };
```
编写一个CRUD选项生成函数，其逻辑就是原CRUD装饰器中的执行逻辑
```typescript
// src/modules/restful/crud.ts
export const registerCrud = async <T extends BaseController<any>>(
    Target: Type<T>,
    options: CrudOptions,
) => {
    const { id, enabled, dtos } = options;
    ...
};
```
### 控制器模块依赖
因为使用配置式路由的原理就是为一群控制器生成一个动态的路由模块，然后把这些模块放入Nestjs自带的RouterModule注册，最后生成一个树形的路由列表。而动态生成的路由模块不知道控制器原来是在哪个模块中，需要用到哪些服务的，所以我们需要为控制器存储固定的依赖模块，这样就可以在生成时导入这些模块了
例如，`PostController`原来是在`ContentModule`中的，所以能使用`ContentModule`中的`PostService`等服务，但是自动生成的`PostRouteModule`是无法使用`PostService`的，而新的控制器是放在`PostRouteModule`中的，所以必须在`ContentModule`的`exports`中添加`PostService`，然后为`PostRouteModule`中自动导入`ContentModule`才可以使控制器能注入`PostService`，这就需要我们使用metadata来存储控制器依赖的模块，以便在导入它的模块中同时导入它依赖的服务的模块
首先定义一个常量用于存储依赖模块
```typescript
// src/modules/restful/constants.ts
export const CONTROLLER_DEPENDS = 'controller_depends';
```
然后添加一个装饰器来存储
```typescript
// src/modules/restful/decorators/crud.decorator.ts
export const Depends = (...depends: Type<any>[]) => SetMetadata(CONTROLLER_DEPENDS, depends ?? []);
```
### 类型
#### Open API配置
`ApiTagOption`用于自定Swagger的标签配置，与Nestjs的Swagger模块的标签选项配置相对应
`ApiDocSource`用于定义Swagger文档的选项，定义的选项是会逐级合并覆盖的。因为每个路由集的选项会覆盖每个版本的选项，每个版本的选项又会向上覆盖整个模块的配置，它的属性如下

- `title`: 用于指定文档的名称
- `description`: 用于指定文档的描述
- `auth`: 用于指定是否启用Auth守卫验证（后续课程才会使用到）
- `tags`: 用于指定稳定的标签列表
```typescript
// src/modules/restful/types.ts
interface ApiTagOption {
    name: string;
    description?: string;
    externalDocs?: ExternalDocumentationObject;
}

export interface ApiDocSource {
    title?: string;
    description?: string;
    auth?: boolean;
    tags?: (string | ApiTagOption)[];
}
```
#### Restful模块配置
`ApiConfig`就是我们整个路由系统（即`RestfulModule`）的配置，其继承自`ApiDocSource`，可以设置一些默认的文档配置
```typescript
// src/modules/restful/types.ts
export interface ApiConfig extends ApiDocSource {
    prefix?: {
        route?: string;
        doc?: string;
    };
    default: string;
    enabled: string[];
    versions: Record<string, ApiVersionOption>;
}
```

- `prefix.route`: 用于指定所有API的总前缀，比如`api`
- `prefix.doc`: 用于指定Open API的文档前缀，比如`docs`
- `default`: 用于指定默认的API版本
- `enabled`: 用于指定启用的版本号列表，其中默认版本无需再加入该数组
- `versions`: 每个API版本的具体配置
#### 版本及路由配置
`ApiVersionOption`是版本配置类型，版本类型继承`ApiDocSource`类型，所以每个版本也可以拥有自己的文档配置，而`routes`则是路由集列表
`ApiRouteOption`是路由集配置，也可以配置自己的Swagger文档选项

- `name`: 用于指定当前路由集的名称
- `path`: 指定当前路由集的路径前缀，由于我们支持嵌套路由，所以这个是总前缀
- `controllers`：指定控制器列表
- `children`: 用于指定嵌套的子路由集
- `doc`: 用于指定路由集的Swagger文档配置
```typescript
// src/modules/restful/types.ts
export interface ApiVersionOption extends ApiDocSource {
    routes?: ApiRouteOption[];
}
export interface ApiRouteOption {
    name: string;
    path: string;
    controllers: Type<any>[];
    children?: ApiRouteOption[];
    doc?: ApiDocSource;
}
```
#### 自定义文档配置
`ApiSwaggerOption`与`ApiDocOption`则用于设置我们生成后的`docs`属性的类型，用于启动Swagger，后续代码会用到
```typescript
// src/modules/restful/types.ts
export interface ApiSwaggerOption extends ApiDocSource {
    version: string;
    path: string;
    include: Type<any>[];
}
export interface ApiDocOption {
    default?: ApiSwaggerOption;
    routes?: { [key: string]: ApiSwaggerOption };
}
```
### 路由配置类
首先我们需要编写一个路由模块的生成类，作用是根据我们传入的配置生成路由模块树供给RouterModule生成路由

- 构造函数接收`configure`实例以获取配置
- `getConfig`方法用于获取当前整个API配置中的某个配置
- `create`方法用于创建路由和Swagger文档，我们在它的子类中编写
```typescript
// src/modules/restful/configure.ts
export abstract class RouterConfigure {
    constructor(protected configure: Configure) {}
    getConfig<T>(key?: string, defaultValue?: any): T {
        return key ? get(this.config, key, defaultValue) : this.config;
    }

    abstract create(_config: ApiConfig): void;
    ...
}
```
#### 属性

- `config`: 为通过`configure`获取的整个API配置
- `routes`: 为生成的路由表，类型即为RouterModule的路由表类型
- `default`: 为默认版本号
- `versions`:为启用的版本号
- `modules`: 为自动生成的路由模块
```typescript
 // src/modules/restful/configure.ts
    protected config!: ApiConfig;
    protected _routes: Routes = [];
    protected _default!: string;
    protected _versions: string[] = [];
    protected _modules: { [key: string]: Type<any> } = {};

    get routes() {
        return this._routes;
    }

    get default() {
        return this._default;
    }

    get versions() {
        return this._versions;
    }

    get modules() {
        return this._modules;
    }
```
#### 生成基本配置
通过过滤，判断，合并对原来的API配置进行一定的处理生成一个新的配置
首先添加用于清理路由，添加路由前缀和生成URL的两个辅助函数函数
```typescript
// src/modules/restful/helpers.ts
/**
 * 路由路径前缀处理
 * @param routePath
 * @param addPrefix
 */
export const trimPath = (routePath: string, addPrefix = true) =>
    `${addPrefix ? '/' : ''}${trim(routePath.replace('//', '/'), '/')}`;

/**
 * 遍历路由及其子孙路由以清理路径前缀
 * @param data
 */
export const getCleanRoutes = (data: ApiRouteOption[]): ApiRouteOption[] =>
    data.map((option) => {
        const route: ApiRouteOption = {
            ...omit(option, 'children'),
            path: trimPath(option.path),
        };
        if (option.children && option.children.length > 0) {
            route.children = getCleanRoutes(option.children);
        } else {
            delete route.children;
        }
        return route;
    });
```
编写配置生成方法
```typescript
// src/modules/restful/configure.ts  
    protected createConfig(config: ApiConfig) {
        if (!config.default) {
            throw new Error('default api version name should been config!');
        }
        const versionMaps = Object.entries(config.versions)
            // 过滤启用的版本
            .filter(([name]) => {
                if (config.default === name) return true;
                return config.enabled.includes(name);
            })
            // 合并版本配置与总配置
            .map(([name, version]) => [
                name,
                {
                    ...pick(config, ['title', 'description', 'auth']),
                    ...version,
                    tags: Array.from(new Set([...(config.tags ?? []), ...(version.tags ?? [])])),
                    routes: getCleanRoutes(version.routes ?? []),
                },
            ]);

        config.versions = Object.fromEntries(versionMaps);
        // 设置所有版本号
        this._versions = Object.keys(config.versions);
        // 设置默认版本号
        this._default = config.default;
        // 启用的版本中必须包含默认版本
        if (!this._versions.includes(this._default)) {
            throw new Error(`Default api version named ${this._default} not exists!`);
        }
        this.config = config;
    }
```
#### 生成路由模块树
遍历每个版本，根据它们的路由列表使用`createRouteModuleTree`生成嵌套的路由模块树，同时额外地，生一套不带版本前缀的默认版本模块树
```typescript
// src/modules/restful/configure.ts
    protected async createRoutes() {
        const versionMaps = Object.entries(this.config.versions);

        // 对每个版本的路由使用'resolveRoutes'方法进行处理
        this._routes = (
            await Promise.all(
                versionMaps.map(async ([name, version]) =>
                    (
                        await createRouteModuleTree(
                            this.configure,
                            this._modules,
                            version.routes ?? [],
                            name,
                        )
                    ).map((route) => ({
                        ...route,
                        path: genRoutePath(route.path, this.config.prefix?.route, name),
                    })),
                ),
            )
        ).reduce((o, n) => [...o, ...n], []);
        // 生成一个默认省略版本号的路由
        const defaultVersion = this.config.versions[this._default];
        this._routes = [
            ...this._routes,
            ...(
                await createRouteModuleTree(
                    this.configure,
                    this._modules,
                    defaultVersion.routes ?? [],
                )
            ).map((route) => ({
                ...route,
                path: genRoutePath(route.path, this.config.prefix?.route),
            })),
        ];
    }
```
`createRouteModuleTree`函数
此函数用于构建路由模块，其逻辑如下

1. 根据生成路由模块名称对应的key，key值为`{父模块名}.{当前的路由集名}`
2. 判断模块名是否唯一
3. 获取路由集中每个控制器的依赖模块及CRUD处理函数
4. 使用`registerCrud`函数处理每个控制器
5. 如果该控制器自己没有设置Swagger标签，则为他们添加父级的标签
6. 把每个控制器的依赖导入到生成的路由模块
7. 生成路由模块，模块的类名使用首字母大写的驼峰命名
8. 把生成的模块放入模块列表，这是为了循环和递归时进行对比，防止路由模块重名（注意，此参数与`RouterConfigure`类的对象的`_modules`属性绑定）
9. 如果有子路由集，则递归执行
10. 最后返回路由树
```typescript
// src/modules/restful/helpers.ts
export const createRouteModuleTree = (
    configure: Configure,
    modules: { [key: string]: Type<any> },
    routes: ApiRouteOption[],
    parentModule?: string,
): Promise<Routes> =>
    Promise.all(
        routes.map(async ({ name, path, children, controllers, doc }) => {
            // 自动创建路由模块的名称
            const moduleName = parentModule ? `${parentModule}.${name}` : name;
            // RouteModule的名称必须唯一
            if (Object.keys(modules).includes(moduleName)) {
                throw new Error('route name should be unique in same level!');
            }
            // 获取每个控制器的依赖模块
            const depends = controllers
                .map((c) => Reflect.getMetadata(CONTROLLER_DEPENDS, c) || [])
                .reduce((o: Type<any>[], n) => {
                    if (o.find((i) => i === n)) return o;
                    return [...o, ...n];
                }, []);
            for (const controller of controllers) {
                const crudRegister = Reflect.getMetadata(CRUD_OPTIONS_REGISTER, controller);
                if (!isNil(crudRegister) && isFunction(crudRegister)) {
                    const crudOptions = isAsyncFn(crudRegister)
                        ? await crudRegister(configure)
                        : crudRegister(configure);
                    registerCrud(controller, crudOptions);
                }
            }
            // 为每个没有自己添加`ApiTags`装饰器的控制器添加Tag
            if (doc?.tags && doc.tags.length > 0) {
                controllers.forEach((controller) => {
                    !Reflect.getMetadata('swagger/apiUseTags', controller) &&
                        ApiTags(
                            ...doc.tags.map((tag) => (typeof tag === 'string' ? tag : tag.name))!,
                        )(controller);
                });
            }
            // 创建路由模块,并导入所有控制器的依赖模块
            const module = CreateModule(`${upperFirst(camelCase(name))}RouteModule`, () => ({
                controllers,
                imports: depends,
            }));
            // 在modules变量中追加创建的RouteModule,防止重名
            modules[moduleName] = module;
            const route: RouteTree = { path, module };
            // 如果有子路由则进一步处理
            if (children)
                route.children = await createRouteModuleTree(
                    configure,
                    modules,
                    children,
                    moduleName,
                );
            return route;
        }),
    );
```
#### 获取路由模块
这个方法用于获取一个树形模块集下的所有已经生成的路由模块，并通过递归的方式获取其子路集映射的模块列表
```typescript
// src/modules/restful/configure.ts
    protected getRouteModules(routes: ApiRouteOption[], parent?: string) {
        const result = routes
            .map(({ name, children }) => {
                const routeName = parent ? `${parent}.${name}` : name;
                let modules: Type<any>[] = [this._modules[routeName]];
                if (children) modules = [...modules, ...this.getRouteModules(children, routeName)];
                return modules;
            })
            .reduce((o, n) => [...o, ...n], [])
            .filter((i) => !!i);
        return result;
    }
```
### 配置构造类
这个类继承自`RouterConfigure`，前者用于添加生成路由树和路由模块树的方法，此类用于生成Swagger文档

- `create`方法用于创建配置，创建路由树和路由模块树，已经创建Swagger文档的配置
- `getModuleImports`方法用于获取所有需要导入到`RestfulModule`的路由模块，以及`RouterModule`
```typescript
// src/modules/restful/restful.ts
@Injectable()
export class Restful extends RouterConfigure {
    async create(config: ApiConfig) {
        this.createConfig(config);
        await this.createRoutes();
        this.createDocs();
    }
  
    getModuleImports() {
        return [...Object.values(this.modules), RouterModule.register(this.routes)];
    }
}
```
#### 属性

- `_docs`使用存储每个版本的键值对映射的文档选项
- `excludeVersionModules`用于在生成文档选项时排除已经添加到文档的模块
```typescript
// src/modules/restful/restful.ts
    protected _docs!: {
        [version: string]: ApiDocOption;
    };

    /**
     * 排除已经添加的模块
     */
    protected excludeVersionModules: string[] = [];

    get docs() {
        return this._docs;
    }
```
#### 生成路由文档
`getRouteDocs`函数用于生成每个路由列表的文档配置，逻辑如下

1. 通过合并覆盖父级获得文档信息
2. `include`包含当前路由集的路由模块
3. 遍历该路由集下的所有路由
4. 排除已经添加到`include`的路由模块
5. 获取当前路由下的自定义的文档配置并合并覆盖上级
6. 如果有子路由集则递归添加
7. 返回路由文档
```typescript
// src/modules/restful/restful.ts    
    protected getRouteDocs(
        option: Omit<ApiSwaggerOption, 'include'>,
        routes: ApiRouteOption[],
        parent?: string,
    ): { [key: string]: ApiSwaggerOption } {
        /**
         * 合并Doc配置
         *
         * @param {Omit<ApiSwaggerOption, 'include'>} vDoc
         * @param {ApiRouteOption} route
         */
        const mergeDoc = (vDoc: Omit<ApiSwaggerOption, 'include'>, route: ApiRouteOption) => ({
            ...vDoc,
            ...route.doc,
            tags: Array.from(new Set([...(vDoc.tags ?? []), ...(route.doc?.tags ?? [])])),
            path: genDocPath(route.path, this.config.prefix?.doc, parent),
            include: this.getRouteModules([route], parent),
        });
        let routeDocs: { [key: string]: ApiSwaggerOption } = {};

        // 判断路由是否有除tags之外的其它doc属性
        const hasAdditional = (doc?: ApiDocSource) =>
            doc && Object.keys(omit(doc, 'tags')).length > 0;

        for (const route of routes) {
            const { name, doc, children } = route;
            const moduleName = parent ? `${parent}.${name}` : name;

            // 加入在版本DOC中排除模块列表
            if (hasAdditional(doc) || parent) this.excludeVersionModules.push(moduleName);

            // 添加到routeDocs中
            if (hasAdditional(doc)) {
                routeDocs[moduleName.replace(`${option.version}.`, '')] = mergeDoc(option, route);
            }
            if (children) {
                routeDocs = {
                    ...routeDocs,
                    ...this.getRouteDocs(option, children, moduleName),
                };
            }
        }
        return routeDocs;
    }
```
#### 生成版本文档
首先排除已经添加到文档的`include`中的模块
```typescript
// src/modules/restful/restful.ts    
    protected filterExcludeModules(routeModules: Type<any>[]) {
        const excludeModules: Type<any>[] = [];
        const excludeNames = Array.from(new Set(this.excludeVersionModules));
        for (const [name, module] of Object.entries(this._modules)) {
            if (excludeNames.includes(name)) excludeModules.push(module);
        }
        return routeModules.filter(
            (rmodule) => !excludeModules.find((emodule) => emodule === rmodule),
        );
    }
```
通过`getDocOption`方法来生成每个版本的文档配置
逻辑为

1. 根据版本的文档配置设置一个默认的文档配置
2. 获取版本的路由集的Swagger配置
3. 过滤掉已经添加到`include`的路由模块
4. 如果还有多余的模块或者在没有当路由没文档的情况下，把`include`添加到选项中
5. 返回文档的版本配置
```typescript
// src/modules/restful/restful.ts
    protected getDocOption(name: string, voption: ApiVersionOption, isDefault = false) {
        const docConfig: ApiDocOption = {};
        // 默认文档配置
        const defaultDoc = {
            title: voption.title!,
            description: voption.description!,
            tags: voption.tags ?? [],
            auth: voption.auth ?? false,
            version: name,
            path: trim(`${this.config.prefix?.doc}${isDefault ? '' : `/${name}`}`, '/'),
        };
        // 获取路由文档
        const routesDoc = isDefault
            ? this.getRouteDocs(defaultDoc, voption.routes ?? [])
            : this.getRouteDocs(defaultDoc, voption.routes ?? [], name);
        if (Object.keys(routesDoc).length > 0) {
            docConfig.routes = routesDoc;
        }
        const routeModules = isDefault
            ? this.getRouteModules(voption.routes ?? [])
            : this.getRouteModules(voption.routes ?? [], name);
        // 文档所依赖的模块
        const include = this.filterExcludeModules(routeModules);
        // 版本DOC中有依赖的路由模块或者版本DOC中没有路由DOC则添加版本默认DOC
        if (include.length > 0 || !docConfig.routes) {
            docConfig.default = { ...defaultDoc, include };
        }
        return docConfig;
    }
```
#### 创建文档配置
遍历所有版本，为每个API版本添加文档配置，并且增加一个默认版本的文档配置
```typescript
// src/modules/restful/restful.ts    
    protected createDocs() {
        const versionMaps = Object.entries(this.config.versions);
        const vDocs = versionMaps.map(([name, version]) => [
            name,
            this.getDocOption(name, version),
        ]);
        this._docs = Object.fromEntries(vDocs);
        const defaultVersion = this.config.versions[this._default];
        // 为默认版本再次生成一个文档
        this._docs.default = this.getDocOption(this._default, defaultVersion, true);
    }
```
#### 构建Open API
根据每个版本的文档配置启动多个Swagger应用
```typescript
// src/modules/restful/restful.ts
    factoryDocs<T extends INestApplication>(app: T) {
        const docs = Object.values(this._docs)
            .map((vdoc) => [vdoc.default, ...Object.values(vdoc.routes ?? {})])
            .reduce((o, n) => [...o, ...n], [])
            .filter((i) => !!i);
        for (const voption of docs) {
            const { title, description, version, auth, include, tags } = voption!;
            const builder = new DocumentBuilder();
            if (title) builder.setTitle(title);
            if (description) builder.setDescription(description);
            if (auth) builder.addBearerAuth();
            if (tags) {
                tags.forEach((tag) =>
                    typeof tag === 'string'
                        ? builder.addTag(tag)
                        : builder.addTag(tag.name, tag.description, tag.externalDocs),
                );
            }
            builder.setVersion(version);
            const document = SwaggerModule.createDocument(app, builder.build(), {
                include: include.length > 0 ? include : [() => undefined as any],
            });
            SwaggerModule.setup(voption!.path, app, document);
        }
    }
```
### 创建模块
编写一个Restful模块，用于导入路由模块以及设置`Restful`这个提供者
```typescript
// src/modules/restful/restful.module.ts
@ModuleBuilder(async (configure) => {
    const restful = new Restful(configure);
    await restful.create(await configure.get('api'));
    return {
        global: true,
        imports: restful.getModuleImports(),
        providers: [
            {
                provide: Restful,
                useValue: restful,
            },
        ],
        exports: [Restful],
    };
})
export class RestfulModule {}
```
### 端口打印
添加两个函数，分别用于根据生成的配置来打印API和DOC的URL，其中`echoApi`调用`echoDcos`
```typescript
// src/modules/restful/helpers.ts
export async function echoApi(configure: Configure, restful: Restful) {
    const appUrl = await configure.get<string>('app.url');
    const apiUrl = await configure.get<string>('app.api');
    console.log(`- ApiUrl: ${chalk.green.underline(apiUrl)}`);
    console.log('- ApiDocs:');
    const { default: defaultDoc, ...docs } = restful.docs;
    echoDocs('default', defaultDoc, appUrl);
    for (const [name, doc] of Object.entries(docs)) {
        console.log();
        echoApiDocs(name, doc, appUrl);
    }
}

function echoDocs(name: string, doc: ApiDocOption, appUrl: string) {
    const getDocPath = (dpath: string) => `${appUrl}/${dpath}`;
    if (!doc.routes && doc.default) {
        console.log(
            `    [${chalk.blue(name.toUpperCase())}]: ${chalk.green.underline(
                getDocPath(doc.default.path),
            )}`,
        );
        return;
    }
    console.log(`    [${chalk.blue(name.toUpperCase())}]:`);
    if (doc.default) {
        console.log(`      default: ${chalk.green.underline(getDocPath(doc.default.path))}`);
    }
    if (doc.routes) {
        Object.entries(doc.routes).forEach(([_routeName, rdocs]) => {
            console.log(
                `      <${chalk.yellowBright.bold(rdocs.title)}>: ${chalk.green.underline(
                    getDocPath(rdocs.path),
                )}`,
            );
        });
    }
}
```
## 更新应用
所有准备工作都做好了，接下来我们就来更改应用，实现配置式路由和Swagger整合的美妙吧
### 内容模块
删除内容模块中原来导入的`controllers`
```typescript
@ModuleBuilder(async (configure) => {
    ...
    return {
        imports: [
            TypeOrmModule.forFeature(Object.values(entities)),
            DatabaseModule.forRepository(Object.values(repositories)),
        ],
        // controllers: Object.values(controllers),
        providers,
        exports: [
            ...Object.values(services),
            PostService,
            DatabaseModule.forRepository(Object.values(repositories)),
        ],
    };
})
export class ContentModule {}
```
### 辅助函数
该函数用于快捷地为控制器的CRUD选项添加一个Hook，Hook中包含一些常用的东西，比如添加一个`ApiOperation`装饰器等等
```typescript
// src/modules/restful/helpers.ts
export function createHookOption(summary?: string): CrudMethodOption {
    return {
        hook: (target, method) => {
            if (!isNil(summary))
                ApiOperation({ summary })(
                    target,
                    method,
                    Object.getOwnPropertyDescriptor(target.prototype, method),
                );
        },
    };
}
```
### 控制器
> 此处仅以`CategoryController`为例，其它的控制器请自行查看代码

更改控制器

- 使用`@ApiTags`装饰器为每个控制器添加标签
- 在`@Crud`装饰器中传入异步函数用于代替原来的固定选项对象，同时我们通过HOOK为每个操作添加一个API注释
- 如果有特殊的额外操作方法，比如`tree`，我们给他单独添加`@ApiOperation`装饰器
```typescript
// src/modules/content/controllers/category.controller.ts
@ApiTags('分类')
@Depends(ContentModule)
@Crud(async () => ({
    id: 'category',
    enabled: [
        {
            name: 'list',
            option: createHookOption('分类查询,以分页模式展示'),
        },
        {
            name: 'detail',
            option: createHookOption('分类详情'),
        },
        {
            name: 'store',
            option: createHookOption('创建分类'),
        },
        {
            name: 'update',
            option: createHookOption('更新分类'),
        },
        {
            name: 'delete',
            option: createHookOption('删除分类'),
        },
        {
            name: 'restore',
            option: createHookOption('恢复分类'),
        },
    ],
    dtos: {
        store: CreateCategoryDto,
        update: UpdateCategoryDto,
    },
}))
@Controller('categories')
export class CategoryController extends BaseControllerWithTrash<CategoryService> {
    constructor(protected service: CategoryService) {
        super(service);
    }

    @Get('tree')
    @ApiOperation({ summary: '树形结构分类查询' })
    @SerializeOptions({ groups: ['category-tree'] })
    async tree(@Query() options: QueryCategoryTreeDto) {
        return this.service.findTrees(options);
    }
}
```
### DTO验证
> 此处列出两个，其它的自行查看代码

为每个验证DTO（包括公共DTO和内容模块中的DTO）中的属性添加上Swagger注释
```typescript
// src/modules/restful/dtos/delete-with-trash.dto.ts
@DtoValidation({ type: 'query' })
export class ListWithTrashedQueryDto extends ListQueryDto implements TrashedOptions {
    @ApiPropertyOptional({
        description:
            '回收站数据过滤,all:包含已软删除和未软删除的数据;only:只包含软删除的数据;none:只包含未软删除的数据',
        enum: SelectTrashMode,
    })
    @IsEnum(SelectTrashMode)
    @IsOptional()
    trashed?: SelectTrashMode;
}

// src/modules/content/dtos/post.dto.ts
@DtoValidation({ groups: ['update'] })
export class UpdatePostDto extends PartialType(CreatePostDto) {
    @ApiProperty({
        description: '待更新的文章ID',
    })
    @IsUUID(undefined, { groups: ['update'], message: '文章ID格式错误' })
    @IsDefined({ groups: ['update'], message: '文章ID必须指定' })
    id!: string;
}
...
```
### 路由配置
版本路由配置
```typescript
// src/routes/v1.ts
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
                ],
            },
            children: [
                {
                    name: 'content',
                    path: 'content',
                    controllers: Object.values(contentControllers),
                },
            ],
        },
    ],
});
```

API配置

```typescript
// src/routes/index.ts
export const api = async (configure: Configure): Promise<ApiConfig> => ({
    title: configure.env('API_TITLE', '3R教室'),
    description: configure.env('API_DESCRIPTION', '3R教室TS全栈开发教程'),
    auth: true,
    prefix: { route: 'api', doc: 'api-docs' },
    default: configure.env('API_DEFAULT_VERSION', 'v1'),
    enabled: [],
    versions: { v1: await v1(configure) },
});
```
在`configs/index.ts`中使用`export * from '../routes';`加入配置
### 环境变量
```shell
# env.example

#Open API
API_TITLE=3R教室
API_DESCRIPTION=3R教室TS全栈开发教程
API_DEFAULT_VERSION=v1
```
### 加载路由模块
与数据库等核心模块一样，可以根据是否传入`api`配置来自动导入`RestfulModule`以及启动Swagger文档
```typescript
// src/modules/core/helpers/app.ts
export async function createBootModule(
    params: AppParams,
    options: Pick<Partial<CreateOptions>, 'meta' | 'modules' | 'globals'>,
): Promise<{ BootModule: Type<any>; modules: ModuleBuildMap }> {
    ...
    if (configure.has('api')) importModules.push(RestfulModule);
}

// src/modules/core/app.ts
    static async create(options: CreateOptions): Promise<CreatorData> {
        const { builder, configs, configure } = options;
        let modules = {};
        try {
            // 根据是否传入api配置来启用open api功能
            if (!isNil(await this._configure.get<ApiConfig>('api', null))) {
                const restful = this._app.get(Restful);
                restful.factoryDocs(this._app);
            }
            // 允许使用关闭监听的钩子
            this._app.enableShutdownHooks();
           ...
        } catch (error) {
            console.log('Create app failed!');
            exit(0);
        }

        return { configure: this._configure, app: this._app, modules };
    }
```
### 设置API Url
同时，为了方便本地开发时访问API和文档，我们需要在`buildConfigure`方法中设置一下`app.url`和`app.api`
```typescript
// src/modules/core/app.ts
   static async buildConfigure(configs: Record<string, any>, option?: ConfigStorageOption) {
        ...
        let appUrl = await configure.get('app.url', undefined);
        if (isNil(appUrl)) {
            const host = await configure.get<string>('app.host');
            const port = await configure.get<number>('app.port')!;
            const https = await configure.get<boolean>('app.https');
            appUrl =
                (await configure.get<boolean>('app.url', undefined)) ??
                `${https ? 'https' : 'http'}://${host!}:${port}`;

            configure.set('app.url', appUrl);
        }
        const routePrefix = await configure.get('api.prefix.route', undefined);
        const apiUrl = routePrefix
            ? `${appUrl}${routePrefix.length > 0 ? `/${routePrefix}` : routePrefix}`
            : appUrl;
        configure.set('app.api', apiUrl);
        return configure;
    }
```
### 启动函数
修改应用启动函数，让它可以接收一个监听回调函数
```typescript
export async function bootApp(
    creator: () => Promise<CreatorData>,
    listened?: (params: CreatorData) => () => Promise<void>,
) {
    const { app, configure } = await creator();
    const { port, host } = await configure.get<AppConfig>('app');
    await app.listen(port, host, listened({ app, configure } as any));
}
```
### 打印URL
最后修改`main.ts`，添加一个监听回调，以便在启动应用后打印出API和Doc的地址
```typescript
// src/main.ts
const creator = createApp({
    ...
});
bootApp(creator, ({ app, configure }) => async () => {
    const restful = app.get(Restful);
    echoApi(configure, restful);
});
```
