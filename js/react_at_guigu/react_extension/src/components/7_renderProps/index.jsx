import React, { Component } from 'react'
import './index.css'

export default class Parent extends Component {
	render() {
		return (
			<div className="parent">
				<h3>我是Parent组件</h3>
				{/* 这里给A传入一个render函数 展示B 同时传递name给B */}
				<A render={(name)=><B name={name}/>}/>
				{/* <A>hello</A> hello不会展示*/}
				{/* <A><B/></A> 展示B组件*/}
			</div>
		)
	}
}

//这样做的好处是什么呢？
//就是A组件预留了一个位置存放子组件 这个子组件不需要写死 而是可以在A标签灵活传入
//对应了Vue里面的插槽技术
class A extends Component {
	state = {name:'tom'}
	render() {
		console.log(this.props);
		const {name} = this.state
		return (
			<div className="a">
				<h3>我是A组件</h3>
				{/* 调用传入的render函数 预留了插槽*/}
				{this.props.render(name)}
				{/* {this.props.children} 需要这样才能拿到标签体*/}
			</div>
		)
	}
}

class B extends Component {
	render() {
		console.log('B--render');
		return (
			<div className="b">
				<h3>我是B组件,{this.props.name}</h3>
			</div>
		)
	}
}
