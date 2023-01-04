<!--
 * @Description  : 质检报告上传
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-01 14:13:30
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-13 14:10:36
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\report-upload\index.vue
-->
<template>
  <div class="report-upload">
    <i-tab>
      <i-action show-all @on-query="getTableData" @on-reset="onReset">
        <i-action-item label="农场名称">
          <el-select
            v-model="searchFields.farmId"
            placeholder="输入农场名称搜索"
            filterable
            clearable
          >
            <el-option
              v-for="item in farmNameList"
              :label="item.name"
              :value="item.id"
              :key="item.id"
            ></el-option>
          </el-select>
        </i-action-item>
        <i-action-item label="品种">
          <el-select
            v-model="searchFields.varietyId"
            placeholder="输入品种名称搜索"
            filterable
            clearable
          >
            <el-option
              v-for="item in varirtyNameList"
              :label="item.dicValue"
              :value="item.dicCode"
              :key="item.id"
            ></el-option>
          </el-select>
        </i-action-item>

        <i-action-item label="品类名称">
          <el-select
            v-model="searchFields.type"
            placeholder="请选择品类名称"
            filterable
            clearable
          >
            <el-option
              v-for="item in typeList"
              :label="item.dicValue"
              :value="item.dicCode"
              :key="item.id"
            ></el-option>
          </el-select>
        </i-action-item>
      </i-action>

      <i-table :data="tableData" border :loading="loading">
        <el-table-column
          prop="farmName"
          label="农场"
          fixed="left"
        ></el-table-column>
        <el-table-column prop="type" label="品类">
          <template slot-scope="{ row }">
            <span>{{ row.type | typeFilter(typeList) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="varietyName" label="品种"></el-table-column>
        <el-table-column prop="plantYear" label="种植年份"></el-table-column>
        <el-table-column
          prop="farmAdministrator"
          label="农场管理员"
        ></el-table-column>
        <el-table-column prop="farmPhone" label="联系电话"></el-table-column>
        <el-table-column
          prop="farmLocation"
          label="农场地址"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column prop="qualityStatus" label="状态">
          <template slot-scope="{ row }">
            <el-tag disable-transitions :style="classObj(row.qualityStatus)">{{
              row.qualityStatus === '0' ? '待上传' : '已上传'
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="createDate"
          label="报告上传时间"
        ></el-table-column>
        <el-table-column prop="creator" label="上传人"></el-table-column>
        <el-table-column label="操作" fixed="right">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              size="small"
              @click="upload(row)"
              v-if="row.qualityStatus === '0'"
              v-btn-permission:1708
              >上传报告</el-button
            >
            <div v-else style="display: flex">
              <el-button
                type="text"
                style="color: #86a617"
                size="small"
                @click="clickCheckReport(row)"
                >查看报告</el-button
              >
              <el-tooltip
                class="item"
                effect="dark"
                content="重新上传将清空原有报告"
                placement="top"
              >
                <el-button
                  type="text"
                  style="color: red"
                  size="small"
                  @click="upload(row)"
                  v-btn-permission:1708
                  >重新上传</el-button
                >
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
      </i-table>

      <check-report
        v-if="isShowCheckDialog"
        :visible.sync="isShowCheckDialog"
        :tableData="tableData"
        :rowDataCheck="rowDataCheck"
      />

      <upload-report
        v-if="isShowUploadDialog"
        :visible.sync="isShowUploadDialog"
        :rowData="rowData"
        @refresh="getTableData()"
      />

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
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import tableMixin from '@/views/mixins/tableMixin'
import CheckReport from './components/CheckReport.vue'
import UploadReport from './components/UploadReport.vue'
import { getList, getType, getFarmSelect } from '@/api/basedata-manage'

export default {
  name: 'report-upload',
  mixins: [tableMixin],
  components: {
    CheckReport,
    UploadReport
  },
  data() {
    return {
      loading: false,
      searchFields: {
        farmId: '',
        varietyId: '',
        type: ''
      },
      farmNameList: [],
      varirtyNameList: [],
      typeList: [],
      tableData: [],
      rowData: {},
      rowDataCheck: {},
      isShowCheckDialog: false,
      isShowUploadDialog: false
    }
  },
  created() {
    this.getTableData()
    this.getFarmNameList()
    this.getVarirtyNameList()
    this.getTypeList()
  },

  filters: {
    typeFilter: (val, typeList) => {
      let result = typeList.filter(res => {
        return res.dicCode.toString() === val.toString()
      })

      if (result.length === 0) {
        return '稻谷'
      }
      return result[0].dicValue
    }
  },

  methods: {
    // 农场
    async getFarmNameList() {
      const { data } = await getFarmSelect({ dicType: 'farm_type' })
      this.farmNameList = data
    },

    //品种
    async getVarirtyNameList() {
      const { data } = await getType({ dicType: 'variety_type' })
      this.varirtyNameList = data
    },

    //品类
    async getTypeList() {
      // const { data } = await getType({ dicType: 'brand_commodity_type' })
      const { data } = await getType({ dicType: 'brand_commodity_type' })
      this.typeList = data
    },

    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        farmId: this.searchFields.farmId,
        varietyId: this.searchFields.varietyId,
        type: this.searchFields.type,
        enterpriseId: this.computedUserInfo.enterpriseId
      }
      console.log('params', params)
      const data = await getList(params)

      return { tableData: data.list, total: data.total }
    },

    onResetImpl() {
      this.searchFields.farmId = ''
      this.searchFields.varietyId = ''
      this.searchFields.type = ''
      this.getTableData()
    },

    clickCheckReport(row) {
      console.log(12)
      this.isShowCheckDialog = true
      this.rowDataCheck = row
    },
    upload(row) {
      this.rowData = row
      this.isShowUploadDialog = true
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
            background: '#bff0d6',
            border: '1px solid #46b27b',
            color: '#46b27b'
          }
        }
      }
    },
    ...mapGetters('user', { computedUserInfo: 'userInfoGetter' })
  }
}
</script>

<style scoped lang="scss">
.report-upload {
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
  border-radius: 0;
}
</style>
