<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 15:45:31
 * @FilePath     : \blockchain-h5\src\pages\entry-1\index.vue
-->
<template>
  <view class="">
    <scroll-view
      scroll-y="true"
      :scroll-into-view="intoViewY"
      scroll-with-animation="true"
      :style="`height: ${innerHeight}px`"
      @scroll="bindScrollY"
    >
      <view class="entry-1" v-if="configInfor">
        <!-- 认证信息 -->
        <Authentication
          id="authentication"
          ref="authentication"
          :mainInfoData="mainInfoData"
        />
        <!-- 产品信息 -->
        <Product
          id="product"
          ref="product"
          v-if="configInfor.product.isShow"
          :menuTree="configInfor.product.subTree"
        />
        <!-- 质检信息 -->
        <Quality
          id="quality"
          ref="quality"
          v-if="configInfor.quality.isShow"
          :menuTree="configInfor.quality.subTree"
        />
        <!-- 种植信息 -->
        <Plant
          id="plant"
          ref="plant"
          v-if="configInfor.plant.isShow"
          :menuTree="configInfor.plant.subTree"
          :templateType="'template1'"
        />
        <!-- 加工信息 -->
        <Machining
          id="machining"
          ref="machining"
          v-if="configInfor.machining.isShow"
          :menuTree="configInfor.machining.subTree"
        />
        <!-- 企业信息 -->
        <Enterprise
          id="enterprise"
          ref="enterprise"
          :qrcode="qrcode"
          v-if="configInfor.enterprise.isShow"
          :menuTree="configInfor.enterprise.subTree"
        />
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
    </scroll-view>
    <!-- 长图文 -->
    <Drag v-if="graphic.title" :data="graphic" :qrcode="qrcode"></Drag>
    <view class="foot-tabbar" @click="tabbarBtn">
      <scroll-view
        scroll-x="true"
        :scroll-into-view="intoViewX"
        scroll-with-animation="true"
        style="white-space: nowrap;"
      >
        <view
          class="foot-tabbar-item"
          :style="{ color: readMoreColor }"
          v-for="(item, index) in tabbarData"
          :key="item.id"
          :class="active == index && activeItem"
          :id="`${item.anchorId}${item.id}`"
          >{{ item.title }}</view
        >
      </scroll-view>
    </view>
  </view>
</template>

