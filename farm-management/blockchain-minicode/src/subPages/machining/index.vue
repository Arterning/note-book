<!--
 * @Description  : 加工列表页
 * @Autor        : Hemingway
 * @Date         : 2021-09-16 15:03:59
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 10:12:36
 * @FilePath     : \blockchain-minicode\src\subPages\machining\index.vue
-->
<template>
  <view class="machining">
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
          <view class="title"><text>加工日期</text></view>
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
                >填写加工信息</view
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
import { queryMachiningList } from '@/api/grain-steps'
import doubleListMixin from '@/mixins/doubleListMixin'

export default {
  name: 'Machining',

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
    // 注册createMachining监听事件
    this.$bus.$on('createMachining', () => {
      this.onRefresh()
    })
  },

  beforeDestroy() {
    // 释放createMachining监听事件
    this.$bus.$off('createMachining', () => {
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
            ((payload.processStartTime = this.watchObj.range[0]),
            (payload.processEndTime = this.watchObj.range[1]))
        } // 日期
        this.watchObj.relation &&
          ((payload.cooEnterpriseId = this.watchObj.relation),
          (payload.cooEnterpriseType = this.isBrandOwner ? '6' : '2')) // 合作企业
        this.watchObj.crop && (payload.varietyId = this.watchObj.crop) // 品种

        try {
          const { code, list, hasNextPage } = await queryMachiningList(payload)
          if (code === '0') {
            this[this.type].list = this[this.type].list.concat(
              (list || []).map(
                ({
                  farmId,
                  farmName,
                  dryYear,
                  processDate,
                  varietyId,
                  varietyName,
                  enterpriseId,
                  enterpriseName,
                  riceFactoryId,
                  riceFactoryName,
                  processInfoId
                }) => {
                  const obj = {
                    title: `[农场] ${farmName}`,
                    farmId,
                    dryYear,
                    processDate,
                    varietyId,
                    enterpriseId,
                    riceFactoryId,
                    enterpriseName,
                    riceFactoryName,
                    processInfoId,
                    farmName,
                    varietyName
                  }
                  if (this.type === 'unhandled') {
                    obj.items = [
                      { text: `烘干年份：${dryYear}` },
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
                      { text: `加工结束日期：${this.formatDate(processDate)}` },
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
          console.log('加工列表查询 error = ', error)
        }
      }
    },

    /**
     * @description: 跳转表单填写
     * @param {String} riceFactoryName 一体化米厂名称
     * @param {String} processInfoId 加工id，关联待处理烘干批次
     * @param {String} farmName 农场名称
     * @param {String} varietyName 品种名称
     * @param {String} dryYear 烘干年份
     * @param {String} enterpriseName 品牌商名称
     * @author: Hemingway
     */
    toInfoCreate({
      riceFactoryName,
      processInfoId,
      farmName,
      varietyName,
      dryYear,
      enterpriseName
    }) {
      console.log(
        riceFactoryName,
        processInfoId,
        farmName,
        varietyName,
        dryYear,
        enterpriseName
      )
      this.$Router.push({
        path: '/machiningInfo',
        query: {
          riceFactoryName: riceFactoryName,
          processInfoId: processInfoId.toString(), // 加工id
          desc: [
            {
              label: '农场',
              value: farmName
            },
            {
              label: '品种',
              value: varietyName
            },
            {
              label: '烘干年份',
              value: dryYear
            },
            {
              label: '品牌商',
              value: enterpriseName
            }
          ] // 明细
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.machining {
  padding: 88rpx 0 200rpx;
}
</style>
