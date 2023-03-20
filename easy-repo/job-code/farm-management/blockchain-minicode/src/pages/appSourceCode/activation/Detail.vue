<template>
  <view class="content">
    <view class="msg"
      >以下信息将同步展示到二维码溯源页面，您可以检查后再提交</view
    >
    <view class="forms-view">
      <uni-forms
        ref="form"
        class="add-brand-form"
        :value="formData"
        label-width="130"
        border
      >
        <uni-forms-item label="包装结束时间" name="">
          <view>{{ pzjcDate }}</view>
        </uni-forms-item>

        <uni-forms-item label="来源农场" name="">
          <view>{{ sourceFarmName }}</view>
        </uni-forms-item>

        <uni-forms-item label="商品名称" name="">
          <view>{{ commodityName }}</view>
        </uni-forms-item>

        <uni-forms-item label="包装规格" name="">
          <view style="word-break: break-all;">{{ normsForMachPack }}</view>
        </uni-forms-item>
        <uni-forms-item required label="该批损坏二维码" name="bad">
          <view class="spec">
            <uni-easyinput
              v-model="formData.bad"
              :input-border="false"
              class="common-input picker-item"
              type="number"
              placeholder="请输入"
            /><text class="unit">个</text>
          </view>

          <!-- <uni-easyinput
              v-model="formData.bad"
              :trim="true"
              type="digit"
              :input-border="false"
            /> -->
        </uni-forms-item>

        <!-- <uni-forms-item label="该批损坏二维码" name="">
          <view class="spec">
            <uni-easyinput
              v-model="formData.packingNumber"
              :input-border="false"
              class="common-input picker-item"
              type="number"
              placeholder="请输入"
              @input="onInput($event, 'packingNumber', item)"
            /><text class="unit">个</text>
          </view>
        </uni-forms-item> -->

        <uni-forms-item required label="绑定溯源码" name="codeStart">
        </uni-forms-item>

        <uni-forms-item
          style="display: none;"
          required
          label="绑定溯源码"
          name="codeEnd"
        >
        </uni-forms-item>
        <view class="trace">
          <text>{{ prefix }}</text>
          <uni-easyinput
            v-model="codeStart"
            :input-border="false"
            class="common-input picker-item"
            type="text"
            placeholder="请输入8位编号"
            @blur="changeCode"
          />
          <view>—</view>
          <text>{{ prefix }}</text>
          <uni-easyinput
            v-model="codeEnd"
            :input-border="false"
            class="common-input picker-item"
            type="text"
            placeholder="请输入8位编号"
          />
        </view>
        <view style="color: red; font-size: 12px;"
          >建议从最小码开始，当前最小码为：00000001</view
        >

        <view class="suyuan">
          <uni-data-checkbox
            v-model="formData.suyuan"
            selected-text-color="#00c853"
            selected-color="#00c853"
            :localdata="range"
            @change="changeSuyuan"
          ></uni-data-checkbox>
          <view class="text" @tap="jump">点击查看溯源信息</view>
        </view>
        <uni-forms-item class="suyuan-item" required label="" name="suyuan">
        </uni-forms-item>
      </uni-forms>
    </view>
    <view class="add-brand">
      <button class="btn" @tap="onSubmit">
        <span>提交并绑定</span>
      </button>
    </view>
  </view>
</template>

