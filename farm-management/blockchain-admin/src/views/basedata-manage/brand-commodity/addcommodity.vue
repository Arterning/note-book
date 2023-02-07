<!--
 * @Description  : 新增商品
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-01 16:27:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-08 16:51:17
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\addcommodity.vue
-->
<template>
  <div class="add-commodity">
    <i-tab>
      <goods-base-info-item :data="data" ref="goodsBase"></goods-base-info-item>
      <size-params-item
        :data="data"
        @getFileIds="getFileIds"
        @uploadOver="uploadOver"
      ></size-params-item>
      <goods-introduce-item
        :data="data"
        @getFileIds="getFileIds"
        @uploadOver="uploadOver"
      ></goods-introduce-item>

      <i-footer>
        <div style="text-align: center">
          <el-button style="margin-right: 20px" size="mini" @click="onCancel"
            >取 消</el-button
          >
          <el-button
            type="primary"
            size="mini"
            @click="clickAdd"
            :loading="addLoading"
            >添 加</el-button
          >
        </div></i-footer
      >
    </i-tab>
  </div>
</template>

<script>
import GoodsBaseInfoItem from './components/GoodsBaseInfoItem.vue'
import SizeParamsItem from './components/SizeParamsItem.vue'
import GoodsIntroduceItem from './components/GoodsIntroduceItem.vue'
import { saveCommodity } from '@/api/basedata-manage'

export default {
  name: 'addcommodity',
  components: {
    GoodsBaseInfoItem,
    GoodsIntroduceItem,
    SizeParamsItem
  },
  data() {
    return {
      addLoading: false,
      data: {
        brandId: '',
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
            commodityAttachment: '',
            limit: 1,
            limitSize: 1,
            params: { fileClass: 'enterpriseImages' }
          }
        ],
        commodityMaterial: {
          commodityDesc: '',
          videoLink: '',
          videoCoverageAttachmentId: ''
        },

        guigeFileInfo: null,
        goodFileInfo: null
      },
      filesIdArr: [] //参数规格图片的id数组
    }
  },
  methods: {
    uploadOver(e) {
      console.log('哈哈', e)
      this.addLoading = !e
    },
    getFileIds(filesArr, type) {
      if (!filesArr.length) return
      if (type === 'guige') {
        this.filesIdArr.push(filesArr[0].fileId)

        this.guigeFileInfo = filesArr[0]
      } else if (type === 'goods') {
        this.data.commodityMaterial.videoCoverageAttachmentId =
          filesArr[0].fileId
      }
    },

    async clickAdd() {
      await this.$refs.goodsBase.$refs.formData.validate()
      try {
        // 将获取到的图片id循环遍历给commodityAttachment字段，filesIdArr获取到的上传
        for (let i = 0; i < this.data.unitList.length; i++) {
          this.data.unitList[i].commodityAttachment = this.filesIdArr[i]

          if (!this.data.unitList[i].weightValue) {
            this.$message.warning(`规格参数，第${i + 1}行,规格不能为空`)
            return
          }
          if (!this.data.unitList[i].commodityAttachment) {
            this.$message.warning(`规格参数，第${i + 1}行,商品图片不能为空`)
            return
          }
          if (!this.data.unitList[i].price) {
            this.$message.warning(`规格参数，第${i + 1}行,商品单价不能为空`)
            return
          }
        }
        this.addLoading = true
        const { code } = await saveCommodity(this.data)
        this.addLoading = false
        if (code === 0 || code === '0') {
          this.$notify({
            title: '成功',
            message: '新增商品成功！',
            type: 'success',
            duration: 2000
          })
          this.$router.back()
          this.$store.dispatch('tagsView/delView', this.$route)
        }
      } catch (error) {
        this.addLoading = false
        console.log('新增商品失败', error)
      }
    },
    onCancel() {
      this.$router.back()
      this.$store.dispatch('tagsView/delView', this.$route)
    }
  }
}
</script>

<style lang="scss" scoped>
.add-commodity {
  height: 100%;
  background-color: #fff;
  position: relative;

  .container {
    margin-bottom: 56px;
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
