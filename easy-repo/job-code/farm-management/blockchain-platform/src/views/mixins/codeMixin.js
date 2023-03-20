/*
 * @Description  : 提取公用js逻辑（发送验证码）
 * @Autor        : Hemingway
 * @Date         : 2021-06-23 11:26:48
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-22 09:57:56
 * @FilePath     : \blockchain-platform\src\views\mixins\codeMixin.js
 */
import { validPhoneNumber } from '@/utils/validate'
import { sendCode } from '@/api/user'

export default {
  data() {
    return {
      send: {
        text: '发送验证码',
        type: 'primary',
        loading: false,
        inCount: false,
        count: 60,
        timer: null
      } // 发送验证码相关状态
    }
  },

  computed: {
    // 发送验证码按钮是否禁用
    isDisabled() {
      return !validPhoneNumber(this.form.username)
    }
  },

  methods: {
    /**
     * @description: 发送验证码
     * @param {String} phone
     * @author: Hemingway
     */
    async handleSendCode(phone, smsType) {
      try {
        this.send.loading = true
        // await sendCode({ phone, isEnterprisePhone: 'N' })
        await sendCode({ phone, username: phone, smsType })
        this.send.loading = false

        this.handleTimer()
      } catch (error) {
        console.log('发送验证码 error = ', error)
        this.send.loading = false
      }
    },

    /**
     * @description: 执行倒计时逻辑
     * @author: Hemingway
     */
    handleTimer() {
      this.send.inCount = true
      this.send.timer = setInterval(() => {
        this.send.count--
        if (this.send.count === 0) {
          this.resetTimer()
        }
      }, 1000)
    },

    /**
     * @description: 重置倒计时
     * @author: Hemingway
     */
    resetTimer() {
      clearInterval(this.send.timer)
      this.send.inCount = false
      this.send.count = 60
    }
  }
}
