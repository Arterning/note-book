<!--
 * @Author: guoxi
 * @Date: 2021-05-25 14:05:47
 * @LastEditTime : 2021-10-21 14:38:36
 * @LastEditors  : Please set LastEditors
 * @Description: 员工列表页
 * @FilePath     : \blockchain-minicode\src\pages\employee\index.vue
-->

<template>
  <view class="content">
    <search-bar
      placeholder="员工姓名搜索"
      :shadow="scrollTop > 8"
      @input="search"
    >
      <uni-icons
        style="margin-left: 32rpx;"
        type="bars"
        size="20"
        color="#212121"
        class="icon"
        @click="open"
      ></uni-icons>
    </search-bar>

    <view style="margin-top: 90rpx;">
      <uni-list :border="false">
        <card v-for="(item, i) in employeeList" :key="i" :data="item">
          <template #default>
            <tag
              :color="item.statusColor[0]"
              :background-color="item.statusColor[1]"
            ></tag>
          </template>
          <template #footer>
            <view class="comm-card__btns">
              <text class="edit" @tap="onAddEmployee('edit', item.sourceData)"
                >编辑</text
              >
              <text class="delete" @tap="onOpenDialog(i, item.id)">删除</text>
            </view>
          </template>
        </card>
      </uni-list>
      <uni-load-more :status="status" @clickLoadMore="query"></uni-load-more>
      <view style="height: 100px;"></view>
    </view>

    <bottom-button
      text="添加员工"
      @click="onAddEmployee('add')"
    ></bottom-button>

    <uni-popup ref="popup" type="dialog">
      <uni-popup-dialog
        type="error"
        mode="base"
        title="确认删除？"
        :before-close="true"
        @close="close"
        @confirm="onConfirm"
      ></uni-popup-dialog>
    </uni-popup>

    <uni-drawer ref="showRight" mode="right" :mask-click="true">
      <view class="drawer">
        <view class="title">账号类型</view>
        <view class="section">
          <view
            v-for="(tag, i) in personTypes"
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
          <button type="primary" class="btn" @click="onConfirmType">
            确定
          </button>
        </view>
      </view>
    </uni-drawer>
  </view>
</template>

<script>
import { getUserList, delUsers } from '@/api/employee'
import SearchBar from '@/components/search-bar'
import Card from '@/components/card'
import BottomButton from '@/components/bottom-button'

export default {
  name: 'Employee',
  components: { SearchBar, Card, BottomButton },
  data() {
    return {
      status: 'more', // 列表加载状态
      personType: [],

      personTypes: [],
      employeeId: 0,
      activeIndex: -1,
      employeeList: [],
      userNameLike: '',
      index: undefined
    }
  },
  computed: {
    accountMap() {
      return this.$store.state.user.dictMap.account_type
    }
  },

  onLoad() {},
  onShow() {
    this.employeeList = []
    this.status = 'more'
    if (this.personTypes.length === 0) {
      Object.keys(this.accountMap).forEach(key => {
        this.personTypes.push({
          code: key,
          text: this.accountMap[key],
          selected: false
        })
      })
    }
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
      this.userNameLike = name
      this.query()
    },
    async query() {
      if (this.status === 'more') {
        this.status = 'loading'
        const params = {
          userNameLike: this.userNameLike,
          types: this.personType,
          pageSize: 0,
          pageNum: 1
        }
        let res = await getUserList(params)
        if (res.code === '0') {
          this.employeeList = res.list.map(el => {
            return {
              id: el.id,
              sourceData: el,
              title: `${el.name}`,
              items: [
                { text: `账号类型：${el.roleName}` },
                { text: `手机号码：${el.phone}` },
                { text: `创建时间：${this.filterDate(el.createDate)}` }
              ]
            }
          })
          uni.stopPullDownRefresh()
          this.status = 'more'
        }
      }
    },
    onSelect(i) {
      this.personTypes[i].selected = !this.personTypes[i].selected
    },

    filterDate(time) {
      const timeArr = time ? time.split(' ') : []
      return timeArr === [] ? '' : timeArr[0]
    },

    onOpenDialog(index, id) {
      this.employeeId = id
      this.activeIndex = index

      this.$refs.popup.open()
    },

    onAddEmployee(type, row) {
      let employeeData = {
        type
      }
      if (type === 'edit') {
        const { id, name, phone, roleName } = row
        const filterItem = this.personTypes.find(
          per => per['text'] === roleName
        )
        employeeData = {
          ...employeeData,
          id,
          name,
          phone,
          roleName,
          personId: filterItem && filterItem.code
        }
        uni.navigateTo({
          url: `/pages/employee/Detail?employeeData=${encodeURIComponent(
            JSON.stringify(employeeData)
          )}`
        })
        return
      }
      uni.navigateTo({
        url: `/pages/employee/Detail?employeeData=${encodeURIComponent(
          JSON.stringify(employeeData)
        )}`
      })
    },

    close() {
      this.$refs.popup.close()
    },

    onConfirm() {
      this.$refs.popup.close()
      this.handleDel(this.activeIndex, this.employeeId)
    },

    async handleDel(index, id) {
      let res = await delUsers(id)
      if (res.code === '0') {
        this.employeeList.splice(index, 1)
        uni.showToast({
          title: '删除成功',
          icon: 'success'
        })
        this.query()
        return
      }
      uni.showToast({
        title: '删除失败',
        icon: 'none'
      })
    },

    onChange(e) {
      this.personType = e.detail.value
    },

    open() {
      // if (this.personTypes.length === 0) {
      //   // 初始化

      // }
      this.$refs.showRight.open()
    },

    onConfirmType() {
      this.personTypes.map(el => {
        if (el.selected) {
          this.personType.push(el.code)
        }
      })

      this.$refs.showRight.close()
      this.query()
    },

    onReset() {
      this.personType = []
      this.personTypes.map(el => {
        el.selected = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.content {
  // 列表
  ::v-deep {
    .uni-list {
      background-color: $uni-bg-color-page;
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

  .comm-card__btns {
    height: 68rpx;
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
}
</style>
