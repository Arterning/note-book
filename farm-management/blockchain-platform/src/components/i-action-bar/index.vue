<!--
 * @Description  : 页面操作栏组件
 * @Autor        : Hemingway
 * @Date         : 2021-12-10 13:02:50
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-03-01 20:58:43
 * @FilePath     : \i-farm-platform\src\components\zl\i-action-bar\index.vue
-->
<template>
  <div class="i-action-bar">
    <slot name="left"></slot>
    <div class="right"><slot></slot></div>
  </div>
</template>

<script>
export default {
  name: 'IActionBar',

  data() {
    return {}
  },

  props: {
    title: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
.i-action-bar {
  height: 56px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .right {
    flex: 1;
    text-align: right;
  }
}
</style>
