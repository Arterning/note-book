<!--
 * @Description  : 租户授权管理
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-19 17:35:16
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\authorize-manage\index.vue
-->
<template>
  <div class="authorize-manage">
    <i-tab>
      <i-action @on-query="getTableDataImpl" @on-reset="onResetImpl" show-all>
        <i-action-item label="企业名称">
          <el-input
            clearable
            v-model="searchFields.name"
            placeholder="输入企业名称搜索"
          ></el-input>
        </i-action-item>
        <i-action-item label="服务模式">
          <el-select
            v-model="searchFields.serviceModelId"
            placeholder="选择服务模式"
            filterable
            clearable
          >
            <el-option
              v-for="item in serviceModeArr"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </i-action-item>

        <i-action-item label="企业类型">
          <el-select
            v-model="searchFields.type"
            placeholder="选择企业类型"
            filterable
            clearable
          >
            <el-option
              v-for="item in companyTypeArr"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </i-action-item>

        <template #action>
          <el-button type="primary" @click="onAdd" v-btn-permission:6704
            >新建</el-button
          >
        </template>
      </i-action>

      <i-table
        border
        :data="tableData"
        :loading="loading"
        :max-height="tableHeight"
        ref="table"
      >
        <el-table-column
          label="企业名称"
          prop="name"
          fixed="left"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column label="企业类型" prop="typeName"></el-table-column>
        <el-table-column
          label="服务模式"
          prop="serviceModelName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="系统管理员姓名"
          prop="linkman"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="管理员手机号"
          prop="linkmanPhone"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="企业地址"
          prop="address"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column label="账号状态" show-overflow-tooltip>
          <template slot-scope="{ row }">
            <el-tag disable-transitions :style="classObj(row.enable)">{{
              row.enable === 'N' ? '已冻结' : '正常'
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="权限状态" show-overflow-tooltip>
          <template slot-scope="{ row }">
            <span>{{
              row.enable !== 'N'
                ? row.status == '0'
                  ? '正常使用'
                  : row.status == '-1'
                  ? '已过期'
                  : '码已用完'
                : '无权限'
            }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          prop="createTime"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          label="创建人"
          prop="creatorName"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column label="操作" fixed="right">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              @click="editBtn(row)"
              style="margin-right: 14px"
              v-btn-permission:6705
              >编辑</el-button
            >
            <el-dropdown trigger="click">
              <span class="el-dropdown-link" style="color: #00c853">
                更多<i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>
                  <el-button
                    type="text"
                    style="color: #000"
                    @click="blockBtn(row.id, row.enable)"
                    >{{ row.enable === 'N' ? '解冻' : '冻结' }}</el-button
                  >
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </i-table>

      <i-footer v-if="tableData && tableData.length > 0">
        <!-- <i-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="current"
          :page-size="size"
          :total="total"
        /> -->
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :page-sizes="[10, 20, 50, 100]"
          :page-size="page.pageSize"
          :current-page="page.pageNum"
          :total="page.total"
          layout="total, sizes, prev, pager, next, jumper"
        >
        </el-pagination>
      </i-footer>
    </i-tab>
  </div>
</template>

<script>
import {
  getTenantList,
  getTenantOperate,
  getDictQuery
} from '@/api/authorize-manage'
import { getRole } from '@/api/tenant-manage'
import { empDictValues } from '@/utils/smallTools'
// import tableMixin from '@/views/mixins/tableMixin'
export default {
  name: 'blockchain-authorize-manage',
  // mixins: [tableMixin],
  data() {
    return {
      loading: false,
      page: {
        pageSize: 10,
        pageNum: 1,
        total: 0
      },
      searchFields: {
        name: '',
        type: '',
        serviceModelId: ''
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
      tableHeight: window.innerHeight - 440
    }
  },
  mounted() {
    // 动态计算表格高度
    this.$nextTick(() => {
      let tableTop =
        this.$refs.table.$refs.table.$el.getBoundingClientRect().top
      this.tableHeight = window.innerHeight - tableTop - 100

      window.onresize = () => {
        return (() => {
          this.tableHeight = window.innerHeight - tableTop - 100
        })()
      }
    })
    this.getTableDataImpl()
    this.getDictQuery()
    this.getRole()
  },

  methods: {
    // 企业类型字典
    async getDictQuery() {
      try {
        let { dictMap } = await getDictQuery()
        this.companyTypeArr = empDictValues(dictMap.TENANT_TYPE)
      } catch (error) {
        console.log('error', error)
      }
    },
    // 服务模式列表
    async getRole() {
      try {
        let { data } = await getRole()
        this.serviceModeArr = data
      } catch (error) {
        console.log('error', error)
      }
    },
    //  * @description: 重置（查询）的具体实现
    onResetImpl() {
      console.log(`重置`)
      this.searchFields = this.$options.data().searchFields
      this.getTableDataImpl()
    },
    // 查询租户授权管理列表
    async getTableDataImpl() {
      let payload = JSON.parse(
        JSON.stringify({
          pageNum: this.page.pageNum,
          pageSize: this.page.pageSize,
          name: ['', undefined, null].includes(this.searchFields.name)
            ? undefined
            : this.searchFields.name,
          type: ['', undefined, null].includes(this.searchFields.type)
            ? undefined
            : this.searchFields.type,
          serviceModelId: ['', undefined, null].includes(
            this.searchFields.serviceModelId
          )
            ? undefined
            : this.searchFields.serviceModelId
        })
      )
      console.log('payload', payload)
      try {
        const { pageNum, pageSize, total, list } = await getTenantList(payload)
        this.tableData = list
        this.page.pageNum = pageNum
        this.page.pageSize = pageSize
        this.page.total = total
      } catch (error) {
        console.log('查询租户授权管理列表 error = ', error)
      }
    },
    // 当前页码改变事件
    handleCurrentChange(val) {
      console.log('handleCurrentChange', val)
      this.page.pageNum = val
      // 分页接口
      this.getTableDataImpl()
      // this.getRoleUser(this.role, this.page.pageSize, this.page.pageNum)
    },
    // 当前条目数改变事件
    handleSizeChange(val) {
      console.log('handleSizeChange', val)
      this.page.pageSize = val
      // 分页接口
      this.getTableDataImpl()
      // this.getRoleUser(this.role, this.page.pageSize, this.page.pageNum)
    },
    // 新增
    onAdd() {
      this.$router.push({
        name: 'add-tenement'
      })
    },
    // 修改
    editBtn(row) {
      this.$router.push({
        name: 'edit-tenement',
        params: { row: row }
      })
    },
    // 冻结/解冻
    blockBtn(id, enable) {
      let payload = {
        id: id,
        operateCode: enable === 'N' ? 'Y' : 'N'
      }
      console.log(`payload`, payload)
      getTenantOperate(payload)
        .then(() => {
          this.$notify({
            message: '设置成功',
            type: 'success',
            duration: 2000,
            onClose: () => {
              this.getTableDataImpl()
            }
          })
        })
        .catch(err => {
          console.log('error', err)
        })
    }
  },
  computed: {
    classObj() {
      return totalStatus => {
        if (totalStatus === 'N') {
          return {
            background: '#f1f1f1',
            border: '1px solid #b1b1b1',
            color: '#000'
          }
        } else {
          return {
            background: '#bef3d0',
            border: '1px solid #3dbf72',
            color: '#00c55b'
          }
        }
        // else if (totalStatus === '2') {
        //   return {
        //     background: '#ffedc0',
        //     border: '1px solid #f76e42',
        //     color: '#f1be28'
        //   }
        // }
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

.el-tag {
  background-color: #fff;
  border: 1px solid #fff;
  color: #fff;
  width: 60px;
  text-align: center;
}
</style>
