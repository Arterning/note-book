<template>
  <view class="content">
    <view class="msg"
      >以下信息将同步展示到二维码溯源页面，您可以检查后再提交</view
    >

    <uni-forms
      ref="form"
      :rules="rules"
      class="add-brand-form"
      label-width="100"
      border
    >
      <view class="msg">品质检测报告</view>
      <uni-forms-item label="检测时间" name="quality_inspection">
        <picker
          class="picker-item"
          mode="date"
          :value="quality_inspection.quality_inspection_date"
          :end="today"
          @change="changeQualityTime"
        >
          <view class="uni-input">{{
            quality_inspection.quality_inspection_date
          }}</view>
        </picker>
        <uni-icons class="unit" type="forward"></uni-icons>
      </uni-forms-item>

      <view class="img-box">
        <view class="upload-container">
          <view
            v-for="(img, i) in quality_inspection.images"
            :key="i"
            class="upload-img"
          >
            <image :src="img" @click="preViewImg(i)" />
            <uni-icons
              class="upload-icon"
              type="close"
              color="red"
              @click="handleDel(i)"
            />
          </view>
          <view class="upload" @click="chooseImg">
            <uni-icons class="upload-icon" type="plusempty" color="#f1f1f1" />
          </view>
        </view>
      </view>

      <view class="msg">安全检测报告</view>
      <uni-forms-item label="检测时间" name="safe_inspection">
        <picker
          class="picker-item"
          mode="date"
          :end="today"
          :value="safe_inspection.safe_inspection_date"
          @change="changeSafeTime"
        >
          <view class="uni-input">{{
            safe_inspection.safe_inspection_date
          }}</view>
        </picker>
        <uni-icons class="unit" type="forward"></uni-icons>
      </uni-forms-item>

      <view class="img-box">
        <view class="upload-container">
          <view
            v-for="(img, i) in safe_inspection.images"
            :key="i"
            class="upload-img"
          >
            <image :src="img" @click="preViewSafeImg(i)" />
            <uni-icons
              class="upload-icon"
              type="close"
              color="red"
              @click="handleSafeDel(i)"
            />
          </view>
          <view class="upload" @click="chooseSafeImg">
            <uni-icons class="upload-icon" type="plusempty" color="#f1f1f1" />
          </view>
        </view>
      </view>
    </uni-forms>

    <view class="add-brand">
      <button class="btn" @tap="onSubmit">
        <span>提交</span>
      </button>
    </view>
  </view>
</template>

