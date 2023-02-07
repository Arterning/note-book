<!--
 * @Description  : 
 * @Autor        : Tyne
 * @Date         : 2021-07-21 17:13:53
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-06 10:26:30
 * @FilePath     : \blockchain-platform\src\views\blockchain\components\settle-audit-dialog.vue
-->
<template>
  <div class="settle-audit-dialog">
    <i-dialog :visible.sync="dialogWindowVisible" class="detail-dialog">
      <template #title>
        入驻审核
        <el-tag
          :type="tagTypeWithStatus(auditData.checkStatus)"
          style="margin-left: 16px"
          v-if="auditData.checkStatus !== '0'"
          >{{ tagNameWithStatus(auditData.checkStatus) }}</el-tag
        >
      </template>

      <i-section-header>申请信息</i-section-header>
      <el-row :gutter="20">
        <el-col :span="8"
          ><div class="grid-content bg-purple">
            <span>企业类型: </span>
            <span style="color: #666">{{
              companyTypeWithType(auditData.type)
            }}</span>
          </div></el-col
        >
        <el-col :span="8"
          ><div class="grid-content bg-purple">
            <span>企业名称: </span>
            <span style="color: #666">{{ auditData.name }}</span>
          </div></el-col
        >
        <el-col :span="8"
          ><div class="grid-content bg-purple">
            <span>统一信用代码: </span>
            <span style="color: #666">{{ auditData.uuid }}</span>
          </div></el-col
        >
      </el-row>
      <el-row :gutter="20" style="margin-top: 15px">
        <el-col :span="8"
          ><div class="grid-content bg-purple">
            <span>法人代表: </span>
            <span style="color: #666">{{ auditData.legalPerson }}</span>
          </div></el-col
        >
        <el-col :span="8"
          ><div class="grid-content bg-purple">
            <span>企业联系电话: </span>
            <span style="color: #666">{{ auditData.officePhone }}</span>
          </div></el-col
        >
        <el-col :span="8"
          ><div class="grid-content bg-purple">
            <span>企业联系地址: </span>
            <span style="color: #666">{{ auditData.address }}</span>
          </div></el-col
        >
      </el-row>

      <i-section-header>企业资质</i-section-header>
      <div :key="index" v-for="(item, index) in imageList">
        <AuditImagePreview :group="item" />
      </div>

      <div v-if="auditData.checkStatus !== '0' && auditData.checkMemo">
        <i-section-header>审核意见</i-section-header>
        <div>
          {{ auditData.checkMemo }}
        </div>
      </div>

      <div slot="footer" class="dialog-footer" v-if="isHandle">
        <el-button @click="handleAudit('3')">不通过</el-button>
        <el-button type="primary" @click="handleAudit('1')">通 过</el-button>
      </div>
    </i-dialog>

    <i-dialog :visible.sync="handleAuditVisible" :title="auditTitle">
      <el-form ref="auditForm" :model="auditForm">
        <el-form-item
          :rules="{
            required: true,
            message: '请填写审核内容'
          }"
          prop="auditContent"
          label=""
        >
          <el-input
            :maxlength="100"
            show-word-limit
            type="textarea"
            :autosize="{ minRows: 3 }"
            :placeholder="placeholder"
            v-model="auditForm.auditContent"
          ></el-input>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="() => (handleAuditVisible = false)">取 消</el-button>
        <el-button @click="confirmAudit" type="primary">确 定</el-button>
      </div>
    </i-dialog>
  </div>
</template>

<script>
import AuditImagePreview from '../components/audit-image-preview'

