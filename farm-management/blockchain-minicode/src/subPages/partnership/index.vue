<template>
  <view class="content">
    <search-bar
      placeholder="企业名称搜索"
      :shadow="scrollTop > 8"
      @input="search"
    >
      <uni-icons
        style="margin-left: 32rpx;"
        type="bars"
        size="20"
        color="#212121"
        class="icon"
        @click="openDrawer"
      ></uni-icons>
    </search-bar>

    <view style="margin-top: 90rpx;">
      <view
        v-for="(item, index) in sourceCodeList"
        :key="index"
        class="comm-card"
      >
        <view class="comm-card__title">
          <text class="title">{{ item.cooEnterpriseName }}</text>
          <view v-if="item.cooStatus === '1'" class="isShipped">{{
            item.cooStatusName
          }}</view>
          <view v-if="item.cooStatus === '0'" class="isShipped2">{{
            item.cooStatusName
          }}</view>
          <view v-if="item.cooStatus === '2'" class="isShipped3">{{
            item.cooStatusName
          }}</view>
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
          <text
            v-if="item.cooStatus === '1'"
            style="color: red;"
            class="detail"
            @tap="cancelCooPar(item)"
            >取消</text
          >
        </view>
      </view>
      <uni-load-more :status="status" @clickLoadMore="query"></uni-load-more>
    </view>
    <bottom-button text="添加合作米厂" @click="onToFormView"></bottom-button>

    <drawer ref="drawer" @reset="onReset" @confirm="onConfirm" />
    <uni-popup ref="popup" type="dialog">
      <uni-popup-dialog
        mode="base"
        title="是否取消合作？"
        :before-close="true"
        @close="close"
        @confirm="confirm"
      ></uni-popup-dialog>
    </uni-popup>
  </view>
</template>

<script>
import drawer from './components/Drawer.vue'
import SearchBar from '@/components/search-bar'
import BottomButton from '@/components/bottom-button'

import { getCooParList, cancelCooPar } from '@/api/partnership'
import { stringifyDate } from '../../utils/tool'

export default {
  name: 'Partnership',
  components: { drawer, SearchBar, BottomButton },
  data() {
    return {
      status: 'more', // 列表加载状态
      enterpriseId: '',
      cooEnterpriseNameLike: '',
      types: [],
      cooStatus: '',
      labelObj: {
        typeName: '企业类型：',
        createDate: '日期：'
      },
      sourceCodeList: []
    }
  },
  onShow() {
    this.sourceCodeList = []
    this.query()
  },

  onReachBottom() {
    console.log('到底了')
    this.query() // 加载下一页
  },

  onPullDownRefresh() {
    console.log('下拉刷新')
    this.query()
  },

  methods: {
    search(name) {
      this.cooEnterpriseNameLike = name
      this.query()
    },
    async query() {
      this.status = 'loading'
      const params = {
        cooEnterpriseNameLike: this.cooEnterpriseNameLike,
        cooStatus: this.cooStatus, //合作状态
        pageSize: '0',
        pageNum: '1'
      }
      const res = await getCooParList(params)
      if (res.code === '0') {
        res.list.map(el => {
          el.createDate = stringifyDate(el.createDate)
        })
        this.sourceCodeList = [...res.list]
        // uni.stopPullDownRefresh()
        this.status = 'nomore'
      }
    },
    cancelCooPar(row) {
      this.enterpriseId = row.id
      this.$refs.popup.open()
    },

    close() {
      this.$refs.popup.close()
    },

    async confirm() {
      this.$refs.popup.close()
      const res = await cancelCooPar({
        id: this.enterpriseId
      })
      if (res.code === '0') {
        this.query()
        // uni.navigateTo({ url: `/subPages/partnership/index` })
      }
      //   this.handleDel(this.activeIndex, this.employeeId)
    },
    openDrawer() {
      this.$refs.drawer.showDrawer()
    },

    onReset() {
      this.types = []
      this.cooStatus = ''
      // this.applyStatus = data.applyStatus
      this.query()
    },

    onConfirm(data) {
      console.log(data)

      this.cooStatus = data.cooStatus

      this.query()
    },

    /**
     * @description: 跳转至合作伙伴详情页
     * @param {Object} row
     * @author: Hemingway
     */
    onDetail(row) {
      this.$Router.push({
        path: '/partnershipDetail',
        query: { enterpriseId: row.cooEnterpriseId, isApprove: false }
      })
    },

    /**
     * @description: 跳转至待选合作伙伴列表页
     * @author: Hemingway
     */
    onToFormView() {
      this.$Router.push({ name: 'partnershipList' })
    }
  }
}
</script>

<style lang="scss">
.content {
  .comm-card {
    width: 718rpx;
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

      .isShipped2 {
        width: 120rpx;
        height: 48rpx;
        text-align: center;
        line-height: 48rpx;
        background-color: #188fff0c;
        border-radius: 8rpx;
        font-family: PingFangSC-Regular;
        font-size: 24rpx;
        color: #1890ff;
        margin-right: 24rpx;
      }

      .isShipped3 {
        width: 120rpx;
        height: 48rpx;
        text-align: center;
        line-height: 48rpx;
        background-color: #8b8b8b0e;
        border-radius: 8rpx;
        font-family: PingFangSC-Regular;
        font-size: 24rpx;
        color: #8b8b8b;
        margin-right: 24rpx;
      }
    }

    .comm-card__body {
      height: 112rpx;
      font-size: 32rpx;
      padding: 20rpx 25rpx;
      display: flex;
      align-items: center;

      .card-item {
        display: flex;
        flex-wrap: wrap;
        align-items: center;

        text {
          display: block;
          margin-right: 40rpx;
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

  ::v-deep .bottom-button {
    width: 100%;
    box-sizing: border-box;
    position: fixed;
    bottom: 0;
    background: rgba(255, 255, 255, 0.8);
    box-shadow: 0 -24rpx 12rpx 1rpx rgba(0, 0, 0, 0.05);
  }
}
</style>
