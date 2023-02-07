<!--
 * @Description  : 卡片式组件
 * @Autor        : Hemingway
 * @Date         : 2021-07-21 19:24:29
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-23 20:38:25
 * @FilePath     : \blockchain-minicode\src\components\card\index.vue
-->
<template>
  <view class="card">
    <view class="wrapper">
      <view v-if="!data.hiddenHeader" class="card__header">
        <slot name="title"
          ><text class="title">{{ data.title }}</text></slot
        >
        <slot>{{ data.extra }}</slot>
      </view>
      <view class="card__body">
        <view v-for="(item, index) in data.items" :key="index" class="row">
          <text>{{ item.text }}</text>
          <slot :needDecoration="item.needDecoration"> </slot>
        </view>
      </view>
      <slot name="footer"></slot>
    </view>
  </view>
</template>

<script>
export default {
  name: 'Card',

  props: {
    data: {
      type: Object,
      default: () => {}
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
.card {
  padding: 16rpx;

  .wrapper {
    background-color: #fff;
    font-size: 32rpx;
    border-radius: 16rpx;
    box-shadow: 0 0 12rpx 1rpx rgba(0, 0, 0, 0.05);

    .card__header {
      // height: 104rpx;
      padding: 28rpx 24rpx;
      color: #212121;
      display: flex;
      justify-content: space-between;
      align-items: center;
      position: relative;
    }

    .card__header::after {
      content: '';
      position: absolute;
      left: 24rpx;
      right: 24rpx;
      bottom: 0;
      border-bottom: 1rpx solid #eee;
      transform: scaleY(0.5);
    }

    .card__body {
      box-sizing: border-box;
      //   height: 168rpx;
      padding: 32rpx 24rpx;
      font-size: 28rpx;
      color: #8b8b8b;
      display: flex;
      flex-direction: column;

      .row {
        display: flex;
        justify-content: space-between;
        align-items: center;
      }

      .row:not(:first-of-type) {
        margin-top: 12rpx;
      }
    }
  }
}
</style>
