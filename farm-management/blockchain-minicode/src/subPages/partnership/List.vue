<!--
 * @Author: guoxi
 * @Date: 2021-05-25 14:05:47
 * @LastEditTime : 2022-07-06 10:30:49
 * @LastEditors  : Your Name
 * @Description: 待选合作伙伴列表页
 * @FilePath     : \blockchain-minicode\src\subPages\partnership\List.vue
-->

<template>
  <view class="content">
    <view class="search">
      <uni-easyinput
        v-model="cooEnterpriseNameLike"
        prefix-icon="search"
        placeholder="企业名称搜索"
        placeholder-style="font-size: 28rpx; color: #8b8b8b; font-family: PingFangSC-Regular;"
        class="search-input"
        :input-border="false"
        @iconClick="query()"
        @input="query()"
      ></uni-easyinput>
      <!-- <uni-icons
        class="bars-icon"
        type="bars"
        size="30"
        @click="open"
      ></uni-icons> -->
      <uni-icons
        class="bars-icon"
        type="location-filled"
        size="30"
        @click="selectCity"
      ></uni-icons>
      <span class="city-name">{{ cityName }}</span>
    </view>

    <view v-for="(item, index) in choiceList" :key="index" class="comm-card">
      <view class="comm-card__title">
        <text class="title">{{ item.cooEnterpriseName }}</text>
      </view>

      <view class="comm-card__body">
        <view class="card-body">
          <p>企业类型：{{ item.typeName }}</p>
          <p>联系电话：{{ item.officePhone }}</p>
          <p>企业地址：{{ item.address }}</p>
        </view>
      </view>
      <view class="comm-card__btns">
        <text class="edit" @tap="onDetail(item)">查看详情</text>
        <text class="edit" @tap="onOpenDialog(item)">选择</text>
      </view>
    </view>

    <!-- <view class="add-commodity">
      <button class="btn" @tap="onAddEmployee('add')">
        <uni-icons type="plusempty" color="#fff"></uni-icons>
        <span>添加员工</span>
      </button>
    </view> -->
    <uni-popup ref="popup" type="dialog">
      <uni-popup-dialog
        type="info"
        mode="base"
        :title="popTitle"
        :before-close="true"
        @close="close"
        @confirm="onConfirm"
      ></uni-popup-dialog>
    </uni-popup>

    <!-- <uni-drawer
      ref="showRight"
      :mask-click="true"
      mode="right"
      :width="300"
      class="drawer"
    >
      <view class="drawer-content">
        <view class="picker-type">
          <view class="title">企业类型</view>
          <view class="picker-right">
            <picker
              mode="selector"
              :value="personType"
              :range="personTypes"
              range-key="name"
              @change="onChange"
            >
              <view class="time-range">
                {{ personTypes[personType]['name'] }}
              </view>
            </picker>
            <uni-icons type="forward" size="16"></uni-icons>
          </view>
        </view>
      </view>
      <view class="drawer-fotter">
        <button type="primary" class="btn reset" @tap="onReset">重置</button>
        <button type="primary" class="btn confirm" @tap="onConfirmType">
          确定
        </button>
      </view>
    </uni-drawer> -->
  </view>
</template>

<script>
import { getChoiceList, choiceCooPar } from '@/api/partnership'
export default {
  name: 'ChoiceList',
  data() {
    return {
      popTitle: '',
      cityName: '',
      personType: 0,
      personId: '',
      personTypes: [],
      cooEnterpriseId: 0,
      cooEnterpriseName: '',
      choiceList: [],
      cooEnterpriseNameLike: ''
    }
  },
  computed: {
    enterpriseMap() {
      return this.$store.state.user.dictMap.enterprise_type
    }
  },

  onLoad(op) {
    this.cityName = op.cityName
  },

  onShow() {
    this.choiceList = []
    this.query()
  },

  methods: {
    async query() {
      // const data = await getEnterTypeList()
      this.personTypes = [{ id: 0, name: '请选择' }]

      Object.keys(this.enterpriseMap).forEach(key => {
        this.personTypes.push({ id: key, name: this.enterpriseMap[key + ''] })
      })
      // const type = this.personTypes[this.personType].id

      const params = {
        types: [6], //企业类型多选框
        // types: type === 0 ? [] : [type], //企业类型多选框
        cooEnterpriseNameLike: this.cooEnterpriseNameLike, //合作伙伴企业名称模糊查询
        addressLike: this.cityName, //企业地址
        pageSize: '0',
        pageNum: '1'
      }
      let res = await getChoiceList(params)
      if (res.code === '0') {
        this.choiceList = res.list
        return
      }
    },

    filterDate(time) {
      const timeArr = time ? time.split(' ') : []
      return timeArr === [] ? '' : timeArr[0]
    },

    onOpenDialog(item) {
      this.cooEnterpriseId = item.cooEnterpriseId
      this.cooEnterpriseName = item.cooEnterpriseName
      this.popTitle = `确定${this.cooEnterpriseName}作为您的合作米厂！`
      this.$refs.popup.open()
    },

    close() {
      this.$refs.popup.close()
    },

    async onConfirm() {
      this.$refs.popup.close()
      const res = await choiceCooPar({
        cooEnterpriseId: this.cooEnterpriseId,
        cooEnterpriseName: this.cooEnterpriseName
      })
      if (res.code === '0') {
        uni.navigateTo({ url: `/subPages/partnership/index` })
      }
      //   this.handleDel(this.activeIndex, this.employeeId)
    },

    onChange(e) {
      this.personType = e.detail.value
    },

    open() {
      this.$refs.showRight.open()
    },

    onConfirmType() {
      const filterItem = this.personTypes.filter(
        per => per.name === this.personTypes[this.personType]['name']
      )[0]
      this.personId = filterItem['id']
      this.$refs.showRight.close()
      this.query()
    },

    onReset() {
      this.personType = 0
      this.$refs.showRight.close()
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

    selectCity() {
      uni.navigateTo({ url: `/subPages/partnership/City` })
    }
  }
}
</script>

<style lang="scss" scoped>
.content {
  display: flex;
  flex-direction: column;

  /* height: calc(100vh - 120rpx); */
  overflow: auto;

  .search {
    padding: 12rpx 0 12rpx 32rpx;
    background-color: #fff;

    .search-input {
      display: inline-block;
      width: 536rpx;
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

      .isAdmin {
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
      padding-left: 25rpx;

      .card-body {
        padding-top: 40rpx;
      }

      p {
        margin: 0;
        margin-bottom: 25rpx;
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        color: #8b8b8b;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
      }
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
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        color: #fff;
        margin-left: 10rpx;
      }
    }
  }

  .city-name {
    width: 115rpx;
    display: inline-block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    vertical-align: sub;
  }

  ::v-deep .uni-popup__info {
    padding: 0 40rpx;
  }
}
</style>
