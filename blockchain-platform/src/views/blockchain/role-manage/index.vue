<!--
 * @Description  : 角色管理
 * @Autor        : Hemingway
 * @Date         : 2021-06-21 21:08:56
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-06 11:49:36
 * @FilePath     : \blockchain-platform\src\views\blockchain\role-manage\index.vue
-->
<template>
  <div class="role-manage">
    <i-tab>
      <i-action @on-query="getTableData" @on-reset="onReset" :loading="loading">
        <i-action-item label="角色名称">
          <el-input
            v-model="roleName"
            @keyup.enter.native="getTableData"
            placeholder="输入角色名称"
            clearable
          >
          </el-input>
        </i-action-item>
      </i-action>

      <!-- 主体内容区域 -->
      <div class="body">
        <i-table
          :loading="loading"
          :data="tableData"
          @selection-change="onSelectionChange"
          ref="table"
        >
          <el-table-column label="角色名称" prop="name"> </el-table-column>
          <el-table-column label="备注" prop="remark"></el-table-column>
          <el-table-column label="创建时间" prop="modifyDate"></el-table-column>
          <el-table-column label="操作">
            <template slot-scope="{ row }">
              <el-button @click="handleClick(row, 0)" type="text"
                >详情</el-button
              >
              <el-button @click="handleClick(row, 1)" type="text"
                >编辑</el-button
              >
            </template>
          </el-table-column>
        </i-table>
      </div>

      <!-- footer区域 -->
      <i-footer v-if="tableData.length > 0">
        <i-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="current"
          :page-size="size"
          :total="total"
        >
        </i-pagination>
      </i-footer>
    </i-tab>

    <i-dialog
      :visible.sync="dialogWindowVisible"
      :title="dialogName"
      ref="dialog"
    >
      <i-section-header title="基本信息" />
      <div v-if="role.disableEdit">
        <ul style="line-height: 20px; list-style: none; padding-left: 0px">
          <li>角色名称: {{ role.name }}</li>
          <li v-if="role.remark" style="margin-top: 10px">
            备注: {{ role.remark }}
          </li>
        </ul>
      </div>
      <div v-else>
        <el-form ref="roleForm" :model="form">
          <el-form-item
            :rules="[{ required: true, message: '角色名不能为空' }]"
            required
            label="角色名称"
            prop="name"
          >
            <el-input v-model="form.name" placeholder="请输入"></el-input>
          </el-form-item>
          <el-form-item label="备注">
            <el-input v-model="form.remark"></el-input>
          </el-form-item>
        </el-form>
      </div>

      <i-section-header title="角色权限" />
      <role-manage-table
        :roleRightListDto="role.roleRightListDto"
        :disableEdit="role.disableEdit"
      />

      <template #footer
        ><el-button @click="goUpdateRole" type="success" v-if="role.disableEdit"
          >修 改</el-button
        >
        <div v-else>
          <el-button @click="roleCreate('修改成功')" type="success"
            >确 定</el-button
          >
          <el-button @click="dialogDismiss">取 消</el-button>
        </div></template
      >
    </i-dialog>
  </div>
</template>

<script>
import tableMixin from '@/views/mixins/tableMixin'

import {
  getRoleListData,
  queryMenuTreeList,
  getRoleJurisdictionData,
  saveRoleRightsTree,
  deleteRole,
  createRole,
  updateRole
} from '@/api/blockchain/roleManage'
import RoleManageTable from '../components/role-manage-table'

