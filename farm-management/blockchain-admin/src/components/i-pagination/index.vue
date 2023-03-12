<!--
 * @Description  : 翻页组件
 * @Autor        : Hemingway
 * @Date         : 2022-03-07 16:03:57
 * @LastEditors  : WuJing
 * @LastEditTime : 2022-04-22 14:05:27
 * @FilePath     : \i-farm-admin\src\components\zl\i-pagination\index.vue
-->
<template>
  <el-pagination
    @size-change="$emit('size-change', $event)"
    @current-change="$emit('current-change', $event)"
    background
    :current-page="currentPage"
    :page-sizes="[10, 20, 50, 100]"
    :page-size="pageSize"
    layout="total, sizes, prev, pager, next, jumper"
    :total="total"
  >
  </el-pagination>
</template>

<script>
export default {
  name: 'IPagination',

  props: ['currentPage', 'pageSize', 'total']
}
</script>
