<!--
 * @Description  : 收粮列表页
 * @Autor        : Hemingway
 * @Date         : 2021-07-21 14:41:57
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 10:28:53
 * @FilePath     : \blockchain-minicode\src\subPages\grain-in\index.vue
-->
<template>
  <view class="grain-in">
    <search-bar
      placeholder="农场名称"
      :shadow="scrollTop > 8"
      @search="onSearch"
      @clear="onClear"
    >
      <uni-icons
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
      :array="dataArray"
      @onSelect="onSelect"
      @onReset="onReset"
    >
      <view class="screen-wrapper"
        ><view class="screen-wrapper--inner">
          <view class="title"><text>收粮日期</text></view>
          <view class="section">
            <uni-datetime-picker
              v-model="watchObj.range"
              type="daterange"
              start="2021-1-1"
              :end="formatDate(new Date())"
              range-separator="至"
            /> </view></view
      ></view>
    </drawer>

    <uni-segmented-control
      :values="['未确认', '已确认']"
      style-type="button"
      active-color="#00c853;"
      @clickItem="onSegmentChange"
    ></uni-segmented-control>

    <view class="content">
      <!-- 未确认列表 -->
      <view v-show="current === 0">
        <uni-list :border="false">
          <card v-for="(item, i) in unconfirmed.list" :key="i" :data="item">
            <template #default>
              <tag
                :text="item.extra"
                :color="item.statusColor[0]"
                :background-color="item.statusColor[1]"
              ></tag>
            </template>
            <template #footer>
              <view
                v-if="!isBrandOwner && item.status === '-1'"
                class="card__footer"
                @click="toUpdate(item.id)"
                >修改内容</view
              >
            </template>
          </card>
        </uni-list>
        <uni-load-more
          :status="status"
          @clickLoadMore="queryList"
        ></uni-load-more>
      </view>

      <!-- 已确认列表 -->
      <view v-show="current === 1">
        <uni-list :border="false">
          <card v-for="(item, i) in confirmed.list" :key="i" :data="item">
            <template #default>
              <tag
                :text="item.extra"
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
    </view>

    <bottom-button
      v-if="!isBrandOwner"
      text="记录收粮信息"
      @click="
        $Router.push({
          name: 'grainInRecord'
        })
      "
    ></bottom-button>
  </view>
</template>

<script>
import BottomButton from '@/components/bottom-button'
import { statusMap } from './statusMap'
import SearchBar from '@/components/search-bar'
import Drawer from '@/components/drawer'
import Card from '@/components/card'
import Tag from '@/components/tag'
import { queryGrainInList } from '@/api/grain-steps'
import doubleListMixin from '@/mixins/doubleListMixin'

export default {
  name: 'ReviewProgress',

  components: {
    SearchBar,
    Drawer,
    Card,
    Tag,
    BottomButton
  },

  mixins: [doubleListMixin],

  data() {
    return {
      // 未确认
      unconfirmed: {
        status: 'more', // 列表加载状态
        pageNum: 1,
        list: []
      },

      // 已确认
      confirmed: {
        isInit: false,
        status: 'more', // 列表加载状态
        pageNum: 1,
        list: []
      },

      // 观测对象
      watchObj: {
        name: '', // 农场名称
        range: [], // 日期范围
        relation: '', //  品牌商/一体化米厂
        crop: '' // 作物品种
      }
    }
  },

  computed: {
    /**
     * @description: 操作类型
     * @return {String}
     * @author: Hemingway
     */
    type() {
      return this.current === 0 ? 'unconfirmed' : 'confirmed'
    },

    /**
     * @description: 查询状态：0 待确认，1 已确认，-1 未通过
     * @return {String}
     * @author: Hemingway
     */
    queryStatus() {
      return this.current === 0 ? '0,-1' : '1'
    },

    /**
     * @description: 确认状态map
     * @author: Hemingway
     */
    checkStatusMap() {
      return this.$store.state.user.dictMap.FARM_CHECK_STATUS
    }
  },

  created() {
    // 注册createGrainIn监听事件
    this.$bus.$on('createGrainIn', () => {
      this.onRefresh()
    })
  },

  beforeDestroy() {
    // 释放createGrainIn监听事件
    this.$bus.$off('createGrainIn', () => {
      this.onRefresh()
    })
  },

  methods: {
    /**
     * @description: 打开抽屉
     * @author: Hemingway
     */

    /**
     * @description: 分段器切换事件
     * @param {Object} e
     * @author: Hemingway
     */
    onSegmentChange(e) {
      this.current = e.currentIndex

      // 确认列表，切换时才加载
      if (this.current === 1 && !this.confirmed.isInit) {
        this.queryList() // 首次加载
        this.confirmed.isInit = true
      }
    },

    /**
     * @description: 列表查询
     * @author: Hemingway
     */
    async queryList() {
      if (this[this.type].status === 'more') {
        this[this.type].status = 'loading'

        const payload = {
          pageNum: this[this.type].pageNum,
          status: this.queryStatus,
          enterpriseId: this.role.enterpriseId,
          enterpriseType: this.role.enterpriseType
        }

        // 添加过滤条件
        this.watchObj.name && (payload.farmName = this.watchObj.name)
        this.watchObj.range.length > 0 &&
          ((payload.startDate = this.watchObj.range[0]),
          (payload.endDate = this.watchObj.range[1]))
        this.watchObj.relation &&
          ((payload.cooEnterpriseId = this.watchObj.relation),
          (payload.cooEnterpriseType = this.isBrandOwner ? '6' : '2'))
        this.watchObj.crop && (payload.varietyId = this.watchObj.crop)

        try {
          const { code, list, hasNextPage } = await queryGrainInList(payload)
          if (code === '0') {
            this[this.type].list = this[this.type].list.concat(
              list.map(
                ({
                  feedListId,
                  status,
                  farmName,
                  reachingDate,
                  carNumber,
                  riceVarietyName,
                  wetRiceWeight,
                  wetRicePrimage,
                  enterpriseName,
                  riceFactoryName
                }) => {
                  return {
                    id: feedListId,
                    status,
                    extra: this.checkStatusMap[status],
                    statusColor: statusMap[status].theme, // 颜色主题
                    title: `[农场] ${farmName}`,
                    items: [
                      { text: `收粮日期：${this.formatDate(reachingDate)}` },
                      { text: `车牌号：${carNumber}` },
                      { text: `品种：${riceVarietyName}` },
                      { text: `湿稻谷重量(kg)：${wetRiceWeight}` },
                      { text: `湿稻谷含水量(%)：${wetRicePrimage}` },
                      {
                        text: `${
                          this.isBrandOwner ? '一体化米厂' : '品牌商'
                        }：${
                          this.isBrandOwner ? riceFactoryName : enterpriseName
                        }`
                      }
                    ]
                  }
                }
              )
            )

            this[this.type].pageNum++
            uni.stopPullDownRefresh()
            this[this.type].status = !hasNextPage ? 'noMore' : 'more'
          }
        } catch (error) {
          console.log('收粮清单列表查询 error = ', error)
        }
      }
    },

    /**
     * @description: 跳转收粮清单（修改）
     * @param {String} id 记录id
     * @author: Hemingway
     */
    toUpdate(id) {
      this.$Router.push({ path: '/grainInRecord', query: { id } })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.grain-in {
  padding: 88rpx 0 200rpx;

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
