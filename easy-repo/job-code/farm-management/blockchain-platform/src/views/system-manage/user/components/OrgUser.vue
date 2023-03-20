<!--
 * @Description  : 用户查询，及用户新建、删除、编辑（解锁）
 * @Autor        : Hemingway
 * @Date         : 2021-12-28 17:59:21
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-04-23 11:12:18
 * @FilePath     : \blockchain-admin\src\views\system-manage\user\components\OrgUser.vue
-->
<template>
  <div class="org-user">
    <div class="scroll-area">
      <!-- 条件选项及操作区域：如有需要，请使用action插槽 -->
      <i-action @on-query="getTableData" @on-reset="onReset" :loading="loading">
        <i-action-item label="关键词">
          <el-input
            placeholder="请输入手机号/姓名/所属组织/角色"
            v-model="searchKey"
            @keyup.enter.native="getTableData"
            clearable
          >
          </el-input>
        </i-action-item>

        <template #action>
          <el-button type="primary" @click="onAdd">新 建</el-button>
        </template>
      </i-action>

      <!-- 主体内容区域 -->
      <div class="body">
        <i-table
          :loading="loading"
          :data="tableData"
          @selection-change="onSelectionChange"
          ref="table"
        >
          <el-table-column type="selection" width="40"> </el-table-column>
          <el-table-column fixed prop="phone" label="手机号"> </el-table-column>
          <el-table-column fixed prop="name" label="姓名"> </el-table-column>
          <el-table-column
            prop="orgName"
            label="所属组织"
            show-overflow-tooltip
          >
          </el-table-column>
          <el-table-column prop="roleName" label="角色"> </el-table-column>
          <el-table-column prop="createDate" label="创建时间">
          </el-table-column>
          <el-table-column prop="creator" label="创建人"> </el-table-column>
          <el-table-column prop="isLocked" label="状态">
            <template v-slot="{ row }">
              {{ row.isLocked === 'Y' ? '已锁' : '正常' }}
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="150">
            <template slot-scope="scope">
              <el-button @click="onEdit(scope.row)" type="text">编辑</el-button>
              <el-button
                @click="onLockOrUnlock(scope.row)"
                type="text"
                v-if="scope.row.isLocked === 'Y'"
              >
                {{ scope.row.isLocked === 'Y' ? '解锁' : '锁定' }}</el-button
              >
              <el-button
                type="text"
                @click="onDelete(scope.row)"
                style="color: #f44336"
                v-if="$store.getters.username !== scope.row.phone"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </i-table>

        <!-- dialog -->
        <i-dialog
          :title="dialogTitle"
          :visible.sync="dialogVisible"
          @close="onCancel"
        >
          <el-form :model="form" :rules="rules" ref="form">
            <el-form-item label="手机号码" prop="phone" :error="error">
              <el-input
                :maxlength="11"
                v-model="form.phone"
                autocomplete="off"
                placeholder="请输入手机号码"
                @blur="onPhoneBlur"
              ></el-input>
            </el-form-item>
            <el-form-item label="姓名">
              <span v-if="isNameExist" class="context">{{
                userInfo.name
              }}</span>
              <el-input
                v-else
                :maxlength="10"
                v-model="form.name"
                autocomplete="off"
                placeholder="请输入姓名"
              />
            </el-form-item>
            <el-form-item label="所属组织">
              <span v-if="isOrgExist" class="context">{{ userInfo.org }}</span>
              <el-cascader
                v-else
                v-model="form.orgInfos"
                placeholder="请选择组织"
                :show-all-levels="false"
                :options="treeData"
                :props="{
                  checkStrictly: true,
                  value: 'orgInfo',
                  children: 'childList'
                }"
                filterable
              ></el-cascader>
            </el-form-item>
            <el-form-item label="角色">
              <span v-if="isRoleExist" class="context">
                {{ userInfo.role }}</span
              >
              <el-select v-else v-model="form.roleId" placeholder="请选择角色">
                <el-option
                  v-for="role in roles"
                  :key="role.id"
                  :label="role.name"
                  :value="role.id"
                >
                </el-option>
              </el-select>
            </el-form-item>
          </el-form>

          <div slot="footer" class="dialog-footer" v-if="!isHidden">
            <el-button @click="onCancel">取 消</el-button>
            <el-button
              type="primary"
              @click="onSubmitUserDialog"
              :disabled="isSubmitDisabled"
              :loading="submitLoading"
              >确 定</el-button
            >
          </div>
        </i-dialog>
      </div>
    </div>

    <!-- footer区域 -->
    <i-footer v-if="tableData.length > 0">
      <template #left>
        <i-footer-selected
          :count="selections.length"
          :isIndeterminate="isIndeterminate"
          @change="onChange"
          ref="selected"
        ></i-footer-selected>

        <!-- 禁用状态 -->
        <el-tooltip
          effect="dark"
          content="页面上暂无选择项"
          placement="top"
          v-if="selections.length === 0"
        >
          <div>
            <el-button type="text" disabled>删除</el-button>
          </div>
        </el-tooltip>
        <!-- 可操作状态 -->
        <el-button
          v-else
          style="color: #f44336"
          type="text"
          @click="onDelete(selections)"
          >删除</el-button
        >
      </template>
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
import { validPhoneNumber } from '@/utils/validate'
import {
  getUser,
  addUser,
  editUser,
  lockOrUnlockUser,
  deleteUser,
  getRole
} from '@/api/system-manage'
import { getUserInfo } from '@/api/user'
import { parseTime } from '@/utils'
import { getOrgInfo } from '@/api/system-manage'
import tableMixin from '@/views/mixins/tableMixin'
import handleTreeData from './mixin'

