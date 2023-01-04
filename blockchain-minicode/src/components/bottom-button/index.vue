<!--
 * @Description  : 适用于页面底部的button
 * @Autor        : Hemingway
 * @Date         : 2021-08-04 12:21:38
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-08-05 21:57:18
 * @FilePath     : \blockchain-minicode\src\components\bottom-button\index.vue
-->
<template>
  <view class="bottom-button">
    <button
      type="primary"
      class="btn"
      :disabled="isDisabled"
      @click="$emit('click')"
    >
      <slot>{{ text }}</slot>
    </button>
  </view>
</template>

<script>
export default {
  name: 'BottomButton',

  props: {
    text: {
      type: String,
      default: ''
    },

    isDisabled: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {}
  },

  computed: {},

  created() {},

  mounted() {},

  methods: {}
}
</script>

<style lang="scss" scoped>
.bottom-button {
  padding: 48rpx 40rpx 40rpx;
  // background-color: $uni-bg-color-page;

  .btn {
    height: 88rpx;
    line-height: 88rpx;
    font-size: 32rpx;
    background-color: $uni-bg-color-btn;
    border-radius: 44rpx;
  }

  .btn::after {
    display: none;
  }
}
</style>
