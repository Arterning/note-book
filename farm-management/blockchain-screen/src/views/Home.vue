<!--
 * @Description  : 首页
 * @Autor        : Hemingway
 * @Date         : 2021-10-06 22:10:11
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-26 14:37:20
 * @FilePath     : \blockchain-screen\src\views\Home.vue
-->
<template>
  <div class="home">
    <div class="header">
      <div class="button" @click="handle">功能介绍</div>
    </div>
    <div class="container">
      <el-row>
        <el-col :span="6">
          <div>
            <command-view class="border" /><enterprise-view
              class="border margin"
            /></div
        ></el-col>
        <el-col :span="12"
          ><div class="margin-left">
            <map-view /><platform-view class="border" /></div
        ></el-col>
        <el-col :span="6"
          ><div>
            <link-view class="border" /><scandata-view
              class="border margin"
            /></div
        ></el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import MapView from '../components/MapView'
import CommandView from '../components/CommandView'
import EnterpriseView from '../components/EnterpriseView'
import LinkView from '../components/LinkView'
import PlatformView from '../components/PlatformView.vue'
import ScandataView from '../components/ScandataView.vue'

export default {
  name: 'Home',
  components: {
    MapView,
    CommandView,
    EnterpriseView,
    LinkView,
    PlatformView,
    ScandataView
  },
  methods: {
    handle() {
      this.$router.push('/solution')
    }
  }
}
</script>
<style lang="scss" scoped>
.home {
  background: #070b26;
  padding: 20px;
  min-width: 1880px;
  min-height: 1060px;
  background-image: url('../assets/img/BG.png');
  background-repeat: no-repeat;
  background-size: 100% 100%; // 满屏

  .header {
    height: 70px;
    color: #fff;

    .button {
      margin-top: 33px;
      width: 120px;
      height: 48px;
      font-size: 16px;
      line-height: 48px;
      text-align: center;
      border-radius: 8px;
      cursor: pointer;
      background-image: url('../assets/img/border.png');
      background-repeat: no-repeat;
    }
  }

  .border {
    border-radius: 8px;
  }

  .margin {
    margin-top: 20px;
  }

  .margin-left {
    margin-left: 20px;
  }
}
</style>
