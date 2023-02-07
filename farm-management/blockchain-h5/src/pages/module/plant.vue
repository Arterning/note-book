<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 10:08:53
 * @FilePath     : \blockchain-h5\src\pages\module\plant.vue
-->
<template>
  <view>
    <!-- 种植信息 -->
    <view class="card-info common-box" :class="themeLinear">
      <view class="card-title" :class="themeColor">
        <text>种植信息</text>
        <view class="certified" :class="themeBackground">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 61.26 64"
            class="status-img"
          >
            <g id="图层_2" data-name="图层 2">
              <g id="图层_1-2" data-name="图层 1">
                <path
                  id="认证"
                  d="M61.23,5.66a2.54,2.54,0,0,0-.89-1.58c-.1-.1-1.68-1.29-4.15-3.17A2.54,2.54,0,0,0,53.82.62C40.27,6.45,32.46.82,32.06.52a2.25,2.25,0,0,0-2.87,0c-.09.1-8,5.93-21.75.1A2.54,2.54,0,0,0,5.07.91C2.6,2.79,1,4,.92,4.08A2.49,2.49,0,0,0,0,5.66,2.33,2.33,0,0,0,.52,7.44c.1.1,5.14,6.63,2.77,16.51C1.9,29.79.42,35.82,1.71,41.46,3.19,48.18,8.23,52.93,16.93,56H17A25.85,25.85,0,0,1,28.7,63.11a2.43,2.43,0,0,0,3.76,0A25,25,0,0,1,44.13,56h.09C53,52.93,58,48.18,59.45,41.46c1.29-5.74-.2-11.77-1.58-17.51A20.31,20.31,0,0,1,60.64,7.44,2.09,2.09,0,0,0,61.23,5.66Zm-43,6.82c1.38,1.88,2.77,4.06,4.25,6.43-1.09.59-2.47,1.39-4.25,2.47-1.29-2.27-2.57-4.45-3.86-6.52ZM15.15,41.26A3.83,3.83,0,0,0,16,38.69V27.51H13.28V23h7.51V36.81l3.07-2.47c.09,1,.39,2.57.79,4.74-2.18,1.58-4.65,3.46-7.42,5.84Zm32.54,4.05H23.06V40.66h2.87V23.26h4.65V40.57h3.76V19.41H24.15V14.86H47.29v4.55H39.08v7.81H46.5v4.64H39.08v8.8h8.61Z"
                />
              </g>
            </g>
          </svg>
          已认证
        </view>
      </view>
      <!-- 种植基地 -->
      <view class="card-content-inner-box" v-if="configureData.plantShow">
        <view class="inner-header">
          种植基地
        </view>
        <view class="inner-body">
          <view class="content-info">
            <view class="content-info-row">
              <text space="emsp" decode="true" class="label" :class="themeColor"
                >基地名称: {{ '' }}</text
              >
              <text class="value">{{ plantInfoData.farmName }}</text>
            </view>
            <view class="content-info-row">
              <text space="emsp" decode="true" class="label" :class="themeColor"
                >基地地址: {{ '' }}</text
              >
              <text class="value">{{ plantInfoData.farmLocation }}</text>
            </view>
          </view>
          <view class="plant-bg">
            <Map
              :class="{ mapbox: true }"
              :def-lyrs="defLyrs"
              @created="onCreatedMap"
            ></Map>
          </view>
        </view>
      </view>
      <!-- 气象实施监测 -->
      <view
        class="card-content-inner-box"
        v-if="configureData.meteorologicalShow"
      >
        <view class="inner-header">
          <view class="meteorological-title">
            <view>
              气象实时监测
            </view>
            <view class="time">
              {{ clock | FormatTime }}
            </view>
          </view>
        </view>
        <view class="inner-body">
          <view class="meteorological-content">
            <view class="meteorological-card">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 80 80"
                :class="meteorologicalSvg"
              >
                <g id="图层_2" data-name="图层 2">
                  <g id="图层_1-2" data-name="图层 1">
                    <g id="温度">
                      <rect class="cls-1" width="80" height="80" />
                      <path
                        d="M24,65.07a9.11,9.11,0,0,1-2.3-17.94v-30a2.29,2.29,0,0,1,4.58,0v30A9.11,9.11,0,0,1,24,65.07Zm9.16-22.23V17.16a9.16,9.16,0,0,0-18.32,0V42.84a16,16,0,1,0,18.32,0Z"
                      />
                      <path
                        d="M49.17,21.71a11.42,11.42,0,1,0,8.09,3.34,11.42,11.42,0,0,0-8.09-3.34Z"
                      />
                      <path
                        d="M49.17,19.43a2.25,2.25,0,0,0,2.26-2.26V12.59a2.29,2.29,0,0,0-4.58,0v4.57a2.25,2.25,0,0,0,2.3,2.26Z"
                      />
                      <path
                        d="M49.17,46.8a2.31,2.31,0,0,0-2.3,2.26v4.56a2.29,2.29,0,1,0,4.58,0V49.1a2.25,2.25,0,0,0-2.26-2.26Z"
                      />
                      <path
                        d="M62.11,42.8a2.29,2.29,0,0,0-2.21-.6,2.3,2.3,0,0,0-1,3.84l3.25,3.21h0a2.29,2.29,0,0,0,1.62.68A2.25,2.25,0,0,0,65.34,46Z"
                      />
                      <path
                        d="M69.77,30.84H65.19a2.28,2.28,0,1,0,0,4.56h4.58a2.21,2.21,0,0,0,1.54-.6A2.24,2.24,0,0,0,72,33.31v-.37a2.24,2.24,0,0,0-2.25-2.1Z"
                      />
                      <path
                        d="M60.5,24.11a2.22,2.22,0,0,0,1.61-.69l3.23-3.21h0a2.25,2.25,0,0,0,0-3.24,2.33,2.33,0,0,0-3.23,0l-3.25,3.24h0a2.26,2.26,0,0,0,1.64,3.88Z"
                      />
                    </g>
                  </g>
                </g>
              </svg>
              <view class="card-content">
                <view class="title" :class="themeColor">温度</view>
                <view>{{ weather.temper }}℃</view>
              </view>
            </view>
            <view class="meteorological-card">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 80 80"
                :class="meteorologicalSvg"
              >
                <g id="图层_2" data-name="图层 2">
                  <g id="图层_1-2" data-name="图层 1">
                    <g id="湿度">
                      <rect class="cls-1" width="80" height="80" />
                      <path
                        d="M32,65.07a9.1,9.1,0,0,1-2.25-17.93V38.8a2.29,2.29,0,1,1,4.57,0v8.34A9.11,9.11,0,0,1,32,65.07Zm-2.38-47.8a2.24,2.24,0,0,1,1.13-1.95,2.21,2.21,0,0,1,2.25,0,2.24,2.24,0,0,1,1.13,1.95v12a2.26,2.26,0,0,1-4.51,0ZM41.15,42.83V17.16a9.16,9.16,0,1,0-18.31,0V42.85a16,16,0,1,0,18.31,0Z"
                      />
                      <path
                        d="M50.3,19.45H61.74a2.28,2.28,0,0,0,0-4.56H50.3a2.28,2.28,0,0,0,0,4.56Z"
                      />
                      <path
                        d="M50.3,28.58h4.58a2.29,2.29,0,0,0,2.26-2.29A2.25,2.25,0,0,0,54.88,24H50.3A2.25,2.25,0,0,0,48,26.29a2.34,2.34,0,0,0,.68,1.62,2.3,2.3,0,0,0,1.62.67Z"
                      />
                      <path
                        d="M61.74,33.12H50.3a2.29,2.29,0,0,0,0,4.57H61.74a2.29,2.29,0,0,0,0-4.57Z"
                      />
                      <path
                        d="M54.88,42.24H50.3a2.28,2.28,0,0,0,0,4.56h4.58a2.3,2.3,0,0,0,2-1.14,2.28,2.28,0,0,0-2-3.42Z"
                      />
                    </g>
                  </g>
                </g>
              </svg>
              <view class="card-content">
                <view class="title" :class="themeColor">湿度</view>
                <view>{{ weather.humidity }}%</view>
              </view>
            </view>
            <view class="meteorological-card">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 80 80"
                :class="meteorologicalSvg"
              >
                <g id="图层_2" data-name="图层 2">
                  <g id="图层_1-2" data-name="图层 1">
                    <g id="风速">
                      <rect class="cls-1" width="80" height="80" />
                      <path
                        d="M57,25.28A75.65,75.65,0,0,1,50.09,22V42.36a49.22,49.22,0,0,0,6.87-.1Z"
                      />
                      <path
                        d="M70,27.77a2.35,2.35,0,0,0-1.75-.6,22.19,22.19,0,0,1-6.7-.48V41.51a33.33,33.33,0,0,0,8-3.21,2.26,2.26,0,0,0,1.19-2V29.45A2.28,2.28,0,0,0,70,27.78Z"
                      />
                      <path
                        d="M27.18,37.52,20.29,32V29.66l6.89-5.48ZM34,16.34a25.44,25.44,0,0,0-4.78-.06,2.27,2.27,0,0,0-1.38.64,2.25,2.25,0,0,0-.66,1.36l-6.94,5.54V10.29a2.3,2.3,0,0,0-4.59,0V69.71a2.3,2.3,0,0,0,4.59,0V37.91l6.9,5.47v.33a2.28,2.28,0,0,0,1.12,1.95,2.25,2.25,0,0,0,1.16.32,2.34,2.34,0,0,0,1.13-.28,29.49,29.49,0,0,1,3.46-1.63Z"
                      />
                      <path
                        d="M45.5,19.72a39.72,39.72,0,0,0-6.87-2.58V42.7a30.45,30.45,0,0,1,6.87-.54Z"
                      />
                      <path
                        d="M11.53,72a2.26,2.26,0,0,1,0-4.52H24.34a2.26,2.26,0,1,1,0,4.52Z"
                      />
                    </g>
                  </g>
                </g>
              </svg>
              <view class="card-content">
                <view class="title" :class="themeColor">风速</view>
                <view>{{ weather.wind }}m/s</view>
              </view>
            </view>
            <view class="meteorological-card">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 80 80"
                :class="meteorologicalSvg"
              >
                <g id="图层_2" data-name="图层 2">
                  <g id="图层_1-2" data-name="图层 1">
                    <g id="降水">
                      <rect class="cls-1" width="80" height="80" />
                      <path
                        d="M37.55,72a2.43,2.43,0,0,1-1.75-.7,2.4,2.4,0,0,1-.73-1.74V54.77a2.47,2.47,0,0,1,4.94,0V69.55a2.47,2.47,0,0,1-.71,1.72,2.44,2.44,0,0,1-1.73.71Z"
                      />
                      <path
                        d="M27.63,67.07a2.42,2.42,0,0,1-2.43-2.43V49.84a2.46,2.46,0,0,1,4.92,0V64.62a2.4,2.4,0,0,1-.71,1.72,2.42,2.42,0,0,1-1.72.72Z"
                      />
                      <path
                        d="M47.39,67.07a2.42,2.42,0,0,1-2.47-2.43V49.84a2.47,2.47,0,1,1,4.94,0V64.62a2.44,2.44,0,0,1-2.43,2.44Z"
                      />
                      <path
                        d="M47.39,8A24.49,24.49,0,0,0,23.86,25.23a16,16,0,0,0-3.62,31.55V49.84a7.39,7.39,0,0,1,14.52-1.92,7.52,7.52,0,0,1,5.52,0,7.39,7.39,0,0,1,14.53,1.92v6.24A24.59,24.59,0,0,0,47.4,8Z"
                      />
                    </g>
                  </g>
                </g>
              </svg>
              <view class="card-content">
                <view class="title" :class="themeColor">降水</view>
                <view>{{ weather.rain }}cm</view>
              </view>
            </view>
          </view>
        </view>
      </view>
      <!-- 种植过程 -->
      <view class="card-content-inner-box">
        <view class="inner-header">
          种植过程
        </view>
        <view class="inner-body">
          <view class="process-content">
            <steps :steps="steps" :current="steps.length - 1">
              <template #date="{date}">
                <view class="date__up">
                  {{ date | dateFormatter('up') }}
                </view>
                <view class="date__down">
                  {{ date | dateFormatter('down') }}
                </view>
              </template>
              <template #scene="{scene}">
                <view class="weather-box">
                  <view class="weather-icon" v-html="scene.weather"> </view>
                  <text class="temper">{{ scene.temper }}℃</text>
                </view>
              </template>
              <template #body="{body}">
                <view v-if="body.desc" class="desc-wrapper">
                  <view
                    v-for="(text, index) in body.desc"
                    :key="index"
                    class="desc-p"
                    >{{ text }}</view
                  >
                </view>
                <view v-if="body.imgs" class="img-wrapper">
                  <image
                    v-for="(img, index) in body.imgs"
                    :key="img"
                    class="img"
                    :src="img"
                    @click="preViewImg(index, body.imgs)"
                  />
                </view>
              </template>
            </steps>
          </view>
        </view>
        <view
          class="inner-foot"
          @click="jump('/entry-1/wholeProcess')"
          v-if="templateType === 'template1'"
        >
          <text>全部过程</text>
          <u-icon
            name="arrow-right"
            size="16"
            color="#212121"
            class="icon"
          ></u-icon>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { preViewImg, getDate, getSteps } from '../../utils/tool'
