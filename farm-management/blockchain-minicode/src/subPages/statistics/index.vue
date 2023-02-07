<!--
 * @Author: guoxi
 * @Date: 2021-05-24 09:55:30
 * @LastEditTime : 2021-11-24 13:49:29
 * @LastEditors  : Please set LastEditors
 * @Description: 
 * @FilePath     : \blockchain-minicode\src\subPages\statistics\index.vue
-->

<template>
  <view class="statistics">
    <view class="statistics-tab">
      <view
        v-for="(item, index) in tabs"
        :key="index"
        class="tab"
        :class="{ active: index === activeIndex }"
        @click="changeTab(index)"
        ><text class="statistics-text">{{ item }}</text></view
      >
    </view>
    <view v-if="activeIndex === 0">
      <view class="sta-title">截至昨日</view>
      <view class="sta-body">
        <view class="sta-body-item">
          <view class="body-item-value">{{ entirety.boundTracing }}</view>
          <view class="body-item-key">溯源码累计</view>
        </view>
        <view class="sta-body-item">
          <view class="body-item-value">{{ entirety.pv }}</view>
          <view class="body-item-key">访问次数</view>
        </view>
        <view class="sta-body-item">
          <view class="body-item-value">{{ entirety.uv }}</view>
          <view class="body-item-key">访问人数</view>
        </view>
      </view>
    </view>
    <view v-else>
      <view class="sta-title2"
        >数据统计
        <uni-icons
          class="bars-icon"
          type="bars"
          size="30"
          @tap="showDrawer"
        ></uni-icons>
      </view>
      <view class="uv">
        <view class="sta-body-item">
          <view class="body-item-value">{{ multidimen.pv }}</view>
          <view class="body-item-key">访问次数</view>
        </view>
        <view v-if="!isopen" class="charts">
          <f2 id="pv" :init="initChartPv" />
        </view>
      </view>
      <view class="pv">
        <view class="sta-body-item">
          <view class="body-item-value">{{ multidimen.uv }}</view>
          <view class="body-item-key">访问人数</view>
        </view>
        <view v-if="!isopen" class="charts">
          <f2 id="uv" :init="initChartUv" />
        </view>
      </view>
    </view>
    <uni-drawer
      ref="showRight"
      mode="right"
      :width="300"
      :mask-click="true"
      class="drawer"
      @change="changeDrawer"
    >
      <view class="drawer-content">
        <view class="title">时间</view>
        <view class="time-container">
          <picker
            mode="date"
            :value="handleDateStart"
            :start="startDate"
            :end="endDate"
            @change="onDateChange($event, 'Start')"
          >
            <view class="time-range">
              <view class="year">{{
                filterDate(handleDateStart)['year']
              }}</view>
              <view class="month">{{
                filterDate(handleDateStart)['month']
              }}</view>
              <view class="day">{{ filterDate(handleDateStart)['day'] }}</view>
            </view>
          </picker>
          <view class="b"></view>
          <picker
            mode="date"
            :value="handleDateEnd"
            :start="startDate"
            :end="endDate"
            @change="onDateChange($event, 'End')"
          >
            <view class="time-range">
              <view class="year">{{ filterDate(handleDateEnd)['year'] }}</view>
              <view class="month">{{
                filterDate(handleDateEnd)['month']
              }}</view>
              <view class="day">{{ filterDate(handleDateEnd)['day'] }}</view>
            </view>
          </picker>
        </view>
        <view class="breeds">
          <view class="breeds-title">商品名称</view>
          <view class="breeds-picker">
            <picker
              mode="selector"
              :value="commodityType"
              :range="commodityList"
              range-key="name"
              @change="onPickerChange($event, 'commodityType')"
            >
              <view>
                {{ commodityList[commodityType]['name'] }}
              </view>
            </picker>
            <uni-icons type="forward" size="16" color="#8b8b8b"></uni-icons>
          </view>
        </view>

        <view class="breeds">
          <view class="breeds-title">页面</view>
          <view class="breeds-picker">
            <picker
              mode="selector"
              :value="pageType"
              :range="pageList"
              range-key="name"
              @change="onPickerChange($event, 'pageType')"
            >
              <view>
                {{ pageList[pageType]['name'] }}
              </view>
            </picker>
            <uni-icons type="forward" size="16" color="#8b8b8b"></uni-icons>
          </view>
        </view>
      </view>
      <view class="drawer-fotter">
        <button type="primary" class="btn reset" @tap="onReset">重置</button>
        <button type="primary" class="btn confirm" @tap="onConfirm">
          确定
        </button>
      </view>
    </uni-drawer>
  </view>
