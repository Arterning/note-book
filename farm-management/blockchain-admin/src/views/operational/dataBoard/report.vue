<template>
  <div class="report">
    <i-tab>
      <el-date-picker
        v-model="plantYear"
        type="year"
        format="yyyy 年"
        value-format="yyyy"
        placeholder="请选择种植年度"
      >
      </el-date-picker>
      <el-row type="flex">
        <el-col :span="20">
          <i-action
            @on-query="getListData"
            @on-reset="onReset"
            show-all
            class="searchs"
          >
            <i-action-item label="农场名称">
              <el-select
                v-model="searchFields.farmId"
                placeholder="请选择农场"
                filterable
              >
                <el-option
                  v-for="item in farmList"
                  :key="item.id"
                  :label="item.dicValue"
                  :value="item.dicCode"
                >
                </el-option>
              </el-select>
            </i-action-item>
            <i-action-item label="品种">
              <el-select
                v-model="searchFields.riceVarietyId"
                placeholder="请选择品种"
                filterable
              >
                <el-option
                  v-for="item in varietiesList"
                  :key="item.id"
                  :label="item.dicValue"
                  :value="item.dicCode"
                >
                </el-option>
              </el-select>
            </i-action-item>
          </i-action>
        </el-col>
        <el-col :span="4"
          ><p style="text-align: center">
            {{ `共 ${number} 批， ` }}{{ weight | toThousands }}{{ ` kg` }}
          </p></el-col
        >
      </el-row>

      <!-- 表格部分  -->
      <i-table border :data="tableData" :loading="loading" :max-height="450">
        <el-table-column
          label="农场"
          prop="farmName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="品种"
          prop="riceVarietyName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="采收日期"
          prop="reachingDate"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="农场主电话"
          prop="phone"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="运输车牌"
          prop="carNumber"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="加工厂"
          prop="enterpriseName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="加工厂电话"
          prop="enterprisePhone"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="湿稻谷(kg)"
          prop="wetRiceWeight"
          show-overflow-tooltip
        >
        </el-table-column>
        <el-table-column
          label="含水量(%)"
          prop="wetRicePrimage"
          show-overflow-tooltip
        ></el-table-column>
        <!-- <el-table-column label="操作" width="90px">
          <template slot-scope="{ row }">
            <el-button type="text" @click="clickReport(row)"
              >种植详情</el-button
            >
          </template>
        </el-table-column> -->
      </i-table>
    </i-tab>
    <!-- footer区域  -->
    <i-footer :title="title" v-if="tableData && tableData.length > 0">
      <i-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="pageNum"
        :page-size="pageSize"
        :total="total"
      />
    </i-footer>
    <!-- <dialog-report
      v-if="dialogReportShow"
      ref="refDialogReport"
      @refReportData="tableData"
    ></dialog-report> -->
  </div>
</template>
<script>
import { getGrainVarietyList, queryGetDict } from '@/api/operational-dataBoard'
import tableMixin from '@/views/mixins/tableMixin'
export default {
  name: 'report',
  mixins: [tableMixin],
  components: {},
  props: {},
  data() {
    return {
      // 页脚标题
      title: '采收明细',
      // 年份
      plantYear: '',
      // 搜索条件对象
      searchFields: {
        // 农场名称
        farmId: '',
        // 品种
        riceVarietyId: ''
      },
      // 农场下拉列表
      farmList: [],
      // 种类下拉列表
      varietiesList: [],
      // 主列表
      tableData: [],
      // 批次
      number: '',
      // 总量
      weight: '',
      // 抽屉组件显示开关
      dialogReportShow: false,
      // tab传值
      loading: false,
      // 页脚
      pageNum: 1,
      pageSize: 10,
      total: 0
    }
  },
  filters: {
    toThousands(val) {
      let _str
      let str = val.toString()
      let flag = str.indexOf('.')
      let before
      let after
      if (flag !== -1) {
        before = str.substring(0, flag).replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
        after = str.substring(flag)
        _str = before + after + ''
      } else {
        _str = str.replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
      }
      return _str
    }
  },
  computed: {},
  watch: {},
  created() {},
  activated() {},
  mounted() {
    // 路由传参
    this.plantYear = this.$route.query.year
    // this.searchFields.riceVarietyId = this.$route.query.type
    // 字典调用
    this.getDictFarm()
    this.getDictArm()
    // 主列表渲染调用
    this.getListData()
  },
  methods: {
    getListData() {
      this.getListDataImpl()
    },
    onReset() {
      this.onResetImpl()
    },
    // 列表
    async getListDataImpl() {
      let params = JSON.parse(
        JSON.stringify({
          plantYear: this.plantYear !== '' ? this.plantYear : undefined,
          varietyType: this.$route.query.type,
          farmId:
            this.searchFields.farmId !== ''
              ? this.searchFields.farmId
              : undefined,
          riceVarietyId:
            this.searchFields.riceVarietyId !== ''
              ? this.searchFields.riceVarietyId
              : undefined,
          pageNum: this.pageNum,
          pageSize: this.pageSize
        })
      )
      this.tableData = this.$options.data().tableData
      this.total = 0
      try {
        let { data } = await getGrainVarietyList(params)
        console.log('采收明细', data)
        this.number = data.number
        this.weight = data.weight
        this.tableData = data.pageInfo.list
        this.total = data.pageInfo.total
      } catch (error) {
        this.$message.error
      }
    },
    // 农场-字典
    async getDictFarm() {
      let params = {
        dicType: 'farm_type'
      }
      try {
        let { data } = await queryGetDict(params)
        console.log('data1', data)
        this.farmList = data
      } catch (error) {
        console.log('error', error)
      }
    },
    // 品种-字典
    async getDictArm() {
      let params = {
        dicType: 'variety_type'
      }
      try {
        let { data } = await queryGetDict(params)
        console.log('data2', data)
        this.varietiesList = data
      } catch (error) {
        console.log('error', error)
      }
    },
    // 重置查询条件
    onResetImpl() {
      // this.searchFields.farmId = ''
      // this.searchFields.riceVarietyId = ''
      this.searchFields = this.$options.data().searchFields
      this.getListData()
    },
    // 跳转详情抽屉
    clickReport(row) {
      this.dialogReportShow = true
      this.$nextTick(() => {
        this.$refs.refDialogReport.init(row)
      })
    },
    // 页面条数变更
    handleSizeChange(val) {
      console.log('val', val)
      this.pageSize = val
      this.getListData()
    },
    // 页面页数变更
    handleCurrentChange(val) {
      console.log('val', val)
      this.pageNum = val
      this.getListData()
    }
  }
}
</script>
<style scoped lang="scss">
.report {
  .searchs {
    margin-top: 20px;
  }
}
</style>
