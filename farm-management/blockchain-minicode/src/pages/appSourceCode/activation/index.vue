<template>
  <view class="activation">
    <search-bar
      placeholder="来源农场搜索"
      :shadow="scrollTop > 8"
      @input="search"
    >
      <uni-icons
        type="bars"
        size="20"
        color="#212121"
        class="icon"
        @click="onRightOpen"
      ></uni-icons>
    </search-bar>

    <uni-segmented-control
      :values="['待绑定', '已绑定']"
      style-type="button"
      active-color="#00c853;"
      @clickItem="onSegmentChange"
    ></uni-segmented-control>

    <view class="content">
      <!-- 未确认列表 -->

      <uni-list v-if="bindingStatus === '0'" :border="false">
        <card v-for="(item, i) in unSourceCodeList" :key="i" :data="item">
          <template #default>
            <tag
              :color="item.statusColor[0]"
              :background-color="item.statusColor[1]"
            ></tag>
          </template>
          <template #footer>
            <view class="card__footer" @click="onDetail(item.sourceData)"
              >绑定溯源码</view
            >
          </template>
        </card>
      </uni-list>
      <uni-list v-else>
        <card v-for="(item, i) in sourceCodeList" :key="i" :data="item">
          <template #default>
            <tag
              :color="item.statusColor[0]"
              :background-color="item.statusColor[1]"
            ></tag>
          </template>
        </card>
      </uni-list>
      <uni-load-more
        :status="status"
        @clickLoadMore="queryList"
      ></uni-load-more>
    </view>

    <drawer
      ref="drawer"
      :width="300"
      :array="dataSingerArray"
      @onSelect="onSelect"
      @onReset="onReset"
      @onSure="onSure"
    >
      <view class="screen-wrapper"
        ><view class="screen-wrapper--inner">
          <view class="title"><text>包装结束日期</text></view>
          <view class="section" style="margin: 0;">
            <uni-datetime-picker
              v-model="watchObj.range"
              type="daterange"
              start="2021-1-1"
              :end="formatDate(new Date())"
              range-separator="至"
            /> </view></view
      ></view>
    </drawer>
  </view>
</template>

<script>
// import drawer from './components/Drawer.vue'
import { getCodeList } from '@/api/appSourceCode'
import SearchBar from '@/components/search-bar'
import Drawer from '@/components/drawer'
import doubleListMixin from '@/mixins/doubleListMixin'
import Card from '@/components/card'
// import { stringifyDate } from '@/utils/tool'

export default {
  name: 'AppSourceCode',
  components: { SearchBar, Drawer, Card },
  mixins: [doubleListMixin],
  data() {
    return {
      bindingStatus: '0',
      sourceFarmId: '', //来源农场下拉框ID
      sourceFarmNameLike: '', //来源农场模糊搜索
      jgcEnterpriseId: '', //一体化米厂下拉框ID
      variety: '', //品种ID
      packDateStart: '', //包装结束日期  范围开始
      packDateEnd: '', //包装结束日期  范围结束
      watchObj: {
        name: '', // 农场名称
        range: [], // 日期范围
        relation: '', //  品牌商/一体化米厂
        crop: '' // 作物品种
      },
      labelObj: {
        applyQuantity: '申请数量：',
        effectQuantity: '未用数量：',
        createDate: '时间：'
      },
      applyStatus: '', // 状态 (1、待支付 2、未发货 3、已发货 4、已签收 9、已作废 )
      sourceCodeList: [],
      unSourceCodeList: []
    }
  },
  computed: {},
  watch: {},

  methods: {
    search(name) {
      this.sourceFarmNameLike = name
      this.queryList()
    },
    onSegmentChange(e) {
      this.bindingStatus = e.currentIndex + ''
      this.queryList()
    },
    async queryList() {
      const params = {
        handleStatusArr: ['1', '2'], //固定值
        bindingStatus: this.bindingStatus, //溯源码绑定状态（0：待绑定、1：已绑定）
        sourceFarmNameLike: this.sourceFarmNameLike,
        jgcEnterpriseId: this.jgcEnterpriseId,
        variety: this.variety,
        packDateStart: this.packDateStart,
        packDateEnd: this.packDateEnd,
        pageNum: '1',
        pageSize: '0'
      }
      const res = await getCodeList(params)
      if (res.code === '0') {
        if (this.bindingStatus === '0') {
          this.unSourceCodeList = res.list.map(el => {
            return {
              id: el.id,
              sourceData: el,
              title: `来源农场：${el.sourceFarmName}`,
              items: [
                { text: `包装结束时间：${el.packDate}` },
                { text: `品种：${el.varietyName}` },
                { text: `商品名称：${el.commodityName}` },
                { text: `包装规格：${el.normsForMachPack}` },
                { text: `一体化米厂：${el.jgcEnterpriseName}` }
              ]
            }
          })
        } else {
          this.sourceCodeList = res.list.map(el => {
            return {
              id: el.id,
              title: `来源农场：${el.sourceFarmName}`,
              items: [
                { text: `包装结束时间：${el.packDate}` },
                { text: `品种：${el.varietyName}` },
                { text: `商品名称：${el.commodityName}` },
                { text: `包装规格：${el.normsForMachPack}` },
                { text: `溯源码编号：${el.srcCodeForMachPack}` },
                { text: `一体化米厂：${el.jgcEnterpriseName}` }
              ]
            }
          })
        }
      }
    },

    onReset(data) {
      this.watchObj.range = []
      this.applyStatus = data.applyStatus
      // this.query()
    },

    onSure() {
      this.jgcEnterpriseId = this.watchObj.relation
      this.variety = this.watchObj.crop
      if (this.watchObj.range.length > 0) {
        this.packDateStart = this.watchObj.range[0]
        this.packDateEnd = this.watchObj.range[1]
      }
      // this.applyStatus = data.applyStatus
      this.queryList()
    },

    onDetail(row) {
      this.$Router.push({
        path: '/CodeActivationDetail',
        query: {
          formData: row
        }
      })
    }
  },

  // 页面周期函数--监听页面加载
  onLoad() {},
  // 页面周期函数--监听页面初次渲染完成
  onReady() {},
  // 页面周期函数--监听页面显示(not-nvue)
  onShow() {
    this.sourceCodeList = []
    this.unSourceCodeList = []
    this.queryList()
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
  .card__footer {
    height: 68rpx;
    font-size: 28rpx;
    // font-weight: 600;
    color: #212121;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
  }

  .card__footer::after {
    content: '';
    border-top: 1rpx solid #eee;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    transform: scaleY(0.5);
  }
}

.activation {
  padding: 88rpx 0 200rpx;
  // 筛选组件插槽
  .screen-wrapper {
    padding-bottom: 24rpx;
    background-color: $uni-bg-color-page;

    .screen-wrapper--inner {
      padding-bottom: 12rpx;
      background-color: #fff;
      border-radius: 0 0 24rpx 24rpx;

      .title {
        padding: 0 24rpx;
        height: 84rpx;
        display: flex;
        justify-content: space-between;
        line-height: 84rpx;
        font-size: 26rpx;
        font-weight: bold;
      }

      .section {
        margin: 0 -16rpx;
        padding: 0 24rpx;
        display: flex;
        flex-wrap: wrap;
        align-content: flex-start;

        ::v-deep .uni-date__x-input {
          font-size: 24rpx;
        }
      }
    }
  }

  ::v-deep .segmented-control {
    margin: 24rpx 16rpx 8rpx;
    background-color: #fff;
  }

  ::v-deep .uni-calendar--fixed {
    bottom: 300rpx !important;
  }
}
</style>
