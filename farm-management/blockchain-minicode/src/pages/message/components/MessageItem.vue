<template>
  <view class="message-item">
    <view class="icon-wrapper">
      <text class="iconfont icon-notifications_outlined"></text>
      <view v-if="!isRead" class="dot"></view>
    </view>
    <view class="desc">
      <view class="desc__header">
        <text class="title">{{ title }}</text>
        <text class="time">{{ time }}</text>
      </view>
      <text class="desc__body">{{ content }}</text>
    </view>
  </view>
</template>

<script>
export default {
  name: 'MessageItem',

  props: {
    title: {
      type: String,
      default: ''
    },
    content: {
      type: String,
      default: ''
    },
    time: {
      type: String,
      default: ''
    },
    isRead: {
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
.message-item {
  padding: 40rpx 32rpx;
  background-color: #fff;
  display: flex;
  align-items: center;
  position: relative;

  .icon-wrapper {
    width: 72rpx;
    height: 72rpx;
    background-color: $theme-color;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;

    .iconfont {
      font-size: 40rpx;
      color: #fff;
    }

    .dot {
      box-sizing: border-box;
      width: 20rpx;
      height: 20rpx;
      background-color: #fa5050;
      border: 4rpx solid #fff;
      border-radius: 50%;
      position: absolute;
      top: 0;
      right: 0;
    }
  }

  .desc {
    width: 0;
    margin-left: 24rpx;
    flex: 1;
    display: flex;
    flex-direction: column;

    .desc__header {
      color: #212121;
      display: flex;
      justify-content: space-between;
      align-items: center;

      .time {
        font-size: 24rpx;
        color: #c7c7c7;
      }
    }

    .desc__body {
      font-size: 28rpx;
      color: #8b8b8b;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
    }
  }
}

.message-item::after {
  content: '';
  border-bottom: 1rpx solid #eee;
  position: absolute;
  left: 32rpx;
  right: 32rpx;
  bottom: 0;
  transform: scaleY(0.5);
}
</style>
