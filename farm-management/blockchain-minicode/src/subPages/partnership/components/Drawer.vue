<template>
  <view>
    <uni-drawer
      ref="showRight"
      mode="right"
      :width="300"
      :mask-click="true"
      class="drawer"
    >
      <view class="drawer">
        <view class="title">合作状态</view>
        <view class="section">
          <view
            v-for="(tag, i) in dictList"
            :key="i"
            class="tag-wrapper"
            @click="onSelect(i)"
            ><span :class="{ tag: true, selected: tag.selected }">{{
              tag.text
            }}</span></view
          >
        </view>

        <view class="footer">
          <button type="plain" class="btn" @click="onReset">重置</button>
          <button type="primary" class="btn" @click="onConfirm">确定</button>
        </view>
      </view>
    </uni-drawer>
  </view>
</template>

<script>
export default {
  name: 'Drawer',

  data() {
    return {
      dict: '',
      dictList: [],
      shipType: 0,
      shipmentsList: []
    }
  },

  computed: {
    cooStatusMap() {
      return this.$store.state.user.dictMap.srv_coo_status
    }
  },

  async created() {
    Object.keys(this.cooStatusMap).forEach(key => {
      this.dictList.push({
        code: key,
        text: this.cooStatusMap[key],
        selected: false
      })
    })
  },

  methods: {
    onSelect(i) {
      this.dictList.map(el => {
        el.selected = false
      })
      this.dictList[i].selected = !this.dictList[i].selected
    },
    showDrawer() {
      this.$refs.showRight.open()
    },

    closeDrawer() {
      this.$refs.showRight.close()
    },

    onShipChange(e, key) {
      this[key] = e.detail.value
    },

    onReset() {
      this.dictList.map(el => {
        el.selected = false
      })
      this.closeDrawer()
      this.dict = ''

      this.$emit('reset', null)
    },

    onConfirm() {
      this.closeDrawer()
      this.dictList.map(el => {
        if (el.selected) {
          this.dict = el.code
        }
      })
      const drawerInfo = {
        cooStatus: this.dict
      }
      this.$emit('confirm', drawerInfo)
    }
  }
}
</script>

<style lang="scss" scoped>
// 抽屉
.drawer {
  box-sizing: border-box;
  height: 100%;
  padding-bottom: 144rpx;

  .title {
    padding-left: 24rpx;
    height: 80rpx;
    line-height: 80rpx;
    font-size: 24rpx;
    color: #8b8b8b;
  }

  .section {
    height: calc(100% - 80rpx);
    overflow-y: scroll;
    margin: 0 -16rpx;
    padding: 0 24rpx 16rpx;
    display: flex;
    flex-wrap: wrap;
    align-content: flex-start;

    .tag-wrapper {
      width: 100%;
      height: 64rpx;
      padding: 0 16rpx;

      .tag {
        box-sizing: border-box;
        display: block;
        height: 48rpx;
        padding: 0 16rpx;
        line-height: 46rpx;
        text-align: center;
        font-size: 24rpx;
        border-radius: 24rpx;
        border: solid 1rpx #e0e0e0;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
      }

      .selected {
        // color: $theme-color;
        // border-color: $theme-color;
        color: #70b913;
        border-color: #70b913;
        background-color: rgba(112, 185, 19, 0.1);
      }
    }
  }

  .section::-webkit-scrollbar {
    display: none;
  }

  .footer {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    padding: 32rpx 24rpx;
    background-color: $uni-bg-color-page;
    display: flex;
    justify-content: space-between;

    .btn {
      margin: 0;
      width: 168rpx;
      height: 80rpx;
      font-size: 32rpx;
      background-color: $uni-bg-color-btn;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 40rpx;
    }

    .btn:first-of-type {
      color: $theme-color;
      background-color: #fff;
      border: 1rpx solid $theme-color;
    }

    .btn::after {
      display: none;
    }
  }
}
</style>