export default {
  name: 'blockchain--blockchain-role-manage',

  components: { RoleManageTable },

  mixins: [tableMixin],

  data() {
    return {
      roleName: '',
      orignalMenuTree: {},
      orignalMenuList: [],
      dialogWindowVisible: false,
      form: {},
      role: {
        id: ''
      },
      dialogName: ''
    }
  },

  mounted() {
    this.init()
  },

  methods: {
    async init() {
      this.getTableData()

      const { menuTree } = await queryMenuTreeList()
      const roleRightTreeDto = { roleMenuList: [...menuTree] }
      this.orignalMenuList = this.analysisTreeMenu(roleRightTreeDto)
      this.orignalMenuTree = roleRightTreeDto
    },

    dialogDismiss() {
      this.dialogWindowVisible = false
    },

    /**
     * @description: 重置的具体实现
     * @author: Hemingway
     */
    onResetImpl() {
      this.roleName = ''
    },

    /**
     * @description: 列表查询逻辑的具体实现
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getTableDataImpl(payload) {
      Object.assign(payload, { roleName: this.roleName }) // 添加额外请求参数

      try {
        const { list, total } = await getRoleListData(payload)
        const tableData = list || []

        return { tableData, total }
      } catch (error) {
        console.log('查询列表 error = ', error)
      }
    },

    // 角色列表的操作按钮（0:详情 1:修改 2:删除）
    handleClick(item, type) {
      if (type === 2) {
        this.roleDelete(item)
      } else {
        this.form = {
          name: item.name,
          remark: item.remark,
          type: item.type
        }
        this.showRole(item, type === 0)
      }
    },

    // 查看角色详情
    async showRole(item, disableEdit) {
      try {
        const { code, roleRightTreeDto } = await getRoleJurisdictionData({
          roleId: item.id
        })
        if (code === '0') {
          item.roleRightTreeDto = roleRightTreeDto // eslint-disable-line
          item.roleRightListDto = this.analysisTreeMenu(roleRightTreeDto) // eslint-disable-line
          this.dialogWindowVisible = true
          this.role = item
          this.role.disableEdit = disableEdit
          this.dialogName = disableEdit ? '角色详情' : '编辑角色'
          this.dialogWindowVisible = true
        }
      } catch (error) {
        console.log(error)
        this.$message.error('请求失败，请稍后重试')
      }
    },

    goUpdateRole() {
      this.role.disableEdit = false
      this.dialogWindowVisible = false
      this.$nextTick(() => {
        this.dialogWindowVisible = true
      })
    },

    // 新增角色
    roleCreate() {
      this.$refs.roleForm.validate(valid => {
        if (valid) {
          if (this.role.id) {
            // 先调新增角色接口，再调角色更新权限接口
            updateRole({
              id: this.role.id,
              oldName: this.role.name,
              enable: this.role.enable,
              ...this.form
            }).then(() => this.roleUpdateRightTree('修改成功'))
          } else {
            // 先调更新角色接口，再调角色更新权限接口
            createRole(this.form).then(res => {
              if (res.code === '0') {
                this.role.id = res.id
                this.roleUpdateRightTree('新增成功')
              }
            })
          }
        } else {
          return false
        }
      })
    },

    // 删除角色
    roleDelete(role) {
      this.$confirm('您将删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        closeOnClickModal: false
      }).then(() => {
        this.roleDeleteQuery(role.id)
      })
    },
    async roleDeleteQuery(roleId) {
      const { code } = await deleteRole({ roleId })
      if (code === '0') {
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
        this.getTableData()
      } else {
        console.log('删除失败')
      }
    },

    // 修改角色权限
    async roleUpdateRightTree(message) {
      this.updateTreeMenuItem(this.role.roleRightTreeDto)
      const treeJson = JSON.stringify({
        roleId: this.role.id,
        roleRightTreeDto: this.role.roleRightTreeDto
      })
      const { code } = await saveRoleRightsTree({
        roleRightsList: treeJson
      })
      if (code === '0') {
        this.$message({
          message,
          type: 'success'
        })
        this.dialogDismiss()
        this.getTableData()
      }
    },

    // 修改原树状数据的值
    updateTreeMenuItem(menu) {
      menu.menuId = menu.id
      if (!menu.roleMenuList) {
        const role = this.role.roleRightListDto.find(({ id }) => id === menu.id)
        menu.isCheck = role['3']
      } else {
        let isCheck = false
        menu.roleMenuList.forEach(item => {
          this.updateTreeMenuItem(item)
          if (item.isCheck) {
            isCheck = true
          }
        })
        menu.isCheck = isCheck
      }
    },

    // 解析树状结构权限变成表格需要的列表
    analysisTreeMenu(roleRightTreeDto) {
      let resultList = []
      if (
        !roleRightTreeDto.roleMenuList &&
        roleRightTreeDto.children &&
        roleRightTreeDto.children.length > 0
      ) {
        roleRightTreeDto.roleMenuList = roleRightTreeDto.children
      }
      roleRightTreeDto.roleMenuList.forEach(menu => {
        this.computedMenuList(menu, resultList)
      })
      return resultList
    },

    // 递归逐条解析
    computedMenuList(menu, resultList, i = 0, cell = {}) {
      if (!menu.roleMenuList && menu.children && menu.children.length > 0) {
        menu.roleMenuList = menu.children
      }
      cell[i] = menu.name
      if (!menu.roleMenuList) {
        cell[i + 1] = menu.isCheck
        cell.id = menu.id
        resultList.push(cell)
      } else {
        menu.roleMenuList.forEach((item, index) => {
          if (index === 0) {
            this.computedMenuList(item, resultList, i + 1, cell)
          } else {
            this.computedMenuList(item, resultList, i + 1)
          }
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.role-manage {
  height: 100%;
  background-color: #fff;
  position: relative;

  .body {
    padding-bottom: 56px; // 预留 i-footer 位置
  }
}
</style>
