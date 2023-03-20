<!--
 * @Description  : 图片上传组件
 * @Autor        : Hemingway
 * @Date         : 2021-06-30 14:00:12
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-16 09:02:43
 * @FilePath     : \blockchain-minicode\src\components\image-upload\index.vue
-->
<template>
  <view class="image-upload">
    <view v-for="(item, index) in previewList" :key="index" class="preview">
      <view v-if="canPreview(item)" class="thumbnail">
        <image
          :src="item"
          mode="aspectFill"
          class="img"
          @click="previewImage(index)"
        />
        <uni-icons
          v-if="canRefresh"
          class="refresh"
          type="refresh-filled"
          size="18"
          color="rgba(255, 0, 0, 0.8)"
          @click="onUpload(index)"
        />
        <uni-icons
          v-if="canDel"
          class="refresh"
          type="close"
          size="18"
          color="rgba(255, 0, 0, 0.8)"
          @click="delImg(index)"
        />
      </view>
      <view
        v-else
        :class="{
          add: true,
          readonly
        }"
        @click="onUpload"
      >
        <text class="iconfont icon-add_photo"></text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'ImageUpload',

  props: {
    // 是否只读
    readonly: {
      type: Boolean,
      default: false
    },
    // 图片路径list
    list: {
      type: Array,
      default: () => []
    },
    // 最大上传数量
    max: {
      type: Number,
      default: 1
    },
    // 上传图片的文件类型标记
    fileClass: {
      type: String,
      default: ''
    },
    // 图片是否可更新
    canRefresh: {
      type: Boolean,
      default: false
    },
    //图片是否可删除
    canDel: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {}
  },

  computed: {
    // 预览list
    previewList() {
      if (this.list.length < this.max) {
        return [...this.list, { finished: false }]
      }

      return this.list
    },

    // 剩余可上传图片数量
    leftCount() {
      return this.max - this.list.length
    }
  },

  created() {},

  mounted() {},

  methods: {
    /**
     * @description: 能否被预览
     * @param {String | Object} item
     * @return {Boolean}
     * @author: Hemingway
     */
    canPreview(item) {
      return typeof item === 'string' ? true : false
    },

    /**
     * @description: 预览大图事件
     * @param {Number} index
     * @author: Hemingway
     */
    previewImage(index) {
      uni.previewImage({
        current: index,
        urls: this.previewList.filter(item => typeof item === 'string')
      })
    },

    /**
     * @description: 删除照片事件
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */
    delImg(index) {
      uni.showModal({
        title: '提示',
        content: '确定要删除照片吗？',
        success: res => {
          if (res.confirm) {
            this.previewList.splice(index, 1)
            this.$emit('delImg', {
              index
            })
          }
        }
      })
    },
    /**
     * @description: 点击上传（更新）图片事件
     * @param {Number} index 索引（函数带参数，表示更新操作）
     * @author: Hemingway
     */
    onUpload(index) {
      if (this.readonly) {
        return
      }
      console.log('index = ', index)
      console.log('this.leftCount = ', this.leftCount)

      uni.chooseImage({
        count: typeof index === 'number' ? 1 : this.leftCount, // 剩余最大可传数
        sizeType: ['original', 'compressed'],
        success: async res => {
          const tempFilePaths = res.tempFilePaths

          uni.showLoading({
            title: '图片上传中...'
          })

          const promiseList = []
          tempFilePaths.forEach(filePath => {
            promiseList.push(this.uploadFile(filePath))
          })

          try {
            const idList = await Promise.all(promiseList)

            if (typeof index === 'number') {
              // 更新图片
              this.$emit('updated', {
                index,
                previewList: res.tempFilePaths,
                idList
              })
            } else {
              // 上传图片
              this.$emit('uploaded', {
                previewList: res.tempFilePaths,
                idList
              })
            }
          } catch (error) {
            console.log('图片上传 error = ', error)
          }

          uni.hideLoading()
        }
      })
    },

    /**
     * @description: 封装上传图片过程为promise
     * @param {String} filePath 临时文件路径
     * @author: Hemingway
     */
    uploadFile(filePath) {
      return new Promise((resolve, reject) =>
        uni.uploadFile({
          url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/uploadFile/w/v1`,
          filePath,
          name: 'file',
          formData: {
            fileClass: this.fileClass,
            sessionId: this.$store.state.user.sessionId,
            clientId: 'poweb'
          },
          success: res => {
            console.log('res = ', res)
            const code = JSON.parse(res.data).code
            if (code === '0') {
              resolve(JSON.parse(res.data).id)
            } else {
              uni.showToast({
                title: '图片上传异常',
                duration: 1000,
                icon: 'none'
              })
              reject()
            }
          },
          fail: () => {
            uni.showToast({
              title: '图片上传异常',
              duration: 1000,
              icon: 'none'
            })
            reject()
          }
        })
      )
    }
  }
}
</script>

<style lang="scss" scoped>
.image-upload {
  padding: 0 20rpx;
  display: flex;
  flex-flow: wrap;

  .preview {
    width: 120rpx;
    height: 120rpx;
    margin-right: 60rpx;

    // 缩略图
    .thumbnail {
      width: 100%;
      height: 100%;
      position: relative;

      .img {
        width: 100%;
        height: 100%;
        border-radius: 10rpx;
      }

      .refresh {
        background: radial-gradient(#fff, transparent);
        border-radius: 50%;
        position: absolute;
        top: 0;
        right: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        // box-shadow: 0 0 16rpx 4rpx rgba(255, 0, 0, 0.5) inset;
        // background: linear-gradient(
        //   to bottom,
        //   transparent,
        //   rgba(255, 0, 0, 0.5)
        // );
      }
    }

    // 添加图片
    .add {
      box-sizing: border-box;
      width: 100%;
      height: 100%;
      background: #fafafa;
      font-size: 20rpx;
      color: #ccc;
      border: 2rpx solid #eee;
      border-radius: 8rpx;
      display: flex;
      justify-content: center;
      align-items: center;

      .iconfont {
        font-size: 48rpx;
        color: #c7c7c7;
      }
    }

    .readonly {
      opacity: 0.6;
    }
  }

  //  第5个以后加上外边距
  .preview:nth-child(n + 5) {
    margin-top: 32rpx;
  }

  // 每行最后一列不加右外边距
  .preview:nth-child(4n) {
    margin-right: 0;
  }

  // 最后一个不需右外边距
  .preview:last-child {
    margin-right: 0;
  }
}
</style>
