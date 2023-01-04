<!--
 * @Description  : 烘干列表页
 * @Autor        : Hemingway
 * @Date         : 2021-09-16 15:03:59
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 16:17:59
 * @FilePath     : \blockchain-minicode\src\subPages\drying\index.vue
-->
<template>
  <view class="drying">
    <search-bar
      :shadow="scrollTop > 8"
      placeholder="农场名称"
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
      <view v-if="type === 'handled'" class="screen-wrapper"
        ><view class="screen-wrapper--inner">
          <view class="title"><text>烘干日期</text></view>
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
      :values="['待处理', '已处理']"
      style-type="button"
      active-color="#00c853;"
      @clickItem="onSegmentChange"
    ></uni-segmented-control>

    <view class="content">
      <!-- 待处理列表 -->
      <view v-show="current === 0">
        <uni-list :border="false">
          <card v-for="(item, i) in unhandled.list" :key="i" :data="item">
            <template #footer>
              <view class="card__footer" @click="toInfoCreate(item)"
                >填写烘干信息</view
              >
            </template>
          </card>
        </uni-list>
        <uni-load-more
          :status="status"
          @clickLoadMore="queryList"
        ></uni-load-more>
      </view>

      <!-- 已处理列表 -->
      <view v-show="current === 1">
        <uni-list :border="false">
          <card v-for="(item, i) in handled.list" :key="i" :data="item"> </card>
        </uni-list>
        <uni-load-more
          :status="status"
          @clickLoadMore="queryList"
        ></uni-load-more>
      </view>
    </view>
  </view>
</template>

<script>
import SearchBar from '@/components/search-bar'
import Drawer from '@/components/drawer'
import Card from '@/components/card'
import { queryDryingList } from '@/api/grain-steps/index'
import doubleListMixin from '@/mixins/doubleListMixin'

export default {
  name: 'Drying',

  components: {
    SearchBar,
    Drawer,
    Card
  },

  mixins: [doubleListMixin],

  data() {
    return {}
  },

  created() {
    // 注册createDrying监听事件
    this.$bus.$on('createDrying', () => {
      this.onRefresh()
    })
  },

  beforeDestroy() {
    // 释放createDrying监听事件
    this.$bus.$off('createDrying', () => {
      this.onRefresh()
    })
  },

  methods: {
    /**
     * @description: 列表查询
     * @author: Hemingway
     */
    async queryList() {
      if (this[this.type].status === 'more') {
        this[this.type].status = 'loading'

        const payload = {
          pageNum: this[this.type].pageNum,
          status: this.type === 'unhandled' ? '0' : '1',
          enterpriseId: this.role.enterpriseId,
          enterpriseType: this.role.enterpriseType
        }

        // 添加过滤条件
        this.watchObj.name && (payload.farmName = this.watchObj.name) // 农场名称
        if (this.type === 'handled') {
          this.watchObj.range.length > 0 &&
            ((payload.dryEndTimeFrom = this.watchObj.range[0]),
            (payload.dryEndTimeTo = this.watchObj.range[1]))
        } // 日期
        this.watchObj.relation &&
          ((payload.cooEnterpriseId = this.watchObj.relation),
          (payload.cooEnterpriseType = this.isBrandOwner ? '6' : '2')) // 合作企业
        this.watchObj.crop && (payload.varietyId = this.watchObj.crop) // 品种

        try {
          const { code, list, hasNextPage } = await queryDryingList(payload)
          if (code === '0') {
            this[this.type].list = this[this.type].list.concat(
              list.map(
                ({
                  farmId,
                  farmName,
                  harvestYear,
                  dryDate,
                  wetRiceWeight,
                  wetRicePrimage,
                  dryRiceWeight,
                  dryRicePrimage,
                  varietyId,
                  varietyName,
                  enterpriseId,
                  enterpriseName,
                  riceFactoryId,
                  riceFactoryName
                }) => {
                  const obj = {
                    title: `[农场] ${farmName}`,
                    farmId,
                    harvestYear,
                    varietyId,
                    enterpriseId,
                    riceFactoryId,
                    riceFactoryName
                  }

                  if (this.type === 'unhandled') {
                    obj.items = [
                      { text: `收割年份：${harvestYear}` },
                      { text: `品种：${varietyName}` },
                      {
                        text: `${
                          this.isBrandOwner ? '一体化米厂' : '品牌商'
                        }：${
                          this.isBrandOwner ? riceFactoryName : enterpriseName
                        }`
                      }
                    ]
                  } else {
                    obj.items = [
                      {
                        text: `烘干日期：${this.formatDate(dryDate)}`
                      },
                      {
                        text: `湿稻谷重量(kg)：${wetRiceWeight}`
                      },
                      {
                        text: `湿稻谷含水量(%)：${wetRicePrimage}`
                      },
                      {
                        text: `干稻谷重量(kg)：${dryRiceWeight}`
                      },
                      {
                        text: `干稻谷含水量(%)：${dryRicePrimage}`
                      },
                      { text: `品种：${varietyName}` },
                      {
                        text: `${
                          this.isBrandOwner ? '一体化米厂' : '品牌商'
                        }：${
                          this.isBrandOwner ? riceFactoryName : enterpriseName
                        }`
                      }
                    ]
                  }

                  return obj
                }
              )
            )

            this[this.type].pageNum++
            uni.stopPullDownRefresh()
            this[this.type].status = !hasNextPage ? 'noMore' : 'more'
          }
        } catch (error) {
          console.log('烘干列表查询 error = ', error)
        }
      }
    },

    /**
     * @description: 跳转表单填写
     * @param {String} farmId 农场id
     * @param {String} harvestYear 收割年份
     * @param {String} varietyId 品种id
     * @param {String} enterpriseId 品牌商id
     * @param {String} riceFactoryName 米厂名称
     * @author: Hemingway
     */
    toInfoCreate({
      riceFactoryId,
      riceFactoryName,
      farmId,
      harvestYear,
      varietyId
    }) {
      this.$Router.push({
        path: '/dryingInfo',
        query: {
          riceFactoryName: riceFactoryName || '', // 当登录用户为品牌商时，可以获取到
          payload: {
            farmId: farmId || '',
            harvestYear: harvestYear || '',
            varietyId: varietyId || '',
            riceFactoryId: riceFactoryId || ''
          }
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.drying {
  padding: 88rpx 0 200rpx;
}
</style>
