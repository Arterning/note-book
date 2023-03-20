<template>
	<view class="content">
		<view class="logo">
			<image src="http://www.itcast.cn/2018czydz/images/gywmban.jpg"></image>
		</view>
		<view class="message">
			<view class="phone">
				<text @click="callPhone">联系电话：400-618-9090(点击拨打)</text>
			</view>
			<view class="address">
				<text>校区地址：浙江省杭州市下沙经济开发区4号大街187号盛泰时代山</text>
			</view>
		</view>
		<map :longitude="longitude" :latitude="latitude" :markers="covers"></map>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				longitude: 116.635672,
				latitude: 40.169419,
				covers:[ {
				width:30,
				height:30,
                latitude: 40.169419,
                longitude: 116.635672,
                iconPath: '../../static/hmlogo.png'
				}]
			}
		},
		methods: {
			callPhone() {
				uni.makePhoneCall({
					phoneNumber: '000000000'
				})
			}
		}
	}
</script>

<style lang="scss">
.content {
	font-size: 30rpx;
	.logo {
		image {
			width: 750rpx;
			height: 320rpx;
		}
	}
	.message {
		margin: 10rpx 20rpx;
		.phone {
			line-height: 80rpx;
			border-bottom: 1px solid #eee;
		}
		.address {
			line-height: 80rpx;
			border-bottom: 1px solid #eee;
		}
	}
	map {
		width: 750rpx;
		height: 750rpx;
	}
}
</style>
