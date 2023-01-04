<!--
 * @Description  : 
 * @Autor        : Hemingway
 * @Date         : 2021-06-21 18:40:17
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-06-22 19:16:41
 * @FilePath     : \blockchain-minicode\src\components\m-sidebar\m-sidebar.vue
-->
<template>
  <view class="sidebar">
    <slot></slot>
  </view>
</template>

<script>
export default {
  name: 'MSidebar',
  // props: {
  //   title: String
  // },
  data() {
    return {}
  }
}
</script>

<style lang="scss">
.sidebar_tab {
  display: flex;
  align-items: center;

  .sidebar_ico {
    border-width: 0rpx;
    width: 2rpx;
    height: 18rpx;
    background: inherit;
    background-color: rgba(19, 98, 253, 1);
    border: none;
    border-radius: 2rpx;
    box-shadow: none;
  }

  .sidebar_title {
    margin-left: 20rpx;
    font-weight: bold;
    font-size: 32rpx;
  }
}
</style>
