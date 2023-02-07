<!--
 * @Description  : 验证手机及设置密码（注册或重置密码时）
 * @Autor        : Hemingway
 * @Date         : 2021-06-23 10:18:41
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-11 19:32:52
 * @FilePath     : \blockchain-minicode\src\pages\login\LoginExtra.vue
-->
<template>
  <div class="login-extra">
    <view class="container">
      <uni-steps
        :options="[{ title: '验证手机' }, { title: '设置密码' }]"
        :active="stepIndex"
      ></uni-steps>

      <!-- 步骤一：验证手机表单 -->
      <uni-forms
        v-if="isFirstStep"
        ref="form"
        v-model="formData"
        :rules="rules"
        border
      >
        <view :class="{ basic: true, active: active === 'phoneNum' }">
          <uni-forms-item name="phoneNum">
            <uni-easyinput
              v-model="formData.phoneNum"
              :trim="true"
              type="number"
              :maxlength="11"
              :input-border="false"
              :clearable="false"
              placeholder="手机号码"
              @focus="active = 'phoneNum'"
              @blur="active = ''"
            />
          </uni-forms-item>
        </view>

        <view :class="{ basic: true, active: active === 'code' }">
          <uni-forms-item name="code">
            <uni-easyinput
              v-model="formData.code"
              :trim="true"
              type="number"
              :maxlength="6"
              :input-border="false"
              :clearable="false"
              placeholder="验证码"
              @focus="active = 'code'"
              @blur="active = ''"
            >
              <template #right>
                <view v-if="!inCountdown" class="code" @click="getCode"
                  ><uni-icons
                    type="paperplane"
                    color="#00c853"
                    size="12"
                  ></uni-icons>
                  <text style="margin-left: 4rpx;">获取验证码</text></view
                >
                <view v-else
                  ><uni-countdown
                    :show-day="false"
                    :second="30"
                    @timeup="inCountdown = false"
                  ></uni-countdown
                ></view>
              </template>
            </uni-easyinput>
          </uni-forms-item>
        </view>
      </uni-forms>

      <!-- 步骤二：重置密码表单 -->
      <uni-forms v-else ref="form2" v-model="formData2" :rules="rules2" border>
        <view
          v-if="isRegister"
          :class="{ basic: true, active: active === 'name' }"
        >
          <uni-forms-item name="name">
            <uni-easyinput
              v-model="formData2.name"
              :trim="true"
              type="text"
              :maxlength="10"
              :input-border="false"
              :clearable="false"
              placeholder="姓名"
              @focus="active = 'name'"
              @blur="active = ''"
            />
          </uni-forms-item>
        </view>

        <view :class="{ basic: true, active: active === 'password' }">
          <uni-forms-item name="password">
            <uni-easyinput
              v-model="formData2.password"
              :trim="true"
              type="password"
              :maxlength="20"
              :input-border="false"
              :clearable="false"
              placeholder="密码"
              @focus="active = 'password'"
              @blur="active = ''"
            />
          </uni-forms-item>
        </view>

        <view :class="{ basic: true, active: active === 'password2' }">
          <uni-forms-item name="password2">
            <uni-easyinput
              v-model="formData2.password2"
              :trim="true"
              type="password"
              :maxlength="20"
              :input-border="false"
              :clearable="false"
              placeholder="确认密码"
              @focus="active = 'password2'"
              @blur="active = ''"
            >
            </uni-easyinput>
          </uni-forms-item>
        </view>
      </uni-forms>
    </view>

    <view class="btn-wrapper">
      <button
        type="primary"
        class="btn"
        :disabled="isDisabled"
        @click="onClick"
      >
        {{ isFirstStep ? '下一步' : '确认' }}
      </button>
    </view>
  </div>
</template>

<script>
import mixin from './mixin'
import { verifyCode, register, resetPassword } from '@/api/user'

