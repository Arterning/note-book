<!--
 * @Description  : 企业入驻分布
 * @Autor        : WuJing
 * @Date         : 2021-10-20 11:23:28
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-09 09:01:50
 * @FilePath     : \blockchain-screen\src\components\EnterpriseView.vue
-->
<template>
  <div class="enterprise-view">
    <v-chart class="chart" :option="option" />
  </div>
</template>

<script>
import { queryEnterprise } from '../api/screen'
export default {
  name: 'EnterpriseView',
  data() {
    return {
      chartInstance: null, //echarts实例对象
      enterpriseData: [],
      option: {
        //坐标轴的偏移量
        grid: {
          top: '10%',
          left: '-8%',
          right: '9%',
          bottom: '-5%',
          //坐标轴文字也跟着移动
          containLabel: true
        },
        //数值
        xAxis: {
          show: false, //不显示x轴的信息
          type: 'value'
        },
        //类目
        yAxis: {
          type: 'category',
          axisLine: {
            show: false //y轴线消失
          },
          axisTick: {
            show: false //y轴坐标点消失
          },
          data: ['品牌商数', '米厂数量', '农场数量', '企业总数'], //从下到上的顺序
          //调整类目项的位置
          axisLabel: {
            //文字样式
            color: '#fff',
            fontSize: 16,
            //把文字调整到了上方
            verticalAlign: 'bottom',
            align: 'right',
            //调整文字上右下左
            padding: [0, -80, 15, 0]
          }
        },
        series: [
          {
            type: 'bar',
            //柱状图的宽度
            barWidth: 10,
            //显示柱状图上的数据
            label: {
              show: true,
              //数据排列在右边
              position: 'right',
              fontSize: 16,
              fontWeight: 'bold',
              color: 'inherit' //自动匹配对应柱状条形的颜色
            },
            //显示背景色
            showBackground: true,
            //给背景配置样式
            backgroundStyle: {
              borderRadius: 5, // 统一设置四个角的圆角大小
              borderWidth: 10, //背景的宽度
              //背景的颜色
              color: 'rgba(110,193,244,0)',
              //边框的颜色
              borderColor: 'rgba(110, 193, 244, 0.2)'
            },
            //x轴上的数据
            data: [0, 0, 0, 0], //从左往右顺序
            //柱状条形的样式
            itemStyle: {
              //给不同条形配置对应的颜色
              color: function (params) {
                var colorList = ['#ff1744', '#00e5ff', '#0066ff', '#00e676']
                return colorList[params.dataIndex]
              },
              borderRadius: [5, 0, 0, 5] //设置柱状角度
            }
          }
        ]
      }
    }
  },
  mounted() {
    this.enterpriseQuery()
  },
  methods: {
    /**
     * @description: 查询入驻企业分布
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */
    async enterpriseQuery() {
      try {
        const { code, data } = await queryEnterprise()
        if (code === '0') {
          //把对象组成数组

          this.enterpriseData = [
            data.enterpriseNum,
            data.riceFactoryNum,
            data.farmNum,
            data.totalNum
          ]
          this.option.series[0].data = this.enterpriseData
        }
      } catch (error) {
        console.log('error = ', error)
      }
    }
  }
}
</script>
<style lang="scss" scoped>
.enterprise-view {
  width: 470px;
  height: 373px;
  background-image: url('../assets/img/enterprise_position.png');
  background-repeat: no-repeat;

  .chart {
    padding-top: 65px;
    height: 308px;
    width: 470px;
  }
}
</style>
