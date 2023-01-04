<!--
 * @Description  : 用户管理
 * @Autor        : Hemingway
 * @Date         : 2021-12-22 19:59:51
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-21 09:29:44
 * @FilePath     : \blockchain-platform\src\views\system-manage\user\index.vue
-->
<template>
  <div class="user">
    <i-tab>
      <i-action @on-query="getTableData" @on-reset="onReset">
        <i-action-item label="姓名">
          <el-input
            v-model="searchFields.userName"
            placeholder="请输入员工姓名搜索"
            clearable
          >
          </el-input>
        </i-action-item>
        <i-action-item label="角色">
          <el-select
            v-model="searchFields.roleId"
            placeholder="请选择角色搜索"
            filterable
            clearable
          >
            <el-option
              v-for="item in roleList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </i-action-item>

        <template #action>
          <el-button type="primary" @click="clickAdd()" v-btn-permission:8114
            >新建</el-button
          >
        </template>
      </i-action>

      <!-- 新建角色弹窗 -->
      <i-dialog
        v-if="visible"
        :title="dialogTitle"
        :visible="visible"
        @close="cancelBtn"
      >
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
          <el-form-item label="姓名" prop="name">
            <el-input placeholder="请输入姓名" v-model="ruleForm.name">
            </el-input>
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input
              placeholder="请输入手机号"
              maxlength="11"
              v-model="ruleForm.phone"
            >
            </el-input>
          </el-form-item>
          <el-form-item label="角色" prop="roleId">
            <el-select placeholder="请选择角色" v-model="ruleForm.roleId">
              <el-option
                v-for="item in roleList"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>

        <span slot="footer" class="dialog-footer">
          <el-button @click="cancelBtn">取 消</el-button>
          <el-button type="primary" @click="onSubmitStaffialog"
            >确 定</el-button
          >
        </span>
      </i-dialog>

      <i-table border :data="tableData" :loading="loading" :maxHeight="450">
        <el-table-column
          prop="name"
          label="姓名"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          prop="phone"
          label="手机号"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          prop="roleName"
          label="角色名称"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column prop="status" label="状态">
          <template slot-scope="{ row }">
            {{ row.status === '1' ? '正常' : '冻结' }}</template
          >
        </el-table-column>
        <el-table-column
          prop="createDate"
          label="创建时间"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column
          prop="creator"
          label="创建人"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              @click="editBtn(row)"
              style="margin-right: 14px"
              v-btn-permission:8115
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
                    @click="blockBtn(row)"
                  >
                    {{ row.status === '1' ? '冻结' : '解冻' }}
                  </el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button
                    type="text"
                    style="color: red"
                    @click="deleteBth(row)"
                    >删除</el-button
                  >
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </i-table>

      <!-- footer区域footer -->
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
import tableMixin from '@/views/mixins/tableMixin'
import {
  getStaffList,
  createStaff,
  editStaff,
  deleteStaff,
  operateAccount,
  getRole
} from '@/api/system-manage'

export default {
  name: 'system-manage--user',
  mixins: [tableMixin],
  data() {
    return {
      loading: false,
      searchFields: {
        userName: '',
        roleId: ''
      },
      roleList: [
        {
          id: '1',
          name: '超级管理员'
        },
        {
          id: '2',
          name: '管理员'
        }
      ],

      tableData: [],

      visible: false,
      dialogTitle: '',
      dialogFlag: '', // add或者edit
      ruleForm: {
        name: '',
        phone: '',
        roleId: '',
        userId: ''
      },
      rules: {
        name: [
          { required: true, message: '请输入姓名', trigger: 'blur' },
          { min: 2, max: 5, message: '长度在 2 到 5 个字符', trigger: 'blur' }
        ],
        phone: [
          {
            required: true,
            message: '请输入手机号',
            trigger: 'blur'
          },
          {
            pattern: /^1[3456789]\d{9}$/,
            message: '请输入合法手机号',
            trigger: 'blur'
          }
        ],
        roleId: [
          {
            required: true,
            message: '请选择角色',
            trigger: 'change'
          }
        ]
      }
    }
  },
  created() {
    this.getRole()
    this.getTableData()
  },
  methods: {
    // 获取角色列表
    async getRole() {
      const { data } = await getRole({
        enable: 'Y'
      })
      this.roleList = data
    },

    // 点击查询事件
    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        userName: this.searchFields.userName,
        roleId: this.searchFields.roleId
      }
      const res = await getStaffList(params)

      return { tableData: res.list, total: res.total }
    },

    onResetImpl() {
      this.searchFields.userName = ''
      this.searchFields.roleId = ''
      this.getTableData()
    },
    clickAdd() {
      console.log(this.searchFields)
      this.dialogTitle = '新建员工'
      this.visible = true
      this.dialogFlag = 'add'
    },

    editBtn(row) {
      this.ruleForm.name = row.name
      this.ruleForm.phone = row.phone
      this.ruleForm.userId = row.id
      // this.ruleForm.roleId = row.roleName
      this.ruleForm.roleId = row.roleId
      this.dialogTitle = '编辑角色'
      this.visible = true
      this.dialogFlag = 'edit'
    },

    onSubmitStaffialog() {
      let api,
        payload = null
      if (this.dialogFlag === 'add') {
        console.log('add')
        api = createStaff
        payload = { ...this.ruleForm, appType: '1' }
        this.$refs.ruleForm.validate(async isOk => {
          if (isOk) {
            this.submit(api, payload, true)
          }
        })
      } else if (this.dialogFlag === 'edit') {
        console.log('edit')
        api = editStaff
        payload = { ...this.ruleForm, appType: '1' }
        console.log('payload', payload)
        this.$refs.ruleForm.validate(async isOk => {
          if (isOk) {
            this.submit(api, payload, false)
          }
        })
      }
    },

    async submit(api, payload, isAddDialog) {
      try {
        await api(payload)
        this.$notify({
          message: `${isAddDialog ? '新建' : '修改'}成功`,
          type: 'success',
          duration: 2000,
          onClose: () => {
            this.getTableData()
            this.visible = false
            this.cancelBtn()
          }
        })
      } catch (error) {
        console.log('角色新建或编辑提交事件 error = ', error)
      }
    },

    cancelBtn() {
      this.visible = false
      this.$refs['ruleForm'].resetFields()
      this.ruleForm.name = ''
      this.ruleForm.phone = ''
      this.ruleForm.roleId = ''
    },

    async deleteBth(row) {
      console.log('id', row.id)
      try {
        const { code } = await deleteStaff({
          userId: row.id,
          roleId: row.roleId
        })
        if (code === '0' || code === 0) {
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success',
            duration: 2000
          })
          this.getTableData()
        }
      } catch (error) {
        console.log('删除员工 err', error)
      }
    },

    // 冻结按钮
    async blockBtn(row) {
      console.log(row)
      let codeNum = ''
      if (row.status === '1') {
        codeNum = '1'
        console.log('111', codeNum)
      } else {
        codeNum = '2'
      }
      try {
        const params = {
          userId: row.id,
          operateType: codeNum
        }

        const { code } = await operateAccount(params)
        if (code === '0' || code === 0) {
          this.$notify({
            title: '成功',
            message: '操作成功！',
            type: 'success',
            duration: 2000
          })
          this.getTableData()
        }
      } catch (error) {
        console.log('删除员工 err', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.user {
  height: 100%;
  background-color: #fff;
  position: relative;
}
</style>
