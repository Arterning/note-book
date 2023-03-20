<!--
 * @Description  : 新建加工信息
 * @Autor        : Hemingway
 * @Date         : 2021-09-23 11:44:23
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 10:07:01
 * @FilePath     : \blockchain-minicode\src\subPages\machining\Info.vue
-->
<template>
  <view class="machining-info">
    <view class="form-wrapper">
      <uni-forms ref="form" v-model="formData" label-width="120" border>
        <uni-group>
          <uni-forms-item
            class="form-item"
            label="加工结束日期"
            name="processTime"
            required
          >
            <uni-datetime-picker
              type="date"
              :value="formData.processTime"
              start="2021-1-1"
              :end="formatDate(new Date())"
              @change="onDateChange"
            />
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="一体化米厂"
            name="riceFactoryName"
          >
            <!-- 当登录用户为品牌商时使用data - riceFactoryName，当登录用户为米厂时使用role - enterpriseName -->
            <text
              :style="{
                lineHeight: '36px'
              }"
              >{{ riceFactoryName || role.enterpriseName }}</text
            >
          </uni-forms-item>
        </uni-group>

        <uni-group class="plain-cell">
          <uni-forms-item
            class="form-item"
            label="副产品(50kg/包)"
            label-width="180"
            required
          >
          </uni-forms-item>
        </uni-group>
        <uni-group class="horizon-cell">
          <uni-forms-item
            v-for="(item, index) in byProductList"
            :key="index"
            class="form-item"
            :label="item.packingTypeName"
            label-width="150"
          >
            <uni-easyinput
              v-model="item.packingNumber"
              :trim="true"
              type="digit"
              :maxlength="8"
              :input-border="false"
              placeholder="请填写件数"
              placeholder-style="color: #bbb;"
              @input="onInput(index, { integerMaxLength: 6 }, $event)"
              @blur="onBlur(index, $event)"
            />
          </uni-forms-item>
        </uni-group>

        <view class="status-container">
          <!-- 标题行 -->
          <view class="title-row">
            <view>
              <text class="is-required">*</text>
              <text class="label-text">是否全部加工完?</text>
              <text class="label-text--mini">(请确认)</text>
            </view>
            <view class="icon-wrapper" @click="fold = !fold">
              <template v-if="fold">
                <text class="text">展开明细</text
                ><uni-icons type="arrowdown" size="8"></uni-icons>
              </template>
              <template v-else>
                <text class="text">收起</text>
                <uni-icons type="arrowup" size="8"></uni-icons>
              </template>
            </view>
          </view>
          <!-- 说明行 -->
          <view v-show="!fold">
            <card :data="desc"> </card>
          </view>
          <!-- 选择行 -->
          <uni-group class="horizon-cell no-label">
            <uni-forms-item class="form-item" name="handleStatus">
              <radio-group class="radio-wrapper" @change="bindRadioChange">
                <label
                  v-for="item in statusList"
                  :key="item.value"
                  class="label"
                >
                  <radio
                    :value="item.value"
                    :checked="item.checked"
                    style="transform: scale(0.7);"
                  />
                  {{ item.name }}
                </label>
              </radio-group>
            </uni-forms-item>
          </uni-group>
        </view>
      </uni-forms>
    </view>

    <bottom-button
      text="提交"
      :is-disabled="isDisabled"
      @click="onSubmit"
    ></bottom-button>
  </view>
</template>

<script>
import Card from '@/components/card'
import BottomButton from '@/components/bottom-button'
import { addMachining } from '@/api/grain-steps'
import { stringifyDate } from '@/utils/tool'

