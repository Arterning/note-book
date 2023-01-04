<!--
 * @Author: guoxi
 * @Date: 2021-05-24 14:06:36
 * @LastEditTime : 2022-07-08 18:38:37
 * @LastEditors  : Your Name
 * @Description: 品牌详情管理页
 * @FilePath     : \blockchain-minicode\src\pages\commodity\BrandGoodsManagement.vue
-->

<template>
  <view class="content">
    <search-bar
      placeholder="品牌名称搜索"
      :shadow="scrollTop > 8"
      @input="search"
    >
    </search-bar>

    <view style="margin-top: 90rpx; margin-bottom: 190rpx;">
      <uni-list :border="false">
        <view v-for="(item, index) in brandList" :key="index" class="comm-card">
          <view class="comm-card__title">
            <text class="title">{{ item.brandName }}</text>
            <view class="edit" @tap="onAddBrand('edit', item)">
              <text class="edit-text">编辑</text>
              <uni-icons type="compose" color="#8b8b8b"></uni-icons>
            </view>
          </view>
          <view class="comm-card__middle">
            <view class="comm-card__time">创建时间：{{ item.createDate }}</view>
            <view class="comm-card__number"
              >商品注册证号：{{ item.brandRegCard }}</view
            >
          </view>
          <view class="comm-card__btns">
            <text class="edit" @tap="onToCommDetail(item.id, item.brandName)"
              >商品详情</text
            >
            <text class="delete" @click="onShowDialog(item.id, index)"
              >删除</text
            >
          </view>
        </view>
      </uni-list>
    </view>
    <bottom-button text="添加品牌" @click="onAddBrand('add')"></bottom-button>
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
import { getBrandList, delBrand } from '@/api/commodity'
import SearchBar from '@/components/search-bar'
import BottomButton from '@/components/bottom-button'

export default {
  name: 'BrandGoodsManagement',
  components: { SearchBar, BottomButton },
  data() {
    return {
      status: 'more', // 列表加载状态
      currentRowId: null,
      activeIndex: -1,
      brandList: [],
      brandNameLike: ''
    }
  },

  onLoad() {},

  onReachBottom() {
    console.log('到底了')
    this.query() // 加载下一页
  },

  onPullDownRefresh() {
    console.log('下拉刷新')
    this.query()
  },

  onShow() {
    this.brandList = []
    this.query()
  },

  methods: {
    search(name) {
      this.brandNameLike = name
      this.query()
    },
    async query() {
      if (this.status === 'more') {
        this.status = 'loading'
        const params = {
          brandNameLike: this.brandNameLike,
          pageSize: '0',
          pageNum: '1'
        }
        let res = await getBrandList(params)
        if (res.code === '0') {
          this.brandList = res.list
          uni.stopPullDownRefresh()
          this.status = 'more'
        }
      }
    },
    onToCommDetail(id, name) {
      uni.navigateTo({
        url: `/pages/commodity/CommManagemant?id=${id}&name=${name}`
      })
    },

    // 添加品牌
    onAddBrand(type, row) {
      if (type === 'edit') {
        const brandInfo = {
          brandName: row.brandName,
          brandRegCard: row.brandRegCard,
          id: row.id
        }
        uni.navigateTo({
          url: `/pages/commodity/AddBrand?type=${type}&brandInfo=${JSON.stringify(
            brandInfo
          )}`
        })
        return
      }
      uni.navigateTo({ url: `/pages/commodity/AddBrand?type=${type}` })
    },

    onClose() {
      this.$refs.popup.close()
    },

    onConfirm() {
      this.$refs.popup.close()
      this.onDelBrand(this.currentRowId, this.activeIndex)
    },

    onShowDialog(id, index) {
      this.$refs.popup.open()
      this.activeIndex = index
      this.currentRowId = id
    },

    async onDelBrand(id, index) {
      let res = await delBrand(id)
      if (res.code === '0') {
        this.brandList.splice(index, 1)
        uni.showToast({
          title: '删除成功',
          icon: 'none'
        })
        this.query()
      } else {
        uni.showToast({
          title: '删除失败',
          icon: 'none'
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  .comm-card {
    width: 718rpx;
    height: 376rpx;
    background-color: #fff;
    border-radius: 16rpx;
    border: solid 1rpx #eee;
    margin-top: 16rpx;

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

      .edit {
        display: flex;
        align-items: center;

        .edit-text {
          font-size: 24rpx;
          color: #8b8b8b;
          margin-right: 10rpx;
        }
      }
    }

    .comm-card__middle {
      font-size: 32rpx;
      color: #8b8b8b;
      padding: 40rpx 25rpx;
    }

    .comm-card__time {
      margin-bottom: 6rpx;
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

  ::v-deep .bottom-button {
    width: 100%;
    box-sizing: border-box;
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(255, 255, 255, 0.8);
    box-shadow: 0 -24rpx 12rpx 1rpx rgba(0, 0, 0, 0.05);
  }
}
</style>
