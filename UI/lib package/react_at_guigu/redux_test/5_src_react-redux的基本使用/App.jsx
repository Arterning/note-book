import React, { Component } from 'react'
import Count from './containers/Count'
import store from './redux/store'

//App组件
//容器组件Count 传递给UI组件reduct中保存的状态 用于操作状态的方法
//UI组件 不要使用任何redux相关的API
export default class App extends Component {
	render() {
		return (
			<div>
				{/* 给容器组件传递store */}
				<Count store={store} />
			</div>
		)
	}
}
