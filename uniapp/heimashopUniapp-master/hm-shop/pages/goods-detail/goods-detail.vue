<template>
	<view>
		<view class="detail">
			<swiper indicator-dots>
				<swiper-item>
					<image src="http://demo.dtcms.net/upload/201504/20/thumb_201504200046589514.jpg"></image>
				</swiper-item>
				<swiper-item>
					<image src="http://demo.dtcms.net/upload/201504/20/thumb_201504200046589514.jpg"></image>
				</swiper-item>
			</swiper>
			<view class="box1">
				<view class="price">
					<text>￥{{detail.sell_price}}</text>
					<text>￥{{detail.market_price}}</text>
				</view>
				<view class="title">
					{{detail.title}}
				</view>
			</view>
			<view class="line"></view>
			<view class="box2">
				<view>货号：{{detail.goods_no}}</view>
				<view>库存：{{detail.stock_quantity}}</view>
			</view>
			<view class="line"></view>
			<view class="box3">
				<view>
					详情介绍
				</view>
				<view>
					<rich-text :nodes="content"></rich-text>
				</view>
			</view>
			<uni-goods-nav class="goods-nav" :fill="true"  :options="options" :buttonGroup="buttonGroup"  @click="onClick" @buttonClick="buttonClick" />
		</view>
	</view>
</template>

<script>
	import uniGoodsNav from '@/components/uni-goods-nav/uni-goods-nav.vue'
	export default {
		data() {
			return {
				id: 0,
				detail: {},
				content: '',
				options: [{
				            icon: 'headphones',
				            text: '客服'
				        }, {
				            icon: 'shop',
				            text: '店铺',
				            info: 0,
				            infoBackgroundColor:'#007aff',
				            infoColor:"red"
				        }, {
				            icon: 'cart',
				            text: '购物车',
				            info: 2
				        }],
				buttonGroup: [{
				          text: '加入购物车',
				          backgroundColor: '#ff0000',
				          color: '#fff'
				        },
				        {
				          text: '立即购买',
				          backgroundColor: '#ffa200',
				          color: '#fff'
				        }
				        ]
			}
		},
		methods: {
			async getDetail() {
				const res = await this.$myRequest({
					url:'/api/goods/getinfo/'+this.id
				})
				this.detail = res.data.message[0]
			},
			async getContent() {
				const res = await this.$myRequest({
					url:'/api/goods/getdesc/'+this.id
				})
				this.content = res.data.message[0].content
			},
			  onClick (e) {
			    uni.showToast({
			      title: `点击${e.content.text}`,
			      icon: 'none'
			    })
			  },
			  buttonClick (e) {
			    console.log(e)
			    this.options[2].info++
				}
		},
		onLoad(opthion) {
			this.id = opthion.id
			this.getDetail()
			this.getContent()
		}
	}
</script>

<style lang="scss">
swiper {
	width: 750rpx;
	height: 476px;
	swiper-item{
		width: 100%;
		height: 100%;
	}
	image {
		height: 100%;
		width: 100%;
	}
}
.box1 {
	.price {
		font-size: 23px;
		color: $index-color;
		line-height: 54px;
		text:nth-child(2) {
			color: #ccc;
			font-size: 19px;
			text-decoration: line-through;
			margin-left: 13px;
		}
	}
	.title {
		font-size: 42rpx;
	}
	.line {
		height: 12rpx;
		width: 750rpx;
		background-color: #EEEEEE;
	}
}
.box2 {
		line-height: 90rpx;
		font-size: 42rpx;
}
.box3 {
	view:nth-child(1) {
		font-size: 42rpx;
		padding-left: 20rpx;
	}
	view:nth-child(2) {
		font-size: 38rpx;
		color: #333333;
	}
}
.goods-nav {
	position: fixed;
	bottom: 0;
	width: 750rpx;
}
</style>
