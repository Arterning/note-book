<!--
 * @Author: guoxi
 * @Date: 2021-05-24 09:55:30
 * @LastEditTime: 2021-10-11 10:17:23
 * @LastEditors: Please set LastEditors
 * @Description: 进度
 * @FilePath     : \blockchain-minicode-h5\src\pages\process\plant.vue
-->

<template>
  <view class="plant">
    <view class="plant-title"> 种植基地 </view>
    <view class="bg-color">
      <view class="plant-bg-title">{{ sectionInfo.name }}</view>
      <view class="plant-bg-des">{{ sectionInfo.detailed_address }}</view>
      <!-- <view class="plant-bg">
        <Map
          :class="{ pointer: show, mapbox: true }"
          :def-lyrs="defLyrs"
          @created="onCreatedMap"
        ></Map>
      </view> -->
    </view>

    <view class="msg-list">
      <view class="list-item">
        <view class="text">物联网实时数据</view>
        <view class="text">2020-10-10 00:00:00</view>
      </view>
      <view style="padding-bottom: 25rpx;">
        <view class="list-item" style="height: 80rpx; line-height: 80rpx;">
          <view class="item-text"
            ><i class="wd"></i>温度：<text>18℃</text></view
          >
          <view class="item-text"
            ><i class="sd"></i>湿度：<text>50.42%</text></view
          >
        </view>
        <view class="list-item" style="height: 80rpx; line-height: 80rpx;">
          <view class="item-text"
            ><i class="fs"></i>风速：<text>4级</text></view
          >
          <view class="item-text"
            ><i class="js"></i>降水：<text>1.5cm</text></view
          >
        </view>
      </view>
    </view>

    <view class="plant-title">
      种植信息
    </view>
    <view class="bg-color" style="padding-top: 15px;">
      <view class="plant-content">
        <mSidebar title="进度详情">
          <view class="row">
            <m-steps
              v-for="(item, index) in stepList"
              :key="index"
              :item="item"
              date="updateTime"
              :activity="0"
            >
              <view slot="status"
                >{{ item.use_type
                }}<view class="common-time"
                  >操作人：{{ item.operator_name }}</view
                ></view
              >
              <view slot="content">
                <!-- <view :key="i" class="common-title">{{ item.use_type }} ></view> -->

                <view v-if="item.image_url" class="common-imgs">
                  <img
                    :src="getImage(item.image_url[0])"
                    @click="preViewImg(0, [getImage(item.image_url[0])])"
                  />
                </view>
              </view>
            </m-steps>
          </view>
        </mSidebar>
      </view>
    </view>
  </view>
</template>

<script>
// import { geth5SrvInfo } from '@/api/myRice'
import mSidebar from '@/components/m-sidebar/m-sidebar.vue'
import MSteps from '../../components/m-steps/m-steps.vue'
import { preViewImg, stringifyDate, getZNImageUrl } from '../../utils/tool'
// import { queryUnionGra } from '@/api/farm-gis/AttBase' //请求数据库的数据
export default {
  name: 'Plant',
  components: {
    mSidebar,
    MSteps
  },
  onLoad(option) {
    this.batchId = option.batchId
  },

  data() {
    return {
      creator: '',
      tx_create_time: '',
      tx_id: '',
      block_number: '',
      batchId: null,
      activity: 0,
      sectionInfo: {
        name: '',
        detailed_address: ''
      },
      plantList: [],
      stepList: [],

      farmID: '', // 农场id
      defLyrs: ['tianImg_c', 'tianCia_c'] //地图地图图层
    }
  },
  created() {
    const data = uni.getStorageSync('h5Info')
    this.farmID = data.sourceFarmId
    this.creator = data.brandName
    this.tx_create_time = data.farming_record_detail.tx_create_time
    this.tx_id = data.farming_record_detail.tx_id
    this.block_number = data.farming_record_detail.block_number
    this.sectionInfo = data.farming_record_detail.section_info
    this.sectionInfo.name = data.sourceFarmName
    const farming_record_detail = data.farming_record_detail.op_detail_list
    farming_record_detail.map(el => {
      el.record_detail_list.map(obj => {
        this.plantList.push({
          updateTime: stringifyDate(el.handle_date),
          image_url: el.image_url,
          operator_name: el.operator_name,
          ...obj
        })
      })
    })

    const bozhong = this.plantList.find(el => el.type === '7')
    if (bozhong) {
      bozhong.use_type = '播种'
      this.stepList.push(bozhong)
    }
    const jifei = this.plantList.find(
      el => el.type_code === 'fertilizeType' && el.use_good === '0'
    )
    if (jifei) {
      jifei.use_type = '基肥'
      this.stepList.push(jifei)
    }

    // OpDetail.RecordDetailList[i].TypeCode==sprayType && OpDetail.RecordDetailList[i].UseGood==”0”
    const chucao = this.plantList.find(
      el => el.type_code === 'sprayType' && el.use_good === '0'
    )
    if (chucao) {
      chucao.use_type = '除草'
      this.stepList.push(chucao)
    }

    // 6）防病治虫，显示农药、操作人、操作时间、图片，筛选条件
    // OpDetail.RecordDetailList[i].TypeCode==sprayType && OpDetail.RecordDetailList[i].UseGood==”1”
    // 7）收割，显示收割时间，筛选条件
    // QRCodeResponse.ReapCreateDate
    const yizai = this.plantList.find(el => el.type === '10')
    if (yizai) {
      yizai.use_type = '移栽'
      this.stepList.push(yizai)
    }

    const fennie = this.plantList.find(
      el =>
        el.type_code === 'fetilizeType' &&
        (el.use_good === '10' || el.use_good === '2')
    )
    if (fennie) {
      fennie.use_type = '分蘖肥'
      this.stepList.push(fennie)
    }

    const fangchong = this.plantList.find(
      el => el.type_code === 'sprayType' && el.use_good === '1'
    )
    if (fangchong) {
      fangchong.use_type = '防病治虫'
      this.stepList.push(fangchong)
    }
    this.stepList.push({
      use_type: '收割',
      updateTime: stringifyDate(data.reap_create_date),
      operator_name: this.stepList[0].operator_name,
      image_url: null
    })
  },
  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    },
    getImage(item) {
      return getZNImageUrl(item)
    },

    validation(val) {
      if (val !== 'null' && val !== '' && val !== '<invalid Value>') {
        return true
      }
      return false
    }
  }
}
</script>

