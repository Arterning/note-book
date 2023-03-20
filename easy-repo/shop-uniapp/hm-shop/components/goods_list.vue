<template>
		<view class="list">
			<view class="list-item" v-for="item in list" :key="item.id" @click="goDetail(item.id)">
				<image :src="item.img_url" mode=""></image>
				<view class="price">
					<text>{{item.sell_price}}</text>
					<text>{{item.market_price}}</text>
				</view>
				<view class="name">
					<text>{{item.title}}</text>
				</view>
			</view>
		</view>
</template>

<script>
	export default{
		props: ["list"],
		methods: {
			goDetail(id) {
				this.$emit("navigator",id)
			}
		}
	}
</script>

<style lang="scss">
	.list {
		display: flex;
		justify-content: space-between;
		flex-wrap: wrap;
		padding: 0 14rpx;
		.list-item {
			width: 355rpx;
			margin: 10rpx 0;
			padding: 14rpx;
			box-sizing: border-box;
			background-color: #fff;
			image {
				width: 100%;
				height: 300rpx;
			}
			.price {
				color: $index-color;
				font-size: 36rpx;
				text:nth-child(2) {
					color: #ccc;
					font-size: 28rpx;
					margin-left: 16rpx;
					text-decoration: line-through;
				}
			}
			.name {
				font-size: 28rpx;
				padding: 10rpx 0 14rpx 0;
			}
		}
	}
</style>
