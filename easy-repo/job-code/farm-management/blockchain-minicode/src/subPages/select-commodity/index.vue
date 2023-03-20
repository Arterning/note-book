<!--
 * @Description  : 商品列表查询
 * @Autor        : Hemingway
 * @Date         : 2021-09-13 17:30:15
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-12 10:47:01
 * @FilePath     : \blockchain-minicode\src\subPages\select-commodity\index.vue
-->
<template>
  <view class="select-commodity">
    <search-bar
      placeholder="商品名称"
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
import { queryCommodityList } from '@/api/grain-steps'
import SearchBar from '@/components/search-bar'
import Card from '@/components/card'
import listMixin from '@/mixins/listMixin'

export default {
  name: 'SelectCommodity',

  components: {
    SearchBar,
    Card
  },

  mixins: [listMixin],

  data() {
    return {
      // 观测对象
      watchObj: {
        name: '' // 商品名称 -> 映射搜索框关键词！（key必须为name）
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
     * @description: 商品点击事件
     * @param {Object} commodity 商品
     * @author: Hemingway
     */
    onClick(commodity) {
      console.log('commodity = ', commodity)
      this.$bus.$emit('chooseCommodity', commodity)
      this.$Router.back()
    },

    /**
     * @description: 查询商品列表，获取商品信息，如种植品种等
     * @author: Hemingway
     */
    async queryList() {
      if (this.status === 'more') {
        this.status = 'loading'

        console.log('this.$Route.query = ', this.$Route.query)

        const payload = {
          pageNum: this.pageNum,
          enterpriseId: this.$Route.query.enterpriseId // 品牌商id
        }
        this.watchObj.name && (payload.commodityNameLike = this.watchObj.name)

        try {
          const { code, list, hasNextPage } = await queryCommodityList(payload)
          if (code === '0') {
            this.list = this.list.concat(
              (list || []).map(
                ({
                  brandId,
                  brandName,
                  commodityId,
                  commodityName,
                  netContent
                }) => {
                  const data = {
                    title: commodityName,
                    items: [
                      {
                        text: `品牌：${brandName}`
                      }
                    ],
                    brandId,
                    brandName,
                    commodityId,
                    commodityName,
                    netContent: JSON.parse(netContent)
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
          console.log('查询商品列表 error = ', error)
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.select-commodity {
  padding: 88rpx 0 40rpx;
}
</style>
