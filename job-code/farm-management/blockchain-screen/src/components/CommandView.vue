<!--
 * @Description  : 控住数据指挥舱
 * @Autor        : WuJing
 * @Date         : 2021-10-20 09:53:20
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-26 16:23:55
 * @FilePath     : \blockchain-screen\src\components\CommandView.vue
-->
<template>
  <div class="command-view">
    <div class="ship">
      <div class="cirlce-enterprise">
        <div class="count">{{ totalEnterpriseNum }}</div>
        <div class="msg">入驻企业数</div>
      </div>
      <div class="cirlce-product">
        <div class="count">{{ commodityNum }}</div>
        <div class="msg">溯源产品数</div>
      </div>
      <div class="cirlce-use">
        <div class="count">{{ traceCodeNum }}</div>
        <div class="msg">溯源码使用数</div>
      </div>
      <div class="cirlce-link">
        <div class="count">{{ uploadsDataNum }}</div>
        <div class="msg">上链数据（条）</div>
      </div>
    </div>
  </div>
</template>

<script>
import { queryCommand } from '../api/screen'
export default {
  name: 'CommandView',
  data() {
    return {
      //入驻企业数
      totalEnterpriseNum: 0,
      //商品总数
      commodityNum: 0,
      //溯源码数
      traceCodeNum: 0,
      //上链数据条数
      uploadsDataNum: 0
    }
  },
  mounted() {
    this.commandQuery()
  },
  methods: {
    /**
     * @description:查询指挥舱数据
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */
    async commandQuery() {
      try {
        const { code, data } = await queryCommand()
        if (code === '0') {
          this.totalEnterpriseNum = data.totalEnterpriseNum
          this.commodityNum = data.commodityNum
          this.traceCodeNum = data.traceCodeNum
          this.uploadsDataNum = data.uploadsDataNum
        }
      } catch (error) {
        console.log('error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.command-view {
  min-height: 568px;
  min-width: 470px;
  background-image: url('../assets/img/command.png');
  background-repeat: no-repeat;
  // background: rgba(0, 102, 255, 0.1);
  .ship {
    padding-top: 95px;
    margin-left: -5px;
    height: 450px;
    display: flex;
    flex-wrap: wrap;

    .msg {
      color: #fff;
      font-size: 18px;
      margin-top: -20px;
    }

    .cirlce-enterprise {
      margin-left: 20px;
      width: 205px;
      height: 205px;
      border-radius: 50%;
      text-align: center;
      background-image: url('../assets/img/green.png');
      background-repeat: no-repeat;

      .count {
        margin-top: 20px;
        font-weight: bold;
        line-height: 120px;
        font-size: 40px;
        color: #00e676;
      }
    }

    .cirlce-product {
      margin-left: 20px;
      width: 205px;
      height: 205px;
      border-radius: 50%;
      background-image: url('../assets/img/yellow.png');
      background-repeat: no-repeat;
      text-align: center;

      .count {
        margin-top: 20px;
        font-weight: bold;
        line-height: 120px;
        font-size: 40px;
        color: #ffd600;
      }
    }

    .cirlce-use {
      margin-left: 20px;
      width: 205px;
      height: 205px;
      border-radius: 50%;
      text-align: center;
      background-image: url('../assets/img/blue.png');
      background-repeat: no-repeat;

      .count {
        margin-top: 20px;
        font-weight: bold;
        line-height: 120px;
        font-size: 40px;
        color: #00e5ff;
      }
    }

    .cirlce-link {
      margin-left: 20px;
      width: 205px;
      height: 205px;
      border-radius: 50%;
      text-align: center;
      background-image: url('../assets/img/red.png');
      background-repeat: no-repeat;

      .count {
        margin-top: 20px;
        font-weight: bold;
        line-height: 120px;
        font-size: 40px;
        color: #ff1744;
      }
    }
  }
}
</style>
