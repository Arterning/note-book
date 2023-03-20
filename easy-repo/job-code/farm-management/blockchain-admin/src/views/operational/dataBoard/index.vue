<template>
  <div class="dataBoard">
    <div class="page-header">
      <i-tab label="数据报表">
        <el-row type="flex">
          <el-col style="width: 300px" class="searchs">
            <el-select
              v-model="checkStatus"
              clearable
              filterable
              placeholder="请选择"
              @change="goodsType"
            >
              <el-option
                v-for="(item, index) in auditStatusList"
                :key="index"
                :label="item.dicValue"
                :value="item.dicCode"
              ></el-option>
            </el-select>
          </el-col>
          <el-col class="searchs">
            <el-date-picker
              v-model="valueYear"
              type="year"
              format="yyyy 年"
              value-format="yyyy"
              placeholder="请选择种植年度"
              :picker-options="pickerOptions0"
              @change="goodsYear"
            >
            </el-date-picker>
          </el-col>
        </el-row>
        <div class="statistics-box">
          <div class="statistics-item">
            <div class="statistics-item-text">
              {{ rawGrainTotal.title }}
              <el-tooltip
                class="item"
                effect="dark"
                content="来源于品牌商的收粮清单记录的湿稻谷重量"
                placement="top"
              >
                <i class="el-icon-info"></i>
              </el-tooltip>
            </div>
            <div class="statistics-item-num">
              {{ rawGrainTotal.numericalValue | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">
              {{ riceTotal.title }}
              <el-tooltip
                class="item"
                effect="dark"
                content="来源于包装时的包装大米重量"
                placement="top"
              >
                <i class="el-icon-info"></i>
              </el-tooltip>
            </div>
            <div class="statistics-item-num">
              {{ riceTotal.numericalValue | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">
              {{ tracingSource.title }}
              <el-tooltip
                class="item"
                effect="dark"
                content="来源于溯源批次绑定的溯源码数量"
                placement="top"
              >
                <i class="el-icon-info"></i>
              </el-tooltip>
            </div>
            <div class="statistics-item-num">
              {{ tracingSource.numericalValue | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">
              {{ numberScannedCodes.title }}
              <el-tooltip
                class="item"
                effect="dark"
                content="来源于有扫码记录的溯源码个数"
                placement="top"
              >
                <i class="el-icon-info"></i>
              </el-tooltip>
            </div>
            <div class="statistics-item-num">
              {{ numberScannedCodes.numericalValue | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">
              {{ SweepCodeNumber.title }}
              <el-tooltip
                class="item"
                effect="dark"
                content="来源于扫码次数的记录"
                placement="top"
              >
                <i class="el-icon-info"></i>
              </el-tooltip>
            </div>
            <div class="statistics-item-num">
              {{ SweepCodeNumber.numericalValue | toThousands }}
            </div>
          </div>
        </div>
      </i-tab>
    </div>
    <div>
      <div class="el-crollbar-out">
        <el-scrollbar class="page-scroll" style="height: 100%">
          <div class="page-body">
            <div class="page-body-item" type="flex">
              <div class="distribution-food">
                <div class="page-body-title-one">原粮品种分布</div>
                <div>
                  <el-button class="btnDet" type="text" @click="details"
                    >查看明细</el-button
                  >
                </div>
              </div>
              <chart-pie
                :width="pieChart1.width"
                :height="pieChart1.height"
                :chartSync="pieChart1.chartSync"
                :echartsData="pieChartDatas1"
              ></chart-pie>
            </div>

            <div class="page-body-item">
              <el-row>
                <el-col :span="16" class="page-body-title">成品大米</el-col>
              </el-row>
              <chart-pie
                :width="pieChart2.width"
                :height="pieChart2.height"
                :chartSync="pieChart2.chartSync"
                :echartsData="pieChartDatas2"
              ></chart-pie>
            </div>
          </div>
          <div class="page-body">
            <div class="page-body-item">
              <div class="page-body-title">成品及溯源码情况</div>
              <chart-bar
                :width="barChart.width"
                :height="barChart.height"
                :echartsData="barChartDatas"
              ></chart-bar>
            </div>
            <div class="page-body-item">
              <el-row>
                <el-col :span="16" class="page-body-title">扫码地区分布</el-col>
                <el-col :span="8">
                  <el-select
                    v-model="brandsChoose"
                    clearable
                    filterable
                    placeholder="商品名称"
                    @change="brandsChooseChange"
                  >
                    <el-option
                      v-for="(item, index) in goodsNameList"
                      :key="index"
                      :label="item.commodityName"
                      :value="item.id"
                    ></el-option>
                  </el-select>
                </el-col>
              </el-row>
              <chart-pie
                :width="pieChart3.width"
                :height="pieChart3.height"
                :chartSync="pieChart3.chartSync"
                :echartsData="pieChartDatas3"
              ></chart-pie>
            </div>
          </div>
          <div class="page-bodys">
            <div class="page-body-items">
              <div class="page-body-title">扫码数据查询</div>
              <el-row type="flex">
                <el-col :span="16">
                  <i-action
                    class="block"
                    @on-query="getList"
                    @on-reset="onGetList"
                  >
                    <span class="demonstration">时间：</span>
                    <el-date-picker
                      v-model="timesValue"
                      type="daterange"
                      range-separator="~"
                      start-placeholder="开始日期"
                      end-placeholder="结束日期"
                      :picker-options="pickerOptions1"
                      value-format="yyyy-MM-dd"
                      :clearable="false"
                    >
                    </el-date-picker>
                    <span class="goods">商品：</span>
                    <el-select
                      v-model="goods"
                      clearable
                      filterable
                      placeholder="商品名称"
                      @change="goodsChange"
                    >
                      <el-option
                        v-for="(item, index) in goodsNameList"
                        :key="index"
                        :label="item.commodityName"
                        :value="item.id"
                      ></el-option>
                    </el-select> </i-action
                ></el-col>
                <el-col :span="5">
                  <el-radio-group
                    v-model="radio"
                    size="mini"
                    style="margin-top: 22px"
                    @change="goodsTable"
                  >
                    <el-radio-button label="图"></el-radio-button>
                    <el-radio-button label="表"></el-radio-button>
                  </el-radio-group>
                  <!-- <el-button-group class="figure">
                    <el-button
                      size="mini"
                      type="primary"
                      plain
                      @click="goodsTable('chart')"
                      :autofocus="['chart'].includes(ifshow)"
                      >图</el-button
                    >
                    <el-button
                      size="mini"
                      type="primary"
                      plain
                      @click="goodsTable('sheet')"
                      :autofocus="['sheet'].includes(ifshow)"
                      >表</el-button
                    >
                  </el-button-group> -->
                </el-col>
                <el-col :span="3">
                  <el-button
                    type="primary"
                    class="exports"
                    @click="getExport"
                    :disabled="exportDisable"
                    :loading="exportLoading"
                    >导出</el-button
                  >
                </el-col>
              </el-row>

              <chart-line
                v-if="['chart'].includes(ifshow)"
                :width="lineChart.width"
                :height="lineChart.height"
                :echartsData="lineChartDatas"
              ></chart-line>
              <goods-list
                v-if="['sheet'].includes(ifshow)"
                ref="goodsData"
                :timesData="timesValue"
                :plantYear="valueYear"
                :varietyType="checkStatus"
                :commodityId="goods"
                :commodityName="commodityName"
                @tableDataChange="e => exportDisChange(e)"
              ></goods-list>
              <!-- @refReportData="goodsList" getScDataLine-->
            </div>
          </div>
        </el-scrollbar>
      </div>
    </div>
  </div>
</template>
<script>
import ChartPie from './components/chartPie'
import ChartBar from './components/chartBar'
import ChartLine from './components/chartLine'
import goodsList from './components/goodsList'
import tableMixin from '@/views/mixins/tableMixin'
import {
  queryGetDict, // 字典
  getBrandList, // 商品
  getGrainGross, // 原粮总量
  getFinishedProductRiceGross, // 成品大米总量
  getBondSyCodeGross, // 绑定溯源码的个数
  getBeSweptCodeGross, // 被扫码的个数
  getSweptCodeGross, //扫码的次数
  getGrainVarietyPie, // 原粮品种分布
  getFinishedProductRicePie, // 成品大米品牌商分类饼图
  getBar, // 成品及溯源码情况
  getScAreaDistributionPie, // 扫码地区分布图
  getScDataLine, // 扫码数据查询
  getExport // 导出
} from '@/api/operational-dataBoard'
export default {
  name: 'dataBoard',
  mixins: [tableMixin],
  components: { ChartPie, ChartLine, ChartBar, goodsList },
  props: {},
  data() {
    return {
      exportDisable: false,
      exportLoading: false,
      radio: '图',
      // 原粮总量
      rawGrainTotal: {
        title: '',
        numericalValue: ''
      },
      // 大米总量
      riceTotal: {
        title: '',
        numericalValue: ''
      },
      // 绑定溯源码个数
      tracingSource: {
        title: '',
        numericalValue: ''
      },
      //  被扫码的个数
      numberScannedCodes: {
        title: '',
        numericalValue: ''
      },
      // 扫码次数
      SweepCodeNumber: {
        title: '',
        numericalValue: ''
      },
      // 图表切换
      ifshow: 'chart',
      // 饼图1
      pieChart1: {
        width: '100%',
        height: '25vh',
        chartSync: 'pieChart1'
      },
      pieChartDatas1: {},
      // 饼图2
      pieChart2: {
        width: '100%',
        height: '25vh',
        chartSync: 'pieChart2'
      },
      pieChartDatas2: {},
      // 饼图3
      pieChart3: {
        width: '100%',
        height: '25vh',
        chartSync: 'pieChart3'
      },
      pieChartDatas3: {},
      // 柱状图1
      barChart: {
        width: '95%',
        height: '25vh'
      },
      barChartDatas: {},
      // 线型图1
      lineChart: {
        width: '100%',
        height: '30vh'
      },
      lineChartDatas: {},
      // 种类
      checkStatus: '1',
      auditStatusList: [],
      // 年度选择
      valueYear: '',
      // 商品名称
      brandsChoose: '',
      brandsChooseList: [],
      pieOneChart: null,
      // 日期选择器
      timesValue: [],
      pickerOptions0: {
        disabledDate: time => {
          return time.getTime() > new Date().getTime()
        }
      },

      data1Map: new Map(),
      data2Map: new Map(),
      pickerMinDate: null,
      pickerMaxDate: null,
      day31: 31 * 24 * 3600 * 1000,
      // 日期使用 maxDate,
      pickerOptions1: {
        onPick: ({ minDate }) => {
          if (minDate && this.pickerMinDate) {
            this.pickerMinDate = null
          } else if (minDate) {
            this.pickerMinDate = minDate.getTime()
          }
        },
        disabledDate: time => {
          if (this.pickerMinDate) {
            return (
              time.getTime() > this.pickerMinDate + this.day31 ||
              time.getTime() < this.pickerMinDate - this.day31 ||
              time.getTime() > Date.now()
            )
          }
          return time.getTime() > Date.now()
        }
      },
      // 商品
      goods: '',
      goodsNameList: [],
      commodityName: '',
      pageNum: 1,
      pageSize: 10
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
  created() {
    this.getTimesValue()
    this.getDictQuery()
    this.getGoods()
    this.getTotalRawGrain(this.valueYear)
    this.getSpeciesDistribution()
    this.getScAreaDistributionPie()
    this.getScDataLine()
  },
  activated() {},
  mounted() {},
  methods: {
    exportDisChange(params) {
      console.log('改变禁用状态', params)
      this.exportDisable = params
    },
    // 商品选择下拉事件
    goodsChange(e) {
      if (!e) {
        this.commodityName = ''
        return
      }
      this.commodityName =
        this.goodsNameList.filter(res => res.id.toString() === e.toString())[0]
          ?.commodityName || ''
    },
    // 扫码数据查询-时间\商品\图\表切换 查询
    goodsListQuery() {
      if (['sheet'].includes(this.ifshow)) {
        this.$nextTick(() => {
          this.$refs.goodsData.init()
        })
      } else if (['chart'].includes(this.ifshow)) {
        this.getScDataLine()
      }
    },
    // 获取扫码数据查询时间
    getTimesValue() {
      this.valueYear = this.$moment(new Date()).format('YYYY')
      this.timesValue = []
      const _now = this.$moment(new Date()).format('YYYY-MM-DD')
      const _before = this.$moment(new Date())
        .add(-30, 'days')
        .format('YYYY-MM-DD')
      this.timesValue.push(_before, _now)
      console.log('this.timesValue', this.timesValue)
    },
    // 获取导出blob
    async getExport() {
      let filesName = '扫码数据查询' + Date.now() + '.xls'
      let params = JSON.parse(
        JSON.stringify({
          plantYear: !['', undefined, null].includes(this.valueYear)
            ? this.valueYear
            : undefined,
          varietyType: !['', undefined, null].includes(this.checkStatus)
            ? this.checkStatus
            : undefined,
          commodityId: !['', undefined, null].includes(this.goods)
            ? this.goods
            : undefined,
          startDate: !['', undefined, null].includes(this.timesValue[0])
            ? this.timesValue[0]
            : undefined,
          endDate: !['', undefined, null].includes(this.timesValue[1])
            ? this.timesValue[1]
            : undefined,
          commodityName: this.commodityName,
          pageNum: this.pageNum,
          pageSize: this.pageSize
        })
      )
      try {
        this.exportLoading = true
        let res = await getExport(params)
        this.exportLoading = false
        this.exportDownLoad(res, filesName)
      } catch (error) {
        this.exportLoading = false
      }
    },
    // 下载导出
    exportDownLoad(data, filename) {
      // 接收的是blob，若接收的是文件流，需要转化一下
      if (typeof window.chrome !== 'undefined') {
        // Chrome version
        const link = document.createElement('a')
        link.href = window.URL.createObjectURL(data)
        link.download = filename
        link.click()
      } else if (typeof window.navigator.msSaveBlob !== 'undefined') {
        // IE version
        const blob = new Blob([data], { type: 'application/force-download' })
        window.navigator.msSaveBlob(blob, filename)
      } else {
        // Firefox version
        const file = new File([data], filename, {
          type: 'application/force-download'
        })
        window.open(URL.createObjectURL(file))
      }
    },
    // 饼图3商品名称改变事件
    brandsChooseChange() {
      this.getScAreaDistributionPie()
    },
    // 扫码地区分布(饼图3)
    async getScAreaDistributionPie() {
      let params = JSON.parse(
        JSON.stringify({
          plantYear: !['', undefined, null].includes(this.valueYear)
            ? this.valueYear
            : undefined,
          varietyType: !['', undefined, null].includes(this.checkStatus)
            ? this.checkStatus
            : undefined,
          commodityId: !['', undefined, null].includes(this.brandsChoose)
            ? this.brandsChoose
            : undefined
        })
      )
      this.pieChartDatas3 = {}
      try {
        let { data } = await getScAreaDistributionPie(params)
        this.pieChartDatas3 = data
        console.log('饼图3', this.pieChartDatas3)
      } catch (error) {
        this.$message.error
      }
    },
    // 扫码数据查询(折线图1)
    async getScDataLine() {
      let params = JSON.parse(
        JSON.stringify({
          plantYear: !['', undefined, null].includes(this.valueYear)
            ? this.valueYear
            : undefined,
          varietyType: !['', undefined, null].includes(this.checkStatus)
            ? this.checkStatus
            : undefined,
          commodityId: !['', undefined, null].includes(this.goods)
            ? this.goods
            : undefined,
          startDate: !['', undefined, null].includes(this.timesValue[0])
            ? this.timesValue[0]
            : undefined,
          endDate: !['', undefined, null].includes(this.timesValue[1])
            ? this.timesValue[1]
            : undefined,
          commodityName: this.commodityName
        })
      )
      this.lineChartDatas = {}
      try {
        let { data } = await getScDataLine(params)
        this.lineChartDatas = data
        if (data.series.length === 0) {
          this.exportDisChange(true)
        } else {
          this.exportDisChange(false)
        }
        this.series = data.series
        console.log('折线图1', this.lineChartDatas)
      } catch (error) {
        this.$message.error
      }
    },
    // 原粮品种分布(饼1-2,柱1)
    async getSpeciesDistribution() {
      let params = JSON.parse(
        JSON.stringify({
          plantYear: !['', undefined, null].includes(this.valueYear)
            ? this.valueYear
            : undefined,
          varietyType: !['', undefined, null].includes(this.checkStatus)
            ? this.checkStatus
            : undefined
        })
      )
      this.pieChartDatas1 = {}
      this.pieChartDatas2 = {}
      this.barChartDatas = {}
      try {
        let res1 = await getGrainVarietyPie(params)
        let res2 = await getFinishedProductRicePie(params)
        let res3 = await getBar(params)
        this.pieChartDatas1 = res1.data
        this.pieChartDatas2 = res2.data
        this.barChartDatas = res3.data
      } catch (error) {
        console.log(error)
      }
    },
    // 种类变更
    goodsType(type) {
      this.getTotalRawGrain(this.valueYear, type)
    },
    // 年份变更
    goodsYear(year) {
      this.getTotalRawGrain(year, this.checkStatus)
    },
    // 总览查询
    async getTotalRawGrain(year, type) {
      let params = JSON.parse(
        JSON.stringify({
          plantYear: !['', undefined, null].includes(year) ? year : undefined,
          varietyType: !['', undefined, null].includes(type) ? type : undefined
        })
      )
      const _defaultDatas = this.$options.data()
      this.rawGrainTotal = _defaultDatas.rawGrainTotal
      this.riceTotal = _defaultDatas.riceTotal
      this.tracingSource = _defaultDatas.tracingSource
      this.numberScannedCodes = _defaultDatas.numberScannedCodes
      this.SweepCodeNumber = _defaultDatas.SweepCodeNumber
      try {
        let res1 = await getGrainGross(params)
        let res2 = await getFinishedProductRiceGross(params)
        let res3 = await getBondSyCodeGross(params)
        let res4 = await getBeSweptCodeGross(params)
        let res5 = await getSweptCodeGross(params)
        this.rawGrainTotal = res1.data
        this.riceTotal = res2.data
        this.tracingSource = res3.data
        this.numberScannedCodes = res4.data
        this.SweepCodeNumber = res5.data
      } catch (error) {
        this.$message.error
      }
    },
    // 大米种类-字典
    async getDictQuery() {
      let params = {
        dicType: 'brand_commodity_type'
      }
      try {
        let { data } = await queryGetDict(params)
        this.auditStatusList = data
      } catch (error) {
        console.log('error', error)
      }
    },
    // 商品-接口 brandsChooseList
    async getGoods() {
      try {
        let { data } = await getBrandList()
        console.log('data9', data)
        this.goodsNameList = data
      } catch (error) {
        console.log(error)
      }
    },

    // 图/表切换
    goodsTable(model) {
      if (model === '图') {
        this.ifshow = 'chart'
      } else {
        this.ifshow = 'sheet'
      }
      this.goodsListQuery()
    },
    //  * @description: 重置（查询）的具体实现
    onGetList() {
      console.log(`重置`)
      this.goods = this.$options.data().goods
      this.commodityName = ''
      this.getTimesValue()
      this.goodsListQuery()
    },
    // 列表
    getList() {
      console.log(`查询`)
      this.goodsListQuery()
    },
    // 查看明细
    details() {
      this.$router.push({
        name: 'report',
        query: {
          year: this.valueYear !== '' ? this.valueYear : undefined,
          type: this.checkStatus !== '' ? this.checkStatus : undefined
        }
      })
    }
  }
}
</script>
<style scoped lang="scss">
.dataBoard {
  height: 100%;
  background-color: #f5f5f5;
  position: relative;
  overflow: auto;

  .body {
    padding-bottom: 56px; // 预留 i-footer 位置
  }

  .page-header {
    background-color: white;
    border-radius: 10px;

    .statistics-box {
      display: flex;
      justify-content: space-between;

      .statistics-item {
        margin-top: 20px;
        width: 20%;
        text-align: center;
        border-right: 1px solid #ccc;
        .statistics-item-text {
          .el-icon-info {
            margin-left: 5px;
          }
        }
        .statistics-item-num {
          line-height: 45px;
          font-size: 28px;
          font-weight: 700;
        }
      }

      .statistics-item:last-child {
        border-right: none;
      }
    }
  }

  .page-body {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;

    .page-body-item {
      background-color: white;
      border-radius: 10px;
      width: calc(50% - 5px);
      padding: 20px;
    }
  }
  .page-bodys {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;

    .page-body-items {
      background-color: white;
      border-radius: 10px;
      width: calc(950% - 5px);
      padding: 20px;
    }
  }
  .searchs {
    margin-top: 20px;
    margin-left: 20px;
  }
  .distribution-food {
    display: inline-block;
    display: flex;
    .page-body-title-one {
      margin: auto 0;
    }
    .btnDet {
      margin-left: 10px;
    }
  }
  #chartPie {
    width: 500px;
    height: 300px;
  }
  // ::deep .el-crollbar-out .page-scroll .el-scrollbar__wrap {
  //   overflow-y: scroll !important;
  //   overflow-x: hidden !important;
  // }
  // .el-crollbar-out {
  //   height: 73vh;
  // }
  // ::deep .el-crollbar-out .is-horizontal {
  //   display: none !important;
  // }

  // ::deep .el-crollbar-out .el-scrollbar__wrap {
  //   overflow-x: hidden !important;
  // }
  .block {
    margin-top: 20px;
    .demonstration {
      margin-top: 10px;
    }
    .goods {
      margin-left: 15px;
      margin-top: 10px;
    }
  }
  .figure {
    margin-top: 25px;
  }
  .exports {
    margin-top: 18px;
  }
  .el-button:foucs {
    background: #00963b;
  }
}
</style>