<script>
import Authentication from '../module/authentication.vue' // 认证信息
import Product from '../module/product.vue' // 产品信息
import Quality from '../module/quality.vue' // 质检信息
import Plant from '../module/plant.vue' // 种植信息
import Machining from '../module/machining.vue' // 加工信息
import Enterprise from '../module/enterprise.vue' // 企业信息
import Drag from '../entry-2/components/Drag.vue'
// import { getImageUrl } from '../../utils/tool'
import templateConfig from '@/static/json/templateConfig.js'
export default {
  name: 'Entry1',
  components: {
    Authentication,
    Product,
    Quality,
    Plant,
    Machining,
    Enterprise,
    Drag
  },

  data() {
    return {
      intoViewY: '',
      intoViewX: '',
      active: '0',
      tabbarData: [
        // { id: '0', title: '认证信息', anchorId: 'authentication' },
        // { id: '1', title: '产品信息', anchorId: 'product' },
        // { id: '2', title: '质检信息', anchorId: 'quality' },
        // { id: '3', title: '种植信息', anchorId: 'plant' },
        // { id: '4', title: '加工信息', anchorId: 'machining' },
        // { id: '5', title: '企业信息', anchorId: 'enterprise' }
      ],
      moduleIdAndAnchorId: [
        { moduleId: '1', anchorId: 'authentication' },
        { moduleId: '2', anchorId: 'product' },
        { moduleId: '3', anchorId: 'quality' },
        { moduleId: '4', anchorId: 'plant' },
        { moduleId: '5', anchorId: 'machining' },
        { moduleId: '6', anchorId: 'enterprise' }
      ],
      // 获取浏览器的视口高度
      innerHeight:
        window.innerHeight ||
        document.documentElement.clientHeight ||
        document.body.clientHeight,
      pageDom: {},
      resObj: {
        scanCodeNum: '无',
        lastTime: '无'
      },
      configInfor: {
        product: {
          isShow: true,
          subTree: []
        },
        quality: {
          isShow: true,
          subTree: []
        },
        plant: {
          isShow: true,
          subTree: []
        },
        machining: {
          isShow: true,
          subTree: []
        },
        enterprise: {
          isShow: true,
          subTree: []
        }
      },
      graphic: {
        title: '',
        imgUrls: []
      }
    }
  },

  watch: {
    // 监听active值的变化，动态变化底部导航栏滚动条
    active(newValue) {
      this.$nextTick(() => {
        if (newValue <= 2) {
          this.intoViewX = this.tabbarData[0].anchorId + this.tabbarData[0].id
          return
        }
        this.intoViewX = this.tabbarData[newValue].anchorId + newValue
      })
    },
    '$store.state.app.configurationInformation': {
      handler(newValue) {
        this.handleConfigInfor(newValue)
        this.initTabbarData(newValue)
      },
      deep: true
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    configurationInformation() {
      return this.$store.getters.configurationInformation
    },
    // 溯源码主要信息
    mainInfoData() {
      return this.$store.getters.mainInfoData
    },
    activeItem() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'active-item-green'
      } else if (colorId === '3') {
        return 'active-item-blue'
      } else {
        return 'active-item-yellow'
      }
    },
    readMoreColor() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return '#00c853'
      } else if (colorId === '3') {
        return '#2196f3'
      } else {
        return '#eabc02'
      }
    }
  },

  created() {
    uni.setNavigationBarTitle({
      title: `${this.mainInfoData.brandName}区块链溯源`
    })
  },

  mounted() {
    this.handleConfigInfor(this.configurationInformation || templateConfig)
    this.initTabbarData(this.configurationInformation || templateConfig)
    this.getLongText()
  },

  methods: {
    /**
     * @description: // 处理长图文信息
     * @author: qjj
     */
    getLongText() {
      let title = this.mainInfoData.title,
        purchaseLink = this.mainInfoData.purchaseLink
      if (title) {
        this.graphic.title =
          title.length > 6 ? title.slice(0, 5) + '...' : title
      }
      if (purchaseLink) {
        // this.graphic.imgUrls = JSON.parse(purchaseLink).map(res =>
        //   getImageUrl(res.id)
        // )
        this.graphic.imgUrls = purchaseLink
      }
    },
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
    },
    /**
     * @description: 初始化tabbarData
     * @author: qjj
     */
    initTabbarData(params) {
      let newArr = []
      if (params) {
        params.forEach(res => {
          if (res.isChecked !== 'Y') return
          newArr.push({
            id: newArr.length && newArr.length,
            title: res.moduleName,
            anchorId: this.moduleIdAndAnchorId.find(
              item => item.moduleId === res.moduleId
            ).anchorId
          })
        })
        this.tabbarData = newArr
      }
    },
    /**
     * @description: 导航栏点击事件
     * @author: qjj
     */
    tabbarBtn(event) {
      this.active = event.target.id.split('').pop()
      let anchorId = this.tabbarData.filter(
        (res, index) => index === Number(this.active)
      )[0].anchorId
      this.$nextTick(() => {
        this.intoViewY = anchorId
      })
    },
    /**
     * @description: 垂直滚动触发事件
     * @author: qjj
     */
    bindScrollY() {
      let authenticationTop = this.$refs.authentication?.$el.getBoundingClientRect()
        .top
      let product = this.$refs.product?.$el.getBoundingClientRect().top
      let quality = this.$refs.quality?.$el.getBoundingClientRect().top
      let plant = this.$refs.plant?.$el.getBoundingClientRect().top
      let machining = this.$refs.machining?.$el.getBoundingClientRect().top
      let enterprise = this.$refs.enterprise?.$el.getBoundingClientRect().top
      let effectHeight = this.innerHeight * 0.5
      if (0 <= authenticationTop && authenticationTop <= effectHeight) {
        this.active = '0'
        return
      }
      if (0 <= product && product <= effectHeight) {
        this.active = this.tabbarData.find(res => res.anchorId === 'product').id
        return
      }

      if (0 <= quality && quality <= effectHeight) {
        this.active = this.tabbarData.find(res => res.anchorId === 'quality').id
        return
      }
      if (0 <= plant && plant <= effectHeight) {
        this.active = this.tabbarData.find(res => res.anchorId === 'plant').id
        return
      }
      if (0 <= machining && machining <= effectHeight) {
        this.active = this.tabbarData.find(
          res => res.anchorId === 'machining'
        ).id
        return
      }
      if (0 <= enterprise && enterprise <= effectHeight) {
        this.active = this.tabbarData.find(
          res => res.anchorId === 'enterprise'
        ).id
        return
      }
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
          id: this.qrcode
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../static/style/template1.scss';
.entry-1 {
  background-color: #fafafa;
  padding: 16rpx;
  // margin-bottom: 80rpx;

  .foot-btn {
    display: flex;
    justify-content: space-between;
    padding: 20rpx 0 10rpx;
    .btn-class {
      width: 48%;
    }
  }
}
.foot-desc {
  padding: 30rpx;
  color: #c9c4c4;
  font-size: 24rpx;
  text-align: center;
  margin-bottom: 90rpx;
}
.foot-tabbar {
  position: fixed;
  bottom: 0;
  border-top: 1px solid #eeeeee;
  background-color: white;
  width: 100%;
  .foot-tabbar-item {
    display: inline-block;
    font-size: 24rpx;
    padding: 30rpx;
    color: #eabc02;
  }
}
</style>
