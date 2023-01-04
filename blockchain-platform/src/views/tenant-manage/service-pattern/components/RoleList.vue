<!--
 * @Description  : 角色列表，及新建、编辑、删除角色
 * @Autor        : Hemingway
 * @Date         : 2021-12-28 17:15:18
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-26 16:45:01
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\service-pattern\components\RoleList.vue
-->
<template>
  <div class="role-list">
    <div class="scroll-area">
      <!-- 条件选项及操作区域：如有需要，请使用action插槽 -->
      <i-action @on-query="getListData" @on-reset="onReset">
        <i-action-item label="服务模式">
          <el-input
            placeholder="请输入服务模式"
            v-model="searchKey"
            @keyup.enter.native="getListData"
            clearable
          >
          </el-input>
        </i-action-item>

        <template #action>
          <el-button type="primary" @click="onAdd" v-btn-permission:6701
            >新 建</el-button
          >
        </template>
      </i-action>

      <div class="list">
        <div
          :class="{ 'role-row': true, active: role.active }"
          v-for="role in listData"
          :key="role.id"
          @click="onChooseRole(role)"
        >
          <div class="complex-label">
            <el-tooltip
              effect="dark"
              content="该角色暂未启用"
              placement="top"
              v-if="role.enable === 'N'"
            >
              <span style="color: #c0c4cc">{{ role.name }}</span>
            </el-tooltip>
            <span v-else>{{ role.name }}</span>

            <span class="remark">{{
              `${role.isSys === 'Y' ? '（系统预设）' : ''}（${
                role.bondedTenantCount
              }人）`
            }}</span>
          </div>
          <div class="action">
            <el-button
              type="text"
              @click="onEdit(role)"
              v-if="role.isSys === 'N'"
              v-btn-permission:6702
              >编辑</el-button
            >
            <template v-if="role.isSys === 'N'">
              <el-tooltip
                effect="dark"
                :content="`${role.bondedTenantCount}个用户正在使用该角色，不能删除`"
                placement="top"
                v-if="role.bondedTenantCount > 0"
              >
                <div style="display: inline-block; margin-left: 10px">
                  <el-button type="text" disabled>删除</el-button>
                </div>
              </el-tooltip>
              <el-button
                v-else
                style="color: #f44336"
                type="text"
                @click="onDelete(role)"
                >删除</el-button
              >
            </template>
          </div>
        </div>
      </div>

      <!-- dialog -->
      <i-dialog :title="dialogTitle" :visible.sync="dialogFormVisible">
        <el-form :model="form">
          <el-form-item label="服务模式">
            <el-input
              v-model="form.name"
              autocomplete="off"
              placeholder="请输入服务模式"
            ></el-input>
          </el-form-item>
          <el-form-item label="使用系统">
            <el-radio-group v-model="form.appType">
              <el-radio
                label="2"
                :disabled="dialogFlag === 'edit'"
                v-if="form.appType === '2' || dialogFlag === 'add'"
                >溯源管理平台</el-radio
              >
              <el-radio
                label="3"
                :disabled="dialogFlag === 'edit'"
                v-if="form.appType === '3' || dialogFlag === 'add'"
                >加工数据采集小程序</el-radio
              >
            </el-radio-group>
          </el-form-item>
          <el-form-item label="限制范围">
            <el-radio-group v-model="form.type">
              <el-radio
                label="1"
                :disabled="dialogFlag === 'edit'"
                v-if="form.type === '1' || dialogFlag === 'add'"
                >限制使用时长</el-radio
              >
              <el-radio
                label="2"
                :disabled="dialogFlag === 'edit'"
                v-if="
                  (form.type === '2' || dialogFlag === 'add') &&
                  form.appType !== '3'
                "
                >限制使用码数</el-radio
              >
              <el-radio
                label="0"
                :disabled="dialogFlag === 'edit'"
                v-if="form.type === '0' || dialogFlag === 'add'"
                >不限制</el-radio
              >
            </el-radio-group>
          </el-form-item>
          <el-form-item label="启用状态">
            <el-radio-group v-model="form.enable">
              <el-radio label="Y">启用</el-radio>
              <el-radio label="N">不启用</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="备注">
            <el-input
              :maxlength="200"
              show-word-limit
              type="textarea"
              :autosize="{ minRows: 3 }"
              placeholder="请输入内容"
              v-model="form.remark"
            >
            </el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取 消</el-button>
          <el-button
            type="primary"
            @click="onSubmitRoleDialog"
            :loading="submitLoading"
            >确 定</el-button
          >
        </div>
      </i-dialog>
    </div>
  </div>
</template>

<script>
import { getRole, addRole, editRole, deleteRole } from '@/api/tenant-manage'

import listMixin from '@/views/mixins/listMixin'

