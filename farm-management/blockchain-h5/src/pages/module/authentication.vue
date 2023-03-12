<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-05 16:14:17
 * @FilePath     : \blockchain-h5\src\pages\module\authentication.vue
-->
<template>
  <view>
    <!-- 认证信息 -->
    <view class="authentication-info common-box">
      <view class="img-box" :style="bgImg">
        <image
          src="../../static/template1-img/区块链认证合格.png"
          class="status-img"
        />
      </view>
      <view class="content-box">
        <view class="content-title"
          >{{ initData.brandName }}·{{ initData.commodityName }}·{{
            initData.unit
          }}</view
        >
        <view class="content-logo">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 61.26 64"
            class="status-img"
          >
            <g id="图层_2" data-name="图层 2">
              <g id="图层_1-2" data-name="图层 1">
                <path
                  id="认证"
                  d="M61.23,5.66a2.54,2.54,0,0,0-.89-1.58c-.1-.1-1.68-1.29-4.15-3.17A2.54,2.54,0,0,0,53.82.62C40.27,6.45,32.46.82,32.06.52a2.25,2.25,0,0,0-2.87,0c-.09.1-8,5.93-21.75.1A2.54,2.54,0,0,0,5.07.91C2.6,2.79,1,4,.92,4.08A2.49,2.49,0,0,0,0,5.66,2.33,2.33,0,0,0,.52,7.44c.1.1,5.14,6.63,2.77,16.51C1.9,29.79.42,35.82,1.71,41.46,3.19,48.18,8.23,52.93,16.93,56H17A25.85,25.85,0,0,1,28.7,63.11a2.43,2.43,0,0,0,3.76,0A25,25,0,0,1,44.13,56h.09C53,52.93,58,48.18,59.45,41.46c1.29-5.74-.2-11.77-1.58-17.51A20.31,20.31,0,0,1,60.64,7.44,2.09,2.09,0,0,0,61.23,5.66Zm-43,6.82c1.38,1.88,2.77,4.06,4.25,6.43-1.09.59-2.47,1.39-4.25,2.47-1.29-2.27-2.57-4.45-3.86-6.52ZM15.15,41.26A3.83,3.83,0,0,0,16,38.69V27.51H13.28V23h7.51V36.81l3.07-2.47c.09,1,.39,2.57.79,4.74-2.18,1.58-4.65,3.46-7.42,5.84Zm32.54,4.05H23.06V40.66h2.87V23.26h4.65V40.57h3.76V19.41H24.15V14.86H47.29v4.55H39.08v7.81H46.5v4.64H39.08v8.8h8.61Z"
                />
              </g>
            </g>
          </svg>
          {{ initData.brandName }}正品
        </view>
        <view class="content-info">
          <view v-if="initData.harvestDate">
            <view space="emsp" decode="true" class="content-info-title"
              >收获时间: {{ '' }}</view
            >
            <view>{{ initData.harvestDate }}</view>
          </view>
          <view v-if="initData.chainCode">
            <view space="emsp" decode="true" class="content-info-title"
              >溯源编号: {{ '' }}</view
            >
            <view>{{ initData.chainCode }}</view>
          </view>
          <!-- <view v-if="initData.hashCode">
            <view space="emsp" decode="true" class="content-info-title"
              >区块链ID: {{ '' }}</view
            >
            <view>{{ initData.hashCode }}</view>
          </view> -->
          <view>
            <view space="emsp" decode="true" class="content-info-title"
              >区块链ID: {{ '' }}</view
            >
            <view
              >86ad770e04a2a104f79dd69c0123e275187fc3661e327d6439afcf77caf07910</view
            >
          </view>
        </view>
      </view>
      <view class="query-record" @click="jump('/entry-1/queryRecord')">
        <text>查询记录</text>
        <u-icon
          name="arrow-right"
          size="16"
          color="#00c853"
          class="icon"
        ></u-icon>
      </view>
    </view>
  </view>
</template>

<script>
import '@/static/json/static.json'
import { getImageUrl } from '../../utils/tool'
export default {
  name: 'Authentication',
  props: ['mainInfoData'],
  data() {
    return {
      bgImg: {
        backgroundImage: ''
      },
      initData: this.mainInfoData
    }
  },

  computed: {
    isStaticData() {
      return this.$store.getters.isStaticData
    },
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    }
  },

  created() {},

  mounted() {
    // 修改背景背景图片路径
    // let imgId =
    //   this.mainInfoData.commodityImageUrl &&
    //   JSON.parse(this.mainInfoData.commodityImageUrl)[0].id
    this.bgImg.backgroundImage = `url(${getImageUrl(
      this.mainInfoData.commodityImageUrl
    )})`
    if (this.isStaticData) {
      let { authentication } = require('../../static/json/static.json')
      this.initData = authentication
      this.bgImg.backgroundImage = `url(static/img/background-img.jpg)`
    }
  },

  methods: {
    /**
     * @description: 跳转二级页面
     * @param {String} path
     * @param {String} type
     * @author: Hemingway
     */
    jump(path) {
      this.$Router.push({
        path,
        query: {
          id: this.qrcode
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.common-box {
  border: 1px solid #eeeeee;
  background-color: white;
  border-radius: 16rpx;
  width: 100%;
  box-sizing: border-box;
}
.authentication-info {
  .img-box {
    position: relative;
    background: no-repeat;
    background-size: 100% 100%;
    height: 0;
    padding-bottom: 100%;
    width: 100%;
    border-radius: 16rpx 16rpx 0 0;
    .logo-img {
      position: absolute;
      top: 22rpx;
      left: 22rpx;
      width: 142rpx;
      height: 30rpx;
    }
    .status-img {
      position: absolute;
      bottom: 22rpx;
      right: 22rpx;
      width: 190rpx;
      height: 120rpx;
    }
  }
  .content-box {
    padding: 32rpx 32rpx 0 32rpx;
    border-bottom: 1px solid #eeeeee;
    .content-title {
      font-weight: 700;
      font-size: 32rpx;
    }
    .content-logo {
      background-color: #e5f9ed;
      color: #00c853;
      padding: 20rpx;
      font-size: 28rpx;
      border-radius: 16rpx;
      display: table;
      margin: 24rpx 0;
      font-weight: 700;
      .status-img {
        width: 40rpx;
        height: 36rpx;
        vertical-align: middle;
        fill: #00c853;
        margin-right: 6rpx;
      }
    }
    .content-info {
      font-size: 28rpx;
      > view {
        margin-bottom: 20rpx;
        display: flex;
        > view {
          width: calc(100% - 140rpx);
          word-break: break-all;
        }
        > .content-info-title {
          width: 140rpx;
          color: #8b8b8b;
        }
      }
    }
  }
  .query-record {
    padding: 35rpx;
    text-align: center;
    font-size: 32rpx;
    color: #00c853;
    font-weight: 700;

    .icon {
      display: inline-block;
      margin-left: 10rpx;
    }
  }
}
</style>
