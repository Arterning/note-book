<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-06 13:54:26
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-26 16:40:14
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\report-upload\components\UploadReport.vue
-->
<template>
  <el-dialog
    title="上传报告"
    :visible="visible"
    width="35%"
    :show-close="false"
  >
    <el-form :model="form" :rules="rules" ref="form">
      <el-form-item label="检测日期" prop="qcInspectionDate">
        <el-date-picker
          v-model="form.qcInspectionDate"
          align="right"
          type="date"
          value-format="yyyy-MM-dd"
          placeholder="选择日期"
          style="width: 83%"
        >
        </el-date-picker>
      </el-form-item>
      <el-form-item label="检测机构" prop="qcOrg">
        <el-input
          v-model="form.qcOrg"
          placeholder="请填写检测机构"
          style="width: 83%"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item label="质检报告">
        <div>
          <i-upload
            :uploadUrl="uploadUrl"
            :needSave="needSaveFlag"
            :limit="limit"
            :limitSize="limitSize"
            accept=".pdf"
            tip="类型：pdf，不超过2MB，仅1个文件"
            ref="upload"
          ></i-upload>
        </div>
      </el-form-item>
    </el-form>
    <el-row slot="footer">
      <el-button @click="onClose" size="small">关 闭</el-button>
      <el-button
        @click="uploadBtn"
        size="small"
        type="primary"
        :loading="uploadBtnLoading"
        >上 传</el-button
      >
    </el-row>
  </el-dialog>
</template>

<script>
import { reportUploada } from '@/api/basedata-manage'
export default {
  name: 'CheckReport',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    rowData: {
      type: Object
    }
  },
  data() {
    return {
      uploadBtnLoading: false,
      limit: 1,
      limitSize: 2,
      params: { fileClass: 'enterpriseImages' },
      uploadUrl: '',
      needSaveFlag: false,
      fileList: [],
      form: {
        qcInspectionDate: '', //检测时间
        qcOrg: '' //检测机构
      },
      rules: {
        qcInspectionDate: [
          { required: true, message: '请选择日期', trigger: 'change' }
        ],
        qcOrg: [{ required: true, message: '请选择检测机构', trigger: 'blur' }]
      },
      picUrl: '' //图片上传的路径
    }
  },
  methods: {
    onClose() {
      this.$emit('update:visible', false)
    },
    async uploadBtn() {
      console.log('kankan', this.$refs.upload)
      await this.$refs.form.validate()
      if (this.$refs.upload.fileList.length === 0) {
        this.$message.warning('请先选择质检报告')
        return
      }
      try {
        const params = {
          farmId: this.rowData.farmId,
          varietyId: this.rowData.varietyId,
          plantYear: this.rowData.plantYear,
          riceFactoryId: this.rowData.riceFactoryId,
          qcDate: this.form.qcInspectionDate,
          qcOrg: this.form.qcOrg,
          file: this.$refs.upload.fileList[0].raw,
          fileClass: 'xxx',
          varietyName: this.rowData.varietyName,
          varietyType: this.rowData.type,
          farmName: this.rowData.farmName
          // qcImages:
        }
        this.uploadBtnLoading = true
        const { code } = await reportUploada(params)
        this.uploadBtnLoading = false
        if (code === '0' || code === 0) {
          this.$notify({
            message: '上传成功',
            type: 'success',
            duration: 2000,
            onClose: () => {
              this.$emit('update:visible', false)
              this.$emit('refresh')
            }
          })
        }
        console.log('params', params)
      } catch (error) {
        console.log('上传失败 err', error)
      }
    },
    uploadFileProcess(event, file, fileList) {
      console.log(11111)
      console.log(event)
      console.log(file)
      console.log(fileList)
    },
    upSuccess(response, file, fileList) {
      console.log(11111, response, file, fileList)
    },
    handleExceed() {
      this.$message.warning(`当前限制选择 1 个文件`)
    },
    beforeUpload() {
      console.log(12)
    }
  }
}
</script>

<style lang="scss" scoped>
::v-deep .el-dialog__footer {
  text-align: center;
  border-top: 1px solid #ccc;
  height: 55px;
}

::v-deep .el-dialog__header {
  border-bottom: 1px solid #ccc;
}
</style>
