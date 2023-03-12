<!--
 * @Description  : i-action的子组件
 * @Autor        : Hemingway
 * @Date         : 2022-03-02 15:15:51
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-04-15 11:49:39
 * @FilePath     : \i-farm-admin\src\components\zl\i-action-item\index.vue
-->
<template>
  <div class="i-action-item">
    <label class="i-action-item__label" v-if="label">{{ label }}</label>
    <div class="i-action-item__content">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IActionItem',

  props: {
    label: {
      type: String,
      default: ''
    }
  },

  data() {
    return {}
  },

  methods: {}
}
</script>

<style lang="scss" scoped>
.i-action-item {
  width: 420px;
  height: 52px;
  padding-bottom: 16px;
  display: flex;

  &__label {
    width: 96px;
    line-height: 36px;
    padding: 0 12px 0 0;
    text-align: right;
    font-size: 14px;
    color: #606266;
  }
}
</style>
