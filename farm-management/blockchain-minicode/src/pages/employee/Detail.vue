<!--
 * @Author: guoxi
 * @Date: 2021-05-25 15:28:53
 * @LastEditTime : 2021-09-08 20:45:38
 * @LastEditors  : Please set LastEditors
 * @Description: 员工详情页
 * @FilePath     : \blockchain-minicode\src\pages\employee\Detail.vue
-->

<template>
  <view class="content">
    <uni-forms
      ref="form"
      class="add-brand-form"
      :value="formData"
      :rules="rules"
      label-width="100"
      border
    >
      <uni-forms-item required label="员工姓名" name="name">
        <search-select
          v-model="formData.name"
          placeholder="请输入"
          :info-list="userList"
          :show-value="showValue"
          :loading="loading"
          type="primary"
          :uni-shadow="true"
          @handleSearch="handleSearch"
          @change="change"
        ></search-select>
        <!-- <uni-easyinput
          v-model="formData.name"
          :input-border="false"
          class="common-input"
          type="text"
          placeholder="请输入"
        /> -->
      </uni-forms-item>
      <uni-forms-item label="手机号" name="phone">
        <uni-easyinput
          v-model="formData.phone"
          disabled
          :input-border="false"
          class="common-input"
          type="digit"
        />
      </uni-forms-item>
      <uni-forms-item required label="账号类型" name="type">
        <picker
          class="picker-item"
          :value="formData.accountType"
          :range="array"
          range-key="name"
          @change="onPickerChange"
        >
          <view class="uni-input">{{
            array[formData.accountType]['name']
          }}</view>
        </picker>
        <uni-icons type="forward"></uni-icons>
      </uni-forms-item>
    </uni-forms>
    <view class="add-brand">
      <button class="btn" @tap="onSubmit">
        <span>确定{{ formData.type === 'add' ? '添加' : '编辑' }}</span>
      </button>
    </view>
    <uni-popup ref="createPopup" background-color="#fff" :before-close="true">
      <popup
        title="系统提示"
        :content="popContent"
        confirm-text="确定"
        :show-close="false"
        @close="close()"
        @confirm="onConfirm"
      />
    </uni-popup>
  </view>
</template>

<script>
import popup from '../../components/Popup'
import searchSelect from '../../components/search-select'
import { userAdd, updateUser, filterUser } from '@/api/employee'
export default {
  name: 'EmployeeDetail',
  components: {
    popup,
    searchSelect
  },
  data() {
    return {
      loading: false,
      // showValue: 'name', // 需要显示的数据，必须与infoList中的name对应
      searchValue: '',
      userList: [],
      popContent: '',
      array: [],
      types: [], // 存放选中的账号类型 id
      formData: {
        name: '',
        phone: '',
        accountType: 0,
        id: 0
      },
      rules: {
        name: {
          rules: [
            {
              required: true,
              errorMessage: '请输入员工姓名'
            }
          ]
        },
        type: {
          rules: [
            {
              required: true,
              errorMessage: '请选择账号类型'
            }
          ]
        }
        // phone: {
        //   rules: [
        //     {
        //       required: true,
        //       errorMessage: '请输入手机号',
        //     },
        //     {
        //       trigger: 'blur',
        //       validateFunction: function (rule, value, data, callback) {
        //         const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
        //         if (!reg.test(value)) {
        //           callback('请输入正确的手机号')
        //         }
        //         return true
        //       },
        //     },
        //   ],
        // },
      }
    }
  },

  computed: {
    accountMap() {
      return this.$store.state.user.dictMap.account_type
    }
  },

  onLoad(option) {
    const employeeData = JSON.parse(decodeURIComponent(option.employeeData))
    this.formData = { ...this.formData, ...employeeData }
    const { type } = this.formData
    Object.keys(this.accountMap).forEach(key => {
      this.array.push({
        id: key,
        name: this.accountMap[key]
      })
    })
    if (type === 'edit') {
      this.formData.accountType = this.array.findIndex(
        ele => ele.id === this.formData.personId
      )
    }
  },

  methods: {
    showValue(item) {
      return item.name + ':' + item.phone
    },
    async handleSearch(name) {
      this.loading = true
      const data = await filterUser(name)
      this.userList = data.resObj
      this.loading = false
      // setTimeout(() => {
      //   this.infoList = this.infoLists
      // }, 2000)
    },
    change(val) {
      this.formData.name = val.name
      this.formData.phone = val.phone

      console.log(val)
    },

    onPickerChange(e) {
      this.formData.accountType = e.target.value
    },
    onSubmit() {
      this.$refs.form
        .submit()
        .then(res => {
          console.log(res)
          const filterItem = this.array.filter(
            ele => ele['name'] === this.array[this.formData.accountType]['name']
          )[0]
          this.types = [`${filterItem['id']}`]
          if (this.formData.accountType === '3') {
            this.popContent = `确定将${this.formData.name}（${this.formData.phone}）设置为系统管理员？`
            this.$refs.createPopup.open()
          } else {
            this.onConfirm()
          }
        })
        .catch(err => {
          console.log('表单错误信息：', err)
        })
    },
    onConfirm() {
      if (this.formData.accountType === 0) this.$refs.createPopup.close()
      if (this.formData.type === 'add') return this.handleAddemployee()
      this.handleEditemployee()
    },

    async handleAddemployee() {
      const { name, phone } = this.formData
      let params = {
        name,
        phone,
        types: this.types ? this.types : []
      }
      let res = await userAdd(params)
      if (res.code === '0') {
        uni.showToast({
          title: '新增成功',
          icon: 'success'
        })
        uni.navigateTo({ url: `/pages/employee/index` })
        return
      }
      uni.showToast({
        title: '新增失败',
        icon: 'none'
      })
    },

    async handleEditemployee() {
      const { name, phone, id } = this.formData
      let params = {
        id,
        name,
        phone,
        types: this.types ? this.types : []
      }
      let res = await updateUser(params)
      if (res.code === '0') {
        uni.showToast({
          title: '编辑成功',
          icon: 'success'
        })
        uni.navigateTo({ url: `/pages/employee/index` })
        return
      }
      uni.showToast({
        title: '编辑失败',
        icon: 'none'
      })
    }
  }
}
</script>

<style lang="scss">
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-top: 1rpx solid #eee;
  border-bottom: 1rpx solid #eee;

  .uni-input {
    height: 72rpx;
    line-height: 72rpx;
    text-align: right;
  }

  .picker-item {
    display: inline-block;
    width: 92%;
  }

  .common-input {
    border: none;
    text-align: right;

    .uni-easyinput__content-input {
      text-align: right;
    }
  }

  .add-brand-form {
    width: 90%;
  }

  .add-brand {
    width: 750rpx;
    height: 138rpx;
    position: fixed;
    bottom: 0;
    background-color: #fafafa;

    .btn {
      border-radius: 44rpx;
      font-size: 32rpx;
      margin: 25rpx 40rpx;
      background-color: #00c853;

      span {
        font-family: PingFangSC-Regular;
        font-size: 32rpx;
        color: #fff;
        margin-left: 10rpx;
      }
    }
  }
}
</style>
