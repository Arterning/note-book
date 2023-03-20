<template>
	<view class="pics">
		<scroll-view class="left" scroll-y>
			<view :class="index === currentIndex ? 'active' : '' " @click="clickItemHandel(index,item.id)" v-for="(item, index) in cates" :key="index">
				{{item.title}}
			</view>
		</scroll-view>
		<scroll-view class="right" scroll-y>
			<view class="item" v-for="item in cateItems" :key="item.id">
				<image :src="item.img_url" @click="viewImg(item.img_url)"></image>
				<text>{{item.title}}</text>
			</view>
			<view v-if="cateItems.length === 0">
				<text>暂时没有资源~~~</text>
			</view>
		</scroll-view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				cates: [],
				currentIndex: 0,
				cateItems: []
			}
		},
		methods: {
			async getImgList() {
				const res = await this.$myRequest({
					url:'/api/getimgcategory'
				})
				// console.log(res)
				this.cates = res.data.message
			},
			async clickItemHandel(index,id) {
				this.currentIndex = index
				const res = await this.$myRequest({
					url:'/api/getimages/'+id
				})
				this.cateItems = res.data.message
				// console.log(this.cateItems)
			},
			viewImg(current) {
				const urls = this.cateItems.map(item => {
					return item.img_url
				})
				uni.previewImage({
					current: current,
					urls
				})
			}
			
		},
		onLoad() {
			this.getImgList()
		}
	}
</script>

<style lang="scss">
page {
	height: 100%;
}
.pics {
	display: flex;
	height: 100%;
	.left {
		height: 100%;
		view {
			width: 200rpx;
			height: 61px;
			line-height: 61px;
			text-align: center;
			border-top: 1px solid #eee;
			border-right: 1px solid #eee;
		}
		.active {
			background-color: $index-color;
			color: #fff;
		}
	}
	.right {
		height: 100%;
		.item {
			margin: 10rpx;
			width: 530rpx;
			image {
				margin: auto;
				width: 520rpx;
				height: 520rpx;
				border-radius: 8rpx;
			}
			text {
				font-size: 30rpx;
			}
		}
	}
}
</style>
