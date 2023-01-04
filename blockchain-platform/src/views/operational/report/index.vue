<!--
 * @Description  : 运营数据报表页
 * @Autor        : qinjj
 * @Date         : 2021-06-21 20:43:26
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-15 18:53:03
 * @FilePath     : \blockchain-platform\src\views\operational\report\index.vue
-->
<template>
  <div class="settle-audit">
    <i-tab :label="tableLable">
      <template :slot="slotCharts">
        <div class="page-header">
          <el-select
            v-model="checkStatus"
            clearable
            filterable
            placeholder="请选择"
            class="dm-select"
            @change="dmChange"
          >
            <el-option
              v-for="(item, index) in auditStatusList"
              :key="index"
              :label="item.dicValue"
              :value="item.dicCode"
            ></el-option>
          </el-select>
          <!-- <i-action :showQuery="false" :loading="loading">
            <i-action-item style="width: 300px">
              <el-select
                v-model="checkStatus"
                clearable
                filterable
                placeholder="请选择大米"
              >
                <el-option
                  v-for="(item, index) in auditStatusList"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                ></el-option>
              </el-select>
            </i-action-item>
          </i-action> -->
          <div class="statistics-box">
            <div class="statistics-item">
              <div class="statistics-item-text">
                合作企业(家)
                <el-tooltip
                  class="item"
                  effect="dark"
                  content="来源于各服务模式下的企业总数，取整数。"
                  placement="top"
                >
                  <i class="el-icon-info"></i>
                </el-tooltip>
              </div>
              <div class="statistics-item-num">
                {{ hzqyNum | toThousands }}
              </div>
            </div>
            <div class="statistics-item">
              <div class="statistics-item-text">
                创建品牌(个)
                <el-tooltip
                  class="item"
                  effect="dark"
                  content="来源于各品牌商下创建品牌的个数，取整数。"
                  placement="top"
                >
                  <i class="el-icon-info"></i>
                </el-tooltip>
              </div>
              <div class="statistics-item-num">
                {{ cjppNum | toThousands }}
              </div>
            </div>
            <div class="statistics-item">
              <div class="statistics-item-text">
                创建溯源码 (个)
                <el-tooltip
                  class="item"
                  effect="dark"
                  content="来源于各品牌商创建溯源码的数量，取整数。"
                  placement="top"
                >
                  <i class="el-icon-info"></i>
                </el-tooltip>
              </div>
              <div class="statistics-item-num">
                {{ cjsymNum | toThousands }}
              </div>
            </div>
            <div class="statistics-item">
              <div class="statistics-item-text">
                绑定溯源码 (个)
                <el-tooltip
                  class="item"
                  effect="dark"
                  content="来源于各品牌商溯源批次绑定的溯源码数量。取整数。"
                  placement="top"
                >
                  <i class="el-icon-info"></i>
                </el-tooltip>
              </div>
              <div class="statistics-item-num">
                {{ bdsymNum | toThousands }}
              </div>
            </div>
            <div class="statistics-item">
              <div class="statistics-item-text">
                扫码次数 (次)
                <el-tooltip
                  class="item"
                  effect="dark"
                  content="来源于各溯源码的扫码次数的记录，取整数。"
                  placement="top"
                >
                  <i class="el-icon-info"></i>
                </el-tooltip>
              </div>
              <div class="statistics-item-num">
                {{ smcsNum | toThousands }}
              </div>
            </div>
          </div>
          <!-- <i-tab label="数据报表">
           
          </i-tab> -->
        </div>
        <div class="page-body">
          <el-row class="page-body-item">
            <!-- <i-section-header title="合作企业 30 家" /> -->
            <div class="page-body-title">合作企业 {{ hzqyNum }} 家</div>
            <div id="chartLine" style="width: 100%; height: 300px"></div>
          </el-row>
          <div class="page-body-item">
            <el-row>
              <el-col :span="16" class="page-body-title"
                >消费者已扫码{{ scanNum }}次</el-col
              >
              <el-col :span="8">
                <el-select
                  v-model="brandsChoose"
                  clearable
                  filterable
                  placeholder="请选择品牌商"
                  @change="brandChange"
                >
                  <el-option
                    v-for="(item, index) in brandsChooseList"
                    :key="index"
                    :label="item.commonName"
                    :value="item.commonValue"
                  ></el-option>
                </el-select>
              </el-col>
            </el-row>
            <div id="chartPie" style="width: 500px; height: 300px"></div>
          </div>
        </div>
      </template>
      <template :slot="slotTable">
        <div class="page-footer">
          <DataTable />
        </div>
      </template>
    </i-tab>
  </div>
