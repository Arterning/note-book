<!--
 * @Description  : 溯源批次查询
 * @Autor        : qinjj
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-26 16:41:44
 * @FilePath     : \blockchain-platform\src\views\traceability-supervision\commodity-handling\index.vue
-->
<template>
  <div class="authorize-manage">
    <i-tab>
      <el-row>
        <el-col
          :span="4"
          style="padding-right: 20px; border-right: 1px dashed #ccc"
        >
          <el-input
            v-model="enterpriseName"
            placeholder="请输入企业名称"
            @keyup.enter.native="getEnterprise"
          >
            <el-button
              slot="append"
              icon="el-icon-search"
              @click="getEnterprise"
            ></el-button
          ></el-input>
          <div
            :style="{
              height: leftHeight + 'px',
              overflow: 'auto',
              marginTop: '10px'
            }"
            v-loading="enterpriseLoading"
          >
            <div
              class="left-item"
              :class="{ 'active-item': currentIndex === index }"
              @click="handleDetial(index, item)"
              v-for="(item, index) in companyArr"
              :key="index"
            >
              <div class="item-company">{{ item.enterpriseName }}</div>
              <div class="item-address">
                <i class="el-icon-location-outline"></i>
                {{ item.detailedAddress }}
              </div>
            </div>
            <div v-if="companyArr.length === 0" class="tips-text">暂无数据</div>
          </div>
        </el-col>
        <el-col :span="20" style="padding-left: 20px">
          <div class="company-info">
            <div class="company-item">
              <span>企业名称：</span>
              <span>{{ enterpriseDetail.enterpriseName }}</span>
            </div>
            <div class="company-item">
              <span>系统管理员：</span>
              <span>{{ enterpriseDetail.legalPerson }}</span>
            </div>
            <div class="company-item">
              <span>电话：</span>
              <span>{{ enterpriseDetail.phone }}</span>
            </div>
          </div>
          <i-action show-all @on-query="getTableData" @on-reset="onReset">
            <i-action-item label="商品名称">
              <el-input
                clearable
                v-model="searchFields.commodityName"
                placeholder="输入商品名称搜索"
              ></el-input>
            </i-action-item>
            <i-action-item label="溯源编号">
              <el-input
                clearable
                v-model="searchFields.tracingCode"
                placeholder="输入溯源编号搜索"
              ></el-input>
            </i-action-item>
            <i-action-item label="种植基地">
              <el-input
                clearable
                v-model="searchFields.farmName"
                placeholder="输入种植基地搜索"
              ></el-input>
            </i-action-item>
            <i-action-item label="加工企业">
              <el-input
                clearable
                v-model="searchFields.enterpriseName"
                placeholder="输入加工企业搜索"
              ></el-input>
            </i-action-item>
            <i-action-item label="是否有质检报告">
              <el-select
                v-model="searchFields.qualityStatus"
                clearable
                filterable
                placeholder="请选择"
              >
                <el-option
                  v-for="(item, index) in administratorNameList"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                ></el-option>
              </el-select>
            </i-action-item>
            <i-action-item label="状态">
              <el-select
                v-model="searchFields.violationStatus"
                clearable
                filterable
                placeholder="请选择"
              >
                <el-option
                  v-for="(item, index) in statusList"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                ></el-option>
              </el-select>
            </i-action-item>
          </i-action>

          <i-table
            border
            :data="tableData"
            :loading="loading"
            :max-height="tableHeight"
            ref="table"
          >
            <el-table-column
              label="商品名称"
              prop="commodityName"
              show-overflow-tooltip
              fixed="left"
              width="120"
            ></el-table-column>
            <el-table-column
              label="所属品牌"
              prop="brandName"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              label="种植基地"
              prop="farmName"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              label="是否有质检报告"
              prop="qualityStatus"
              show-overflow-tooltip
              width="120"
            >
              <template slot-scope="{ row }">
                <span>{{ row.qualityStatus === '1' ? '是' : '否' }}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="加工企业"
              prop="enterpriseName"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              label="包装件数"
              prop="packingUnit"
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              label="溯源码个数"
              prop="bondNumber"
              show-overflow-tooltip
              width="100"
            ></el-table-column>
            <el-table-column
              label="溯源码绑定编号"
              prop="bondCode"
              show-overflow-tooltip
              min-width="160"
            ></el-table-column>
            <el-table-column label="状态" width="100">
              <template slot-scope="{ row }">
                <el-tag
                  disable-transitions
                  :style="classObj(row.violationStatus)"
                  >{{ row.violationStatus === '0' ? '正常' : '已禁用' }}</el-tag
                >
              </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" width="170">
              <template slot-scope="{ row }">
                <el-button
                  type="text"
                  @click="detailBtn(row)"
                  style="margin-right: 14px"
                  >溯源详情</el-button
                >
                <el-button
                  type="text"
                  @click="
                    () => {
                      currentRow = row
                      reasonVisible = true
                    }
                  "
                  style="margin-right: 14px; color: red"
                  v-if="row.violationStatus === '0'"
                  >禁用</el-button
                >
                <el-button
                  v-else
                  type="text"
                  @click="handleRecovery(row)"
                  style="margin-right: 14px; color: #00c853"
                  :loading="cancelLoading"
                  >恢复</el-button
                >
              </template>
            </el-table-column>
          </i-table>
        </el-col>
      </el-row>
      <el-drawer
        v-if="dialogVisible"
        :visible.sync="dialogVisible"
        direction="rtl"
        :show-close="false"
        size="375px"
      >
        <template slot="title">
          <p class="title-detail">溯源详情</p>
        </template>
        <iframe
          frameborder="0"
          border="0"
          :src="`${iframeUrl}/blockchain/h5/`"
          id="H5Page"
        ></iframe>
        <div class="dialogFooter">
          <el-button plain size="mini" @click="dialogVisible = false"
            >关闭</el-button
          >
        </div>
      </el-drawer>
      <i-dialog :visible.sync="reasonVisible" title="禁用理由">
        <el-input
          type="textarea"
          placeholder="请输入禁用理由，如：该批大米查出铬超标。"
          v-model="reason"
          maxlength="100"
          rows="5"
          show-word-limit
          style="width: 100%"
        >
        </el-input>
        <div slot="footer" class="dialog-footer">
          <el-button @click="reasonVisible = false">取 消</el-button>
          <el-button
            type="primary"
            @click="reasonHandle"
            :loading="cancelLoading"
            >确 定</el-button
          >
        </div>
      </i-dialog>
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
import {
  enterpriseInfoList,
  violationCommodityList,
  updateSyCodeStatus
} from '@/api/commodity-handling.js'
import tableMixin from '@/views/mixins/tableMixin'
export default {
  name: 'blockchain-complaint-goods',
  mixins: [tableMixin],
  data() {
    return {
      cancelLoading: false,
      reasonVisible: false,
      dialogVisible: false,
      iframeUrl: null,
      reason: '',
      enterpriseName: '', // 品牌商
      loading: false,
      enterpriseLoading: false,
      currentIndex: '',
      dialogShow: false,
      companyArr: [],
      enterpriseDetail: {
        enterpriseName: '',
        legalPerson: '',
        phone: '',
        detailedAddress: ''
      },
      searchFields: {
        commodityName: '',
        tracingCode: '',
        farmName: '',
        enterpriseName: '',
        violationStatus: '',
        qualityStatus: ''
      },
      serviceModeArr: [
        { id: 1, name: '服务模式1' },
        { id: 2, name: '服务模式2' }
      ],
      companyTypeArr: [
        {
          id: 3,
          name: '企业类型1'
        },
        {
          id: 4,
          name: '企业类型2'
        }
      ],
      tableData: [],
      administratorNameList: [
        { name: '是', id: '1' },
        { name: '否', id: '0' }
      ],
      statusList: [
        { name: '正常', id: '0' },
        { name: '已禁用', id: '1' }
      ],
      leftHeight: window.innerHeight - 300,
      tableHeight: window.innerHeight - 440,
      currentRow: ''
    }
  },
  methods: {
    // 禁用理由提交按钮
    async reasonHandle() {
      if (!this.reason) {
        this.$message.warning('请先填写禁用理由')
        return
      }
      try {
        this.cancelLoading = true
        await updateSyCodeStatus({
          packingId: this.currentRow.packingInfoId,
          packingBatch: this.currentRow.packingBatch,
          status: '1',
          disableReason: this.reason
        })
        this.$notify({
          title: '成功',
          message: '禁用成功！',
          type: 'success',
          duration: 2000
        })
        this.cancelLoading = false
        this.reasonVisible = false
        this.reason = ''
        this.getTableData()
      } catch (error) {
        this.cancelLoading = false
      }
    },
    // 恢复按钮
    handleRecovery(row) {
      this.$confirm('此操作将恢复该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          try {
            this.cancelLoading = true
            await updateSyCodeStatus({
              packingId: row.packingInfoId,
              packingBatch: row.packingBatch,
              status: '0'
            })
            this.$notify({
              title: '成功',
              message: '恢复成功！',
              type: 'success',
              duration: 2000
            })
            this.cancelLoading = false
            this.getTableData()
          } catch (error) {
            this.cancelLoading = false
          }
        })
        .catch(() => {})
    },
    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        commodityName: this.searchFields.commodityName,
        tracingCode: this.searchFields.tracingCode,
        farmName: this.searchFields.farmName,
        enterpriseName: this.searchFields.enterpriseName,
        violationStatus: this.searchFields.violationStatus,
        qualityStatus: this.searchFields.qualityStatus,
        riceFactoryId: this.enterpriseDetail.enterpriseId
      }
      const data = await violationCommodityList(params)

      return { tableData: data?.list || [], total: data?.total || 0 }
    },

    onResetImpl() {
      this.searchFields.commodityName = ''
      this.searchFields.enterpriseName = ''
      this.searchFields.tracingCode = ''
      this.searchFields.farmName = ''
      this.searchFields.violationStatus = ''
      this.searchFields.qualityStatus = ''
      this.getTableData()
    },
    handleBtn() {
      this.dialogShow = true
    },
    async detailBtn(rowData) {
      this.dialogVisible = true
      await this.$nextTick()
      const iframe = document.querySelector('#H5Page')
      console.log('详情点击', rowData)
      // 处理兼容性问题
      if (iframe.attachEvent) {
        iframe.attachEvent('onload', () => {
          iframe.contentWindow.postMessage(
            {
              machPackingId: rowData.packingInfoId,
              packingBatchId: rowData.packingBatch
              // qrCode: rowData.bondCode.split(',')[0].split('-')[0]
            },
            '*'
          )
        })
      } else {
        iframe.onload = () => {
          iframe.contentWindow.postMessage(
            {
              machPackingId: rowData.packingInfoId,
              packingBatchId: rowData.packingBatch
              // qrCode: rowData.bondCode.split(',')[0].split('-')[0]
            },
            '*'
          )
        }
      }
    },
    // 查询企业信息
    async getEnterprise() {
      try {
        this.enterpriseLoading = true
        let result = await enterpriseInfoList({
          enterpriseName: this.enterpriseName
        })
        this.enterpriseLoading = false
        if (result.code === '0') {
          this.companyArr = result.data
          if (this.companyArr && this.companyArr.length > 0) {
            this.handleDetial(0, this.companyArr[0])
          }
        }
      } catch (error) {
        this.enterpriseLoading = false
      }
    },
    // 企业详情点击事件
    handleDetial(index, item) {
      this.currentIndex = index
      this.enterpriseDetail = item
      this.getTableData()
    }
  },
  created() {
    this.getEnterprise()
    this.iframeUrl = process.env.VUE_APP_BASE_URL
  },
  mounted() {
    // 监听浏览器视口高度变化

    this.$nextTick(() => {
      let tableTop =
        this.$refs.table.$refs.table.$el.getBoundingClientRect().top
      this.tableHeight = window.innerHeight - tableTop - 100

      window.onresize = () => {
        return (() => {
          this.leftHeight = window.innerHeight - 300
          this.tableHeight = window.innerHeight - tableTop - 100
        })()
      }
    })
  },
  computed: {
    classObj() {
      return totalStatus => {
        if (totalStatus === '0') {
          return {
            background: '#f1f1f1',
            border: '1px solid #b1b1b1',
            color: '#000'
          }
        } else {
          return {
            background: '#bef3d0',
            border: '1px solid #3dbf72',
            color: ' #00c55b'
          }
        }
      }
    }
  }
}
</script>

<style scoped lang="scss">
.authorize-manage {
  height: 100%;
  background-color: #fff;
  position: relative;

  iframe {
    width: 100%;
    height: 100%;
  }

  .active-item {
    color: #00c853 !important;
    .item-address {
      color: #00c853 !important;
    }
  }

  .left-item {
    padding: 10px 5px;
    border-bottom: 1px solid #ccc;
    line-height: 25px;
    cursor: pointer;
    .item-company {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    .item-address {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-size: 14px;
      color: #aaa;
    }
  }

  .tips-text {
    color: #bbb;
    text-align: center;
  }

  .company-info {
    padding: 20px;
    background: #eee;
    display: flex;
    border-radius: 5px;
    margin-bottom: 20px;
    .company-item {
      width: 33%;
    }
  }
}

.el-tag {
  background-color: #fff;
  border: 1px solid #fff;
  color: #fff;
  width: 60px;
  text-align: center;
}

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
  padding: 5px 10px;
  display: block;
  height: 8% !important;
  z-index: 999;

  .title-detail {
    font-size: 16px;
    color: #000 !important;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif !important;
  }
  .title-productBatch {
    font-size: 12px;
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
</style>
