<!--
 * @Author: guoxi
 * @Date: 2021-05-24 14:06:36
 * @LastEditTime : 2022-07-08 18:12:24
 * @LastEditors  : Your Name
 * @Description: 商品列表页
 * @FilePath     : \blockchain-minicode\src\pages\commodity\CommManagemant.vue
-->

<template>
  <view class="content">
    <view class="search-box">
      <view class="search">
        <uni-easyinput
          v-model="commodityNameLike"
          class="search-input"
          :input-border="false"
          prefix-icon="search"
          placeholder="商品名称搜索"
          placeholder-style="font-size: 28rpx; color: #8b8b8b;"
          @iconClick="query()"
          @input="query()"
        ></uni-easyinput>
      </view>
    </view>
    <view style="margin-top: 90rpx;">
      <uni-list :border="false">
        <view
          v-for="(item, index) in commityList"
          :key="index"
          class="comm-card"
        >
          <view class="comm-card__title" @tap="onToCommDetail(item.id)">
            <text class="title">{{ item.commodityName }}</text>
            <uni-icons type="forward" color="#8b8b8b"></uni-icons>
          </view>
          <view class="comm-card__time">创建时间：{{ item.createDate }}</view>
          <view class="comm-card__btns">
            <text class="edit" @click="onAddCommodity('edit', item.id)"
              >编辑</text
            >
            <text class="delete" @click="onShowDialog(item.id, index)"
              >删除</text
            >
          </view>
        </view>
      </uni-list>
    </view>
    <view class="add-commodity">
      <button type="primary" class="btn" @tap="onAddCommodity('add')">
        <uni-icons type="plusempty" color="#fff"></uni-icons>
        <span>添加商品</span>
      </button>
    </view>
    <uni-popup ref="popup" type="dialog">
      <uni-popup-dialog
        type="error"
        mode="base"
        title="确认删除？"
        :before-close="true"
        @close="onClose"
        @confirm="onConfirm"
      ></uni-popup-dialog>
    </uni-popup>
  </view>
</template>

<script>
import { getCommodityList, delCommodity } from '@/api/commodity'
export default {
  name: 'CommManagemant',
  data() {
    return {
      status: 'more', // 列表加载状态
      currentRowId: null,
      activeIndex: -1,
      brandId: null,
      brandName: '',
      commityList: [],
      commodityNameLike: ''
    }
  },

  onLoad(options) {
    console.log('kankan', options)
    this.brandId = options.id
    this.brandName = decodeURI(options.name)
  },

  onReachBottom() {
    console.log('到底了')
    this.query() // 加载下一页
  },

  onPullDownRefresh() {
    console.log('下拉刷新')
    this.query()
  },

  onShow() {
    this.commityList = []
    this.query()
  },

  methods: {
    async query() {
      if (this.status === 'more') {
        this.status = 'loading'
        const params = {
          brandId: this.brandId,
          commodityNameLike: this.commodityNameLike,
          pageSize: '0',
          pageNum: '1'
        }
        let res = await getCommodityList(params)
        if (res.code === '0') {
          this.commityList = res.list
          uni.stopPullDownRefresh()
          this.status = 'more'
        }
      }
    },

    onToCommDetail(id) {
      uni.navigateTo({
        url: `/pages/commodity/CommodityDetail?id=${id}&name=${this.brandName}`
      })
    },

    onAddCommodity(type, commodityId) {
      if (type === 'edit') {
        uni.navigateTo({
          url: `/pages/commodity/AddCommodity?type=${type}&brandId=${this.brandId}&commodityId=${commodityId}&name=${this.brandName}`
        })
        return
      }
      uni.navigateTo({
        url: `/pages/commodity/AddCommodity?type=${type}&brandId=${this.brandId}&name=${this.brandName}`
      })
    },

    onClose() {
      this.$refs.popup.close()
    },

    onConfirm() {
      this.$refs.popup.close()
      this.handleDel(this.currentRowId, this.activeIndex)
    },

    onShowDialog(id, index) {
      this.$refs.popup.open()
      this.activeIndex = index
      this.currentRowId = id
    },

    async handleDel(id, index) {
      let res = await delCommodity(id)
      if (res.code === '0') {
        this.commityList.splice(index, 1)
        uni.showToast({
          title: '删除成功',
          icon: 'none'
        })
        this.query()
        return
      }
      uni.showToast({
        title: '删除失败',
        icon: 'none'
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.content {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 220rpx);
  margin: 0 auto;
  overflow: auto;

  .search-box {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1;
  }

  .search {
    padding: 12rpx 32rpx;
    background-color: #fff;
    display: flex;
    align-items: center;

    .search-input {
      display: inline-block;
      width: 100%;
      background-color: #f5f5f5;
      border-radius: 32px;
      border: none;
    }

    .input-placeholder {
      font-size: 28rpx;
      color: #8b8b8b;
    }

    .bars-icon {
      display: inline-block;
      vertical-align: middle;
      margin-left: 16rpx;
    }
  }

  .comm-card {
    width: 718rpx;
    height: 280rpx;
    background-color: #fff;
    border-radius: 16rpx;
    border: solid 1rpx #eee;
    margin: 16rpx auto 0;

    .comm-card__title {
      height: 104rpx;
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 25rpx;

      .title {
        font-size: 32rpx;
        color: #212121;
      }
    }

    .comm-card__time {
      height: 72rpx;
      font-size: 32rpx;
      padding-left: 25rpx;
      color: #8b8b8b;
    }

    .comm-card__btns {
      height: 104rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      border-top: 1rpx solid #eee;

      .edit,
      .delete {
        flex: 1;
        height: 100%;
        font-size: 28rpx;
        display: flex;
        justify-content: center;
        align-items: center;
      }

      .edit {
        color: #212121;
      }

      .delete {
        border-left: 1rpx solid #eee;
        color: #e53935;
      }
    }
  }

  .add-commodity {
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
        margin-left: 10rpx;
      }
    }
  }
}
</style>
