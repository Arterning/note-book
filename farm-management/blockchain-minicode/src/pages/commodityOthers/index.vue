<!--
 * @Description  : 其他素材
 * @Autor        : WuJing
 * @Date         : 2021-11-10 10:58:23
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-22 10:00:15
 * @FilePath     : \blockchain-minicode\src\pages\commodityOthers\index.vue
-->

<template>
  <view class="commodityOthers">
    <view class="link">
      <view class="title"
        >线上直购<text class="note">（请上传商城链接）</text></view
      >
      <textarea maxlength="2000" :value="text" @input="inputValue"> </textarea>
    </view>
    <view>
      <view class="title"
        >商品长图介绍<text class="note">（请上传图片，1M以内）</text></view
      >
      <view class="introduce">
        <view class="titleInput">
          <view>标题：</view>
          <view>
            <input
              class="textarea"
              maxlength="10"
              :value="title"
              placeholder="请输入4~10个字的标题，如：一粒米的旅行"
              @input="inputTitle"
            />
            <text v-if="lessFour" class="rule">*标题不能少于4个字</text>
          </view>
        </view>

        <view class="img">
          <view>长图：</view>
          <view class="image-upload">
            <image-upload
              file-class="imgLong"
              :can-del="true"
              :max="4"
              :list="imgList"
              @uploaded="onUpload('imgLong', $event)"
              @delImg="delImg('imgLong', $event)"
            />
          </view>
        </view>
      </view>
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
      title: '', //标题
      imgList: [], //图片内容数组
      id: [], //图片id数组
      lessFour: false //标题字数是否大于0小于4
    }
  },
  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    },
    isDisabled() {
      //第二个模块填写了不填完整不能提交，第一个模块填了或第二个模块填完整了都可以提交
      //不能提交的情况：1.长图介绍只填了其中一个 2.标题不在4~6个字之间的
      if (
        (!this.title && this.imgList.length > 0) ||
        (this.title && this.imgList.length === 0) ||
        this.lessFour
      ) {
        return true
      } else {
        return false
      }
    }
  },
  mounted() {
    this.otherQuery()
  },
  methods: {
    //视频链接输入方法
    inputValue(e) {
      this.text = e.detail.value
    },
    //标题输入方法
    inputTitle(e) {
      this.title = e.detail.value
      //判断输入的字数大于0，小于4
      if (this.title.length > 0 && this.title.length < 4) {
        this.lessFour = true
      } else {
        this.lessFour = false
      }
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
     * @description: 商品文字介绍提交事件
     * @author: Wujing
     */
    async onSubmit() {
      uni.showModal({
        title: '提示',
        content: '确认信息填写无误？',
        success: async res => {
          if (res.confirm) {
            try {
              let str = this.id[0] //存id字符串
              //处理数据：后端需要用,隔开的id字符串
              for (let i = 1; i < this.id.length; i++) {
                str = str + ',' + this.id[i]
              }
              const { code } = await updateMaterial({
                enterpriseId: this.role.enterpriseId,
                purchaseLink: this.text, //购买链接
                otherMaterialTitle: this.title, //标题

                otherMaterialAttachmentId: str ? str : -1 //图片id，如果没有图片，上传要传个-1（后端用来判断是删除还是更新）
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
              // 提交表单
            } catch (error) {
              console.log('更多素材提交 error = ', error)
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
    otherQuery() {
      let otherId = this.$Route.query.otherAttachmentId //图片id
      this.text = this.$Route.query.purchaseLink
      //图片回填：模拟图片上传事件
      if (otherId && otherId !== '-1') {
        console.log('test 0000', otherId)
        this.title = this.$Route.query.otherAttachmentTitle
        const idList = otherId.split(',') //拿到后端数据变成数组
        const previewList = idList.map(item => this.getImageUrl(item))
        this.onUpload('imgLong', { previewList, idList })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.commodityOthers {
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

  .introduce {
    width: 94%;
    margin: auto;
    border-radius: 10rpx;
    height: 300rpx;

    .titleInput {
      background: #fff;
      padding: 10rpx 0 20rpx 20rpx;
      border-radius: 10rpx;

      input {
        margin-top: 10rpx;
        border: solid 1rpx #808080;
        background: #fff;
        margin-left: 0rpx;
        border-radius: 10rpx;
        padding: 15rpx;
        font-size: 26rpx;
        height: 30rpx;
        width: 90%;
      }

      .rule {
        color: #f00;
        font-size: 24rpx;
      }
    }

    .img {
      background: #fff;
      padding: 10rpx 0 20rpx 20rpx;
      border-radius: 10rpx;
      margin-top: 20rpx;

      .image-upload {
        margin-top: 10rpx;
        margin-left: -20rpx;
      }
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

  .bottom {
    //定位到底部按钮
    position: absolute;
    bottom: 0;
    width: 100%;
  }
}
</style>
