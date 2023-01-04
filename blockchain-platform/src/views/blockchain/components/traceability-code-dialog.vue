<template>
  <i-dialog :visible.sync="visible" :before-close="handleClose" title="发货">
    <el-form ref="traceabilityCode" :model="formData" :rules="rules">
      <template v-if="type === 0">
        <el-form-item label="数量：" prop="downloadNum" required>
          <el-input
            v-model="formData.downloadNum"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item label="溯源码编号：">
          <div class="sy_code">SY-DM00000001 - SY-DM00010000</div>
        </el-form-item>
      </template>
      <template v-if="type === 1">
        <el-form-item label="快递单号：" prop="expressNum" required>
          <el-input
            v-model="formData.expressNum"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
      </template>
      <template v-if="type === 2">
        <el-form-item label="作废理由：" prop="reason" required>
          <el-input
            v-model="formData.reason"
            :rows="2"
            type="textarea"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item label="退款金额：" prop="refundAmount" required>
          <el-input
            v-model="formData.refundAmount"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
      </template>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="handleConfirm" type="primary">{{
        confirmText
      }}</el-button>
    </span>
  </i-dialog>
</template>

<script>
export default {
  name: 'TraceabilityCodeDialog',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    type: {
      type: Number,
      default: 0
    }
  },

  data() {
    return {
      formData: {
        downloadNum: '',
        expressNum: '',
        reason: '',
        refundAmount: ''
      },
      rules: {
        downloadNum: [
          { required: true, message: '请输入下载数量', trigger: 'blur' },
          {
            validator: (rule, value, callback) => {
              if (Number(value) > 50000) {
                return callback('下载数量不能超过 50000')
              }
              return callback()
            }
          }
        ],
        expressNum: [
          { required: true, message: '请输入快递单号', trigger: 'blur' }
        ],
        reason: [
          { required: true, message: '请输入作废理由', trigger: 'blur' }
        ],
        refundAmount: [
          { required: true, message: '请输入退款金额', trigger: 'blur' }
        ]
      }
    }
  },

  computed: {
    confirmText() {
      const texts = ['下 载', '确认发货', '确 认']
      return texts[this.type]
    }
  },

  methods: {
    handleClose() {
      this.$emit('close')
    },

    handleConfirm() {
      this.$refs.traceabilityCode.validate(valid => {
        if (valid) {
          for (const key in this.formData) {
            if (this.formData[key] === '') delete this.formData[key]
          }
          this.$emit('confirm', this.formData)
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.sy_code {
  color: #606266;
  font-size: 12px;
  height: 32px;
}
</style>