</template>

<script>
import { getEntiretyStatis, getMultidimenStatis } from '@/api/statistics'
// import { getOptions } from './chart'
import { getCommodityList } from '@/api/commodity'

import f2 from '@/components/i-uni-f2/f2.vue'

export default {
  name: 'Statistics',
  components: {
    f2
  },

  data() {
    return {
      pageType: 0,
      pageList: [{ id: '', name: '请选择' }],
      commodityList: [{ id: '', name: '请选择' }],
      commodityType: 0,
      isopen: false,
      handleDateStart: '',
      handleDateEnd: '',
      uvChartOp: [],
      // {
      //   option: getOptions('rgb(0 200 83 / 100%)', 'rgb(0 200 83 / 20%)')
      // },
      pvChartOp: [],
      //  {
      //   option: getOptions('rgb(255 152 0 / 100%)', 'rgb(255 152 0 / 20%)')
      // },
      entirety: {
        boundTracing: 0,
        pv: 0,
        uv: 0
      },
      multidimen: {
        pv: 0,
        uv: 0
      },
      activeIndex: 0,
      tabs: ['整体数据', '多维数据']
    }
  },
  computed: {
    accessTypeMap() {
      return this.$store.state.user.dictMap.srv_access_type
    }
  },

  onShow() {
    this.getCommodityList()
    this.getpageList()
    this.initData()
  },

  methods: {
    async getpageList() {
      // const res = await getPages({ dicType: 'srv_access_type' })
      // const pages = res.map(el => {
      //   return { name: el.dicValue, id: el.id }
      // })
      this.pageList = [{ id: '', name: '请选择' }]
      Object.keys(this.accessTypeMap).forEach(key => {
        this.pageList.push({ id: key, name: this.accessTypeMap[key] })
      })
    },
    async getCommodityList() {
      const res = await getCommodityList({ brandId: '' })
      if (res.code === '0') {
        this.commodityList = res.list.map(item => {
          return {
            id: item.id,
            name: item.commodityName,
            netContent: item.netContent
          }
        })
        this.commodityList = [{ id: '', name: '请选择' }, ...this.commodityList]
      }
    },
    onPickerChange(e, key) {
      this[key] = e.detail.value
    },
    filterDate(date) {
      const dateArr = date ? date.split('-') : []
      return {
        year: dateArr[0],
        month: dateArr[1],
        day: dateArr[2]
      }
    },
    onDateChange(e, key) {
      if (key === 'Start') {
        this.handleDateEnd = ''
      } else {
        //兼容ios
        if (isNaN(Date.parse(this.handleDateStart))) {
          this.handleDateStart = this.handleDateStart.replace(/-/g, '/')
        }
        if (isNaN(Date.parse(e.detail.value))) {
          e.detail.value = e.detail.value.replace(/-/g, '/')
        }
        var sdate = new Date(this.handleDateStart)
        var now = new Date(e.detail.value)
        var days = now.getTime() - sdate.getTime()
        var day = parseInt(days / (1000 * 60 * 60 * 24))
        if (day < 0) {
          this.handleDateEnd = ''
          uni.showToast({
            title: '结束日期必须大于开始日期，请重新选择！',
            icon: 'none'
          })

          return
        }
      }
      this[`handleDate${key}`] = e.detail.value
    },
    async initChartPv(F2, config) {
      await this.$parent.initData()
      const chart = new F2.Chart(config)
      const defs = {
        key: {
          type: 'timeCat',
          range: [0, 1],
          tickCount: 3
        }
      }
      chart.axis(false)
      chart.source(this.$parent.pvChartOp, defs)
      chart
        .line()
        .position('key*value')
        .shape('smooth')
        .color('l(0) 0:#ff9800 0.5:#ff980080 1:#ff980033')
      chart
        .area()
        .position('key*value')
        .shape('smooth')
        .color('l(0) 0:#ff9800 0.5:#ff980080 1:#ff980033')
      chart.render()
      return chart
    },
    async initChartUv(F2, config) {
      await this.$parent.initData()
      const chart = new F2.Chart(config)
      const defs = {
        key: {
          type: 'timeCat',
          range: [0, 1],
          tickCount: 3
        }
      }
      chart.axis(false)
      chart.source(this.$parent.uvChartOp, defs)
      chart
        .line()
        .position('key*value')
        .shape('smooth')
        .color('l(0) 0:#00c853 0.5:#00c85380 1:#00c85333')
      chart
        .area()
        .position('key*value')
        .shape('smooth')
        .color('l(0) 0:#00c853 0.5:#00c85380 1:#00c85333')
      chart.render()
      return chart
    },
    async initData() {
      if (this.activeIndex === 0) {
        const res = await getEntiretyStatis()
        this.entirety = res.resObj
        console.log(res)
      } else {
        const res = await getMultidimenStatis({
          statisticsDateStart: this.handleDateStart,
          statisticsDateEnd: this.handleDateEnd,
          commodityId: this.commodityList[this.commodityType].id,
          accType: this.pageType,
          pageSize: '0',
          pageNum: '1'
        })
        this.multidimen.pv = res.resObj.pvStatisCount.pvTotal
        this.multidimen.uv = res.resObj.uvStatisCount

        this.uvChartOp = res.resObj.uvStatis.map(el => {
          return { key: el.statisticsDate, value: el.uv }
        })
        this.pvChartOp = res.resObj.pvStatis.map(el => {
          return { key: el.statisticsDate, value: el.pv }
        })
      }
    },
    changeTab(index) {
      this.activeIndex = index
      this.initData()
    },
    showDrawer() {
      // this.isopen = true
      this.$refs.showRight.open()
    },
    changeDrawer(status) {
      this.isopen = status
      if (!this.isopen) this.initData()
    },
    onReset() {
      this.handleDateStart = ''
      this.handleDateEnd = ''
      this.pageType = 0
      this.commodityType = 0
    },

    onConfirm() {
      this.$refs.showRight.close()
      this.initData()
    }
  }
}
</script>

