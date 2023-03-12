<!--
 * @Description  : 区块链证书
 * @Autor        : Hemingway
 * @Date         : 2021-10-21 10:26:40
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-06-13 13:45:28
 * @FilePath     : \blockchain-h5\src\pages\certificate\index.vue
-->
<template>
  <view class="certificate">
    <view class="container">
      <view class="topic">{{ name }}</view>
      <view class="group">
        <view class="title">区块链ID：</view>
        <view class="text">{{ tx_id }}</view>
      </view>
      <view class="group">
        <view class="title">唯一编码：</view>
        <view class="text">{{ block_number }}</view>
      </view>
      <view class="bottom-wrapper group">
        <view class="column">
          <view class="title">上链时间：</view>
          <view class="text">{{ tx_create_time }}</view>
        </view>
        <view class="column">
          <view class="title">上传者身份：</view>
          <view class="text">{{ creator }}</view>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'Certificate',

  data() {
    return {
      name: '',
      creator: '',
      tx_create_time: '',
      tx_id: '',
      block_number: ''
    }
  },

  created() {
    const {
      cername,
      tx_id,
      block_number,
      tx_create_time,
      creator
    } = this.$Route.query

    this.name = cername
    this.tx_id = tx_id
    this.block_number = block_number
    this.tx_create_time = tx_create_time
    this.creator = creator
  }
}
</script>

<style lang="scss" scoped>
.certificate {
  width: 100%;
  height: 135.73vw;
  background: url(../../static/img/certificate.png) no-repeat;
  background-size: 100% 100%;
  position: relative;

  .container {
    position: absolute;
    top: 362rpx;
    left: 96rpx;
    right: 96rpx;
    font-size: 22rpx;
    color: #212121;

    .topic {
      text-align: center;
      font-size: 28rpx;
      font-weight: bold;
      color: #354380;
    }

    .group {
      margin-top: 36rpx;

      .title {
        color: #354380;
      }

      .text {
        margin-top: 20rpx;
        word-break: break-all;
      }
    }

    .bottom-wrapper {
      display: flex;
      flex-direction: row;

      .column {
        flex: 1;
      }
    }
  }
}
</style>
