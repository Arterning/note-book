<!--
 * @Description  : 区块链明细数据
 * @Autor        : qinjj
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-26 13:51:57
 * @FilePath     : \blockchain-platform\src\views\blockchain-query\detail.vue
-->
<template>
  <div class="authorize-manage">
    <i-action @on-query="getTableData" @on-reset="onReset" showAll>
      <i-action-item label="上链时间">
        <el-date-picker
          v-model="time"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          format="yyyy-MM-dd"
          value-format="yyyy-MM-dd"
        >
        </el-date-picker>
      </i-action-item>
      <i-action-item label="数据区块">
        <el-select
          v-model="searchFields.dataType"
          clearable
          filterable
          placeholder="请选择数据区块"
        >
          <el-option
            v-for="(item, index) in blockList"
            :key="index"
            :label="item.label"
            :value="item.value"
          ></el-option>
        </el-select>
      </i-action-item>
      <i-action-item label="状态">
        <el-select
          v-model="searchFields.status"
          clearable
          filterable
          placeholder="请选择状态"
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
      :max-height="540"
      ref="table"
    >
      <el-table-column
        label="种植基地"
        prop="plantBase"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="作物品类"
        prop="cropsName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="作物品种"
        prop="varietyName"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="数据区块"
        prop="dataType"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="最早上链时间"
        prop="firstUploadTime"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="状态" prop="status" show-overflow-tooltip>
        <template slot-scope="{ row }">
          <el-tag disable-transitions :style="classObj(row.status)">{{
            row.status === '0' ? '正常' : '已修改'
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        label="最新修改时间"
        prop="latestUpdateTime"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="上传单位"
        prop="uploadOrg"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column label="操作" fixed="right">
        <template slot-scope="{ row }">
          <el-button
            type="text"
            @click="detailBtn(row)"
            style="margin-right: 14px"
            >查看详情</el-button
          >
        </template>
      </el-table-column>
    </i-table>
    <i-dialog :visible.sync="lsVisible" title="链上数据详情" width="60%">
      <div v-if="newData.length >= 1">
        <i-section-header
          :title="`最新内容（时间：${newTime}）`"
          class="title-box"
        />
        <p>{{ newData[0].content }}</p>
      </div>

      <div v-if="historyData.length > 0">
        <i-section-header title="历史内容" class="title-box" />
        <div v-for="(item, index) in historyData" :key="index">
          <p>
            <span>时间：</span>
            {{ item.uploadTime }}
          </p>
          <p>
            <span>内容：</span>
            {{ item.content }}
          </p>
        </div>
      </div>
      <div style="text-align: center" v-if="newData.length === 0">暂无数据</div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="lsVisible = false">关 闭</el-button>
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
  </div>
</template>

<script>
import tableMixin from '@/views/mixins/tableMixin'
import { getList, getHistory, getDict } from '@/api/blockchain-index.js'
// import { parseTime } from '@/utils'
export default {
  name: 'blockchain-detail',
  mixins: [tableMixin],
  data() {
    return {
      lsVisible: false,
      time: '',
      searchFields: {
        queryStartTime: '',
        queryEndTime: '',
        status: '',
        dataType: ''
      },
      auditStatusList: [],
      tableData: [],
      statusList: [
        {
          name: '正常',
          id: '0'
        },
        {
          name: '已修改',
          id: '1'
        }
      ],
      blockList: [],
      newData: [],
      historyData: [],
      newTime: '',
      tableHeight: window.innerHeight - 440
    }
  },
  filters: {
    toThousands(val) {
      let str = val.toString()
      return str.replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
    }
  },
  async created() {
    // 获取数据区块链
    this.getDict()
    this.getTableData()
  },
  mounted() {},
  methods: {
    // 获取数据区块链
    async getDict() {
      this.blockList = []
      let result = await getDict()
      if (result.dictMap.DATA_TYPE) {
        for (let item in result.dictMap.DATA_TYPE) {
          this.blockList.push({
            label: result.dictMap.DATA_TYPE[item],
            value: item
          })
        }
      }
    },
    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        queryStartTime: this.time ? this.time[0] : '',
        queryEndTime: this.time ? this.time[1] : '',
        status: this.searchFields.status,
        dataType: this.searchFields.dataType
        // dataType: '2'
      }
      const data = await getList(params)

      return {
        tableData: data?.list || [],
        total: data?.total || 0
      }
    },

    onResetImpl() {
      this.time = ''
      this.searchFields.queryStartTime = ''
      this.searchFields.queryEndTime = ''
      this.searchFields.status = ''
      this.searchFields.dataType = ''
      this.getTableData()
    },

    async detailBtn(row) {
      this.newData = []
      this.historyData = []
      let result = await getHistory({ queryKey: row.id })
      if (!result.data) return
      let data = result.data.map(res => {
        res.uploadTime = res.uploadTime.replace(/T|Z/g, ' ')
        res.content = JSON.parse(res.content).upload_context
        return res
      })
      this.newTime = row.latestUpdateTime || row.firstUploadTime
      this.newData = data.slice(0, 1)
      this.historyData = data.slice(1)
      this.lsVisible = true
    }
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
}
</style>
