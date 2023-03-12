<!--
 * @Description  : 长图介绍
 * @Autor        : Hemingway
 * @Date         : 2021-11-05 16:41:30
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-06-13 13:45:53
 * @FilePath     : \blockchain-h5\src\pages\graphic\index.vue
-->
<template>
  <view class="graphic">
    <image
      v-for="(url, index) in imgUrls"
      :key="url"
      class="img"
      :src="url"
      mode="widthFix"
      @click="preViewImg(index, imgUrls)"
    ></image>
  </view>
</template>

<script>
import { preViewImg } from '@/utils/tool'
export default {
  name: 'Graphic',

  data() {
    return {
      imgUrls: []
    }
  },

  created() {
    const { title, imgUrls } = this.$Route.query
    uni.setNavigationBarTitle({
      title
    })
    this.imgUrls = imgUrls
  },

  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    }
  }
}
</script>

<style lang="scss" scoped>
.graphic {
  .img {
    display: block;
    width: 100%;
  }
}
</style>
