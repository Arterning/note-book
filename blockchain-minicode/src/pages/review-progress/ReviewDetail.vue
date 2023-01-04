<!--
 * @Description  : 审核详情页（只读：审核中/通过/未通过/取消资格）
 * @Autor        : Hemingway
 * @Date         : 2021-07-22 11:51:14
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-14 11:35:55
 * @FilePath     : \blockchain-minicode\src\pages\review-progress\ReviewDetail.vue
-->
<template>
  <view class="review-detail">
    <info :status="status" :note="note"> </info>

    <view class="plain-cell">申请信息</view>
    <view class="panel">
      <view class="row">
        <text class="label">企业类型：</text>
        <text class="value">{{ enterpriseMap[type] }}</text>
      </view>
      <view class="row">
        <text class="label">企业名称：</text>
        <text class="value">{{ enterpriseName }}</text>
      </view>
      <view class="row">
        <text class="label">法人代表：</text>
        <text class="value">{{ legalPerson }} </text>
      </view>
      <view class="row">
        <text class="label">联系电话：</text>
        <text class="value">{{ phoneNum }}</text>
      </view>
      <view class="row">
        <text class="label">联系地址：</text>
        <text class="value">{{ address }}</text>
      </view>
      <view class="row">
        <text class="label">统一社会信用代号：</text>
        <text class="value">{{ creditCode }}</text>
      </view>
    </view>

    <view class="plain-cell">企业资质</view>
    <view class="panel">
      <view class="certificate">
        <view class="title">法人代表身份证（人像面）</view>
        <image-upload readonly :max="1" :list="previewMap.idCardFront" />
      </view>
      <view class="certificate">
        <view class="title">法人代表身份证（国徽面）</view>
        <image-upload readonly :max="1" :list="previewMap.idCardReverse" />
      </view>
      <view class="certificate">
        <view class="title">企业营业执照</view>
        <image-upload readonly :max="1" :list="previewMap.businessLicense" />
      </view>

      <!-- 品牌商资质 start -->
      <template v-if="isBrandOwner">
        <view class="certificate">
          <view class="title">品牌商标注册证</view>
          <image-upload
            readonly
            :max="1"
            :list="previewMap.registrationCertificate"
          />
        </view>

        <view class="certificate">
          <view class="title">食品经营许可证</view>
          <image-upload
            readonly
            :max="1"
            :list="previewMap.foodBusinessLicense"
          />
        </view>
      </template>
      <!-- 品牌商资质 end -->

      <!-- 合作商资质 start -->
      <template v-else>
        <view class="certificate">
          <view class="title">食品生产许可证</view>
          <image-upload
            readonly
            :max="1"
            :list="previewMap.riceProcessingLicense"
          />
        </view>

        <view class="certificate">
          <view class="title">工厂操作环境</view>
          <!-- 烘干、仓储、加工（加工+包装） -->
          <view style="display: flex;">
            <!-- 烘干 环境 -->
            <view v-if="['6', '3'].includes(type)" class="image-wrapper">
              <image-upload
                readonly
                :max="1"
                :list="previewMap.dryEnvironment"
              />
              <text class="text">烘干</text>
            </view>

            <!-- 仓储 环境 -->
            <view v-if="['6', '4'].includes(type)" class="image-wrapper">
              <image-upload
                readonly
                :max="1"
                :list="previewMap.storageEnvironment"
              />
              <text class="text">仓储</text>
            </view>

            <!-- 加工+包装 环境 -->
            <template v-if="['6', '5'].includes(type)">
              <view class="image-wrapper">
                <image-upload
                  readonly
                  :max="1"
                  :list="previewMap.processingEnvironment"
                />
                <text class="text">加工</text>
              </view>
              <view class="image-wrapper">
                <image-upload
                  readonly
                  :max="1"
                  :list="previewMap.packageEnvironment"
                />
                <text class="text">包装</text>
              </view>
            </template>
          </view>
        </view>
      </template>
      <!-- 合作商资质 end -->
    </view>

    <!-- 未通过，发起整改后重新申请 -->
    <bottom-button
      v-if="status === '3'"
      text="修改申请"
      @click="
        $Router.push({
          name: 'application',
          query: { id, handleType: 'update' }
        })
      "
    ></bottom-button>
  </view>
