/*
 * @Description  : 提取公用js逻辑：获取验证码/获取公钥（封装加密方法）
 * @Autor        : Hemingway
 * @Date         : 2021-06-23 11:26:48
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-06-28 10:12:49
 * @FilePath     : \blockchain-minicode\src\pages\login\mixin.js
 */
import { JSEncrypt } from 'wxmp-rsa'
const encrypt = new JSEncrypt()
import { getMsgCode, getPublicKey } from '@/api/user'

export default {
  data() {
    return {}
  },

  methods: {
    /**
     * @description: 获取验证码
     * @author: Hemingway
     */
    async getCode() {
      try {
        // 手机号校验
        await this.$refs.form.validateField('phoneNum')

        // 获取验证码
        const { code } = await getMsgCode({
          phone: this.formData.phoneNum,
          smsType: this.smsType
        })
        if (code === '0') {
          this.inCountdown = true
          uni.showToast({
            title: '发送成功',
            icon: 'success',
            duration: 2000
          })
        }
      } catch (error) {
        console.log('获取验证码 error = ', error)
      }
    },

    /**
     * @description: rsa加密
     * @param {String} raw 被加密的字符
     * @author: Hemingway
     */
    async encrypt(raw) {
      try {
        const publicKey = await this.getKey()
        if (publicKey) {
          // console.log('publicKey = ', publicKey)
          encrypt.setPublicKey(publicKey)
          return encrypt.encryptLong(raw)
        }
      } catch (error) {
        console.log('rsa加密 error = ', error)
      }
    },

    /**
     * @description: 获取公钥
     * @author: Hemingway
     */
    async getKey() {
      try {
        const { code, publicKey } = await getPublicKey()
        if (code === '0') {
          return publicKey
        }
      } catch (error) {
        console.log('获取加密公钥 error = ', error)
      }
    }
  }
}
