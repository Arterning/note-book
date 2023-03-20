<!--
 * @Description  : 扫码数据
 * @Autor        : WuJing
 * @Date         : 2021-10-20 17:21:26
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-03 15:04:12
 * @FilePath     : \blockchain-screen\src\components\ScandataView.vue
-->
<template>
  <div class="scandata-view">
    <div class="time">
      <el-date-picker
        class="picker"
        v-model="time"
        type="monthrange"
        align="right"
        unlink-panels
        range-separator="到"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        valueFormat="yyyy-MM"
        :editable="false"
        @change="handle()"
        :clearable="false"
        :picker-options="pickerOptions"
      >
      </el-date-picker>
    </div>
    <v-chart class="chart" :option="option" />
  </div>
</template>

<script>
import { queryScanData } from '../api/screen'
export default {
  name: 'ScandataView',

  data() {
    return {
      time: [],
      startMonth: '',
      endMonth: '',
      chartInstance: null, //echarts实例对象
      scanNums: [], //扫码数
      scanMonth: [], //扫码月份
      option: {
        //设置标题
        title: {
          text: '用户扫码次数',
          //标题样式
          textStyle: {
            color: '#fff',
            fontSize: 16,
            fontWeight: 'normal'
          },

          //标题布局
          textAlign: 'left',
          padding: [
            0, // 上
            10, // 右
            5, // 下
            40 // 左
          ]
        },
        //坐标移动
        grid: {
          left: '5%',
          top: '20%',
          right: '8%',
          bottom: '2%',
          //坐标轴也跟着移动
          containLabel: true
        },
        tooltip: {
          //坐标轴上触发提示栏
          trigger: 'axis'
        },
        xAxis: {
          data: [],
          type: 'category',
          //x轴的开头不需要具备间隙
          boundaryGap: false,
          //不要x轴的坐标轴
          axisLine: {
            show: false
          },
          //不要x轴的刻度线
          axisTick: {
            show: false
          },
          //x轴上标签的样式设置
          axisLabel: {
            //字体
            show: true,
            //距离刻度线的距离
            margin: 20,
            color: '#fff',
            fontSize: 14
          }
        },
        yAxis: {
          type: 'value',
          //y轴标签的样式配置
          axisLabel: {
            show: true,
            fontSize: 14,
            color: '#fff'
          },
          //网格背景的设置
          splitLine: {
            show: true,
            lineStyle: {
              color: '#041D52',
              width: 1,
              //实线
              type: 'solid'
            }
          }
        },
        series: {
          data: [],
          type: 'line',
          lineStyle: {
            //渐变设置
            color: new this.$echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: '#00e676'
              }, //0%的颜色值
              {
                offset: 1,
                color: '#00e5ff'
              } //100%的颜色值
            ]),
            width: 4 //线条宽度
          },
          //折点的颜色设置
          itemStyle: {
            color: '#00e5ff'
          },
          //把折线设置为曲线
          smooth: true
        }
      },
      //限制时间选择是21年之后到当前月的
      pickerOptions: {
        disabledDate(e) {
          return (
            (e.getFullYear() === new Date().getFullYear() &&
              e.getMonth() > new Date().getMonth()) ||
            e.getFullYear() < 2021 ||
            e.getFullYear() > new Date().getFullYear()
          )
        }
      }
    }
  },
  mounted() {
    this.initTime()
    this.scanDataQuery()
  },
  methods: {
    //初始年月（近6个月）计算
    initTime() {
      let d = new Date() //定义了一个时间类型d, 来获取当前的时间
      let Year = d.getFullYear() //获取当前年
      let Month = d.getMonth() + 1 //获取当前月
      let mydate = new Date(Year, Month - 6)
      let agoYear = mydate.getFullYear() //获取6个月前的年
      let tempMonth = mydate.getMonth() + 1 //获取6个月前的月
      let agoMonth = tempMonth
      if (tempMonth < 10) {
        //处理月份为个位数的时候加个0
        agoMonth = '0' + tempMonth
      }
      this.startMonth = agoYear + '-' + agoMonth
      this.endMonth = Year + '-' + Month
      this.time = [this.startMonth, this.endMonth]
    },
    //选中事件后的回调
    handle() {
      this.startMonth = this.time[0]
      this.endMonth = this.time[1]
      this.scanDataQuery()
    },
    /**
     * @description: 查询扫码数据
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */
    async scanDataQuery() {
      try {
        const { code, list } = await queryScanData({
          //后端需要传年月日的格式，就加个-01
          queryStartMonth: this.startMonth + '-01',
          queryEndMonth: this.endMonth + '-01'
        })
        if (code === '0' && list) {
          //初始化数组
          this.scanMonth = []
          this.scanNums = []
          list.map(item => {
            //获取对象的key
            this.scanMonth.push(Object.keys(item)[0])
            //获取对象的value
            this.scanNums.push(Object.values(item)[0])
          })
          this.option.xAxis.data = this.scanMonth
          this.option.series.data = this.scanNums
        }
      } catch (error) {
        console.log('error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.scandata-view {
  color: #fff;
  height: 568px;
  width: 470px;
  background-image: url('../assets/img/scanData.png');
  background-repeat: no-repeat;

  .time {
    padding-top: 105px;
    margin-left: 20px;
    height: 48px;

    .picker {
      text-align: center;
      width: 100%;
      border: none;
      background-color: rgb(6, 20, 60);
    }
    //左右时间范围的设置
    /deep/.el-range-input {
      color: #fff;
      font-size: 18px;
      background-color: rgb(6, 20, 60);
      background-image: url('../assets/img/input.png');
      background-repeat: no-repeat;
      width: 160px;
      height: 48px;
    }
    //中间文字的设置
    /deep/.el-range-separator {
      color: #fff;
      font-size: 18px;
      margin-left: 20px;
      margin-right: 20px;
    }
    //图标的设置
    /deep/.el-input__icon {
      color: #fff;
      font-size: 18px;
      margin-right: 10px;
    }
  }

  .chart {
    margin-top: 35px;
    height: 350px;
    width: 100%;
    margin-bottom: 22px;
  }
}
</style>
