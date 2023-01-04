<!--
 * @Description  : 区块链数据查询
 * @Autor        : qinjj
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-22 10:37:53
 * @FilePath     : \blockchain-platform\src\views\blockchain-query\blockchain-index.vue
-->
<template>
  <div class="authorize-manage">
    <i-tab :label="tableLable">
      <template v-slot:0>
        <i-action :showQuery="false">
          <i-action-item>
            <el-date-picker
              v-model="searchYear"
              type="year"
              format="yyyy年"
              value-format="yyyy"
              placeholder="请选择年份"
              @change="getSlData"
            >
            </el-date-picker>
          </i-action-item>
        </i-action>
        <div class="statistics-box">
          <div class="statistics-item">
            <div class="statistics-item-text">上链数据总数 (条)</div>
            <div class="statistics-item-num">
              {{ slData.totalUploads | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">种植数据 (条)</div>
            <div class="statistics-item-num">
              {{ slData.plantCount | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">加工数据 (条)</div>
            <div class="statistics-item-num">
              {{ slData.processCount | toThousands }}
            </div>
          </div>
          <div class="statistics-item">
            <div class="statistics-item-text">质检数据 (条)</div>
            <div class="statistics-item-num">
              {{ slData.checkCount | toThousands }}
            </div>
          </div>
        </div>
        <div class="echarts-box">
          <div class="time-box">
            <el-date-picker
              v-model="startYear"
              type="year"
              format="yyyy年"
              placeholder="选择开始年份"
              style="width: 150px"
              @change="startYearChange"
              value-format="yyyy"
            >
            </el-date-picker>
            ~
            <el-date-picker
              v-model="endYear"
              type="year"
              format="yyyy年"
              placeholder="选择结束年份"
              style="width: 150px"
              @change="endYearChange"
              value-format="yyyy"
            >
            </el-date-picker>
          </div>
          <div id="chartLine" style="width: 100%; height: 420px"></div>
        </div>
      </template>
      <template v-slot:1>
        <Detail></Detail>
      </template>
    </i-tab>
  </div>
</template>

<script>
import * as echarts from 'echarts'
import Detail from './detail.vue'
import { getYearly, getRange } from '@/api/blockchain-index.js'
export default {
  name: 'blockchain-index',
  components: { Detail },
  data() {
    return {
      searchYear: '',
      tableLable: ['综合统计', '明细数据'],
      startYear: '',
      endYear: '',
      slData: {
        totalUploads: 0,
        plantCount: 0,
        processCount: 0,
        checkCount: 0
      },
      xAxisData: [],
      upChainData: [],
      plantData: [],
      matchData: [],
      qualityData: []
    }
  },
  filters: {
    toThousands(val) {
      let str = val.toString()
      return str.replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
    }
  },
  created() {
    let currentYear = new Date().getFullYear()
    this.searchYear = currentYear + '年'
    this.startYear = currentYear - 6 + '年'
    this.endYear = currentYear + '年'
    this.getSlData()
    this.getEchartData()
  },
  mounted() {
    //初始化线图
    this.lineChartFun()
  },
  methods: {
    // 获取上链数据
    async getSlData() {
      let result = await getYearly({ queryYear: this.searchYear })
      this.slData = result.data || {
        totalUploads: 0,
        plantCount: 0,
        processCount: 0,
        checkCount: 0
      }
    },
    // 获取折线图数据
    async getEchartData() {
      let xAxisData = []
      let upChainData = []
      let plantData = []
      let matchData = []
      let qualityData = []
      let result = await getRange({
        yearStart: this.startYear,
        yearEnd: this.endYear
      })
      if (result.data) {
        result.data.forEach(res => {
          xAxisData.push(res.statisticYear)
          upChainData.push(res.totalUploads)
          plantData.push(res.plantCount)
          matchData.push(res.processCount)
          qualityData.push(res.checkCount)
        })
        this.xAxisData = xAxisData
        this.upChainData = upChainData
        this.plantData = plantData
        this.matchData = matchData
        this.qualityData = qualityData
      } else {
        this.xAxisData = []
        this.upChainData = []
        this.plantData = []
        this.matchData = []
        this.qualityData = []
      }
      this.lineChartFun()
    },
    // 线图
    lineChartFun() {
      var chartDom = document.getElementById('chartLine')
      let pieOneChart = echarts.init(chartDom)
      var option
      option = {
        // title: {
        //   text: 'Stacked Line'
        // },
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          data: ['上链数据总数', '种植数据', '加工数据', '质检数据']
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        toolbox: {
          feature: {
            saveAsImage: {}
          }
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: this.xAxisData
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: '上链数据总数',
            type: 'line',
            // stack: 'Total',
            label: {
              show: true,
              position: 'top'
            },
            data: this.upChainData
          },
          {
            name: '种植数据',
            type: 'line',
            // stack: 'Total',
            label: {
              show: true,
              position: 'top'
            },
            data: this.plantData
          },
          {
            name: '加工数据',
            type: 'line',
            // stack: 'Total',
            label: {
              show: true,
              position: 'top'
            },
            data: this.matchData
          },
          {
            name: '质检数据',
            type: 'line',
            // stack: 'Total',
            label: {
              show: true,
              position: 'top'
            },
            data: this.qualityData
          }
        ]
      }
      option && pieOneChart.setOption(option)
    },
    startYearChange() {
      if (this.startYear && this.endYear) {
        if (this.startYear <= this.endYear) {
          this.getEchartData()
        } else {
          this.$message({
            showClose: true,
            message: '开始年份不能大于结束年份',
            type: 'error'
          })
        }
      } else if (!this.startYear && !this.endYear) {
        this.getEchartData()
      }
    },
    endYearChange() {
      if (this.startYear && this.endYear) {
        if (this.startYear <= this.endYear) {
          this.getEchartData()
        } else {
          this.$message({
            showClose: true,
            message: '开始年份不能大于结束年份',
            type: 'error'
          })
        }
      } else if (!this.startYear && !this.endYear) {
        this.getEchartData()
      }
    }
  },
  computed: {}
}
</script>

<style scoped lang="scss">
.authorize-manage {
  height: 100%;
  background-color: #fff;
  position: relative;

  .statistics-box {
    display: flex;
    justify-content: space-between;
    border-bottom: 1px solid #dfe4ed;
    padding-bottom: 20px;
    .statistics-item {
      width: 25%;
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

  .echarts-box {
    padding-top: 20px;

    .time-box {
      padding-bottom: 20px;
    }
  }
}
</style>
