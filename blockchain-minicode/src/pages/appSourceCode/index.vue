<template>
  <view class="content">
    <view class="nav">
      <uni-icons
        class="bars-icon"
        type="bars"
        color="#212121"
        size="20"
        @tap="openDrawer"
      ></uni-icons>
    </view>

    <view
      v-for="(item, index) in sourceCodeList"
      :key="index"
      class="comm-card"
    >
      <view class="comm-card__title">
        <text class="title">{{ item.applyCode }}</text>
        <view class="isShipped">{{ shipmentsFilter(item.applyStatus) }}</view>
      </view>

      <view class="comm-card__body">
        <view class="card-item">
          <text v-for="(key, i) in Object.keys(labelObj)" :key="i"
            >{{ labelObj[key] }}{{ item[key] }}</text
          >
        </view>
      </view>
      <view class="comm-card__btns">
        <text class="detail" @tap="onDetail(item)">查看详情</text>
      </view>
    </view>

    <view class="add-brand">
      <button class="btn" @tap="onToFormView">
        <uni-icons type="plusempty" color="#fff"></uni-icons>
        <span>申请溯源码</span>
      </button>
    </view>

    <drawer ref="drawer" @reset="onReset" @confirm="onConfirm" />
  </view>
</template>

<script>
import drawer from './components/Drawer.vue'
import { getApplyList } from '@/api/appSourceCode'
export default {
  name: 'AppSourceCode',
  components: { drawer },
  data() {
    return {
      labelObj: {
        applyQuantity: '申请数量：',
        effectQuantity: '未用数量：',
        createDate: '时间：'
      },
      applyStatus: '', // 状态 (1、待支付 2、未发货 3、已发货 4、已签收 9、已作废 )
      sourceCodeList: [
        // {
        //   applyStatus: 2,
        //   applyCode: '',
        //   applyQuantity: '',
        //   effectQuantity: '',
        //   createDate: ''
        // }
      ]
    }
  },
  computed: {},
  watch: {},

  methods: {
    async query() {
      const params = {
        enterpriseName: '',
        applyStatus: this.applyStatus, //状态 (1:待支付 2：未发货 3：已发货 4、已签收 9、已作废 )
        pageNum: 1,
        pageSize: 10
      }
      const res = await getApplyList(params)
      if (res.code === '0') {
        this.sourceCodeList = [...res.list]
      }
    },

    shipmentsFilter(shipped) {
      const shipments = {
        1: '待支付',
        2: '未发货',
        3: '已发货',
        4: '已签收',
        9: '已作废'
      }
      return shipments[shipped]
    },

    openDrawer() {
      this.$refs.drawer.showDrawer()
    },

    onReset(data) {
      this.applyStatus = data.applyStatus
      // this.query()
    },

    onConfirm(data) {
      this.applyStatus = data.applyStatus
      this.query()
    },

    onDetail(row) {
      uni.navigateTo({ url: `/pages/appSourceCode/Detail?id=${row.id}` })
    },

    onToFormView() {
      uni.navigateTo({ url: '/pages/appSourceCode/SourceCodeForm' })
    }
  },

  // 页面周期函数--监听页面加载
  onLoad() {},
  // 页面周期函数--监听页面初次渲染完成
  onReady() {},
  // 页面周期函数--监听页面显示(not-nvue)
  onShow() {
    this.sourceCodeList = []
    this.query()
  },
  // 页面周期函数--监听页面隐藏
  onHide() {},
  // 页面周期函数--监听页面卸载
  onUnload() {},
  // 页面处理函数--监听用户下拉动作
  onPullDownRefresh() {
    uni.stopPullDownRefresh()
  },
  // 页面处理函数--监听用户上拉触底
  onReachBottom() {}
  // 页面处理函数--监听页面滚动(not-nvue)
  /* onPageScroll(event) {}, */
  // 页面处理函数--用户点击右上角分享
  /* onShareAppMessage(options) {}, */
}
</script>

<style lang="scss">
.content {
  // display: flex;
  // flex-direction: column;
  // align-items: center;
  // justify-content: center;
  border-bottom: 1rpx solid #eee;
  height: calc(100vh - 121rpx);
  overflow-y: auto;

  .nav {
    width: 750rpx;
    height: 88rpx;
    background-color: #fff;
    display: flex;
    align-items: center;
    justify-content: flex-end;

    .bars-icon {
      margin-right: 32rpx;
    }
  }

  .comm-card {
    width: 718rpx;
    height: 431rpx;
    background-color: #fff;
    border-radius: 16rpx;
    border: solid 1rpx #eee;
    margin: 16rpx auto;

    .comm-card__title {
      height: 104rpx;
      display: flex;
      justify-content: space-between;
      align-items: center;
      border-bottom: 1rpx solid #eee;

      .title {
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        color: #212121;
        text-indent: 25rpx;
      }

      .isShipped {
        width: 120rpx;
        height: 48rpx;
        text-align: center;
        line-height: 48rpx;
        background-color: #00c85317;
        border-radius: 8rpx;
        font-family: PingFangSC-Regular;
        font-size: 24rpx;
        color: #00c853;
        margin-right: 24rpx;
      }
    }

    .comm-card__body {
      height: 224rpx;
      font-size: 32rpx;
      padding: 0 25rpx;
      display: flex;
      align-items: center;

      .card-item {
        display: flex;
        flex-wrap: wrap;
        align-items: center;

        text {
          display: block;
          width: 315rpx;
          margin-bottom: 25rpx;
          white-space: nowrap;
          font-family: PingFangSC-Regular;
          font-size: 32rpx;
          color: #8b8b8b;
        }
      }
    }

    .comm-card__btns {
      height: 104rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      border-top: 1rpx solid #eee;

      .detail {
        flex: 1;
        height: 100%;
        font-size: 28rpx;
        display: flex;
        justify-content: center;
        align-items: center;
        color: #212121;
      }
    }
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
