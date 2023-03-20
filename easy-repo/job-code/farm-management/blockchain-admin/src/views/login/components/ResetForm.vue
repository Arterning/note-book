<!--
 * @Description  : 重置密码组件（账号验证 -> 重置密码）
 * @Autor        : Hemingway
 * @Date         : 2021-12-14 17:02:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-02 15:59:14
 * @FilePath     : \blockchain-admin\src\views\login\components\ResetForm.vue
-->
<template>
  <el-form
    ref="form"
    :model="form"
    :rules="rules"
    autocomplete="on"
    label-position="left"
    class="reset-form"
  >
    <div class="title-container">
      找回密码
      <el-tooltip
        content="密码长度为 8 ~ 20位（含大写字母、小写字母、数字、特殊字符两项及以上）"
        placement="top"
        class="tooltip"
      >
        <i class="el-icon-info"> </i>
      </el-tooltip>
    </div>

    <!-- 账号验证 -->
    <template v-if="isValidateAccout">
      <!-- 用户名：手机号 -->
      <el-form-item prop="username">
        <!-- <span class="svg-container">
          <svg-icon icon-class="user" />
        </span> -->
        <el-input
          :maxlength="11"
          ref="username"
          v-model="form.username"
          placeholder="手机号码"
          name="username"
          type="text"
          tabindex="1"
        />
      </el-form-item>
      <!-- 验证码 -->
      <el-form-item
        prop="code"
        :class="{ hidden: current === 0 }"
        style="position: relative"
      >
        <span class="svg-container">
          <svg-icon icon-class="message" />
        </span>
        <el-input
          :maxlength="6"
          v-model="form.code"
          type="text"
          placeholder="验证码"
          name="code"
          tabindex="2"
        />
        <div class="send">
          <el-button
            v-if="!send.inCount"
            size="small"
            plain
            :type="send.type"
            :loading="send.loading"
            @click="handleSendCode(form.username, '02')"
            :disabled="!form.username"
            >发送验证码</el-button
          >
          <div v-else class="count">
            {{ send.count + ' s' }}
          </div>
        </div>
      </el-form-item>
    </template>
    <!-- 重置密码 -->
    <template v-else>
      <!-- 新密码 -->
      <el-tooltip
        v-model="capsTooltip1"
        content="Caps lock is On"
        placement="right"
        manual
      >
        <el-form-item prop="password1">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType1"
            ref="password1"
            v-model="form.password1"
            :type="passwordType1"
            placeholder="请设置新密码"
            name="password1"
            tabindex="2"
            @keyup.native="checkCapslock($event, 1)"
            @blur="capsTooltip1 = false"
            @keyup.enter.native="$refs.password2.focus()"
          />
          <span class="show-pwd" @click="showPwd(1)">
            <svg-icon
              :icon-class="passwordType1 === 'password' ? 'eye' : 'eye-open'"
            />
          </span>
        </el-form-item>
      </el-tooltip>
      <!-- 二次输入新密码 -->
      <el-tooltip
        v-model="capsTooltip2"
        content="Caps lock is On"
        placement="right"
        manual
      >
        <el-form-item prop="password2">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType2"
            ref="password2"
            v-model="form.password2"
            :type="passwordType2"
            placeholder="请确认新密码"
            name="password2"
            tabindex="2"
            autocomplete="on"
            @keyup.native="checkCapslock($event, 2)"
            @blur="capsTooltip2 = false"
          />
          <span class="show-pwd" @click="showPwd(2)">
            <svg-icon
              :icon-class="passwordType2 === 'password' ? 'eye' : 'eye-open'"
            />
          </span>
        </el-form-item>
      </el-tooltip>
    </template>

    <!-- current为0时，校验手机号是否合法；current为1时，校验验证码是否正确；current为2时，提交 -->
    <el-button
      :loading="stepsMap[current].loading"
      type="primary"
      style="width: 100%; margin-top: 9px"
      :disabled="isDisabled"
      @click.native.prevent="stepsMap[current].callback"
      >{{ stepsMap[current].btnText }}</el-button
    >

    <div class="back-container">
      <span>已有账号，</span
      ><span class="back" @click="$emit('login')">返回登录</span>
    </div>
  </el-form>
</template>

