<!--
 * @Description  : 页面头部组件
 * @Autor        : Hemingway
 * @Date         : 2021-12-10 13:02:50
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-01-06 11:35:59
 * @FilePath     : \i-farm-platform\src\components\zl\i-header\index.vue
-->
<template>
  <header class="i-header">
    <div>
      <span class="title">{{ title }}</span>
      <span class="label">{{ label }}</span>
    </div>
    <slot> </slot>
  </header>
</template>

<script>
export default {
  name: 'IHeader',

  data() {
    return {}
  },

  props: {
    title: {
      type: String,
      default: ''
    },
    label: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
.i-header {
  height: 56px;
  padding: 0 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #eee;

  .title {
    font-size: 18px;
    font-weight: bold;
    color: #212121;
  }
  .label {
    font-size: 14px;
    color: #8b8b8b;
  }
}
</style>
