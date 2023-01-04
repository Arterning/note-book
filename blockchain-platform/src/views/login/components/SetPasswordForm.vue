<!--
 * @Description  : 设置密码组件
 * @Autor        : Hemingway
 * @Date         : 2021-12-31 12:14:50
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-08 16:41:49
 * @FilePath     : \blockchain-platform\src\views\login\components\SetPasswordForm.vue
-->
<template>
  <el-form ref="form" :model="form" :rules="rules">
    <!-- 新密码 -->
    <el-tooltip
      v-model="capsTooltip1"
      content="Caps lock is On"
      placement="right"
      manual
    >
      <el-form-item label="密码" prop="password1">
        <el-input
          :key="passwordType1"
          ref="password1"
          v-model="form.password1"
          :type="passwordType1"
          placeholder="请输入密码"
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
      <el-form-item label="确认密码" prop="password2">
        <el-input
          :key="passwordType2"
          ref="password2"
          v-model="form.password2"
          :type="passwordType2"
          placeholder="请确认密码"
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
  </el-form>
</template>

<script>
import { checkPasswordLevel } from '@/utils/validate'

export default {
  name: 'SetPasswordForm',

  data() {
    // 校验密码
    const validatePassword = (rule, value, callback) => {
      if (this.form.password1 !== this.form.password2) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        checkPasswordLevel(rule, value, callback)
      }
    }
    return {
      form: {
        password1: '',
        password2: ''
      },
      rules: {
        password1: [{ trigger: 'blur', validator: checkPasswordLevel }],
        password2: [{ trigger: 'blur', validator: validatePassword }]
      },
      passwordType1: 'password',
      passwordType2: 'password',
      capsTooltip1: false,
      capsTooltip2: false,
      loading: false
    }
  },

  watch: {
    form: {
      handler() {
        if (this.form.password1 && this.form.password2) {
          this.$emit('submit-disabled', false)
        } else {
          this.$emit('submit-disabled', true)
        }
      },
      deep: true
    }
  },

  created() {},

  mounted() {
    if (this.form.password1 === '') {
      this.$refs.password1.focus()
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
     * @description: 设置密码逻辑
     * @author: Hemingway
     */
    async submit() {
      if (this.form.password1 === this.form.password2) {
        this.$refs.form.validate(valid => {
          if (valid) {
            // 调用密码接口，成功后返回登录
            this.loading = true
            this.$store
              .dispatch('user/setPassword', this.form)
              .then(() => {
                this.loading = false

                this.$message({
                  message: `密码设置成功，下次登录系统可使用密码登录`,
                  type: 'success',
                  duration: 3000,
                  onClose: () => {
                    this.$emit('set-password-success')
                  }
                })
              })
              .catch(error => {
                this.loading = false
                console.log('设置密码 error = ', error)
              })
          }
        })
      }
    }
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #283443;
$input: #212121;
$cursor: #00c853;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .set-password-form .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.set-password-form {
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $input;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style>

<style lang="scss" scoped>
// 查看密码
.show-pwd {
  position: absolute;
  right: 42px;
  top: 0;
  font-size: 16px;
  color: #889aa4;
  cursor: pointer;
  user-select: none;
}
</style>
