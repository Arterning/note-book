<!--
 * @Description  : 审核进度页
 * @Autor        : Hemingway
 * @Date         : 2021-07-21 14:41:57
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 10:27:29
 * @FilePath     : \blockchain-minicode\src\pages\review-progress\index.vue
-->
<template>
  <view class="review-progress">
    <search-bar
      placeholder="企业名称"
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
      :array="dataArray"
      @onSelect="onSelect"
      @onReset="onReset"
    >
    </drawer>

    <uni-list :border="false">
      <card v-for="(item, i) in list" :key="i" :data="item">
        <template #default>
          <tag
            :text="item.extra"
            :color="item.statusColor[0]"
            :background-color="item.statusColor[1]"
          ></tag>
        </template>
        <template #footer>
          <view
            v-if="item.status === '0'"
            class="card__footer"
            @click="toUpdate(item.id)"
            >修改申请</view
          >
          <view v-else class="card__footer" @click="toDetail(item.id)"
            >查看详情</view
          >
        </template>
      </card>
    </uni-list>
    <uni-load-more :status="status" @clickLoadMore="queryList"></uni-load-more>
  </view>
</template>

<script>
import { statusMap } from './statusMap'
import Drawer from '@/components/drawer'
import Card from '@/components/card'
import Tag from '@/components/tag'
import SearchBar from '@/components/search-bar'
import { queryApplicationList } from '@/api/user'
import listMixin from '@/mixins/listMixin'

export default {
  name: 'ReviewProgress',

  components: {
    Drawer,
    Card,
    Tag,
    SearchBar
  },

  mixins: [listMixin],

  data() {
    return {
      settleTypeList: [], // 企业类型选择列表

      // 观测对象
      watchObj: {
        name: '', // 企业名称 -> 映射搜索框关键词！（key必须为name）
        type: '' // 企业类型
      }
    }
  },

  computed: {
    /**
     * @description: 企业类型map
     * @author: Hemingway
     */
    enterpriseMap() {
      return this.$store.state.user.dictMap.enterprise_type
    },

    /**
     * @description: 审核状态map
     * @author: Hemingway
     */
    checkStatusMap() {
      return this.$store.state.user.dictMap.enterpriseCheckStatus
    },

    /**
     * @description: 筛选器数据源
     * @param {*}
     * @return {Array}
     * @author: Hemingway
     */
    dataArray() {
      const arr = []
      arr.push({
        title: '企业类型',
        field: 'type', // 映射watchObj
        tags: this.settleTypeList
      })

      return arr
    }
  },

  methods: {
    /**
     * @description: 打开抽屉
     * @author: Hemingway
     */
    onRightOpen() {
      if (this.settleTypeList.length === 0) {
        // 初始化筛选器
        Object.keys(this.enterpriseMap).forEach(key => {
          this.settleTypeList.push({
            code: key,
            text: this.enterpriseMap[key],
            selected: false
          })
        })
      }

      this.$refs.drawer.open()
    },

    /**
     * @description: 筛选项点击事件
     * @param {Object} e
     * @author: Hemingway
     */
    onSelect(e) {
      this.watchObj[e.field] = e.codeList.join()
    },

    /**
     * @description: 重置筛选项
     * @author: Hemingway
     */
    onReset() {
      Object.keys(this.watchObj).forEach(key => {
        if (key !== 'name') {
          this.watchObj[key] = ''
        }
      })
    },

    /**
     * @description: 查询审核列表
     * @author: Hemingway
     */
    async queryList() {
      if (this.status === 'more') {
        this.status = 'loading'

        const payload = { pageNum: this.pageNum }
        this.watchObj.name && (payload.name = this.watchObj.name)
        this.watchObj.type && (payload.type = this.watchObj.type)

        try {
          const { code, list, hasNextPage } = await queryApplicationList(
            payload
          )
          if (code === '0') {
            this.list = this.list.concat(
              list.map(
                ({ id, checkStatus, name, type, createDate, checkTime }) => {
                  const data = {
                    id,
                    status: checkStatus,
                    extra: this.checkStatusMap[checkStatus],
                    statusColor: statusMap[checkStatus].theme, // 颜色主题
                    title: name,
                    items: []
                  }
                  data.items.push({
                    text: `企业类型：${this.enterpriseMap[type]}`
                  })
                  data.items.push({
                    text: `申请日期：${createDate.slice(0, 10)}`
                  })
                  data.items.push({
                    text: `审核日期：${(checkTime && checkTime.slice(0, 10)) ||
                      '[待审核]'}`
                  })
                  return data
                }
              )
            )

            this.pageNum++
            uni.stopPullDownRefresh()
            this.status = !hasNextPage ? 'noMore' : 'more'
          }
        } catch (error) {
          console.log('审核进度列表查询 error = ', error)
        }
      }
    },

    /**
     * @description: 跳转入驻申请（修改）
     * @param {String} id 记录id
     * @param {String} handleType 操作类型
     * @author: Hemingway
     */
    toUpdate(id) {
      this.$Router.push({
        path: '/application',
        query: { id, handleType: 'update' }
      })
    },

    /**
     * @description: 跳转入驻审核详情
     * @param {String} id 记录id
     * @author: Hemingway
     */
    toDetail(id) {
      this.$Router.push({ path: '/reviewDetail', query: { id } })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/listMixin.scss';

.review-progress {
  padding: 88rpx 0 40rpx;
}
</style>
