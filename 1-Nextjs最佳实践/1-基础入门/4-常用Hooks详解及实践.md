## 准备
关于react的基础知识，比如jsx等等其新官网已经说的非常明确，指直接查看[官网中文文档](https://zh-hans.reactjs.org/learn/describing-the-ui)即可，本教材不再赘述，我们直接开始学习hooks的灵活使用
依赖
```bash
~ pnpm add clsx immer lodash
~ pnpm add @types/lodash -D
```
## `useState`
此Hook用于改变组件内的状态,例
```tsx
export const StateDemo: FC = () => {
    const [count, setCount] = useState(1);
    const [isShow, toggleShow] = useState(true);

    return (
        <div className={clsx($styles.container, 'tw-w-[20rem]')}>
            <h2 className="tw-text-center">useState Demo</h2>
            {isShow && <p className="tw-text-center tw-py-5">{count}</p>}
            <div className="tw-flex tw-justify-around">
                <Button onClick={() => setCount(count + 1)} type="dashed" ghost>
                    增加
                </Button>
                <Button onClick={() => setCount(count - 1)} type="dashed" ghost>
                    减少
                </Button>
                <Button onClick={() => toggleShow(!isShow)} type="dashed" ghost>
                    {isShow ? '显示' : '隐藏'}
                </Button>
            </div>
        </div>
    );
};
```
## `useEffect`
在状态不同的生命周期执行副作用
### 简单用法
每次状态更新都执行所有没有依赖的`useEffect`,以下代码'toggle ghost'这一条在`resize`浏览器时也会触发
```tsx
export const EffectDemo: FC = () => {
    const [ghost, setGhost] = useState<boolean>(false);
    const [width, setWidth] = useState(window.innerWidth);
    const toggleGhostBtn = () => setGhost(!ghost);
    const resizeHandle = () => setWidth(window.innerWidth);
    useEffect(() => {
        console.log('浏览器宽度改变');
        window.addEventListener('resize', resizeHandle);
    },);
    useEffect(() => {
        console.log('切换幽灵按钮');
    });
    return (
        <div className={$styles.container}>
            <h2 className="tw-text-center">useEffect Demo</h2>
            <p className="tw-text-center tw-py-5">{ghost ? 'ghost' : '普通'}按钮</p>
            <Button type="primary" onClick={toggleGhostBtn} ghost={ghost}>
                切换按钮样式
            </Button>
            <p className="pt-5">宽度为: {width}</p>
        </div>
    );
};
```
### 依赖更新
通过`useEffect`的第二个参数,可以指定其依赖的变量,只有此变量的状态更改时才会执行副作用函数,如果第二个参数为空,则只在第一次渲染和重新渲染时触发
```tsx
const EffectDemo: FC = () => {
    ...
    useEffect(() => {
        console.log('浏览器宽度改变');
        window.addEventListener('resize', resizeHandle);
    }, [width]);
    useEffect(() => {
        console.log('切换幽灵按钮');
    }, [ghost]);
    useEffect(() => {
        console.log('只在第一次或重新渲染组件时触发');
    }, []);
};
```
### 清理监听
在监听`width`的`useEffect`中,每次改变`width`的状态,都会添加一个`resize`事件,这会极大的耗费浏览器占用的内存,通过一个返回值的方式,即可在下一次`width`状态改变后与添加新的`resize`监听前,取消上次添加的`resize`监听事件
```tsx
const EffectDemo: FC = () => {
    ...
    useEffect(() => {
        window.addEventListener('resize', resizeHandle);
        return () => {
            window.removeEventListener('resize', resizeHandle);
        };
    }, [width]);
};
```
### 异步执行
在`useEffect`中执行异步函数的语法如下,其实就是在原函数里调用一个`async`打头的立即函数
```typescript
useEffect(() => {
    (async () => {})();
});
```
以下示例代码让按钮在变成`ghost`之后1s再变红色
```tsx

const EffectDemo: FC = () => {
    const [red, setRed] = useState<boolean>(false);
     useEffect(() => {
        (async () => {
            await new Promise((resolve, reject) => {
                setTimeout(() => resolve(true), 1000);
            });
            setRed(ghost);
        })();
    }, [ghost]);
    return (
        <div>
            <Button type="primary" onClick={toggleGhostBtn} ghost={ghost} danger={red}>
                切换按钮样式
            </Button>
        </div>
    );
};
```
## `useContext`
用于向后代组件透传一个值,以创建一个语言选择器为例
### 定义类型
```typescript
type LangType = {
    name: string;
    label: string;
    data: Locale;
};
type LangState = {
    lang: LangType;
    setLang: (lang: LangType) => void;
};
```
### 定义语言列表
```typescript
export const langs: LangType[] = [
    {
        name: 'en_US',
        label: '🇺🇸 english(US)',
        data: enUS,
    },
    {
        name: 'zh_CN',
        label: '🇨🇳 简体中文',
        data: zhCN,
    },
];
```
### 创建`context`
```typescript
export const LangContext = createContext<LangState>({
    lang: langs[0],
    setLang: (lang: LangType) => {},
});
```
### 创建`provider`包装器
```tsx
const LangProvider: FC<LangState & { children?: ReactNode }> = ({ lang, setLang, children }) => {
    return <LangContext.Provider value={{ lang, setLang }}>{children}</LangContext.Provider>;
};
```
### 创建`Lang`组件
```tsx
export const Lang: FC<{ children?: ReactNode }> = ({ children }) => {
    const [lang, setLang] = useState<LangType>(langs[0]);
    return (
        <LangProvider lang={lang} setLang={setLang}>
            {children}
        </LangProvider>
    );
};
```
### `App.tsx`包装
把`App.tsx`中的所有节点包含于`Lang`组件
```tsx
const App: FC = () => {
    return (
        <Lang>
        ...
          <ContextDemo />
        </Lang>
    );
};
export default App;
```
### 语言选择组件
```tsx
export const LangConfig: FC = () => {
    const { lang, setLang } = useContext(LangContext);
    const changeLang = (value: string) => {
        const current = langs.find((item) => item.name === value);
        current && setLang(current);
    };
    return (
        <Select defaultValue={lang.name} style={{ width: 120 }} onChange={changeLang}>
            {langs.map(({ name, label }) => (
                <Select.Option key={name} value={name}>
                    {label}
                </Select.Option>
            ))}
        </Select>
    );
```
### Demo组件
```tsx
const ContextDemo: FC = () => {
    const { lang } = useContext(LangContext);
    return (
        <div className={$styles.container}>
            <h2 className="tw-text-center">useContext Demo</h2>
            <p className="tw-text-center tw-py-5">当前语言: {lang.label}</p>
            <div className="tw-w-full">
                <h3>Antd语言切换测试</h3>
                <div className="tw-w-full tw-my-4">
                    <LangConfig />
                </div>
            </div>
            <Pagination defaultCurrent={0} total={500} />
        </div>
    );
};
export default ContextDemo;
```
更改布局组件
```tsx
const MasterLayout: FC<{ children?: ReactNode }> = ({ children }) => {
    const { lang } = useContext(LangContext);

    return (
        <ConfigProvider
            locale={lang.data}
            theme={{
                ...
            }}
        >
            <StyleProvider hashPriority="high">
                <div className={$styles.masterLayout}>{children}</div>
            </StyleProvider>
        </ConfigProvider>
    );
};
```
## `useReducer`
使用`Context`+`useReducer`可以实现轻量级的全局状态管理
以实现一个简单的应用配置功能为例(包含标题设置和暗黑模式切换)
### 编写类型
```typescript
export enum ThemeMode {
    LIGHT = 'light',
    DARK = 'dark',
}

export type ThemeState = {
    mode: `${ThemeMode}`;
    compact: boolean;
};

export enum ThemeActionType {
    CHANGE_MODE = 'change_mode',
    CHANGE_COMPACT = 'change_compact',
}

export type ThemeAction =
    | { type: `${ThemeActionType.CHANGE_MODE}`; value: `${ThemeMode}` }
    | { type: `${ThemeActionType.CHANGE_COMPACT}`; value: boolean };

export type ThemeContextType = {
    state: ThemeState;
    dispatch: Dispatch<ThemeAction>;
};
```
### 默认配置
```typescript
export const defaultThemeConfig: ThemeState = {
    mode: 'light',
    compact: false,
};

```
### 创建`Context`
```typescript
// 透传配置状态与dispatch
export const ThemeContext = createContext<ThemeContextType | null>(null);
```
### 状态操作
为了确保数据的唯一性不被污染,使用[immer.js][]操作数据
```typescript
export const ThemeReducer: Reducer<ThemeState, ThemeAction> = produce((draft, action) => {
    switch (action.type) {
        case 'change_mode':
            draft.mode = action.value;
            break;
        case 'change_compact':
            draft.compact = action.value;
            break;
        default:
            break;
    }
});
```
### 包装器组件

- 合并默认配置和初始化配置
- 把配置状态和`dispatch`传给`ThemeContext`
```tsx
export const Theme: FC<{ data?: ThemeState; children?: ReactNode }> = ({ data = {}, children }) => {
    const [state, dispatch] = useReducer(ThemeReducer, data, (initData) => ({
        ...defaultThemeConfig,
        ...initData,
    }));
    useEffect(() => {
        if (state.mode === 'dark') {
            document.documentElement.setAttribute('data-theme', 'tw-dark');
        } else {
            document.documentElement.removeAttribute('data-theme');
        }
    }, [state.mode]);
    return <ThemeContext.Provider value={{ state, dispatch }}>{children}</ThemeContext.Provider>;
};
```
### 主题选择组件
```tsx
export const ThemeConfig: FC = () => {
    const context = useContext(ThemeContext);
    if (isNil(context)) return null;
    const { state, dispatch } = context;
    const toggleMode = () =>
        dispatch({ type: 'change_mode', value: state.mode === 'light' ? 'dark' : 'light' });
    const toggleCompact = () => dispatch({ type: 'change_compact', value: !state.compact });
    return (
        <>
            <Switch
                checkedChildren="🌛"
                unCheckedChildren="☀️"
                onChange={toggleMode}
                checked={state.mode === 'dark'}
                defaultChecked={state.mode === 'dark'}
            />
            <Switch
                checkedChildren="紧凑"
                unCheckedChildren="正常"
                onChange={toggleCompact}
                checked={state.compact}
                defaultChecked={state.compact}
            />
        </>
    );
};
```
### Demo组件
```tsx
const ReducerDemo: FC = () => {
    const context = useContext(ThemeContext);
    if (isNil(context)) return null;
    const {
        state: { mode, compact },
    } = context;
    return (
        <div className={$styles.container}>
            <h2 className="tw-text-center">useReducer Demo</h2>
            <p>主题模式: 「{mode === 'dark' ? '暗黑' : '明亮'}」</p>
            <p>是否紧凑: 「{compact ? '是' : '否'}」</p>
            <div className="tw-w-full">
                <ThemeConfig />
            </div>
        </div>
    );
};
```
### 在`App.tsx`中包装
```tsx

const App = () => {
    return (
        <Lang>
            <Theme>
                <MasterLayout>
                    ...
                    <ReducerDemo />
                </MasterLayout>
            </Theme>
        </Lang>
    );
};
```
### 更改布局组件
```typescript
const MasterLayout: FC<{ children?: ReactNode }> = ({ children }) => {
    const { lang } = useContext(LangContext);
    const themeContext = useContext(ThemeContext);
    const [antdThemes, setAntdThemes] = useState<MappingAlgorithm[]>([theme.defaultAlgorithm]);
    let themeMode = defaultThemeConfig.mode;
    let themeCompact = defaultThemeConfig.compact;
    if (!isNil(themeContext)) {
        themeMode = themeContext.state.mode;
        themeCompact = themeContext.state.compact;
    }
    useEffect(() => {
        if (!themeCompact) {
            setAntdThemes(() =>
                themeMode === 'dark' ? [theme.darkAlgorithm] : [theme.defaultAlgorithm],
            );
        } else {
            setAntdThemes(() =>
                themeMode === 'dark'
                    ? [theme.darkAlgorithm, theme.compactAlgorithm]
                    : [theme.defaultAlgorithm, theme.compactAlgorithm],
            );
        }
    }, [themeMode, themeCompact]);

    return (
        <ConfigProvider
            locale={lang.data}
            theme={{
                algorithm: antdThemes,
                components: {
                    Layout: {
                        colorBgBody: '',
                    },
                },
            }}
        >
            <StyleProvider hashPriority="high">
                <div className={$styles.masterLayout}>{children}</div>
            </StyleProvider>
        </ConfigProvider>
    );
};
```
