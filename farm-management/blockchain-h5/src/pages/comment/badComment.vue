<!--
 * @Author       : qinjj
 * @Date         : 2021-05-24 09:55:30
 * @LastEditlastTime : 2021-11-19 16:08:41
 * @LastEditors  : Your Name
 * @Description  : 入口页
 * @FilePath     : \blockchain-h5\src\pages\comment\badComment.vue
-->
<template>
  <view class="badComment">
    <u--form
      labelPosition="left"
      labelWidth="90"
      :model="formData"
      :rules="rules"
      ref="form1"
    >
      <u-form-item label="投诉人：" prop="complainant" borderBottom required>
        <u--input
          v-model="formData.complainant"
          border="none"
          placeholder="请输入投诉人"
        ></u--input>
      </u-form-item>
      <u-form-item
        label="联系电话："
        prop="complainantsHotline"
        borderBottom
        required
      >
        <u--input
          v-model="formData.complainantsHotline"
          border="none"
          placeholder="请输入联系电话"
        ></u--input>
      </u-form-item>
      <u-form-item label="联系邮箱：" prop="mailbox" borderBottom>
        <u--input
          v-model="formData.mailbox"
          border="none"
          placeholder="请输入联系邮箱"
        ></u--input>
      </u-form-item>
      <u-form-item
        label="投诉理由："
        prop="complaintsReasons"
        borderBottom
        required
      >
        <u--textarea
          v-model="formData.complaintsReasons"
          placeholder="请输入投诉理由"
          count
          maxlength="200"
        ></u--textarea>
      </u-form-item>
    </u--form>
    <view class="image-box">
      <view class="image-title">
        附件：
      </view>
      <u-upload
        :fileList="fileList1"
        @afterRead="afterRead"
        @delete="deletePic"
        name="1"
        multiple
        :maxCount="10"
      ></u-upload>
    </view>
    <view class="footer-box">
      <u-button
        type="success"
        text="提交"
        class="footer-btn"
        @click="submit"
      ></u-button>
      <u-button
        type="success"
        :plain="true"
        text="取消"
        class="footer-btn"
        @click="backBtn"
      ></u-button>
    </view>
    <u-toast ref="uToast"></u-toast>
  </view>
</template>

<script>
import { submitComplaintDetails } from '@/api/myRice'
export default {
  name: 'badComment',

  data() {
    return {
      formData: {
        complainant: '',
        complainantsHotline: '',
        mailbox: '',
        complaintsReasons: '',
        complaintImages: ''
      },
      rules: {
        complainant: {
          type: 'string',
          required: true,
          message: '请填写投诉人',
          trigger: ['blur', 'change']
        },
        complainantsHotline: [
          {
            type: 'string',
            required: true,
            message: '请填写联系电话',
            trigger: ['blur', 'change']
          },
          {
            pattern: /^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$/g,
            // 正则检验前先将值转为字符串
            transform(value) {
              return String(value)
            },
            message: '联系电话格式不正确',
            trigger: ['blur', 'change']
          }
        ],
        mailbox: {
          pattern: /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.(com|cn|net)$/g,
          // 正则检验前先将值转为字符串
          transform(value) {
            return String(value)
          },
          message: '邮箱格式不正确',
          trigger: ['blur', 'change']
        },
        complaintsReasons: {
          type: 'string',
          required: true,
          message: '请填写投诉理由',
          trigger: ['blur', 'change']
        }
      },
      fileList1: [],
      imgIds: []
    }
  },

  computed: {
    qrcode() {
      return this.$store.getters.code
    }
  },

  created() {},

  methods: {
    /**
     * @description: 提交按钮事件
     * @author: qinjj
     */
    submit() {
      let that = this
      this.$refs.form1
        .validate()
        .then(async () => {
          let params = {
            ...that.formData,
            tracingCode: that.qrcode,
            complaintImages: that.imgIds.join(',')
          }
          let result = await submitComplaintDetails(params)
          if (result.code === '0') {
            that.$refs.uToast.show({
              type: 'success',
              message: '提交成功',
              iconUrl: 'https://cdn.uviewui.com/uview/demo/toast/success.png',
              complete() {
                that.backBtn()
              }
            })
          }
        })
        .catch(() => {
          uni.$u.toast('校验失败')
        })
      console.log('提交', this.imgIds)
    },
    /**
     * @description: 新增图片
     * @author: qinjj
     */
    async afterRead(event) {
      // 当设置 mutiple 为 true 时, file 为数组格式，否则为对象格式
      let lists = [].concat(event.file)
      let fileListLen = this[`fileList${event.name}`].length
      lists.map(item => {
        this[`fileList${event.name}`].push({
          ...item,
          status: 'uploading',
          message: '上传中'
        })
      })
      for (let i = 0; i < lists.length; i++) {
        const result = await this.uploadFilePromise(lists[i].url)
        let item = this[`fileList${event.name}`][fileListLen]
        this[`fileList${event.name}`].splice(
          fileListLen,
          1,
          Object.assign(item, {
            status: 'success',
            message: '',
            url: result
          })
        )
        fileListLen++
        console.log('新增完成', this[`fileList1`])
      }
    },
    /**
     * @description: 删除图片
     * @author: qinjj
     */
    deletePic(event) {
      this[`fileList${event.name}`].splice(event.index, 1)
      this.imgIds.splice(event.index, 1)
    },
    /**
     * @description: 上传接口
     * @author: qinjj
     */
    uploadFilePromise(url) {
      let that = this
      return new Promise(resolve => {
        uni.uploadFile({
          url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/uploadFile/w/v1`, // 仅为示例，非真实的接口地址
          filePath: url,
          name: 'file',
          formData: {
            fileClass: 'complaintImages'
          },
          success: res => {
            if (res.statusCode === 200) {
              let data = JSON.parse(res.data)
              that.imgIds.push(data.id)
            }
            setTimeout(() => {
              resolve(res.data.data)
            }, 1000)
          }
        })
      })
    },
    /**
     * @description: 取消按钮
     * @author: qinjj
     */
    backBtn() {
      window.history.go(-1)
    }
  }
}
</script>

<style lang="scss" scoped>
.badComment {
  padding: 16rpx 32rpx;
  background: #fafafa;
  .image-box {
    .image-title {
      margin: 10px 0;
      font-size: 15px;
    }
  }
  .footer-box {
    display: flex;
    justify-content: space-around;
    padding-top: 30rpx;
    border-top: 1px solid rgb(214, 215, 217);
    .footer-btn {
      width: 30%;
    }
  }
}
</style>
