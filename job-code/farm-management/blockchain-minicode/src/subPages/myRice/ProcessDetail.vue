<!--
 * @Author: guoxi
 * @Date: 2021-05-25 15:28:53
 * @LastEditTime : 2021-07-08 14:04:27
 * @LastEditors  : Please set LastEditors
 * @Description: 溯源详情
 * @FilePath     : \blockchain-minicode\src\subPages\myRice\ProcessDetail.vue
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
      <uni-forms-item
        v-for="(item, index) in formData.formItems.ruleFormItems"
        :key="index"
        :required="item.required"
        :label="item.label"
        :name="item.name"
      >
        <view
          >{{ formData[item.name] }}
          <span v-if="item.unit" class="unit">{{ item.unit }}</span></view
        >
      </uni-forms-item>
      <uni-forms-item name="" label="照片"> </uni-forms-item>
      <view class="report-card">
        <view
          v-for="(item, index) in formData.images"
          :key="index"
          class="card-img"
          ><img :src="item" @click="preViewImg(index)"
        /></view>
      </view>
    </uni-forms>
  </view>
</template>

<script>
import { getProgressSrvInfo } from '@/api/myRice'
import {
  warEndFormItems,
  hgEndFormItems,
  jgFormItems,
  bzFormItems
} from '../../common/constant'
import { getImageUrl, preViewImg } from '../../utils/tool'
export default {
  name: 'ProcessDetail',

  onLoad(option) {
    const info = JSON.parse(decodeURIComponent(option.info))
    this.procInstId = info.procInstId
    this.stepType = info.stepType
    this.query()
  },

  data() {
    return {
      procInstId: '',
      stepType: '',
      formData: {
        ids: [],
        images: [],
        ccFactoryName: '',
        storageBatch: '',
        startDate: '',
        endDate: '',
        outWeight: '',
        enterWeight: '',
        hgBatch: '',
        hgFactoryName: '',
        hgDate: '',
        hgHsl: '',
        hgWeight: '',

        jgBatch: '',
        upFlowEnterpriseName: '',
        jgDate: '',
        ccWeight: '',

        packDate: '',
        machPackNormList: '',
        formItems: warEndFormItems
      }
    }
  },

  methods: {
    getImage(item) {
      return getImageUrl(item.id)
    },
    preViewImg(i) {
      preViewImg(i, this.formData.images)
    },
    async query() {
      const params = {
        procInstId: this.procInstId,
        stepType: this.stepType
      }
      let res = await getProgressSrvInfo(params)
      const data = res.resObj
      if (this.stepType === 'CC') {
        uni.setNavigationBarTitle({
          title: '仓储信息'
        })
        this.formData.ccFactoryName = data.ccFactoryName
        this.formData.storageBatch = data.id
        this.formData.startDate = data.startDate
        this.formData.endDate = data.endDate
        this.formData.outWeight = data.outWeight
        this.formData.enterWeight = data.enterWeight
        this.formData.formItems = warEndFormItems
        this.formData.ids = data.ccStartAttachJson || []
      } else if (this.stepType === 'HG') {
        uni.setNavigationBarTitle({
          title: '烘干信息'
        })
        this.formData.formItems = hgEndFormItems
        this.formData.hgBatch = data.hgBatch
        this.formData.hgFactoryName = data.hgFactoryName
        this.formData.hgDate = data.hgDate
        this.formData.hgHsl = data.hgHsl
        this.formData.hgWeight = data.hgWeight
        this.formData.ids = data.hgAttachJson || []
      } else if (this.stepType === 'JG') {
        uni.setNavigationBarTitle({
          title: '加工信息'
        })

        this.formData.formItems = jgFormItems
        this.formData.jgBatch = data.jgBatch
        this.formData.upFlowEnterpriseName = data.upFlowEnterpriseName
        this.formData.jgDate = data.jgDate
        this.formData.ccWeight = data.ccWeight
        this.formData.hgWeight = data.hgWeight
        this.formData.ids = data.jgAttachJson || []
      } else {
        uni.setNavigationBarTitle({
          title: '包装信息'
        })
        this.formData.formItems = bzFormItems
        this.formData.packDate = data.packDate
        this.formData.upFlowEnterpriseName = data.upFlowEnterpriseName
        this.formData.machPackNormList = data.normsForMachPack
        this.formData.ids = data.packAttachJson || []
      }

      this.formData.ids.map(el => {
        this.formData.images.push(this.getImage(el))
      })
      // if (res.code === '0') {
      //   this.processList = res.list
      // }
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

  .report-card {
    padding: 32rpx;
    background-color: #fff;

    .card-img {
      display: inline-block;
      margin-right: 32rpx;
      width: 112rpx;
      height: 112rpx;

      img {
        width: 100%;
        height: 100%;
      }
    }
  }
}
</style>
