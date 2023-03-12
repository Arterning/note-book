<!--
 * @Author: guoxi
 * @Date: 2021-05-25 10:29:19
 * @LastEditTime : 2021-08-05 10:03:05
 * @LastEditors  : Please set LastEditors
 * @Description: 弹窗组件
 * @FilePath     : \blockchain-minicode\src\components\Popup.vue
-->
<template>
  <view class="uni-popup-container">
    <view class="popup-title">{{ title }}</view>
    <view class="popup-content">{{ content }}</view>
    <view v-if="confirmText" class="popup-btns">
      <text class="btn" @tap="onConfirm">{{ confirmText }}</text>
    </view>
    <uni-icons
      v-if="showClose"
      type="close"
      size="34"
      color="#fff"
      class="close"
      @tap="onClose"
    ></uni-icons>
  </view>
</template>

<script>
export default {
  name: 'Popup',
  props: {
    title: {
      type: String,
      default: ''
    },
    content: {
      type: String,
      default: ''
    },
    confirmText: {
      type: String,
      default: ''
    },
    showClose: {
      type: Boolean,
      default: true
    }
  },

  data() {
    return {}
  },

  computed: {},

  created() {},

  mounted() {},

  methods: {
    onConfirm() {
      this.$emit('confirm')
    },

    onClose() {
      this.$emit('close')
    }
  }
}
</script>

<style lang="scss" scoped>
.uni-popup-container {
  width: 590rpx;
  border-radius: 8rpx;
  background-color: #fff;
  position: relative;

  .popup-title {
    height: 88rpx;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 32rpx;
    color: #212121;
  }

  .popup-content {
    padding: 28rpx 41rpx 37rpx;
    line-height: 48rpx;
    text-align: center;
    font-size: 32rpx;
    color: #8b8b8b;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .popup-btns {
    height: 88rpx;
    font-size: 32rpx;
    color: #00c853;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
  }

  .popup-btns::after {
    content: '';
    border-top: 1rpx solid #eee;
    transform: scaleY(0.5);
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
  }

  .close {
    position: absolute;
    bottom: -144rpx;
    left: 50%;
    transform: translateX(-50%);
  }
}
</style>
