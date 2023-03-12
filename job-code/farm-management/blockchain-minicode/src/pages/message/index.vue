<!--
 * @Description  : 消息页
 * @Autor        : Hemingway
 * @Date         : 2021-06-22 11:35:19
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-12 15:09:54
 * @FilePath     : \blockchain-minicode\src\pages\message\index.vue
-->
<template>
  <view class="message">
    <uni-list :border="false">
      <message-item
        v-for="(item, index) in list"
        :key="index"
        class="compontent"
        :title="item.title"
        :content="item.content"
        :time="item.time"
        :is-read="item.isRead"
        @click.native="toDetail(item)"
      />
    </uni-list>
    <uni-load-more :status="status" @clickLoadMore="queryList"></uni-load-more>
  </view>
</template>

<script>
import MessageItem from './components/MessageItem.vue'
import { queryMessageList, readMessageInfo } from '@/api/user'
import listMixin from '@/mixins/listMixin'

export default {
  name: 'Message',

  components: {
    MessageItem
  },

  mixins: [listMixin],

  methods: {
    /**
     * @description: 查询消息列表
     * @author: Hemingway
     */
    async queryList() {
      if (this.status === 'more') {
        this.status = 'loading'

        try {
          const { code, list, hasNextPage } = await queryMessageList({
            pageNum: this.pageNum
          })
          if (code === '0') {
            this.list = this.list.concat(
              list.map(item => ({
                id: item.id,
                businessId: item.businessId,
                enterpriseId: item.enterpriseId,

                title: item.title,
                content: item.content,
                time: item.sendTime,
                isRead: item.messageStatus === '0' ? false : true,
                detailRouteName:
                  item.messageType === '1'
                    ? 'reviewDetail'
                    : 'partnershipDetail'
              }))
            )

            this.pageNum++
            uni.stopPullDownRefresh()
            this.status = !hasNextPage ? 'noMore' : 'more'
          }
        } catch (error) {
          console.log('查询消息列表 error', error)
        }
      }
    },

    /**
     * @description: 跳转消息（审核类或邀请类）详情
     * @param {String} id 消息记录id
     * @param {String} businessId 业务id
     * @param {String} detailRouteName 跳转至详情页的路由名称
     * @param {String} enterpriseId 企业id（邀请消息跳转至企业详情需要该参数）
     * @author: Hemingway
     */
    toDetail({ id, isRead, businessId, detailRouteName, enterpriseId }) {
      if (!isRead) {
        this.onMessageRead(id)
      }

      const isApprove = detailRouteName === 'partnershipDetail' ? true : false
      this.$Router.push({
        name: detailRouteName,
        query: { id: businessId, isApprove, enterpriseId }
      })
    },

    /**
     * @description: 将消息状态置于已读
     * @param {String} id 消息记录id
     * @author: Hemingway
     */
    async onMessageRead(id) {
      try {
        const { code } = await readMessageInfo({ id })
        if (code === '0') {
          console.log(
            'before = ',
            this.list.find(item => item.id === id).isRead
          )
          this.list.find(item => item.id === id).isRead = true
          console.log('after = ', this.list.find(item => item.id === id).isRead)
        }
      } catch (error) {
        console.log('更新消息记录 error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.message {
  padding-bottom: 40rpx;

  .compontent:last-child ::v-deep .message-item::after {
    display: none;
  }
}
</style>
