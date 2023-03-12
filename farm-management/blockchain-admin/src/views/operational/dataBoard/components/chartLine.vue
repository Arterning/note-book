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
    // 组件渲染数据
    echartsData: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      chart: null
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
      console.log('this.echarts-线形图配置', this.echarts)
      this.setOptions()
      // this.setOptions(this.echartsData)
    },
    // 线形图配置
    setOptions() {
      let _datas = []
      let _legendDatas = []
      if (
        this.echartsData.series &&
        Array.isArray(this.echartsData.series) &&
        this.echartsData.series.length > 0
      ) {
        _legendDatas = this.echartsData.series
        this.echartsData.series.map(el => {
          let _seriesObj = {
            name: el.name,
            type: 'line',
            // stack: 'Total',
            lineGap: 2,
            lineWidth: 10,
            emphasis: {
              focus: 'series'
            },
            data: el.data
          }
          _datas.push(_seriesObj)
        })
      }

      let lineOption = {
        // 标题
        // title: {
        //   text: 'Stacked Line'
        // },
        // 悬停展示
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            //去掉移动的指示线
            type: 'none'
          }
        },
        // 图例组件-通过点击图例控制哪些系列不显示
        legend: {
          orient: 'vertical',
          left: '81%',
          top: '30%',
          itemGap: 5,
          selectedMode: true,
          formatter: name => {
            if (_legendDatas.length > 0) {
              let value
              _legendDatas.map(el => {
                if (name === '被扫溯源码(个)') {
                  value = this.echartsData.beSweptNumber
                } else if (el.name === name) {
                  value = eval(el.data.join('+'))
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
                width: 100,
                align: 'right'
              }
            }
          }
        },
        // 直角坐标系内绘图网格组件
        grid: {
          x: 10,
          y: 100,
          x2: 300,
          y2: 10,
          borderWidth: 1,
          left: '2%',
          right: '20%',
          bottom: '2%',
          containLabel: true
        },
        // 直角坐标系 grid 中的 x 轴
        xAxis: [
          {
            type: 'category',
            boundaryGap: false,
            axisLabel: {
              show: true,
              fontSize: 12,
              //主要设置其间隔，间隔为3
              interval: 2
            },
            data: this.echartsData.xaxis
          }
        ],
        // 直角坐标系 grid 中的 y 轴
        yAxis: [
          {
            type: 'value'
          }
        ],
        series: _datas
      }
      this.chart.setOption(lineOption, true)
    }
  }
}
</script>
<style scoped lang="scss">
#line {
  width: 300px;
  height: 300px;
}
</style>
