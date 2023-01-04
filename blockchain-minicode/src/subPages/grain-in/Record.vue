<!--
 * @Description  : 新建/编辑 收粮信息
 * @Autor        : Hemingway
 * @Date         : 2021-09-06 14:25:39
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-18 18:30:29
 * @FilePath     : \blockchain-minicode\src\subPages\grain-in\Record.vue
-->
<template>
  <view class="record">
    <view class="form-wrapper">
      <uni-forms ref="form" v-model="formData" label-width="120" border>
        <uni-group>
          <uni-forms-item
            class="form-item"
            label="收粮日期"
            name="reachingDate"
            required
          >
            <uni-datetime-picker
              type="date"
              :value="formData.reachingDate"
              start="2021-1-1"
              :end="formatDate(new Date())"
              @change="onDateChange"
            />
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="品牌商"
            name="enterpriseName"
            required
          >
            <view
              @click="
                $Router.push({
                  name: 'selectEnterprise'
                })
              "
            >
              <text
                :style="{
                  lineHeight: '36px',
                  marginRight: '4rpx',
                  color: !formData.enterpriseName ? '#bbb' : ''
                }"
                >{{ formData.enterpriseName || '请选择' }}</text
              >
              <uni-icons type="arrowright" size="12"></uni-icons>
            </view>
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
              >{{ role.enterpriseName }}</text
            >
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="来源农场"
            name="farmName"
            required
          >
            <view
              @click="
                $Router.push({
                  name: 'selectFarm'
                })
              "
            >
              <text
                :style="{
                  lineHeight: '36px',
                  marginRight: '4rpx',
                  color: !formData.farmName ? '#bbb' : ''
                }"
                >{{ formData.farmName || '请选择' }}</text
              >
              <uni-icons type="arrowright" size="12"></uni-icons>
            </view>
          </uni-forms-item>

          <uni-forms-item
            v-if="formData.varietyName"
            class="form-item"
            label="品种"
          >
            <text
              :style="{
                lineHeight: '36px'
              }"
              >{{ formData.varietyName }}</text
            >
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="车牌号"
            label-width="150"
            name="carNumber"
            required
          >
            <uni-easyinput
              v-model="formData.carNumber"
              :trim="true"
              type="text"
              :maxlength="18"
              :input-border="false"
              placeholder="请填写送货车牌号"
              placeholder-style="color: #bbb;"
            />
          </uni-forms-item>
          <uni-forms-item
            class="form-item"
            label="湿稻谷重量(kg)"
            label-width="150"
            name="wetRiceWeight"
            required
          >
            <uni-easyinput
              v-model="formData.wetRiceWeight"
              :trim="true"
              type="digit"
              :maxlength="9"
              :input-border="false"
              placeholder="请填写重量"
              placeholder-style="color: #bbb;"
              @input="onInput('wetRiceWeight', { integerMaxLength: 7 }, $event)"
              @blur="onBlur('wetRiceWeight', $event)"
            />
          </uni-forms-item>
          <uni-forms-item
            class="form-item"
            label="湿稻谷含水量(%)"
            label-width="150"
            name="wetRicePrimage"
            required
          >
            <uni-easyinput
              v-model="formData.wetRicePrimage"
              :trim="true"
              type="digit"
              :maxlength="4"
              :input-border="false"
              placeholder="请填写含水量"
              placeholder-style="color: #bbb;"
              @input="
                onInput('wetRicePrimage', { integerMaxLength: 2 }, $event)
              "
              @blur="onBlur('wetRicePrimage', $event)"
            />
          </uni-forms-item>
        </uni-group>

        <uni-group class="plain-cell">
          <uni-forms-item
            class="form-item"
            label="照片"
            label-width="180"
            required
          >
          </uni-forms-item>
        </uni-group>
        <uni-group class="horizon-cell">
          <uni-forms-item class="form-item" name="images">
            <image-upload
              can-refresh
              :max="8"
              file-class="feedList"
              :list="preview.images"
              @uploaded="onUpload('images', $event)"
              @updated="onUpdate('images', $event)"
            />
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
import ImageUpload from '@/components/image-upload'
import BottomButton from '@/components/bottom-button'
import {
  queryGrainInDetail,
  addGrainIn,
  updateGrainIn
} from '@/api/grain-steps'
import { stringifyDate } from '@/utils/tool'

