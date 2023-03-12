<!--
 * @Author: guoxi
 * @Date: 2021-05-24 09:55:30
 * @LastEditTime : 2021-10-21 14:47:14
 * @LastEditors  : Please set LastEditors
 * @Description: 检测报告
 * @FilePath     : \blockchain-minicode\src\pages\report\BlockChainReport.vue
-->

<template>
  <view class="report">
    <view
      v-if="inspection_report.quality_report_image_url.length > 0"
      class="report-title"
    >
      品质检测报告
    </view>
    <view
      v-if="inspection_report.quality_report_image_url.length > 0"
      class="report-card"
    >
      <view
        v-for="(item, index) in inspection_report.quality_report_image_url"
        :key="index"
        class="card-img"
        ><img
          :src="item"
          @click="preViewImg(index, inspection_report.quality_report_image_url)"
      /></view>
    </view>

    <view
      v-if="inspection_report.safe_report_image_url.length > 0"
      class="report-title"
    >
      安全检测报告
    </view>
    <view
      v-if="inspection_report.safe_report_image_url.length > 0"
      class="report-card"
    >
      <view
        v-for="(item, index) in inspection_report.safe_report_image_url"
        :key="index"
        class="card-img"
        ><img
          :src="item"
          @click="preViewImg(index, inspection_report.safe_report_image_url)"
      /></view>
    </view>

    <view
      v-if="inspection_report.water_report_image_url.length > 0"
      class="report-title"
    >
      水质检测报告
    </view>
    <view
      v-if="inspection_report.water_report_image_url.length > 0"
      class="report-card"
    >
      <view
        v-for="(item, index) in inspection_report.water_report_image_url"
        :key="index"
        class="card-img"
        ><img
          :src="item"
          @click="preViewImg(index, inspection_report.water_report_image_url)"
      /></view>
    </view>

    <view
      v-if="inspection_report.soil_report_image_url.length > 0"
      class="report-title"
    >
      土壤检测报告
    </view>
    <view
      v-if="inspection_report.soil_report_image_url.length > 0"
      class="report-card"
    >
      <view
        v-for="(item, index) in inspection_report.soil_report_image_url"
        :key="index"
        class="card-img"
        ><img
          :src="item"
          @click="preViewImg(index, inspection_report.soil_report_image_url)"
      /></view>
    </view>
  </view>
</template>

<script>
import {
  stringifyDate,
  preViewImg,
  getImageUrl,
  getZNImageUrl
} from '../../utils/tool'
export default {
  name: 'Index',
  onLoad(option) {
    const flag = option.flag
    this.isFlag = flag === '1' ? false : true
  },
  data() {
    return {
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

  created() {
    if (this.isFlag) {
      const data = uni.getStorageSync('h5Info')
      this.creator = data.brandName
      this.tx_create_time = data.inspection_report.tx_create_time
      this.tx_id = data.inspection_report.tx_id
      this.block_number = data.inspection_report.block_number

      this.inspection_report.quality_inspection_date = stringifyDate(
        data.inspection_report.quality_inspection_date
      )
      if (this.validation(data.inspection_report.quality_report_image_url)) {
        const urls = JSON.parse(data.inspection_report.quality_report_image_url)
        urls.map(el => {
          this.inspection_report.quality_report_image_url.push(
            getImageUrl(el.id)
          )
        })
      }

      this.inspection_report.safe_inspection_date = stringifyDate(
        data.inspection_report.safe_inspection_date
      )
      if (this.validation(data.inspection_report.safe_report_image_url)) {
        const urls = JSON.parse(data.inspection_report.safe_report_image_url)
        urls.map(el => {
          this.inspection_report.safe_report_image_url.push(getImageUrl(el.id))
        })
      }

      this.inspection_report.water_inspection_date = stringifyDate(
        data.inspection_report.water_inspection_date
      )

      if (this.validation(data.inspection_report.water_report_image_url)) {
        const urls = data.inspection_report.water_report_image_url.split(',')
        urls.map(el => {
          this.inspection_report.water_report_image_url.push(getZNImageUrl(el))
        })
      }

      this.inspection_report.soil_inspection_date = stringifyDate(
        data.inspection_report.soil_inspection_date
      )
      if (this.validation(data.inspection_report.soil_report_image_url)) {
        const urls = data.inspection_report.soil_report_image_url.split(',')
        urls.map(el => {
          this.inspection_report.soil_report_image_url.push(getZNImageUrl(el))
        })
      }
    }
  },
  methods: {
    validation(val) {
      if (val !== 'null' && val !== '' && val !== '<invalid Value>') {
        return true
      }
      return false
    },
    preViewImg(i, images) {
      preViewImg(i, images)
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
  background-color: #fafafa;
  padding-left: 33rpx;
  font-family: PingFangSC-Regular;
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