export default {
  name: 'LoginExtra',

  mixins: [mixin], // getCode

  data() {
    return {
      // 表单
      formData: {
        phoneNum: '', // 手机号
        code: '' // 验证码
      },
      // 表单域规则
      rules: {
        phoneNum: {
          rules: [
            {
              required: true,
              errorMessage: '请输入手机号码'
            },
            {
              minLength: 11,
              maxLength: 11,
              errorMessage: '请确认手机号码长度（11位）'
            },
            {
              pattern: '^1[3456789]\\d{9}$',
              errorMessage: '请输入正确的手机号码'
            }
          ],
          validateTrigger: 'blur'
        },
        code: {
          rules: [
            {
              required: true,
              errorMessage: '请输入验证码'
            },
            {
              minLength: 6,
              maxLength: 6,
              errorMessage: '请确认验证码长度（6位）'
            }
          ],
          validateTrigger: 'blur'
        }
      },
      // 表单2
      formData2: {
        name: '', // 姓名
        password: '', // 密码
        password2: '' // 确认密码
      },
      // 表单域2规则
      rules2: {
        name: {
          rules: [
            {
              required: true,
              errorMessage: '请输入姓名'
            },
            {
              minLength: 2,
              maxLength: 10,
              errorMessage: '请确认姓名长度（2-10位）'
            }
          ],
          validateTrigger: 'blur'
        },
        password: {
          rules: [
            {
              required: true,
              errorMessage: '请输入密码'
            },
            {
              minLength: 8,
              maxLength: 20,
              errorMessage: '请确认登录密码长度（8-20位）'
            },
            {
              pattern:
                '^.*(?=.{8,20})(?=.*\\d)(?=.*[a-zA-Z])(?=.*[!@#$%^&*()_+.]).*$',
              errorMessage:
                '密码强度不够，应包含数字、字母及指定特殊字符（如!@#$%^&*()_+.）'
            }
          ],
          validateTrigger: 'blur'
        },
        password2: {
          rules: [
            {
              required: true,
              errorMessage: '请输入确认密码'
            },
            {
              minLength: 8,
              maxLength: 20,
              errorMessage: '请确认登录密码长度（8-20位）'
            },
            {
              pattern:
                '^.*(?=.{8,20})(?=.*\\d)(?=.*[a-zA-Z])(?=.*[!@#$%^&*()_+.]).*$',
              errorMessage:
                '密码强度不够，应包含数字、字母及指定特殊字符（如!@#$%^&*()_+.）'
            }
          ],
          validateTrigger: 'blur'
        }
      },
      smsType: '', // 短信类型
      active: '', // 激活的输入框name字段
      isRegister: true, // true-注册；false-重置密码（根据smsType进行判断）
      inCountdown: false, // true-倒计时中；false-非倒计时中
      isFirstStep: true, // 是否处在步骤一
      certifiCate: '' // 存储步骤一结果
    }
  },

  computed: {
    /**
     * @description: 步骤索引
     * @author: Hemingway
     */
    stepIndex() {
      return this.isFirstStep ? 0 : 1
    },

    /**
     * @description: 表单按钮是否禁用
     * @author: Hemingway
     */
    isDisabled() {
      if (this.isFirstStep) {
        // 步骤一：验证手机表单

        return !this.formData.phoneNum || !this.formData.code
      } else {
        // 步骤二：设置密码表单

        if (this.isRegister) {
          // 注册
          return (
            !this.formData2.name ||
            !this.formData2.password ||
            !this.formData2.password2
          )
        } else {
          // 忘记密码
          return !this.formData2.password || !this.formData2.password2
        }
      }
    }
  },

  watch: {
    smsType(value) {
      if (value === '02') {
        uni.setNavigationBarTitle({
          title: '重置密码'
        })

        this.isRegister = false
      }
    }
  },

  created() {},

  mounted() {
    this.smsType = this.$Route.query.smsType
  },

  methods: {
    /**
     * @description: 按钮点击事件（含验证手机、设置密码）
     * @author: Hemingway
     */
    async onClick() {
      if (this.isFirstStep) {
        // 步骤一：验证手机表单

        try {
          const { phoneNum, code } = await this.$refs.form.validate()

          // 校验验证码
          const res = await verifyCode({
            phone: phoneNum,
            code,
            msgType: this.smsType
          })
          console.log('res = ', res)
          if (res.code === '0') {
            this.certifiCate = res.certifiCate

            // 进入下一步骤
            this.isFirstStep = false
          }
        } catch (error) {
          console.log('验证手机 error = ', error)
        }
      } else {
        // 步骤二：设置密码表单

        if (this.formData2.password === this.formData2.password2) {
          try {
            // 表单校验和参数准备
            const { name, password } = await this.$refs.form2.validate()
            const secretPassword = await this.encrypt(password)

            // 入参
            const paramObj = {
              phone: this.formData.phoneNum
            }
            this.isRegister
              ? Object.assign(paramObj, {
                  certificate: this.certifiCate,
                  name,
                  password: secretPassword
                })
              : Object.assign(paramObj, {
                  certifiCate: this.certifiCate,
                  name,
                  newPassWord: secretPassword
                })

            // 表单提交
            const api = this.isRegister ? register : resetPassword
            const res = await api(paramObj)
            if (res.code === '0') {
              uni.showToast({
                title: `${this.isRegister ? '注册成功' : '重置密码成功'}`,
                icon: 'success',
                duration: 2000,
                complete: () => {
                  setTimeout(() => {
                    this.$Router.replaceAll({ name: 'login' })
                  }, 1500)
                }
              })
            }
          } catch (error) {
            console.log('设置密码 error = ', error)
          }
        } else {
          uni.showToast({
            title: '两次输入的密码不一致，请确认',
            icon: 'none',
            duration: 2000
          })
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.login-extra {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #fff;

  .container {
    margin: 80rpx 80rpx 0;
    position: relative;

    ::v-deep .uni-steps {
      margin-bottom: 80rpx;
    }

    ::v-deep .uni-forms-item {
      border-top: none;

      .uni-error-message {
        text-align: left;
      }
    }

    ::v-deep .uni-easyinput {
      .uni-easyinput__content {
        text-align: left;

        .uni-easyinput__content-input {
          padding: 0 !important;
        }
      }
    }

    // 发送验证码
    .code {
      font-size: 24rpx;
      color: #8b8b8b;
    }

    .basic {
      position: relative;
    }

    .basic::after {
      content: '';
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      border-bottom: 1rpx solid #eee;
      transform: scaleY(0.5);
    }

    .active::after {
      border-color: $theme-color;
    }
  }

  // 按钮组
  .btn-wrapper {
    margin: 120rpx 40rpx 0;

    .btn {
      height: 88rpx;
      line-height: 88rpx;
      background-color: $uni-bg-color-btn;
      font-size: 32rpx;
      border-radius: 44rpx;
    }

    .btn::after {
      display: none;
    }
  }
}
</style>