export default {
  name: 'Record',

  components: {
    ImageUpload,
    BottomButton
  },

  data() {
    return {
      recordId: '',

      // 表单
      formData: {
        reachingDate: '' || this.formatDate(new Date()), // 收粮日期
        enterpriseName: '', // 品牌商名称
        enterpriseId: '', // 品牌商id
        farmName: '', // 农场名称
        farmId: '', // 农场id
        varietyName: '', // 品种id
        varietyId: '', // 品种id
        phone: '', // 农场管理员电话
        carNumber: '', // 车牌号
        wetRiceWeight: '', // 湿稻谷重量
        wetRicePrimage: '', // 湿稻谷含水量
        images: '' // 照片
      },
      // 表单域规则
      rules: {
        reachingDate: {
          rules: [
            {
              required: true,
              errorMessage: '请选择收粮日期'
            }
          ],
          validateTrigger: 'blur'
        },
        enterpriseName: {
          rules: [
            {
              required: true,
              errorMessage: '请选择品牌商'
            }
          ],
          validateTrigger: 'blur'
        },
        farmName: {
          rules: [
            {
              required: true,
              errorMessage: '请选择来源农场'
            }
          ],
          validateTrigger: 'blur'
        },
        carNumber: {
          rules: [
            {
              required: true,
              errorMessage: '请填写送货车牌号'
            },
            {
              pattern:
                '^([京津晋冀蒙辽吉黑沪苏浙皖闽赣鲁豫鄂湘粤桂琼渝川贵云藏陕甘青宁新][ABCDEFGHJKLMNPQRSTUVWXY][1-9DF][1-9ABCDEFGHJKLMNPQRSTUVWXYZ]\\d{3}[1-9DF]|[京津晋冀蒙辽吉黑沪苏浙皖闽赣鲁豫鄂湘粤桂琼渝川贵云藏陕甘青宁新][ABCDEFGHJKLMNPQRSTUVWXY][\\dABCDEFGHJKLNMxPQRSTUVWXYZ]{5})$', // 参考：http://www.qilin668.com/a/5e1edf12ce73e33.html
              errorMessage: '请填写正确的车牌号'
            }
          ],
          validateTrigger: 'blur'
        },
        wetRiceWeight: {
          rules: [
            {
              required: true,
              errorMessage: '请填写湿稻谷重量'
            }
          ],
          validateTrigger: 'blur'
        },
        wetRicePrimage: {
          rules: [
            {
              required: true,
              errorMessage: '请填写湿稻谷含水量'
            }
          ],
          validateTrigger: 'blur'
        },
        images: {
          rules: [
            {
              required: true,
              errorMessage: '请上传照片'
            }
          ],
          validateTrigger: 'blur'
        }
      },
      // 图片预览
      preview: {
        images: []
      },
      // 图片id
      id: {
        images: []
      },
      delList: [] // 无效图片id
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
      const {
        enterpriseName,
        farmName,
        carNumber,
        wetRiceWeight,
        wetRicePrimage,
        images
      } = this.formData

      return (
        !enterpriseName ||
        !farmName ||
        !carNumber ||
        !wetRiceWeight ||
        !wetRicePrimage ||
        !images
      )
    }
  },

  watch: {},

  created() {
    // 注册chooseFarm监听事件
    this.$bus.$on(
      'chooseFarm',
      ({ farmName, farmId, varietyName, varietyId, phone }) => {
        this.formData.farmName = farmName
        this.formData.farmId = farmId
        this.formData.varietyName = varietyName
        this.formData.varietyId = varietyId
        this.formData.phone = phone
        this.$refs.form.setValue('farmName', farmName) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      }
    )
    // 注册chooseBrandOwer监听事件
    this.$bus.$on('chooseBrandOwer', ({ enterpriseId, enterpriseName }) => {
      console.log('chooseBrandOwer = ', enterpriseId, enterpriseName)
      this.formData.enterpriseId = enterpriseId
      this.formData.enterpriseName = enterpriseName
      this.$refs.form.setValue('enterpriseName', enterpriseName) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
    })
  },

  mounted() {
    const id = this.$Route.query.id
    console.log('id = ', id)
    if (id) {
      console.log('修改收粮清单')
      this.queryDetailAndFillForm(id)
      this.recordId = id
    }
  },

  beforeDestroy() {
    // 释放chooseFarm监听事件
    this.$bus.$off('chooseFarm')
    // 释放chooseBrandOwer监听事件
    this.$bus.$off('chooseBrandOwer')
  },

  onReady() {
    // 需要在onReady中设置规则
    this.$refs.form.setRules(this.rules)
  },

  methods: {
    /**
     * @description: 收粮日期选择事件
     * @param {String} e
     * @author: Hemingway
     */
    onDateChange(e) {
      this.$refs.form.setValue('reachingDate', e) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.reachingDate = e
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
     * @description: 图片上传事件
     * @param {String} field 域字段
     * @param {Object} event
     * @author: Hemingway
     */
    onUpload(field, event) {
      console.log(field, event)

      const { previewList, idList } = event
      this.preview[field] = [...this.preview[field], ...previewList] // 图片预览数组
      this.id[field] = [...this.id[field], ...idList] // 图片id数组

      this.$refs.form.setValue(field, this.id[field].join(',')) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData[field] = this.id[field].join(',')
    },

    /**
     * @description: 图片更新事件
     * @param {String} field 域字段
     * @param {Object} event
     * @author: Hemingway
     */
    onUpdate(field, event) {
      console.log(field, event)

      const { index, previewList, idList } = event
      this.preview[field].splice(index, 1, ...previewList) // 图片预览数组
      this.id[field].splice(index, 1, ...idList) // 图片id数组
      this.delList.push(this.id[field][index]) // 无效id数组

      this.$refs.form.setValue(field, this.id[field].join(',')) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData[field] = this.id[field].join(',')
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
            try {
              const {
                enterpriseName,
                farmName,
                carNumber,
                wetRiceWeight,
                wetRicePrimage,
                images
              } = await this.$refs.form.validate()

              console.log(
                'this.$refs.form.validate() = ',
                enterpriseName,
                farmName,
                carNumber,
                wetRiceWeight,
                wetRicePrimage,
                images
              )

              // 参数准备
              const payload = {
                reachingDate: this.formData.reachingDate,
                enterpriseName,
                enterpriseId: this.formData.enterpriseId,
                riceFactoryName: this.role.enterpriseName,
                riceFactoryId: this.role.enterpriseId,
                farmName: this.formData.farmName,
                farmId: this.formData.farmId,
                riceVarietyName: this.formData.varietyName,
                riceVarietyId: this.formData.varietyId,
                phone: this.formData.phone,
                carNumber,
                wetRiceWeight,
                wetRicePrimage,
                imageIds: images // 附件id集合（待后端优化接口）
              }

              console.log('payload = ', payload)

              // Object.assign(payload, {
              //   effective: this.effectiveIdList.join(','),
              //   invalid: this.delList.join(',')
              // }) // 附件

              // 提交类型：新增还是修改
              let api = null
              if (this.recordId) {
                payload.feedListId = this.recordId
                api = updateGrainIn
              } else {
                api = addGrainIn
              }

              const { code } = await api(payload)
              if (code === '0') {
                uni.showToast({
                  title: '请等待农场确认',
                  icon: 'success',
                  duration: 2000,
                  complete: () => {
                    setTimeout(() => {
                      // uni.navigateBack()
                      // uni.switchTab({
                      //   url: '/pages/home/index'
                      // })
                      this.$bus.$emit('createGrainIn')
                      this.$Router.back()
                    }, 1500)
                  }
                })
              }
              // 提交表单
            } catch (error) {
              console.log('收粮清单表单提交 error = ', error)
            }
          }
        }
      })
    },

    /**
     * @description: 根据id查询详情并填充表单
     * @param {String} id
     * @author: Hemingway
     */
    async queryDetailAndFillForm(id) {
      try {
        const { code, feedListDto } = await queryGrainInDetail({
          feedListId: id
        })
        if (code === '0') {
          const {
            reachingDate, // 收粮日期
            enterpriseName, // 品牌商名称
            enterpriseId, // 品牌商id
            farmName, // 农场名称
            farmId, // 农场id
            riceVarietyName, // 品种id
            riceVarietyId, // 品种id
            phone, // 农场管理员电话
            carNumber, // 车牌号
            wetRiceWeight, // 湿稻谷重量
            wetRicePrimage, // 湿稻谷含水量
            imageIds // 照片
          } = feedListDto

          this.formData.reachingDate = this.formatDate(reachingDate)
          this.formData.enterpriseName = enterpriseName
          this.formData.enterpriseId = enterpriseId
          this.formData.farmName = farmName
          this.formData.farmId = farmId
          this.formData.varietyName = riceVarietyName
          this.formData.varietyId = riceVarietyId
          this.formData.phone = phone
          this.formData.carNumber = carNumber
          this.formData.wetRiceWeight = wetRiceWeight
          this.formData.wetRicePrimage = wetRicePrimage

          // 图片回填：模拟图片上传事件
          const idList = imageIds.split(',')
          const previewList = idList.map(id => this.getImageUrl(id))
          this.onUpload('images', { previewList, idList })
        }
      } catch (error) {
        console.log('入驻审核详情查询 error = ', error)
      }
    },

    /**
     * @description: 获取图片下载url
     * @param {String} id
     * @return {String} url
     * @author: Hemingway
     */
    getImageUrl(id) {
      return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
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

.record {
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
