<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 14:07:34
 * @FilePath     : \blockchain-h5\src\pages\entry-1\queryRecord\index.vue
-->
<template>
  <view class="query-record">
    <view class="common-box">
      <table class="table">
        <colgroup>
          <col width="20%" />
          <col width="40%" />
          <col width="40%" />
        </colgroup>
        <thead>
          <tr>
            <td>
              查询次数
            </td>
            <td>
              查询时间
            </td>
            <td>
              位置
            </td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in scanInfoData" :key="index">
            <td>
              {{ scanInfoData.length - index }}
            </td>
            <td>
              {{ item.visitTime }}
            </td>
            <td>
              {{ item.address || '--' }}
            </td>
          </tr>
        </tbody>
      </table>
    </view>
    <view class="back-btn" @click="handleBack">
      <u-icon name="close" size="20" color="white" class="icon"></u-icon>
    </view>
  </view>
</template>

<script>
import { getScanInfo } from '@/api/myRice'
export default {
  name: 'QueryRecord',
  components: {},

  data() {
    return {
      scanInfoData: []
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    // 是否请求接口
    isStaticData() {
      return this.$store.getters.isStaticData
    },
    themeBackground() {
      let colorId = this.$store.getters.colorId.toString()
      console.log('颜色', colorId)
      if (colorId === '2') {
        return 'theme1-background'
      } else if (colorId === '3') {
        return 'theme2-background'
      } else {
        return ''
      }
    }
  },

  created() {
    if (this.isStaticData) {
      let {
        authentication: { queryRecord }
      } = require('../../../static/json/static.json')
      this.scanInfoData = queryRecord
      return
    }
    this.scanInfo()
  },

  mounted() {},

  methods: {
    /**
     * @description: // 获取扫码星系
     * @author: qjj
     */
    async scanInfo() {
      const { data } = await getScanInfo({
        qrCode: this.qrcode
      })
      if (!data) return
      this.scanInfoData = data
    },
    /**
     * @description: // 回退按钮
     * @author: qjj
     */
    handleBack() {
      window.history.go(-1)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../../static/style/template1.scss';
.query-record {
  background-color: #fafafa;
  padding: 16rpx;
  .common-box {
    border: 1px solid #eeeeee;
    background-color: white;
    border-radius: 16rpx;
    width: 100%;
    box-sizing: border-box;

    .table {
      width: 100%;
      font-size: 24rpx;
      text-align: center;
      border-collapse: collapse;
      thead {
        color: #008080;
        tr > td {
          border-bottom: 1px solid #00bfbf;
          padding: 20rpx 0;
        }
      }
      tbody > tr > td {
        padding: 20rpx 0;
      }
    }
  }

  .back-btn {
    position: fixed;
    right: 0;
    bottom: 0;
    width: 80rpx;
    height: 80rpx;
    border-radius: 80rpx 0 0 0;
    background: red;
    .icon {
      position: absolute;
      right: 12rpx;
      bottom: 12rpx;
    }
  }
}
</style>
