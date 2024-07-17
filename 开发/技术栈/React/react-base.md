
## 创建虚拟DOM

使用原生JS创建虚拟DOM
```javascript
<script type="text/javascript" > 
    //1.创建虚拟DOM
    const VDOM = React.createElement('h1',{id:'title'},React.createElement('span',{},'Hello,React'))
    //2.渲染虚拟DOM到页面
    ReactDOM.render(VDOM,document.getElementById('test'))
</script>
```


使用JSX语法创建虚拟DOM
```javascript
	<script type="text/babel" > /* 此处一定要写babel */
		//1.创建虚拟DOM
		const VDOM = (  /* 此处一定不要写引号，因为不是字符串 */
			<h1 id="title">
				<span>Hello,React</span>
			</h1>
		)
		//2.渲染虚拟DOM到页面
		ReactDOM.render(VDOM,document.getElementById('test'))
	</script>
```


# 虚拟DOM的本质
1. 本质是Object类型的对象（一般对象）
2. 虚拟DOM比较“轻”，真实DOM比较“重”，因为虚拟DOM是React内部在用，无需真实DOM上那么多的属性。
3. 虚拟DOM最终会被React转化为真实DOM，呈现在页面上。

## JSX的语法规则
1. 定义虚拟DOM时，不要写引号。
2. 标签中混入JS表达式时要用{}。
3. 样式的类名指定不要用class，要用className。
4. 内联样式，要用style={{key:value}}的形式去写。
5. 只有一个根标签
6. 标签必须闭合
7. 标签首字母
(1).若小写字母开头，则将该标签转为html中同名元素，若html中无该标签对应的同名元素，则报错。
(2).若大写字母开头，react就去渲染对应的组件，若组件没有定义，则报错。

```javascript
const VDOM = (
			<div>
				<h2 className="title" id={myId.toLowerCase()}>
					<span style={{color:'white',fontSize:'29px'}}>{myData.toLowerCase()}</span>
				</h2>
				<h2 className="title" id={myId.toUpperCase()}>
					<span style={{color:'white',fontSize:'29px'}}>{myData.toLowerCase()}</span>
				</h2>
				<input type="text"/>
			</div>
		)
```


```javascript
const data = ['Angular','React','Vue']
//1.创建虚拟DOM
const VDOM = (
    <div>
        <h1>前端js框架列表</h1>
        <ul>
            {
                data.map((item,index)=>{
                    return <li key={index}>{item}</li>
                })
            }
        </ul>
    </div>
)
//2.渲染虚拟DOM到页面
ReactDOM.render(VDOM,document.getElementById('test'))
```

## 组件的State

状态(state)不可直接更改，必须通过调用setState更改
```javascript
class Weather extends React.Component{
			
        constructor(props){
            console.log('constructor');
            super(props)

            //初始化状态
            this.state = {isHot:false,wind:'微风'}
            this.changeWeather = this.changeWeather.bind(this)
        }

        render(){
            const {isHot,wind} = this.state
            return <h1 onClick={this.changeWeather}>今天天气很{isHot ? '炎热' : '凉爽'}，{wind}</h1>
        }

		changeWeather() {
            console.log('changeWeather');
            const isHot = this.state.isHot
            this.setState({isHot:!isHot})
            console.log(this);
		}
}
```

state的简写方式

```javascript
class Weather extends React.Component{

    state = {isHot:false,wind:'微风'}

    render(){
        const {isHot,wind} = this.state
        return <h1 onClick={this.changeWeather}>今天天气很{isHot ? '炎热' : '凉爽'}，{wind}</h1>
    }

    changeWeather = () => {
        const isHot = this.state.isHot
        this.setState({isHot:!isHot})
    }

}
```
在 React 中，箭头函数与普通函数的最大区别在于它们不绑定自己的 this 值，而是继承自其所在的作用域。因此，在 React 中，使用箭头函数定义的方法可以很方便地获取组件实例的 this 值，而无需手动绑定。

在传统的 JavaScript 中，this 的值是在函数被调用时确定的。如果函数被作为对象的方法调用，那么 this 就是该对象；如果函数不是作为对象的方法调用，那么 this 就是全局对象。为了确保 this 始终指向正确的对象，传统的 JavaScript 代码通常需要使用 bind()、call() 或 apply() 等方法手动绑定 this。

而在 React 中，组件方法的 this 值通常需要指向组件实例本身，以便在方法内部可以访问组件的状态和属性。使用箭头函数定义组件方法，可以避免手动绑定 this 的问题，从而使代码更加简洁易懂。

## props

