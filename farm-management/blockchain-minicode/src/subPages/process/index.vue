<!--
 * @Author: guoxi
 * @Date: 2021-05-24 09:55:30
 * @LastEditTime : 2021-11-24 13:51:48
 * @LastEditors  : Please set LastEditors
 * @Description: 进度
-->

<template>
  <view class="process">
    <view class="process-title">
      加工过程
    </view>
    <view class="process-content">
      <mSidebar title="进度详情">
        <view class="row">
          <m-steps :item="dry" date="hg_date" :activity="0">
            <text slot="status">烘干</text>
            <view slot="content">
              <view class="common-title"
                >烘干公司：{{ dry.hg_factory_name }}</view
              >
              <view class="common-imgs">
                <img
                  v-for="(item, index) in dry.image_url"
                  :key="index"
                  :src="item"
                  @click="preViewImg(index, dry.image_url)"
                />
              </view>
            </view>
          </m-steps>

          <m-steps :item="process" date="process_date" :activity="0">
            <view slot="status">加工</view>
            <view slot="content">
              <view class="common-title"
                >加工公司：{{ process.process_enterprise_name }}</view
              >
              <view class="common-imgs">
                <img
                  v-for="(item, index) in process.process_image_url"
                  :key="index"
                  :src="item"
                  @click="preViewImg(index, process.process_image_url)"
                />
                <!-- <img :src="process.process_image_url" /> -->
              </view>
            </view>
          </m-steps>

          <m-steps :item="packages" date="pack_date" :activity="0">
            <view slot="status">包装</view>
            <view slot="content">
              <view class="common-title"
                >规格：{{ packages.each_bag_weight }}</view
              >
              <view class="common-imgs">
                <!-- <img :src="packages.package_image_url" /> -->
                <img
                  v-for="(item, index) in packages.package_image_url"
                  :key="index"
                  :src="item"
                  @click="preViewImg(index, packages.package_image_url)"
                />
              </view>
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
import { stringifyDate, preViewImg } from '../../utils/tool'
export default {
  name: 'Process',
  components: {
    mSidebar,
    MSteps
  },

  data() {
    return {
      creator: '',
      tx_create_time: '',
      tx_id: '',
      block_number: '',
      activity: 0,
      dry: {
        hg_date: '',
        hg_factory_name: '',
        image_url: []
      },
      process: {
        process_date: '',
        process_enterprise_name: '',
        process_image_url: []
      },
      warehouse: {
        cc_factory_name: '',
        duration_date: '',
        image_url: [],
        ref_temperature: '',
        start_date: '',
        day: 0
      },
      packages: {
        each_bag_weight: '',
        pack_date: '',
        package_image_url: []
      }
    }
  },
  created() {
    const data = uni.getStorageSync('h5Info')
    this.packages.each_bag_weight = data.normsForMachPack
    this.creator = data.brandName
    this.tx_create_time = data.process_record_detail.tx_create_time
    this.tx_id = data.process_record_detail.tx_id
    this.block_number = data.process_record_detail.block_number
    this.dry.hg_date = stringifyDate(data.process_record_detail.dry.hg_date)
    this.dry.hg_factory_name = data.hgEnterpriseName
    const dry_url = data.jgcEnterpriseImgs.attachments.dryEnvironment
    if (dry_url.length > 0) {
      this.dry.image_url.push(this.getImageUrl(dry_url[0].id))
    }

    this.warehouse.duration_date = stringifyDate(
      data.process_record_detail.warehouse.duration_date
    )
    this.warehouse.day = this.DateMinus(
      stringifyDate(data.process_record_detail.warehouse.start_date),
      stringifyDate(data.process_record_detail.warehouse.duration_date)
    )

    this.warehouse.cc_factory_name =
      data.process_record_detail.warehouse.cc_factory_name
    this.warehouse.ref_temperature =
      data.process_record_detail.warehouse.ref_temperature
    this.warehouse.start_date = data.process_record_detail.warehouse.start_date
    if (this.validation(data.process_record_detail.warehouse.image_url)) {
      const urls = JSON.parse(data.process_record_detail.warehouse.image_url)
      urls.map(el => {
        this.warehouse.image_url.push(this.getImageUrl(el.id))
      })
    }

    // this.process = data.process_record_detail.process
    this.process.process_date = stringifyDate(
      data.process_record_detail.process.process_date
    )
    this.process.process_enterprise_name = data.jgcEnterpriseName

    const process_url = data.jgcEnterpriseImgs.attachments.processingEnvironment
    if (process_url.length > 0) {
      this.process.process_image_url.push(this.getImageUrl(process_url[0].id))
    }

    this.warehouse.process_date = stringifyDate(this.warehouse.process_date)

    // this.packages = data.process_record_detail.package
    this.packages.pack_date = stringifyDate(
      data.process_record_detail.package.pack_date
    )

    const packages_url = data.jgcEnterpriseImgs.attachments.packageEnvironment
    if (packages_url.length > 0) {
      this.packages.package_image_url.push(this.getImageUrl(packages_url[0].id))
    }
  },
  methods: {
    getImageUrl(id) {
      return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
    },
    preViewImg(i, images) {
      preViewImg(i, images)
    },
    DateMinus(date1, date2) {
      //兼容ios
      if (isNaN(Date.parse(date1))) {
        date1 = date1.replace(/-/g, '/')
      }
      if (isNaN(Date.parse(date2))) {
        date2 = date2.replace(/-/g, '/')
      }
      var sdate = new Date(date1)
      var now = new Date(date2)
      var days = now.getTime() - sdate.getTime()
      var day = parseInt(days / (1000 * 60 * 60 * 24))
      return day
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
.process {
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
      height: 150px;
    }

    .node {
      left: 156rpx;
      width: 24rpx;
      height: 24rpx;
      background-color: #eee;
    }

    .wrapper {
      padding-left: 220rpx;

      .wrapper {
        font-family: PingFangSC-Regular;
        font-size: 28rpx;
        color: #212121;
      }
    }
  }
}

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
    right: 37rpx;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #00c853;
  }
}

.process-content {
  width: 100%;
  background-color: #fff;
  margin-top: 40rpx;
  // padding: 40rpx;

  .row {
    padding: 24rpx 0;
  }

  .common-title {
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    margin-top: 26rpx;
    color: #8b8b8b;
  }

  .common-imgs {
    width: 500rpx;
    margin: 32rpx 0 41rpx 0;

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
</style>
