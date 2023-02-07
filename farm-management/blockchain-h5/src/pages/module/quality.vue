<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-05 14:51:10
 * @FilePath     : \blockchain-h5\src\pages\module\quality.vue
-->
<template>
  <view>
    <!-- 质检信息 -->
    <view class="card-info common-box" :class="themeLinear">
      <view class="card-title" :class="themeColor">
        <text>质检信息</text>
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
      <view class="card-content-box">
        <view class="content-info">
          <view class="content-info-row" v-if="configureData.dateShow">
            <text space="emsp" decode="true" class="label" :class="themeColor"
              >检测日期: {{ '' }}</text
            >
            <text class="value">{{ checkInfoData.checkDate }}</text>
          </view>
          <view class="content-info-row" v-if="configureData.mechanismShow">
            <text space="emsp" decode="true" class="label" :class="themeColor"
              >检测机构: {{ '' }}</text
            >
            <text class="value">{{ checkInfoData.checkOrg }}</text>
          </view>
          <!-- <view class="content-info-row" v-if="configureData.resultShow">
            <text space="emsp" decode="true" class="label" :class="themeColor"
              >检测结果: {{ '' }}</text
            >
            <text class="value">{{ checkInfoData.qualified }}</text>
          </view> -->
        </view>
        <view class="report-card section">
          <image
            v-for="(image, index) in images"
            :key="index"
            :src="image"
            class="card-img"
            @click="preViewImg(index, images)"
          />
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { preViewImg, getImageUrl } from '../../utils/tool'
import { getCheckInfo } from '@/api/myRice'
export default {
  name: 'Quality',
  props: ['menuTree'],
  data() {
    return {
      images: [], // 商品介绍视频
      // 页面配置信息
      configureData: {
        dateShow: true,
        mechanismShow: true,
        resultShow: true
      },
      checkInfoData: {}
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
    // 主题颜色
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

  created() {
    if (this.isStaticData) {
      let { quality } = require('../../static/json/static.json')
      this.checkInfoData = quality || {}
      this.images = quality.checkResults || []
      return
    }
    this.checkInfo()
  },

  mounted() {
    this.getChecked()
  },

  methods: {
    getChecked() {
      // 获取检测日期isChecked
      this.configureData.dateShow =
        this.menuTree?.subTree &&
        this.menuTree.subTree.find(res => res.menuId === '18').isChecked === 'Y'
          ? true
          : false
      // 获取检测机构isChecked
      this.configureData.mechanismShow =
        this.menuTree?.subTree &&
        this.menuTree.subTree.find(res => res.menuId === '19').isChecked === 'Y'
          ? true
          : false
      // 获取检测结果isChecked
      // this.configureData.resultShow =
      //   this.menuTree?.subTree &&
      //   this.menuTree.subTree.find(res => res.menuId === '20').isChecked === 'Y'
      //     ? true
      //     : false
    },
    preViewImg(i, images) {
      preViewImg(i, images)
    },
    /**
     * @description: 获取页面初始化信息
     * @author: qjj
     */
    async checkInfo() {
      const { data } = await getCheckInfo({
        qrCode: this.qrcode,
        machPackingId: this.packingInfo && this.packingInfo.machPackingId,
        packingBatchId: this.packingInfo && this.packingInfo.packingBatchId
      })
      if (!data) return
      this.checkInfoData = data || {}
      this.images =
        (data.checkImages &&
          data.checkImages.split(',').map(id => getImageUrl(id))) ||
        []
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../static/style/template1.scss';
.card-content-box {
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

  .report-card {
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;
    .card-img {
      width: 48%;
      height: 330rpx;
      margin-bottom: 20rpx;
      border: 2rpx solid #eee;
      box-sizing: border-box;
    }
  }
}
</style>
