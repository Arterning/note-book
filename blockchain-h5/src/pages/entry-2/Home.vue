<!--
 * @Author       : guoxi
 * @Date         : 2021-05-24 09:55:30
 * @LastEditTime : 2022-08-03 16:13:02
 * @LastEditors  : Your Name
 * @Description  : 溯源首页
 * @FilePath     : \blockchain-h5\src\pages\entry-2\Home.vue
-->
<template>
  <view class="home">
    <!-- 轮播 -->
    <!-- <swiper
      class="swiper"
      :indicator-dots="imgUrls.length > 1"
      :autoplay="true"
      :interval="2000"
      indicator-color="#8b8b8b"
      indicator-active-color="#e0e0e0"
      :circular="true"
    >
      <swiper-item v-for="(item, index) in imgUrls" :key="index">
        <img class="swiper-image" :src="item" mode="aspectFill" />
      </swiper-item>
    </swiper> -->
    <!-- 图片 -->
    <!-- <view class="swiper">
      <img
        class="swiper-image"
        src="../../static/img/home/banner.png"
        mode="aspectFill"
      />
    </view> -->
    <!-- 产品、企业信息 -->
    <Product
      v-if="configInfor.product.isShow"
      :menuTree="configInfor.product.subTree"
    />
    <Enterprise
      v-if="configInfor.enterprise.isShow"
      :menuTree="configInfor.enterprise.subTree"
    />
    <!-- 长图文 -->
    <Drag v-if="graphic.title" :data="graphic" :qrcode="qrcode"></Drag>

    <view class="card-container">
      <view class="card">
        <view class="card__left">
          <image src="/static/img/home/1.jpg" class="image" />
          <view class="desc">
            <text class="title">安全与品质</text>
            <!-- <text>等级：{{ qualityLevel }}</text> -->
          </view>
        </view>
        <view
          v-if="isShowReport"
          class="card__right"
          @click="jump('/report', '1')"
        >
          <text>质检信息</text>
          <u-icon
            name="arrow-right"
            size="16"
            color="#212121"
            class="icon"
          ></u-icon>
        </view>
      </view>

      <view class="card">
        <view class="card__left">
          <image src="/static/img/home/2.jpg" class="image" />
          <view class="desc">
            <text class="title">种植管理</text>
            <!-- <text>收割时间：{{ mainInfoData.harvestDate || '' }}</text> -->
          </view>
        </view>
        <!-- 种植过程 - 为大通湖定制 -->
        <view class="card__right" @click="jump('/plant', '2')">
          <text>种植信息</text>
          <u-icon
            name="arrow-right"
            size="16"
            color="#212121"
            class="icon"
          ></u-icon>
        </view>
      </view>

      <view class="card">
        <view class="card__left">
          <image src="/static/img/home/3.jpg" class="image" />
          <view class="desc">
            <text class="title">加工管理</text>
            <!-- <text>包装时间：见外包装</text> -->
          </view>
        </view>
        <view class="card__right" @click="jump('/process', '3')">
          <text>加工信息</text>
          <u-icon
            name="arrow-right"
            size="16"
            color="#212121"
            class="icon"
          ></u-icon>
        </view>
      </view>
      <view class="foot-btn">
        <u-button
          type="success"
          :plain="true"
          text="给我好评"
          size="small"
          class="btn-class"
          icon="thumb-up-fill"
          @click="jump('/goodComment')"
        ></u-button>
        <u-button
          type="error"
          :plain="true"
          text="举报投诉"
          class="btn-class"
          size="small"
          icon="bell-fill"
          @click="jump('/badComment')"
        ></u-button>
      </view>
    </view>
    <view class="foot-desc">
      技术支持：中联智慧农业股份有限公司
    </view>
    <!-- <template v-if="purchaseLink">
      <bottom-btn text="线上直购" @click.native="jumpBuy">
        <template #main>
          <u-icon
            name="shopping-cart"
            size="24"
            color="#fff"
            class="icon"
          ></u-icon>
          <text> 线上直购</text>
        </template>
      </bottom-btn>
    </template> -->
  </view>
