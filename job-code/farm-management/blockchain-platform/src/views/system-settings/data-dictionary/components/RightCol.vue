<!--
 * @Description  : 字典项
 * @Autor        : Hemingway
 * @Date         : 2022-04-27 09:36:48
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-04-27 19:38:49
 * @FilePath     : \blockchain-platform\src\views\system-settings\data-dictionary\components\RightCol.vue
-->
<template>
  <div class="right-col">
    <div class="body">
      <i-action
        @on-query="getTableData"
        @on-reset="onReset"
        :loading="loading"
        :show-all="true"
        :show-query="false"
      >
        <template #action>
          <el-button
            type="primary"
            @click="$emit('pop', { isType: 'newDicValue' })"
            >新 建</el-button
          >
        </template>
      </i-action>
      <i-table
        :loading="loading"
        :data="tableData"
        @selection-change="onSelectionChange"
        ref="table"
      >
        <el-table-column label="名称">
          <template slot-scope="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="编码">
          <template slot-scope="scope">
            <span>{{ scope.row.code }}</span>
          </template>
        </el-table-column>

        <el-table-column label="值">
          <template slot-scope="scope">
            <span>{{ scope.row.value }}</span>
          </template>
        </el-table-column>

        <el-table-column label="是否可见">
          <template slot-scope="scope">
            <span>{{ scope.row.enable === 'Y' ? '是' : '否' }}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作">
          <template slot-scope="scope">
            <template v-if="scope.row.isSys === 'N'">
              <el-button
                @click="$emit('pop', { ...scope.row, isType: 'editDicValue' })"
                type="text"
                >编辑</el-button
              >
              <el-button @click="deleteCell(scope.row)" type="text"
                >删除</el-button
              >
            </template>
          </template>
        </el-table-column>
      </i-table>
    </div>

    <i-footer v-if="tableData.length > 0">
      <i-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="current"
        :page-size="size"
        :total="total"
      >
      </i-pagination>
    </i-footer>
  </div>
</template>

<script>
import { getDictionaryValue } from '@/api/systemSettings'
import tableMixin from '@/views/mixins/tableMixin'

export default {
  name: 'RightCol',

  mixins: [tableMixin],

  props: {
    dictionaryId: {
      type: String,
      default: ''
    }
  },

  data() {
    return {}
  },

  watch: {
    dictionaryId(newVal) {
      if (newVal) {
        this.getTableData()
      }
    }
  },

  methods: {
    /**
     * @description: 列表查询逻辑的具体实现（获取字典）
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getTableDataImpl(payload) {
      Object.assign(payload, { dictionaryId: this.dictionaryId }) // 添加额外请求参数

      try {
        const { list, total } = await getDictionaryValue(payload)
        const tableData = list || []

        return { tableData, total }
      } catch (error) {
        console.log('查询字典列表 error = ', error)
      }
    },

    deleteCell(row) {
      this.$confirm(`此操作将永久删除字典项：${row.name}，是否继续？`, '提示', {
        confirmButtonText: '确 定',
        cancelButtonText: '取 消',
        type: 'warning'
      })
        .then(() => {
          this.$emit('row-delete', { isType: 'VALUE', id: row.id })
        })
        .catch(() => {})
    }
  }
}
</script>
