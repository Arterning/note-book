<template>
  <view class="content">
    <uni-forms
      ref="form"
      class="add-brand-form"
      :value="formData"
      :rules="rules"
      label-width="100"
      border
    >
      <uni-forms-item required label="申请数量" name="applyQuantity">
        <uni-easyinput
          v-model="formData.applyQuantity"
          :input-border="false"
          class="common-input"
          type="digit"
          placeholder="请输入"
        />
      </uni-forms-item>
      <!-- <view class="msg"
        >溯源码编号：
        <view class="msg-code">{{ startCode }} - {{ endCode }}</view></view
      > -->
      <uni-forms-item required label="收件人姓名" name="reciverName">
        <uni-easyinput
          v-model="formData.reciverName"
          :input-border="false"
          class="common-input"
          type="text"
          placeholder="请输入"
        />
      </uni-forms-item>
      <uni-forms-item required label="收件人手机号" name="reciverPhone">
        <uni-easyinput
          v-model="formData.reciverPhone"
          :input-border="false"
          class="common-input"
          type="number"
          placeholder="请输入"
        />
      </uni-forms-item>

      <uni-forms-item class="form-item" label="收货地址" name="region" required>
        <picker mode="region" :value="region" @change="bindRegionChange">
          <view
            :style="{
              lineHeight: '36px',
              color: !formData.region ? '#bbb' : ''
            }"
          >
            {{ formData.region || '省市区县等' }}
          </view>
        </picker>
      </uni-forms-item>

      <uni-forms-item
        class="form-item"
        label="详细地址"
        name="detailAddress"
        required
      >
        <uni-easyinput
          v-model="formData.detailAddress"
          :trim="true"
          type="text"
          :maxlength="30"
          :input-border="false"
          placeholder="乡镇、街道、楼牌号等"
          placeholder-style="color: #bbb;"
        />
      </uni-forms-item>
    </uni-forms>

    <view class="add-brand">
      <button class="btn" @tap="onSubmit">
        <span>确认申请</span>
      </button>
    </view>
  </view>
</template>

<script>
import { getMincode, applySourceCode } from '@/api/appSourceCode'
export default {
  name: 'SourceCodeForm',
  components: {},
  data() {
    return {
      region: [], // 所在地区选择
      prefix: '',
      minCode: '',
      startCode: '',
      endCode: '',
      formData: {
        applyQuantity: '',
        reciverName: '',
        reciverPhone: '',
        reciverAddr: '',
        region: '',
        detailAddress: ''
      },
      rules: {
        applyQuantity: {
          rules: [
            {
              required: true,
              errorMessage: '请输入申请数量'
            }
          ]
        },
        reciverName: {
          rules: [
            {
              required: true,
              errorMessage: '请输入收件人姓名'
            }
          ]
        },
        reciverPhone: {
          rules: [
            {
              required: true,
              errorMessage: '请输入收件人手机号'
            },
            {
              pattern:
                '^((\\d{3,4})-(\\d{7,8})|(\\d{3,4})-(\\d{7,8})-(\\d{1,4})|1[3456789]\\d{9})$', // 参考：https://c.runoob.com/front-end/854
              errorMessage:
                '请填写正确的电话或手机号码，如010-11112222、13344445555'
            }
          ]
        },
        region: {
          rules: [
            {
              require: true,
              errorMessage: '请选择所在地区'
            }
          ],
          validateTrigger: 'blur'
        },
        detailAddress: {
          rules: [
            {
              require: true,
              errorMessage: '请填写详细地址'
            }
          ],
          validateTrigger: 'blur'
        }
      }
    }
  },
  onShow() {
    this.query()
  },

  watch: {
    'formData.applyQuantity': {
      handler(newVal) {
        if (!newVal) return (this.endCode = this.startCode)
        let str = ''
        const MAX_LEN = 8,
          NEW_VAL_LEN = Number(newVal.length),
          VAL_LEN = (Number(this.minCode) + '').length - NEW_VAL_LEN

        const len = MAX_LEN - NEW_VAL_LEN - VAL_LEN
        for (let index = 0; index < len; index++) {
          str += '0'
        }
        str = str + (Number(newVal) - 1 + Number(this.minCode))
        this.endCode = this.prefix + str
      },
      immediate: true
    }
  },

  methods: {
    async query() {
      const res = await getMincode()

      if (res.code === '0') {
        const { prefix, minCode } = res.data
        this.prefix = prefix
        this.minCode = minCode
        this.startCode = prefix + minCode
        this.endCode = prefix + minCode
      }
    },
    bindRegionChange(e) {
      const region = e.detail.value.join('')

      this.$refs.form.setValue('region', region)
      this.formData.region = region
    },

    onSubmit() {
      this.$refs.form
        .submit()
        .then(async res => {
          res['reciverAddr'] = res.region + res.detailAddress
          const params = {
            startCode: Number(this.minCode),
            ...res
            // applyQuantity: 1000,
            // reciverName: 'zhangsan',
            // reciverAddr: '深圳市福田区中康路小学xxxx',
            // reciverPhone: '13366888555',
          }
          const data = await applySourceCode(params)
          console.log(data)
          if (data.code === '0')
            uni.navigateTo({ url: '/pages/appSourceCode/index' })
        })
        .catch(err => {
          console.log('表单错误信息：', err)
        })
    }
  },

  // 页面周期函数--监听页面加载
  onLoad() {},
  // 页面周期函数--监听页面初次渲染完成
  onReady() {},
  // 页面周期函数--监听页面显示(not-nvue)
  onHide() {},
  // 页面周期函数--监听页面卸载
  onUnload() {},
  // 页面处理函数--监听用户下拉动作
  onPullDownRefresh() {
    uni.stopPullDownRefresh()
  },
  // 页面处理函数--监听用户上拉触底
  onReachBottom() {}
  // 页面处理函数--监听页面滚动(not-nvue)
  /* onPageScroll(event) {}, */
  // 页面处理函数--用户点击右上角分享
  /* onShareAppMessage(options) {}, */
}
</script>

<style lang="scss">
.content {
  // display: flex;
  // flex-direction: column;
  // align-items: center;
  // justify-content: center;
  border-top: 1rpx solid #eee;
  border-bottom: 1rpx solid #eee;
  height: calc(100vh - 121rpx);
  overflow-y: auto;

  .msg {
    width: 100vw;
    line-height: 80rpx;
    background-color: #fafafa;
    font-family: PingFangSC-Regular;
    font-size: 22rpx;
    color: #8b8b8b;
    text-indent: 32rpx;

    .msg-code {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }

  .uni-input {
    height: 72rpx;
    line-height: 72rpx;
    text-align: right;
  }

  .common-input {
    border: none;
    text-align: right;

    .uni-easyinput__content-input {
      text-align: right;
    }
  }

  .mr40 .uni-easyinput__content-input {
    margin-right: 40rpx;
  }

  .add-brand-form {
    width: 90%;
    height: calc(100vh - 80rpx);

    .uni-forms-item {
      padding-left: 16rpx;
      padding-right: 33rpx;
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
}
</style>
