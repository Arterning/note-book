<!--
 * @Description  : 密码设置
 * @Autor        : Hemingway
 * @Date         : 2022-03-18 16:34:57
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-04-23 11:50:06
 * @FilePath     : \blockchain-admin\src\views\profile\components\SetPassword.vue
-->
<template>
  <!-- 密码设置 dialog -->
  <i-dialog :visible.sync="visible">
    <template #title>
      {{ isSet ? '更改密码' : '设置密码' }}
      <el-tooltip
        content="密码长度为 8 ~ 20位（含大写字母、小写字母、数字、特殊字符两项及以上）"
        placement="top"
        class="tooltip"
      >
        <i class="el-icon-info"> </i>
      </el-tooltip>
    </template>
    <el-form :model="form" :rules="rules" ref="form">
      <!-- 旧密码 -->
      <el-tooltip
        v-model="capsTooltip0"
        content="Caps lock is On"
        placement="right"
        manual
        v-if="isSet"
      >
        <el-form-item label="旧密码" prop="password0">
          <el-input
            tabindex="0"
            ref="password0"
            :maxlength="20"
            v-model="form.password0"
            :type="passwordType0"
            autocomplete="off"
            placeholder="请输入旧密码"
            @keyup.native="checkCapslock($event, 0)"
            @blur="capsTooltip0 = false"
          ></el-input>
          <span class="show-pwd" @click="showPwd(0)">
            <svg-icon
              :icon-class="passwordType0 === 'password' ? 'eye' : 'eye-open'"
            />
          </span>
        </el-form-item>
      </el-tooltip>
      <!-- 新密码 -->
      <el-tooltip
        v-model="capsTooltip1"
        content="Caps lock is On"
        placement="right"
        manual
      >
        <el-form-item :label="hasPassword ? '新密码' : '密码'" prop="password1">
          <el-input
            ref="password1"
            :maxlength="20"
            v-model="form.password1"
            :type="passwordType1"
            autocomplete="off"
            :placeholder="hasPassword ? '请输入新密码' : '请输入密码'"
            @keyup.native="checkCapslock($event, 1)"
            @blur="capsTooltip1 = false"
          ></el-input
          ><span class="show-pwd" @click="showPwd(1)">
            <svg-icon
              :icon-class="passwordType1 === 'password' ? 'eye' : 'eye-open'"
            />
          </span> </el-form-item
      ></el-tooltip>
      <!-- 确认新密码 -->
      <el-tooltip
        v-model="capsTooltip2"
        content="Caps lock is On"
        placement="right"
        manual
      >
        <el-form-item label="确认密码" prop="password2">
          <el-input
            ref="password2"
            :maxlength="20"
            v-model="form.password2"
            :type="passwordType2"
            autocomplete="off"
            :placeholder="hasPassword ? '请确认新密码' : '请确认密码'"
            @keyup.native="checkCapslock($event, 2)"
            @blur="capsTooltip2 = false"
          ></el-input
          ><span class="show-pwd" @click="showPwd(2)">
            <svg-icon
              :icon-class="passwordType2 === 'password' ? 'eye' : 'eye-open'"
            /> </span></el-form-item
      ></el-tooltip>
    </el-form>

    <div slot="footer" class="dialog-footer">
      <el-button @click="onCancel">取 消</el-button>
      <el-button
        type="primary"
        @click="onSubmit"
        :disabled="isSubmitDisabled"
        :loading="loading"
        >确 定</el-button
      >
    </div>
  </i-dialog>
</template>

<script>
import { checkPasswordLevel } from '@/utils/validate'

export default {
  name: 'SetPassword',

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
      visible: false,
      loading: false,
      passwordType0: 'password',
      passwordType1: 'password',
      passwordType2: 'password',
      capsTooltip0: false,
      capsTooltip1: false,
      capsTooltip2: false,
      form: {
        password0: '', // 旧密码
        password1: '', // 新密码
        password2: '' // repeat新密码
      },
      rules: {
        password0: [{ trigger: 'blur', validator: checkPasswordLevel }],
        password1: [{ trigger: 'blur', validator: checkPasswordLevel }],
        password2: [{ trigger: 'blur', validator: validatePassword }]
      },
      isShow: false // 手动干预是否进入修改密码状态
    }
  },

  computed: {
    hasPassword() {
      return this.$store.getters.hasPassword
    },
    // 是否是设置密码状态
    isSet() {
      return this.hasPassword && this.isShow
    },
    // 修改密码提交按钮时候禁用
    isSubmitDisabled() {
      let flag = !this.form.password1 || !this.form.password2

      if (this.hasPassword) {
        flag = flag || !this.form.password0
      }

      return flag
    }
  },

  watch: {
    hasPassword: {
      handler(newV, oldV) {
        if (oldV === undefined) {
          this.isShow = newV || false // 初始化isShow
        }
      },
      immediate: true
    }
  },

  methods: {
    /**
     * @description: 检查caps lock
     * @param {Number} flag 0或1或2
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
     * @description: 密码设置/修改提交事件
     * @author: Hemingway
     */
    onSubmit() {
      if (this.form.password1 === this.form.password2) {
        this.$refs.form.validate(async valid => {
          if (valid) {
            this.loading = true

            let dispatch, payload
            const { password0, password1 } = this.form
            if (this.isSet) {
              // 修改密码
              dispatch = 'user/updatePassword'
              payload = { oldPassword: password0, newPassword: password1 }
            } else {
              // 设置密码
              dispatch = 'user/setPassword'
              payload = { password1 }
            }

            try {
              await this.$store.dispatch(dispatch, payload)

              this.loading = false

              this.$message({
                message: `${
                  this.isSet
                    ? '密码更改成功'
                    : '密码设置成功，下次登录系统可使用密码登录'
                }`,
                type: 'success',
                duration: 3000,
                onClose: () => {
                  this.onCancel()
                  setTimeout(() => {
                    this.isShow = true
                  }, 3000)
                }
              })
            } catch (error) {
              this.loading = false
              console.log('设置密码 error = ', error)
            }
          }
        })
      }
    },

    /**
     * @description: 终止操作，重置状态
     * @author: Hemingway
     */
    onCancel() {
      this.visible = false

      this.$refs.form.resetFields()

      this.passwordType0 = 'password'
      this.passwordType1 = 'password'
      this.passwordType2 = 'password'
    }
  }
}
</script>

<style lang="scss" scoped>
$dark_gray: #889aa4;

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
