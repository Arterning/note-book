<!--
 * @Description  : 溯源码创建
 * @Autor        : SunXiuqiong
 * @Date         : 2022-05-31 17:32:13
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-07 10:18:56
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\create\index.vue
-->
<template>
  <div class="create">
    <i-tab>
      <!-- 头部查询部分 -->
      <i-action @on-query="getTableData" @on-reset="onReset">
        <i-action-item label="品牌名称">
          <el-select
            v-model="searchFields.brandId"
            placeholder="请选择品牌名称"
            filterable
            clearable
          >
            <el-option
              v-for="item in brandList"
              :key="item.brandId"
              :label="item.brandName"
              :value="item.brandId"
            >
            </el-option>
          </el-select>
        </i-action-item>
        <i-action-item label="商品名称">
          <el-select
            v-model="searchFields.commodityId"
            placeholder="请选择商品名称"
            filterable
            clearable
          >
            <el-option
              v-for="item in commoditListSearch"
              :key="item.id"
              :label="item.commodityName"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </i-action-item>

        <template #action>
          <el-button
            type="primary"
            @click="onToggleSYCodeDialog()"
            v-btn-permission:1713
            >创建溯源码</el-button
          >
        </template>
      </i-action>

      <!-- 弹窗 -->
      <i-dialog v-if="visible" :title="dialogTitle" :visible.sync="visible">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
          <el-form-item
            label="溯源码批次"
            prop="chaincodeBatch"
            class="is-required"
          >
            <el-input
              disabled
              v-model="ruleForm.chaincodeBatch"
              placeholder="请先选择所属品牌与商品名称"
            ></el-input>
          </el-form-item>
          <el-form-item label="所属品牌" prop="brandId">
            <el-select
              placeholder="请选择品牌名称"
              filterable
              v-model="ruleForm.brandId"
              @change="changeVal"
            >
              <el-option
                v-for="item in brandList"
                :key="item.brandId"
                :label="item.brandName"
                :value="item.brandId"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="商品名称" prop="commodityId">
            <el-select
              placeholder="请选择商品名称"
              filterable
              v-model="ruleForm.commodityId"
              @change="onChange"
            >
              <el-option
                v-for="item in commoditList"
                :key="item.commodityId"
                :label="item.commodityName"
                :value="item.commodityId"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="溯源码数量" prop="chaincodeNum">
            <el-input
              placeholder="请输入数量"
              auto-complete="on"
              tabindex="1"
              :minlength="1"
              :maxlength="5"
              v-model.number="ruleForm.chaincodeNum"
            >
              <template slot="append">个</template>
            </el-input>
          </el-form-item>

          <el-form-item
            label="模板选择"
            class="template-select-view is-required"
          >
            <el-row>
              <span class="title">模板选择</span>
              <el-radio-group v-model="ruleForm.templateId">
                <template v-for="template in templates">
                  <el-radio :label="template.id" :key="template.id">
                    {{ template.label }}
                  </el-radio>
                </template>
              </el-radio-group>
            </el-row>
            <el-row>
              <span class="title">主题色</span>
              <el-radio-group v-model="ruleForm.colorId">
                <template v-for="color in colors">
                  <el-radio :label="color.id" :key="color.id">
                    {{ color.label }}
                  </el-radio>
                </template>
              </el-radio-group></el-row
            >
            <el-row>
              <div style="display: flex">
                <span>预览</span>
                <div class="tempreview">
                  <iframe
                    @load="onH5Load"
                    scrolling="auto"
                    frameborder="0"
                    border="0"
                    :src="`${iframeUrl}/blockchain/h5/`"
                    ref="tempreview-iframe"
                  ></iframe>
                </div>
              </div>
            </el-row>
          </el-form-item>
        </el-form>

        <span slot="footer" class="dialog-footer">
          <el-button @click="onToggleSYCodeDialog(false)">取 消</el-button>
          <el-button type="primary" @click="submit">创 建</el-button>
        </span>
      </i-dialog>

      <!-- 表格部分 -->
      <i-table border :data="tableData" :loading="loading" :maxHeight="450">
        <el-table-column prop="chainCodeBatch" label="溯源码批次">
        </el-table-column>
        <el-table-column prop="commodityName" label="商品名称">
        </el-table-column>
        <el-table-column
          prop="brandName"
          label="所属品牌"
          show-overflow-tooltip
        >
        </el-table-column>
        <el-table-column prop="chaincodeNum" label="溯源码个数">
        </el-table-column>
        <el-table-column label="绑定状态">
          <template slot-scope="{ row }">
            <el-tag disable-transitions :style="classObj(row.bondStatus)">{{
              row.bondStatus === '1'
                ? '未绑定'
                : row.bondStatus === '2'
                ? '部分绑定'
                : '全部绑定'
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="unBondNum" label="未绑定数量"></el-table-column>
        <el-table-column prop="createDate" label="创建时间"> </el-table-column>
        <el-table-column prop="creator" label="创建人" width="100px">
        </el-table-column>
        <el-table-column label="操作" width="90px">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              @click="clickDown(row)"
              :disabled="row.unBondNum === '0'"
              >下载</el-button
            >
          </template>
        </el-table-column>
      </i-table>

      <!-- footer区域 -->
      <i-footer v-if="tableData && tableData.length > 0">
        <i-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="current"
          :page-size="size"
          :total="total"
        />
      </i-footer>
    </i-tab>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import tableMixin from '@/views/mixins/tableMixin'
import {
  getList,
  getbrandList,
  getChaincodeBatch,
  createCode,
  downFile,
  getGoods
} from '@/api/traceability-manage'

export default {
  name: 'traceability-manage--create',
  mixins: [tableMixin],
  data() {
    //创建溯源码校验方法
    var checkChaincodeNum = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('数量不能为空'))
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入数字值'))
        } else {
          if (value === 0) {
            callback(new Error('请至少选择1个'))
          } else {
            callback()
          }
        }
      }, 100)
    }
    return {
      searchFields: {
        brandId: '',
        commodityId: ''
      },
      brandList: [],
      commoditList: [], //存储所属品牌下面的商品列表
      commoditListSearch: [],
      tableData: [],
      dialogTitle: '创建溯源',
      loading: false,
      visible: false,
      ruleForm: {
        chaincodeBatchId: '', //溯源码批次Id
        chaincodeBatch: '', //溯源码批次
        brandId: '', //品牌
        commodityId: '', //商品
        chaincodeNum: '', //创建溯源码数量
        templateId: 1, //模板id
        colorId: 1 //颜色id
      },
      rules: {
        brandId: [
          { required: true, message: '请选择品牌名称', trigger: 'change' }
        ],
        commodityId: [
          {
            required: true,
            message: '请选择商品名称'
          }
        ],
        chaincodeNum: [
          {
            required: true,
            message: '数量不能为空'
          },
          { validator: checkChaincodeNum, trigger: 'blur' }
        ]
      },
      templates: [
        { label: '模板一', id: 1 },
        { label: '模板二', id: 2 }
      ],
      colors: [
        { label: '黄色', id: 1 },
        { label: '绿色', id: 2 },
        { label: '蓝色', id: 3 }
      ],
      iframeUrl: null
    }
  },
  watch: {
    'ruleForm.brandId': {
      handler() {
        this.ruleForm.commodityId = null
        this.clearRuleFormValidate('commodityId')
      }
    },
    'ruleForm.templateId': {
      handler(templateId) {
        this.postMessage({ templateId, colorId: this.ruleForm.colorId })
      }
    },
    'ruleForm.colorId': {
      handler(colorId) {
        this.postMessage({ colorId, templateId: this.ruleForm.templateId })
      }
    }
  },
  created() {
    this.getTableData()
    this.getbrandList()
    this.getGoods()
    this.iframeUrl = process.env.VUE_APP_BASE_URL
  },

  methods: {
    // 向h5传值
    postMessage(payload) {
      console.log(payload)
      this.$refs['tempreview-iframe']?.contentWindow?.postMessage?.(
        payload,
        '*'
      )
    },

    // 获取所有品牌
    async getbrandList() {
      console.log('chaxun', this.computedUserInfo)
      const { data } = await getbrandList({
        pageSize: -1,
        enterpriseId: this.computedUserInfo.enterpriseId
      })
      this.brandList = data
    },

    // 获取所有商品
    async getGoods() {
      const { data } = await getGoods()
      this.commoditListSearch = data
    },

    //选择品牌下的对应商品的changeVal方法
    changeVal(value) {
      // this.searchFields.commodityId = null
      for (var i = 0; i < this.brandList.length; i++) {
        if (value === this.brandList[i].brandId) {
          this.commoditList = this.brandList[i].commodities
        }
      }
    },

    // 当品牌跟商品同时选择后才获取溯源码批次
    onChange() {
      if (this.ruleForm.brandId !== '' && this.ruleForm.commodityId !== '') {
        console.log(this.ruleForm.brandId)
        console.log(this.ruleForm.commodityId)
        this.getChaincodeBatch()
      }
    },

    // 获取溯源码批次
    async getChaincodeBatch() {
      const { data } = await getChaincodeBatch({
        brandId: this.ruleForm.brandId,
        commodityId: this.ruleForm.commodityId
      })
      this.ruleForm.chaincodeBatch = data.chainCodeBatch
      this.ruleForm.chaincodeBatchId = data.id
    },

    // 点击下载事件
    async clickDown(row) {
      try {
        const res = await downFile(
          { chaincodeBatchId: row.id },
          `溯源码-${row.commodityName}-${row.chainCodeBatch}.xls`
        )
        console.log(res)
        // this.download(res, 'application/json;charset=UTF-8', 'test.excel')
      } catch (error) {
        console.log('下载 err' + error)
      }
    },

    download(res, type, filename) {
      // 创建blob对象，解析流数据
      const blob = new Blob([res], {
        // 设置返回的文件类型
        // type: 'application/pdf;charset=UTF-8' 表示下载文档为pdf，如果是word则设置为msword，excel为excel
        type: type
      })
      // 这里就是创建一个a标签，等下用来模拟点击事件
      const a = document.createElement('a')
      // 兼容webkix浏览器，处理webkit浏览器中href自动添加blob前缀，默认在浏览器打开而不是下载
      const URL = window.URL || window.webkitURL
      // 根据解析后的blob对象创建URL 对象
      const herf = URL.createObjectURL(blob)
      // 下载链接
      a.href = herf
      // 下载文件名,如果后端没有返回，可以自己写a.download = '文件.pdf'
      a.download = filename
      document.body.appendChild(a)
      // 点击a标签，进行下载
      a.click()
      // 收尾工作，在内存中移除URL 对象
      document.body.removeChild(a)
      window.URL.revokeObjectURL(herf)
    },

    // 点击重置，清空选择项
    onResetImpl() {
      this.searchFields.commodityId = ''
      this.searchFields.brandId = ''
      this.getTableData()
    },

    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        brandId: this.searchFields.brandId,
        commodityId: this.searchFields.commodityId
      }
      const { data, total } = await getList(params)

      return { tableData: data, total }
    },

    submit() {
      this.$refs.ruleForm.validate(async isOk => {
        if (isOk) {
          // 校验通过
          try {
            await createCode(this.ruleForm)
            this.$notify({
              title: '成功',
              message: '创建成功！',
              type: 'success',
              duration: 2000
            })
            this.getTableData()
            this.onToggleSYCodeDialog(false)
          } catch (error) {
            return false
          }
        }
      })
    },
    onH5Load() {
      // console.log(this.ruleForm.templateId, this.ruleForm.colorId)
      this.postMessage({
        templateId: this.ruleForm.templateId,
        colorId: this.ruleForm.colorId
      })
    },

    async onToggleSYCodeDialog(flag = true) {
      if (flag) {
        Object.keys(this.ruleForm).forEach(key => (this.ruleForm[key] = null))
        this.ruleForm.templateId = this.templates?.[0]?.id || null
        this.ruleForm.colorId = this.colors?.[0]?.id || null
      }

      this.visible = flag
    },

    clearRuleFormValidate(prop) {
      this.$nextTick(() => {
        this.$refs.ruleForm.clearValidate(prop)
      })
    }
  },
  computed: {
    classObj() {
      return totalStatus => {
        if (totalStatus === '1') {
          return {
            background: '#f1f1f1',
            border: '1px solid #b1b1b1',
            color: '#000'
          }
        } else if (totalStatus === '2') {
          return {
            background: '#f7e1b3',
            border: '1px solid #bf9e4a',
            color: '#be9330'
          }
        } else if (totalStatus === '3') {
          return {
            background: '#a7dbc0',
            border: '1px solid #40745b',
            color: '#489974'
          }
        }
      }
    },
    ...mapGetters('user', { computedUserInfo: 'userInfoGetter' })
  }
}
</script>

<style lang="scss" scoped>
.create {
  height: 100%;
  background-color: #fff;
  position: relative;

  ::v-deep .el-row {
    margin-left: -50px;
  }
  .title {
    display: inline-block;
    width: 80px;
    padding-right: 20px;
  }

  .tempreview {
    width: 300px;
    height: 380px;
    margin: 12px 0 0 53px;
    border: 1px solid #ccc;

    iframe {
      width: 100%;
      height: 100%;
    }
  }
}
</style>
<style lang="scss">
.template-select-view {
  display: flex;
  flex-direction: column;
  .el-form-item__content {
    margin-left: 85px !important;
  }
}
.el-tag {
  background-color: #fff;
  border: 1px solid #fff;
  color: #fff;
  width: 80px;
  text-align: center;
}
</style>
