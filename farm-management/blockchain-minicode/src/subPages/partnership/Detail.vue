<!--
 * @Description  : 合作伙伴详情页
 * @Autor        : guoxi
 * @Date         : 2021-06-20 17:01:00
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-08-07 18:18:37
 * @FilePath     : \blockchain-minicode\src\subPages\partnership\Detail.vue
-->
<template>
  <view class="application">
    <view class="form-wrapper">
      <uni-forms ref="form" label-width="140" border>
        <uni-group>
          <uni-forms-item
            class="form-item enterprise"
            label="企业类型"
            name="enterpriseType"
          >
            {{ enterprise.enterpriseText }}
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="企业名称"
            name="enterpriseName"
          >
            {{ enterprise.name }}
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="统一社会信用代号"
            label-width="150"
            name="creditCode"
          >
            {{ enterprise.uuid }}
          </uni-forms-item>

          <uni-forms-item class="form-item" label="法人代表" name="legalPerson">
            {{ enterprise.legalPerson }}
          </uni-forms-item>

          <uni-forms-item
            class="form-item"
            label="企业联系电话"
            name="phoneNum"
          >
            {{ enterprise.officePhone }}
          </uni-forms-item>
        </uni-group>

        <uni-group class="plain-cell">
          <uni-forms-item class="form-item" label="企业联系地址">
          </uni-forms-item>
        </uni-group>
        <uni-group class="horizon-cell">
          <uni-forms-item class="form-item" name="address">
            {{ enterprise.address }}
          </uni-forms-item>
        </uni-group>

        <uni-group
          v-if="
            enterprise.attachments && enterprise.attachments.businessLicense
          "
          class="plain-cell"
        >
          <uni-forms-item class="form-item" label="企业营业执照">
          </uni-forms-item>
        </uni-group>
        <uni-group
          v-if="
            enterprise.attachments && enterprise.attachments.businessLicense
          "
          class="horizon-cell"
        >
          <uni-forms-item class="form-item">
            <view
              v-for="(item, index) in enterprise.attachments.businessLicense"
              :key="index"
              class="card-img"
              ><img
                :src="getImage(item)"
                @tap="preViewImg(i, enterprise.attachments.businessLicense)"
            /></view>
          </uni-forms-item>
        </uni-group>

        <uni-group
          v-if="
            enterprise.attachments &&
              enterprise.attachments.riceProcessingLicense
          "
          class="plain-cell"
        >
          <uni-forms-item class="form-item" label="大米生产加工许可证">
          </uni-forms-item>
        </uni-group>

        <uni-group
          v-if="
            enterprise.attachments &&
              enterprise.attachments.riceProcessingLicense
          "
          class="horizon-cell"
        >
          <uni-forms-item class="form-item">
            <view
              v-for="(item, index) in enterprise.attachments
                .riceProcessingLicense"
              :key="index"
              class="card-img"
              ><img
                :src="getImage(item)"
                @tap="
                  preViewImg(i, enterprise.attachments.riceProcessingLicense)
                "
            /></view>
          </uni-forms-item>
        </uni-group>

        <uni-group
          v-if="enterprise.attachments && enterprise.attachments.dryEnvironment"
          class="plain-cell"
        >
          <uni-forms-item class="form-item" label="加工环境"> </uni-forms-item>
        </uni-group>
        <uni-group
          v-if="enterprise.attachments && enterprise.attachments.dryEnvironment"
          class="horizon-cell"
        >
          <uni-forms-item class="form-item" name="operatingEnvironment">
            <view
              v-for="(item, index) in enterprise.attachments.dryEnvironment"
              :key="index"
              class="card-img"
              ><img
                :src="getImage(item)"
                @tap="preViewImg(i, enterprise.attachments.dryEnvironment)"
            /></view>
          </uni-forms-item>
        </uni-group>
        <uni-group class="plain-cell">
          <p class="form-item" style="height: 20rpx;"></p>
        </uni-group>
        <uni-group>
          <uni-forms-item
            class="form-item enterprise"
            label="审核结果"
            name="a"
          >
            <span class="enterprise-active">合格</span>
          </uni-forms-item>
        </uni-group>
      </uni-forms>
      <view v-if="isApprove" class="btn-tools">
        <view class="cancel btn" @tap="cooOrNon('2')">不合作</view>
        <view class="submit btn" @tap="cooOrNon('1')">确认合作</view>
      </view>
    </view>
  </view>