``` javascript
class Person extends React.Component {

    //构造器是否接收props，取决于：是否希望在构造器中通过this访问props
    constructor(props){    
        super(props)
        console.log('constructor',this.props);
	}

    render(){
        const {name,age,sex} = this.props
        return (
            <ul>
                <li>姓名：{name}</li>
                <li>性别：{sex}</li>
                <li>年龄：{age+1}</li>
            </ul>
        )
    }
}

// set prop types
Person.propTypes = {
    name:PropTypes.string.isRequired,
    sex:PropTypes.string,
    age:PropTypes.number,
    speak:PropTypes.func,
}

// set default props
Person.defaultProps = {
    sex:'男',
    age:18
}
```

函数组件使用props

```javascript
function Person (props){
    const {name,age,sex} = props
    return (
        <ul>
            <li>姓名：{name}</li>
            <li>性别：{sex}</li>
            <li>年龄：{age}</li>
        </ul>
    )
}
```


## refs

```javascript
class Demo extends React.Component{
    //展示左侧输入框的数据
    showData = ()=>{
        const {input1} = this.refs
        alert(input1.value)
    }
    //展示右侧输入框的数据
    showData2 = ()=>{
        const {input2} = this.refs
        alert(input2.value)
    }
    render(){
        return(
            <div>
                <input ref="input1" type="text" placeholder="点击按钮提示数据"/>&nbsp;
                <button onClick={this.showData}>点我提示左侧的数据</button>&nbsp;
                <input ref="input2" onBlur={this.showData2} type="text" placeholder="失去焦点提示数据"/>
            </div>
        )
    }
}
```

## 事件处理

```javascript
class Demo extends React.Component{
    
    //创建ref容器
    myRef = React.createRef()

    //展示左侧输入框的数据
    showData = (event)=>{
        console.log(event.target);
        alert(this.myRef.current.value);
    }

    //展示右侧输入框的数据
    showData2 = (event)=>{
        alert(event.target.value);
    }

    render(){
        return(
            <div>
                <input ref={this.myRef} type="text" placeholder="点击按钮提示数据"/>&nbsp;
                <button onClick={this.showData}>点我提示左侧的数据</button>&nbsp;
                <input onBlur={this.showData2} type="text" placeholder="失去焦点提示数据"/>&nbsp;
            </div>
        )
    }
}
```

1. 通过onXxx属性指定事件处理函数(注意大小写)
* React使用的是自定义(合成)事件, 而不是使用的原生DOM事件, 为了更好的兼容性
* React中的事件是通过事件委托方式处理的(委托给组件最外层的元素), 为了的高效
2. 通过event.target得到发生事件的DOM元素对象, 不要过度使用ref
        

## 受控组件和非受控组件


### 受控组件

表单的内容存放在state中

```javascript
class Login extends React.Component{
    //初始化状态
    state = {
        username:'', //用户名
        password:'' //密码
    }

    saveUsername = (event)=>{
        this.setState({username:event.target.value})
    }

    savePassword = (event)=>{
        this.setState({password:event.target.value})
    }

    handleSubmit = (event)=>{
        event.preventDefault() //阻止表单提交
        const {username,password} = this.state
        alert(`你输入的用户名是：${username},你输入的密码是：${password}`)
    }

    render(){
        return(
            <form onSubmit={this.handleSubmit}>
                用户名：<input onChange={this.saveUsername} type="text" name="username"/>
                密码：<input onChange={this.savePassword} type="password" name="password"/>
                <button>登录</button>
            </form>
        )
    }
}
```


### 非受控组件

表单的内容并没有存放在state中

```javascript
class Login extends React.Component{
    handleSubmit = (event)=>{
        event.preventDefault() //阻止表单提交
        const {username,password} = this
        alert(`你输入的用户名是：${username.value},你输入的密码是：${password.value}`)
    }
    render(){
        return(
            <form onSubmit={this.handleSubmit}>
                用户名：<input ref={c => this.username = c} type="text" name="username"/>
                密码：<input ref={c => this.password = c} type="password" name="password"/>
                <button>登录</button>
            </form>
        )
    }
}
```

### 函数柯里化改写

```javascript
class Login extends React.Component{

    state = {
        username:'', //用户名
        password:'' //密码
    }

    saveFormData = (dataType)=>{
        return (event)=>{
            this.setState({[dataType]:event.target.value})
        }
    }

    //表单提交的回调
    handleSubmit = (event)=>{
        event.preventDefault() //阻止表单提交
        const {username,password} = this.state
        alert(`你输入的用户名是：${username},你输入的密码是：${password}`)
    }
    render(){
        return(
            <form onSubmit={this.handleSubmit}>
                用户名：<input onChange={this.saveFormData('username')} type="text" name="username"/>
                密码：<input onChange={this.saveFormData('password')} type="password" name="password"/>
                <button>登录</button>
            </form>
        )
    }
}
```

