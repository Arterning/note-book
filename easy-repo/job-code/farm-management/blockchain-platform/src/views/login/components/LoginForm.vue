<!--
 * @Description  : 登录组件（密码登录或验证码登录；短信登录成功后设置密码）
 * @Autor        : Hemingway
 * @Date         : 2021-12-14 17:02:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-02 16:21:12
 * @FilePath     : \blockchain-platform\src\views\login\components\LoginForm.vue
-->
<template>
  <el-form
    ref="form"
    :model="form"
    :rules="rules"
    autocomplete="on"
    label-position="left"
    class="login-form"
  >
    <!-- <div class="title-container">
      <el-radio-group v-model="isPasswordLogin" size="mini">
        <el-radio-button :label="true" :disabled="isPasswordLoginDisabled"
          >密码登录</el-radio-button
        >
        <el-radio-button :label="false">短信登录</el-radio-button>
      </el-radio-group>
    </div> -->

    <!-- 用户名：手机号 -->

    <el-form-item prop="username" :error="error" label="">
      <el-input
        :maxlength="11"
        ref="username"
        v-model="form.username"
        placeholder="请输入手机号"
        name="username"
        type="text"
        tabindex="1"
        autocomplete="off"
      />
    </el-form-item>

    <template v-if="isPasswordLogin">
      <!-- 密码 -->
      <el-tooltip
        v-model="capsTooltip"
        content="Caps lock is On"
        placement="right"
        manual
      >
        <el-form-item prop="password" label="" class="password-item">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="form.password"
            :type="passwordType"
            placeholder="请输入密码"
            name="password"
            tabindex="2"
            autocomplete="off"
            @keyup.native="checkCapslock"
            @blur="capsTooltip = false"
            @keyup.enter.native="onLogin"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon
              :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"
            />
          </span>
        </el-form-item>
      </el-tooltip>

      <div class="forget" @click="$emit('reset', form.username)">
        忘记密码？
      </div>
    </template>
    <template v-else>
      <!-- 验证码 -->
      <el-form-item prop="code" label="">
        <el-input
          :maxlength="6"
          v-model="form.code"
          type="text"
          placeholder="验证码"
          name="code"
          tabindex="2"
          autocomplete="off"
          @keyup.enter.native="onLogin"
        />
        <div class="send">
          <el-button
            v-if="!send.inCount"
            size="small"
            :type="send.type"
            :loading="send.loading"
            @click="handleSendCode(form.username, '01')"
            :disabled="isDisabled"
            >发送验证码</el-button
          >
          <div v-else class="count">
            {{ send.count + ' s' }}
          </div>
        </div>
      </el-form-item>
    </template>

    <el-button
      :loading="loading"
      type="primary"
      style="width: 100%; margin-top: 3px; background-color: #05dbfe"
      @click.native.prevent="onLogin"
      :disabled="isSubmitDisabled"
      >登录</el-button
    >
  </el-form>
</template>

<script>
import mixin from '../../mixins/codeMixin'
import { validPhoneNumber, checkPasswordLevel } from '@/utils/validate'
import { checkUsername } from '@/api/user'

