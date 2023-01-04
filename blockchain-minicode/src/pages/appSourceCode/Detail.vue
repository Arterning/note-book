<template>
  <view class="content">
    <view class="comm__container">
      <view class="comm__row b_bottom">
        <text class="comm-title">申请日期</text>
        <text class="comm-detail">{{ detailInfo.createDate }}</text>
      </view>
      <view class="comm__row">
        <text class="comm-title">申请数量</text>
        <text class="comm-detail">{{ detailInfo.applyQuantity }}</text>
      </view>
      <view class="msg"
        >溯源码编号：
        <view
          >{{ detailInfo.codePrefix + detailInfo.startCodeStr }} -
          {{ detailInfo.codePrefix + detailInfo.endCodeStr }}
        </view>
      </view>
      <view class="comm__row b_bottom">
        <text class="comm-title">收件人姓名</text>
        <text class="comm-detail">{{ detailInfo.reciverName }}</text>
      </view>
      <view class="comm__row">
        <text class="comm-title">收件人手机号</text>
        <text class="comm-detail">{{ detailInfo.reciverPhone }}</text>
      </view>
    </view>
    <view class="comm__container b_top mt_16">
      <view class="comm__row b_bottom">
        <text class="comm-title">订单状态</text>
        <text class="comm-detail">{{
          statusFilter(detailInfo.applyStatus)
        }}</text>
      </view>
      <view class="comm__row b_bottom">
        <text class="comm-title">发货日期</text>
        <text class="comm-detail">{{ detailInfo.qualityGuaPeriod || '' }}</text>
      </view>
      <view class="comm__row">
        <text class="comm-title">快递单号</text>
        <text class="comm-detail">{{ detailInfo.expressNum || '' }}</text>
      </view>
    </view>

    <view class="add-brand">
      <button class="btn" @tap="onSubmit">
        <span>确认签收</span>
      </button>
    </view>
  </view>
</template>

<script>
import { getApplyDetail } from '@/api/appSourceCode'
export default {
  name: 'Detail',
  components: {},
  data() {
    return {
      applyId: '', // 申请id
      detailInfo: {
        id: '',
        enterpriseCode: '', //企业编号
        enterpriseName: '', //企业名称
        applyCode: '', //申请编号
        applyQuantity: 0, //数量
        codePrefix: '', //溯源码前缀
        startCodeStr: '', //溯源码开始
        endCodeStr: ' ', //溯源码结束
        applyStatus: '', //状态 (1:待支付 2：未发货 3：已发货 4、已签收 9、已作废 )
        reciverName: '',
        reciverAddr: '',
        reciverPhone: '',
        creator: '',
        createDate: ''
      } // 订单详情信息
    }
  },
  computed: {},

  // 页面周期函数--监听页面加载
  onLoad(option) {
    this.applyId = option.id
  },

  onShow() {
    this.query()
  },

  methods: {
    async query() {
      const params = {
        applyId: this.applyId //申请ID
      }
      const res = await getApplyDetail(params)
      if (res.code === '0') {
        this.detailInfo = { ...res.data }
      }
    },

    statusFilter(status) {
      const statusObj = {
        1: '待支付',
        2: '未发货',
        3: '已发货',
        4: '已签收',
        9: '已作废'
      }
      return statusObj[status]
    },

    onSubmit() {}
  }
}
</script>

<style lang="scss">
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  border-top: 1rpx solid #eee;
  border-bottom: 1rpx solid #eee;
  height: calc(100vh - 121rpx);
  overflow-y: auto;

  .comm__container {
    width: 100%;
    background: #fff;
    display: flex;
    flex-direction: column;
    align-items: center;
    border-bottom: 1rpx solid #eee;

    .msg {
      width: 100vw;

      /* height: 80rpx; */
      line-height: 80rpx;
      background-color: #fafafa;
      font-family: PingFangSC-Regular;
      font-size: 22rpx;
      color: #8b8b8b;
      text-indent: 32rpx;
    }
  }

  .comm__row {
    width: calc(100vw - 66rpx);
    height: 104rpx;
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 28rpx;

    .comm-title {
      color: #8b8b8b;
    }

    .comm-detail {
      color: #212121;
    }
  }

  .b_bottom {
    border-bottom: 1rpx solid #eee;
  }

  .b_top {
    border-top: 1rpx solid #eee;
  }

  .mt_16 {
    margin-top: 16rpx;
  }

  .add-brand {
    width: 750rpx;
    height: 138rpx;
    position: fixed;
    bottom: 0;
    background-color: #fafafa;

    .btn {
      border-radius: 44rpx;
      font-size: 32rpx;
      margin: 25rpx 40rpx;
      background-color: #00c853;

      span {
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        color: #fff;
        margin-left: 10rpx;
      }
    }
  }
}
</style>
