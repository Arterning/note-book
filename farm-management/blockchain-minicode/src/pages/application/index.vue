<!--
 * @Description  : （品牌商和合作商）入驻申请页
 * @Autor        : Hemingway
 * @Date         : 2021-06-20 17:01:00
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-25 09:10:21
 * @FilePath     : \blockchain-minicode\src\pages\application\index.vue
-->
<template>
  <view class="application">
    <view class="form-wrapper">
      <uni-forms ref="form" v-model="formData" label-width="120" border>
        <view class="tip">温馨提示：以下资料用于溯源平台入驻资质审核。</view>
        <uni-group>
          <uni-forms-item
            class="form-item enterprise"
            label="企业类型"
            name="enterpriseType"
            required
          >
            <radio-group class="radio-wrapper" @change="bindEnterpriseChange">
              <label
                v-for="(item, index) in settleTypeList"
                :key="index"
                class="label"
              >
                <radio
                  :value="item.code"
                  :checked="
                    formData.enterpriseType.split(',').includes(item.code)
                  "
                  style="transform: scale(0.7);"
                  :disabled="enterpriseIds.length > 0"
                />{{ item.text }}</label
              >
            </radio-group>
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="企业名称"
            name="enterpriseName"
            required
          >
            <uni-easyinput
              v-model="formData.enterpriseName"
              :trim="true"
              type="text"
              :maxlength="30"
              :input-border="false"
              placeholder="请填写企业名称"
              placeholder-style="color: #bbb;"
              :disabled="enterpriseIds.length > 0"
            />
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="法人代表"
            name="legalPerson"
            required
          >
            <uni-easyinput
              v-model="formData.legalPerson"
              :trim="true"
              type="text"
              :maxlength="10"
              :input-border="false"
              placeholder="请填写企业法人姓名"
              placeholder-style="color: #bbb;"
              :disabled="enterpriseIds.length > 0"
            />
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="联系电话"
            name="phoneNum"
            required
          >
            <uni-easyinput
              v-model="formData.phoneNum"
              :trim="true"
              type="text"
              :maxlength="16"
              :input-border="false"
              placeholder="请填写企业联系电话"
              placeholder-style="color: #bbb;"
              :disabled="enterpriseIds.length > 0"
            />
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="统一代码"
            name="creditCode"
            required
          >
            <uni-easyinput
              v-model="formData.creditCode"
              :trim="true"
              type="text"
              :maxlength="18"
              :input-border="false"
              placeholder="请填写统一社会信用代码"
              placeholder-style="color: #bbb;"
              :disabled="enterpriseIds.length > 0"
            />
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="所在地区"
            name="region"
            required
          >
            <picker
              mode="region"
              :value="region"
              :disabled="enterpriseIds.length > 0"
              @change="bindRegionChange"
            >
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
              :disabled="enterpriseIds.length > 0"
            />
          </uni-forms-item>
        </uni-group>

        <uni-group class="plain-cell">
          <uni-forms-item
            class="form-item"
            label="法人代表身份证（人像面）"
            label-width="180"
            required
          >
          </uni-forms-item>
        </uni-group>
        <uni-group class="horizon-cell">
          <uni-forms-item class="form-item" name="frontSide">
            <image-upload
              :can-refresh="!(handleType === 'append')"
              :max="1"
              file-class="idCardFront"
              :list="preview.frontSide"
              @uploaded="onUpload('frontSide', $event)"
              @updated="onUpdate('frontSide', $event)"
            />
          </uni-forms-item>
        </uni-group>

        <uni-group class="plain-cell">
          <uni-forms-item
            class="form-item"
            label="法人代表身份证（国徽面）"
            :label-width="180"
            required
          >
          </uni-forms-item>
        </uni-group>
        <uni-group class="horizon-cell">
          <uni-forms-item class="form-item" name="backSide">
            <image-upload
              :can-refresh="!(handleType === 'append')"
              :max="1"
              file-class="idCardReverse"
              :list="preview.backSide"
              @uploaded="onUpload('backSide', $event)"
              @updated="onUpdate('backSide', $event)"
            />
          </uni-forms-item>
        </uni-group>

        <uni-group class="plain-cell">
          <uni-forms-item class="form-item" label="企业营业执照" required>
          </uni-forms-item>
        </uni-group>
        <uni-group class="horizon-cell">
          <uni-forms-item class="form-item" name="businessLicense">
            <image-upload
              :can-refresh="!(handleType === 'append')"
              :max="1"
              file-class="businessLicense"
              :list="preview.businessLicense"
              @uploaded="onUpload('businessLicense', $event)"
              @updated="onUpdate('businessLicense', $event)"
            />
          </uni-forms-item>
        </uni-group>

        <!-- 品牌商资质 start -->
        <view v-if="isBrandOwner">
          <uni-group class="plain-cell">
            <uni-forms-item class="form-item" label="品牌商标注册证" required>
            </uni-forms-item>
          </uni-group>
          <uni-group class="horizon-cell">
            <uni-forms-item class="form-item" name="brandLicense">
              <image-upload
                :can-refresh="!(handleType === 'append' && isBrandOwner)"
                :max="1"
                file-class="registrationCertificate"
                :list="preview.brandLicense"
                @uploaded="onUpload('brandLicense', $event)"
                @updated="onUpdate('brandLicense', $event)"
              />
            </uni-forms-item>
          </uni-group>

          <uni-group class="plain-cell">
            <uni-forms-item class="form-item" label="食品经营许可证" required>
            </uni-forms-item
          ></uni-group>
          <uni-group class="horizon-cell">
            <uni-forms-item class="form-item" name="grainBusinessLicense">
              <image-upload
                :can-refresh="!(handleType === 'append' && isBrandOwner)"
                :max="1"
                file-class="foodBusinessLicense"
                :list="preview.grainBusinessLicense"
                @uploaded="onUpload('grainBusinessLicense', $event)"
                @updated="onUpdate('grainBusinessLicense', $event)"
              />
            </uni-forms-item>
          </uni-group>
        </view>
        <!-- 品牌商资质 end -->

        <!-- 合作商资质 start -->
        <view v-if="isFactory">
          <uni-group class="plain-cell">
            <uni-forms-item class="form-item" label="食品生产许可证" required>
            </uni-forms-item>
          </uni-group>
          <uni-group class="horizon-cell">
            <uni-forms-item class="form-item" name="produceLicence">
              <image-upload
                :max="1"
                file-class="riceProcessingLicense"
                :can-refresh="!(handleType === 'append' && isFactory)"
                :list="preview.produceLicence"
                @uploaded="onUpload('produceLicence', $event)"
                @updated="onUpdate('produceLicence', $event)"
              />
            </uni-forms-item>
          </uni-group>

          <uni-group class="plain-cell">
            <uni-forms-item class="form-item" label="工厂操作环境" required>
            </uni-forms-item>
          </uni-group>
          <uni-group class="horizon-cell">
            <uni-forms-item class="form-item" name="operatingEnvironment">
              <view style="display: flex;">
                <view class="image-wrapper">
                  <image-upload
                    :can-refresh="!(handleType === 'append' && isFactory)"
                    :max="1"
                    file-class="dryEnvironment"
                    :list="preview.operatingEnvironment_1"
                    @uploaded="onUpload('operatingEnvironment_1', $event)"
                    @updated="onUpdate('operatingEnvironment_1', $event)"
                  />
                  <text class="text">烘干</text>
                </view>
                <view class="image-wrapper">
                  <image-upload
                    :can-refresh="!(handleType === 'append' && isFactory)"
                    :max="1"
                    file-class="storageEnvironment"
                    :list="preview.operatingEnvironment_2"
                    @uploaded="onUpload('operatingEnvironment_2', $event)"
                    @updated="onUpdate('operatingEnvironment_2', $event)"
                  />
                  <text class="text">仓储</text>
                </view>
                <view class="image-wrapper">
                  <image-upload
                    :can-refresh="!(handleType === 'append' && isFactory)"
                    :max="1"
                    file-class="processingEnvironment"
                    :list="preview.operatingEnvironment_3"
                    @uploaded="onUpload('operatingEnvironment_3', $event)"
                    @updated="onUpdate('operatingEnvironment_3', $event)"
                  />
                  <text class="text">加工</text>
                </view>
                <view class="image-wrapper">
                  <image-upload
                    :can-refresh="!(handleType === 'append' && isFactory)"
                    :max="1"
                    file-class="packageEnvironment"
                    :list="preview.operatingEnvironment_4"
                    @uploaded="onUpload('operatingEnvironment_4', $event)"
                    @updated="onUpdate('operatingEnvironment_4', $event)"
                  />
                  <text class="text">包装</text>
                </view>
              </view>
            </uni-forms-item>
          </uni-group>
        </view>
        <!-- 合作商资质 end -->
      </uni-forms>
    </view>

    <bottom-button
      v-if="!(enterpriseIds.length > 0)"
      text="提交"
      :is-disabled="isDisabled"
      @click="onSubmit"
    ></bottom-button>
    <view v-else class="notice-wrapper">
      <view>
        <uni-icons
          type="checkbox"
          color="#00c853"
          size="14"
          style="margin-right: 8rpx;"
        ></uni-icons
        ><text>您已完成企业入驻！</text>
      </view>
      <text class="notice">(注：一家企业只能入驻一种企业类型)</text>
    </view>
  </view>