export default {
  name: 'LoginForm',

  mixins: [mixin], // handleSendCode

  props: {
    username: {
      type: String,
      default: ''
    },
    isPasswordLogin: {
      type: Boolean,
      default: true
    }
  },

  watch: {
    username: {
      handler(newVal) {
        this.form.username = newVal
      },
      immediate: true
    },
    'form.username': {
      handler(newVal) {
        newVal && this.checkUsernameStatus(newVal)
      },
      immediate: true
    }
  },

  computed: {
    // 提交按钮是否禁用
    isSubmitDisabled() {
      let flag = !this.form.username

      if (this.isPasswordLogin) {
        flag = flag || !this.form.password
      } else {
        flag = flag || !this.form.code
      }

      return flag
    }
  },

  data() {
    // 校验密码
    const validatePassword = (rule, value, callback) => {
      checkPasswordLevel(rule, value, callback)
    }
    // 校验验证码
    const validateCode = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('验证码不能少于6位数字'))
      } else if (!/^\d{6}$/.test(value)) {
        callback(new Error('验证码含有非数字字符'))
      } else {
        callback()
      }
    }
    return {
      error: '', // 账号错误提示
      hasAlerted: {}, // 验证码登录alert状态对象，记录未设置密码手机号的alert状态
      isPasswordLoginDisabled: false,
      // isPasswordLogin: true, // true 密码登录；false 短信登录
      form: {
        username: '',
        password: '',
        code: ''
      },
      rules: {
        username: [
          {
            required: true,
            trigger: 'blur',
            validator: this.validatePhoneNumber
          }
        ],
        password: [
          { required: true, trigger: 'blur', validator: validatePassword }
        ],
        code: [{ required: true, trigger: 'blur', validator: validateCode }]
      },
      passwordType: 'password',
      capsTooltip: false,
      loading: false,
      redirect: undefined,
      otherQuery: {}
    }
  },

  mounted() {
    if (this.form.username === '') {
      this.$refs.username.focus()
    } else if (this.form.password === '') {
      this.$refs.password.focus()
    }
  },

  methods: {
    checkCapslock(e) {
      const { key } = e
      this.capsTooltip = key && key.length === 1 && key >= 'A' && key <= 'Z'
    },

    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },

    /**
     * @description: 手机号校验逻辑
     * @param {Object} rule
     * @param {String} value
     * @param {Function} callback
     * @author: Hemingway
     */
    validatePhoneNumber(rule, value, callback) {
      if (!validPhoneNumber(value)) {
        callback(new Error('请输入正确的手机号'))
      } else if (this.error) {
        callback(new Error(this.error)) // rule规则之外的手动干预
      } else {
        callback()
      }
    },

    /**
     * @description: 账号检测（用户是否存在，以及是否需要设置密码）
     * @param {String} username
     * @author: Hemingway
     */
    async checkUsernameStatus(username) {
      this.error = ''

      if (validPhoneNumber(username)) {
        try {
          const { data } = await checkUsername({
            phone: this.form.username
          })
          if (data.isExist !== 'Y') {
            this.error = '该手机号的用户不存在，请确认'
          }
        } catch (error) {
          console.log('检测账号状态 error = ', error)
        }
      }
    },

    checkUsernameCallback(username) {
      this.isPasswordLogin = false // 跳转至验证码登录
      this.isPasswordLoginDisabled = true // 禁用密码登录
      this.hasAlerted[username] = true // 更新提示记录
    },

    /**
     * @description: 登录逻辑
     * @author: Hemingway
     */
    onLogin() {
      this.$refs.form.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store
            .dispatch('user/login', {
              ...this.form,
              isPasswordLogin: this.isPasswordLogin
            })
            .then(isSetPassWord => {
              this.loading = false

              // 如果需要设置密码
              if (isSetPassWord !== 'Y') {
                this.$emit('set-password')
              } else {
                this.$emit('login-success')
              }
              // 临时修改，登陆成功跳转到个人中心页面
              // this.$router.push('/profile')
            })
            .catch(error => {
              this.loading = false
              console.log('用户登录 error = ', error)
            })
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import './countdown.scss';
.login-form {
  $dark_gray: #889aa4;
  $light_gray: #eee;
  position: relative;
  padding: 15px 20px;

  ::v-deep .el-form-item__content {
    display: flex;
    align-items: center;
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    height: 48px;
    text-align: center;
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 2px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  .forget {
    position: absolute;
    right: 17px;
    bottom: 60px;
    font-size: 14px;
    color: rgb(33, 214, 216);
  }

  .forget:hover {
    color: #022464;
    cursor: pointer;
  }
}

.password-item {
  ::v-deep .el-form-item__label {
    margin-left: 10px;
  }
}

.send {
  right: 4px;
  margin-top: -7px;
}
</style>
