<!--
 * @Author: guoxi
 * @Date: 2021-05-24 14:06:36
 * @LastEditTime : 2022-07-08 18:31:31
 * @LastEditors  : Your Name
 * @Description: 商品添加页
 * @FilePath     : \blockchain-minicode\src\pages\commodity\AddCommodity.vue
-->

<template>
  <view class="content">
    <uni-forms
      ref="form"
      class="add-brand-form"
      :value="formData"
      label-width="100"
      border
    >
      <uni-forms-item label="商品名称" name="commodityName">
        <uni-row>
          <uni-col :span="24">
            <uni-easyinput
              v-model="formData.commodityName"
              type="text"
              :input-border="false"
              placeholder="请输入"
            />
          </uni-col>
        </uni-row>
      </uni-forms-item>
      <uni-forms-item label="所属品牌" name="brandType">
        <uni-row>
          <uni-col :span="24">
            <!-- <picker
              class="picker-item"
              :value="formData.brandRegCard"
              :range="brandList"
              :range-key="'brandName'"
              @change="onPickerChange"
            >
            </picker> -->
            <view class="picker-input">{{ brandName }}</view>
            <!-- <uni-icons type="forward" size="16"></uni-icons> -->
          </uni-col>
        </uni-row>
      </uni-forms-item>

      <uni-row>
        <uni-col :span="24">
          <uni-forms-item label="质量等级" name="qualityGrade">
            <uni-easyinput
              v-model="formData.qualityGrade"
              type="text"
              :input-border="false"
              placeholder="请输入"
            />
          </uni-forms-item>
        </uni-col>
      </uni-row>
      <uni-forms-item label="保质期" name="qualityGuaPeriod">
        <uni-row>
          <uni-col :span="22">
            <uni-easyinput
              v-model="formData.qualityGuaPeriod"
              type="digit"
              :input-border="false"
              placeholder="请输入"
            />
          </uni-col>
          <uni-col :span="2" style="text-align: right;">月</uni-col>
        </uni-row>
      </uni-forms-item>
      <uni-forms-item label="净含量" name="netContent">
        <uni-icons type="plus" size="30" @tap="addNet"></uni-icons>
      </uni-forms-item>

      <uni-row v-for="(item, index) in netContent" :key="index">
        <uni-col :span="2">
          <view v-if="index > 0" style="text-align: center;">
            <uni-icons type="clear" size="14" @tap="delNet(index)"></uni-icons>
          </view>
        </uni-col>
        <uni-col :span="4">
          <span style="font-size: 14px; color: #333;">规格{{ index + 1 }}</span>
        </uni-col>
        <uni-col :span="14">
          <uni-easyinput
            v-model="item.weightValue"
            type="digit"
            :input-border="false"
            placeholder="请输入"
          />
        </uni-col>
        <uni-col :span="4">
          <view style="text-align: center; font-size: 14px; color: #333;">
            kg /
            <picker
              style="display: inline-block;"
              :range="array"
              @change="changeUnit($event, item)"
            >
              <view class="picker"> {{ item.weightUnit }} </view>
            </picker>
          </view>
        </uni-col>
      </uni-row>
      <uni-forms-item
        class="images"
        label="商品图片（尺寸718*402，用于溯源宣传）"
        name="images"
      >
      </uni-forms-item>

      <view class="img-box">
        <view class="upload-container">
          <view v-for="(img, i) in formData.images" :key="i" class="upload-img">
            <image :src="img" @click="preViewImg(i)" />
            <uni-icons
              class="upload-icon"
              type="close"
              color="red"
              @click="handleDel(i)"
            />
          </view>
          <view class="upload" @click="chooseImg">
            <uni-icons class="upload-icon" type="plusempty" color="#f1f1f1" />
          </view>
        </view>
      </view>
    </uni-forms>

    <view class="add-brand">
      <button type="primary" class="btn" @tap="onAddCommodity">
        <span>确定{{ initBrandInfo.type === 'add' ? '添加' : '编辑' }}</span>
      </button>
    </view>
  </view>
</template>

