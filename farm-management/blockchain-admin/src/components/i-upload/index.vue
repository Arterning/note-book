<!--
 * @Description  : IUpload
 * @Autor        : Hemingway
 * @Date         : 2022-04-19 15:13:17
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-21 09:03:25
 * @FilePath     : \blockchain-admin\src\components\i-upload\index.vue
-->
<template>
  <div class="i-upload">
    <el-upload
      :auto-upload="needSave"
      :action="action"
      :data="data"
      :name="name"
      :list-type="needSave ? 'picture-card' : 'text'"
      :limit="limit"
      :multiple="multiple"
      :accept="accept"
      :on-exceed="handleExceed"
      :before-upload="beforeUpload"
      :on-success="handleSuccess"
      :on-error="handleError"
      :on-remove="handleRemove"
      :on-change="handleChange"
      :file-list="fileList"
      :disabled="disabled"
      ref="upload"
      :class="{ finished: finished }"
    >
      <!-- 自动上传 -->
      <template v-if="needSave">
        <i class="el-icon-plus"></i>
      </template>
      <!-- 手动上传 -->
      <template v-else>
        <el-button size="small" type="primary" :loading="loading"
          >选取文件</el-button
        >
        <div class="el-upload__tip" slot="tip" v-if="tip">
          <i class="el-icon-info"></i> {{ tip }}
        </div>
      </template>

      <!-- 自动上传 -->
      <template #file="{ file }" v-if="needSave">
        <el-tooltip :content="file.name" placement="top">
          <!-- <div v-loading="true">上传中</div> -->
          <div
            style="width: 100%; height: 100%"
            v-loading="uploadLoading && beforeFile.uid === file.uid"
            element-loading-text="上传中..."
          >
            <img
              class="el-upload-list__item-thumbnail"
              :src="thumbnailUrl(file)"
              alt=""
            />
            <span class="el-upload-list__item-actions">
              <span
                class="el-upload-list__item-preview"
                @click="handlePreview(file)"
              >
                <i class="el-icon-zoom-in"></i>
              </span>
              <span
                class="el-upload-list__item-delete"
                @click="handleDownload(file)"
              >
                <i class="el-icon-download"></i>
              </span>
              <span
                class="el-upload-list__item-delete"
                v-if="!disabled"
                @click="handleRemove(file)"
              >
                <i class="el-icon-delete"></i>
              </span>
            </span>
          </div>
        </el-tooltip>
      </template>
    </el-upload>

    <el-dialog :visible.sync="previewVisible" v-if="!customPreview">
      <pdf :src="previewUrl" v-if="previewPDF" />
      <img width="100%" :src="previewUrl" alt="" v-else />
    </el-dialog>
  </div>
</template>

<script>
import pdf from 'vue-pdf'

// 默认校验函数
const defaultValidation = (file, resolve, reject) => { // eslint-disable-line
  resolve()
}

