<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-20 10:50:57
 * @LastEditors  : SunXiuqiong
 * @LastEditTime : 2022-07-20 16:55:26
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\authorize-manage\add.vue
-->
<template>
  <div class="add">
    <i-tab> <add-edit :rows="rows"></add-edit></i-tab>
  </div>
</template>

<script>
import AddEdit from './components/AddEdit.vue'

export default {
  name: 'add-tenement',
  components: {
    AddEdit
  },
  data() {
    return {
      rows: undefined
    }
  },
  mounted() {
    this.rows = {
      model: 'add',
      row: undefined
    }
    console.log(`AddEdit`, this.rows.model)
  }
}
</script>

<style scoped>
.add {
  height: 100%;
  background-color: #fff;
  position: relative;
}
</style>
