<!--
 * @Description  : 步骤组件
 * @Autor        : Hemingway
 * @Date         : 2021-11-11 11:03:52
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-06-29 14:13:47
 * @FilePath     : \blockchain-h5\src\components\steps\index.vue
-->
<template>
  <view class="steps">
    <view v-for="(step, index) in steps" :key="step.title" class="step">
      <view class="date"
        ><slot v-if="current >= index" name="date" :date="step.date">{{
          step.date
        }}</slot>
      </view>

      <view class="scale">
        <view
          v-if="index !== steps.length - 1"
          class="line"
          :style="{
            backgroundColor: current - 1 < index ? '#c7c7c7' : stepsColor
          }"
        ></view>
        <view class="node">
          <view v-if="current < index" class="dot"> </view>
          <u-icon
            name="checkmark-circle-fill"
            size="16"
            :color="stepsColor"
            class="checked"
          ></u-icon>
        </view>
      </view>

      <view class="header">
        <view class="">
          <text :style="{ color: current < index ? '#8b8b8b' : stepsColor }">{{
            step.title
          }}</text>
          <!-- 新增天气插槽 -->
          <slot v-if="current >= index" name="scene" :scene="step.scene"></slot>
        </view>

        <slot v-if="current >= index" name="extra"
          ><span
            class="extra"
            :style="{ color: current < index ? '#8b8b8b' : '#212121' }"
            >{{ step.extra }}</span
          ></slot
        >
      </view>

      <view v-if="current >= index" class="body">
        <slot name="body" :body="step.body"> </slot>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'Steps',

  props: {
    // 节点数据源
    steps: {
      type: Array,
      default: () => []
    },
    // 已完成节点索引
    current: {
      type: Number,
      default: 0
    }
  },
  computed: {
    stepsColor() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return '#00c853'
      } else if (colorId === '3') {
        return '#2196f3'
      } else {
        return '#eabc02'
      }
    }
  }
}
</script>

<style lang="scss">
.steps {
  padding-left: 176rpx;
  font-size: 28rpx;
  color: #00c853;

  .step {
    position: relative;
    padding-bottom: 60rpx;

    .date {
      position: absolute;
      top: 0;
      left: -176rpx;
      color: #212121;
    }

    .scale {
      position: absolute;
      top: 24rpx;
      left: -48rpx;
      height: 100%;

      .line {
        width: 2rpx;
        height: 100%;
        background-color: #00c853;
      }

      .node {
        position: absolute;
        width: 32rpx;
        height: 32rpx;
        left: -16rpx;
        top: -16rpx;
        border-radius: 50%;
        background-color: #fff;
        display: flex;
        justify-content: center;
        align-items: center;

        .dot {
          width: 16rpx;
          height: 16rpx;
          border-radius: 50%;
          background-color: #c7c7c7;
        }
      }
    }

    .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      font-size: 32rpx;

      .extra {
        font-size: 24rpx;
      }
    }
  }

  .step:last-of-type {
    padding-bottom: 40rpx;
  }
}
</style>
