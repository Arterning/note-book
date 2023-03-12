<!--
 * @Description  : 标签组件
 * @Autor        : Hemingway
 * @Date         : 2021-07-22 10:54:22
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-07-22 11:12:23
 * @FilePath     : \blockchain-minicode\src\pages\review-progress\components\Tag.vue
-->
<template>
  <view class="tag" :style="{ color, backgroundColor }">{{ text }}</view>
</template>

<script>
export default {
  name: 'Tag',

  props: {
    text: {
      type: String,
      default: ''
    },
    color: {
      type: String,
      default: ''
    },
    backgroundColor: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
.tag {
  width: 120rpx;
  height: 48rpx;
  font-size: 24rpx;
  border-radius: 8rpx;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
