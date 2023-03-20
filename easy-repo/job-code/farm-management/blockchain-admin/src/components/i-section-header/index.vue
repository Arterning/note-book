<!--
 * @Description  : 区块标题组件
 * @Autor        : Hemingway
 * @Date         : 2022-03-08 19:06:18
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-03-11 13:46:53
 * @FilePath     : \i-farm-admin\src\components\zl\i-section-header\index.vue
-->
<template>
  <div class="i-section-header">
    <div class="left-side">
      <!-- title -->
      <span class="title">
        <slot>{{ title }}</slot></span
      >

      <!-- 备注说明 -->
      <span v-if="remark" class="remark">{{ `（ ${remark} ）` }}</span>
      <slot v-else name="remark"></slot>
    </div>

    <div class="right-side">
      <!-- 操作按钮区域 -->
      <slot name="action"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ISectionHeader',

  props: {
    title: {
      type: String,
      default: ''
    },
    remark: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
.i-section-header {
  padding: 10px 0;
  display: flex;
  justify-content: space-between;

  .left-side {
    position: relative;

    .title {
      height: 36px;
      line-height: 36px;
      padding-left: 14px;
      font-size: 16px;
    }

    .remark {
      font-size: 12px;
      color: #8b8b8b;
    }
  }

  .left-side::after {
    content: '';
    display: block;
    width: 4px;
    height: 16px;
    background-color: #00c853;
    top: 10px;
    left: 0;
    position: absolute;
  }
}
</style>
