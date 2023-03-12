<!--
 * @Author: guoxi
 * @Date: 2021-05-25 14:05:47
 * @LastEditTime : 2022-07-06 16:22:05
 * @LastEditors  : Your Name
 * @Description: 员工列表页
 * @FilePath     : \blockchain-minicode\src\pages\report\List.vue
-->
<template>
  <view class="report-list">
    <search-bar
      placeholder="来源农场搜索"
      :shadow="scrollTop > 8"
      @input="search"
    >
      <uni-icons
        style="margin-left: 32rpx;"
        type="bars"
        size="20"
        color="#212121"
        class="icon"
        @click="onRightOpen"
      ></uni-icons>
    </search-bar>

    <drawer
      ref="drawer"
      :width="300"
      :array="dataSingerArray"
      @onSelect="onSelect"
      @onReset="onReset"
      @onSure="onSure"
    >
      <view class="drawer">
        <view class="title">加工年份</view>
        <view class="section">
          <view
            v-for="(tag, i) in years"
            :key="i"
            class="tag-wrapper tag-year"
            @click="onSelectYear(i)"
            ><span :class="{ tag: true, selected: tag.selected }">{{
              tag.text
            }}</span></view
          >
        </view>
      </view>
    </drawer>

    <uni-segmented-control
      :values="['待上传', '已上传']"
      style-type="button"
      active-color="#00c853;"
      @clickItem="onSegmentChange"
    ></uni-segmented-control>

    <uni-list v-if="params.checkStatus === '0'" :border="false">
      <card v-for="(item, i) in unReportList" :key="i" :data="item">
        <template #default>
          <tag
            :color="item.statusColor[0]"
            :background-color="item.statusColor[1]"
          ></tag>
        </template>

        <template #footer>
          <view class="card__footer" @click="upload(item.sourceData)"
            >上传报告</view
          >
        </template>
      </card>
    </uni-list>

    <uni-list v-else :border="false">
      <card v-for="(item, i) in reportList" :key="i" :data="item">
        <template #default>
          <tag
            :color="item.statusColor[0]"
            :background-color="item.statusColor[1]"
          ></tag>
        </template>
        <template #footer>
          <view class="comm-card__btns">
            <text class="edit" @tap="select(item)">查看报告</text>
            <text
              v-if="isUpload(item)"
              class="delete"
              @tap="upload(item.sourceData)"
              >重新上传</text
            >
          </view>
        </template>
      </card>
    </uni-list>
    <uni-load-more :status="status" @clickLoadMore="queryList"></uni-load-more>
  </view>
</template>

