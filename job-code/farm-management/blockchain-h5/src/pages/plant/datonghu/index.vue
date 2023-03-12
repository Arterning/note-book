<!--
 * @Description  : 种植过程 - 为大通湖定制
 * @Autor        : Hemingway
 * @Date         : 2021-12-16 19:09:33
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-06-13 18:16:52
 * @FilePath     : \blockchain-h5\src\pages\plant\datonghu\index.vue
-->
<template>
  <view class="plant">
    <scroll-view :scroll-y="true">
      <view class="plant-title"> 种植地块 </view>
      <!-- 基地信息 -->
      <view class="base">
        <view class="plant-bg-title">{{ sectionInfo.name + '6号地块' }}</view>
        <view class="plant-bg-des">{{ sectionInfo.detailed_address }}</view>
        <view class="plant-bg">
          <Map
            :class="{ pointer: show, mapbox: true }"
            :def-lyrs="defLyrs"
            @created="onCreatedMap"
          ></Map>
        </view>
      </view>
      <!-- 实时气象 -->
      <view v-if="isShow" class="real-msg">
        <view class="header-wrapper">
          <view class="header">
            <text class="title">气象监测数据</text>
            <text class="time">{{ actDate }}</text>
          </view>
          <text>通过气象设备实时监测气象情况，为数字化种植提供依据。</text>
        </view>

        <view class="msg-wrapper">
          <view class="msg"
            ><i class="wd"></i>温度：<text class="value"
              >{{ airTemperature }}℃</text
            ></view
          >

          <view class="msg"
            ><i class="sd"></i>湿度：<text class="value"
              >{{ relativeHumidity }}%</text
            ></view
          >

          <view class="msg"
            ><i class="fs"></i>风速：<text class="value"
              >{{ averageWindSpeed }}级</text
            ></view
          >

          <view class="msg"
            ><i class="js"></i>降水：<text class="value"
              >{{ rainfall }}cm</text
            ></view
          >
        </view>
      </view>

      <view class="plant-title">
        种植信息（种植方式：再生稻）
      </view>
      <view class="plant-content">
        <view
          v-for="item in stages"
          :key="item.name"
          class="stage"
          @click="onChooseStage(item.name, item.steps)"
        >
          {{ item.name }}
        </view>
      </view>
    </scroll-view>

    <bottom-btn text="点击查看产品区块链证书" @click.native="onView" />
  </view>
</template>

<script>
import { getFarmWeatherInfo } from '@/api/myRice'
import { preViewImg } from '@/utils/tool'
import Map from '@/components/farm-map/Map' //引入地图
import { queryUnionGra } from '@/api/farm-gis/AttBase' //请求数据库的数据
import IDrawBase from '../IDrawBase' //重载地块颜色
import common from '../common.js' //引入初始化地图
import BottomBtn from '@/components/bottom-btn'
import stages from './stages.js'

