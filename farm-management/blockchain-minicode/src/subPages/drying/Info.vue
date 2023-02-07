<!--
 * @Description  : 新建烘干信息
 * @Autor        : Hemingway
 * @Date         : 2021-09-23 11:44:23
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 16:21:55
 * @FilePath     : \blockchain-minicode\src\subPages\drying\Info.vue
-->
<template>
  <view class="drying-info">
    <view class="form-wrapper">
      <uni-forms ref="form" v-model="formData" label-width="120" border>
        <uni-group>
          <uni-forms-item
            class="form-item"
            label="烘干结束日期"
            name="dryEndTime"
            required
          >
            <uni-datetime-picker
              type="date"
              :value="formData.dryEndTime"
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

          <!-- 收粮清单 -->
          <uni-forms-item
            class="form-item"
            label="收粮清单"
            name="feedListArray"
            required
          >
            <view
              @click="
                $Router.push({
                  path: '/selectGrainIn',
                  query: {
                    payload: queryGrainInPayload,
                    selectedList
                  }
                })
              "
            >
              <text
                :style="{
                  lineHeight: '36px',
                  marginRight: '4rpx',
                  color: !formData.feedListArray ? '#bbb' : ''
                }"
                >{{ selectedCarNumbers || '请选择' }}</text
              >
              <uni-icons type="arrowright" size="12"></uni-icons>
            </view>
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="干稻谷含水量(%)"
            label-width="150"
            name="dryRiceprimage"
            required
          >
            <uni-easyinput
              v-model="formData.dryRiceprimage"
              :trim="true"
              type="digit"
              :maxlength="4"
              :input-border="false"
              placeholder="请填写含水量"
              placeholder-style="color: #bbb;"
              @input="
                onInput('dryRiceprimage', { integerMaxLength: 2 }, $event)
              "
              @blur="onBlur('dryRiceprimage', $event)"
            />
          </uni-forms-item>
          <uni-forms-item
            v-show="dryRiceWeight"
            class="form-item"
            label="干稻谷重量(kg)"
          >
            <text
              :style="{
                lineHeight: '36px'
              }"
              >{{ dryRiceWeight }}</text
            >
          </uni-forms-item>
        </uni-group>
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
import BottomButton from '@/components/bottom-button'
import { addDrying } from '@/api/grain-steps'
import { stringifyDate } from '@/utils/tool'

export default {
  name: 'DryingInfo',

  components: {
    BottomButton
  },

  data() {
    return {
      riceFactoryName: '', // 一体化米厂名称
      queryGrainInPayload: {}, // 目标收粮信息 必要查询参数

      selectedList: [], // 所选收粮清单列表
      selectedCarNumbers: '', // 所选收粮清单车牌号

      // 表单
      formData: {
        dryEndTime: '' || this.formatDate(new Date()), // 烘干结束日期
        feedListArray: '', // 收粮清单 id集合
        dryRiceprimage: '' // 干稻谷含水量
      },
      // 表单域规则
      rules: {
        dryEndTime: {
          rules: [
            {
              required: true,
              errorMessage: '请选择烘干结束日期'
            }
          ],
          validateTrigger: 'blur'
        },
        feedListArray: {
          rules: [
            {
              required: true,
              errorMessage: '请选择收粮清单'
            }
          ],
          validateTrigger: 'blur'
        },
        dryRiceprimage: {
          rules: [
            {
              required: true,
              errorMessage: '请填写干稻谷含水量'
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

    // 干稻谷总重量
    dryRiceWeight() {
      if (this.selectedList.length === 0 || !this.formData.dryRiceprimage) {
        return 0
      }

      return this.selectedList.reduce((sum, current) => {
        return (
          parseFloat(sum) +
          ((1 - 0.01 * parseFloat(current.wetRicePrimage)) /
            (1 - 0.01 * parseFloat(this.formData.dryRiceprimage))) *
            current.wetRiceWeight
        ).toFixed(1)
      }, 0)
    },

    /**
     * @description: 表单按钮是否禁用
     * @return {Boolean}
     * @author: Hemingway
     */
    isDisabled() {
      const { feedListArray, dryRiceprimage } = this.formData

      return !feedListArray || !dryRiceprimage
    }
  },

  watch: {},

  created() {
    // 注册chooseGrainIn监听事件
    this.$bus.$on('chooseGrainIn', selectedList => {
      this.selectedList = selectedList
      this.selectedCarNumbers = selectedList
        .map(item => item.carNumber)
        .join(', ') // 用于展示选择项

      const idList = selectedList.map(item => item.feedListId).join(',') // 选择的清单id
      this.formData.feedListArray = idList
      this.$refs.form.setValue('feedListArray', idList) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
    })
  },

  mounted() {
    // this.$AppReady.then(() => {
    //   console.log('this.$Route.query = ', this.$Route.query)
    // })
    console.log('this.$Route.query = ', this.$Route.query)

    this.riceFactoryName = this.$Route.query.riceFactoryName
    this.queryGrainInPayload = this.$Route.query.payload
  },

  beforeDestroy() {
    // 释放chooseGrainIn监听事件
    this.$bus.$off('chooseGrainIn')
  },

  onReady() {
    // 需要在onReady中设置规则
    this.$refs.form.setRules(this.rules)
  },

  methods: {
    /**
     * @description: 烘干结束日期选择事件
     * @param {String} e
     * @author: Hemingway
     */
    onDateChange(e) {
      this.$refs.form.setValue('dryEndTime', e) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.dryEndTime = e
    },

    /**
     * @description: 输入事件 -> 限制整数长度
     * @param {String} field 域字段
     * @param {Number} integerMaxLength 整数部分最大长度
     * @param {String} event
     * @author: Hemingway
     */
    onInput(field, { integerMaxLength }, event) {
      this.$nextTick(() => {
        if (event.split('.')[0].length > integerMaxLength) {
          this.formData[field] = event.slice(0, -1)
        }
      })
    },

    /**
     * @description: 失焦事件
     * @param {String} field 域字段
     * @param {Object} event
     * @author: Hemingway
     */
    onBlur(field, event) {
      this.formData[field] = Number(
        parseFloat(event.detail.value || 0)
      ).toFixed(1)
    },

    /**
     * @description: 表单提交事件
     * @author: Hemingway
     */
    onSubmit() {
      uni.showModal({
        title: '提示',
        content: '确认信息填写无误？',
        success: async res => {
          if (res.confirm) {
            // 参数准备
            const payload = {
              dryEndTime: this.formData.dryEndTime,
              feedListArray: this.formData.feedListArray,
              dryRiceprimage: this.formData.dryRiceprimage,
              dryRiceWeight: this.dryRiceWeight
            }

            console.log('payload = ', payload)
            try {
              const { code } = await addDrying(payload)
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
                      this.$bus.$emit('createDrying')
                      this.$Router.back()
                    }, 1500)
                  }
                })
              }
            } catch (error) {
              console.log('烘干表单提交 error = ', error)
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

.drying-info {
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
}
</style>