</template>

<script>
import ImageUpload from '@/components/image-upload'
import BottomButton from '@/components/bottom-button'
import Info from './components/Info.vue'
import { queryApplicationDetail } from '@/api/user'

export default {
  name: 'ReviewDetail',

  components: {
    ImageUpload,
    BottomButton,
    Info
  },

  data() {
    return {
      id: '', // 审核记录的id
      status: '', // 审核状态

      note: '', // 备注
      type: '', // 企业类型
      enterpriseName: '', // 企业名称
      legalPerson: '', // 法人代表
      phoneNum: '', // 联系电话
      address: '', // 联系地址
      creditCode: '', // 统一社会信用代号

      // 资质图片预览
      previewMap: {
        idCardFront: [],
        idCardReverse: [],
        businessLicense: [],

        registrationCertificate: [], // 品牌商标注册证
        foodBusinessLicense: [], // 食品经营许可证

        riceProcessingLicense: [], // 食品生产许可证
        dryEnvironment: [], // 烘干
        storageEnvironment: [], // 仓储
        processingEnvironment: [], // 加工
        packageEnvironment: [] // 包装
      } // 图片预览
    }
  },

  computed: {
    /**
     * @description: 企业类型map
     * @author: Hemingway
     */
    enterpriseMap() {
      return this.$store.state.user.dictMap.enterprise_type
    },

    /**
     * @description: 企业类型是否是品牌商
     * @return {Boolean}
     * @author: Hemingway
     */
    isBrandOwner() {
      return this.type === '2'
    },

    /**
     * @description: 上传图片时fileClass的map
     * @author: Hemingway
     */
    fileClassMap() {
      return this.$store.state.user.dictMap.fileClass
    }
  },

  mounted() {
    this.id = this.$Route.query.id
    this.queryDetail(this.id)
  },

  methods: {
    /**
     * @description: 查询入驻审核详情
     * @param {String} id
     * @author: Hemingway
     */
    async queryDetail(id) {
      try {
        const { code, enterpriseChangeDto } = await queryApplicationDetail({
          id
        })
        if (code === '0') {
          const {
            checkStatus,
            checkMemo,
            type,
            name,
            legalPerson,
            officePhone,
            address,
            uuid,
            attachments
          } = enterpriseChangeDto
          this.status = checkStatus
          this.note = checkMemo
          this.type = type
          this.enterpriseName = name
          this.legalPerson = legalPerson
          this.phoneNum = officePhone
          this.address = address
          this.creditCode = uuid

          // 图片资源准备
          for (let key in attachments) {
            const idList = attachments[key].map(item => item.id)
            const previewList = idList.map(id => this.getImageUrl(id))
            this.previewMap[key] = previewList
          }
        }
      } catch (error) {
        console.log('入驻审核详情查询 error = ', error)
      }
    },

    /**
     * @description: 获取图片下载url
     * @param {String} id
     * @return {String} url
     * @author: Hemingway
     */
    getImageUrl(id) {
      return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
    }
  }
}
</script>

<style lang="scss" scoped>
.review-detail {
  .plain-cell {
    height: 80rpx;
    line-height: 80rpx;
    padding-left: 32rpx;
    font-size: 24rpx;
    color: #8b8b8b;
    // background-color: $uni-bg-color-page;
  }

  .panel {
    padding: 40rpx 32rpx;
    background-color: #fff;

    // 基本信息
    .row {
      font-size: 28rpx;

      .label {
        color: #8b8b8b;
      }

      .value {
        color: #212121;
      }
    }

    .row:not(:first-of-type) {
      margin-top: 16rpx;
    }

    // 企业资质
    .certificate {
      .title {
        margin-bottom: 24rpx;
        font-size: 28rpx;
        color: #212121;
      }

      // 工厂环境
      .image-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;

        .text {
          margin-top: 16rpx;
          font-size: 24rpx;
          color: #8b8b8b;
        }
      }
    }

    .certificate:not(:last-of-type) {
      margin-bottom: 24rpx;
      padding-bottom: 36rpx;
      position: relative;
    }

    .certificate:not(:last-of-type)::after {
      content: '';
      border-bottom: 1rpx solid #eee;
      transform: scaleY(0.5);
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
    }
  }
}
</style>