<script>
import { getMincode, bindSourceCode } from '@/api/appSourceCode'
import { stringifyDate } from '@/utils/tool'
export default {
  name: 'Packing',

  data() {
    return {
      target: null,
      id: '',
      pzjcDate: '',
      sourceFarmName: '',
      commodityName: '',
      normsForMachPack: '',
      range: [{ value: 0, text: '溯源信息已完整' }],
      prefix: '',
      codeStart: null,
      codeEnd: null,
      suyuan: null,
      formData: {
        bad: 0,
        codeStart: null,
        codeEnd: null,
        suyuan: null
      },
      rules: {
        bad: {
          rules: [
            {
              required: true,
              errorMessage: '请输入损坏溯源码个数'
            }
          ]
        },
        codeStart: {
          rules: [
            {
              required: true,
              errorMessage: '请输入起始溯源码'
            }
          ]
        },
        codeEnd: {
          rules: [
            {
              required: true,
              errorMessage: '请输入结束溯源码'
            }
          ]
        },
        suyuan: {
          rules: [
            {
              required: true,
              errorMessage: '请勾选溯源信息完整'
            }
          ]
        }
      }
    }
  },
  onReady() {
    this.$refs.form.setRules(this.rules)
  },

  onLoad() {
    const formData = this.$Route.query.formData
    this.target = this.$Route.query.formData
    this.pzjcDate = stringifyDate(formData.packDate)
    this.sourceFarmName = formData.sourceFarmName
    this.commodityName = formData.commodityName
    this.normsForMachPack = formData.normsForMachPack
    this.id = formData.id
    // this.formData = { ...this.formData, ...formData }
  },

  mounted() {
    this.getPrefix()
  },

  methods: {
    async getPrefix() {
      const res = await getMincode()
      if (res.code === '0') {
        const { prefix } = res.data
        this.prefix = prefix
      }
    },
    changeCode() {
      if (this.codeStart.length === 8) {
        let num = 0
        const arr = this.normsForMachPack.split(';')
        arr.map(el => {
          num += parseInt(el.split('*')[1])
        })
        num += parseInt(this.formData.bad)
        const end = parseInt(this.codeStart) - 1 + num
        const start = this.codeStart.substr(0, (end + '').length)
        this.codeEnd = start + '' + end
      }
      if (this.codeStart !== null && this.codeEnd !== null) {
        this.$refs.form.setValue('codeStart', this.codeStart)
        this.$refs.form.setValue('codeEnd', this.codeEnd)
      }
    },
    //     onInput(e, key) {
    //       this[key] = e
    // debugger

    //     },
    jump() {
      uni.navigateTo({
        url: `/subPages/myRice/ProcessFlow?info=${encodeURIComponent(
          JSON.stringify({
            batchId: this.id
          })
        )}`
      })
    },
    changeSuyuan() {
      this.$refs.form.setValue('suyuan', 0)
    },
    onSubmit() {
      this.$refs.form
        .validate()
        .then(async () => {
          uni.showModal({
            title: '提示',
            content: '确认信息填写无误？',
            success: async res => {
              if (res.confirm) {
                const params = {
                  id: this.id,
                  bindSrcCodeList: [
                    {
                      codeStartId: '',
                      codeEndId: '',
                      codeStart: this.prefix + this.codeStart,
                      codeEnd: this.prefix + this.codeEnd
                    }
                  ]
                }
                const data = await bindSourceCode(params)
                if (data.code === '0') {
                  this.$Router.push({
                    name: 'CodeActivation'
                  })
                }
                // 参数准备
              }
            }
          })
        })
        .catch(err => {
          console.log('表单错误信息：', err)
        })
    }
  }
}
</script>

<style lang="scss">
.content {
  .suyuan {
    position: relative;
    margin-top: 20rpx;

    .text {
      position: absolute;
      right: 0;
      top: 0;
      color: #00c853;
      font-size: 14px;
    }
  }

  ::v-deep .suyuan-item .uni-forms-item--border {
    border: unset;

    .uni-forms-item__label {
      display: none;
    }

    .uni-error-message {
      padding-left: 0 !important;
      text-align: left !important;
      top: -48px !important;
    }
  }

  .forms-view {
    background: #fff;
    width: 100%;
  }

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

  .weight {
    line-height: 72rpx;
  }

  .qrcode {
    width: 160rpx;
    height: 80rpx;
    text-align: center;
    line-height: 80rpx;
    background-color: #fafafa;
    font-family: PingFangSC-Regular;
    font-size: 24rpx;
    color: #00c853;
    position: absolute;
    right: 0;
  }

  .uni-input {
    height: 72rpx;
    line-height: 72rpx;
    text-align: right;
  }

  .picker-item {
    display: inline-block;
  }

  .common-input {
    border: none;
    text-align: right;

    .uni-easyinput__content-input {
      text-align: right;
    }
  }

  .trace {
    /* display: flex;
    align-items: center; */

    text {
      font-family: PingFangSC-Regular;
      font-size: 28rpx;
      color: #101010;
      width: 100rpx;
    }

    ::v-deep .uni-easyinput__content-input {
      text-align: left;
    }

    .common-input {
      flex: 1;
    }
  }

  .add-brand-form {
    display: inline-block;
    margin-left: 5%;
    width: 90%;
    height: calc(100vh - 80rpx);
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
