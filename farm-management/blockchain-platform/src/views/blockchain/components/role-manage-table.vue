<!--
 * @Description  : 
 * @Autor        : Tyne
 * @Date         : 2021-07-21 16:00:10
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-07-27 15:17:47
 * @FilePath     : /agriculture-web/src/views/blockchain/components/role-manage-table.vue
-->
<template>
  <div class="role-manage-table">
    <el-table
      :data="disableEdit ? disableData : tableData"
      border
      style="width: 100%"
    >
      <!-- <el-table-column prop="0" min-width="80" label="终端"> </el-table-column> -->
      <el-table-column prop="0" min-width="80" label="一级菜单">
      </el-table-column>
      <el-table-column prop="1" min-width="80" label="二级菜单">
      </el-table-column>
      <el-table-column prop="2" min-width="80" label="子页面">
      </el-table-column>
      <el-table-column width="80" label="选择">
        <template slot-scope="{ row }">
          <el-checkbox :disabled="disableEdit" v-model="row['3']"></el-checkbox>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'RoleManageTable',
  props: {
    roleRightListDto: {
      type: Array,
      default() {
        return new Array()
      }
    },
    disableEdit: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {}
  },

  computed: {
    tableData() {
      return this.roleRightListDto
    },
    /**
     * @description: 查看详情的时候要过滤掉没选中的权限
     * @param {*}
     * @return {*}
     * @author: Tyne
     */
    disableData() {
      // 拥有的权限不一定带了父级的名称，所以处理一下
      const roleCopyList = JSON.parse(JSON.stringify(this.roleRightListDto))
      const length = roleCopyList.length
      let title0 = ''
      let title1 = ''
      let title2 = ''
      for (let index = 0; index < length; index++) {
        const element = roleCopyList[index]
        if (!title0 && !element['3']) title0 = element['0']
        if (!title1 && !element['3']) title1 = element['1']
        if (!title2 && !element['3']) title2 = element['2']
        if (title0 && !element['0'] && element['3']) {
          element['0'] = title0
          title0 = ''
        }
        if (title1 && !element['1'] && element['3']) {
          element['1'] = title1
          title1 = ''
        }
        if (title2 && !element['2'] && element['3']) {
          element['2'] = title2
          title2 = ''
        }
      }
      return roleCopyList.filter(item => {
        return item['3']
      })
      // return this.roleRightListDto.filter(item => {
      //   return item['3'] === true
      // })
    }
  },

  created() {},

  mounted() {},

  methods: {}
}
</script>

<style>
.role-manage-table {
  overflow: auto;
}
</style>