</template>

<script>
import ImageUpload from '@/components/image-upload'
import BottomButton from '@/components/bottom-button'
import { queryApplicationDetail, submitApplication } from '@/api/user'

// 图片上传的fileClass，与字段field定义的映射关系
const fieldNameMap = {
  idCardFront: 'frontSide', // 身份证 人像面
  idCardReverse: 'backSide', // 身份证 国徽面
  businessLicense: 'businessLicense', // 企业营业执照
  registrationCertificate: 'brandLicense', // 品牌商标注册证 - 品牌商
  foodBusinessLicense: 'grainBusinessLicense', // 食品经营许可证 - 品牌商
  riceProcessingLicense: 'produceLicence', // 食品生产许可证 - 烘干厂/仓储厂/加工厂/一体化厂
  dryEnvironment: 'operatingEnvironment_1', // 工厂操作环境 - 烘干厂/一体化厂
  storageEnvironment: 'operatingEnvironment_2', // 工厂操作环境 - 仓储厂/一体化厂
  processingEnvironment: 'operatingEnvironment_3', // 工厂操作环境 - 加工厂/一体化厂
  packageEnvironment: 'operatingEnvironment_4' // 工厂操作环境 - 加工厂/一体化厂
}

export default {
  name: 'Application',

  components: {
    ImageUpload,
    BottomButton
  },

  data() {
    return {
      handleType: '', // 操作类型
      recordId: '', // 记录id（当修改表单时有用）
      region: [], // 所在地区选择

      // 表单
      formData: {
        enterpriseType: '', // 企业类型: 2-品牌商；6-烘干加工一体化厂；
        enterpriseName: '', // 企业名称
        legalPerson: '', // 法人代表
        phoneNum: '', // 企业联系电话
        creditCode: '', // 统一代码
        region: '', // 所在地区
        detailAddress: '', // 详细地址
        frontSide: '', // 身份证 人像面
        backSide: '', // 身份证 国徽面
        businessLicense: '', // 企业营业执照
        brandLicense: '', // 品牌商标注册证 - 品牌商
        grainBusinessLicense: '', // 食品经营许可证 - 品牌商
        produceLicence: '', // 食品生产许可证 - 烘干厂/仓储厂/加工厂/一体化厂
        operatingEnvironment: '' // 工厂操作环境 - 烘干厂/仓储厂/加工厂/一体化厂
      },
      // 表单域规则
      rules: {
        enterpriseType: {
          rules: [
            {
              required: true,
              errorMessage: '请选择企业类型'
            }
          ],
          validateTrigger: 'blur'
        },
        enterpriseName: {
          rules: [
            {
              required: true,
              errorMessage: '请填写企业名称'
            }
          ],
          validateTrigger: 'blur'
        },

        legalPerson: {
          rules: [
            {
              required: true,
              errorMessage: '请填写法人代表'
            }
          ],
          validateTrigger: 'blur'
        },
        phoneNum: {
          rules: [
            {
              required: true,
              errorMessage: '请填写企业联系电话'
            },
            {
              pattern:
                '^((\\d{3,4})-(\\d{7,8})|(\\d{3,4})-(\\d{7,8})-(\\d{1,4})|1[3456789]\\d{9})$', // 参考：https://c.runoob.com/front-end/854
              errorMessage:
                '请填写正确的电话或手机号码，如010-11112222、13344445555'
            }
          ],
          validateTrigger: 'blur'
        },
        creditCode: {
          rules: [
            {
              required: true,
              errorMessage: '请填写统一代码（15或18位）'
            },
            {
              pattern:
                '^([0-9A-HJ-NPQRTUWXY]{2}\\d{6}[0-9A-HJ-NPQRTUWXY]{10}|[1-9]\\d{14})$', // 参考：http://www.qilin668.com/a/5e217a2e6ee1fs4.html
              errorMessage: '请填写正确的统一代码'
            }
          ],
          validateTrigger: 'blur'
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
        },
        frontSide: {
          rules: [
            {
              required: true,
              errorMessage: '请上传法人代表身份证（人像面）'
            }
          ],
          validateTrigger: 'blur'
        },
        backSide: {
          rules: [
            {
              required: true,
              errorMessage: '请上传法人代表身份证（国徽面）'
            }
          ],
          validateTrigger: 'blur'
        },
        businessLicense: {
          rules: [
            {
              required: true,
              errorMessage: '请上传企业营业执照'
            }
          ],
          validateTrigger: 'blur'
        },
        brandLicense: {
          rules: [
            {
              required: true,
              errorMessage: '请上传品牌商标注册证'
            }
          ],
          validateTrigger: 'blur'
        },
        grainBusinessLicense: {
          rules: [
            {
              required: true,
              errorMessage: '请上传食品经营许可证'
            }
          ],
          validateTrigger: 'blur'
        },
        produceLicence: {
          rules: [
            {
              required: true,
              errorMessage: '请上传食品生产许可证'
            }
          ],
          validateTrigger: 'blur'
        },
        operatingEnvironment: {
          rules: [
            {
              required: true,
              errorMessage: '请上传工厂操作环境图片'
            },
            {
              validateFunction: (rule, value, data, callback) => {
                if (value.split(',').length < this.needLength) {
                  callback('请完善工厂操作环境图片')
                }
                return true
              }
            }
          ],
          validateTrigger: 'blur'
        }
      },
      // 图片预览
      preview: {
        frontSide: [],
        backSide: [],
        businessLicense: [],
        brandLicense: [],
        grainBusinessLicense: [],
        produceLicence: [],
        operatingEnvironment_1: [], // 烘干
        operatingEnvironment_2: [], // 仓储
        operatingEnvironment_3: [], // 加工
        operatingEnvironment_4: [] // 包装
      },
      // 图片id
      id: {
        frontSide: [],
        backSide: [],
        businessLicense: [],
        brandLicense: [],
        grainBusinessLicense: [],
        produceLicence: [],
        operatingEnvironment_1: [], // 烘干
        operatingEnvironment_2: [], // 仓储
        operatingEnvironment_3: [], // 加工
        operatingEnvironment_4: [] // 包装
      },
      needLength: 0, // 工厂操作环境需要上传的图片数量
      delList: [] // 无效图片id
    }
  },

  computed: {
    /**
     * @description: 企业类型map
     * @author: Hemingway
     */
    enterpriseMap() {
      return this.$store.state.user.dictMap.enterprise_type
    },

    /**
     * @description: 根据企业类型map转list
     * @author: Hemingway
     */
    settleTypeList() {
      let arr = []
      Object.keys(this.enterpriseMap).forEach(key => {
        arr.push({ code: key, text: this.enterpriseMap[key] })
      })
      return arr
    },

    /**
     * @description: 上传图片时fileClass的map
     * @author: Hemingway
     */
    fileClassMap() {
      return this.$store.state.user.dictMap.fileClass
    },

    /**
     * @description: 基于fileClass的扩展map（根据fileClass寻找fieldName）
     * @author: Hemingway
     */
    fileClassMapPlus() {
      if (this.fileClassMap) {
        const map = {}
        for (let key in this.fileClassMap) {
          map[key] = {
            text: this.fileClassMap[key],
            fieldName: fieldNameMap[key]
          }
        }
        return map
      }

      return {}
    },

    // 获取用户信息
    userInfo() {
      return this.$store.state.user.userInfo
    },

    // 获取用户角色对应的企业类型id集合
    enterpriseIds() {
      return (
        this.userInfo && this.userInfo.roles.filter(role => role.enterpriseType)
      )
    },

    /**
     * @description: 企业类型是否是品牌商
     * @return {Boolean}
     * @author: Hemingway
     */
    isBrandOwner() {
      return this.formData.enterpriseType.split(',').includes('2')
    },

    /**
     * @description: 企业类型是否是一体化米厂
     * @return {Boolean}
     * @author: Hemingway
     */
    isFactory() {
      return this.formData.enterpriseType.split(',').includes('6')
    },

    /**
     * @description: 统计工厂环境 idList
     * @return {Array}
     * @author: Hemingway
     */
    concatEnvironmentIdList() {
      let arr = []

      if (this.isFactory) {
        // 一体化厂
        this.needLength = 4 // eslint-disable-line
        arr = [
          ...this.id.operatingEnvironment_1,
          ...this.id.operatingEnvironment_2,
          ...this.id.operatingEnvironment_3,
          ...this.id.operatingEnvironment_4
        ]

        // 更新工厂环境字段
        if (this.$refs.form.setValue) {
          this.$refs.form.setValue('operatingEnvironment', arr.join(',')) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
        }
        this.formData.operatingEnvironment = arr.join(',') // eslint-disable-line
      }

      return arr
    },

    /**
     * @description: 提交表单时，图片id集合
     * @return {Array}
     * @author: Hemingway
     */
    effectiveIdList() {
      const {
        frontSide,
        backSide,
        businessLicense,
        brandLicense,
        produceLicence,
        grainBusinessLicense
      } = this.id
      let arr = [...frontSide, ...backSide, ...businessLicense]

      if (this.isBrandOwner) {
        // 品牌商
        arr = arr.concat(brandLicense, grainBusinessLicense)
      }
      if (this.isFactory) {
        // 合作商
        arr = arr.concat(produceLicence, this.concatEnvironmentIdList)
      }

      return arr
    },

    /**
     * @description: 表单按钮是否禁用
     * @return {Boolean}
     * @author: Hemingway
     */
    isDisabled() {
      const {
        enterpriseType,
        enterpriseName,
        legalPerson,
        phoneNum,
        creditCode,
        region,
        detailAddress,
        frontSide,
        backSide,
        businessLicense,
        brandLicense, // 品牌商标注册证 - 品牌商
        grainBusinessLicense, // 食品经营许可证 - 品牌商
        produceLicence, // 食品生产许可证 - 烘干厂/仓储厂/加工厂/一体化厂
        operatingEnvironment // 工厂操作环境 - 烘干厂/仓储厂/加工厂/一体化厂
      } = this.formData

      let flag =
        !enterpriseType ||
        !enterpriseName ||
        !legalPerson ||
        !phoneNum ||
        !creditCode ||
        !region ||
        !detailAddress ||
        !frontSide ||
        !backSide ||
        !businessLicense

      if (this.isBrandOwner) {
        // 品牌商
        flag = flag || !brandLicense || !grainBusinessLicense
      }
      if (this.isFactory) {
        // 一体化米厂
        flag =
          flag ||
          !produceLicence ||
          !operatingEnvironment ||
          this.concatEnvironmentIdList.length < this.needLength
      }

      return flag
    }
  },

  watch: {
    handleType(value) {
      if (value === 'update') {
        uni.setNavigationBarTitle({
          title: '修改入驻申请'
        })
      } else if (value === 'append') {
        // 目前，不支持入驻多种企业类型
        // uni.setNavigationBarTitle({
        //   title: '添加入驻申请'
        // })
      }
    }
  },

  mounted() {
    this.handleType = this.$Route.query.handleType
    this.recordId = this.$Route.query.id
    if (this.recordId) {
      this.queryDetailAndFillForm(this.recordId)
    }
  },

  onReady() {
    // 需要在onReady中设置规则
    this.$refs.form.setRules(this.rules)
  },

  methods: {
    /**
     * @description: 企业类型选择事件
     * @param {Object} e
     * @author: Hemingway
     */
    bindEnterpriseChange(e) {
      // const enterprise = e.detail.value.join(',') // 目前，不支持入驻多种企业类型
      const enterprise = e.detail.value

      this.$refs.form.setValue('enterpriseType', enterprise) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.enterpriseType = enterprise
    },

    /**
     * @description: 所在地区选择事件
     * @param {Object} e
     * @author: Hemingway
     */
    bindRegionChange(e) {
      const region = e.detail.value.join('')

      this.$refs.form.setValue('region', region) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData.region = region
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

      if (field.includes('_')) {
        // 工厂环境

        this.$refs.form.setValue(
          field.slice(0, -2),
          this.concatEnvironmentIdList.join(',')
        ) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
        this.formData[field.slice(0, -2)] = this.concatEnvironmentIdList.join(
          ','
        )
      } else {
        // 非工厂环境

        this.$refs.form.setValue(field, this.id[field].join(',')) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
        this.formData[field] = this.id[field].join(',')
      }
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
      this.delList.push(this.id[field][index]) // 无效id数组
      this.id[field].splice(index, 1, ...idList) // 图片id数组

      this.$refs.form.setValue(field, this.id[field].join(',')) // 参见：https://ext.dcloud.net.cn/plugin?name=uni-forms
      this.formData[field] = this.id[field].join(',')
    },

    /**
     * @description: 表单提交事件
     * @author: Hemingway
     */
    async onSubmit() {
      try {
        const {
          enterpriseType,
          enterpriseName,
          legalPerson,
          phoneNum,
          creditCode,
          region,
          detailAddress
        } = await this.$refs.form.validate()

        console.log(
          'this.$refs.form.validate() = ',
          enterpriseType,
          enterpriseName,
          legalPerson,
          phoneNum,
          creditCode,
          region,
          detailAddress
        )
        // 参数准备
        const payload = {
          type: enterpriseType,
          name: enterpriseName,
          legalPerson,
          officePhone: phoneNum,
          uuid: creditCode,
          region,
          detailAddress
        }

        console.log('payload = ', payload)

        Object.assign(payload, {
          effective: this.effectiveIdList.join(','),
          invalid: this.delList.join(',')
        }) // 附件
        if (this.recordId && this.handleType === 'update') {
          payload.id = this.recordId
        } // 提交类型：新增还是修改

        const { code } = await submitApplication(payload)
        if (code === '0') {
          uni.showToast({
            title: '请等待系统审核',
            icon: 'success',
            duration: 2000,
            complete: () => {
              setTimeout(() => {
                this.$Router.push({
                  name: 'reviewProgress'
                })
              }, 1500)
            }
          })
        }
        // 提交表单
      } catch (error) {
        console.log('入驻申请表单提交 error = ', error)
      }
    },

    /**
     * @description: 根据id查询详情并填充表单
     * @param {String} id
     * @author: Hemingway
     */
    async queryDetailAndFillForm(id) {
      try {
        const { code, enterpriseChangeDto } = await queryApplicationDetail({
          id
        })
        if (code === '0') {
          const {
            type,
            name,
            legalPerson,
            officePhone,
            uuid,
            region,
            detailAddress,
            attachments
          } = enterpriseChangeDto

          // 企业类型
          if (this.handleType === 'update') {
            this.bindEnterpriseChange({
              detail: {
                value: type
              }
            })
          } else if (this.handleType === 'append') {
            this.bindEnterpriseChange({
              detail: {
                // value: this.settleTypeList.map(item => item.code) // 目前，不支持入驻多种企业类型
                value: type
              }
            })
          }

          this.formData.enterpriseName = name // 企业名称
          this.formData.legalPerson = legalPerson // 法人代表
          this.formData.phoneNum = officePhone // 联系电话
          this.formData.creditCode = uuid // 统一代码
          // 所在地区
          this.$refs.form.setValue('region', region)
          this.formData.region = region
          this.formData.region = region // 所在地区
          this.formData.detailAddress = detailAddress // 详细地址

          // 图片回填：模拟图片上传事件
          for (let key in attachments) {
            const field = this.fileClassMapPlus[key].fieldName
            const idList = attachments[key].map(item => item.id)
            const previewList = idList.map(id => this.getImageUrl(id))
            this.onUpload(field, { previewList, idList })
          }
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
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/formMixin.scss';

.application {
  // 企业选择
  .enterprise ::v-deep .uni-forms-item__content {
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

  // 工厂环境
  .image-wrapper {
    display: flex;
    flex-direction: column;
    align-items: center;

    .text {
      margin-top: 16rpx;
      font-size: 24rpx;
      color: #8b8b8b;
    }
  }

  // 提示
  .notice-wrapper {
    padding: 60rpx;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    font-size: 28rpx;

    .notice {
      margin-top: 16rpx;
      font-size: 24rpx;
      font-weight: lighter;
      color: #8b8b8b;
    }
  }

  .tip {
    color: #8b8b8b;
    font-size: 24rpx;
    margin-top: 20rpx;
    margin-bottom: 20rpx;
    margin-left: 20rpx;
  }
}
</style>
