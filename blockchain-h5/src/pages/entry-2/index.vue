<!--
 * @Author       : guoxi
 * @Date         : 2021-05-24 09:55:30
 * @LastEditlastTime : 2021-11-19 16:08:41
 * @LastEditors  : Your Name
 * @Description  : 入口页
 * @FilePath     : \blockchain-h5\src\pages\entry-2\index.vue
-->
<template>
  <view class="index">
    <view class="entry-2">
      <!-- 认证信息 -->
      <Authentication :mainInfoData="mainInfoData" />
    </view>
    <view class="bottom-btn" @click.native="toHome">
      <view class="btn">深度溯源</view>
    </view>
  </view>
</template>

<script>
import Authentication from '../module/authentication.vue' // 认证信息
export default {
  name: 'Index',

  components: {
    Authentication
  },

  data() {
    return {
      isShow: true
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    // 溯源码主要信息
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
    /**
     * @description: 跳转首页
     * @author: Hemingway
     */
    toHome() {
      this.$Router.push({
        path: '/home',
        query: {
          id: this.qrcode
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.index {
  .entry-2 {
    padding: 16rpx;
  }

  .bottom-btn {
    box-sizing: border-box;
    width: 100%;
    height: 138rpx;
    padding: 24rpx;

    .btn {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100%;
      font-size: 32rpx;
      color: #fff;
      background-color: #00c853;
      border-radius: 44rpx;
    }

    .btn::after {
      border: none;
    }
  }
}
</style>
