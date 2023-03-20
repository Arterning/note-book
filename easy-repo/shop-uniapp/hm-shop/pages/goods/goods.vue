<template>
	<view class="bgc">
		<goods-list :list="list" @navigator="goGoodsDetail"></goods-list>
		<view class="fotter" v-if="listEnd">
			<text>我是有底线的！！！</text>
		</view>
	</view>
</template>

<script>
	import goodsList from '../../components/goods_list.vue'
	export default {
		data() {
			return {
				pageindex: 1,
				list: [],
				listEnd: false
			}
		},
		methods: {
			// 获取商品列表
			async getList(callback) {
				const res =  await this.$myRequest({
					url:'/api/getgoods?pageindex='+this.pageindex
				})
				this.list = [...this.list, ...res.data.message]
				callback && callback()
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
		},
		onLoad() {
			this.getList()
		},
		// 触底事件
		onReachBottom() {
			if (this.list.length < this.pageindex*10) return this.listEnd = true
			this.pageindex ++
			this.getList()
		},
		// 下拉刷新事件
		onPullDownRefresh() {
			this.list = []
			this.pageindex = 1
			this.listEnd = false
			this.getList(()=> {
				uni.stopPullDownRefresh()
			})
		}
	}
</script>

<style>
.bgc {
		background-color: #eee;
	}
.fotter {
			width: 100%;
			height: 60rpx;
			line-height: 60rpx;
			text-align: center;
			font-size: 36rpx;
			color: #808080;
}
</style>