高阶函数：
如果一个函数符合下面2个规范中的任何一个，那该函数就是高阶函数。
* 若A函数，接收的参数是一个函数，那么A就可以称之为高阶函数。
* 若A函数，调用的返回值依然是一个函数，那么A就可以称之为高阶函数。

常见的高阶函数有：Promise、setTimeout、arr.map()等等

函数的柯里化：通过函数调用继续返回函数的方式，实现多次接收参数最后统一处理的函数编码形式

```javascript
function sum(a){
    return(b)=>{
        return (c)=>{
            return a+b+c
        }
    }
}
```

## 组件生命周期
1. 初始化阶段: 由ReactDOM.render()触发---初次渲染
        1.	constructor()
        2.	getDerivedStateFromProps 
        3.	render()
        4.	componentDidMount() 一般在这个钩子中做一些初始化的事，例如：开启定时器、发送网络请求、订阅消息
2. 更新阶段: 由组件内部this.setSate()或父组件重新render触发
        1.	getDerivedStateFromProps
        2.	shouldComponentUpdate()
        3.	render()
        4.	getSnapshotBeforeUpdate
        5.	componentDidUpdate()
3. 卸载组件: 由ReactDOM.unmountComponentAtNode()触发
        1.	componentWillUnmount() 一般在这个钩子中做一些收尾的事，例如：关闭定时器、取消订阅消息


```javascript
class Count extends React.Component{
    //构造器
    constructor(props){
        console.log('Count---constructor');
        super(props)
        //初始化状态
        this.state = {count:0}
    }

    add = ()=>{
        const {count} = this.state
        this.setState({count:count+1})
    }

    //卸载组件按钮的回调
    death = ()=>{
        ReactDOM.unmountComponentAtNode(document.getElementById('test'))
    }

    //强制更新按钮的回调
    force = ()=>{
        this.forceUpdate()
    }
    
    //若state的值在任何时候都取决于props，那么可以使用getDerivedStateFromProps
    static getDerivedStateFromProps(props,state){
        console.log('getDerivedStateFromProps',props,state);
        return null
    }

    //在更新之前获取快照
    getSnapshotBeforeUpdate(){
        console.log('getSnapshotBeforeUpdate');
        return 'getSnapshotBeforeUpdate'
    }

    //组件挂载完毕的钩子
    componentDidMount(){
        console.log('Count---componentDidMount');
    }

    //组件将要卸载的钩子
    componentWillUnmount(){
        console.log('Count---componentWillUnmount');
    }

    //控制组件更新的“阀门”
    shouldComponentUpdate(){
        console.log('Count---shouldComponentUpdate');
        return true
    }

    //组件更新完毕的钩子
    componentDidUpdate(preProps,preState,snapshotValue){
        console.log('Count---componentDidUpdate',preProps,preState,snapshotValue);
    }
    
    render(){
        console.log('Count---render');
        const {count} = this.state
        return(
            <div>
                <h2>当前求和为：{count}</h2>
                <button onClick={this.add}>点我+1</button>
                <button onClick={this.death}>卸载组件</button>
                <button onClick={this.force}>不更改任何状态中的数据，强制更新一下</button>
            </div>
        )
    }
}
```

## 虚拟DOM中key的选择

react/vue中的key有什么作用？key的内部原理是什么？为什么遍历列表时，key最好不要用index?
      
* 虚拟DOM中key的作用：
        1. key是虚拟DOM对象的标识, 在更新显示时key起着极其重要的作用。

        2. 当状态中的数据发生变化时，react会根据【新数据】生成【新的虚拟DOM】

        3. 随后React进行【新虚拟DOM】与【旧虚拟DOM】的diff比较，比较规则如下：

        a. 旧虚拟DOM中找到了与新虚拟DOM相同的key：
                    (1).若虚拟DOM中内容没变, 直接使用之前的真实DOM
                    (2).若虚拟DOM中内容变了, 则生成新的真实DOM，随后替换掉页面中之前的真实DOM

        b. 旧虚拟DOM中未找到与新虚拟DOM相同的key
                    根据数据创建新的真实DOM，随后渲染到到页面
                        
* 用index作为key可能会引发的问题：
        1. 若对数据进行：逆序添加、逆序删除等破坏顺序操作:
                        会产生没有必要的真实DOM更新 ==> 界面效果没问题, 但效率低。

        2. 如果结构中还包含输入类的DOM：
                        会产生错误DOM更新 ==> 界面有问题。
                        
        3. 注意！如果不存在对数据的逆序添加、逆序删除等破坏顺序操作，
            仅用于渲染列表用于展示，使用index作为key是没有问题的。
        
* 开发中如何选择key:
        1. 最好使用每条数据的唯一标识作为key, 比如id、手机号、身份证号、学号等唯一值。
        2. 如果确定只是简单的展示数据，用index也是可以的。