<script>
import { getReportList } from '@/api/report'
import SearchBar from '@/components/search-bar'
import doubleListMixin from '@/mixins/doubleListMixin'
import Drawer from '@/components/drawer'
import Card from '@/components/card'
export default {
  name: 'ReportList',
  components: { SearchBar, Drawer, Card },
  mixins: [doubleListMixin],
  data() {
    return {
      watchObj: {
        name: '', // 农场名称
        range: [], // 日期范围
        relation: '', //  品牌商/一体化米厂
        crop: '' // 作物品种
      },
      years: [
        {
          text: new Date().getFullYear(),
          value: new Date().getFullYear(),
          selected: false
        }
      ],
      status: 'more', // 列表加载状态
      params: {
        uploadPage: '0',
        taskStatus: '1',
        isOriginal: '0',
        checkStatus: '0', //0：待上传、1：已上传
        sourceFarmNameLike: '', //来源农场名称模糊查询
        jgcEnterpriseId: '', //一体化米厂下拉框ID
        variety: '',
        pageSize: '0', //不分页默认传0即可
        pageNum: '1' //不分页默认传1即可
      },
      reportList: [],
      unReportList: []
    }
  },

  onShow() {
    this.reportList = []
    this.queryList()
  },

  onReachBottom() {
    console.log('到底了')
    this.queryList() // 加载下一页
  },
  onPullDownRefresh() {
    console.log('下拉刷新')
    this.queryList()
  },

  methods: {
    isUpload(item) {
      //兼容ios
      if (isNaN(Date.parse(item.sourceData.modifyDate))) {
        item.sourceData.modifyDate = item.sourceData.modifyDate.replace(
          /-/g,
          '/'
        )
      }
      const dateTime = new Date(item.sourceData.modifyDate)
      dateTime.setDate(dateTime.getDate() + 1)
      return dateTime > new Date()
    },
    search(name) {
      this.params.sourceFarmNameLike = name
      this.queryList()
    },
    onSegmentChange(e) {
      this.params.checkStatus = e.currentIndex + ''
      if (e.currentIndex === 1) {
        this.params.uploadPage = null
      } else {
        this.params.uploadPage = '0'
      }
      this.queryList() // 首次加载
    },
    onSelectYear(i) {
      this.years[i].selected = !this.years[i].selected
    },

    select({
      sourceData: { pzjcDate, pzjcAttachJson, aqjcDate, aqjcAttachJson }
    }) {
      this.$Router.push({
        path: '/Report',
        query: { flag: '1', pzjcDate, pzjcAttachJson, aqjcDate, aqjcAttachJson }
      })
    },

    upload(params) {
      this.$Router.push({
        path: '/reportUpload',
        query: {
          pzjcDate: params.pzjcDate || '',
          pzjcAttachJson: params.pzjcAttachJson || '',
          aqjcDate: params.aqjcDate || '',
          aqjcAttachJson: params.aqjcAttachJson || '',
          id: params.id || ''
        }
      })
    },
    async queryList() {
      if (this.status === 'more') {
        this.status = 'loading'
        let res = await getReportList(this.params)
        if (res.code === '0') {
          if (this.params.checkStatus === '0') {
            this.unReportList = res.list.map(el => {
              return {
                isMutliple: false,
                id: el.id,
                sourceData: el,
                title: `[来源农场] ${el.sourceFarmName}`,
                items: [
                  { text: `加工年份：${new Date().getFullYear()}年` },
                  { text: `品种：${el.varietyName}` },
                  { text: `一体化米厂：${el.jgcEnterpriseName}` }
                ]
              }
            })
          } else {
            this.reportList = res.list.map(el => {
              return {
                isMutliple: false,
                id: el.id,
                sourceData: el,
                title: `[来源农场] ${el.sourceFarmName}`,
                items: [
                  { text: `加工年份：${new Date().getFullYear()}年` },
                  { text: `品种：${el.varietyName}` },
                  { text: `一体化米厂：${el.jgcEnterpriseName}` }
                ]
              }
            })
          }

          uni.stopPullDownRefresh()
          this.status = 'more'
        }
      }
    },

    filterDate(time) {
      const timeArr = time ? time.split(' ') : []
      return timeArr === [] ? '' : timeArr[0]
    },

    onChange(e) {
      this.personType = e.detail.value
    },

    onReset() {
      this.years[0].selected = false
      this.params.jgDateEnd = null
      this.params.jgDateStart = null

      this.personId = ''
      // this.$refs.showRight.close()
      // this.query()
    },
    onSure() {
      this.params.jgcEnterpriseId = this.watchObj.relation
      this.params.variety = this.watchObj.crop
      this.queryList()
    }
  }
}
</script>

<style lang="scss" scoped>
.report-list {
  padding: 88rpx 0 200rpx;

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

  ::v-deep .segmented-control {
    margin: 24rpx 16rpx 8rpx;
    background-color: #fff;
  }

  .card__footer {
    height: 68rpx;
    font-size: 28rpx;
    color: #212121;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    border-top: 1rpx solid #dddddd4f;
  }

  .comm-card__btns {
    height: 68rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: PingFangSC-Semibold;
    font-size: 28rpx;
    color: #212121;
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

    .title {
      padding-left: 24rpx;
      height: 80rpx;
      line-height: 80rpx;
      font-size: 24rpx;
      font-weight: bold;
    }

    .section {
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
