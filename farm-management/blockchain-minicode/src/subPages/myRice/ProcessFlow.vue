<!--
 * @Author: guoxi
 * @Date: 2021-05-24 09:55:30
 * @LastEditTime : 2021-07-08 11:40:35
 * @LastEditors  : Please set LastEditors
 * @Description: 进度
 * @FilePath     : \blockchain-minicode\src\subPages\myRice\ProcessFlow.vue
-->

<template>
  <view class="process">
    <!-- <view class="process-title"> 烘干批次：xxxxxx </view> -->
    <view class="process-content">
      <mSidebar class="a" title="进度详情">
        <view class="row">
          <m-steps :item="dry" date="updateTime" :activity="0">
            <text slot="status">检测报告</text>
            <view slot="content">
              <view
                class="process-btn btn-disable"
                @tap="jump('BlockChainReport')"
                >检测报告</view
              >
            </view>
          </m-steps>

          <m-steps :item="dry" date="updateTime" :activity="0">
            <view slot="status">生产加工环节</view>
            <view slot="content">
              <view class="process-btn btn-disable" @tap="jump('Process', '3')"
                >加工数据</view
              >
            </view>
          </m-steps>

          <m-steps :item="dry" date="updateTime" :activity="0">
            <view slot="status">种植环节</view>
            <view slot="content">
              <view class="process-btn btn-disable" @tap="jump('Plant')"
                >种植信息</view
              >
            </view>
          </m-steps>
        </view>
      </mSidebar>
    </view>
  </view>
</template>

<script>
import mSidebar from '@/components/m-sidebar/m-sidebar.vue'
import MSteps from '../../components/m-steps/m-steps.vue'
import { getSrvInfoByJgBatch } from '@/api/myRice'
import { stringifyDate } from '@/utils/tool'
export default {
  name: 'Process',
  components: {
    mSidebar,
    MSteps
  },

  onLoad(option) {
    const info = JSON.parse(decodeURIComponent(option.info))

    this.batchId = info.batchId
    this.getData()
  },
  data() {
    return {
      dry: {
        updateTime: null
      },
      batchId: '',
      commInfo: null,
      uniId: '',
      varietyName: '', //品种
      detailed_address: '', //产地
      commodityName: '', // 商品名称
      qualityGuaPeriod: '', // 保质期
      reap_create_date: '', //收割时间
      pack_date: '',
      qualityGrade: ''
    }
  },
  mounted() {},

  methods: {
    async getData() {
      let data = await getSrvInfoByJgBatch({
        machPackingId: this.batchId
      })

      this.commInfo = data.resObj.commInfo

      if (data.resObj.blockchainInfo.data) {
        const res = JSON.parse(data.resObj.blockchainInfo.data)
        res['normsForMachPack'] = data.resObj.normsForMachPack
        res['brandName'] = this.commInfo.brandName
        res['commodityAttachment'] = this.commInfo.commodityAttachment
        res['hgEnterpriseImgs'] = data.resObj.hgEnterpriseImgs
        res['hgEnterpriseName'] = data.resObj.hgEnterpriseName
        res['jgcEnterpriseImgs'] = data.resObj.jgcEnterpriseImgs
        res['jgcEnterpriseName'] = data.resObj.jgcEnterpriseName
        res['sourceFarmId'] = data.resObj.sourceFarmId
        res['sourceFarmName'] = data.resObj.sourceFarmName
        uni.setStorageSync('h5Info', res)
        this.reap_create_date = stringifyDate(res.reap_create_date)
        this.pack_date = stringifyDate(
          res.process_record_detail.package.pack_date
        )
        this.detailed_address = res.detailed_address
      }
      this.varietyName = this.commInfo.varietyName
      this.commodityName = this.commInfo.commodityName
      this.qualityGuaPeriod = this.commInfo.qualityGuaPeriod
      this.qualityGrade = this.commInfo.qualityGrade
    },
    jump(url) {
      this.$Router.push({ name: url })
    }
  }
}
</script>

<style lang="scss">
.process-title {
  position: relative;
  height: 80rpx;
  line-height: 80rpx;
  text-indent: 33rpx;
  background-color: #fafafa;
  font-family: PingFangSC-Regular;
  font-size: 24rpx;
  color: #8b8b8b;

  .process-doc {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    right: 73rpx;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #00c853;
  }
}

.process-content {
  width: 100%;
  background-color: #fff;
  padding: 40rpx;

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

.process-btn {
  width: 178rpx;
  height: 56rpx;
  line-height: 56rpx;
  display: inline-block;
  margin: 24rpx 24rpx 0 0;
  border-radius: 28rpx;
  border: solid 1rpx #e0e0e0;
  font-size: 28rpx;
  text-align: center;
  cursor: pointer;
  pointer-events: none;
}

.btn-disable {
  background-color: #00c853;
  color: #fff;
  pointer-events: unset;
}

.a {
  position: relative;
  left: -120rpx;
}
</style>
