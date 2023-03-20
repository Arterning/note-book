<template>
	<view class="news">
		<view class="list">
			<view class="item" v-for="item in newsList" :key="item.id" @click="toNewsDetail(item.id)">
				<image :src="item.img_url" />
				<view class="txt">
					<text class="title">
						{{item.title}}
					</text>
					<view class="info">
						<text>发表时间：{{item.add_time}}</text>
						<text>浏览：{{item.click}}</text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				newsList: []
			}
		},
		methods: {
			async getNewsList() {
				const res = await this.$myRequest({
					url:'/api/getnewslist'
				})
				this.newsList = res.data.message
			},
			toNewsDetail(id) {
				uni.navigateTo({
					url:'../news-detail/news-detail?id='+id
				})
			}
		},
		onLoad() {
			this.getNewsList()
		}
	}
</script>

<style lang="scss">
	.list {
		display: flex;
		flex-wrap: wrap;
		.item {
			display: flex;
			padding: 0 18rpx;
			width: 750rpx;
			height: 162rpx;
			border-bottom: 1px solid $index-color;
			image {
				display: block;
				margin: auto;
				min-width: 192rpx;
				min-height: 144rpx;
				max-width: 192rpx;
				max-height: 144rpx;
			}
			.txt {
				display: flex;
				flex-direction: column;
				margin: 9rpx 0 9rpx 15rpx;
				justify-content: space-between;
				font-size: 22rpx;
				.title {
					font-size: 28rpx;
				}
				.info text:nth-child(2){
					margin-left: 30rpx;
				}
			}
		}
	}
</style>