</template>

<script>
import {
  cooperativeEnterpriseNumber,
  brandNumber,
  syCodeNumber,
  bondSyCodeNumber,
  scanCodeNumber,
  cooperativeEnterpriseBar,
  scanCodePie,
  enterpriseRelations,
  getDict
} from '@/api/operational'
import tableMixin from '@/views/mixins/tableMixin'
import * as echarts from 'echarts'
import DataTable from './dataTable.vue'
export default {
  name: 'blockchain--settle-audit',

  components: { DataTable },

  mixins: [tableMixin],

  data() {
    return {
      checkStatus: '1', // 搜索的状态
      companyType: '', // 搜索的企业类型
      companyName: '', // 搜索的公司名
      // 年度选择
      valueYear: '',
      // 选择品牌商
      brandsChoose: '',
      brandsChooseList: [],
      pieOneChart: null,
      mapInfo: {},
      auditId: '',
      type: 0,
      auditData: {},
      slotCharts: 0,
      slotTable: 1,
      tableLable: ['数据看板', '数据查询'],
      hzqyNum: 0,
      cjppNum: 0,
      cjsymNum: 0,
      bdsymNum: 0,
      smcsNum: 0,
      enterpriseNum: 0,
      scanNum: 0,
      auditStatusList: []
    }
  },

  filters: {
    toThousands(val) {
      let str = val.toString()
      return str.replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
    }
  },

  computed: {
    permissionBtn() {
      return this.$store.getters.permissionBtn
    }
  },
  created() {
    // 根据菜单判断权限
    let permissionBtn = this.$store.getters.permissionBtn
    if (permissionBtn.includes('1101') && !permissionBtn.includes('1102')) {
      this.tableLable = ['数据看板']
    } else if (
      !permissionBtn.includes('1101') &&
      permissionBtn.includes('1102')
    ) {
      this.tableLable = ['数据查询']
      this.slotCharts = 1
      this.slotTable = 0
    }
    getDict({ dicType: 'brand_commodity_type' }).then(res => {
      this.auditStatusList = res.data
      console.log('daomi', res)
    })
  },

  mounted() {
    this.getTableData()
    // this.initInfo()
  },

  methods: {
    // 饼图
    pieChartFun(data = []) {
      var chartDom = document.getElementById('chartPie')
      this.pieOneChart = echarts.init(chartDom)
      var option
      option = {
        title: {
          show: true,
          text: '单位：次',
          right: 0,
          bottom: '10%',
          textStyle: {
            fontWeight: 'normal',
            fontSize: 16,
            width: 100,
            align: 'right'
          }
        },
        tooltip: {
          trigger: 'item'
        },
        // legend: {
        //   top: '5%',
        //   left: 'right',
        //   // right: 'auto',
        //   orient: 'vertical'
        // },
        legend: {
          orient: 'vertical',
          left: '70%',
          top: '5%',
          itemGap: 5,
          formatter: name => {
            if (data.length > 0) {
              let value
              data.map(el => {
                if (el.name === name) {
                  value = el.value
                }
              })
              return [`{name|${name}}`, `{value|${value}}`].join('')
            } else {
              return `{name|${name}}`
            }
          },
          textStyle: {
            rich: {
              name: {
                width: -50,
                align: 'right'
              },
              value: {
                width: 95,
                align: 'right'
              }
            }
          }
        },
        series: [
          {
            // name: 'Access From',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 10,
              borderColor: '#fff',
              borderWidth: 2
            },
            label: {
              show: false,
              position: 'center'
            },
            emphasis: {
              label: {
                // show: true,
                fontSize: '40',
                fontWeight: 'bold'
              }
            },
            labelLine: {
              show: false
            },
            data: data
          }
        ]
      }
      option && this.pieOneChart.setOption(option)
    },
    // 柱状图
    lineChartFun(xaxis = [], data = []) {
      var chartDom = document.getElementById('chartLine')
      let lineOneChart = echarts.init(chartDom)
      var option
      option = {
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        grid: {
          bottom: 70
        },
        xAxis: {
          type: 'category',
          data: xaxis,
          axisLabel: {
            rotate: -20
            // hideOverlap: false
          }
        },
        yAxis: {
          type: 'value',
          name: '单位：家',
          nameLocation: 'center',
          nameTextStyle: {
            lineHeight: 56
          }
        },
        series: [
          {
            data: data,
            type: 'bar',
            showBackground: true,
            // barWidth: 60,
            backgroundStyle: {
              color: 'rgba(180, 180, 180, 0.2)'
            }
          }
        ]
      }
      option && lineOneChart.setOption(option)
    },
    async getTableDataImpl() {
      cooperativeEnterpriseNumber({ varietyType: this.checkStatus }).then(
        res => {
          this.hzqyNum = res.data.value
        }
      )
      brandNumber({ varietyType: this.checkStatus }).then(res => {
        this.cjppNum = res.data.value
      })
      syCodeNumber({ varietyType: this.checkStatus }).then(res => {
        this.cjsymNum = res.data.value
      })
      bondSyCodeNumber({ varietyType: this.checkStatus }).then(res => {
        this.bdsymNum = res.data.value
      })
      scanCodeNumber({ varietyType: this.checkStatus }).then(res => {
        this.smcsNum = res.data.value
      })
      // 合作企业柱状图
      cooperativeEnterpriseBar({ varietyType: this.checkStatus }).then(
        ({ data }) => {
          this.enterpriseNum = data.number
          this.lineChartFun(data.xaxis, data.series[0].data)
        }
      )
      // 扫码次数饼状图
      scanCodePie({ varietyType: this.checkStatus }).then(({ data }) => {
        this.scanNum = data.number
        this.pieChartFun(data.data)
      })
      // 品牌商下拉框
      enterpriseRelations().then(res => {
        this.brandsChooseList = res.data
      })
    },
    // 品牌商选择事件
    brandChange(e) {
      scanCodePie({ enterpriseId: e }).then(({ data }) => {
        this.scanNum = data.number
        this.pieChartFun(data.data)
      })
    },
    // 种类选择事件
    dmChange() {
      this.getTableDataImpl()
    }
  }
}
</script>

<style lang="scss" scoped>
.settle-audit {
  height: 100%;
  background-color: #f5f5f5;
  position: relative;
  overflow: auto;
  ::v-deep #pane-1 {
    background: white;
  }

  .body {
    padding-bottom: 56px; // 预留 i-footer 位置
  }

  .page-header {
    background-color: white;
    border-radius: 10px;

    .dm-select {
      padding-top: 20px;
      padding-left: 20px;
    }

    .statistics-box {
      display: flex;
      justify-content: space-between;
      padding: 20px;

      .statistics-item {
        width: 20%;
        text-align: center;
        border-right: 1px solid #ccc;
        display: flex;
        justify-content: space-around;
        flex-direction: column;
        align-items: center;
        height: 70px;
        .statistics-item-text {
          .el-icon-info {
            margin-left: 5px;
          }
        }
        .statistics-item-num {
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

      .page-body-title {
        border-left: 3px solid #00c853;
        padding-left: 10px;
      }
    }
  }

  .page-footer {
    background-color: white;
    border-radius: 10px;
    padding: 20px 20px 0;
    .page-footer-title {
      border-left: 3px solid #00c853;
      padding-left: 10px;
      margin-bottom: 20px;
    }
  }
}
</style>