<script>
import { uploadReport } from '@/api/report'
import {
  chooseImg,
  preViewImg,
  getImageUrl,
  getDateFormat
} from '../../utils/tool'
export default {
  name: 'DryingInfo',

  data() {
    return {
      target: null,
      today: getDateFormat(new Date()),
      imgMaxNum: 10,
      qualityImgPaths: [],
      safeImgPaths: [],

      quality_inspection: {
        quality_inspection_date: '请选择',
        ids: [],
        images: []
      },
      safe_inspection: {
        safe_inspection_date: '请选择',
        ids: [],
        images: []
      },
      formData: {
        quality_inspection: [],
        safe_inspection: []
      },
      rules: {
        quality_inspection: {
          rules: [
            // {
            //   required: true,
            //   errorMessage: '请至少上传一张照片'
            // },
            // {
            //   validateFunction: (rule, value, data, callback) => {
            //     if (value.length === 0) {
            //       callback('请至少上传一张照片')
            //     }
            //     return true
            //   }
            // }
          ]
        },
        safe_inspection: {
          rules: [
            // {
            //   required: true,
            //   errorMessage: '请至少上传一张照片'
            // },
            // {
            //   validateFunction: (rule, value, data, callback) => {
            //     console.log('value:', value)
            //     console.log('rule:', rule)
            //     console.log('data:', data)
            //     if (value.length === 0) {
            //       callback('请至少上传一张照片')
            //     }
            //     return true
            //   }
            // }
          ]
        }
      }
    }
  },

  onLoad() {
    this.target = this.$Route.query

    this.quality_inspection.quality_inspection_date = getDateFormat(
      this.target.pzjcDate
    )
    this.quality_inspection.ids = this.target.pzjcAttachJson || []

    this.quality_inspection.ids.map(el => {
      this.quality_inspection.images.push(this.getImage(el))
    })

    this.safe_inspection.safe_inspection_date = getDateFormat(
      this.target.aqjcDate
    )
    this.safe_inspection.ids = this.target.aqjcAttachJson || []

    this.safe_inspection.ids.map(el => {
      this.safe_inspection.images.push(this.getImage(el))
    })
  },

  mounted() {},

  methods: {
    getImage(item) {
      return getImageUrl(item.id)
    },
    chooseImg() {
      chooseImg(
        this.quality_inspection,
        this.imgMaxNum,
        this.qualityImgPaths,
        this
      )
    },
    preViewImg(i) {
      preViewImg(i, this.quality_inspection.images)
    },

    handleDel(i) {
      uni.showModal({
        title: '提示',
        content: '是否要删除该图片',
        success: res => {
          if (res.confirm) {
            this.qualityImgPaths.splice(i, 1)
            this.quality_inspection.images.splice(i, 1)
            this.quality_inspection.ids.splice(i, 1)
          }
        }
      })
    },
    chooseSafeImg() {
      chooseImg(this.safe_inspection, this.imgMaxNum, this.safeImgPaths, this)
    },
    preViewSafeImg(i) {
      preViewImg(i, this.safe_inspection.images)
    },

    handleSafeDel(i) {
      uni.showModal({
        title: '提示',
        content: '是否要删除该图片',
        success: res => {
          if (res.confirm) {
            this.safeImgPaths.splice(i, 1)
            this.safe_inspection.images.splice(i, 1)
            this.safe_inspection.ids.splice(i, 1)
          }
        }
      })
    },

    changeSafeTime(e) {
      this.safe_inspection.safe_inspection_date = e.target.value
    },
    changeQualityTime(e) {
      this.quality_inspection.quality_inspection_date = e.target.value
    },

    filterArrayId(arr, name) {
      const id = arr.filter(ele => ele.name === name)[0]['id']
      return id
    },

    async onSubmit() {
      if (this.quality_inspection.images.length > 0) {
        this.formData.quality_inspection = this.quality_inspection.images
        this.$refs.form.setValue(
          'quality_inspection',
          this.quality_inspection.images
        )
      }
      if (this.safe_inspection.images.length > 0) {
        this.formData.safe_inspection = this.safe_inspection.images
        this.$refs.form.setValue('safe_inspection', this.safe_inspection.images)
      }

      this.$refs.form.validate().then(async res => {
        const data = await uploadReport({
          id: this.target.id,
          aqjcAttachJsonStr: JSON.stringify(this.safe_inspection.ids),
          pzjcAttachJsonStr: JSON.stringify(this.quality_inspection.ids),
          aqjcDate: `${this.safe_inspection.safe_inspection_date} 00:00:00`,
          pzjcDate: `${this.quality_inspection.quality_inspection_date} 00:00:00`
        })

        if (data.code === '0') {
          uni.navigateTo({
            url: `/pages/report/List`
          })
        }
      })
    }
  }
}
</script>

<style lang="scss">
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-top: 1rpx solid #eee;
  border-bottom: 1rpx solid #eee;
  height: calc(100vh - 121rpx);
  overflow-y: auto;

  .msg {
    width: 100%;
    height: 80rpx;
    line-height: 80rpx;
    background-color: #fafafa;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #8b8b8b;
    text-indent: 34rpx;
  }

  .warn {
    width: 100vw;
    height: 80rpx;
    line-height: 80rpx;
    background-color: #fafafa;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #fa4251;
    text-indent: 34rpx;
    position: relative;
    left: -20px;
  }

  .qrcode {
    width: 100vw;
    height: 80rpx;
    text-align: center;
    line-height: 80rpx;
    background-color: #fafafa;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #00c853;
    position: relative;
    left: -20px;
  }

  .uni-input {
    height: 72rpx;
    line-height: 72rpx;
    text-align: right;
  }

  .picker-item {
    display: inline-block;
    width: calc(100% - 56rpx);
  }

  .common-input {
    border: none;
    text-align: right;

    .uni-easyinput__content-input {
      text-align: right;
    }
  }

  .add-brand-form {
    width: 90%;
    height: calc(100vh - 80rpx);

    .msg {
      width: 100vw;
      position: relative;
      right: 40rpx;
    }
  }

  .add-brand {
    width: 750rpx;
    height: 138rpx;
    position: fixed;
    bottom: 0;
    background-color: #fafafa;

    .btn {
      border-radius: 44rpx;
      font-size: 32rpx;
      margin: 25rpx 40rpx;
      background-color: #00c853;

      span {
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        color: #fff;
        margin-left: 10rpx;
      }
    }
  }

  .file-picker {
    width: 400rpx;
    height: 181rpx;
  }

  .unit {
    font-size: 28rpx;
    color: #212121;
    margin-left: 18rpx;
  }
}
</style>
