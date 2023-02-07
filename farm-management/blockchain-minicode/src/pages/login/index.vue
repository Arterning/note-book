<!--
 * @Description  : 登录页
 * @Autor        : Hemingway
 * @Date         : 2021-06-22 11:05:33
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-18 16:49:34
 * @FilePath     : \blockchain-minicode\src\pages\login\index.vue
-->
<template>
  <view class="login">
    <view class="container">
      <view class="title">{{ title }}</view>

      <uni-forms ref="form" v-model="formData" :rules="rules" border>
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

        <!-- 密码（账号登录） -->
        <view
          v-if="loginTypeFlag"
          :class="{ basic: true, active: active === 'password' }"
        >
          <uni-forms-item name="password">
            <uni-easyinput
              v-model="formData.password"
              :trim="true"
              type="password"
              :maxlength="20"
              :input-border="false"
              :clearable="false"
              placeholder="登录密码"
              @focus="active = 'password'"
              @blur="active = ''"
            />
          </uni-forms-item>
        </view>
        <!-- 验证码（验证码登录） -->
        <view v-else :class="{ basic: true, active: active === 'code' }">
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

      <!-- 忘记密码（账号登录） -->
      <view
        v-if="loginTypeFlag"
        class="login__bottom"
        @click="$Router.push({ name: 'loginExtra', query: { smsType: '02' } })"
        ><uni-icons type="help" color="#00c853" size="12"></uni-icons
        ><text style="margin-left: 4rpx;">忘记密码</text></view
      >
    </view>

    <view class="btn-wrapper">
      <button
        type="primary"
        class="btn"
        :disabled="isDisabled"
        @click="onLogin"
      >
        登录
      </button>
      <button
        type="primary"
        class="btn"
        @click="loginTypeFlag = !loginTypeFlag"
      >
        {{ titleReverse }}
      </button>
    </view>

    <view
      class="register"
      @click="$Router.push({ name: 'loginExtra', query: { smsType: '00' } })"
      >注册账号</view
    >
  </view>
</template>

<script>
import mixin from './mixin'
import { accountLogin, smsLogin } from '@/api/user'
import { afterLogin } from '@/utils/auth'

export default {
  name: 'Login',

  mixins: [mixin], // getCode

  data() {
    return {
      // 表单
      formData: {
        phoneNum: '', // 手机号
        password: '', // 密码
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
        password: {
          rules: [
            {
              required: true,
              errorMessage: '请输入登录密码'
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
      loginTypeFlag: true, // true-账号登录；false-验证码登录
      smsType: '01', // 短信类型
      active: '', // 激活的输入框name字段
      title: '账号密码登录', // 标题text
      titleReverse: '短信验证登录', // 切换认证text
      inCountdown: false // true-倒计时中；false-非倒计时中
    }
  },

  computed: {
    /**
     * @description: 表单按钮是否禁用
     * @author: Hemingway
     */
    isDisabled() {
      if (this.loginTypeFlag) {
        return !this.formData.phoneNum || !this.formData.password // 验证账号登录表单
      } else {
        return !this.formData.phoneNum || !this.formData.code // 验证验证码登录表单
      }
    }
  },

  watch: {
    loginTypeFlag(value) {
      if (value) {
        this.title = '账号密码登录'
        this.titleReverse = '短信验证登录'
      } else {
        this.title = '短信验证登录'
        this.titleReverse = '账号密码登录'
      }
    }
  },

  methods: {
    /**
     * @description: 登录事件（含账号登录、短信登录）
     * @author: Hemingway
     */
    async onLogin() {
      // 表单校验及登录
      let api = null
      const paramObj = {}
      try {
        const { phoneNum, password, code } = await this.$refs.form.validate()
        if (this.loginTypeFlag) {
          // 账号登录
          paramObj.password = await this.encrypt(password)
          paramObj.username = phoneNum
          api = accountLogin
        } else {
          // 验证码登录
          paramObj.phone = phoneNum
          paramObj.code = code // 验证码
          api = smsLogin
        }

        const res = await api(paramObj)
        if (res.code === '0') {
          // 获取权限

          uni.showToast({
            title: '登录成功',
            icon: 'success',
            duration: 2000,
            complete: () => {
              this.$store.commit('setSessionId', res.sessionId)
              uni.setStorageSync('sessionId', res.sessionId)

              setTimeout(async () => {
                console.log('----------0 已登录未配置权限')
                await afterLogin()
                console.log('----------8 已完成整个登录过程')
                this.$Router.replaceAll({ name: 'home' })
              }, 1000)
            }
          })
        }
      } catch (error) {
        console.log('登录 error = ', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.login {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #fff;

  .container {
    margin: 80rpx 80rpx 0;
    position: relative;

    .title {
      margin-bottom: 24rpx;
      font-size: 40rpx;
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

    // 忘记密码
    .login__bottom {
      margin-top: 32rpx;
      text-align: right;
      font-size: 24rpx;
      color: #8b8b8b;
      position: absolute;
      right: 0;
      display: flex;
      align-items: center;
    }

    // 获取验证码
    .code {
      font-size: 24rpx;
      color: #8b8b8b;
      display: flex;
      align-items: center;
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

    .btn:last-child {
      margin-top: 24rpx;
      background-color: $uni-bg-color;
      color: #212121;
      border: 1px solid #e0e0e0;
    }

    .btn:last-child:active {
      opacity: 0.6;
    }

    .btn::after {
      display: none;
    }
  }

  .register {
    width: 100%;
    height: 138rpx;
    line-height: 138rpx;
    text-align: center;
    position: fixed;
    bottom: 0;
    color: $theme-color;
  }
}
</style>
