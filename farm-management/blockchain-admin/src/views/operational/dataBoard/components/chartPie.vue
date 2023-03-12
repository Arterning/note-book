<template>
  <div
    :ref="echarts"
    :id="echarts"
    :class="echarts"
    :style="{ width: width, height: height }"
  ></div>
</template>
<script>
import * as echarts from 'echarts'
export default {
  name: '',
  components: {},
  props: {
    // 画布宽
    width: {
      type: String,
      default: '50%'
    },
    // 画布高
    height: {
      type: String,
      default: '100px'
    },
    // 图表标识
    chartSync: {
      type: String,
      default: ''
    },
    // 组件渲染数据
    echartsData: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      chart: null
      // sidebarElm: null
    }
  },
  computed: {
    echarts() {
      // return `echarts${Math.round(Math.random() * 100000)}`
      return `echarts${Math.random() * 100000}`
    }
  },
  watch: {
    echartsData: {
      handler(data) {
        if (data) {
          if (this.chart) {
            this.$nextTick(() => {
              this.setOptions(data)
            })
          }
        }
      },
      deep: true
      // immediate: true
    }
  },
  created() {},
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
    // 监听resize
    window.addEventListener('resize', () => {
      this.chart.resize()
    })
  },
  updated() {},
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    // 释放图表实例,dispose 会释放内部占用的一些资源和事件绑定
    this.chart.dispose()
    // 解除实例的引用做不到，所以重新赋值为null
    this.chart = null
    // 取消监听resize
    window.removeEventListener('resize', () => {
      this.chart.resize()
    })
  },
  methods: {
    // 初始化实例
    initChart() {
      this.chart = echarts.init(document.getElementById(this.echarts))
      console.log('this.echarts-饼图配置', this.echarts)
      this.setOptions()
      // this.setOptions(this.echartsData)
    },
    // 饼图配置
    setOptions() {
      let _datas = []
      let _legendDatas = []
      let _title = {}
      if (
        this.echartsData.data &&
        Array.isArray(this.echartsData.data) &&
        this.echartsData.data.length > 0
      ) {
        _datas = this.echartsData.data
        _legendDatas = this.echartsData.data
        console.log('饼图数据', _legendDatas)
      }
      if (['pieChart1'].includes(this.chartSync)) {
        _title = {
          text: '单位：kg',
          right: '8%',
          legend: {
            left: '70%'
          }
        }
      } else if (['pieChart2'].includes(this.chartSync)) {
        _title = {
          text: '单位：kg',
          right: '8%',
          legend: {
            left: '60%'
          }
        }
      } else if (['pieChart3'].includes(this.chartSync)) {
        _title = {
          text: '单位：次',
          right: '10%',
          legend: {
            left: '70%'
          }
        }
      }
      let pieOption = {
        // 标题
        title: {
          show: true,
          text: _title.text,
          right: _title.right,
          bottom: '10%',
          textStyle: {
            fontWeight: 'normal',
            fontSize: 16,
            width: 100,
            align: 'right'
          }
        },
        // 悬停展示
        tooltip: {
          trigger: 'item'
        },
        // 图例组件-通过点击图例控制哪些系列不显示
        legend: {
          orient: 'vertical',
          left: _title.legend.left || '70%',
          top: '20%',
          itemGap: 5,
          selectedMode: true,
          formatter: name => {
            if (_legendDatas.length > 0) {
              let value
              _legendDatas.map(el => {
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
                width: 120,
                align: 'right'
              }
            }
          }
        },
        series: [
          {
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 10,
              borderColor: '#fff',
              borderWidth: 2
            },
            center: ['45%', '50%'],
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
            data: _datas
          }
        ]
      }
      this.chart.setOption(pieOption, true)
    }
  }
}
</script>
<style scoped lang="scss">
#pie {
  width: 600px;
  height: 300px;
}
</style>