<style lang="scss">
.statistics-tab {
  text-align: center;
  height: 88rpx;
  line-height: 88rpx;
  background-color: #fff;

  .tab {
    width: 50%;
    display: inline-block;
    font-family: PingFangSC-Regular;
    font-size: 28rpx;
    color: #8b8b8b;
  }

  .active .statistics-text {
    position: relative;
    color: #212121;

    &::before {
      content: '';
      display: block;
      position: absolute;
      width: 40rpx;
      height: 4rpx;
      top: 60rpx;
      left: 50%;
      transform: translateX(-50%);
      background-color: #00c853;
    }
  }
}

.sta-title {
  height: 80rpx;
  line-height: 80rpx;
  text-indent: 33rpx;
  background-color: #fafafa;
  font-family: PingFangSC-Regular;
  font-size: 24rpx;
  font-weight: normal;
  font-stretch: normal;
  letter-spacing: 0;
  color: #8b8b8b;
}

.sta-title2 {
  position: relative;
  height: 104rpx;
  background-color: #fafafa;
  line-height: 104rpx;
  text-indent: 33rpx;
  font-family: PingFangSC-Regular;
  font-size: 24rpx;
  font-weight: normal;
  font-stretch: normal;
  letter-spacing: 0;
  color: #8b8b8b;

  .bars-icon {
    position: absolute;
    right: 37rpx;
  }
}

