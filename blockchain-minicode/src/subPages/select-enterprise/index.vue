<!--
 * @Description  : 企业列表查询，选择品牌商
 * @Autor        : Hemingway
 * @Date         : 2021-09-13 17:30:15
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-27 10:42:28
 * @FilePath     : \blockchain-minicode\src\pages\select-enterprise\index.vue
-->
<template>
  <view class="select-enterprise">
    <search-bar
      placeholder="企业名称"
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
import { queryEnterpriseRelations } from '@/api/grain-steps'
import SearchBar from '@/components/search-bar'
import Card from '@/components/card'
import listMixin from '@/mixins/listMixin'

export default {
  name: 'SelectEnterprise',

  components: { SearchBar, Card },

  mixins: [listMixin],

  data() {
    return {
      // 观测对象
      watchObj: {
        name: '' // 企业名称 -> 映射搜索框关键词！（key必须为name）
      }
    }
  },

  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    }
  },

  watch: {},

  created() {},

  mounted() {},

  methods: {
    /**
     * @description: 品牌商点击事件
     * @param {Object} brandOwer 品牌商
     * @author: Hemingway
     */
    onClick(brandOwer) {
      console.log('farm = ', brandOwer)
      this.$bus.$emit('chooseBrandOwer', brandOwer)
      this.$Router.back()
    },

    /**
     * @description: 查询品牌商列表
     * @author: Hemingway
     */
    async queryList() {
      const payload = {
        enterpriseId: this.role.enterpriseId,
        enterpriseType: this.role.enterpriseType
      }
      this.watchObj.name && (payload.enterpriseName = this.watchObj.name)

      try {
        const { code, relations } = await queryEnterpriseRelations(payload)
        if (code === '0') {
          this.list = this.list.concat(
            (relations.enterprises || []).map(
              ({ enterpriseId, enterpriseName, address, detailAddress }) => {
                const data = {
                  id: enterpriseId,
                  title: enterpriseName, // 品牌商名称
                  items: [
                    {
                      text: `地址：${address}${detailAddress}`
                    }
                  ],
                  enterpriseId,
                  enterpriseName
                }
                return data
              }
            )
          )

          uni.stopPullDownRefresh()
          this.status = 'noMore' // 接口未分页，直接写死
        }
      } catch (error) {
        console.log('审核进度列表查询 error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.select-enterprise {
  padding: 88rpx 0 40rpx;
}
</style>
