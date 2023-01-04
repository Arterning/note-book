<!--
 * @Description  : 溯源素材列表
 * @Autor        : WuJing
 * @Date         : 2021-11-10 10:12:11
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 10:25:51
 * @FilePath     : \blockchain-minicode\src\pages\material\index.vue
-->
<template>
  <view class="material">
    <uni-list :border="false">
      <uni-list-item
        title="商品文字介绍"
        :show-arrow="true"
        :clickable="true"
        @click="text()"
      >
      </uni-list-item>

      <uni-list-item
        title="商品视频介绍"
        :show-arrow="true"
        :clickable="true"
        @click="video()"
      >
      </uni-list-item>

      <uni-list-item
        title="更多素材"
        :show-arrow="true"
        :clickable="true"
        @click="other()"
      >
      </uni-list-item>
    </uni-list>
  </view>
</template>

<script>
import { queryMaterial } from '@/api/user'
export default {
  data() {
    return {
      commodityDesc: '', //商品描述
      videoLink: '', //视频链接
      videoCoverageAttachmentId: '', //视频封面id
      purchaseLink: '', //购买链接
      otherAttachmentTitle: '', //标题
      otherAttachmentId: '' //附件图片id
    }
  },
  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    }
  },
  created() {
    // 返回时刷新接口
    this.$bus.$on('refresh', () => {
      this.materialQuery()
    })
  },
  mounted() {
    this.materialQuery()
  },
  beforeDestroy() {
    // 释放refresh监听事件
    this.$bus.$off('refresh')
  },
  methods: {
    /**
     * @description: 商品素材查询
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */

    async materialQuery() {
      try {
        const { code, data } = await queryMaterial({
          enterpriseId: this.role.enterpriseId
        })
        if (code === '0' && data) {
          this.commodityDesc = data.commodityDesc //商品介绍文字
          this.videoLink = data.videoLink //视频链接
          this.videoCoverageAttachmentId = data.videoCoverageAttachmentId //视频封面id
          this.purchaseLink = data.purchaseLink //购买链接
          this.otherAttachmentTitle = data.otherAttachmentTitle //附件标题
          this.otherAttachmentId = data.otherAttachmentId //附件长图id
        }
      } catch (error) {
        console.log('文字查询记录 error = ', error)
      }
    },
    //文字传值
    text() {
      this.$Router.push({
        path: '/commodityText',
        query: { commodityText: this.commodityDesc }
      })
    },
    //视频传值
    video() {
      this.$Router.push({
        path: '/commodityVideo',
        query: {
          videoLink: this.videoLink,
          videoCoverageAttachmentId: this.videoCoverageAttachmentId
        }
      })
    },
    //更多素材传值
    other() {
      this.$Router.push({
        path: '/commodityOthers',
        query: {
          purchaseLink: this.purchaseLink,
          otherAttachmentTitle: this.otherAttachmentTitle,
          otherAttachmentId: this.otherAttachmentId
        }
      })
    }
  }
}
</script>
<style lang="scss" scoped>
::v-deep .uni-list-item__content-title {
  font-size: 30rpx;
}
</style>
