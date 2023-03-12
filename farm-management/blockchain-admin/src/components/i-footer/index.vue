<!--
 * @Description  : 页面尾部组件
 * @Autor        : Hemingway
 * @Date         : 2021-12-10 13:02:50
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-20 14:15:19
 * @FilePath     : \i-farm-admin\src\components\i-footer\index.vue
-->
<template>
  <footer class="i-footer">
    <slot name="left"></slot>
    <div class="right"><slot></slot></div>
  </footer>
</template>

<script>
export default {
  name: 'IFooter',

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
.i-footer {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 56px;
  padding-bottom: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  z-index: 9;
  border-top: 1px solid #eee;
  box-shadow: 0 -6px 6px -6px #eee;

  .right {
    flex: 1;
    text-align: right;
  }
}
</style>
