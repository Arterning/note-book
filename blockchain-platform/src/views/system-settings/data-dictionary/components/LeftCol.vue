<!--
 * @Description  : 字典
 * @Autor        : Hemingway
 * @Date         : 2022-04-27 09:36:24
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-04-27 19:31:08
 * @FilePath     : \blockchain-platform\src\views\system-settings\data-dictionary\components\LeftCol.vue
-->
<template>
  <div class="left-col">
    <div class="body">
      <i-action @on-query="getTableData" @on-reset="onReset" :loading="loading">
        <i-action-item label="类型">
          <el-select
            v-model="type"
            placeholder="请选择类型"
            clearable
            filterable
          >
            <el-option
              v-for="item in dictTypeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            />
          </el-select>
        </i-action-item>
        <i-action-item label="关键词">
          <el-input
            placeholder="请输入名称/编码"
            v-model="searchKey"
            @keyup.enter.native="getListData"
            clearable
          >
          </el-input>
        </i-action-item>

        <template #action>
          <el-button type="primary" @click="$emit('pop', { isType: 'newDic' })"
            >新 建</el-button
          >
        </template>
      </i-action>

      <i-table
        :loading="loading"
        :data="tableData"
        @selection-change="onSelectionChange"
        @row-click="$emit('row-click', $event)"
        ref="table"
      >
        <el-table-column label="类型">
          <template slot-scope="scope">
            <span>{{ typeFilter(scope.row.type) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="名称" prop="name" />
        <el-table-column label="编码" prop="code" />
        <el-table-column label="是否可见">
          <template slot-scope="scope">
            <span>{{ scope.row.enable === 'Y' ? '是' : '否' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <template v-if="scope.row.isSys === 'N'">
              <el-button
                @click="$emit('pop', { ...scope.row, isType: 'editDic' })"
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
import { getDataDictionary } from '@/api/systemSettings'
import tableMixin from '@/views/mixins/tableMixin'

export default {
  name: 'LeftCol',

  mixins: [tableMixin],

  data() {
    return {
      // left：筛选条件
      type: '', // 类型筛选
      searchKey: '' // 关键词
    }
  },

  props: {
    dictTypeList: {
      type: Array,
      default: () => []
    }
  },

  computed: {
    dictType() {
      return this.$store.getters.systemDic.dictType
    }
  },

  created() {},

  mounted() {
    this.getTableData()
  },

  methods: {
    /**
     * @description: 重置的具体实现
     * @author: Hemingway
     */
    onResetImpl() {
      this.type = ''
      this.searchKey = ''
    },

    /**
     * @description: 列表查询逻辑的具体实现（获取字典）
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getTableDataImpl(payload) {
      Object.assign(payload, { searchKey: this.searchKey, type: this.type }) // 添加额外请求参数

      try {
        const { list, total } = await getDataDictionary(payload)
        const tableData = list || []

        return { tableData, total }
      } catch (error) {
        console.log('查询字典列表 error = ', error)
      }
    },

    deleteCell(row) {
      this.$confirm(`此操作将永久删除字典：${row.name}，是否继续？`, '提示', {
        confirmButtonText: '确 定',
        cancelButtonText: '取 消',
        type: 'warning'
      })
        .then(() => {
          this.$emit('row-delete', { isType: 'DIC', id: row.id })
        })
        .catch(() => {})
    },

    typeFilter(id) {
      return this.dictType[id] || ''
    }
  }
}
</script>
