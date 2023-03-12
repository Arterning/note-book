<!--
 * @Description  : 解决方案
 * @Autor        : WuJing
 * @Date         : 2021-10-26 14:24:24
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-03 10:30:02
 * @FilePath     : \blockchain-screen\src\views\Solution.vue
-->
<template>
  <div class="solution">
    <div class="header">
      <div class="button" @click="handle">返回大屏</div>
    </div>
    <div class="container">
      <el-row>
        <el-col :span="6">
          <div class="introduction">
            <img
              class="picture"
              src="../assets/img/certificate.png"
              alt=""
            /></div
        ></el-col>
        <el-col :span="12"
          ><div class="margin-left">
            <div class="management"></div>
            <div class="approve"></div></div
        ></el-col>
        <el-col :span="6"><div class="advantage"></div></el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Solution',
  components: {},
  methods: {
    handle() {
      this.$router.push('/')
    }
  }
}
</script>

<style lang="scss" scoped>
.solution {
  background: #070b26;
  padding: 20px;
  min-width: 1880px;
  min-height: 1060px;
  background-image: url('../assets/img/bg-introduce.png');
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

  .introduction {
    width: 470px;
    height: 925px;
    background-image: url('../assets/img/introduction.png');
    background-repeat: no-repeat;
    text-align: center;
    padding-top: 45px;

    .picture {
      margin-top: 40px;
      height: 530px;
      border-radius: 8px;
    }
  }

  .management {
    background-image: url('../assets/img/management.png');
    background-repeat: no-repeat;
    width: 900px;
    height: 480px;
  }

  .approve {
    margin-top: 10px;
    background-image: url('../assets/img/approve.png');
    background-repeat: no-repeat;
    width: 900px;
    height: 480px;
    margin-bottom: 5px;
  }

  .margin-left {
    margin-left: 20px;
  }

  .advantage {
    background-image: url('../assets/img/advantage.png');
    background-repeat: no-repeat;
    width: 470px;
    height: 961px;
  }
}
</style>
