<!--
 * @Description  : 拖拽上传文件
 * @Autor        : Tyne
 * @Date         : 2022-01-23 19:56:31
 * @LastEditors  : Tyne
 * @LastEditTime : 2022-04-18 17:18:50
 * @FilePath     : /i-farm-platform/src/components/zl/i-drag-upload/index.vue
-->

<template>
  <div class="drag-upload-content">
    <el-upload
      :action="actionURL"
      :on-remove="handleRemove"
      :on-exceed="handleExceed"
      :on-error="handleError"
      :on-success="handleSuccess"
      :data="updateParam"
      :before-upload="beforeUpload"
      :limit="limit"
      :file-list="fileList"
      drag
    >
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">
        只能上传jpg/png文件，且不超过20M
      </div>
    </el-upload>
    <ul class="el-upload-list el-upload-list--text"></ul>
  </div>
</template>

<script>
import { getImgUrl } from '@/utils/business'
export default {
  name: 'IDragUpload',
  props: {
    uploadUrl: {
      type: String,
      default: '/agriculturecomponent/file/uploadFile/w/v1'
    },
    fileClass: {
      type: String,
      default: ''
    },
    limit: {
      type: Number,
      default: 6
    },
    orignalFileList: {
      type: Array,
      default() {
        return []
      }
    }
  },
  data() {
    return {
      updateParam: {},
      fileList: [],
      deleteIds: [], // 删除的资源ID数组
      fileIds: [], // 保留的资源ID数组
      isCancel: false, // 是否是取消上传图片
      uploading: false // 是否正在上传文件
    }
  },
  watch: {
    orignalFileList: {
      handler(val) {
        this.fileList = val.map(e => ({
          id: e,
          name: e + '.png',
          url: getImgUrl(e)
        }))
        this.fileIds = [].concat(val)
        this.deleteIds = []
      },
      deep: true
    }
  },
  computed: {
    actionURL() {
      return (
        process.env.VUE_APP_BASE_URL +
        (process.env.VUE_APP_URL_PREFIX || '') +
        this.uploadUrl
      )
    }
  },
  beforeMount() {
    const val = this.orignalFileList
    this.fileList = val.map(e => ({
      id: e,
      name: e + '.png',
      url: getImgUrl(e)
    }))
    this.fileIds = [].concat(val)
    this.deleteIds = []
  },
  mounted() {
    if (this.$store.getters.token) {
      this.updateParam['sessionId'] = this.$store.getters.token
      this.updateParam['fileClass'] = this.fileClass
      this.updateParam['isOpen'] = 'Y'
      this.updateParam['clientId'] = 'poweb'
    }
  },
  methods: {
    handleRemove(file) {
      // 删除文件
      if (!file.id && !file.response) {
        return
      }
      const id = file.id || file.response.id
      this.deleteIds.push(id)
      // 从上传成功数组中，移除被删除的文件ID
      this.fileIds.splice(
        this.fileIds.findIndex(item => item === id),
        1
      )
      this.fileList.splice(
        this.fileList.findIndex(item => (item.id || item.response.id) === id),
        1
      )
      const deleteStr =
        this.deleteIds.length !== 0 ? this.deleteIds.join(',') : ''
      const fileStr = this.fileIds.length !== 0 ? this.fileIds.join(',') : ''
      // 实时返回结果数组（存储了文件ID）
      this.$emit('HandleFile', {
        fileId: fileStr,
        deleteIds: deleteStr,
        fileList: this.fileList
      })
    },

    handleExceed() {
      // 上传限制
      this.uploading = false
      this.$message(`最多上传${this.limit}张`)
    },
    handleError() {
      // 上传失败
      this.uploading = false
      this.$message.error('上传失败')
    },
    handleSuccess(response, file, fileList) {
      console.log('fileList:', fileList)
      console.log('file:', file)
      // 上传成功
      this.uploading = false
      if (this.checkRes(response)) {
        const param = this.filterData(file)
        // 实时返回结果数组（存储了文件ID）
        this.$emit('HandleFile', { ...param, fileList: this.fileList })
      } else {
        // 上传失败重置数组
        this.fileList = [...this.fileList]
      }
    },
    beforeUpload(file) {
      // 上传前 钩子函数 。 上传附件大小不能超过10M
      this.uploading = true
      const extStart = file.name.lastIndexOf('.')
      const ext = file.name.substring(extStart, file.name.length).toUpperCase()
      if (ext !== '.PNG' && ext !== '.JPG' && ext !== '.JPEG') {
        // 判断文件是否是图片类型
        this.$message({
          message: '此类型文件不能上传',
          type: 'warning'
        })
        this.isCancel = true
        this.uploading = false
        return false
      }
      const fileSize = file.size
      if (!fileSize || fileSize === 0) {
        this.$message({
          message: '文件或已损坏，请正确选择文件',
          type: 'warning'
        })
        this.uploading = false
        return false
      }
      if (fileSize > 20 * 1024 * 1024) {
        this.$message({
          message: '上传文件大小不能超过20M',
          type: 'warning'
        })
        this.uploading = false
        return false
      }
    },
    checkRes(res) {
      // 验证返回值
      if (res.code === '20' || res.code === '21') {
        let message
        if (res.code === '20') {
          message = '缺少sessionId，请重新登入'
        } else if (res.code === '21') {
          message = '无效的sessionId，请重新登入'
        }
        this.$message.error(message)
        return false
      } else if (res.code !== '0') {
        let message
        if (
          res.subErrors.length !== 0 &&
          res.subErrors[0].message &&
          res.subErrors[0].message.length !== 0
        ) {
          // 判断错误信息
          message = res.subErrors[0].message
        } else {
          message = res.msg
        }
        this.$message.error(message)
        return false
      }
      return true
    },
    filterData(file) {
      // 过滤数据
      this.fileList.push(file)
      if (file.response.id) {
        this.fileIds.push(file.response.id)
      } else {
        this.$message.error('后台返回文件id为nil')
      }
      // 拼接回传数据
      const deleteStr =
        this.deleteIds.length !== 0 ? this.deleteIds.join(',') : ''
      const fileStr = this.fileIds.length !== 0 ? this.fileIds.join(',') : ''
      return { fileId: fileStr, deleteIds: deleteStr }
    },
    forbidClick() {}, // 只是为了禁止点击冒泡
    close() {
      // this.dialogImageUrl = ''
    }
  }
}
</script>

<style scoped lang="scss">
.drag-upload-content {
  width: 100%;
  padding: 0 105px;
}
</style>
