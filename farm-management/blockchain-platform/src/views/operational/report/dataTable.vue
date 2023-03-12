<!--
 * @Description  : 运营数据报表-数据查询界面
 * @Autor        : qinjj
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-15 17:25:56
 * @FilePath     : \blockchain-platform\src\views\operational\report\dataTable.vue
-->
<template>
  <div class="authorize-manage">
    <i-action @on-query="getTableData" @on-reset="onReset" showAll>
      <i-action-item label="时间">
        <el-date-picker
          v-model="timeSlot"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          @change="timeChange"
          value-format="yyyy-MM-dd"
          :clearable="false"
        >
        </el-date-picker>
      </i-action-item>
      <i-action-item label="服务模式">
        <el-select
          v-model="searchFields.serviceModelId"
          clearable
          filterable
          placeholder="请选择服务模式"
        >
          <el-option
            v-for="(item, index) in serviceModeList"
            :key="index"
            :label="item.commonName"
            :value="item.commonValue"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="企业">
        <el-select
          v-model="searchFields.enterpriseId"
          clearable
          filterable
          placeholder="请选择企业"
        >
          <el-option
            v-for="(item, index) in enterpriseList"
            :key="index"
            :label="item.commonName"
            :value="item.commonValue"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="品牌名">
        <el-select
          v-model="searchFields.brandId"
          clearable
          filterable
          placeholder="请选择品牌名"
        >
          <el-option
            v-for="(item, index) in brandNameList"
            :key="index"
            :label="item.commonName"
            :value="item.commonValue"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="商品名">
        <el-select
          v-model="searchFields.commodityId"
          clearable
          filterable
          placeholder="请选择商品名"
        >
          <el-option
            v-for="(item, index) in goodsList"
            :key="index"
            :label="item.commonName"
            :value="item.commonValue"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="指标">
        <el-radio-group v-model="searchFields.status" @change="getTableData">
          <el-radio label="1" style="margin-bottom: 5px">创建溯源码数</el-radio>
          <el-radio label="2">绑定溯源码数</el-radio>
          <el-radio label="3">被扫溯源码数</el-radio>
          <el-radio label="4">扫码次数</el-radio>
          <el-radio label="5">扫码分布省份</el-radio>
        </el-radio-group>
      </i-action-item>
    </i-action>
    <!-- <el-button
      type="primary"
      icon="el-icon-download"
      style="margin-bottom: 10px"
      >导 出</el-button
    > -->

    <i-table border :data="tableData" :loading="loading">
      <el-table-column
        label="时间"
        prop="scopeDate"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="企业"
        prop="enterpriseName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="品牌名"
        prop="brandName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="商品名"
        prop="commodityName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        :label="columnValue"
        prop="value"
        show-overflow-tooltip
      ></el-table-column>
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
  </div>
</template>

<script>
import {
  serviceModelList,
  commodityList,
  enterpriseList,
  brandList,
  dimensionalityList
} from '@/api/operational'
import tableMixin from '@/views/mixins/tableMixin'
import { parseTime } from '@/utils'
export default {
  name: 'operational-dataTable',
  mixins: [tableMixin],
  data() {
    return {
      timeSlot: [],
      searchFields: {
        startDate: '',
        endDate: '',
        serviceModelId: '',
        enterpriseId: '',
        brandId: '',
        commodityId: '',
        status: '1'
      },
      serviceModeList: [],
      enterpriseList: [],
      brandNameList: [],
      goodsList: [],
      tableData: [],
      checkList: [],
      columnValue: '创建溯源码数',
      tableHeight: window.innerHeight - 430
    }
  },
  filters: {
    toThousands(val) {
      let str = val.toString()
      return str.replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
    }
  },
  watch: {
    'searchFields.status': {
      handler(newV) {
        switch (newV) {
          case '1':
            this.columnValue = '创建溯源码数（个）'
            break
          case '2':
            this.columnValue = '绑定溯源码数（个）'
            break
          case '3':
            this.columnValue = '被扫溯源码数（个）'
            break
          case '4':
            this.columnValue = '扫码次数（次）'
            break
          case '5':
            this.columnValue = '扫码分布省份（个）'
            break
        }
      },
      immediate: true
    }
  },
  created() {
    // 初始化下拉框的选项
    this.initSelectData()
  },
  mounted() {
    // 初始化时间段 一年
    let startTime = parseTime(
      new Date(new Date().getTime() - 365 * 24 * 60 * 60 * 1000),
      '{y}-{m}-{d}'
    )
    let endTime = parseTime(new Date(), '{y}-{m}-{d}')
    this.timeSlot = [startTime, endTime]
    this.searchFields.startDate = startTime
    this.searchFields.endDate = endTime
    this.getTableData()
  },
  methods: {
    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        ...this.searchFields
      }
      const data = await dimensionalityList(params)

      return { tableData: data?.list || [], total: data?.total || 0 }
    },

    onResetImpl() {
      let startTime = parseTime(
        new Date(new Date().getTime() - 365 * 24 * 60 * 60 * 1000),
        '{y}-{m}-{d}'
      )
      let endTime = parseTime(new Date(), '{y}-{m}-{d}')
      this.searchFields = {
        startDate: startTime,
        endDate: endTime,
        serviceModelId: '',
        enterpriseId: '',
        brandId: '',
        commodityId: '',
        status: '1'
      }
      this.getTableData()
    },

    // 时间选择改变时间
    timeChange(val) {
      let startTiming = new Date(val[0]).getTime()
      let endTiming = new Date(val[1]).getTime()
      if (endTiming - startTiming > 365 * 24 * 60 * 60 * 1000) {
        this.$message.error('查询时间段不能超过一年')
        return
      }
      this.searchFields.startDate = val[0]
      this.searchFields.endDate = val[1]
    },

    // 初始化下拉框下拉选项
    initSelectData() {
      serviceModelList().then(res => {
        this.serviceModeList = res.data
      })
      commodityList().then(res => {
        this.goodsList = res.data
      })
      enterpriseList().then(res => {
        this.enterpriseList = res.data
      })
      brandList().then(res => {
        this.brandNameList = res.data
      })
    }
  }
}
</script>

<style scoped lang="scss">
.authorize-manage {
  height: 100%;
  background-color: #fff;
  padding-bottom: 55px;
}
</style>
