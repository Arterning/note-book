<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-20 10:51:36
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-09 10:46:57
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\authorize-manage\edit.vue
-->
<template>
  <div class="edit">
    <i-tab> <add-edit :rows="rows"></add-edit></i-tab>
  </div>
</template>

<script>
import AddEdit from './components/AddEdit.vue'

export default {
  name: 'edit-tenement',
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
      model: 'edit',
      row: this.$route.params.row
    }
    console.log(`AddEdit`, this.rows.model)
  }
}
</script>

<style scoped>
.edit {
  height: 100%;
  background-color: #fff;
  position: relative;
}
</style>