<script>
// /* eslint-disable */
import mixin from '../../mixins/codeMixin'
import { validPhoneNumber, checkPasswordLevel } from '@/utils/validate'
import { validatePhone, validateCode } from '@/api/user'
export default {
  name: 'ResetForm',

  mixins: [mixin], // handleSendCode

  props: {
    username: {
      type: String,
      default: ''
    }
  },

  watch: {
    username: {
      handler(newVal) {
        this.form.username = newVal
      },
      immediate: true
    }
  },

  data() {
    // 校验手机号
    const validatePhoneNumber = (rule, value, callback) => {
      if (!validPhoneNumber(value)) {
        callback(new Error('请输入正确的手机号'))
      } else {
        callback()
      }
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
    // 校验密码
    const validatePassword = (rule, value, callback) => {
      if (this.form.password1 !== this.form.password2) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        checkPasswordLevel(rule, value, callback)
      }
    }

    return {
      current: 0, // 当前步骤：0 验证手机号；1 校验验证码；2 提交
      stepsMap: {
        0: {
          btnText: '下一步',
          loading: false,
          callback: this.handleValidateUsername
        },
        1: {
          btnText: '下一步',
          loading: false,
          callback: this.handleValidateCode
        },
        2: { btnText: '确认', loading: false, callback: this.handleReset }
      }, // 各步骤相关状态
      form: {
        username: '',
        code: '',
        certificate: '', // code 对应的验证码
        password1: '',
        password2: ''
      },
      rules: {
        username: [
          { required: true, trigger: 'blur', validator: validatePhoneNumber }
        ],
        code: [{ required: true, trigger: 'blur', validator: validateCode }],
        password1: [
          { required: true, trigger: 'blur', validator: checkPasswordLevel }
        ],
        password2: [
          { required: true, trigger: 'blur', validator: validatePassword }
        ]
      },
      passwordType1: 'password',
      passwordType2: 'password',
      capsTooltip1: false,
      capsTooltip2: false
    }
  },

  computed: {
    // true 账号校验步骤；false 重置密码步骤
    isValidateAccout() {
      return this.current === 0 || this.current === 1
    },

    // 下一步按钮是否禁用
    isDisabled() {
      const { username, code, certificate, password1, password2 } = this.form
      let flag = !username

      if (this.current === 1) {
        flag = flag || !code
      } else if (this.current === 2) {
        flag = flag || !certificate || !password1 || !password2
      }

      return flag
    }
  },

  mounted() {
    if (this.form.username === '') {
      this.$refs.username.focus()
    }
  },

  methods: {
    /**
     * @description: 检查caps lock
     * @param {Number} flag 1或2
     * @author: Hemingway
     */
    checkCapslock(e, flag) {
      const { key } = e
      this[`capsTooltip${flag}`] =
        key && key.length === 1 && key >= 'A' && key <= 'Z'
    },

    /**
     * @description: 切换密码输入框显示和隐藏
     * @param {Number} flag 1或2
     * @author: Hemingway
     */
    showPwd(flag) {
      if (this[`passwordType${flag}`] === 'password') {
        this[`passwordType${flag}`] = ''
      } else {
        this[`passwordType${flag}`] = 'password'
      }
      this.$nextTick(() => {
        this.$refs[`password${flag}`].focus()
      })
    },

    /**
     * @description: 校验用户手机号
     * @author: Hemingway
     */
    handleValidateUsername() {
      this.$refs.form.validateField('username', async error => {
        if (error) {
          return
        } else {
          try {
            this.stepsMap[0].loading = true
            await validatePhone({ phone: this.form.username })
            this.current++
            this.stepsMap[0].loading = false
          } catch (error) {
            console.log('校验用户手机号 error = ', error)
            this.stepsMap[0].loading = false
          }
        }
      })
    },

    /**
     * @description: 校验验证码
     * @author: Hemingway
     */
    async handleValidateCode() {
      // 校验表单手机号和验证码，然后调接口
      let errCount = 0
      this.$refs.form.validateField(['username', 'code'], error => {
        if (error) {
          errCount++
        }
      })

      if (errCount > 0) {
        return
      } else {
        try {
          this.stepsMap[1].loading = true
          const { certifiCate } = await validateCode({
            phone: this.form.username,
            code: this.form.code,
            msgType: '02'
          })
          this.current++
          this.form.certificate = certifiCate
          this.stepsMap[1].loading = false
        } catch (error) {
          console.log('校验用户手机号 error = ', error)
          this.stepsMap[1].loading = false
        }
      }
    },

    /**
     * @description: 重置密码逻辑
     * @author: Hemingway
     */
    handleReset() {
      if (this.form.password1 === this.form.password2) {
        this.$refs.form.validate(valid => {
          if (valid) {
            // 调用重置密码接口，成功后返回登录
            this.stepsMap[2].loading = true

            this.$store
              .dispatch('user/resetPassword', this.form)
              .then(() => {
                this.stepsMap[2].loading = false

                this.$message({
                  message: `密码重置成功，即将导航至登录页面`,
                  type: 'success',
                  duration: 3000,
                  onClose: () => {
                    this.$emit('reset-success', this.form.username)
                  }
                })
              })
              .catch(() => {
                this.stepsMap[2].loading = false
              })
          }
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import './countdown.scss';
.reset-form {
  $dark_gray: #889aa4;
  $light_gray: #eee;
  position: relative;
  padding: 10px;
  height: 249px;

  .tooltip {
    position: relative;
    top: 1px;
    left: 4px;
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    height: 26px;
    padding-left: 15px;
    font-size: 18px;
    color: #00c853;
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  .hidden {
    // display: none;
    visibility: hidden;
  }

  .back-container {
    position: relative;
    top: 10px;
    height: 15px;
    font-size: 14px;
    text-align: center;

    .back {
      color: #00c853;
    }
    .back:hover {
      cursor: pointer;
    }
  }
}
</style>
