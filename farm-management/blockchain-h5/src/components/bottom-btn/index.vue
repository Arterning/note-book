<!--
 * @Description  : 页面底部按钮
 * @Autor        : Hemingway
 * @Date         : 2021-11-10 15:58:24
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-11 09:21:55
 * @FilePath     : \blockchain-h5\src\components\bottom-btn\index.vue
-->
<template>
  <view class="bottom-btn">
    <view class="btn"
      ><slot name="main">{{ text }}</slot></view
    >
  </view>
</template>

<script>
export default {
  name: 'ButtomBtn',

  props: {
    text: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
.bottom-btn {
  position: fixed;
  bottom: 0;
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
</style>
