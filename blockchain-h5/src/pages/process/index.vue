<!--
 * @Author       : guoxi
 * @Date         : 2021-05-24 09:55:30
 * @LastEditTime : 2022-07-07 13:43:33
 * @LastEditors  : Your Name
 * @Description  : 加工过程
 * @FilePath     : \blockchain-h5\src\pages\process\index.vue
-->
<template>
  <view class="process">
    <scroll-view :scroll-y="true">
      <!-- 加工信息 -->
      <Machining
        v-if="configInfor.machining.isShow"
        :menuTree="configInfor.machining.subTree"
      />
    </scroll-view>
  </view>
</template>

<script>
import { preViewImg } from '../../utils/tool'
import Machining from '../module/machining.vue' // 加工信息
export default {
  name: 'Process',

  components: {
    Machining
  },

  data() {
    return {}
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    configInfor() {
      return this.$store.getters.configInfor
    },
    mainInfoData() {
      return this.$store.getters.mainInfoData
    }
  },

  created() {
    uni.setNavigationBarTitle({
      title: `${this.mainInfoData.brandName}区块链溯源`
    })
  },

  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    }
  }
}
</script>

<style lang="scss" scoped>
.process {
  padding: 16rpx;
  margin-top: -22rpx;
  // ::v-deep .uni-scroll-view {
  //   height: calc(100vh - 138rpx);
  // }
}
</style>