export default {
  name: 'MachiningInfo',

  components: {
    Card,
    BottomButton
  },

  data() {
    return {
      riceFactoryName: '', // 一体化米厂名称
      processInfoId: '', // 加工id

      // 副产品列表
      byProductList: [
        {
          packingType: '1',
          packingUnit: '',
          packingTypeName: '皮糠',
          packingNumber: ''
        },
        {
          packingType: '2',
          packingUnit: '',
          packingTypeName: '小碎',
          packingNumber: ''
        },
        {
          packingType: '3',
          packingUnit: '',
          packingTypeName: '大碎',
          packingNumber: ''
        },
        {
          packingType: '4',
          packingUnit: '',
          packingTypeName: '二选',
          packingNumber: ''
        },
        {
          packingType: '5',
          packingUnit: '',
          packingTypeName: '三选',
          packingNumber: ''
        },
        {
          packingType: '6',
          packingUnit: '',
          packingTypeName: '选白米',
          packingNumber: ''
        }
      ],

      fold: true, // 是否折叠 状态选择说明
      desc: {}, // 说明内容
      statusList: [
        {
          name: '还有剩余未脱谷',
          value: '2',
          checked: false
        },
        {
          name: '已全部完成脱谷',
          value: '1',
          checked: false
        }
      ],

      // 表单
      formData: {
        processTime: '' || this.formatDate(new Date()), // 加工结束日期
        handleStatus: '' // 加工状态（1：加工完，2：有剩余未加工完）
      },
      // 表单域规则
      rules: {
        processTime: {
          rules: [
            {
              required: true,
              errorMessage: '请选择加工结束日期'
            }
          ],
          validateTrigger: 'blur'
        },
        handleStatus: {
          rules: [
            {
              required: true,
              errorMessage: '请选择加工状态'
            }
          ],
          validateTrigger: 'blur'
        }
      }
    }
  },

  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    },

    /**
     * @description: 表单按钮是否禁用
     * @return {Boolean}
     * @author: Hemingway
     */
    isDisabled() {
      let flag = !this.formData.handleStatus
      this.byProductList.forEach(item => {
        if (!item.packingNumber) {
          flag = true
        }
      })
      return flag
    }
  },

  mounted() {
    console.log('this.$Route.query = ', this.$Route.query)
    this.riceFactoryName = this.$Route.query.riceFactoryName
    this.processInfoId = this.$Route.query.processInfoId
    this.desc = {
      hiddenHeader: true,
      items: this.$Route.query.desc.map(item => ({
        text: `${item.label}: ${item.value}`
      }))
    }
  },

  onReady() {
    // 需要在onReady中设置规则
    this.$refs.form.setRules(this.rules)
  },

  methods: {
    /**
     * @description: 加工结束日期选择事件
     * @param {String} e
     * @author: Hemingway
     */
    onDateChange(e) {
      this.$refs.form.setValue('processTime', e) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.processTime = e
    },

    /**
     * @description: 输入事件 -> 限制整数长度
     * @param {Number} index 在数组中的索引
     * @param {Number} integerMaxLength 整数部分最大长度
     * @param {String} event
     * @author: Hemingway
     */
    onInput(index, { integerMaxLength }, event) {
      this.$nextTick(() => {
        if (event.split('.')[0].length > integerMaxLength) {
          this.byProductList[index].packingNumber = event.slice(0, -1)
        }
      })
    },

    /**
     * @description: 失焦事件
     * @param {Number} index 在数组中的索引
     * @param {Object} event
     * @author: Hemingway
     */
    onBlur(index, event) {
      this.byProductList[index].packingNumber = Number(
        parseFloat(event.detail.value || 0)
      ).toFixed(1)
    },

    /**
     * @description: 加工状态选择事件
     * @param {Object} e
     * @author: Hemingway
     */
    bindRadioChange(e) {
      const status = e.detail.value

      this.$refs.form.setValue('handleStatus', status) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.handleStatus = status
    },

    /**
     * @description: 表单提交事件
     * @author: Hemingway
     */
    async onSubmit() {
      uni.showModal({
        title: '提示',
        content: '确认信息填写无误？',
        success: async res => {
          if (res.confirm) {
            // 参数准备
            const payload = {
              processTime: this.formData.processTime,
              packingInfos: JSON.stringify(this.byProductList),
              processInfoId: this.processInfoId,
              handleStatus: this.formData.handleStatus
            }

            console.log('payload = ', payload)
            try {
              const { code } = await addMachining(payload)
              if (code === '0') {
                uni.showToast({
                  title: '提交成功',
                  icon: 'success',
                  duration: 2000,
                  complete: () => {
                    setTimeout(() => {
                      // uni.navigateBack()
                      // uni.switchTab({
                      //   url: '/pages/home/index'
                      // })
                      this.$bus.$emit('createMachining')
                      this.$Router.back()
                    }, 1500)
                  }
                })
              }
            } catch (error) {
              console.log('加工表单提交 error = ', error)
            }
          }
        }
      })
    },

    /**
     * @description: 格式化日期
     * @param {*} date
     * @author: Hemingway
     */
    formatDate(date) {
      return stringifyDate(date)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/formMixin.scss';

.machining-info {
  // 日期选择器
  ::v-deep .uni-date {
    padding-left: 150rpx;

    .uni-date__x-input {
      height: 68rpx;
      font-size: 24rpx;
    }

    .uni-date__icon-clear {
      display: none;
    }
  }

  // 副产品
  .horizon-cell ::v-deep .uni-forms-item {
    .uni-forms-item__label {
      display: unset;
    }

    .uni-easyinput .uni-easyinput__content {
      text-align: right;
    }
  }

  // 加工状态
  .status-container {
    margin-top: 16rpx;
    padding: 0 16rpx;
    background: #fff;

    .title-row {
      display: flex;
      justify-content: space-between;
      height: 104rpx;
      line-height: 104rpx;
      border-bottom: 1rpx solid #eee;

      .is-required {
        color: #f00;
        font-weight: bold;
      }

      .label-text {
        font-size: 14px;
        color: #333;
      }

      .label-text--mini {
        font-size: 24rpx;
        font-weight: lighter;
        color: #8b8b8b;
      }

      .icon-wrapper {
        display: flex;
        align-items: center;
        color: #8b8b8b;

        .text {
          margin-right: 4rpx;
          font-size: 24rpx;
          font-weight: lighter;
        }
      }
    }

    .no-label {
      ::v-deep .uni-forms-item {
        .uni-forms-item__label {
          display: none;
        }
      }

      ::v-deep .uni-group__content {
        padding: 0;
      }

      // 单选框
      ::v-deep .uni-forms-item__content {
        display: flex;
        align-items: center;
        justify-content: flex-end;

        .radio-wrapper {
          display: flex;
          flex-wrap: wrap;
        }

        .label:not(:last-of-type) {
          margin-right: 32rpx;
        }
      }
    }
  }
}
</style>