export default {
  name: 'RoleList',

  mixins: [listMixin],

  data() {
    return {
      searchKey: '', // rolename

      // dialog
      dialogFlag: '', // add或者edit
      dialogTitle: '',
      dialogFormVisible: false,

      //   表单相关
      // form: {
      //   id: '',
      //   name: '',
      //   enable: '',
      //   remark: ''
      // },
      // form: {
      //   roleId: '',
      //   roleName: '',
      //   enable: '',
      //   remark: ''
      // },
      form: {
        // id: '', //角色id
        name: '',
        appType: '2',
        type: '1',
        enable: 'Y',
        remark: ''
      },
      submitLoading: false,
      listData: [
        {
          roleId: '1',
          roleName: '系统管理员',
          isSys: 'Y',
          bondedTenantCount: '10',
          enable: 'Y'
        },
        {
          roleId: '2',
          roleName: '湖南业务部管理员',
          isSys: 'N',
          bondedTenantCount: '20',
          enable: 'Y'
        }
      ]
    }
  },

  computed: {},

  watch: {
    dialogFlag(newV) {
      if (newV === 'add') {
        this.dialogTitle = '新建服务模式'
      } else if (newV === 'edit') {
        this.dialogTitle = '编辑服务模式'
      }
    }
  },

  mounted() {
    this.getListData()
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
     * @description: 查询一级组织下的角色信息
     * @param {Object} payload 预制请求参数对象
     * @author: Hemingway
     */
    async getListDataImpl(payload) {
      console.log('payload', payload)
      const { activeRoleId } = payload
      if (activeRoleId) {
        delete payload.activeRoleId
      }
      if (this.searchKey) {
        payload.roleName = this.searchKey
      }
      try {
        const { data } = await getRole(payload)
        console.log('data', data)
        return { listData: data }
      } catch (error) {
        console.log('查询角色列表 error = ', error)
      }
    },

    /**
     * @description: 新建角色
     * @author: Hemingway
     */
    onAdd() {
      // this.resetForm()
      this.dialogFlag = 'add'
      ;(this.form = {
        name: '',
        appType: '2',
        type: '1',
        enable: 'Y',
        remark: ''
      }),
        (this.dialogFormVisible = true)
    },

    /**
     * @description: 编辑角色
     * @param {Objecy} role
     * @author: Hemingway
     */
    onEdit(role) {
      // const { id, name, enable, remark } = role
      let _form = {
        id: role.id,
        name: role.name,
        appType: role.appType,
        type: role.type,
        enable: role.enable,
        remark: role.remark ? role.remark : ''
      }
      this.form = Object.assign({}, this.form, _form)
      console.log('编辑', this.form)
      this.dialogFlag = 'edit'
      this.dialogFormVisible = true
    },

    /**
     * @description: 删除角色
     * @param {Objecy} role
     * @author: Hemingway
     */
    onDelete(role) {
      this.$confirm(`此操作将永久删除角色：${role.name}，是否继续？`, '提示', {
        confirmButtonText: '确 定',
        cancelButtonText: '取 消',
        type: 'warning'
      })
        .then(async () => {
          try {
            await deleteRole({ id: role.id })
            this.$notify({
              message: '删除成功',
              type: 'success',
              duration: 2000,
              onClose: () => {
                this.getListData()
              }
            })
          } catch (error) {
            console.log('角色删除 error = ', error)
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
        name: '',
        enable: 'Y',
        remark: '',
        appType: '2'
      }
    },

    /**
     * @description: 角色新建或编辑提交事件
     * @author: Hemingway
     */
    onSubmitRoleDialog() {
      let api,
        payload = null

      if (this.dialogFlag === 'add') {
        api = addRole

        payload = { ...this.form }

        this.handleSubmitRoleDialog(api, payload, true)
      } else if (this.dialogFlag === 'edit') {
        api = editRole
        payload = { ...this.form }
        console.log('edit', payload)

        this.handleSubmitRoleDialog(api, payload, false)
      }
    },

    /**
     * @description: 角色新建或编辑提交事件——请求接口
     * @param {Object} api
     * @param {Object} payload
     * @param {Boolean} isAddDialog
     * @author: Hemingway
     */
    async handleSubmitRoleDialog(api, payload, isAddDialog) {
      console.log('payload', payload)
      this.submitLoading = true
      try {
        // const { id } = await api(payload)
        await api(payload)
        this.$notify({
          message: `${isAddDialog ? '新建' : '修改'}成功`,
          type: 'success',
          duration: 2000,
          onClose: () => {
            // this.getListData({ activeRoleId: id })
            this.getListData()
            this.dialogFormVisible = false
          }
        })

        this.submitLoading = false
      } catch (error) {
        console.log('角色新建或编辑提交事件 error = ', error)
        this.submitLoading = false
      }
    },

    /**
     * @description: 角色选择事件
     * @param {Object} role
     * @author: Hemingway
     */
    onChooseRole(role) {
      this.listData.forEach(item => {
        if (role.roleId !== item.id) {
          item.active = false
        } else {
          item.active = true
        }
      })

      this.$emit('onChooseRole', role)
    }
  }
}
</script>

<style lang="scss" scoped>
.role-list {
  height: calc(100vh - 204px);

  .scroll-area {
    height: 100%;
    padding: 0 16px;
    overflow-y: scroll;
    background-color: #fff;

    .list {
      background-color: #fff;
      font-size: 14px;

      .role-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 44px;
        padding: 0 16px;
        border-bottom: 1px solid #eee;

        .complex-label {
          .remark {
            color: #8b8b8b;
          }
        }
      }

      .role-row:hover {
        background: #e3f7cd;
      }

      .role-row.active {
        // background-color: #ebfbf1;
        background-color: red;
      }
    }
  }

  .scroll-area::-webkit-scrollbar {
    display: none;
  }
}
</style>
