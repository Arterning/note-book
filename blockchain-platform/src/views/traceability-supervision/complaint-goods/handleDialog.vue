<!--
 * @Description  : 处理弹框
 * @Autor        : qinjj
 * @Date         : 2022-07-06 13:54:26
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-09 17:33:13
 * @FilePath     : \blockchain-platform\src\views\traceability-supervision\complaint-goods\handleDialog.vue
-->
<template>
  <el-dialog
    title="处理投诉商品"
    :visible="visible"
    width="35%"
    :show-close="false"
  >
    <el-form :model="rowData" :rules="rules" label-width="150px" ref="form">
      <el-form-item label="投诉人姓名" prop="complainant">
        <el-input
          v-model="rowData.complainant"
          style="width: 83%"
          disabled
        ></el-input>
      </el-form-item>
      <el-form-item label="联系电话" prop="complainantsHotline">
        <el-input
          v-model="rowData.complainantsHotline"
          style="width: 83%"
          disabled
        ></el-input>
      </el-form-item>
      <el-form-item label="联系邮箱" prop="mailbox">
        <el-input
          v-model="rowData.mailbox"
          style="width: 83%"
          disabled
        ></el-input>
      </el-form-item>
      <el-form-item label="投诉商品" prop="commodityName">
        <el-input
          v-model="rowData.commodityName"
          style="width: 83%"
          disabled
        ></el-input>
      </el-form-item>
      <el-form-item label="溯源编号" prop="tracingCode">
        <el-input
          v-model="rowData.tracingCode"
          style="width: 83%"
          disabled
        ></el-input>
      </el-form-item>
      <el-form-item label="投诉理由" prop="complaintsReasons">
        <el-input
          type="textarea"
          v-model="rowData.complaintsReasons"
          style="width: 83%"
          disabled
          :rows="5"
        ></el-input>
      </el-form-item>
      <el-form-item label="处理结果" prop="handleResult">
        <el-input
          type="textarea"
          v-model="rowData.handleResult"
          style="width: 83%"
          clearable
          :rows="5"
        ></el-input>
      </el-form-item>
    </el-form>
    <el-row slot="footer">
      <el-button @click="onClose" size="small">关 闭</el-button>
      <el-button @click="uploadBtn" size="small" type="primary"
        >提 交</el-button
      >
    </el-row>
  </el-dialog>
</template>

<script>
import { submitInfo } from '@/api/complaint-goods.js'
export default {
  name: 'handleDialog',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    rowData: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      rules: {
        handleResult: [
          { required: true, message: '请输入处理结果', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    onClose() {
      this.$emit('update:visible', false)
    },
    uploadBtn() {
      let that = this
      this.$refs.form.validate(valid => {
        if (valid) {
          let params = {
            complaintId: this.rowData.id,
            handleResult: this.rowData.handleResult
          }
          submitInfo(params).then(res => {
            if (res.code === '0') {
              this.$notify({
                title: '成功',
                message: '操作成功！',
                type: 'success',
                duration: 1000
              })
              that.$emit('onRefresh')
            }
          })
        } else {
          return false
        }
      })
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
