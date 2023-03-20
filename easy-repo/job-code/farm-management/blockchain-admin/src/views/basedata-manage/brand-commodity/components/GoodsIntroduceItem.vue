<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-11 17:12:33
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 10:36:25
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\components\GoodsIntroduceItem.vue
-->
<template>
  <div class="goods-introduce">
    <i-section-header title="商品介绍">
      <div slot="action" v-if="showEditBtnIntroduce === true">
        <el-row v-if="isWrite === true">
          <el-button type="text" size="mini" @click="clickCancel"
            >取消</el-button
          >
          <el-button
            type="text"
            size="mini"
            @click="clickSave"
            :disabled="canSave"
            >保存</el-button
          >
        </el-row>
        <el-button
          slot="action"
          type="text"
          size="mini"
          v-else
          @click="clickWrite"
          >编辑</el-button
        >
      </div>
    </i-section-header>
    <slot></slot>

    <el-form :model="data" ref="formData" :rules="rules">
      <span>文字介绍</span><br />
      <el-form-item>
        <el-input
          v-model="data.commodityMaterial.commodityDesc"
          type="textarea"
          :disabled="inputDisableIntroduce"
          maxlength="500"
          show-word-limit
          row="4"
          style="width: 75%"
          placeholder="请输入商品文字介绍,如:芜湖数字大米是......"
        ></el-input>
      </el-form-item>

      <span>视频介绍</span>
      <el-row>
        <el-col :span="11">
          <el-form-item>
            <span>视频链接</span>
            <span class="remark"
              >（尺寸：16:9，请将视频转化为链接，链接长度不超过100个字符）</span
            >
            <el-input
              v-model="data.commodityMaterial.videoLink"
              type="textarea"
              maxlength="150"
              show-word-limit
              rows="2"
              style="width: 70%"
              :disabled="inputDisableIntroduce"
            >
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item>
            <span>视频封面图片</span>
            <span class="remark">（尺寸和视频保持一致，1M以内）</span>
            <i-upload
              @success="uploadSuccess"
              ref="uploadRef"
              uploadUrl="/blockchaincomponent/file/uploadFile/w/v1"
              :params="params"
              :limit="limit"
              :limitSize="limitSize"
              :disabled="inputDisableIntroduce"
              accept=".jpg,.png"
              :files="filesArr"
              @uploadOver="uploadOver"
            ></i-upload>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import { updateCommodity } from '@/api/basedata-manage'

export default {
  name: 'GoodsIntroduceItem',
  props: {
    showEditBtnIntroduce: {
      type: Boolean,
      default: false
    },
    inputDisableIntroduce: {
      type: Boolean,
      default: false
    },
    data: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      isWrite: false,
      limit: 1,
      limitSize: 1,
      params: { fileClass: 'enterpriseImages' },
      filesArr: [],
      rules: {
        textInt: [
          { required: true, message: '请输入文字介绍', trigger: 'blur' }
        ]
      },
      canSave: false
    }
  },

  mounted() {
    console.log('this.data', this.data)
  },
  watch: {
    data(newVal) {
      this.data = newVal
      this.echoPic()
    }
  },
  methods: {
    // 图片上传完成后的回调
    uploadOver(e) {
      this.canSave = !e
      this.$emit('uploadOver', e)
    },
    echoPic() {
      const picId = this.data.commodityMaterial.videoCoverageAttachmentId
      if (!picId) return
      this.filesArr = [
        {
          name: picId + '.jpg',
          url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${picId}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`,
          response: {
            code: '0',
            msg: 'OK',
            id: picId,
            url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${picId}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
          },
          fileId: picId
        }
      ]
      console.log(this.filesArr)
    },
    uploadSuccess() {
      console.log('heihei', this.$refs.uploadRef.fileList)
      this.$emit('getFileIds', this.$refs.uploadRef.fileList, 'goods')
    },
    clickWrite() {
      this.isWrite = true
      this.$emit('update:inputDisableIntroduce', false)
    },
    clickCancel() {
      this.isWrite = false
      this.$emit('update:inputDisableIntroduce', true)
    },
    async clickSave() {
      this.$refs['formData'].validate(async valid => {
        if (valid) {
          try {
            let filesArr = this.$refs.uploadRef.fileList
            if (filesArr.length === 0) {
              this.data.commodityMaterial.videoCoverageAttachmentId = ''
            } else {
              this.data.commodityMaterial.videoCoverageAttachmentId =
                filesArr[0].fileId
            }
            const { code } = await updateCommodity(this.data)
            if (code === 0 || code === '0') {
              this.$notify({
                title: '成功',
                message: '修改成功！',
                type: 'success',
                duration: 2000
              })
              this.isWrite = false
              this.$emit('update:inputDisableIntroduce', true)
            }
          } catch (error) {
            console.log('编辑保存失败', error)
          }
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
.goods-introduce {
  font-size: 13px;
  padding-bottom: 40px;
  ::v-deep .el-form-item--medium .el-form-item__content {
    font-size: 13px;
  }
}
.remark {
  font-size: 12px;
  margin: 6px 0 0 10px;
  color: #606266;
}
</style>
