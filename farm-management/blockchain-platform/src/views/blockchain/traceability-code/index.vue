<!--
 * @Description  : 溯源码管理
 * @Autor        : jiagui_chen
 * @Date         : 2021-06-21 21:09:57
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-06 11:17:15
 * @FilePath     : \blockchain-platform\src\views\blockchain\traceability-code\index.vue
-->
<template>
  <div class="traceability-code">
    <i-tab>
      <i-action @on-query="getTableData" @on-reset="onReset" :loading="loading">
        <i-action-item label="状态">
          <el-select v-model="applyStatus" clearable placeholder="请选择状态">
            <el-option
              v-for="(item, index) in shipList"
              :key="index"
              :label="item.name"
              :value="item.id"
            ></el-option>
          </el-select>
        </i-action-item>
        <i-action-item label="企业名称">
          <el-input
            placeholder="请输入企业名称"
            v-model="enterpriseName"
            @keyup.enter.native="getTableData"
            clearable
          />
        </i-action-item>
      </i-action>

      <!-- 主体内容区域 -->
      <div class="body">
        <i-table
          :loading="loading"
          :data="tableData"
          @selection-change="onSelectionChange"
          ref="table"
        >
          <el-table-column
            label="流程号"
            prop="applyCode"
            show-overflow-tooltip
          >
          </el-table-column>
          <el-table-column
            label="企业名称"
            prop="enterpriseName"
            show-overflow-tooltip
          >
          </el-table-column>
          <el-table-column
            label="申请时间"
            prop="createDate"
            show-overflow-tooltip
          >
          </el-table-column>
          <el-table-column
            label="申请数量"
            prop="applyQuantity"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            label="收件人姓名"
            prop="reciverName"
          ></el-table-column>
          <el-table-column
            label="收件人手机号"
            prop="reciverPhone"
          ></el-table-column>
          <el-table-column
            label="收货地址"
            prop="reciverAddr"
            show-overflow-tooltip
          >
          </el-table-column>
          <el-table-column label="任务" prop="applyStatus">
            <template slot-scope="{ row }">
              <div>
                {{ filterStatus(row.applyStatus) }}
              </div>
            </template>
          </el-table-column>
          <el-table-column
            label="发货日期"
            prop="abnormal_number"
          ></el-table-column>
          <el-table-column label="快递单号" prop="expressNum"></el-table-column>
          <el-table-column label="操作人" prop="creator"></el-table-column>
          <el-table-column label="操作">
            <template slot-scope="{ row }">
              <div v-if="row.applyStatus === '2'">
                <el-button @click="handleClick(row, 0)" type="text"
                  >下载</el-button
                >
                <el-button @click="handleClick(row, 1)" type="text"
                  >发货</el-button
                >
              </div>
              <div v-if="row.applyStatus === '3'">
                <el-button @click="handleClick(row, 2)" type="text"
                  >作废</el-button
                >
              </div>
            </template>
          </el-table-column>
        </i-table>
      </div>

      <!-- footer区域 -->
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
    </i-tab>

    <traceability-code-dialog
      v-if="showTraCodeDialog"
      :visible="showTraCodeDialog"
      :type="type"
      @close="onClose"
      @confirm="onConfirm"
    />
  </div>
</template>

<script>
import tableMixin from '@/views/mixins/tableMixin'
import {
  getTracingCodeList,
  sendout,
  tracingCodeCancel,
  downloadSyCode
} from '@/api/blockchain'
import TraceabilityCodeDialog from '../components/traceability-code-dialog.vue'

export default {
  name: 'blockchain--blockchain-traceability-code',

  components: { TraceabilityCodeDialog },

  mixins: [tableMixin],

  data() {
    return {
      applyStatus: '',
      enterpriseName: '',

      applyId: '',
      showTraCodeDialog: false,
      type: 0, // 0: 下载 1: 发货 2: 作废
      shipList: [
        { id: 1, name: '待支付' },
        { id: 2, name: '未发货' },
        { id: 3, name: '已发货' },
        { id: 4, name: '已签收' },
        { id: 9, name: '已作废' }
      ]
    }
  },

  mounted() {
    this.getTableData()
  },

  methods: {
    dataURLtoBlob(base64Str) {
      const bstr = atob(base64Str)
      let n = bstr.length
      const u8arr = new Uint8Array(n)
      while (n--) {
        u8arr[n] = bstr.charCodeAt(n)
      }
      return new Blob([u8arr], { type: 'txt/plain;charset=UTF-8' }) // 该类型为txt
    },

    filterStatus(status) {
      const statusObj = {
        1: '待支付',
        2: '待发货',
        3: '已发货',
        4: '已签收',
        9: '已作废'
      }
      return status ? statusObj[status] : ''
    },

    /**
     * 点击操作按钮执行对应逻辑代码
     * @method
     */
    handleClick(row, type) {
      this.type = type
      this.applyId = row.id
      this.showTraCodeDialog = type !== 0
      if (type === 0) return this.downloadCode()
    },

    onClose() {
      this.showTraCodeDialog = false
    },

    onConfirm(form) {
      if (this.type === 1) return this.handleShip(form)
      if (this.type === 2) return this.handleCancel(form)
    },

    // 发货
    handleShip(form) {
      const params = { ...form, applyId: this.applyId }
      sendout(params)
        .then(res => {
          if (res.code === '0') {
            this.onClose()
            this.$message.success('发货成功')
            this.getTableData()
          }
        })
        .catch(err => {
          console.log(err)
        })
    },

    // 作废
    handleCancel(form) {
      const params = { ...form, applyId: this.applyId }
      tracingCodeCancel(params)
        .then(res => {
          if (res.code === '0') {
            this.onClose()
            this.$message.success('作废成功')
            this.getTableData()
          }
        })
        .catch(err => {
          console.log(err)
        })
    },

    // 下载溯源码
    async downloadCode() {
      try {
        await downloadSyCode({
          applyId: this.applyId,
          fileName: `${new Date().getTime()}.xls`
        })
      } catch (error) {
        console.log('文件下载 error ', error)
      }
    },

    filterData(queryString) {
      return this.tableData
        .filter(item => {
          return (
            item.enterpriseName
              .toLowerCase()
              .indexOf(queryString.toLowerCase()) !== -1
          )
        })
        .map(item => Object.assign({}, { value: item.enterpriseName }))
    },

    /**
     * @description: 重置的具体实现
     * @author: Hemingway
     */
    onResetImpl() {
      this.enterpriseName = ''
      this.applyStatus = ''
    },

    /**
     * @description: 列表查询逻辑的具体实现
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getTableDataImpl(payload) {
      Object.assign(payload, {
        enterpriseName: this.enterpriseName,
        applyStatus: this.applyStatus
      }) // 添加额外请求参数

      try {
        const { list, total } = await getTracingCodeList(payload)
        const tableData = list || []

        return { tableData, total }
      } catch (error) {
        console.log('查询列表 error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.traceability-code {
  height: 100%;
  background-color: #fff;
  position: relative;

  .body {
    padding-bottom: 56px; // 预留 i-footer 位置
  }
}
</style>
