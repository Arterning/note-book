<!--
 * @Description  : 农场列表查询，选择来源农场。
 * @Autor        : Hemingway
 * @Date         : 2021-09-13 17:30:15
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-16 17:14:36
 * @FilePath     : \blockchain-minicode\src\pages\select-farm\index.vue
-->
<template>
  <view class="select-farm">
    <search-bar
      placeholder="农场名称"
      :shadow="scrollTop > 8"
      @search="onSearch"
      @clear="onClear"
    />

    <uni-list :border="false">
      <card
        v-for="(item, i) in list"
        :key="i"
        :data="item"
        @click.native="onClick(item)"
      >
      </card>
    </uni-list>
    <uni-load-more :status="status" @clickLoadMore="queryList"></uni-load-more>
  </view>
</template>

<script>
import { queryFarmList } from '@/api/grain-steps'
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
        name: '' // 农场名称 -> 映射搜索框关键词！（key必须为name）
      }
    }
  },

  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    }
  },

  created() {},

  mounted() {},

  methods: {
    /**
     * @description: 农场点击事件
     * @param {Object} farm 农场
     * @author: Hemingway
     */
    onClick(farm) {
      console.log('farm = ', farm)
      this.$bus.$emit('chooseFarm', farm)
      this.$Router.back()
    },

    /**
     * @description: 查询农场列表，获取农场信息，如种植品种等
     * @author: Hemingway
     */
    async queryList() {
      if (this.status === 'more') {
        this.status = 'loading'

        const payload = {
          pageNum: this.pageNum,
          enterpriseId: this.role.enterpriseId
        }
        this.watchObj.name && (payload.farmName = this.watchObj.name)

        try {
          const { code, list, hasNextPage } = await queryFarmList(payload)
          if (code === '0') {
            this.list = this.list.concat(
              (list || []).map(
                ({
                  farmId,
                  farmName,
                  farmLocation,
                  varietyName,
                  varietyId,
                  phone
                }) => {
                  const data = {
                    title: farmName,
                    items: [
                      {
                        text: `品种：${varietyName}`
                      },
                      {
                        text: `地址：${farmLocation}`
                      }
                    ],
                    farmName,
                    farmId,
                    varietyName,
                    varietyId,
                    phone
                  }
                  return data
                }
              )
            )

            this.pageNum++
            uni.stopPullDownRefresh()
            this.status = !hasNextPage ? 'noMore' : 'more'
          }
        } catch (error) {
          console.log('查询农场列表 error = ', error)
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.select-farm {
  padding: 88rpx 0 40rpx;
}
</style>