.sta-body {
  width: 718rpx;
  height: 184rpx;
  background-color: #fff;
  border-radius: 16rpx;
  border: solid 1rpx #eee;
  margin: 0 auto;
}

.sta-body-item {
  position: relative;
  display: inline-block;
  width: 33%;
  height: 104rpx;
  margin-top: 40rpx;
  border-right: 1rpx solid rgba(0, 0, 0, 0.1);

  &:last-child {
    border-right: none;
  }

  .body-item-value {
    position: absolute;
    left: 50%;
    top: 0;
    transform: translateX(-50%);
    font-family: PingFangSC-Semibold;
    font-size: 48rpx;
    font-weight: bold;
    font-stretch: normal;
    letter-spacing: 0;
    color: #212121;
  }

  .body-item-key {
    position: absolute;
    left: 50%;
    bottom: 0;
    transform: translateX(-50%);
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    font-weight: normal;
    font-stretch: normal;
    letter-spacing: 0;
    color: #8b8b8b;
    white-space: nowrap;
  }
}

.pv,
.uv {
  position: relative;
  width: 718rpx;
  height: 184rpx;
  background-color: #fff;
  border-radius: 16rpx;
  border: solid 1rpx #eee;
  margin: 8rpx auto;
}

.charts {
  position: absolute;
  top: 0;
  right: 0;
  width: 66%;
  height: 184rpx;
}

.drawer {
  width: 100%;

  .drawer-content {
    margin-top: 40rpx;

    .title {
      height: 80rpx;
      padding: 0 32rpx;
      display: flex;
      align-items: center;
      font-size: 24rpx;
      color: #8b8b8b;
    }

    .time-container {
      height: 48rpx;
      padding: 0 32rpx;
      margin-bottom: 26rpx;
      display: flex;

      .time-range {
        display: flex;

        .year {
          width: 100rpx;
          height: 48rpx;
          border-radius: 8rpx;
          border: solid 1rpx #e0e0e0;
          font-size: 24rpx;
          color: #212121;
          text-align: center;
          line-height: 48rpx;
        }

        .month {
          margin: 0 8rpx;
        }

        .month,
        .day {
          width: 64rpx;
          height: 48rpx;
          border-radius: 8rpx;
          border: solid 1rpx #e0e0e0;
          font-size: 24rpx;
          color: #212121;
          text-align: center;
          line-height: 48rpx;
        }
      }

      .b {
        width: 24rpx;
        margin: 0 4rpx;
        position: relative;

        &::after {
          content: '';
          position: absolute;
          width: 24rpx;
          height: 2rpx;
          background-color: #e0e0e0;
          top: 50%;
        }
      }
    }

    .breeds {
      padding: 0 32rpx;
      display: flex;
      font-size: 12px;
      align-items: center;
      height: 40px;
      color: #8b8b8b;
      justify-content: space-between;

      .breeds-picker {
        display: flex;
        align-items: center;
        justify-content: center;
      }

      // .breeds-btn {
      //   // width: 168rpx;
      //   height: 48rpx;
      //   line-height: 48rpx;
      //   border-radius: 24rpx;
      //   border: 1rpx solid #00c853;
      //   color: #212121;
      // }
    }
  }

  .drawer-fotter {
    width: 600rpx;
    height: 112rpx;
    background-color: #fff;
    position: fixed;
    bottom: 0;
    display: flex;
    align-items: center;

    .btn {
      width: 252rpx;
      height: 88rpx;
      font-size: 32rpx;
      line-height: 88rpx;
      border-radius: 44rpx;
      border: 1rpx solid #00c853;
    }

    .reset {
      color: #00c853;
      background: #fff;
    }

    .confirm {
      color: #fff;
      background: #00c853;
    }
  }
}
</style>