</template>

<script>
import { getEnterpriseById, cooOrNon } from '@/api/partnership'
import { preViewImg } from '../../utils/tool'
export default {
  name: 'PartnershipDetail',

  data() {
    return {
      id: '',
      isApprove: false,
      businessId: '',
      enterprise: {
        name: '',
        type: '',
        enterpriseText: '',
        uuid: '',
        legalPerson: '',
        officePhone: '',
        address: '',
        attachments: {}
      }
    }
  },

  computed: {
    /**
     * @description: 企业类型map
     * @author: Hemingway
     */
    enterpriseMap() {
      return this.$store.state.user.dictMap.enterprise_type
    }
  },

  mounted() {
    this.isApprove = this.$Route.query.isApprove
    this.businessId = this.$Route.query.id // 业务id
    this.id = this.$Route.query.enterpriseId // 企业id
    this.queryEnterpriseDetail(this.$Route.query.enterpriseId)
  },

  methods: {
    getImage(item) {
      return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${item.id}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
    },
    preViewImg(i, images) {
      const urls = images.map(el => this.getImage(el))
      preViewImg(i, urls)
    },

    /**
     * @description: 查询企业详情
     * @param {String} id 企业id
     * @author: Hemingway
     */
    async queryEnterpriseDetail(id) {
      try {
        const { code, enterpriseDto } = await getEnterpriseById({
          EnterpriseId: id
        })
        if (code === '0') {
          enterpriseDto.enterpriseText = enterpriseDto.type
            .split(',')
            .map(type => this.enterpriseMap[type])
            .join(' ')
          this.enterprise = enterpriseDto
        }
      } catch (error) {
        console.log('查询企业详情 error = ', error)
      }
    },

    async cooOrNon(status) {
      const { code } = await cooOrNon({
        id: this.businessId,
        cooStatus: status
      })
      if (code === '0') {
        uni.showToast({
          title: '已审批完成！',
          icon: 'success',
          duration: 2000,
          complete: () => {
            setTimeout(() => {
              this.$Router.back()
            }, 1500)
          }
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.application {
  ::v-deep .uni-forms-item__content {
    line-height: 72rpx;
  }
  // 表单域分组
  ::v-deep .uni-group {
    margin-top: 0 !important;

    .group-conent-padding {
      padding: 0 16rpx;
    }
  }

  // 表单域
  ::v-deep .uni-forms-item {
    border: none;
    padding: 16rpx 0;
    position: relative;
  }

  .form-item:not(:last-child) ::v-deep .uni-forms-item::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    border-bottom: 1rpx solid #eee;
    transform: scaleY(0.5);
  }

  // 企业选择
  .enterprise ::v-deep .uni-forms-item__content {
    display: flex;
    align-items: center;
    justify-content: flex-end;

    .picker-cell {
      margin-right: 8rpx;
    }
  }

  // 整行表单域
  .horizon-cell {
    ::v-deep .uni-forms-item {
      .uni-forms-item__label {
        display: none;
      }

      .uni-error-message {
        padding-left: 0 !important;
        text-align: left;
      }
    }

    ::v-deep .uni-easyinput {
      .uni-easyinput__content {
        text-align: left;

        .uni-easyinput__content-textarea {
          min-height: 84rpx;
        }
      }
    }
  }

  // 普通简单行
  .plain-cell ::v-deep .uni-group {
    background: #f5f5f5;

    .uni-forms-item {
      padding: 4rpx 0;
    }

    .label-text {
      font-size: 24rpx;
    }
  }
}

.enterprise-active {
  color: #00c853;
}

.btn-tools {
  margin: 65rpx 0 25rpx 0;
  text-align: center;

  .btn {
    display: inline-block;
    width: 315rpx;
    height: 88rpx;
    line-height: 88rpx;
    background-color: #fff;
    border-radius: 44rpx;
    border: solid 1rpx #eee;
    text-align: center;
  }

  .cancel {
    margin-right: 40rpx;
  }

  .submit {
    background-color: #00c853;
    color: #fff;
  }
}

.card-img {
  /* display: inline-block; */
  float: left;
  margin-right: 32rpx;
  width: 112rpx;
  height: 112rpx;

  img {
    width: 100%;
    height: 100%;
  }
}
</style>
