<!--
 * @Description  : i-footer的子组件
 * @Autor        : Hemingway
 * @Date         : 2022-03-07 17:07:08
 * @LastEditors  : WuJing
 * @LastEditTime : 2022-04-02 17:19:42
 * @FilePath     : \i-farm-platform\src\components\zl\i-footer-selected\index.vue
-->
<template>
  <div class="i-footer-selected">
    <el-checkbox
      :indeterminate="isIndeterminate"
      v-model="checked"
      class="checkbox"
      @change="$emit('change', $event)"
    ></el-checkbox
    ><span>{{ `已选择 ${count} 项` }}</span>
  </div>
</template>

<script>
export default {
  name: 'IFooterSelected',

  props: {
    count: {
      type: Number,
      default: 0
    },
    isIndeterminate: {
      type: Boolean,
      default: false
    },
    checkedStatus: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {
      checked: false
    }
  },
  watch: {
    checkedStatus(val) {
      console.log('监听事件', val)
      this.checked = val
    }
  }
}
</script>

<style lang="scss" scoped>
.i-footer-selected {
  margin-right: 10px;
  padding-right: 11px;
  height: 36px;
  line-height: 36px;
  position: relative;
  font-size: 14px;

  .checkbox {
    margin-right: 10px;
    width: 40px;
    text-align: center;
  }
}

.i-footer-selected::after {
  content: '';
  display: block;
  width: 1px;
  height: 20px;
  background-color: #8b8b8b;
  top: 8px;
  right: 0;
  position: absolute;
}
</style>
