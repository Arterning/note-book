<!--
 * @Description  : 平台入口
 * @Autor        : WuJing
 * @Date         : 2021-10-20 16:15:44
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-26 11:14:59
 * @FilePath     : \blockchain-screen\src\components\PlatformView.vue
-->
<template>
  <div class="platform-view">
    <div class="content">
      <div>
        <div><img src="../assets/img/app.png" class="picture" /></div>
        <div class="text">智农云APP（种植端）</div>
      </div>

      <div>
        <div><img src="../assets/img/min.jpg" class="picture" /></div>
        <div class="text">溯源平台小程序（加工端）</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PlatformView',
  data() {
    return {}
  }
}
</script>
<style lang="scss" scoped>
.platform-view {
  margin-top: 20px;
  color: #fff;
  background-image: url('../assets/img/platform.png');
  background-repeat: no-repeat;
  height: 372px;
  width: 900px;
  padding-bottom: 10px;
  text-align: center;

  .content {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
    padding-top: 94px;

    .picture {
      width: 208px;
      height: 208px;
      background-color: rgba(0, 102, 255, 0.4);
      // background-image: url('../assets/img/min.jpg');
      // background-repeat: no-repeat;
    }

    .text {
      margin-top: 20px;
      font-size: 18px;
    }
  }
}
</style>
