<!--
 * @Description  : 手机号设置
 * @Autor        : Hemingway
 * @Date         : 2022-03-18 16:46:23
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-24 10:27:22
 * @FilePath     : \blockchain-admin\src\views\profile\components\SetPhone.vue
-->
<template>
  <i-dialog :title="stepsMap[current].title" :visible.sync="visible">
    <el-form :model="form" :rules="rules" ref="form">
      <!-- 身份验证 -->
      <template v-if="current === 0">
        <el-form-item label="当前手机号">
          <el-input
            :maxlength="11"
            v-model="oldPhone"
            placeholder="手机号码"
            type="text"
            tabindex="1"
            disabled
          />
          <el-tooltip
            content="如果该手机号不能接收验证码，请联系农场管理员"
            placement="top"
            class="tooltip"
          >
            <i class="el-icon-info"> </i>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="验证码" prop="oldCode" style="position: relative">
          <el-input
            :maxlength="6"
            v-model="form.oldCode"
            type="text"
            placeholder="验证码"
            tabindex="2"
          />
          <div class="send">
            <el-button
              v-if="!send.inCount"
              size="small"
              type="text"
              :loading="send.loading"
              @click.stop="handleSendCode(oldPhone, '04')"
              >发送验证码</el-button
            >
            <div v-else class="count">
              {{ send.count + ' s' }}
            </div>
          </div>
        </el-form-item>
      </template>

      <!-- 更改手机号 -->
      <template v-else>
        <el-form-item label="新手机号" prop="phone">
          <el-input
            :maxlength="11"
            v-model="form.phone"
            placeholder="手机号码"
            name="username"
            type="text"
            tabindex="1"
          />
        </el-form-item>
        <el-form-item label="验证码" prop="code" style="position: relative">
          <el-input
            :maxlength="6"
            v-model="form.code"
            type="text"
            placeholder="验证码"
            tabindex="2"
          />
          <div class="send">
            <el-button
              v-if="!send.inCount"
              size="small"
              type="text"
              :loading="send.loading"
              @click.stop="handleSendCode(form.phone, '03')"
              :disabled="isSendCodeDisabled"
              >发送验证码</el-button
            >
            <div v-else class="count">
              {{ send.count + ' s' }}
            </div>
          </div>
        </el-form-item>
      </template>
    </el-form>

    <div slot="footer" class="dialog-footer">
      <el-button @click="onCancel">取 消</el-button>
      <el-button
        type="primary"
        @click="stepsMap[current].callback"
        :disabled="isSubmitDisabled"
        :loading="loading"
        >{{ stepsMap[current].btnText }}</el-button
      >
    </div>
  </i-dialog>
</template>

<script>
import mixin from '../../mixins/codeMixin'
import { validPhoneNumber } from '@/utils/validate'
import { validateCode } from '@/api/user'

export default {
  name: 'SetPhone',

  mixins: [mixin], // handleSendCode

  props: {
    oldPhone: {
      type: String,
      default: ''
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

    return {
      stepsMap: {
        0: {
          title: '身份验证',
          callback: this.onNext,
          btnText: '下一步'
        },
        1: {
          title: '更改手机号',
          callback: this.onSubmit,
          btnText: '确 定'
        }
      },
      visible: false,
      loading: false, // 按钮的loading
      current: 0, // 当前步骤：0 验证当前手机号；1 提交
      form: {
        oldCode: '',

        phone: '',
        code: ''
      },
      rules: {
        oldCode: [{ trigger: 'blur', validator: validateCode }],
        phone: [{ trigger: 'blur', validator: validatePhoneNumber }],
        code: [{ trigger: 'blur', validator: validateCode }]
      },
      oldCertificate: '', // 旧手机短信凭证
      certificate: ''
    }
  },

  computed: {
    // 发送验证码按钮是否禁用
    isSendCodeDisabled() {
      return !validPhoneNumber(this.form.phone)
    },
    // 下一步按钮是否禁用
    isSubmitDisabled() {
      const { oldCode, phone, code } = this.form

      if (this.current === 0) {
        return !oldCode
      } else {
        return !phone || !code
      }
    }
  },

  methods: {
    /**
     * @description: 下一步，校验旧手机
     * @author: Hemingway
     */
    async onNext() {
      let errCount = 0
      this.$refs.form.validateField(['oldCode'], error => {
        if (error) {
          errCount++
        }
      })

      if (errCount > 0) {
        return
      } else {
        try {
          this.loading = true
          const { certifiCate } = await validateCode({
            phone: this.oldPhone,
            code: this.form.oldCode,
            msgType: '04'
          })
          this.oldCertificate = certifiCate
          this.loading = false

          this.current++
          if (this.send.timer) {
            this.resetTimer()
          }
        } catch (error) {
          console.log('校验用户手机号 error = ', error)
          this.loading = false
        }
      }
    },

    /**
     * @description: 修改手机号提交事件
     * @author: Hemingway
     */
    async onSubmit() {
      let errCount = 0
      this.$refs.form.validateField(['phone', 'code'], error => {
        if (error) {
          errCount++
        }
      })

      if (errCount > 0) {
        return
      } else {
        try {
          this.loading = true
          const { certifiCate } = await validateCode({
            phone: this.form.phone,
            code: this.form.code,
            msgType: '03'
          })
          this.certificate = certifiCate

          await this.$store.dispatch('user/updatePhone', {
            oldCertificate: this.oldCertificate,
            phone: this.form.phone,
            certificate: this.certificate,
            code: this.form.code
          })

          this.loading = false

          this.$message({
            message: '手机号更改成功，3s后请重新登录',
            type: 'success',
            duration: 3000,
            onClose: async () => {
              await this.$store.dispatch('user/resetToken')
              this.$router.push(`/login?redirect=${this.$route.fullPath}`)
            }
          })
        } catch (error) {
          console.log('校验用户手机号&提交修改手机号 error = ', error)
          this.loading = false
        }
      }
    },

    /**
     * @description: 终止操作，重置状态
     * @author: Hemingway
     */
    onCancel() {
      this.visible = false

      this.current = 0
      this.$refs.form.resetFields()
      if (this.send.timer) {
        this.resetTimer()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import './countdown.scss';
</style>
