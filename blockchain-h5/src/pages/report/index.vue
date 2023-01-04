<!--
 * @Author       : guoxi
 * @Date         : 2021-05-24 09:55:30
 * @LastEditTime : 2022-07-07 13:44:32
 * @LastEditors  : Your Name
 * @Description  : 检测报告
 * @FilePath     : \blockchain-h5\src\pages\report\index.vue
-->
<template>
  <view class="report">
    <Quality
      v-if="configInfor.quality.isShow"
      :menuTree="configInfor.quality.subTree"
    />
    <!-- <bottom-btn text="点击查看产品区块链证书" @click.native="onView" /> -->
  </view>
</template>

<script>
import { preViewImg } from '../../utils/tool'
import Quality from '../module/quality.vue' // 质检信息

export default {
  name: 'Report',
  components: { Quality },
  data() {
    return {
      creator: '',
      tx_create_time: '',
      tx_id: '',
      block_number: '',
      list: [] // 图片
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    configInfor() {
      return this.$store.getters.configInfor
    },
    mainInfoData() {
      return this.$store.getters.mainInfoData
    }
  },

  created() {
    uni.setNavigationBarTitle({
      title: `${this.mainInfoData.brandName}区块链溯源`
    })
  },

  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    },

    /**
     * @description: 跳转区块链证书页面
     * @author: guoxi
     */
    onView() {
      uni.navigateTo({
        url: `/certificate?id=${this.qrcode}&cername=检测报告区块链&tx_id=${this.tx_id}&block_number=${this.block_number}&tx_create_time=${this.tx_create_time}&creator=${this.creator}`
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.report {
  padding: 16rpx;
  margin-top: -22rpx;
  // ::v-deep .uni-scroll-view {
  //   height: calc(100vh - 138rpx);

  //   .report-title {
  //     position: relative;
  //     width: 750rpx;
  //     height: 80rpx;
  //     line-height: 80rpx;
  //     padding-left: 32rpx;
  //     font-size: 24rpx;
  //     color: #8b8b8b;
  //   }

  //   .report-card {
  //     margin: 0 -14rpx;
  //     padding: 32rpx 32rpx 0;
  //     background-color: #fff;
  //     display: flex;
  //     flex-wrap: wrap;

  //     .card-img {
  //       width: 210rpx;
  //       height: 280rpx;
  //       margin: 0 14rpx 32rpx;
  //       border: 2rpx solid #eee;
  //       box-sizing: border-box;
  //     }
  //   }

  //   .postscript {
  //     padding: 30rpx 32rpx;
  //     font-size: 24rpx;
  //     color: #212121;

  //     .text:not(:first-of-type) {
  //       margin-top: 14rpx;
  //     }
  //   }
  // }
}
</style>
