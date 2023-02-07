<!--
 * @Description  : ITable（1.属性：loading、max-height、data；2.插槽：default插槽；3.事件：selection-change、sort-change、row-click）
 * @Autor        : Hemingway
 * @Date         : 2022-03-07 19:21:50
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-04-27 14:36:29
 * @FilePath     : \i-farm-platform\src\components\zl\i-table\index.vue
-->
<template>
  <el-table
    ref="table"
    border
    v-loading="loading"
    :max-height="maxHeight"
    :data="data"
    :highlight-current-row="true"
    :header-cell-style="{ backgroundColor: '#fafafa' }"
    :cell-style="{ padding: '7px 0' }"
    @selection-change="$emit('selection-change', $event)"
    @sort-change="$emit('sort-change', $event)"
    @row-click="onRowClick"
    ><slot></slot
  ></el-table>
</template>

<script>
export default {
  name: 'ITable',

  props: {
    // 表格最大高度
    maxHeight: {
      type: Number,
      default: 555
    },
    // 是否加载中
    loading: {
      type: Boolean,
      default: false
    },
    // 数据源
    data: {
      type: Array,
      default: () => []
    }
  },
  methods: {
    onRowClick(row) {
      this.$emit('row-click', row)
    }
  }
}
</script>
