t<!--
 * @Description  : 新建包装信息
 * @Autor        : Hemingway
 * @Date         : 2021-09-23 11:44:23
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-11 15:49:55
 * @FilePath     : \blockchain-minicode\src\subPages\packing\Info.vue
-->
<template>
  <view class="packing-info">
    <view class="form-wrapper">
      <uni-forms ref="form" v-model="formData" label-width="120" border>
        <uni-group>
          <uni-forms-item
            class="form-item"
            label="包装结束日期"
            name="packingTime"
            required
          >
            <uni-datetime-picker
              type="date"
              :value="formData.packingTime"
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
            <text
              :style="{
                lineHeight: '36px'
              }"
              >{{ riceFactoryName }}</text
            >
          </uni-forms-item>
        </uni-group>

        <uni-group>
          <uni-forms-item
            class="form-item"
            label="商品名称"
            name="commodityName"
            required
          >
            <view
              @click="
                $Router.push({
                  name: 'selectCommodity',
                  query: {
                    enterpriseId
                  }
                })
              "
            >
              <text
                :style="{
                  lineHeight: '36px',
                  marginRight: '4rpx',
                  color: !formData.commodityName ? '#bbb' : ''
                }"
                >{{ formData.commodityName || '请选择' }}</text
              >
              <uni-icons type="arrowright" size="12"></uni-icons>
            </view>
          </uni-forms-item>

          <!-- 商品规格列表 -->
          <template v-if="packingInfos.length > 0">
            <uni-forms-item
              v-for="(item, index) in packingInfos"
              :key="index"
              class="form-item"
              :label="item.spec"
              label-width="150"
            >
              <uni-easyinput
                v-model="item.packingNumber"
                :trim="true"
                type="number"
                :maxlength="6"
                :input-border="false"
                placeholder="请填写件数"
                placeholder-style="color: #bbb;"
                @input="onInput(index, $event)"
                @blur="onBlur(index, $event)"
              />
            </uni-forms-item>
          </template>
        </uni-group>

        <view class="status-container">
          <!-- 标题行 -->
          <view class="title-row">
            <view>
              <text class="is-required">*</text>
              <text class="label-text">是否全部包装完?</text>
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
import { addPacking } from '@/api/grain-steps'
import { stringifyDate } from '@/utils/tool'

export default {
  name: 'PackingInfo',

  components: {
    Card,
    BottomButton
  },

  data() {
    return {
      riceFactoryName: '', // 一体化米厂名称
      enterpriseId: '', // 品牌商id
      processInfoId: '', // 操作id

      fold: true, // 是否折叠 状态选择说明
      desc: {}, // 说明内容
      statusList: [
        {
          name: '还有剩余未包装',
          value: '0',
          checked: false
        },
        {
          name: '已全部完成包装',
          value: '1',
          checked: false
        }
      ],

      // 包装产物
      packingInfos: [],

      // 表单
      formData: {
        packingTime: '' || this.formatDate(new Date()), // 包装结束日期
        commodityId: '', // 商品id
        commodityName: '', // 商品名称
        handleStatus: '' // 包装状态（1：包装完，0：有剩余未包装完）
      },
      // 表单域规则
      rules: {
        packingTime: {
          rules: [
            {
              required: true,
              errorMessage: '请选择包装结束日期'
            }
          ],
          validateTrigger: 'blur'
        },
        commodityName: {
          rules: [
            {
              required: true,
              errorMessage: '请选择商品名称'
            }
          ],
          validateTrigger: 'blur'
        },
        handleStatus: {
          rules: [
            {
              required: true,
              errorMessage: '请选择包装状态'
            }
          ],
          validateTrigger: 'blur'
        }
      }
    }
  },

  computed: {
    /**
     * @description: 表单按钮是否禁用
     * @return {Boolean}
     * @author: Hemingway
     */
    isDisabled() {
      let flag = !this.formData.handleStatus || !this.formData.commodityName
      this.packingInfos &&
        this.packingInfos.forEach(item => {
          if (!item.packingNumber) {
            flag = true
          }
        })
      return flag
    }
  },

  created() {
    // 注册chooseCommodity监听事件
    this.$bus.$on(
      'chooseCommodity',
      ({ commodityId, commodityName, netContent }) => {
        this.formData.commodityId = commodityId
        this.formData.commodityName = commodityName

        // 包装产物
        this.packingInfos = netContent.map(item => {
          if (!item.id) {
            item.id = new Date().getTime() + Math.round(Math.random() * 1000)
          }
          item.spec = `${item.weightValue}kg/${item.weightUnit}`
          item.packingNumber = ''
          return item
        })
      }
    )
  },

  mounted() {
    console.log('this.$Route.query = ', this.$Route.query)
    this.riceFactoryName = this.$Route.query.riceFactoryName
    this.enterpriseId = this.$Route.query.enterpriseId
    this.processInfoId = this.$Route.query.processInfoId
    this.desc = {
      hiddenHeader: true,
      items: this.$Route.query.desc.map(item => ({
        text: `${item.label}: ${item.value}`
      }))
    }
  },

  beforeDestroy() {
    // 释放chooseCommodity监听事件
    this.$bus.$off('chooseCommodity')
  },

  onReady() {
    // 需要在onReady中设置规则
    this.$refs.form.setRules(this.rules)
  },

  methods: {
    /**
     * @description: 包装结束日期选择事件
     * @param {String} e
     * @author: Hemingway
     */
    onDateChange(e) {
      this.$refs.form.setValue('packingTime', e) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.packingTime = e
    },

    /**
     * @description: 输入事件 -> 模拟 view to data
     * @param {Number} index 在数组中的索引
     * @param {String} event
     * @author: Hemingway
     */
    onInput(index, event) {
      console.log('onInput', event)

      this.$set(
        this.packingInfos,
        index,
        Object.assign(this.packingInfos[index], {
          packingNumber: event
        })
      ) // 【注】动态渲染的uni-easyinput是个傻子，需要手动设置值
    },

    /**
     * @description: 失焦事件
     * @param {Number} index 在数组中的索引
     * @param {Object} event
     * @author: Hemingway
     */
    onBlur(index, event) {
      if (event.detail.value === '') {
        this.$set(
          this.packingInfos,
          index,
          Object.assign(this.packingInfos[index], {
            packingNumber: '0' // 设置默认值
          })
        ) // 【注】动态渲染的uni-easyinput是个傻子，需要手动设置值
      }
    },

    /**
     * @description: 包装状态选择事件
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
              packingTime: this.formData.packingTime,
              commodityName: this.formData.commodityName,
              commodityId: this.formData.commodityId,
              packingInfos: JSON.stringify(
                this.packingInfos.filter(item => item.packingNumber !== '0')
              ), // 包装规格相关信息
              packingId: this.processInfoId, // 操作id
              handleStatus: this.formData.handleStatus
            }

            console.log('payload = ', payload)
            try {
              const { code } = await addPacking(payload)
              if (code === '0') {
                uni.showToast({
                  title: '提交成功',
                  icon: 'success',
                  duration: 2000,
                  complete: () => {
                    setTimeout(() => {
                      this.$bus.$emit('createPacking')
                      this.$Router.back()
                    }, 1500)
                  }
                })
              }
            } catch (error) {
              console.log('包装表单提交 error = ', error)
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

.packing-info {
  ::v-deep .uni-group {
    margin-top: 16rpx !important;
  }
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

  // 包装状态
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
