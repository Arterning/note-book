<!--
 * @Author       : guoxi
 * @Date         : 2021-05-24 09:55:30
 * @LastEditTime : 2022-07-07 13:43:06
 * @LastEditors  : Your Name
 * @Description  : 种植过程
 * @FilePath     : \blockchain-h5\src\pages\plant\index.vue
-->
<template>
  <view class="plant">
    <Plant
      v-if="configInfor.plant.isShow"
      :menuTree="configInfor.plant.subTree"
      :templateType="'template2'"
    />
  </view>
</template>

<script>
import Plant from '../module/plant.vue' // 种植信息
import { preViewImg } from '../../utils/tool'

export default {
  name: 'PlantPage',

  components: {
    Plant
  },

  // mixins: [common, IDrawBase],

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

  mounted() {},

  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    }
  }
}
</script>

<style lang="scss" scoped>
.plant {
  padding: 16rpx;
  margin-top: -22rpx;
}
</style>