export default {
  name: 'IUpload',

  components: { pdf },

  props: {
    // 上传文件是否需要保存到文件服务器
    needSave: {
      type: Boolean,
      default: true
    },
    // 文件上传url
    uploadUrl: {
      type: String,
      default: '/agriculturecomponent/file/uploadFile/w/v1'
    },
    // 接口请求额外参数 （1）自动上传：fileClass、isOpen（2）手动上传：业务接口所需参数
    params: {
      type: Object,
      default: () => {}
    },
    // 上传的文件字段名
    name: {
      type: String,
      default: 'file'
    },
    // 是否支持文件多选
    multiple: {
      type: Boolean,
      default: false
    },
    // 文件最大数量
    limit: {
      type: Number,
      default: 1
    },
    // 单个文件最大尺寸，单位 M
    limitSize: {
      type: Number,
      default: 10
    },
    // 文件类型，如：".jpg,.png,.doc"
    accept: {
      type: String,
      default: ''
    },
    // 上传提示
    tip: {
      type: String,
      default: ''
    },
    // 自定义校验函数（需自己处理异常信息的展示，函数结束前必须调用 resolve 或 reject）
    validation: {
      type: Function,
      default: defaultValidation
    },
    // 是否需要自定义预览
    customPreview: {
      type: Boolean,
      default: false
    },
    // 默认文件列表（文件回显场景）
    files: {
      type: Array,
      default: () => []
    },
    // 是否禁用
    disabled: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {
      fileList: [], // 文件列表
      invalidList: [], // 失效文件列表
      loading: false, // 选取按钮加载效果
      hasValidation: false, // 是否存在自定义校验

      // 预览
      previewPDF: '',
      previewUrl: '',
      previewVisible: false,

      successCount: 0, // 成功计数
      errNameList: [], // 自动上传：错误文件name列表
      uploadLoading: false,
      beforeFile: {}
    }
  },

  computed: {
    // 上传url
    action() {
      const uploadUrl = (process.env.VUE_APP_URL_PREFIX || '') + this.uploadUrl

      return process.env.VUE_APP_BASE_URL + uploadUrl
    },
    // formData参数
    data() {
      const defaultParams = {
        sessionId: this.$store.getters.token,
        clientId: 'poweb'
      }
      return { ...defaultParams, ...this.params }
    },
    // 是否已达到最大上传数量
    finished() {
      return this.fileList.length >= this.limit
    },

    // 无效图片ids字符串拼接
    invalid() {
      return this.invalidList.map(file => file.fileId || file.id).join(',')
    }
  },

  watch: {
    // 初始化fileList
    files: {
      handler(newVal) {
        console.log(newVal)
        this.fileList = newVal
        newVal && newVal.length > 0 && (this.fileList = newVal)
      },
      immediate: true
    },
    // 判断是否存在自定义校验函数
    validation: {
      handler(newVal) {
        this.hasValidation = newVal !== defaultValidation
      },
      immediate: true
    },
    fileList: {
      handler(newVal) {
        // 确保准确捕获到初始化状态（图片回显场景必须！）
        this.$nextTick(() => {
          this.$emit(
            'change',
            this.needSave
              ? {
                  // 待上传文件列表
                  fileList: newVal,
                  // 有效图片ids字符串拼接
                  effective: newVal
                    .filter(file => {
                      return file && file.status === 'success'
                    })
                    .map(file => (file.response || file).id)
                    .join(','),
                  // 无效图片ids字符串拼接
                  invalid: this.invalid
                }
              : { fileList: newVal }
          )
        })
      }
    }
  },

  methods: {
    /**
     * @description: 文件选择超出事件
     * @param {Object} files 本次选择的文件
     * @param {Array} fileList 所有选择的文件
     * @author: Hemingway
     */
    handleExceed(files, fileList) { // eslint-disable-line
      this.$message.warning(`当前限制选择 ${this.limit} 个文件`)
    },

    /**
     * @description: 文件上传前逻辑（插入自定义校验）
     * @param {Object} file 本次选择的文件
     * @author: Hemingway
     */
    beforeUpload(file) {
      // 自动上传常规性检查，异常后续处理
      if (this.needSave && this.errNameList.includes(file.name)) {
        return false // 阻止继续上传，将触发 on-remove
      }

      // 自定义校验
      if (this.hasValidation) {
        this.loading = true
        return new Promise(async (resolve, reject) => { // eslint-disable-line
          try {
            await this.validation(file, resolve, reject)
            this.loading = false
          } catch (error) {
            console.log('上传文件自定义校验 error: ', error)

            this.loading = false
          }
        })
      }
      // 上传图片展示加载图标
      this.beforeFile = file
      this.uploadLoading = true
      this.$emit('uploadOver', false)
    },

    /**
     * @description: 文件改变事件
     * @param {Object} file 本次改变的文件
     * @param {Array} fileList
     * @author: Hemingway
     */
  async  handleChange(file, fileList) { // eslint-disable-line
      // 文件添加时机，执行常规性校验
      if (file.status === 'ready') {
        const res = this.handleCommonCheck(file)

        if (res === undefined) {
          // 正常

          this.fileList.push(file)
        } else {
          // 异常

          this.$message.warning(res.message)

          // 自动上传：标记错误文件，在 beforeUpload 阻止
          // 手动上传：手动维护列表
          this.needSave
            ? this.errNameList.push(file.name)
            : (this.fileList = [...this.fileList])
        }
      }
    },

    /**
     * @description: 文件常规性检查
     * @param {Object} file
     * @return {Promise}
     * @author: Hemingway
     */
    handleCommonCheck(file) {
      // 检查文件类型

      const suffixName = file.name
        .slice(file.name.lastIndexOf('.'))
        .toLocaleLowerCase()
      if (!this.accept.split(',').includes(suffixName)) {
        return new Error(`${file.name}：不支持的文件类型`)
      }

      // 检查文件名
      if (this.fileList.map(file => file.name).includes(file.name)) {
        return new Error(`${file.name}：文件已在列表中存在`)
      }

      // 检查文件大小
      if (file.size > this.limitSize * 1024 * 1024) {
        return new Error(
          `${file.name}：文件大小超出最大限制 ${this.limitSize} MB`
        )
      }
    },

    /**
     * @description: 文件上传成功事件
     * @param {Object} response 本次上传的响应
     * @param {Object} file 本次上传的文件
     * @param {Array} fileList 已上传 + 本次上传
     * @author: Hemingway
     */
    handleSuccess(response, file, fileList) {
      this.uploadLoading = false
      this.$emit('uploadOver', true)
      if (response.code !== '0') {
        let subErrors = response.subErrors
        const message =
          (subErrors && subErrors.length > 0 && subErrors[0].message) ||
          response.msg
        this.$message({
          message: `${file.name}：${message}`,
          type: 'error',
          duration: 3000
        })

        // 删除出错的文件
        this.fileList.splice(
          this.fileList.findIndex(item => item.name === file.name),
          1
        )

        // 上传失败判断
        if (
          fileList.filter(file => !['success', 'fail'].includes(file.status))
            .length === 0
        ) {
          this.$emit('error')
          this.resetVariable()
        }
      } else {
        // 自动上传：添加fileId属性（额外的动机：修改fileList）
        if (this.needSave) {
          this.fileList.splice(
            this.fileList.findIndex(item => item.name === file.name),
            1,
            { ...file, fileId: response.id }
          )
        }

        // 上传成功判断
        if (++this.successCount === fileList.length) {
          this.$emit('success')

          this.resetVariable()
        }
      }
    },

    /**
     * @description: 文件上传失败事件
     * @param {} err
     * @param {Object} file 本次选择的文件
     * @param {Array} fileList 已上传 + 本次上传
     * @author: Hemingway
     */
    handleError(err, file, fileList) {
      this.uploadLoading = false
      this.$emit('uploadOver', true)
      this.$message({
        message: `${file.name}：${'网络异常'}`, // 暂不了解 err 的数据结构
        type: 'error',
        duration: 3000
      })

      // 删除出错的文件
      this.fileList.splice(
        this.fileList.findIndex(item => item.name === file.name),
        1
      )

      // 上传失败判断
      if (
        fileList.filter(file => !['success', 'fail'].includes(file.status))
          .length === 0
      ) {
        this.$emit('error')
        this.resetVariable()
      }
    },

    /**
     * @description: 一次导入之后，需重置变量
     * @author: Hemingway
     */
    resetVariable() {
      this.successCount = 0
      this.errNameList = []
    },

    /**
     * @description: 文件预览事件
     * @param {Object} file
     * @author: Hemingway
     */
    handlePreview(file) {
      console.log('yulan', file)
      const suffixName = file.name
        .slice(file.name.lastIndexOf('.') + 1)
        .toLocaleLowerCase()

      if (!['jpg', 'png', 'jpeg', 'gif', 'pdf'].includes(suffixName)) {
        this.$message(`.${suffixName} 格式文件暂不支持预览`)
        return false
      }

      if (this.customPreview) {
        this.$emit('preview', { previewUrl: file.url })
      } else {
        this.previewPDF = suffixName === 'pdf'
        this.previewUrl = file.url
        this.previewVisible = true
      }
    },

    /**
     * @description: 缩略图url
     * @param {Object} file
     * @return {String}
     * @author: Hemingway
     */
    thumbnailUrl(file) {
      const suffixName = file.name
        .slice(file.name.lastIndexOf('.') + 1)
        .toLocaleLowerCase()

      if (['jpg', 'png', 'jpeg', 'gif'].includes(suffixName)) {
        return file.url
      } else if (['pdf', 'doc', 'xls', 'ppt'].includes(suffixName)) {
        return require(`@/assets/images/fileThumbnail/${suffixName}.png`)
      } else if ('xlsx' === suffixName) {
        return require(`@/assets/images/fileThumbnail/xls.png`)
      } else {
        return ''
      }
    },

    /**
     * @description: 文件下载事件
     * @param {Object} file
     * @author: Hemingway
     */
    handleDownload(file) {
      const win = window.open(file.url, file.name, '')

      const timer = setInterval(() => {
        if (win.document.title === file.name) {
          clearInterval(timer)
        } else {
          win.document.title = file.name
        }
      }, 1 / 48)
    },

    /**
     * @description: 文件移除事件
     * @param {Object} file 本次移除的文件
     * @param {Array} fileList 移除后的文件列表
     * @author: Hemingway
     */
    handleRemove(file, fileList) {
      // 自动上传：移除该文件对应的 errName
      if (file.status === 'ready') {
        this.fileList = fileList

        if (this.needSave) {
          // 自动上传
          this.errNameList.splice(
            this.errNameList.findIndex(name => name === file.name),
            1
          ) // 更新errName
        }
      }

      // 手动删除已上传文件，交互
      if (file.status === 'success') {
        this.$confirm(
          `此操作将永久删除文件：${file.name}，是否继续？`,
          '提示',
          {
            confirmButtonText: '确 定',
            cancelButtonText: '取 消',
            type: 'warning'
          }
        )
          .then(() => {
            this.fileList.splice(
              this.fileList.findIndex(item => item.name === file.name),
              1
            )

            // 自动上传：统计文件状态
            if (this.needSave) {
              this.invalidList.push(file)
            }
          })
          .catch(() => {})
      }
    },

    /**
     * @description: 手动上传，提交到业务接口
     * @author: Hemingway
     */
    submit() {
      this.$refs.upload.submit()
    }
  }
}
</script>

<style lang="scss" scoped>
.finished ::v-deep {
  .el-upload.el-upload--picture-card {
    display: none;
  }
}
::v-deep
  .el-upload-list.el-upload-list--picture-card.is-disabled
  + .el-upload.el-upload--picture-card {
  background-color: #f5f7fa;
  border-color: #dfe4ed;
  color: #c0c4cc;
  cursor: not-allowed;
}
</style>
