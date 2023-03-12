<template>
	<view class="home">
		<swiper indicator-dots circular>
			<swiper-item v-for="item in swipers" :key="item.id">
				<image :src="item.img"/>
			</swiper-item>
		</swiper>
		<!-- 导航区 -->
		<view class="nav">
			<view class="nav-item" v-for="(item, index) in navs" :key="index" @click="navItemHandel(item.url)">
				<view class="iconfont" :class="item.icon"></view>
				<text>{{item.title}}</text>
			</view>
		</view>
		<!-- 推荐商品 -->
		<view class="hot-goods">
			<view class="tit">
				<text>推荐商品</text>
			</view>
			<goods-list :list="list" @navigator="goGoodsDetail"></goods-list>
		</view>
	</view>
</template>

<script>
	import goodsList from '../../components/goods_list.vue'
	export default {
		data() {
			return {
				// 轮播图数据
				swipers: [],
				// 商品列表数据
				list: [],
				// 导航数据
				navs: [{
					icon:'icon-ziyuan',
					title:'黑马超市',
					url:'/pages/goods/goods'
				},{
					icon:'icon-guanyuwomen',
					title:'联系我们',
					url:'/pages/contact/contact'
				},{
					icon:'icon-tupian',
					title:'社区图片',
					url:'/pages/pics/pics'
				},{
					icon:'icon-shipin',
					title:'学习视频',
					url:'/pages/videos/videos'
				}]
			}
		},
		onLoad() {
			this.getSwipers()
			this.getList()
		},
		methods: {
			// 获取轮播图数据
			async getSwipers() {
				const res =  await this.$myRequest({
					url:'/api/getlunbo'
				})
				this.swipers = res.data.message
				console.log(this.swipers)
			},
			// 获取商品列表
			async getList() {
				const res =  await this.$myRequest({
					url:'/api/getgoods?pageindex=1'
				})
				this.list = res.data.message
				console.log(this.list)
			},
			//导航栏点击事件
			 navItemHandel(url) {
				 uni.navigateTo({
				 	url: url
				 })
			 },
			 // 跳转到商品详情页
			 goGoodsDetail(id) {
			 	uni.navigateTo({
			 		url:'/pages/goods-detail/goods-detail?id='+id
			 	})
			 }
		},
		components:{
			"goods-list": goodsList
		}
	}
</script>

<style lang="scss">
	swiper {
		width: 750rpx;
		height: 380rpx;
		image {
			width: 100%;
			height: 100%;
		}
	}
	// 导航区
	.nav {
		display: flex;
		.nav-item {
			width: 25%;
			text-align: center;
			.iconfont {
				margin: 20rpx auto;
				width: 120rpx;
				height: 120rpx;
				border-radius: 60rpx;
				background-color: $index-color;
				color: #fff;
				font-size: 50rpx;
				line-height: 120rpx;
			}
			.icon-tupian {
				font-size: 45rpx;
			}
			text {
				font-size: 30rpx;
			}
		}
	}
	// 推荐商品
	.hot-goods {
		margin-top: 10px;
		background-color: #eee;
		overflow: hidden;
		.tit {
			height: 100rpx;
			margin: 6rpx 0;
			font-size: 36rpx;
			line-height: 100rpx;
			color: $index-color;
			letter-spacing: 40rpx;
			text-align: center;
			background-color: #fff;
		}
		
	}
</style>
