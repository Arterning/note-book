<!--
 * @Description  : 商品视频介绍
 * @Autor        : WuJing
 * @Date         : 2021-11-10 10:57:42
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-22 09:43:49
 * @FilePath     : \blockchain-minicode\src\pages\commodityVideo\index.vue
-->
<template>
  <view class="commodityVideo">
    <view class="link">
      <view class="title"
        >视频链接<text class="note">（请上传视频链接）</text></view
      >
      <textarea maxlength="2000" :value="text" @input="inputValue"> </textarea>
    </view>
    <view>
      <view class="title"
        >视频封面图片<text class="note">（请上传图片，1M以内）</text></view
      >
      <view class="img"
        ><image-upload
          file-class="imgList"
          :can-del="true"
          :max="1"
          :list="imgList"
          @uploaded="onUpload('imgList', $event)"
          @delImg="delImg('imgList', $event)"
      /></view>
    </view>
    <view class="bottom">
      <bottom-button
        text="提交"
        :is-disabled="isDisabled"
        @click="onSubmit"
      ></bottom-button>
    </view>
  </view>
</template>

<script>
import ImageUpload from '@/components/image-upload'
import BottomButton from '@/components/bottom-button'
import { updateMaterial } from '@/api/user'

export default {
  components: {
    BottomButton,
    ImageUpload
  },

  data() {
    return {
      text: '', //视频链接
      imgList: [], //图片内容数组
      id: [] //图片id数组
    }
  },
  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    },
    isDisabled() {
      //当有图片没有连接或者没有图片有链接的时候，才不能提交
      if (
        (this.imgList.length === 0 && this.text) ||
        (this.imgList.length > 0 && !this.text)
      ) {
        return true
      } else {
        return false
      }
    } //规则：两个都必填才能false
  },
  mounted() {
    this.videoQuery()
  },

  methods: {
    inputValue(e) {
      this.text = e.detail.value
    },
    /**
     * @description: 图片上传事件
     * @param {String} field 域字段
     * @param {Object} event
     * @author: Wujing
     */
    onUpload(field, event) {
      console.log('onUpload', field, event)
      const { previewList, idList } = event
      this.imgList = [...this.imgList, ...previewList] // 图片预览数组
      this.id = [...this.id, ...idList] // 图片id数组
    },
    /**
     * @description:图片删除事件
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */

    delImg(field, event) {
      const { index } = event
      this.imgList.splice(index, 1) // 图片预览数组
      this.id.splice(index, 1) // 图片id数组
    },
    /**
     * @description: 获取图片下载url
     * @param {String} id
     * @return {String} url
     * @author: wujing
     */
    getImageUrl(id) {
      return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
    },
    /**
     * @description: 商品视频介绍提交事件
     * @author: Wujing
     */
    async onSubmit() {
      uni.showModal({
        title: '提示',
        content: '确认信息填写无误？',
        success: async res => {
          if (res.confirm) {
            try {
              const { code } = await updateMaterial({
                enterpriseId: this.role.enterpriseId,
                videoLink: this.text,
                //图片删除要传-1标记
                videoCoverageAttachmentId: this.id[0] ? this.id[0] : -1
              })
              if (code === '0') {
                uni.showToast({
                  title: '提交成功',
                  icon: 'success',
                  duration: 2000,
                  complete: () => {
                    setTimeout(() => {
                      this.$bus.$emit('refresh')
                      this.$Router.back()
                    }, 1500)
                  }
                })
              }
            } catch (error) {
              console.log('商品视频介绍提交 error = ', error)
            }
          }
        }
      })
    },
    /**
     * @description: 视频介绍素材查询
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */
    videoQuery() {
      let videoId = this.$Route.query.videoCoverageAttachmentId
      if (videoId && videoId !== '-1') {
        this.text = this.$Route.query.videoLink
        //图片回填：模拟图片上传事件
        const idList = [videoId]
        const previewList = [this.getImageUrl(videoId)]
        this.onUpload('imgList', { previewList, idList })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.commodityVideo {
  margin-top: 20rpx;

  .link {
    margin-bottom: 30rpx;

    textarea {
      background: #fff;
      margin: auto;
      width: 88%;
      height: 200rpx;
      border-radius: 10rpx;
      padding: 25rpx;
      font-size: 26rpx;
    }
  }

  .title {
    margin-left: 30rpx;
    font-size: 34rpx;
    margin-bottom: 10rpx;

    .note {
      font-size: 28rpx;
      color: #808080;
    }
  }

  .img {
    background: #fff;
    padding: 20rpx 0 20rpx 0;
    width: 94%;
    margin: auto;
    border-radius: 10rpx;
  }

  .bottom {
    //定位到底部按钮
    position: absolute;
    bottom: 0;
    width: 100%;
  }
}
</style>