<style lang="scss" scoped>
.plant {
  background-color: #fff;

  ::v-deep .steps {
    .date {
      left: 32rpx;
      width: 84rpx;
      font-size: 28rpx;
      color: #8b8b8b;
    }

    .tail {
      left: 166rpx;
      border-left: 2px solid #e4e7ed;
      height: 125px;
    }

    .node {
      left: 156rpx;
      width: 24rpx;
      height: 24rpx;
      background-color: #eee;
    }

    .wrapper {
      width: auto;
      padding-left: 220rpx;

      .wrapper {
        font-family: PingFangSC-Regular;
        font-size: 28rpx;
        color: #212121;
      }
    }
  }

  .bg-color {
    background-color: #fff;
  }

  .plant-title {
    position: relative;
    height: 80rpx;
    line-height: 80rpx;
    text-indent: 33rpx;
    background-color: #eee;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #8b8b8b;

    .plant-doc {
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      right: 37rpx;
      font-family: PingFangSC-Regular;
      font-size: 24rpx;
      color: #00c853;
    }
  }

  .msg-list {
    display: none;
    width: 718rpx;
    border-radius: 16rpx;
    background-color: #fff;
    margin: 16rpx auto;
    border-left: 1rpx solid #fafafa;
    border-right: 1rpx solid #fafafa;

    .list-item {
      display: flex;
      height: 104rpx;
      line-height: 104rpx;
      border-bottom: 1rpx solid #fafafa;

      .text:nth-child(1) {
        flex: 1;
        text-indent: 24rpx;
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        font-weight: normal;
        font-stretch: normal;
        letter-spacing: 0;
        color: #212121;
      }

      .text:nth-child(2) {
        flex: 1;
        position: relative;
        left: -24rpx;
        text-align: right;
        font-family: PingFangSC-Regular;
        font-size: 28rpx;
        font-weight: normal;
        font-stretch: normal;
        letter-spacing: 0;
        color: #8b8b8b;
      }

      .item-text {
        flex: 1;
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        font-weight: normal;
        font-stretch: normal;
        letter-spacing: 0;
        color: #8b8b8b;

        i {
          display: inline-block;
          width: 40rpx;
          height: 40rpx;
          vertical-align: middle;
          margin: 0 30rpx;
        }

        // .wd {
        //   background: url(../../static/img/WD.png) no-repeat;
        //   background-size: 40rpx 40rpx;
        // }

        // .fs {
        //   background: url(../../static/img/FS.png) no-repeat;
        //   background-size: 40rpx 40rpx;
        // }

        // .sd {
        //   background: url(../../static/img/SD.png) no-repeat;
        //   background-size: 40rpx 40rpx;
        // }

        // .js {
        //   background: url(../../static/img/JS.png) no-repeat;
        //   background-size: 40rpx 40rpx;
        // }

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

  .plant-bg {
    width: 686rpx;
    height: 384rpx;
    //background: url(../../static/img/plant.jpg) no-repeat;
    background-size: 686rpx 384rpx;
    border-radius: 16rpx;
    margin: 0 auto;
    padding-bottom: 40rpx;
  }

  .plant-bg-title {
    width: 686rpx;
    margin: 0 auto 16rpx auto;
    padding-top: 40rpx;
    font-family: PingFangSC-Regular;
    font-size: 28rpx;
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

  .plant-content-title {
    width: 144rpx;
    height: 56rpx;
    line-height: 56rpx;
    text-align: center;
    background-color: #00c8531a;
    border-radius: 28rpx;
    color: #00c853;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    margin: 33rpx 0 0 95rpx;
  }

  .plant-content {
    width: 100%;
    background-color: #fff;
    // padding: 40rpx 0 0 40rpx;

    .row {
      padding: 24rpx 0;
    }

    .common-title {
      font-family: PingFangSC-Regular;
      font-size: 24rpx;
      margin: 32rpx 0;
      color: #8b8b8b;
    }

    .common-imgs {
      width: 500rpx;
      margin: 40rpx 0;

      img {
        display: inline-block;
        margin-right: 17rpx;
        width: 102rpx;
        height: 102rpx;

        &:last-child {
          margin-right: unset;
        }
      }
    }

    .common-time {
      position: absolute;
      right: 33rpx;
      top: 0;
      font-family: PingFangSC-Regular;
      font-size: 24rpx;
      color: #8b8b8b;
    }
  }
}
</style>
