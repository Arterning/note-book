<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-01 14:13:03
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-14 17:05:57
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\company-info\index.vue
-->
<template>
  <div class="company-info">
    <i-tab>
      <div class="container">
        <i-section-header>企业介绍</i-section-header>
        <div class="hr-style"></div>
        <div class="introduce-area">
          <el-row>
            <el-col :span="2">企业名称:</el-col>
            <el-col :span="9" class="text">{{ data.enterpriseName }}</el-col>
            <el-col :span="2">企业地址:</el-col>
            <el-col :span="9" class="text">{{ data.enterpriseAddress }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="2">企业介绍:</el-col>
            <el-col :span="22">
              <el-input
                type="textarea"
                class="text"
                placeholder="请输入企业介绍，如：xxx公司是......"
                v-model="data.enterpriseDetails"
                maxlength="500"
                show-word-limit
                rows="3"
                :disabled="isEdit"
              ></el-input>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="2">相关图片</el-col>
            <el-col :span="22" class="text"
              >用于溯源页面展示，建议上传农产品相关资质，可上传jpg、png图片，单个附件不超过1M，最多可上传15个</el-col
            >
          </el-row>
          <el-row>
            <i-upload
              accept=".jpg,.png"
              uploadUrl="/blockchaincomponent/file/uploadFile/w/v1"
              :params="params"
              :limit="limit"
              :limitSize="limitSize"
              :files="fileArr"
              :disabled="isEdit"
              @uploadOver="
                e => {
                  this.loading = !e
                }
              "
            ></i-upload>
          </el-row>
        </div>
        <i-section-header>
          <span>悬浮资源位配置</span>
          <span style="margin-left: 18px" class="text"
            >用于溯源页面展示，标题请限制在4-6个字之间，如购买渠道</span
          >
        </i-section-header>
        <div class="hr-style"></div>
        <div class="location-area">
          <el-row>
            <el-col :span="1" style="line-height: 32px">品类</el-col>
            <el-col :span="9">
              <el-select
                v-model="data.varietyType"
                placeholder="请选择品类"
                filterable
                clearable
                style="width: 60%"
                size="mini"
                :disabled="isEdit"
              >
                <el-option
                  v-for="item in typeArr"
                  :label="item.dicValue"
                  :value="item.dicCode"
                  :key="item.id"
                  class="option-style"
                ></el-option>
              </el-select>
            </el-col>

            <el-col :span="1" style="line-height: 32px">标题</el-col>
            <el-col :span="13">
              <el-input
                size="mini"
                v-model="data.enterAttachTitle"
                minlength="4"
                maxlength="6"
                :disabled="isEdit"
              ></el-input>
            </el-col>
          </el-row>
          <el-row style="margin-top: 10px">
            <el-col :span="1">链接</el-col>
            <el-col :span="9">
              <el-input
                type="textarea"
                style="width: 60%; background: #fff"
                v-model="data.enterUrl"
                rows="3"
                :disabled="isEdit"
              ></el-input>
            </el-col>
            <el-col :span="1">预览</el-col>

            <el-col :span="13">
              <div
                style="
                  width: 80px;
                  height: 80px;
                  position: relative;
                  cursor: pointer;
                "
                @click="handlePrivew"
              >
                <img
                  style="width: 80px; height: 80px"
                  src="@/assets/images/title.png"
                />
                <span
                  style="
                    position: absolute;
                    color: #fff;
                    font-size: 10px;
                    left: 0;
                    right: 0;
                    bottom: 15px;
                    margin: auto;
                    width: 75px;
                    text-align: center;
                  "
                  >{{
                    data.enterAttachTitle ? data.enterAttachTitle : ''
                  }}</span
                >
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
      <i-footer>
        <el-button
          size="mini"
          @click="e => (isEdit = false)"
          v-btn-permission:1703
          v-if="isEdit"
          >修 改</el-button
        >
        <el-button
          type="primary"
          size="mini"
          @click="onSubmit"
          :loading="loading"
          v-btn-permission:1703
          v-else
          >提 交</el-button
        >
      </i-footer>
    </i-tab>
    <el-drawer
      v-if="dialogVisible"
      :visible.sync="dialogVisible"
      direction="rtl"
      :show-close="false"
      size="375px"
    >
      <template slot="title">
        <p class="title-detail">预览</p>
      </template>
      <iframe
        frameborder="0"
        border="0"
        :src="iframeUrl"
        style="width: 100%; height: 100%"
        id="H5Page"
      ></iframe>
      <div class="dialogFooter">
        <el-button plain size="mini" @click="dialogVisible = false"
          >关闭</el-button
        >
      </div>
    </el-drawer>
  </div>
</template>

<script>
import {
  getCompanyInfo,
  updateCompanyInfo,
  getType
} from '@/api/basedata-manage'

export default {
  name: 'company-info',
  data() {
    return {
      data: {},
      loading: false,
      limit: 15,
      limitSize: 1,
      params: { fileClass: 'enterpriseImages' },
      fileArr: [],
      typeArr: [],
      dialogVisible: false,
      iframeUrl: '',
      isEdit: true
    }
  },
  created() {
    this.getCompanyInfo()
    this.getType()
  },
  methods: {
    // 点击图标预览
    handlePrivew() {
      if (!this.data.enterUrl) {
        this.$message.warning('请先填写链接!')
        return
      }
      this.iframeUrl = this.data.enterUrl
      this.dialogVisible = true
    },
    async getCompanyInfo() {
      this.fileArr = []
      const { data } = await getCompanyInfo()
      var picIdArr = []
      if (data.enterAttachmentId) {
        picIdArr = data.enterAttachmentId.split(',')
      }
      picIdArr.forEach(item => {
        if (item) {
          this.fileArr.push({
            name: item + '.jpg',
            url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${item}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`,
            response: {
              code: '0',
              msg: 'OK',
              id: item,
              url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${item}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
            },
            fileId: item
          })
        }
      })
      console.log('this.fileArr', this.fileArr)

      this.data = data
    },
    async getType() {
      const { data } = await getType({ dicType: 'brand_commodity_type' })
      this.typeArr = data
    },

    async onSubmit() {
      this.loading = true

      var fileidArr = []
      this.fileArr.forEach(item => {
        fileidArr.push(item.fileId)
      })
      this.data.enterAttachmentId = fileidArr.toString()
      console.log('data', this.data)

      try {
        const { code } = await updateCompanyInfo(this.data)
        if (code === 0 || code === '0') {
          this.$notify({
            title: '成功',
            message: '修改成功！',
            type: 'success',
            duration: 2000
          })
          this.isEdit = true
          this.loading = false
          this.getCompanyInfo()
        }
      } catch (error) {
        this.loading = false
        console.log('修改企业信息err', error)
      }
    }
  }
}
</script>

<style scoped lang="scss">
::v-deep .el-drawer {
  border-radius: 10px;
}
::v-deep .el-drawer__wrapper {
  height: 667px;
  margin-top: 100px;
  border-radius: 6px;
}
::v-deep .el-drawer__header {
  border-bottom: 1px solid #ccc;
  padding: 10px;
  display: block;
  height: 10% !important;
  z-index: 999;

  .title-detail {
    font-size: 16px;
    color: #000 !important;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif !important;
  }
}
::v-deep .el-drawer__body {
  padding: 0 0 47px 0;
  margin-top: -31px;
  overflow: hidden;
}

::v-deep .dialogFooter {
  height: 7%;
  width: 100%;
  position: absolute;
  bottom: 0;
  left: 0;
  text-align: center;
  line-height: 42px;
  border-top: 1px solid #ccc;
}

.company-info {
  height: 100%;
  background-color: #fff;
  position: relative;
  font-size: 13px;

  .text {
    font-size: 12px;
    color: #393a35;
  }
  .container {
    width: 75%;
    padding-bottom: 56px;

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
    }

    .introduce-area {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      padding: 10px 0;
      // height: 220px;
    }
    .location-area {
      padding: 10px 0;
      height: 120px;
      display: flex;
      flex-direction: column;
      justify-content: space-between;

      .option-style {
        height: 20px;
        width: 10px;
      }
    }
  }
}

::v-deep .i-section-header .left-side .title {
  font-size: 14px;
}
::v-deep .el-textarea__inner {
  border-radius: 0;
  // background-color: #fafafa;
}

::v-deep .i-footer .right {
  text-align: center;
}
</style>

<style scoped>
/* 修改i-upload样式  */
::v-deep .el-upload--picture-card {
  border: 1px solid #c0ccda;
}
::v-deep .el-row {
  margin-top: 20px;
}
</style>