import Steps from '@/components/steps'
import IDrawBase from '../plant/IDrawBase' //重载地块颜色
import common from '../plant/common.js' //引入初始化地图
import Map from '@/components/farm-map/Map' //引入地图
import { getPlantInfo } from '@/api/myRice'
import { QuerySectionSByFarm } from '@/api/farm-gis/commonApi'
export default {
  name: 'Plant',
  components: { Steps, Map },
  props: ['menuTree', 'templateType'],
  filters: {
    /**
     * @description: 根据UI设计稿处理日期展示
     * @param {String} date
     * @param {String} flagStr
     * @return {String}
     * @author: Hemingway
     */
    dateFormatter(date, flagStr) {
      const isDate = !isNaN(Date.parse(date))
      if (isDate) {
        if (flagStr === 'up') {
          return date.slice(5)
        } else {
          return date.slice(0, 4)
        }
      } else {
        if (flagStr === 'up') {
          return date.slice(0, 3)
        } else {
          return date.slice(3)
        }
      }
    },
    FormatTime: function(val) {
      return getDate(val, 'year') //根据传参不同显示的时间类型也不同
    }
  },
  mixins: [common, IDrawBase],
  data() {
    return {
      steps: [], // 加工步骤时间线
      allSteps: [], // 全部时间线
      defLyrs: ['tianVec_c', 'tianCva_c'], //地图地图图层
      farmID: '', // 农场id
      plantInfoData: {},
      weather: {},
      farmRecords: [],
      clock: Date.parse(new Date()),
      times: null,
      configureData: {
        plantShow: true,
        meteorologicalShow: true
      },
      excludeArr: ['日期', '农事图片', '信息采集员']
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    isStaticData() {
      return this.$store.getters.isStaticData
    },
    packingInfo() {
      return this.$store.getters.packingInfo
    },
    themeColor() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'theme1-color'
      } else if (colorId === '3') {
        return 'theme2-color'
      } else {
        return ''
      }
    },
    themeBackground() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'theme1-background'
      } else if (colorId === '3') {
        return 'theme2-background'
      } else {
        return ''
      }
    },
    themeLinear() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'theme1-linear'
      } else if (colorId === '3') {
        return 'theme2-linear'
      } else {
        return ''
      }
    },
    meteorologicalSvg() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'meteorological-green-svg'
      } else if (colorId === '3') {
        return 'meteorological-blue-svg'
      } else {
        return 'meteorological-yellow-svg'
      }
    }
  },

  watch: {
    // 深度监听，及时更新页面
    menuTree: {
      handler() {
        this.getChecked()
      },
      deep: true
    }
  },

  async created() {
    if (this.isStaticData) {
      let { plant } = require('../../static/json/static.json')
      this.plantInfoData = plant || {}
      this.weather = plant.realWeatherDto || {}
      this.farmRecords = plant.farmRecords || []
      this.$store.commit('app/SET_FARM_RECORDS', this.farmRecords)
      this.farmID = plant.farmId || ''
      this.init()
      this.getChecked()
      return
    }
    await this.plantInfo()
    this.getChecked()
  },

  mounted() {
    let _this = this
    this.times = setInterval(function() {
      _this.clock = Date.parse(new Date())
    }, 1000)
  },
  destroyed() {
    // 清除定时器
    if (this.times) {
      clearInterval(this.times)
    }
  },

  methods: {
    getChecked() {
      if (!this.menuTree?.subTree) return
      // 获取种植基地isChecked
      this.configureData.plantShow =
        this.menuTree.subTree.find(res => res.menuId === '22').isChecked === 'Y'
          ? true
          : false
      // 获取气象实施检测isChecked
      this.configureData.meteorologicalShow =
        this.menuTree.subTree.find(res => res.menuId === '23').isChecked === 'Y'
          ? true
          : false
      // 根据配置信息获取steps数据
      getSteps(this.menuTree).then(res => {
        // 如果是模板一就只加载前三个流程，模板二加载全部流程
        console.log('种植流程', res)
        if (this.templateType === 'template1') {
          this.steps = res.slice(0, 3)
        } else {
          this.steps = res
        }
      })
    },
    /**
     * add by yumanshan  gis map
     */
    async getAreaData() {
      let result = await QuerySectionSByFarm({ farmId: this.farmID })
      if (result.code === '0') {
        this.areaData = result.data
        //获取地块数据
        this.areaData = this.coorFilter(this.areaData)
        //渲染地块数据
        this.feas = this.drawFeas(this.areaData)
        if (this.feas) {
          //获取地块边界
          let extent = window.cxgis.ol.factory.MapFactory.getExtent(
            this.map,
            this.feas
          )
          this.map.getView().fit(extent, { padding: [50, 50, 50, 50] })
        }
      }
    },
    /**
     * 创建地图完成后回调进行数据加载
     */
    init() {
      this.getFarmCoors(this.farmID).then(res => {
        if (res === null || res.length === 0) {
          console.log('没有农场编号或者编号错误!')
          return
        }
        this.map
          .getView()
          .setCenter(
            this.transformLocate([
              res[0]['farm_longitude'],
              res[0]['farm_latitude']
            ])
          )
      })
      this.getAreaData()
    },
    preViewImg(i, images) {
      preViewImg(i, images)
    },
    /**
     * @description: 跳转二级页面
     * @param {String} path
     * @param {String} type
     * @author: Hemingway
     */
    jump(path) {
      this.$Router.push({
        path,
        query: {
          id: this.qrcode,
          templateType: this.templateType
        }
      })
    },
    /**
     * @description: 获取页面初始化信息
     * @author: qjj
     */
    async plantInfo() {
      const { data } = await getPlantInfo({
        qrCode: this.qrcode,
        machPackingId: this.packingInfo && this.packingInfo.machPackingId,
        packingBatchId: this.packingInfo && this.packingInfo.packingBatchId
      })
      if (!data) return
      this.plantInfoData = data || {}
      this.weather = data.realWeatherDto || {}
      this.farmRecords = data.farmRecords || []
      this.$store.commit('app/SET_FARM_RECORDS', this.farmRecords)
      this.farmID = data.farmId || ''
      this.init()
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../static/style/template1.scss';
.card-content-inner-box {
  .section {
    margin-top: 48rpx;
  }

  .content-info {
    .content-info-row {
      margin-bottom: 25rpx;
    }

    .label {
      color: #eabc02;
    }
  }

  .img-box {
    width: 100%;
    box-sizing: border-box;
    height: 330rpx;
    margin-top: 40rpx;

    image {
      width: 100%;
      height: 330rpx;
    }
  }

  .inner-header > .meteorological-title {
    display: flex;
    justify-content: space-between;
    .time {
      font-weight: normal;
      color: #8b8b8b;
    }
  }

  .inner-body {
    // 气象内容
    .meteorological-content {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      .meteorological-card:nth-child(1),
      .meteorological-card:nth-child(2) {
        margin-bottom: 40rpx;
      }
      .meteorological-card {
        width: 45%;
        box-sizing: border-box;
        height: 80rpx;
        display: flex;
        .meteorological-yellow-svg {
          width: 84rpx;
          height: 84rpx;
          border-radius: 20rpx;
          background-color: $theme-color-yellow-linear;
          fill: $theme-color-yellow;
          .cls-1 {
            fill: none;
          }
        }
        .meteorological-green-svg {
          width: 84rpx;
          height: 84rpx;
          border-radius: 20rpx;
          background-color: $theme-color-green-linear;
          fill: $theme-color-green;
          .cls-1 {
            fill: none;
          }
        }
        .meteorological-blue-svg {
          width: 84rpx;
          height: 84rpx;
          border-radius: 20rpx;
          background-color: $theme-color-blue-linear;
          fill: $theme-color-blue;
          .cls-1 {
            fill: none;
          }
        }
        .card-content {
          margin-left: 60rpx;
          font-size: 32rpx;
          .title {
            color: #eabc02;
          }
        }
      }
    }

    // 时间线
    .process-content {
      background-color: #fff;

      .desc-wrapper {
        padding-top: 24rpx;

        .desc-p {
          word-break: break-all;
          color: #8b8b8b;
        }
      }

      .img-wrapper {
        padding: 32rpx 0 0;
        display: flex;
        flex-wrap: wrap;

        .img {
          display: block;
          margin-right: 16rpx;
          width: 120rpx;
          height: 120rpx;
        }
      }

      .weather-box {
        display: inline-block;
        margin-left: 10rpx;
        color: #8b8b8b;

        .weather-icon {
          width: 40rpx;
          height: 40rpx;
          vertical-align: middle;
          display: inline-block;
        }
        .temper {
          font-size: 28rpx;
          margin-left: 8rpx;
        }
      }
    }

    // 地图
    .plant-bg {
      width: 100%;
      height: 384rpx;
      background-size: 100% 384rpx;
      border-radius: 16rpx;
      margin: 0 auto;
    }
  }

  .inner-foot {
    padding: 28rpx;
    text-align: center;
    background-color: #fafafa;
    color: #8b8b8b;
    .icon {
      display: inline-block;
      margin-left: 16rpx;
    }
  }
}
</style>
