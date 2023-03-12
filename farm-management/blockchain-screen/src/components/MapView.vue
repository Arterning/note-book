<!--
 * @Description  : 地图组件
 * @Autor        : Hemingway
 * @Date         : 2021-10-08 15:32:00
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-03 10:56:40
 * @FilePath     : \blockchain-screen\src\components\MapView.vue
-->
<template>
  <div class="map-view">
    <div class="title">
      <span class="count">{{ count }}</span>
    </div>
    <v-chart class="chart" autoresize :option="option" />
  </div>
</template>

<script>
import 'echarts-china-cities-js/echarts-china-cities-js/an1_hui1_wu2_hu2'
import { queryFarmMap } from '../api/screen'

export default {
  name: 'MapView',

  data() {
    return {
      count: 0,
      option: {
        geo: {
          map: '芜湖',
          roam: true,
          zoom: 1.1, // 地图比例
          itemStyle: {
            areaColor: '#14448C',
            borderColor: '#6987B5',
            shadowColor: `#337495`,
            shadowOffsetX: 1,
            shadowOffsetY: 1,
            shadowBlur: 2
          },
          emphasis: {
            itemStyle: {
              areaColor: '#14448C'
            },
            label: {
              color: '#fff' // 字体颜色
            }
          },
          label: {
            color: '#fff' // 颜色不一致会出现闪烁问题
          }
        },
        series: [
          {
            name: '设备分布',
            type: 'scatter', // 散点图
            coordinateSystem: 'geo', // 应用地图geo
            // 散点图样式设置
            itemStyle: {
              color: '#00E5FF' // 散点图颜色
            },
            data: [],
            symbol:
              'path://M512 64C317.92 64 160 221.92 160 416c0 187.36 315.424 520.032 328.832 534.08C494.88 956.448 503.264 960 512 960c0.224 0 0.48 0 0.704 0 8.992 0 17.472-4.192 23.392-10.944l109.216-125.12C790.432 646.176 864 508.928 864 416 864 221.92 706.08 64 512 64zM512 576c-88.384 0-160-71.616-160-160s71.616-160 160-160 160 71.616 160 160S600.384 576 512 576z', // 标记样式
            symbolSize: 20, // 标记的大小
            symbolKeepAspect: true // 保持长宽比
          }
        ]
      }
    }
  },

  mounted() {
    this.queryFarm()
  },

  methods: {
    /**
     * @description: 查询农场的数量和坐标
     * @author: WuJing
     */
    async queryFarm() {
      try {
        const { code, proFarmList, farmTotal } = await queryFarmMap({
          provinceCode: '340000', // 安徽省
          cityCode: '340200 ', // 芜湖市
          countyCode: '340208' // 三山区
        })
        if (code === '0' && proFarmList) {
          this.option.series[0].data = proFarmList.map(item => ({
            name: item.farmName,
            value: item.value
          }))

          this.count = farmTotal
        }
      } catch (error) {
        console.log('农场查询 error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.map-view {
  color: #fff;
  width: 900px;
  height: 568px;

  .title {
    width: 240px;
    height: 64px;
    font-size: 24px;
    line-height: 65px;
    background-image: url('../assets/img/farmCount.png');
    background-repeat: no-repeat;
    position: absolute;

    .count {
      margin-left: 145px;
    }
  }

  .chart {
    height: 568px;
    width: 900px;
  }
}
</style>
