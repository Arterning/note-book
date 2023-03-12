<!--
 * @Description  : 区块链证书
 * @Autor        : Hemingway
 * @Date         : 2021-10-21 10:26:40
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 16:48:13
 * @FilePath     : \blockchain-h5\src\pages\forbidden\index.vue
-->
<template>
  <view class="certificate">
    {{ `该码因"${reason}"被暂停使用` }}
  </view>
</template>

<script>
export default {
  name: 'Forbidden',

  data() {
    return {
      reason: ''
    }
  },

  created() {
    const { reason } = this.$Route.query
    this.reason = reason
  }
}
</script>

<style lang="scss" scoped>
.certificate {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