</template>

<script>
import { addAccRecord } from '@/api/myRice'
import { getFingerprintInfo, getImageUrl } from '../../utils/tool'
import Drag from './components/Drag.vue'
import Product from '../module/product.vue' // 产品信息
import Enterprise from '../module/enterprise.vue' // 企业信息

export default {
  name: 'Home',

  components: {
    Drag,
    Product,
    Enterprise
  },

  data() {
    return {
      imgUrls: [require('../../static/img/home/banner.png')],
      commInfo: null, // 商品基本信息
      uniId: '', // 指纹
      commodityName: '', // 商品名称
      detailed_address: '', // 产地
      qualityGuaPeriod: '', // 保质期

      textDesc: {
        content: '',
        shadowStyle: {
          backgroundImage:
            'linear-gradient(-180deg, rgba(255, 255, 255, 0) 0%, #fff 90%)',
          paddingTop: '100px',
          marginTop: '-100px',
          position: 'relative'
        }
      }, // 商品介绍文字
      video: {
        videoUrl: '',
        poster: ''
      }, // 商品介绍视频
      purchaseLink: '', // 商品购买链接
      graphic: {
        title: '',
        imgUrls: []
      }, // 商品介绍图文附件

      qualityGrade: '', // 商品品质等级
      reap_create_date: '', // 收割时间
      enterpriseNum: '', // 品牌商企业id

      isShowReport: true // 是否展示报告入口
    }
  },

  computed: {
    configInfor() {
      return this.$store.getters.configInfor
    },
    configurationInformation() {
      return this.$store.getters.configurationInformation
    },
    qualityLevel() {
      return this.$store.getters.qualityLevel
    },
    mainInfoData() {
      return this.$store.getters.mainInfoData
    },
    // 溯源信息
    allInfo() {
      return this.$store.getters.allInfo
    },
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    // 计算产地名称：截取到市
    origin() {
      let res = ''

      const reg = /\S+?(市|州)/
      res =
        this.detailed_address &&
        this.detailed_address.match(reg)[0].replace('中国', '')

      return res
    }
  },

  async created() {
    uni.setNavigationBarTitle({
      title: `${this.mainInfoData.brandName}区块链溯源`
    })
    this.uniId = await getFingerprintInfo() // 指纹
  },

  mounted() {
    this.handleConfigInfor(this.configurationInformation)
    // 处理拖拽溯源信息
    let title = this.mainInfoData.title,
      purchaseLink = this.mainInfoData.purchaseLink
    if (title) {
      this.graphic.title = title.length > 6 ? title.slice(0, 5) + '...' : title
    }
    if (purchaseLink) {
      this.graphic.imgUrls = JSON.parse(purchaseLink).map(res =>
        getImageUrl(res.id)
      )
    }
  },

  methods: {
    /**
     * @description: // 处理配置信息
     * @author: qjj
     */
    handleConfigInfor(arr) {
      this.$store.commit('app/SET_COLOR_ID', arr[0].colorId || '1')
      arr.forEach(res => {
        switch (res.moduleId) {
          case '2':
            this.configInfor.product.isShow = res.isChecked === 'Y'
            this.configInfor.product.subTree = res.subTree || []
            break
          case '3':
            this.configInfor.quality.isShow = res.isChecked === 'Y'
            this.configInfor.quality.subTree = res.subTree || []
            break
          case '4':
            this.configInfor.plant.isShow = res.isChecked === 'Y'
            this.configInfor.plant.subTree = res.subTree || []
            break
          case '5':
            this.configInfor.machining.isShow = res.isChecked === 'Y'
            this.configInfor.machining.subTree = res.subTree || []
            break
          case '6':
            this.configInfor.enterprise.isShow = res.isChecked === 'Y'
            this.configInfor.enterprise.subTree = res.subTree || []
            break
        }
      })
      this.$store.commit('app/SET_CONFIG_INFOR', this.configInfor)
    },

    /**
     * @description: 跳转二级页面
     * @param {String} path
     * @param {String} type
     * @author: Hemingway
     */
    jump(path) {
      // this.addRecord(type)

      this.$Router.push({
        path,
        query: {
          id: this.qrcode
        }
      })
    },

    /**
     * @description: 统计浏览器访问次数
     * @param {String} type
     * @author: Hemingway
     */
    addRecord(type) {
      addAccRecord({
        enterpriseId: this.commInfo.enterpriseId,
        brandId: this.commInfo.brandId,
        brandName: this.commInfo.brandName,
        commodityId: this.commInfo.commodityId,
        commodityName: this.commInfo.commodityName,
        tracingCode: this.qrcode, // 溯源码编号
        uniId: this.uniId, // 指纹
        accType: type // 首页标识（'0'首页 或 '1'检测报告 或 '2'种植过程 或 '3'加工过程）
      })
    },

    /**
     * @description: 跳转至购物页面
     * @author: Hemingway
     */
    jumpBuy() {
      location.href = this.purchaseLink
    }
  }
}
</script>

