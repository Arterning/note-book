<template>
	<view>
		<view class="article">
			<view class="title">
				<text>{{newsData.title}}</text>
			</view>
			<view class="date">
				<text>发表时间：{{newsData.add_time | formatDate}}</text>
				<text>浏览：{{newsData.click}}</text>
			</view>
			<view class="content">
				<rich-text :nodes="newsData.content" space></rich-text>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				pageId: 0,
				newsData: {}
			}
		},
		methods: {
			async getNews() {
				const res = await this.$myRequest({
					url:'/api/getnew/'+this.pageId.id
				})
				this.newsData = res.data.message[0]
				console.log(this.newsData)
			}
		},
		onLoad(opthon) {
			this.pageId = opthon
			this.getNews()
		}
	}
</script>

<style lang="scss">
.article {
	font-size: 15px;
	padding: 0 10px;
	.title {
		margin: 20rpx 0;
		text-align: center;
	}
	.date {
		display: flex;
		justify-content: space-between;
	}
	
}
</style>
