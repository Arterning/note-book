<!--
 * @Description  : 商品详情(编辑)
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-01 16:08:03
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-06 19:08:21
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\details.vue
-->
<template>
  <div class="details">
    <i-tab>
      <div>
        <goods-base-info-item
          :data="data"
          :showEditBtn="true"
          :inputDisable.sync="iptDiasble"
        >
          <div class="hr-style"></div>
        </goods-base-info-item>
        <size-params-item
          :data="data"
          :showEditBtnParams="true"
          :inputDisableParams.sync="iptDisableParams"
          mode="view"
        >
          <div class="hr-style"></div>
        </size-params-item>
        <goods-introduce-item
          :data="data"
          :showEditBtnIntroduce="true"
          :inputDisableIntroduce.sync="iptDiasbleIntroduce"
        >
          <div class="hr-style"></div>
        </goods-introduce-item>
      </div>

      <i-footer>
        <div style="text-align: center">
          <el-button size="small" @click="clickRetrunBTn">返 回</el-button>
        </div>
      </i-footer>
    </i-tab>
  </div>
</template>

<script>
import GoodsBaseInfoItem from './components/GoodsBaseInfoItem.vue'
import SizeParamsItem from './components/SizeParamsItem.vue'
import GoodsIntroduceItem from './components/GoodsIntroduceItem.vue'
import { getBrandCommodityInfo } from '@/api/basedata-manage'

export default {
  name: 'details',
  components: {
    GoodsBaseInfoItem,
    GoodsIntroduceItem,
    SizeParamsItem
  },
  props: {
    id: {
      required: true
    }
  },
  data() {
    return {
      iptDiasble: true,
      iptDisableParams: true,
      iptDiasbleIntroduce: true,
      data: {
        brandCommodityId: '', //商品id
        field_5: '', //品牌id
        type: '',
        commodityName: '',
        qualityGrade: '',
        standard: '',
        qualityGuaPeriod: '',
        unitList: [
          {
            weightUnit: '',
            weightValue: '',
            price: '',
            commodityAttachment: ''
          }
        ],
        commodityMaterial: {
          commodityDesc: '',
          videoLink: '',
          videoCoverageAttachmentId: ''
        }
      }
    }
  },
  created() {
    this.getBrandCommodityInfo()
  },
  methods: {
    async getBrandCommodityInfo() {
      const { data } = await getBrandCommodityInfo({
        brandCommodityId: this.id
      })
      this.data = data
      console.log(112233, this.data)
    },
    clickRetrunBTn() {
      this.$router.back()
      this.$store.dispatch('tagsView/delView', this.$route)
    }
  }
}
</script>

<style scoped lang="scss">
.details {
  height: 100%;
  background-color: #fff;
  position: relative;

  .container {
    margin-bottom: 56px;
    width: 90%;
    overflow-y: auto;
  }
  .base-info {
    ::v-deep .el-descriptions-item__label:not(.is-bordered-label) {
      line-height: 33px;
      font-size: 13px;
      color: #000;
    }
  }
  .size-params {
    font-size: 13px;
    ::v-deep .el-form-item--medium .el-form-item__content {
      line-height: 28px;
    }
  }
  .goods-introduce {
    font-size: 13px;
    ::v-deep .el-form-item--medium .el-form-item__content {
      font-size: 13px;
    }
  }
}

.remark {
  font-size: 12px;
  margin: 6px 0 0 10px;
  color: #606266;
}

.hr-style {
  width: 100%;
  height: 1px;
  background-image: linear-gradient(
    to right,
    #ccc 0%,
    #ccc 50%,
    transparent 50%
  );
  background-size: 14px 1px;
  background-repeat: repeat-x;
  margin-top: -8px;
  margin-bottom: 10px;
}
</style>

<style scoped>
::v-deep .i-section-header .left-side .title {
  font-size: 13px;
}

::v-deep .el-input-group__append {
  color: #000;
  width: 10px !important;
}

/* 修改i-upload样式  */
::v-deep .el-upload--picture-card {
  border: 1px solid #c0ccda;
  width: 70px;
  height: 70px;
  line-height: 70px;
}
</style>