<style lang="scss" scoped>
.home {
  padding: 16rpx 16rpx 0;
  margin-top: -22rpx;
  .foot-btn {
    display: flex;
    justify-content: space-between;
    padding: 20rpx 0 10rpx;
    .btn-class {
      width: 48%;
    }
  }

  .foot-desc {
    padding: 0 30rpx 30rpx;
    color: #c9c4c4;
    font-size: 24rpx;
    text-align: center;
  }

  .swiper {
    height: 402rpx;

    .swiper-image {
      width: 100%;
      border-radius: 16rpx;
    }
  }

  .introduction {
    position: relative;
    margin-top: 24rpx;
    padding: 104rpx 24rpx 0;
    background-image: linear-gradient(to bottom, #fcf8ea 0%, #fff 100%);
    border-radius: 16rpx;

    .title {
      position: absolute;
      top: 0;
      left: 50%;
      transform: translateX(-50%);
      height: 104rpx;
      line-height: 104rpx;
      font-size: 36rpx;
      color: #eabc02;
      display: flex;
      align-items: center;

      .circle {
        width: 34rpx;
        height: 16rpx;
      }

      .circle.mirror {
        transform: rotateY(180deg);
      }

      .text {
        margin: 0 18rpx;
      }
    }

    .introduction-inner {
      padding: 24rpx 40rpx 0;
      font-size: 32rpx;
      color: #212121;
      background-color: #fff;
      border-radius: 16rpx 16rpx 0 0;

      .brief {
        .brief-row {
          margin-top: 12rpx;

          .label {
            color: #eabc02;
          }
        }
      }

      .section {
        margin-top: 48rpx;
      }

      .video {
        width: 100%;
        height: 330rpx;
      }

      .image {
        display: block;
        width: 100%;
      }
    }
  }

  // 卡片组
  .card-container {
    margin-top: 24rpx;
    padding-bottom: 20rpx;

    .card {
      height: 176rpx;
      margin-top: 16rpx;
      padding: 0 32rpx;
      display: flex;
      justify-content: space-between;
      align-items: center;
      border: 2rpx solid #eee;
      box-sizing: border-box;
      border-radius: 16rpx;
      background-color: #fff;
      font-size: 24rpx;

      .card__left {
        display: flex;
        height: 100%;
        box-sizing: border-box;

        .image {
          width: 120rpx;
          height: 120rpx;
          margin: auto 36rpx auto 0;
        }

        .desc {
          padding: 42rpx 0;
          display: flex;
          justify-content: center;
          flex-direction: column;

          .title {
            font-size: 32rpx;
            font-weight: bold;
          }
        }
      }

      .card__right {
        display: flex;
        align-items: center;

        .icon {
          margin-left: 16rpx;
        }
      }
    }
  }

  // 立即购买
  ::v-deep .bottom-btn {
    position: static;

    .btn {
      background-color: #dfba25;
    }

    .icon {
      margin-right: 24rpx;
    }
  }
}
</style>
