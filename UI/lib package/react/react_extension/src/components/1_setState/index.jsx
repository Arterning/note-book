import React, { Component } from 'react'

export default class Demo extends Component {

	state = {count:0}

	add = ()=>{
		//对象式的setState 可以看到比较麻烦 需要先获取原来的 然后再更新
		/* //1.获取原来的count值
		const {count} = this.state
		//2.更新状态
		this.setState({count:count+1},()=>{
			console.log(this.state.count); //这里输出的是1 
		})
		//这里输出的是0 因为state的更新是异步进行的
		//如果需要获得更新后的state 需要使用上面的那种形式 也就是回调函数
		//console.log('12行的输出',this.state.count); //0 */

		//函数式的setState 比较简洁
		this.setState( state => ({count:state.count+1}))
	}

	render() {
		return (
			<div>
				<h1>当前求和为：{this.state.count}</h1>
				<button onClick={this.add}>点我+1</button>
			</div>
		)
	}
}