export default {
  name: 'SettleAuditDialog',
  components: { AuditImagePreview },
  props: {
    auditData: {
      type: Object,
      default() {
        return {}
      }
    },
    auditType: {
      type: String,
      default: '0'
    },
    mapInfo: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  data() {
    return {
      newRoleData: {},
      form: {},
      dialogWindowVisible: false,
      handleAuditVisible: false,
      auditForm: {
        auditContent: ''
      },
      checkStatus: '',
      isHandle: false,
      auditTitle: '',
      placeholder: ''
    }
  },
  computed: {
    imageList() {
      const orignalImageMap = this.auditData.attachments
      if (!orignalImageMap) return []
      const idCardFront = orignalImageMap.idCardFront
        ? orignalImageMap.idCardFront
        : []
      const idCardReverse = orignalImageMap.idCardReverse
        ? orignalImageMap.idCardReverse
        : []
      const businessLicense = orignalImageMap.businessLicense
        ? orignalImageMap.businessLicense
        : []
      const registrationCertificate = orignalImageMap.registrationCertificate
        ? orignalImageMap.registrationCertificate
        : []
      const riceProcessingLicense = orignalImageMap.riceProcessingLicense
        ? orignalImageMap.riceProcessingLicense
        : []
      const foodBusinessLicense = orignalImageMap.foodBusinessLicense
        ? orignalImageMap.foodBusinessLicense
        : []
      const dryEnvironment = orignalImageMap.dryEnvironment
        ? orignalImageMap.dryEnvironment
        : []
      const storageEnvironment = orignalImageMap.storageEnvironment
        ? orignalImageMap.storageEnvironment
        : []
      const processingEnvironment = orignalImageMap.processingEnvironment
        ? orignalImageMap.processingEnvironment
        : []
      const packageEnvironment = orignalImageMap.packageEnvironment
        ? orignalImageMap.packageEnvironment
        : []

      const imageList = [
        {
          title: '法人代表身份证',
          imageList: [
            ...idCardFront.map(element => ({
              imageName: '正面',
              imageUrl: this.getImgUrl(element.id)
            })),
            ...idCardReverse.map(element => ({
              imageName: '反面',
              imageUrl: this.getImgUrl(element.id)
            }))
          ]
        },
        {
          title: '企业营业执照',
          imageList: [
            ...businessLicense.map(element => ({
              imageUrl: this.getImgUrl(element.id)
            })),
            ...registrationCertificate.map(element => ({
              imageUrl: this.getImgUrl(element.id)
            }))
          ]
        },
        {
          title: '大米生产加工许可证',
          imageList: [
            ...riceProcessingLicense.map(element => ({
              imageUrl: this.getImgUrl(element.id)
            })),
            ...foodBusinessLicense.map(element => ({
              imageUrl: this.getImgUrl(element.id)
            }))
          ]
        },
        {
          title: '加工环境',
          imageList: [
            ...dryEnvironment.map(element => ({
              imageName: '烘干',
              imageUrl: this.getImgUrl(element.id)
            })),
            ...storageEnvironment.map(element => ({
              imageName: '仓储',
              imageUrl: this.getImgUrl(element.id)
            })),
            ...processingEnvironment.map(element => ({
              imageName: '加工',
              imageUrl: this.getImgUrl(element.id)
            })),
            ...packageEnvironment.map(element => ({
              imageName: '包装',
              imageUrl: this.getImgUrl(element.id)
            }))
          ]
        }
      ]
      return imageList.filter(element => element.imageList.length > 0)
    }
  },
  watch: {
    checkStatus(val) {
      if (val === '4') {
        this.auditTitle = '资格取消理由'
        this.placeholder = '请输入取消资格理由'
      } else {
        this.auditTitle = '入驻审核意见'
        this.placeholder = '请输入审核意见'
      }
    }
  },
  methods: {
    showDialog(isHandle) {
      this.isHandle = isHandle
      this.dialogWindowVisible = true
    },
    handleAudit(checkStatus) {
      this.handleAuditVisible = true
      this.auditForm.auditContent = ''
      this.checkStatus = checkStatus
      this.$nextTick(() => {
        this.$refs.auditForm.resetFields()
      })
    },
    closeAuditHandle() {
      this.handleAuditVisible = false
      this.dialogWindowVisible = false
    },
    confirmAudit() {
      this.$refs.auditForm.validate(valid => {
        if (valid) {
          this.$emit('confirmAudit', {
            id: this.auditData.id,
            checkMemo: this.auditForm.auditContent,
            checkStatus: this.checkStatus
          })
          this.closeAuditHandle()
        } else {
          return false
        }
      })
    },
    tagTypeWithStatus(status) {
      let type = ''
      switch (status) {
        case '0':
          type = 'info'
          break
        case '1':
          type = 'success'
          break
        case '2':
          type = 'info'
          break
        case '3':
          type = 'warning'
          break
        case '4':
          type = 'danger'
          break
        default:
          break
      }
      return type
    },
    tagNameWithStatus(status) {
      let type = ''
      switch (status) {
        case '0':
          type = '未审核'
          break
        case '1':
          type = '通过'
          break
        case '2':
          type = '审核中'
          break
        case '3':
          type = '未通过'
          break
        case '4':
          type = '已取消资格'
          break
        default:
          break
      }
      return type
    },
    companyTypeWithType(type) {
      try {
        const companyMap = this.mapInfo.enterprise_type
        const nameList = type.split(',').map(i => companyMap[i])
        const typeName = nameList.join(',')
        return typeName
      } catch (error) {
        return ''
      }
    },
    getImgUrl(id) {
      let url = `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&sessionId=${this.$store.getters.sessionId}&clientId=poweb`
      url = url.replace('/nfapi/', '/')
      return url
    }
  }
}
</script>

<style lang="scss" scoped>
.detail-dialog {
  ::v-deep .el-dialog__body {
    padding-top: 0;
  }
}
</style>
