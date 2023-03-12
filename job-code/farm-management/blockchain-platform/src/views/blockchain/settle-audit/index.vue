<!--
 * @Description  : 入驻审核页
 * @Autor        : Hemingway
 * @Date         : 2021-06-21 20:43:26
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-31 18:19:09
 * @FilePath     : \blockchain-platform\src\views\blockchain\settle-audit\index.vue
-->
<template>
  <div class="settle-audit">
    <i-tab>
      <i-action @on-query="getTableData" @on-reset="onReset" :loading="loading">
        <i-action-item label="审核状态">
          <el-select
            v-model="checkStatus"
            clearable
            filterable
            placeholder="请选择审核状态"
          >
            <el-option
              v-for="(item, index) in auditStatusList"
              :key="index"
              :label="item.name"
              :value="item.id"
            ></el-option>
          </el-select>
        </i-action-item>
        <i-action-item label="企业类型">
          <el-select
            v-model="companyType"
            clearable
            filterable
            placeholder="请选择企业类型"
          >
            <el-option
              v-for="(item, index) in companyList"
              :key="index"
              :label="item.companyName"
              :value="item.companyType"
            ></el-option>
          </el-select>
        </i-action-item>
        <i-action-item label="企业名称">
          <el-input
            v-model="companyName"
            :trigger-on-focus="false"
            @keyup.enter.native="refreshData"
            placeholder="请输入企业名称"
            clearable
          >
          </el-input>
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
          <el-table-column label="流程号" show-overflow-tooltip prop="id">
          </el-table-column>
          <el-table-column label="企业类型" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>
                {{ companyTypeWithType(scope.row.type) }}
              </span>
            </template>
          </el-table-column>
          <el-table-column
            label="企业名称"
            show-overflow-tooltip
            prop="name"
            min-width="100"
          >
          </el-table-column>
          <el-table-column
            label="统一信用代号"
            prop="uuid"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            label="企业联系电话"
            show-overflow-tooltip
            prop="officePhone"
          >
          </el-table-column>
          <el-table-column
            label="企业联系地址"
            prop="address"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            label="申请日期"
            prop="createDate"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column label="审核状态">
            <template slot-scope="scope">
              <el-tag
                :type="tagTypeWithStatus(scope.row.checkStatus)"
                style="padding: 0px 5px"
                >{{ tagNameWithStatus(scope.row.checkStatus) }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column
            label="审核日期"
            prop="checkTime"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            label="审核人"
            prop="approver"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column label="操作" fixed="right">
            <template slot-scope="{ row }">
              <el-button
                v-if="row.checkStatus === '0' || row.checkStatus === '2'"
                @click="handleClick(row, 0)"
                type="text"
                >审核</el-button
              >
              <el-button
                @click="handleClick(row, 1)"
                type="text"
                style="color: #67c23a"
                >详情</el-button
              >
              <el-button
                v-if="row.checkStatus === '1'"
                @click="handleClick(row, 2)"
                type="text"
                style="color: #f56c6c"
                >取消资格</el-button
              >
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

      <settleAuditDialog
        ref="settleDialog"
        @confirmAudit="confirmAudit"
        :audit-data="auditData"
        :mapInfo="mapInfo"
      />
    </i-tab>
  </div>
</template>

<script>
import settleAuditDialog from '../components/settle-audit-dialog'
import {
  getListData,
  getInfoDict,
  getAuditDetail,
  checkEnterprise
} from '@/api/blockchain/settleAudit'
import tableMixin from '@/views/mixins/tableMixin'

export default {
  name: 'blockchain--settle-audit',

  mixins: [tableMixin],

  components: { settleAuditDialog },

  data() {
    return {
      checkStatus: '', // 搜索的状态
      companyType: '', // 搜索的企业类型
      companyName: '', // 搜索的公司名

      mapInfo: {},
      auditId: '',
      type: 0,
      auditData: {}
    }
  },

  computed: {
    dialogRef() {
      return this.$refs.settleDialog
    },
    companyList() {
      if (!this.mapInfo.enterprise_type) return []
      const keys = Object.keys(this.mapInfo.enterprise_type)
      const names = keys.map(it => ({
        companyType: it,
        companyName: this.mapInfo.enterprise_type[it]
      }))
      return names
    },
    auditStatusList() {
      if (!this.mapInfo.enterpriseCheckStatus) return []
      const keys = Object.keys(this.mapInfo.enterpriseCheckStatus)
      const names = keys.map(it => ({
        id: it,
        name: this.mapInfo.enterpriseCheckStatus[it]
      }))
      return names
    }
  },

  mounted() {
    this.getTableData()
    this.initInfo()
  },

  methods: {
    async initInfo() {
      const { code, dictMap } = await getInfoDict()
      if (code === '0') {
        this.mapInfo = dictMap
      }
    },

    /**
     * 点击操作按钮执行对应逻辑代码
     * @method
     */
    async handleClick(row, type) {
      this.type = type
      this.auditId = row.id
      if (type === 2) {
        this.cancleAudit(row.id)
      } else {
        this.auditDetail(row, type === 0)
      }
    },
    async auditDetail(row, isHandle) {
      const auditId = row.id
      const { code, enterpriseChangeDto } = await getAuditDetail({
        id: auditId
      })
      if (code === '0') {
        this.auditData = enterpriseChangeDto
        this.dialogRef.showDialog(isHandle)
        // 变为审核中
        if (isHandle && row.checkStatus === '0') {
          row.checkStatus = '2' // eslint-disable-line
          await checkEnterprise({
            id: auditId,
            checkStatus: '2'
          })
        }
      }
    },
    cancleAudit(id) {
      // 取消资格
      this.auditData = { id }
      this.dialogRef.handleAudit('4')
    },

    /**
     * @description: 重置的具体实现
     * @author: Hemingway
     */
    onResetImpl() {
      this.checkStatus = ''
      this.companyType = ''
      this.companyName = ''
    },

    /**
     * @description: 列表查询逻辑的具体实现
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getTableDataImpl(payload) {
      Object.assign(payload, {
        checkStatus: this.checkStatus,
        type: this.companyType,
        name: this.companyName
      }) // 添加额外请求参数

      try {
        const { list, total } = await getListData(payload)
        const tableData = list || []

        return { tableData, total }
      } catch (error) {
        console.log('查询入驻审核列表 error = ', error)
      }
    },

    async confirmAudit(param) {
      const { code } = await checkEnterprise(param)
      if (code === '0') {
        let toastMessage = '审核成功'
        if (param.checkStatus === '4') {
          toastMessage = '取消资格成功'
        }
        this.dialogRef.closeAuditHandle()
        this.getTableData()
        this.$message({ message: toastMessage, type: 'success' })
      }
    },
    tagTypeWithStatus(status) {
      let type = ''
      switch (status) {
        case '0':
          type = 'info'
          break
        case '1':
          type = 'success'
          break
        case '2':
          type = 'info'
          break
        case '3':
          type = 'warning'
          break
        case '4':
          type = 'danger'
          break
        default:
          break
      }
      return type
    },
    tagNameWithStatus(status) {
      let type = ''
      switch (status) {
        case '0':
          type = '未审核'
          break
        case '1':
          type = '通过'
          break
        case '2':
          type = '审核中'
          break
        case '3':
          type = '未通过'
          break
        case '4':
          type = '已取消资格'
          break
        default:
          break
      }
      return type
    },
    companyTypeWithType(type) {
      try {
        const companyMap = this.mapInfo.enterprise_type
        const nameList = type.split(',').map(i => companyMap[i])
        const typeName = nameList.join(',')
        return typeName
      } catch (error) {
        return ''
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.settle-audit {
  height: 100%;
  background-color: #fff;
  position: relative;

  .body {
    padding-bottom: 56px; // 预留 i-footer 位置
  }
}
</style>
