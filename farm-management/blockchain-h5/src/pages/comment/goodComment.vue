<!--
 * @Author       : qinjj
 * @Date         : 2021-05-24 09:55:30
 * @LastEditlastTime : 2021-11-19 16:08:41
 * @LastEditors  : Your Name
 * @Description  : 入口页
 * @FilePath     : \blockchain-h5\src\pages\comment\goodComment.vue
-->
<template>
  <view class="goodComment">
    <view class="comment-item">
      <text class="comment-title">新鲜度：</text>
      <u-rate
        :count="count"
        v-model="freshnessValue"
        class="comment-icon"
        activeIcon="thumb-up-fill"
        inactiveIcon="thumb-up"
        activeColor="#5ac725"
        size="26"
      ></u-rate>
      <text class="comment-text">{{ freshnessText }}</text>
    </view>
    <view class="comment-item">
      <text class="comment-title" decode>外&emsp;观：</text>
      <u-rate
        :count="count"
        v-model="appearanceValue"
        class="comment-icon"
        activeIcon="thumb-up-fill"
        inactiveIcon="thumb-up"
        activeColor="#5ac725"
        size="26"
      ></u-rate>
      <text class="comment-text">{{ appearanceText }}</text>
    </view>
    <view class="comment-item">
      <text class="comment-title" decode>口&emsp;味：</text>
      <u-rate
        :count="count"
        v-model="flavorValue"
        class="comment-icon"
        activeIcon="thumb-up-fill"
        inactiveIcon="thumb-up"
        activeColor="#5ac725"
        size="26"
      ></u-rate>
      <text class="comment-text">{{ flavorText }}</text>
    </view>
    <view>
      <view style="marginBottom: 10rpx;fontSize: 30rpx">
        评价理由：
      </view>
      <u--textarea
        v-model="reason"
        placeholder="请输入评价理由，如：大米很新鲜，外观晶莹剔透，口味软糯，值得推荐。"
        height="200"
        count
        maxlength="500"
      ></u--textarea>
    </view>
    <view class="footer-box">
      <u-button
        type="success"
        text="提交"
        class="footer-btn"
        @click="submit"
      ></u-button>
      <u-button
        type="success"
        :plain="true"
        text="取消"
        class="footer-btn"
        @click="backBtn"
      ></u-button>
    </view>
    <u-toast ref="goodToast"></u-toast>
  </view>
</template>

<script>
import { submitGoodComment } from '@/api/myRice'
export default {
  name: 'goodComment',

  data() {
    return {
      count: 5,
      freshnessValue: 0,
      appearanceValue: 0,
      flavorValue: 0,
      reason: ''
    }
  },

  computed: {
    qrcode() {
      return this.$store.getters.code
    },
    freshnessText() {
      let result = ''
      switch (this.freshnessValue) {
        case 1:
          result = '非常差'
          break
        case 2:
          result = '差'
          break
        case 3:
          result = '一般'
          break
        case 4:
          result = '比较满意'
          break
        case 5:
          result = '非常满意'
          break
      }
      return result
    },
    appearanceText() {
      let result = ''
      switch (this.appearanceValue) {
        case 1:
          result = '非常差'
          break
        case 2:
          result = '差'
          break
        case 3:
          result = '一般'
          break
        case 4:
          result = '比较满意'
          break
        case 5:
          result = '非常满意'
          break
      }
      return result
    },
    flavorText() {
      let result = ''
      switch (this.flavorValue) {
        case 1:
          result = '非常差'
          break
        case 2:
          result = '差'
          break
        case 3:
          result = '一般'
          break
        case 4:
          result = '比较满意'
          break
        case 5:
          result = '非常满意'
          break
      }
      return result
    }
  },

  created() {},

  methods: {
    /**
     * @description: 提交按钮事件
     * @author: qinjj
     */
    async submit() {
      if (this.freshnessValue === 0) {
        this.$refs.goodToast.show({
          type: 'error',
          icon: false,
          title: '失败主题',
          message: '请先评价新鲜度',
          iconUrl: 'https://cdn.uviewui.com/uview/demo/toast/error.png'
        })
        return
      }
      if (this.appearanceValue === 0) {
        this.$refs.goodToast.show({
          type: 'error',
          icon: false,
          title: '失败主题',
          message: '请先评价外观',
          iconUrl: 'https://cdn.uviewui.com/uview/demo/toast/error.png'
        })
        return
      }
      if (this.flavorValue === 0) {
        this.$refs.goodToast.show({
          type: 'error',
          icon: false,
          title: '失败主题',
          message: '请先评价口味',
          iconUrl: 'https://cdn.uviewui.com/uview/demo/toast/error.png'
        })
        return
      }
      let params = {
        freshness: this.freshnessValue,
        appearance: this.appearanceValue,
        taste: this.flavorValue,
        tracingCode: this.qrcode,
        appraisalReason: this.reason
      }
      let result = await submitGoodComment(params)
      if (result.code === '0') {
        this.$refs.goodToast.show({
          type: 'success',
          message: '提交成功',
          iconUrl: 'https://cdn.uviewui.com/uview/demo/toast/success.png',
          complete() {
            window.history.go(-1)
          }
        })
      }
      console.log('提交', this.imgIds)
    },
    /**
     * @description: 取消按钮
     * @author: qinjj
     */
    backBtn() {
      window.history.go(-1)
    }
  }
}
</script>

<style lang="scss" scoped>
.goodComment {
  padding: 32rpx;
  background: #fafafa;
  .comment-item {
    display: flex;
    height: 60rpx;
    line-height: 60rpx;
    margin-bottom: 30rpx;
    .comment-title {
      font-size: 30rpx;
      margin-right: 10rpx;
    }
    .comment-icon {
      margin-right: 40rpx;
    }
    .comment-text {
      font-size: 28rpx;
    }
  }
  .footer-box {
    display: flex;
    justify-content: space-around;
    margin-top: 30rpx;
    .footer-btn {
      width: 30%;
    }
  }
}
</style>