export default {
  name: 'OrgUser',

  mixins: [tableMixin],

  data() {
    return {
      orgId: '',
      searchKey: '', // 模糊匹配组织或用户

      // dialog
      dialogFlag: '', // add或者edit
      dialogTitle: '',
      dialogVisible: false,
      submitLoading: false,

      //   表单相关
      form: {
        id: '',
        name: '',
        phone: '',
        orgInfos: [], // orgInfo数组
        roleId: [] // roleId
      },
      rules: {
        phone: [{ trigger: 'blur', validator: this.validatePhoneNumber }]
      },
      error: '',

      roles: [], // role列表
      treeData: [], // 组织机构树

      // 用于展示已入库的用户信息
      userInfo: {
        name: '',
        org: '',
        role: ''
      },
      isNameExist: false,
      isOrgExist: false,
      isRoleExist: false,
      isHidden: false // 是否打断新建操作
    }
  },

  props: {
    org: {
      type: Object,
      default: () => {}
    }
  },

  computed: {
    // dialog的提交按钮是否禁用
    isSubmitDisabled() {
      let flag =
        !this.form.name ||
        !this.form.phone ||
        !(this.form.orgInfos.length > 0) ||
        !this.form.roleId ||
        !validPhoneNumber(this.form.phone)

      return flag
    },

    // treeData转list
    orgList() {
      const handleOrgTreeToList = treeData => {
        let res = []

        treeData.forEach(({ id, label, orgInfo, childList }) => {
          res.push({
            id,
            label,
            orgInfo
          })

          if (childList && childList.length > 0) {
            res = res.concat(handleOrgTreeToList(childList))
          }
        })

        return res
      }

      return handleOrgTreeToList(this.treeData)
    }
  },

  watch: {
    dialogFlag(newV) {
      if (newV === 'add') {
        this.resetForm()
        this.dialogTitle = '新建用户'
      } else if (newV === 'edit') {
        this.dialogTitle = '编辑用户'
      }
    },

    org: {
      handler(newV) {
        this.orgId = newV.id
        this.getTableData()
      },
      deep: true
    }
  },

  methods: {
    /**
     * @description: 重置（查询）的具体实现
     * @author: Hemingway
     */
    onResetImpl() {
      this.searchKey = ''
    },

    /**
     * @description: 列表查询逻辑的具体实现
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getTableDataImpl(payload) {
      payload.orgId = this.orgId
      if (this.searchKey) {
        payload.keywords = this.searchKey
      }

      try {
        const { list, total } = await getUser(payload)
        const tableData = (list || []).map(
          ({
            id,
            userName,
            telephone,
            organizationName,
            orgInfos,
            roleInfos,
            isLocked,
            createDate,
            creator
          }) => ({
            id,
            name: userName,
            phone: telephone,
            orgInfos, // 用于编辑回显
            roleInfos, // 用于编辑回显
            orgName: organizationName, // 用于列表展示
            roleName: roleInfos.map(role => role.name).join(', '), // 用于列表展示
            isLocked,
            createDate: parseTime(createDate, '{y}-{m}-{d}'),
            creator
          })
        )

        return { tableData, total }
      } catch (error) {
        console.log('账号查询 error = ', error)
      }
    },

    /**
     * @description: 手机号校验逻辑
     * @param {Object} rule
     * @param {String} value
     * @param {Function} callback
     * @author: Hemingway
     */
    validatePhoneNumber(rule, value, callback) {
      if (!validPhoneNumber(value)) {
        callback(new Error('请输入正确的手机号'))
      } else if (this.error) {
        callback(new Error(this.error)) // rule规则之外的手动干预
      } else {
        callback()
      }
    },

    /**
     * @description: 手机号失焦或键盘enter事件
     * @param {Object} e
     * @author: Hemingway
     */
    async onPhoneBlur(e) {
      if (this.dialogFlag === 'add') {
        // 新建用户时检测手机号
        e.preventDefault()

        this.error = ''

        if (validPhoneNumber(this.form.phone)) {
          try {
            const { userName, organRoles } = await getUserInfo({
              phone: this.form.phone
            })

            if (userName) {
              this.form.name = userName
              this.userInfo.name = userName
              this.isNameExist = true
            } else {
              this.form.name = ''
              this.userInfo.name = ''
              this.isNameExist = false
            }

            if (organRoles && organRoles.length > 0) {
              const orgRole = organRoles.find(
                item => item.tenantId === this.$store.getters.orgId
              )

              if (orgRole) {
                // 已在当前组织中存在，仅需处理展示

                this.error = '该手机号的用户已存在，请确认'
                this.isHidden = true
                this.userInfo.org = orgRole.organizationName
                this.isOrgExist = true
                this.userInfo.role = orgRole.roleName
                this.isRoleExist = true
              } else {
                // 不在当前组织中存在

                this.resetOrgRole()
              }
            } else {
              // 不在当前组织中存在

              this.resetOrgRole()
            }
          } catch (error) {
            console.log('查询用户信息 error = ', error)
          }
        }
      }
    },

    /**
     * @description: 重置dialog
     * @author: Hemingway
     */
    resetDialog() {
      this.form.name = ''
      this.userInfo.name = ''
      this.isNameExist = false
      this.resetOrgRole()
    },

    /**
     * @description: 重置组织及角色状态
     * @author: Hemingway
     */
    resetOrgRole() {
      this.isHidden = false
      this.form.orgInfos = []
      this.userInfo.org = ''
      this.isOrgExist = false
      this.form.roleId = ''
      this.userInfo.role = ''
      this.isRoleExist = false
    },

    /**
     * @description: 新建用户
     * @author: Hemingway
     */
    async onAdd() {
      this.dialogFlag = 'add'
      this.dialogVisible = true

      if (this.roles.length === 0) {
        await this.getRoleInfo()
      }
      if (this.treeData.length === 0) {
        await this.getOrgInfo()
      }
    },

    /**
     * @description: 编辑用户
     * @param {Object} user
     * @author: Hemingway
     */
    async onEdit(user) {
      this.dialogFlag = 'edit'
      this.dialogVisible = true

      if (this.roles.length === 0) {
        await this.getRoleInfo()
      }
      if (this.treeData.length === 0) {
        await this.getOrgInfo()
      }

      const { id, name, phone, orgInfos, roleInfos } = user
      this.form.id = id
      this.form.name = name
      this.form.phone = phone

      this.form.orgInfos =
        orgInfos &&
        orgInfos.map(({ id }) => {
          const node = this.orgList.find(org => org.id === id) // 找到
          return node.orgInfo
        }) // orgInfo数组
      this.form.roleId = roleInfos && roleInfos[0].id // roleId
    },

    /**
     * @description: 查询一级组织下的角色信息
     * @author: Hemingway
     */
    async getRoleInfo() {
      try {
        const { roleInfoList } = await getRole()
        if (roleInfoList && roleInfoList.length > 0) {
          this.roles = roleInfoList
        } else {
          this.roles = []
        }
      } catch (error) {
        console.log('查询角色列表 error = ', error)
      }
    },

    /**
     * @description: 查询组织机构
     * @author: Hemingway
     */
    async getOrgInfo() {
      try {
        const { orgInfoList } = await getOrgInfo()
        handleTreeData(orgInfoList)
        this.treeData = [orgInfoList]
      } catch (error) {
        console.log('查询组织机构 error = ', error)
      }
    },

    /**
     * @description: 锁定或解锁用户
     * @param {Object} user
     * @author: Hemingway
     */
    onLockOrUnlock(user) {
      this.$confirm(
        `此操作将${user.isLocked === 'Y' ? '解锁' : '锁定'}用户：${
          user.name
        }，是否继续？`,
        '提示',
        {
          confirmButtonText: '确 定',
          cancelButtonText: '取 消',
          type: 'warning'
        }
      )
        .then(async () => {
          try {
            await lockOrUnlockUser({
              userIds: user.id,
              isLocked: user.isLocked === 'Y' ? 'N' : 'Y'
            })
            this.$message({
              message: `${user.isLocked === 'Y' ? '解锁' : '锁定'}成功`,
              type: 'success',
              duration: 2000,
              onClose: () => {
                user.isLocked = user.isLocked === 'Y' ? 'N' : 'Y'
              }
            })
          } catch (error) {
            console.log(
              `用户${user.isLocked === 'Y' ? '解锁' : '锁定'} error = `,
              error
            )
          }
        })
        .catch(() => {})
    },

    /**
     * @description: 删除用户
     * @param {Object | Array} user
     * @author: Hemingway
     */
    onDelete(user) {
      let users = []
      if (!(user instanceof Array)) {
        users.push(user)
      } else {
        users = user
      }

      this.$confirm(
        `此操作将永久删除用户：${users
          .map(user => user.name)
          .join(', ')}，是否继续？`,
        '提示',
        {
          confirmButtonText: '确 定',
          cancelButtonText: '取 消',
          type: 'warning'
        }
      )
        .then(async () => {
          try {
            await deleteUser({ userId: users.map(user => user.id).join(',') })
            this.$message({
              message: '删除成功',
              type: 'success',
              duration: 2000,
              onClose: () => {
                this.getTableData()
                this.$parent.$parent.$parent.$parent.$refs.orgTree.getOrgInfo()
              }
            })
          } catch (error) {
            console.log('用户删除 error = ', error)
          }
        })
        .catch(() => {})
    },

    /**
     * @description: 重置form
     * @author: Hemingway
     * @example:
     */
    resetForm() {
      this.form = {
        id: '',
        name: '',
        phone: '',
        orgInfos: [],
        roleId: ''
      }
    },

    /**
     * @description: 用户新建或编辑提交事件
     * @author: Hemingway
     */
    onSubmitUserDialog() {
      let api,
        payload = null
      const { id, name, phone, orgInfos, roleId } = this.form

      if (this.dialogFlag === 'add') {
        api = addUser

        payload = {
          userName: name,
          phone,
          orgInfos: orgInfos.map(({ id, type }) => ({ id, type })),
          roleInfos: [roleId].map(id => ({ id }))
        }

        this.handleSubmitUserDialog(api, payload, true)
      } else if (this.dialogFlag === 'edit') {
        api = editUser
        payload = {
          userId: id,
          userName: name,
          phone,
          orgInfos: orgInfos.map(({ id, type }) => ({ id, type })),
          roleInfos: [roleId].map(id => ({ id }))
        }

        this.handleSubmitUserDialog(api, payload, false)
      }
    },

    /**
     * @description: 用户新建或编辑提交事件——请求接口
     * @param {Object} api
     * @param {Object} payload
     * @param {Boolean} isAddDialog
     * @author: Hemingway
     */
    handleSubmitUserDialog(api, payload, isAddDialog) {
      this.$refs.form.validate(async valid => {
        if (valid) {
          this.submitLoading = true
          try {
            await api(payload)

            this.$message({
              message: `${isAddDialog ? '新建' : '修改'}成功`,
              type: 'success',
              duration: 2000,
              onClose: () => {
                this.dialogVisible = false
                this.getTableData()
                this.$parent.$parent.$parent.$parent.$refs.orgTree.getOrgInfo()
              }
            })
            this.submitLoading = false
          } catch (error) {
            console.log('用户新建或提交事件 error = ', error)
            this.submitLoading = false
          }
        }
      })
    },

    onCancel() {
      this.dialogVisible = false
      this.$refs.form.resetFields()
      this.resetDialog()
    }
  }
}
</script>

<style lang="scss" scoped>
.org-user {
  position: relative;

  .scroll-area {
    height: calc(100% - 40px);
    padding-bottom: 20px;
    overflow-y: scroll;

    .body {
      padding-bottom: 56px; // 预留 i-footer 位置

      .context {
        padding-left: 15px;
      }
    }
  }

  .scroll-area::-webkit-scrollbar {
    display: none;
  }

  ::v-deep .i-footer {
    margin: 0 -16px;
    bottom: -16px;
  }
}
</style>
