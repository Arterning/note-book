<template>
  <view class="steps">
    <view class="date">
      <view class="">
        <text>{{ getDate(item[date]) }}</text>
        <slot name="dateTop"></slot>
      </view>
    </view>
    <view class="tail" :class="{ 'active-tail': activity === index }"></view>
    <view class="node" :class="{ 'active-node': activity === index }"></view>
    <view class="wrapper">
      <view class="wrapper-status">
        <text>{{ item[wrapperStatus] }}</text>
        <slot name="status"></slot>
      </view>
      <view class="wrapperTitle">
        <slot name="content"></slot>
      </view>
    </view>
  </view>
</template>

<script>
/*
	activity: 当前进度条状态
	wrapperStatus：流程状态对应字段
	wrapperTitle：详情对应字段
	index：进度条
	date：时间
	 */
export default {
  name: 'MSteps',
  props: {
    item: {
      type: Object,
      default: null
    },
    activity: {
      type: Number,
      default: 0
    },
    wrapperStatus: {
      type: String,
      default: 'status'
    },
    wrapperTitle: {
      type: String,
      default: 'content'
    },
    date: {
      type: String,
      default: 'date'
    },
    index: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {}
  },
  methods: {
    getDate(str) {
      if (!str) return ''
      const arr = str.split('-')
      return `${arr[1]}-${arr[2]} ${arr[0]}`
    }
  }
}
</script>

<style lang="scss">
.steps {
  position: relative;
  padding-bottom: 40rpx;

  .tail {
    position: absolute;
    left: 200rpx;
    height: 100%;
    border-left: 4rpx solid #e4e7ed;
  }

  &:last-child {
    .tail {
      display: none;
    }
  }

  .date {
    position: absolute;
    width: 84px;
    height: 60px;
    top: -6rpx;
    color: #aaa;
    text-align: center;
    font-size: 26rpx;
  }

  .node {
    left: 190rpx;
    width: 24rpx;
    height: 24rpx;
    position: absolute;
    background-color: #e4e7ed;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .wrapper {
    position: relative;
    width: calc(100vw - 280rpx);
    padding-left: 240rpx;
    top: -6rpx;

    .wrapper-status {
      font-family: PingFangSC-Regular;
      font-size: 28rpx;
      color: #212121;
    }
  }
}
</style>
