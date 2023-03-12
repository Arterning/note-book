<!--
 * @Author: guoxi
 * @Date: 2021-05-24 14:06:36
 * @LastEditTime : 2021-10-18 21:48:50
 * @LastEditors  : Please set LastEditors
 * @Description: 商品详情页
 * @FilePath     : \blockchain-minicode\src\pages\commodity\CommodityDetail.vue
-->

<template>
  <view class="content">
    <view class="comm__row b_bottom">
      <text class="comm-title">商品信息</text>
      <text class="comm-detail">{{ commodityData.commodityName }}</text>
    </view>
    <view class="comm__row">
      <text class="comm-title">所属品牌</text>
      <text class="comm-detail">{{ commodityData.brandName }}</text>
    </view>
    <view class="comm__row b_bottom">
      <text class="comm-title">净含量</text>
      <text class="comm-detail">{{
        getNetContent(commodityData.netContent)
      }}</text>
    </view>
    <view class="comm__row b_bottom">
      <text class="comm-title">质量等级</text>
      <text class="comm-detail">{{ commodityData.qualityGrade }}</text>
    </view>

    <view class="comm__row">
      <text class="comm-title">保质期</text>
      <text class="comm-detail">{{ commodityData.qualityGuaPeriod }}月</text>
    </view>
  </view>
</template>

<script>
import { detailCommodity } from '@/api/commodity'
export default {
  name: 'CommodityDetail',
  data() {
    return {
      commodityId: '',
      commodityData: {
        brandName: '',
        commodityName: '',
        netContent: '',
        qualityGrade: '',
        qualityGuaPeriod: ''
      }
    }
  },

  onLoad(option) {
    this.commodityId = option.id
  },

  onShow() {
    this.query()
  },

  methods: {
    async query() {
      let res = await detailCommodity(this.commodityId)
      if (res.code === '0') this.commodityData = { ...res.resObj }
    },
    getNetContent(netContent) {
      let str = ''

      if (netContent) {
        str += JSON.parse(netContent)
          .map(el => `${el.weightValue}kg/${el.weightUnit}`)
          .join('; ')
      }

      return str
    }
  }
}
</script>

<style lang="scss" scoped>
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0 34rpx;
  border-top: 1px solid #eee;
  border-bottom: 1px solid #eee;

  .comm__row {
    width: 100%;
    height: 104rpx;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .comm-title {
      font-size: 28rpx;
      color: #8b8b8b;
    }

    .comm-detail {
      font-size: 28rpx;
      color: #212121;
    }
  }

  .b_bottom {
    border-bottom: 1px solid #eee;
  }
}
</style>
