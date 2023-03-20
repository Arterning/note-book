<!--
 * @Description  : 投诉商品
 * @Autor        : qinjj
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-14 17:52:09
 * @FilePath     : \blockchain-platform\src\views\traceability-supervision\complaint-goods\index.vue
-->
<template>
  <div class="authorize-manage">
    <i-tab>
      <i-action @on-query="getTableData" @on-reset="onReset">
        <i-action-item label="投诉商品">
          <el-input
            clearable
            v-model="searchFields.commodityName"
            placeholder="输入企业名称搜索"
          ></el-input>
        </i-action-item>
        <i-action-item label="企业名称">
          <el-input
            clearable
            v-model="searchFields.enterpriseName"
            placeholder="输入企业名称搜索"
          ></el-input>
        </i-action-item>

        <!-- <template #action>
          <el-button type="primary" @click="onAdd">新建</el-button>
        </template> -->
      </i-action>

      <i-table border :data="tableData" :loading="loading" :max-height="450">
        <el-table-column
          label="投诉商品"
          prop="commodityName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="所属品牌"
          prop="brandName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="所属企业"
          prop="riceFactoryName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="溯源编码"
          prop="tracingCode"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="投诉人姓名"
          prop="complainant"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="联系电话"
          prop="complainantsHotline"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="投诉时间"
          prop="complaintsDate"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column label="状态">
          <template slot-scope="{ row }">
            <el-tag disable-transitions :style="classObj(row.status)">{{
              row.status === '0' ? '未处理' : '已处理'
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="处理人姓名" prop="handler"></el-table-column>
        <el-table-column label="处理时间" prop="handleDate"></el-table-column>
        <el-table-column label="操作" fixed="right">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              @click="detailBtn(row)"
              style="margin-right: 14px"
              >详情</el-button
            >
            <el-button
              type="text"
              @click="handleBtn(row)"
              style="margin-right: 14px"
              :disabled="row.status !== '0'"
              v-btn-permission:7105
              >处理</el-button
            >
          </template>
        </el-table-column>
      </i-table>

      <i-footer v-if="tableData && tableData.length > 0">
        <i-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="current"
          :page-size="size"
          :total="total"
        />
      </i-footer>
    </i-tab>
    <handle-dialog
      :visible.sync="dialogShow"
      :rowData="rowData"
      @onRefresh="refresh"
      v-if="dialogShow"
    />
  </div>
</template>

<script>
import handleDialog from './handleDialog.vue'
import tableMixin from '@/views/mixins/tableMixin'
import { getList } from '@/api/complaint-goods.js'
export default {
  name: 'blockchain-complaint-goods',
  components: { handleDialog },
  mixins: [tableMixin],
  data() {
    return {
      dialogShow: false,
      searchFields: {
        commodityName: '',
        enterpriseName: ''
      },
      tableData: [],
      rowData: {}
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    refresh() {
      this.dialogShow = false
      this.getTableData()
    },
    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        commodityName: this.searchFields.commodityName,
        enterpriseName: this.searchFields.enterpriseName
      }
      const { data } = await getList(params)

      return { tableData: data?.list || [], total: data?.total || 0 }
    },

    onResetImpl() {
      this.searchFields.commodityName = ''
      this.searchFields.enterpriseName = ''
      this.getTableData()
    },
    handleBtn(row) {
      this.dialogShow = true
      this.rowData = row
    },
    detailBtn(row) {
      this.$router.push({
        name: 'blockchain-complaint-goods-detail',
        params: { row: row }
      })
    }
  },
  computed: {
    classObj() {
      return totalStatus => {
        if (totalStatus === '0') {
          return {
            background: '#f1f1f1',
            border: '1px solid #b1b1b1',
            color: '#000'
          }
        } else {
          return {
            background: '#bef3d0',
            border: '1px solid #3dbf72',
            color: ' #00c55b'
          }
        }
      }
    }
  }
}
</script>

<style scoped lang="scss">
.authorize-manage {
  height: 100%;
  background-color: #fff;
  position: relative;
}

.el-tag {
  background-color: #fff;
  border: 1px solid #fff;
  color: #fff;
  width: 60px;
  text-align: center;
}
</style>
