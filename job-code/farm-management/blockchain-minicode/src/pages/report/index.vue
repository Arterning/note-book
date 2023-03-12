<!--
 * @Author: guoxi
 * @Date: 2021-05-24 09:55:30
 * @LastEditTime : 2022-07-06 11:34:12
 * @LastEditors  : Your Name
 * @Description: 检测报告
 * @FilePath     : \blockchain-minicode\src\pages\report\index.vue
-->

<template>
  <view class="report">
    <view class="report-title">
      品质检测报告
      <view class="report-time">{{
        inspection_report.quality_inspection_date
      }}</view>
    </view>
    <view class="report-card">
      <view
        v-for="(item, index) in inspection_report.quality_report_image_url"
        :key="index"
        class="card-img"
        ><img
          :src="getImage(item)"
          @tap="preViewImg(i, inspection_report.quality_report_image_url)"
      /></view>
    </view>

    <view class="report-title">
      安全检测报告
      <view class="report-time">{{
        inspection_report.safe_inspection_date
      }}</view>
    </view>
    <view class="report-card">
      <view
        v-for="(item, index) in inspection_report.safe_report_image_url"
        :key="index"
        class="card-img"
        ><img
          :src="getImage(item)"
          @tap="preViewImg(i, inspection_report.safe_report_image_url)"
      /></view>
    </view>
  </view>
</template>

<script>
import { getImageUrl, preViewImg, stringifyDate } from '../../utils/tool'
export default {
  name: 'Index',

  data() {
    return {
      target: null,
      creator: '',
      tx_create_time: '',
      tx_id: '',
      block_number: '',
      inspection_report: {
        quality_inspection_date: '',
        quality_report_image_url: [],

        safe_inspection_date: '',
        safe_report_image_url: [],

        water_inspection_date: '',
        water_report_image_url: [],
        soil_inspection_date: '',
        soil_report_image_url: []
      },
      isFlag: true
    }
  },
  onLoad() {
    let query = this.$Route.query
    const flag = query.flag
    this.isFlag = flag === '1' ? false : true
    this.target = query
  },
  onShow() {
    this.initData()
  },

  methods: {
    initData() {
      if (!this.isFlag) {
        this.inspection_report.quality_inspection_date = stringifyDate(
          this.target.pzjcDate
        )
        this.inspection_report.quality_report_image_url = this.target.pzjcAttachJson

        this.inspection_report.safe_inspection_date = stringifyDate(
          this.target.aqjcDate
        )
        this.inspection_report.safe_report_image_url = this.target.aqjcAttachJson
      }
    },
    getImage(item) {
      return getImageUrl(item.id)
    },
    preViewImg(i, images) {
      const urls = images.map(el => this.getImage(el))
      preViewImg(i, urls)
    }
  }
}
</script>

<style lang="scss" scoped>
.report-title {
  position: relative;
  width: 750rpx;
  height: 80rpx;
  line-height: 80rpx;
  padding-left: 32rpx;
  font-size: 24rpx;
  color: #212121;

  .report-time {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    right: 73rpx;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #8b8b8b;
  }
}

.report-card {
  padding: 32rpx;
  background-color: #fff;

  .card-img {
    display: inline-block;
    margin-right: 32rpx;
    width: 112rpx;
    height: 112rpx;

    img {
      width: 100%;
      height: 100%;
    }
  }
}

.add-brand {
  width: 750rpx;
  height: 138rpx;
  position: fixed;
  bottom: 0;
  background-color: #fafafa;

  .btn {
    border-radius: 44rpx;
    font-size: 32rpx;
    margin: 25rpx 40rpx;
    background-color: #00c853;

    span {
      font-family: PingFangSC-Regular;
      font-size: 32rpx;
      color: #fff;
      margin-left: 10rpx;
    }
  }
}
</style>