<script>
import { chooseImg, preViewImg, getImageUrl } from '../../utils/tool'
import {
  addCommodity,
  updateCommodity,
  detailCommodity
  // getBrandList
} from '@/api/commodity'
export default {
  name: 'AddCommodity',
  data() {
    return {
      imgMaxNum: 10,
      imgPaths: [],
      array: ['袋', '箱'],
      netContent: [
        { id: new Date().getTime(), weightUnit: '袋', weightValue: 10 }
      ],
      brandType: 0,
      brandName: '',
      initBrandInfo: {
        type: 'add', // add: 新增, edit: 编辑
        brandId: 0,
        commodityId: 0
      },
      brandList: [],
      formData: {
        commodityAttachment: null,
        images: [],
        ids: [],
        commodityName: '',
        qualityGrade: '',
        qualityGuaPeriod: '',
        brandRegCard: 0
      },
      imagesArr: [],
      rules: {
        commodityName: {
          rules: [
            {
              required: true,
              errorMessage: '请输入品牌名称'
            }
          ]
        },
        brandType: {
          rules: [
            {
              required: true,
              errorMessage: '请选择所属品牌'
            }
          ]
        },
        netContent: {
          rules: [
            {
              required: true,
              errorMessage: '请输入净含量'
            }
          ]
        },
        qualityGrade: {
          rules: [
            {
              required: true,
              errorMessage: '请输入质量等级'
            }
          ]
        },
        qualityGuaPeriod: {
          rules: [
            {
              required: true,
              errorMessage: '请输入保质期'
            }
          ]
        },
        images: {
          rules: [
            // {
            //   required: false,
            //   errorMessage: '请至少上传一张照片'
            // },
            // {
            //   validateFunction: (rule, value, data, callback) => {
            //     console.log('value:', value)
            //     console.log('rule:', rule)
            //     console.log('data:', data)
            //     if (value.length === 0) {
            //       callback('请至少上传一张照片')
            //     }
            //     return true
            //   }
            // }
          ]
        }
      }
    }
  },

  onReady() {
    // 需要在onReady中设置规则
    this.$refs.form.setRules(this.rules)
    this.$refs.form.setValue('netContent', this.netContent)
  },

  onLoad(options) {
    this.initBrandInfo = {
      type: options.type,
      brandId: options.brandId
    }
    this.brandName = decodeURI(options.name)
    if (options.type === 'edit') {
      uni.setNavigationBarTitle({
        title: '编辑商品'
      })
      this.initBrandInfo.commodityId = options.commodityId
    }
  },

  mounted() {
    this.query()
  },

  methods: {
    chooseImg() {
      chooseImg(this.formData, this.imgMaxNum, this.imgPaths, this)
    },
    preViewImg(i) {
      preViewImg(i, this.formData.images)
    },

    handleDel(i) {
      uni.showModal({
        title: '提示',
        content: '是否要删除该图片',
        success: res => {
          if (res.confirm) return this.removePath(i)
        }
      })
    },

    removePath(index) {
      this.imgPaths.splice(index, 1)
      this.formData.images.splice(index, 1)
      this.formData.ids.splice(index, 1)
    },
    changeUnit(e, item) {
      item.weightUnit = this.array[parseInt(e.detail.value)]
    },
    addNet() {
      const netContent = {
        id: new Date().getTime(),
        weightUnit: '袋',
        weightValue: 0
      }
      this.netContent.push(netContent)
      this.$refs.form.setValue('netContent', netContent)
    },
    delNet(index) {
      this.netContent.splice(index, 1)
    },
    // async queryGoods() {
    //   const params = {
    //     brandNameLike: '',
    //     pageSize: '0',
    //     pageNum: '1'
    //   }
    //   let res = await getBrandList(params)
    //   this.brandList = res.list
    // },
    async query() {
      // await this.queryGoods()
      const { commodityId } = this.initBrandInfo
      if (commodityId) {
        let res = await detailCommodity(commodityId)
        if (res.code === '0') {
          this.netContent = JSON.parse(res.resObj.netContent)
          this.formData = { ...res.resObj, brandType: 0 }
          this.formData.images = []
          this.formData.ids = []
          if (this.formData.commodityAttachment) {
            this.formData.ids = JSON.parse(this.formData.commodityAttachment)
            this.formData.ids.map(el => {
              this.formData.images.push(getImageUrl(el.id))
            })
          }
        }
        this.brandType = this.brandList.findIndex(
          el => el.brandRegCard === res.resObj.brandRegCard
        )
      }
    },

    onPickerChange(e) {
      this.brandType = e.target.value
    },

    onAddCommodity() {
      this.$refs.form
        .submit()
        .then(res => {
          res.netContent = JSON.stringify(this.netContent)
          res['commodityAttachment'] = JSON.stringify(this.formData.ids)
          const { type } = this.initBrandInfo
          if (type === 'add') {
            this.addCommodity(res)
            return
          }
          this.updateCommodity(res)
        })
        .catch(err => {
          console.log(err)
        })
    },

    async addCommodity(params) {
      const { brandId } = this.initBrandInfo
      params = {
        ...params,
        brandId
      }
      let res = await addCommodity(params)
      if (res.code === '0') {
        uni.showToast({
          title: '新增成功',
          icon: 'none'
        })
        uni.navigateTo({
          url: `/pages/commodity/CommManagemant?id=${this.initBrandInfo.brandId}&name=${this.brandName}`
        })
        // uni.navigateBack({
        //   delta: 1,
        // })
        return
      }
      uni.showToast({
        title: '新增失败',
        icon: 'none'
      })
    },

    async updateCommodity(params) {
      const { commodityId: id, brandId } = this.initBrandInfo
      params = {
        ...params,
        id,
        brandId
      }
      let res = await updateCommodity(params)
      if (res.code === '0') {
        uni.showToast({
          title: '更新成功',
          icon: 'none'
        })
        uni.navigateTo({
          url: `/pages/commodity/CommManagemant?id=${this.formData.brandId}&name=${this.formData.brandName}`
        })
        // uni.navigateBack({
        //   delta: 1,
        // })
        return
      }
      uni.showToast({
        title: '更新失败',
        icon: 'none'
      })
    },

    openPopup(msg) {
      this.popMsg = msg
      this.$refs.hintPopup.open()
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

.uni-row {
  display: flex;
  align-items: center;
}

.uni-col-2 {
  text-align: right;
  margin-left: 18rpx;
}

.uni-forms-item__content .data-v-1359f286 {
  text-align: right;
  line-height: 36px;
}

.images .uni-forms-item__label {
  width: 100% !important;
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

  .picker-item {
    display: inline-block;
    width: 92%;
  }

  .picker-input {
    height: 72rpx;
    line-height: 72rpx;
    text-align: right;
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
