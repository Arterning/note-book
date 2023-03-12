<!--
 * @Description  : 审核结果提示组件
 * @Autor        : Hemingway
 * @Date         : 2021-07-22 15:18:31
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-07-29 12:29:17
 * @FilePath     : \blockchain-minicode\src\pages\review-progress\components\Info.vue
-->
<template>
  <view class="info">
    <view class="main">
      <view class="icon-wrapper" :style="{ backgroundColor: theme[0] }">
        <text class="iconfont icon-xiaoxi"></text>
      </view>
      <view class="desc">
        {{ desc }}
      </view>
    </view>
    <!-- 如果有备注，展现备注 -->
    <view v-if="note" class="note" :style="{ color: theme[0] }">{{
      note
    }}</view>
  </view>
</template>

<script>
import { statusMap } from '../statusMap'

export default {
  name: 'Info',

  props: {
    status: {
      type: String,
      default: ''
    },
    note: {
      type: String,
      default: ''
    }
  },

  data() {
    return {}
  },

  computed: {
    // 颜色主题
    theme() {
      if (!this.status) {
        return []
      }
      return statusMap[this.status].theme
    },
    // 描述语
    desc() {
      if (!this.status) {
        return ''
      }
      return statusMap[this.status].desc
    }
  },

  created() {},

  mounted() {},

  methods: {}
}
</script>

<style lang="scss" scoped>
.info {
  .main {
    padding: 40rpx 32rpx;
    background-color: #fff;
    display: flex;
    align-items: center;
    position: relative;

    .icon-wrapper {
      width: 72rpx;
      height: 72rpx;
      border-radius: 50%;
      display: flex;
      justify-content: center;
      align-items: center;
      position: relative;

      .iconfont {
        font-size: 40rpx;
        color: #fff;
      }
    }

    .desc {
      margin-left: 24rpx;
      font-size: 32rpx;
      color: #212121;
    }
  }

  .note {
    height: 104rpx;
    line-height: 104rpx;
    padding-left: 32rpx;
    font-size: 28rpx;
    background-color: #fff;
    position: relative;
  }

  .note::after {
    content: '';
    border-top: 1rpx solid #eee;
    transform: scaleY(0.5);
    position: absolute;
    top: 0;
    left: 32rpx;
    right: 32rpx;
  }
}
</style>
