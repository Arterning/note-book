<!--
 * @Description  : 查看报告弹窗
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-06 13:54:26
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-02 13:48:25
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\report-upload\components\CheckReport.vue
-->
<template>
  <el-dialog
    title="查看报告"
    :visible="visible"
    width="35%"
    :show-close="false"
  >
    <el-form :model="rowDataCheck" :rules="rules">
      <el-form-item label="检测日期" prop="qcInspectionDate">
        <el-date-picker
          disabled
          v-model="rowDataCheck.qcInspectionDate"
          align="right"
          type="date"
          placeholder="选择日期"
          style="width: 83%"
        >
        </el-date-picker>
      </el-form-item>
      <el-form-item label="检测机构" prop="qcOrg">
        <el-input
          v-model="rowDataCheck.qcOrg"
          clearable
          placeholder="请填写检测机构"
          disabled
          style="width: 83%"
        ></el-input>
      </el-form-item>
      <el-form-item label="质检报告">
        <el-image
          v-for="(url, index) in filesArr"
          :key="index"
          style="width: 100px; height: 100px"
          :src="url"
          :preview-src-list="filesArr"
        >
        </el-image>
      </el-form-item>
    </el-form>
    <el-row slot="footer">
      <el-button @click="onClose" size="small">关 闭</el-button>
    </el-row>
  </el-dialog>
</template>

<script>
export default {
  name: 'CheckReport',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    tableData: {
      type: Array,
      default: null
    },
    rowDataCheck: {
      Type: Object,
      default: null
    }
  },
  data() {
    return {
      filesArr: [],
      picIdArr: [] //用于图片回显的id
      // url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
    }
  },
  mounted() {
    this.echoPic()
  },
  methods: {
    //回显图片
    echoPic() {
      this.picIdArr = this.rowDataCheck.qcImages.split(',')
      for (var i = 0; i < this.picIdArr.length; i++) {
        this.filesArr.push(
          `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${this.picIdArr[i]}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
        )
      }
    },

    onClose() {
      this.$emit('update:visible', false)
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