export default {
  name: 'Plant',

  components: {
    Map,
    BottomBtn
  },

  mixins: [common, IDrawBase],

  data() {
    return {
      actDate: '', //物联网更新时间
      airTemperature: '', //温度：
      relativeHumidity: '', //湿度
      averageWindSpeed: '', //风速
      rainfall: '', //降水
      isShow: false,
      creator: '',
      tx_create_time: '',
      tx_id: '',
      block_number: '',
      sectionInfo: {
        name: '',
        detailed_address: ''
      },
      stages, // 生育期数据
      farmID: '', // 农场id
      defLyrs: ['tianVec_c', 'tianCva_c'] //地图地图图层
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    }
  },

  created() {
    const data = uni.getStorageSync('h5Info')

    this.creator = data.brandName
    this.farmID = data.sourceFarmId

    this.sectionInfo.name = data.sourceFarmName
    this.sectionInfo.detailed_address = data.detailed_address

    const { tx_create_time, tx_id, block_number } = data.farming_record_detail

    this.tx_create_time = tx_create_time
    this.tx_id = tx_id
    this.block_number = block_number
  },

  mounted() {
    this.getFarmWeatherInfo()
  },

  methods: {
    /**
     * @description: 获取气象信息
     * @author: Hemingway
     */
    async getFarmWeatherInfo() {
      try {
        const { data } = await getFarmWeatherInfo(this.farmID)
        if (data.actDate !== null) {
          this.actDate = data.actDate
          this.airTemperature = data.airTemperature
          this.relativeHumidity = data.relativeHumidity
          this.averageWindSpeed = data.averageWindSpeed
          this.rainfall = data.rainfall

          this.isShow = true
        }
      } catch (error) {
        console.log('查询气象信息 error = ', error)
      }
    },

    preViewImg(i, images) {
      preViewImg(i, images)
    },

    /**
     * add by yumanshan  gis map
     */
    /**
     * 获取地块数据（含坐标）
     * @method
     * @returns {Promise}
     */
    getAreaData() {
      const param = {
        _major: 3,
        _minor: 0,
        _exp: `a.minor=? and a.id=b.id and c.farm_id=${this.farmID} and c.is_active='Y' and c.status='Y' and b.farm_id=c.farm_id and b.name=c.name and b.is_active='Y' and b.status='Y'`,
        cols: 'a:*;c:*',
        vals: '2',
        types: 'i',
        orderby: 'id desc',
        tables: `a;b:3,2;c:3,1`
      }
      return queryUnionGra(param).then(res => {
        this.drawFeas(res)
      })
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

    /**
     * @description: 跳转区块链证书页面
     * @author: guoxi
     */
    onView() {
      uni.navigateTo({
        url: `/certificate?id=${this.qrcode}&cername=种植过程区块链&tx_id=${this.tx_id}&block_number=${this.block_number}&tx_create_time=${this.tx_create_time}&creator=${this.creator}`
      })
    },

    /**
     * @description: 跳转生育期详情
     * @param {String} name 生育期名称
     * @param {Array} steps 生育期环节
     * @author: Hemingway
     */
    onChooseStage(name, steps) {
      uni.navigateTo({
        url: `/plant-datonghu-stage?id=${
          this.qrcode
        }&name=${name}&steps=${encodeURIComponent(JSON.stringify(steps))}`
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.plant {
  ::v-deep .uni-scroll-view {
    height: calc(100vh - 138rpx);

    .plant-title {
      position: relative;
      height: 80rpx;
      line-height: 80rpx;
      text-indent: 33rpx;
      font-family: PingFangSC-Regular;
      font-size: 24rpx;
      color: #8b8b8b;
    }

    .base {
      padding: 40rpx 32rpx;
      background-color: #fff;

      .plant-bg-title {
        width: 686rpx;
        margin: 0 auto 16rpx auto;
        font-size: 32rpx;
        font-weight: normal;
        font-stretch: normal;
        letter-spacing: 0;
        color: #212121;
      }

      .plant-bg-des {
        width: 686rpx;
        font-family: PingFangSC-Regular;
        font-size: 28rpx;
        color: #8b8b8b;
        margin: 0 auto 25rpx auto;
      }

      .plant-bg {
        width: 686rpx;
        height: 384rpx;
        background-size: 686rpx 384rpx;
        border-radius: 16rpx;
        margin: 0 auto;
      }
    }

    // 气象数据
    .real-msg {
      margin: 16rpx 16rpx 0;
      background-color: #fff;
      border: 2rpx solid #eee;
      border-radius: 16rpx;

      .header-wrapper {
        position: relative;
        padding: 36rpx 24rpx;
        font-size: 28rpx;
        color: #8b8b8b;

        .header {
          margin-bottom: 24rpx;
          display: flex;
          justify-content: space-between;
          align-items: center;

          .title {
            color: #212121;
            font-size: 32rpx;
          }

          .time {
            color: #00c853;
          }
        }
      }

      .header-wrapper::after {
        content: '';
        position: absolute;
        border-bottom: 2rpx solid #eee;
        transform: scaleY(0.5);
        left: 24rpx;
        right: 24rpx;
        bottom: 0;
      }

      .msg-wrapper {
        padding: 40rpx 24rpx 16rpx;
        display: flex;
        flex-wrap: wrap;

        .msg {
          margin-bottom: 24rpx;
          flex: 50%;
          font-size: 28rpx;
          font-weight: normal;
          font-stretch: normal;
          letter-spacing: 0;
          color: #8b8b8b;

          .value {
            color: #212121;
            font-size: 28rpx;
          }

          i {
            display: inline-block;
            width: 40rpx;
            height: 40rpx;
            vertical-align: middle;
            margin: 0 30rpx;
          }

          .wd {
            background: url(../../../static/img/WD.png) no-repeat;
            background-size: 40rpx 40rpx;
          }

          .fs {
            background: url(../../../static/img/FS.png) no-repeat;
            background-size: 40rpx 40rpx;
          }

          .sd {
            background: url(../../../static/img/SD.png) no-repeat;
            background-size: 40rpx 40rpx;
          }

          .js {
            background: url(../../../static/img/JS.png) no-repeat;
            background-size: 40rpx 40rpx;
          }

          text {
            font-family: PingFangSC-Regular;
            font-size: 32rpx;
            font-weight: normal;
            letter-spacing: 0;
            color: #212121;
          }
        }
      }
    }

    .plant-content {
      box-sizing: border-box;
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      grid-template-rows: 120rpx;
      grid-gap: 8rpx;

      .stage {
        height: 120rpx;
        line-height: 120rpx;
        text-align: center;
        color: #212121;
        background-color: #fff;
        font-size: 28rpx;
      }
    }
  }
}
</style>
