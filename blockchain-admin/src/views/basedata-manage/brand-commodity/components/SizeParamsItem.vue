<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-11 17:10:39
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-08 15:38:17
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\components\SizeParamsItem.vue
-->
<template>
  <div class="size-params">
    <i-section-header title="规格参数">
      <div slot="action" v-if="showEditBtnParams === true">
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
    <el-form :model="data" ref="formData" class="custom-form">
      <el-row style="margin-bottom: 6px">
        <el-col :span="8"> <span>规格</span> </el-col>
        <el-col :span="10">
          <span>商品图片</span>
          <span class="remark">
            用于溯源，尺寸716*716,建议左上角放品牌logo。</span
          ></el-col
        >
        <el-col :span="6">商品单价</el-col>
      </el-row>
      <el-row v-for="(item, index) in computedUnitList" :key="index">
        <el-col :span="8">
          <el-form-item>
            <!-- 删除按钮 -->
            <el-button
              type="text"
              icon="el-icon-delete"
              class="delete-btn"
              size="mini"
              v-if="index !== 0 && showEditBtnParams === false"
              @click="clickDelete(index)"
            ></el-button>
            <!-- 删除按钮 -->
            <el-input
              v-model="item.weightValue"
              placeholder="请输入"
              style="width: 150px"
              size="mini"
              :disabled="inputDisableParams"
            >
              <template slot="append">kg</template>
            </el-input>
            /
            <el-select
              v-model="item.weightUnit"
              style="width: 61px"
              size="mini"
              :disabled="inputDisableParams"
            >
              <el-option label="袋" value="袋" />
              <el-option label="件" value="件" />
              <el-option label="罐" value="罐" />
              <el-option label="个" value="个" /> </el-select
            ><br />
            <el-button
              type="text"
              style="margin: 14px 0 0 108px"
              size="mini"
              @click="clickAdd"
              v-if="index === 0 && showEditBtnParams === false"
              >新增</el-button
            >
          </el-form-item>
        </el-col>
        <el-col :span="10">
          <el-form-item>
            <i-upload
              @success="uploadSuccess(index)"
              ref="uploadRef"
              accept=".jpg,.png"
              uploadUrl="/blockchaincomponent/file/uploadFile/w/v1"
              :params="item.params"
              :limit="item.limit"
              :limitSize="item.limitSize"
              :disabled="inputDisableParams"
              :files="item.filesArr"
              @uploadOver="uploadOver"
            ></i-upload>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item>
            <el-input
              v-model="item.price"
              placeholder="请输入"
              style="width: 164px"
              size="mini"
              :disabled="inputDisableParams"
            >
              <template slot="append">元</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import { updateCommodity } from '@/api/basedata-manage'
export default {
  name: 'SizeParamsItem',
  props: {
    mode: { type: String, default: 'insert' },
    showEditBtnParams: {
      type: Boolean,
      default: false
    },
    inputDisableParams: {
      type: Boolean,
      default: false
    },
    data: {
      type: Object,
      default: null,
      required: false
    }
  },
  data() {
    return {
      limit: 1,
      limitSize: 1,
      params: { fileClass: 'enterpriseImages' },
      isWrite: false,
      canSave: false
    }
  },
  created() {
    console.log('data1111111', this.data)
  },
  watch: {
    data(newVal) {
      this.data = newVal
      // this.echoPic()
    }
  },
  computed: {
    computedIsInsertMode() {
      return this.mode === 'insert'
    },
    computedIsViewMode() {
      return this.mode === 'view'
    },
    computedUnitList() {
      let hangleArr = this.data.unitList.map(element => {
        element.weightUnit = element.weightUnit || '袋'
        let picId = element.commodityAttachment || ''
        element.params = { fileClass: 'commodityType' }
        if (picId) {
          element.filesArr = [
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
        }
        return element
      })
      console.log('结果', hangleArr)
      return hangleArr
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

    uploadSuccess(index) {
      this.data.unitList[index].commodityAttachment =
        this.$refs.uploadRef[index].files[0].fileId
      this.$emit('getFileIds', this.$refs.uploadRef[index].files, 'guige')
    },
    clickAdd() {
      this.data.unitList.push({
        weightUnit: '袋',
        weightValue: '',
        price: '',
        commodityAttachment: '',
        limit: 1,
        limitSize: 1,
        params: { fileClass: 'commodityType' },
        files: []
      })
      console.log(this.data.unitList)
      console.log(this.data)
    },

    clickDelete(index) {
      console.log(index)
      this.data.unitList.splice(index, 1)
    },

    clickWrite() {
      this.isWrite = true
      this.$emit('update:inputDisableParams', false)
    },

    clickCancel() {
      this.isWrite = false
      this.$emit('update:inputDisableParams', true)
    },
    async clickSave() {
      this.$refs['formData'].validate(async valid => {
        if (valid) {
          try {
            // 处理图片保存逻辑
            this.data.unitList = this.computedUnitList.map(res => {
              if (res.filesArr.length === 0) {
                res.commodityAttachment = ''
              } else {
                res.commodityAttachment = res.filesArr[0].fileId
              }
              return res
            })
            // 校验
            for (let i = 0; i < this.data.unitList.length; i++) {
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
            const { code } = await updateCommodity(this.data)
            if (code === 0 || code === '0') {
              this.$notify({
                title: '成功',
                message: '修改成功！',
                type: 'success',
                duration: 2000
              })
              this.isWrite = false
              this.$emit('update:inputDisableParams', true)
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

<style lang="scss" scoped>
.size-params {
  font-size: 13px;
  ::v-deep .el-form-item--medium .el-form-item__content {
    line-height: 28px;
  }
}
.custom-form {
  padding-left: 35px;
  .delete-btn {
    color: red;
    position: absolute;
    z-index: 1;
    left: -22px;
    cursor: pointer;
  }
}
.remark {
  font-size: 12px;
  margin: 6px 0 0 10px;
  color: #606266;
}
</style>
