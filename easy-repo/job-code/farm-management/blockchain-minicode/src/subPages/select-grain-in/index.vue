<!--
 * @Description  : 农场列表查询，选择来源农场。
 * @Autor        : Hemingway
 * @Date         : 2021-09-13 17:30:15
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-23 16:15:59
 * @FilePath     : \blockchain-minicode\src\pages\select-grain-in\index.vue
-->
<template>
  <view class="select-grain-in">
    <search-bar
      placeholder="车牌号"
      :shadow="scrollTop > 8"
      @search="onSearch"
      @clear="onClear"
    />

    <uni-list :border="false">
      <card v-for="(item, i) in list" :key="i" :data="item"
        ><template #default>
          <checkbox-group
            class="checkbox-wrapper"
            @change="onRadioChange(i, $event)"
          >
            <checkbox
              :value="item.feedListId"
              :checked="item.selected"
              style="transform: scale(0.7);"
            />
          </checkbox-group>
        </template>
      </card>
    </uni-list>
    <uni-load-more :status="status" @clickLoadMore="queryList"></uni-load-more>

    <view class="bottom-wrapper">
      <view class="flex">
        本次烘干重量
        <text class="text">{{ amount }}</text
        >kg
      </view>
      <view class="flex btn-group">
        <button type="primary" class="btn" @click="handleCheckAll">
          {{ isSelectedAll ? '取消全选' : '全选' }}
        </button>
        <button type="primary" class="btn" @click="handleConfirm">
          确定
        </button>
      </view>
    </view>
  </view>
</template>

<script>
import { queryToDryingList } from '@/api/grain-steps'
import SearchBar from '@/components/search-bar'
import Card from '@/components/card'
import listMixin from '@/mixins/listMixin'

export default {
  name: 'SelectFarm',

  components: {
    SearchBar,
    Card
  },

  mixins: [listMixin],

  data() {
    return {
      // 观测对象
      watchObj: {
        name: '' // 车牌号 -> 映射搜索框关键词！（key必须为name）
      }
    }
  },

  computed: {
    // 计算是否已经全选列表
    isSelectedAll() {
      const unSelected = this.list.find(item => item.selected === false)
      return !unSelected
    },
    // 选中的列表
    selectedList() {
      return this.list.filter(item => item.selected === true)
    },
    // 计算待烘干重量
    amount() {
      if (this.selectedList.length === 0) {
        return 0
      }

      return this.selectedList.reduce((sum, current) => {
        console.log('current = ', current)
        return (parseFloat(sum) + parseFloat(current.wetRiceWeight)).toFixed(1)
      }, 0)
    }
  },

  created() {},

  mounted() {},

  methods: {
    /**
     * @description: 查询待烘干的清单列表
     * @author: Hemingway
     */
    async queryList() {
      if (this.status === 'more') {
        this.status = 'loading'

        const payload = {
          pageNum: this.pageNum,
          ...this.$Route.query.payload
        }

        this.watchObj.name && (payload.carNumber = this.watchObj.name)

        try {
          const { code, list } = await queryToDryingList(payload)
          if (code === '0') {
            this.list = this.list.concat(
              (list || []).map(
                ({
                  feedListId,
                  carNumber,
                  wetRiceWeight,
                  wetRicePrimage,
                  reachingDate
                }) => {
                  const data = {
                    title: `[车牌号] ${carNumber}`,
                    items: [
                      {
                        text: `湿稻谷重量(kg)：${wetRiceWeight}`
                      },
                      {
                        text: `湿稻谷含水量(%)：${wetRicePrimage}`
                      },
                      {
                        text: `收粮日期：${reachingDate &&
                          this.formatDate(reachingDate)}`
                      }
                    ],
                    wetRiceWeight, // 湿稻谷重量
                    wetRicePrimage, // 湿稻谷含水量
                    feedListId, // id
                    selected: this.$Route.query.selectedList
                      .map(item => item.feedListId)
                      .includes(feedListId), // 默认未被选择
                    carNumber
                  }
                  return data
                }
              )
            )

            this.pageNum++
            uni.stopPullDownRefresh()
            this.status = 'noMore' // 接口未分页，直接写死
          }
        } catch (error) {
          console.log('查询农场列表 error = ', error)
        }
      }
    },

    /**
     * @description: 收粮清单 单个点击事件
     * @param {Number} i 索引
     * @param {Object} e
     * @author: Hemingway
     */
    onRadioChange(i, e) {
      this.list[i].selected = e.detail.value.length !== 0 // 修改选中状态
    },

    /**
     * @description: 收粮清单 全选与取消全选事件
     * @author: Hemingway
     * @example:
     */
    handleCheckAll() {
      const unSelected = this.list.find(item => item.selected === false)
      if (!unSelected) {
        // 取消全选
        this.list.forEach(item => {
          item.selected = false
        })
      } else {
        // 全选
        this.list.forEach(item => {
          item.selected = true
        })
      }
    },

    /**
     * @description: 选择提交事件
     * @author: Hemingway
     */
    handleConfirm() {
      const selectedList = this.list.filter(item => item.selected === true)
      this.$bus.$emit('chooseGrainIn', selectedList)
      this.$Router.back()
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.select-grain-in {
  padding: 88rpx 0 200rpx;

  // 底部操作栏
  .bottom-wrapper {
    box-sizing: border-box;
    width: 100%;
    padding: 48rpx 32rpx 40rpx;
    position: fixed;
    bottom: 0;
    background: rgba(255, 255, 255, 0.8);
    box-shadow: 0 -24rpx 12rpx 1rpx rgba(0, 0, 0, 0.05);
    display: flex;
    align-items: center;

    .flex {
      flex: 1;
      font-size: 24rpx;

      .text {
        margin: 0 8rpx;
        color: $theme-color;
      }
    }

    .btn-group {
      display: flex;

      .btn {
        height: 68rpx;
        line-height: 68rpx;
        font-size: 32rpx;
        background-color: $uni-bg-color-btn;
        border-radius: 34rpx;
      }

      .btn:first-of-type {
        flex: 1;
        color: #00c853;
        background-color: #fff;
        border: 1rpx solid #00c853;
      }

      .btn:last-of-type {
        flex: none;
        margin-left: 32rpx;
      }

      .btn::after {
        display: none;
      }
    }
  }
}
</style>
