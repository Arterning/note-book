<!--
 * @Description  : i-action的子组件
 * @Autor        : Hemingway
 * @Date         : 2022-03-02 15:15:51
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-16 16:59:38
 * @FilePath     : \blockchain-platform\src\components\i-action-item\index.vue
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
    width: 115px;
    line-height: 36px;
    padding: 0 12px 0 0;
    text-align: right;
    font-size: 14px;
    color: #606266;
  }
}
</style>
