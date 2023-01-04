<!--
 * @Author: guoxi
 * @Date: 2021-05-24 14:06:36
 * @LastEditTime : 2021-06-21 16:57:32
 * @LastEditors  : Please set LastEditors
 * @Description: 品牌添加
 * @FilePath     : \blockchain-minicode\src\pages\commodity\AddBrand.vue
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
      <uni-forms-item label="品牌名称" name="brandName">
        <uni-easyinput
          v-model="formData.brandName"
          type="text"
          placeholder="请输入"
          :input-border="false"
        />
      </uni-forms-item>
      <uni-forms-item label="商标注册证号" name="brandRegCard">
        <uni-easyinput
          v-model="formData.brandRegCard"
          type="text"
          placeholder="请输入"
          :input-border="false"
        />
      </uni-forms-item>
    </uni-forms>
    <view class="add-brand">
      <button type="primary" class="btn" @tap="onAddBrand">
        <!-- <uni-icons type="plusempty" color="#fff"></uni-icons> -->
        <span>确定{{ type === 'add' ? '添加' : '编辑' }}</span>
      </button>
    </view>
  </view>
</template>

<script>
import { addBrand, updateBrand } from '@/api/commodity'
export default {
  name: 'AddBrand',
  data() {
    return {
      type: 'add', // edit: 编辑, add: 新增
      brandId: null,
      formData: {
        brandName: '',
        brandRegCard: ''
      },
      rules: {
        brandName: {
          rules: [
            {
              required: true,
              errorMessage: '请输入品牌名称'
            }
          ]
        },
        brandRegCard: {
          rules: [
            {
              required: true,
              errorMessage: '请输入商标注册证号'
            }
          ]
        }
      }
    }
  },

  onLoad(option) {
    this.type = option.type
    if (option.type === 'edit') {
      uni.setNavigationBarTitle({
        title: '编辑品牌'
      })
      const brandInfo = JSON.parse(option.brandInfo)
      this.formData = { ...brandInfo }
      this.brandId = brandInfo.id
      return
    }
    this.formData = { brandName: '', brandRegCard: '' }
  },

  mounted() {},

  methods: {
    onAddBrand() {
      this.$refs.form
        .submit()
        .then(res => {
          if (this.type === 'add') {
            this.addBrand(res)
            return
          }
          this.updateBrand(res)
        })
        .catch(err => {
          console.log(err)
        })
    },

    async addBrand(params) {
      // this.$http
      //   .post('http://10.33.2.224:8902/iablockchain/srvBrand/add', {
      //     ...params,
      //   })
      //   .then((res) => {
      //     const { statusCode, dataResult } = res.data
      //     if (statusCode === '0') {
      // uni.showToast({
      //   title: '新增成功',
      //   icon: 'none',
      // })
      // uni.navigateBack({
      //   delta: 1,
      // })
      //     }
      //   })
      //   .catch(() => {
      //     uni.showToast({
      //       title: '新增失败',
      //       icon: 'none',
      //     })
      //   })
      let res = await addBrand(params)
      if (res.code === '0') {
        uni.showToast({
          title: '新增成功',
          icon: 'none'
        })
        uni.navigateBack({
          delta: 1
        })
        return
      }
      uni.showToast({
        title: '新增失败',
        icon: 'none'
      })
    },

    async updateBrand(params) {
      params = {
        ...params,
        id: this.brandId
      }
      let res = await updateBrand(params)
      if (res.code === '0') {
        uni.showToast({
          title: '更新成功',
          icon: 'none'
        })
        uni.navigateBack({
          delta: 1
        })
        return
      }
      uni.showToast({
        title: '更新失败',
        icon: 'none'
      })
      // this.$http
      //   .post('http://10.33.2.224:8902/iablockchain/srvBrand/add', {
      //     ...params,
      //   })
      //   .then((res) => {
      //     const { statusCode, dataResult } = res.data
      //     if (statusCode === '0') {
      //       uni.showToast({
      //         title: '更新成功',
      //         icon: 'none',
      //       })
      //       uni.navigateBack({
      //         delta: 1,
      //       })
      //     }
      //   })
      //   .catch(() => {
      //     uni.showToast({
      //       title: '更新失败',
      //       icon: 'none',
      //     })
      //   })
    }
  }
}
</script>

<style>
.uni-easyinput__content-input.data-v-20076044 {
  text-align: right;
}

.uni-forms-item__label .label-text.data-v-1359f286 ::before {
  content: '*';
  color: #dd524d;
}
</style>

<style lang="scss" scoped>
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-top: 1rpx solid #eee;
  border-bottom: 1rpx solid #eee;

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
        margin-left: 10rpx;
      }
    }
  }
}
</style>
