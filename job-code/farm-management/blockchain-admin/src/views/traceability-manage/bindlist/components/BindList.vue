<!--
 * @Description  : 已绑定列表
 * @Autor        : SunXiuqiong
 * @Date         : 2022-06-06 10:52:26
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-21 11:23:59
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\bindlist\components\BindList.vue
-->
<template>
  <div class="bindlist-area">
    <i-action :showAll="showAll" @on-query="getTableData" @on-reset="onReset">
      <i-action-item label="品种名称">
        <el-select
          v-model="condition.varietyId"
          placeholder="请选择品种名称"
          filterable
          clearable
        >
          <el-option
            :label="item.varietyName"
            :value="item.varietyId"
            v-for="item in varietyInfos"
            :key="item.varietyId"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="商品名称">
        <el-select
          v-model="condition.commodityId"
          placeholder="请选择商品名称"
          filterable
          clearable
        >
          <el-option
            :label="item.commodityName"
            :value="item.commodityId"
            v-for="item in commodityInfos"
            :key="item.commodityId"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="来源农场">
        <el-select
          v-model="condition.farmId"
          placeholder="请选择来源农场"
          filterable
          clearable
        >
          <el-option
            :label="item.farmName"
            :value="item.farmId"
            v-for="item in farmInfos"
            :key="item.farmId"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="米厂">
        <el-select
          v-model="condition.riceFactoryId"
          placeholder="请选择米厂"
          filterable
          clearable
        >
          <el-option
            v-for="item in riceFactoryInfos"
            :label="item.riceFactoryName"
            :value="item.riceFatoryId"
            :key="item.riceFatoryId"
          ></el-option>
        </el-select>
      </i-action-item>
    </i-action>
    <!-- 表格区域 -->

    <i-table
      :data="tableData"
      :maxHeight="450"
      border
      style="width: 99%"
      :loading="loading"
    >
      <el-table-column label="产品批次" fixed="left" minWidth="190px">
        <template slot-scope="{ row }">
          <el-button type="text" size="small" @click="CheckDeltail(row)">{{
            row.productBatch
          }}</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="varietyName" label="品种名称" minWidth="130">
      </el-table-column>
      <el-table-column prop="commodityName" label="商品名称" minWidth="130">
      </el-table-column>
      <el-table-column prop="brandName" label="所属品牌" minWidth="150">
      </el-table-column>
      <el-table-column
        prop="packingDate"
        label="包装日期"
        minWidth="120"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        prop="packingUnit"
        label="包装规格"
        minWidth="120"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        prop="riceFactoryName"
        label="米厂"
        minWidth="130"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        prop="farmName"
        label="来源农场"
        minWidth="130"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column prop="bondNumber" label="溯源码绑定个数" minWidth="120">
      </el-table-column>
      <el-table-column
        prop="bondCode"
        label="溯源码绑定编号"
        minWidth="240"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        prop="modifier"
        label="操作人"
        minWidth="110"
        show-overflow-tooltip
      >
      </el-table-column>
      <el-table-column
        prop="bondDate"
        label="绑定时间"
        minWidth="130"
        show-overflow-tooltip
      >
      </el-table-column>
    </i-table>

    <!-- <detail-dialog :visible.sync="dialogVisible" /> -->
    <el-drawer
      v-if="dialogVisible"
      :visible.sync="dialogVisible"
      :direction="direction"
      :show-close="false"
      size="375px"
    >
      <template slot="title">
        <p class="title-detail">溯源详情</p>
        <p class="title-productBatch">
          (&nbsp;产品批次&nbsp;:&nbsp;<span>{{ proBatch }}</span>
          &nbsp;)
        </p>
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

    <i-footer v-if="tableData && tableData.length > 0">
      <i-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="current"
        :page-size="size"
        :total="total"
      />
    </i-footer>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getBindList, getConditionList } from '@/api/traceability-manage'
import tableMixin from '@/views/mixins/tableMixin'
import IFooter from '@/components/i-footer' // 尾部栏
import IPagination from '@/components/i-pagination'
// import DetailDialog from './DetailDialog.vue'

export default {
  name: 'BindList',
  mixins: [tableMixin],
  components: {
    IPagination,
    IFooter
    // DetailDialog
  },

  data() {
    return {
      direction: 'rtl',
      loading: false,
      bondStatus: 1,
      condition: {
        varietyId: '',
        commodityId: '',
        farmId: '',
        riceFactoryId: ''
      },
      varietyInfos: [], //品种信息
      commodityInfos: [], //商品信息
      farmInfos: [], //农场信息
      riceFactoryInfos: [], //米厂信息
      dialogVisible: false,
      title: '溯源详情',
      showAll: true,
      tableData: [],
      proBatch: '',
      iframeUrl: null
    }
  },
  created() {
    this.getConditionList()
    this.getTableData()
    this.iframeUrl = process.env.VUE_APP_BASE_URL
    // this.iframeUrl = 'https://localhost:8080'
  },

  computed: {
    ...mapGetters('user', { computedUserInfo: 'userInfoGetter' })
  },

  methods: {
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        varietyId: this.condition.varietyId,
        commodityId: this.condition.commodityId,
        farmId: this.condition.farmId,
        riceFactoryId: this.condition.riceFactoryId,
        bondStatus: this.bondStatus,
        enterpriseId: this.computedUserInfo.enterpriseId
      }
      const { data } = await getBindList(params)

      return { tableData: data?.list || [], total: data?.total || 0 }
    },

    // 筛选查询
    async getConditionList() {
      try {
        const { data } = await getConditionList()
        this.varietyInfos = data.varietyInfos
        this.commodityInfos = data.commodityInfos
        this.farmInfos = data.farmInfos
        this.riceFactoryInfos = data.riceFactoryInfos
      } catch (error) {
        console.log('查询错误 err:' + error)
      }
    },

    // 点击重置，清空选择项
    onResetImpl() {
      this.condition = {
        varietyId: '',
        brandId: '',
        farmId: '',
        riceFactoryId: ''
      }
      this.getTableData()
    },

    async CheckDeltail(rowData) {
      this.proBatch = rowData.productBatch
      this.dialogVisible = true
      await this.$nextTick()
      const iframe = document.querySelector('#H5Page')
      // 处理兼容性问题
      if (iframe.attachEvent) {
        iframe.attachEvent('onload', () => {
          iframe.contentWindow.postMessage(
            {
              machPackingId: rowData.packingInfoId,
              packingBatchId: rowData.packingBatch,
              qrCode: rowData.bondCode.split(',')[0].split('-')[0]
            },
            '*'
          )
        })
      } else {
        iframe.onload = () => {
          iframe.contentWindow.postMessage(
            {
              machPackingId: rowData.packingInfoId,
              packingBatchId: rowData.packingBatch,
              qrCode: rowData.bondCode.split(',')[0].split('-')[0]
            },
            '*'
          )
        }
      }
    }
  }
}
</script>

<style lang="scss" scpoed>
.el-drawer__wrapper {
  height: 667px;
  margin-top: 100px;
  border-radius: 6px;
}
.el-drawer__header {
  border-bottom: 1px solid #ccc;
  padding: 10px;
  display: block;
  height: 10% !important;
  z-index: 999;
}

.el-drawer__body {
  padding: 0 0 47px 0;
  margin-top: -31px;
  overflow: hidden;
}

.dialogFooter {
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

<style scoped lang="scss">
.bindlist-area {
  iframe {
    width: 100%;
    height: 100%;
  }
}

// ::v-deep .el-table--scrollable-y .el-table__body-wrapper {
//   overflow-y: hidden;
// }
</style>
