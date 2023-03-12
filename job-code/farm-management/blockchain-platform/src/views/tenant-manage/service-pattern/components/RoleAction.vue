<!--
 * @Description  : 角色人员查看，及角色权限管理
 * @Autor        : Hemingway
 * @Date         : 2021-12-28 17:59:21
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 14:20:34
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\service-pattern\components\RoleAction.vue
-->
<template>
  <div class="role-action">
    <i-section-header>服务模式权限及服务企业</i-section-header>
    <el-tabs
      type="card"
      v-model="activeName"
      @tab-click="handleTabClick"
      v-if="role.id"
      class="inner-tab"
    >
      <el-tab-pane label="服务权限" name="first" class="right-setting">
        <section class="section">
          <div class="title">
            <!-- <svg-icon class-name="section-icon" icon-class="chart" />数据权限 -->
            数据权限
          </div>
          <el-cascader
            :options="areaList"
            :props="{
              label: 'name',
              value: 'code',
              checkStrictly: 'true'
            }"
            collapse-tags
            clearable
            style="width: 360px"
            v-model="regionPower"
            @change="changeOver"
            ref="cascader"
          ></el-cascader>
          <!-- <el-popover placement="top" v-model="visible">
            <div slot="reference" class="dropBtn">
              <span v-if="codeObj.dataPermissionAddress">{{
                codeObj.dataPermissionAddress
              }}</span>
              <span v-else>请选择查看允许的地区</span>
              <i
                :class="visible ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"
              ></i>
            </div>

            <div class="popoverBox">
              <linkageSelect
                v-model="visible"
                :codeObj="codeObj"
                @changeOver="changeOver"
              ></linkageSelect>
            </div>
          </el-popover> -->
        </section>
        <section class="section">
          <div class="title">菜单与权限</div>
          <el-tree
            ref="tree"
            :data="treeData"
            :render-after-expand="false"
            :default-expand-all="false"
            :default-expanded-keys="['-1', '-2', '-3']"
            :expand-on-click-node="true"
            node-key="id"
            show-checkbox
            :props="defaultProps"
            @check-change="onCheckedChange"
            v-loading="loading"
          >
          </el-tree>
        </section>
      </el-tab-pane>
      <el-tab-pane label="服务企业" name="second" class="role-person">
        <el-table
          :data="tableData"
          v-loading="loading2"
          :header-cell-style="{ backgroundColor: '#fafafa' }"
          :cell-style="{ padding: '7px 0' }"
        >
          <el-table-column prop="defaultTenantName" label="企业名称">
          </el-table-column>
          <el-table-column prop="name" label="管理员"> </el-table-column>
          <el-table-column prop="phone" label="手机号"> </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
    <span class="empty" v-else>选择左侧角色，以设置角色权限和查看角色人员</span>

    <i-footer v-show="showSubmitBtn && activeName === 'first'">
      <el-button
        type="primary"
        @click="onSubmitRoleRights"
        :loading="submitLoading"
        v-show="role.isSys === 'N'"
        >提交</el-button
      >
    </i-footer>

    <i-footer v-if="activeName === 'second' && tableData.length > 0">
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
  </div>
</template>

<script>
import { getRoleRight, getRoleUser, submitRoleRight } from '@/api/tenant-manage'
import areaList from './select/linkageSelect/area.js'
// import linkageSelect from './select/linkageSelect/index.vue'
// import { array } from 'jszip/lib/support'

export default {
  name: 'RoleAction',
  // components: {
  //   linkageSelect
  // },
  data() {
    return {
      areaList,
      regionPower: [],
      ifIsCheck: false,
      loading: false,
      loading2: false,
      activeName: 'first', // first 权限设置；second 角色人员
      showSubmitBtn: false,
      page: {
        pageSize: 10,
        pageNum: 1,
        total: 0
      },
      tableData: [],

      treeData: [], // 权限列表
      defaultProps: {
        label: 'name', // 节点标签
        children: 'children' // 子树
      },

      submitLoading: false,
      visible: false,
      codeObj: {
        dataPermissionProvince: '',
        dataPermissionCity: '',
        dataPermissionArea: '',
        dataPermissionAddress: ''
      },
      //默认勾选数组
      DEFAULT_CHECKED_KEYS: [],
      // 菜单和功能权限 _privileges
      privileges: []
    }
  },

  props: {
    role: {
      type: Object,
      default: () => {}
    }
  },

  watch: {
    role: {
      async handler(newRole, oldRole) {
        if (newRole.id !== oldRole.id) {
          this.loading = true
          await this.handlerChooseRole(newRole)
          this.loading = false
        }
      },
      deep: true
    }
  },

  created() {
    let newArr = JSON.parse(JSON.stringify(areaList))
    newArr.unshift({
      name: '全国',
      code: 'all'
    })
    this.areaList = newArr
  },

  methods: {
    changeOver() {
      let nodesObj = this.$refs['cascader'].getCheckedNodes()[0].pathNodes
      this.codeObj = {
        dataPermissionProvince:
          nodesObj[0]?.data?.code === 'all'
            ? ''
            : nodesObj[0]?.data?.code
            ? nodesObj[0]?.data?.code
            : '',
        dataPermissionCity: nodesObj[1]?.data?.code || '',
        dataPermissionArea: nodesObj[2]?.data?.code || '',
        dataPermissionAddress: `${nodesObj[0]?.data?.name}${
          nodesObj[1]?.data?.name || ''
        }${nodesObj[2]?.data?.name || ''}`
      }
    },
    // 当前页码改变事件
    handleCurrentChange(val) {
      console.log('handleCurrentChange', val)
      this.page.pageNum = val
      this.getRoleUser(this.role, this.page.pageSize, this.page.pageNum)
    },

    // 当前条目数改变事件
    handleSizeChange(val) {
      console.log('handleSizeChange', val)
      this.page.pageSize = val
      this.getRoleUser(this.role, this.page.pageSize, this.page.pageNum)
    },

    /**
     * @description: 角色选择处理逻辑
     * @param {Object} newRole
     * @return {Promise}
     * @author: Hemingway
     */
    async handlerChooseRole(newRole) {
      this.treeData = []
      this.DEFAULT_CHECKED_KEYS = []
      this.activeName = 'first' // 角色切换，右侧容器默认切到权限设置视角
      await this.getRoleRight(newRole)
    },

    /**
     * @description: 查询指定角色权限树
     * @param {Object} role
     * @author: Hemingway
     */
    async getRoleRight(role) {
      try {
        const { data } = await getRoleRight({
          roleId: role.relatedRoleId,
          appType: role.appType
          // roleId: role.id
        })
        let echoArr = []
        this.codeObj.dataPermissionAddress = data.dataPermissionAddress
        this.codeObj.dataPermissionProvince = data.dataPermissionProvince
        this.codeObj.dataPermissionCity = data.dataPermissionCity
        this.codeObj.dataPermissionArea = data.dataPermissionArea
        if (data.dataPermissionProvince) {
          echoArr.push(data.dataPermissionProvince)
        }
        if (data.dataPermissionCity) {
          echoArr.push(data.dataPermissionCity)
        }
        if (data.dataPermissionArea) {
          echoArr.push(data.dataPermissionArea)
        }
        if (echoArr.length === 0) {
          echoArr.push('all')
        }
        this.regionPower = echoArr
        console.log('回显', this.regionPower)
        if (Array.isArray(data.menuTree) && data.menuTree.length > 0) {
          // 显示提交按钮
          this.showSubmitBtn = true
          // 创建 菜单及功能权限树
          this.treeData = []
          // 格式化 菜单和功能树
          this.treeFormat(data.menuTree)
          // 菜单及功能权限树 赋值
          this.treeData = data.menuTree
          // 勾选状态回显
          this.DEFAULT_CHECKED_KEYS = []
          this.defaultChecked(data.menuTree)
          this.$refs.tree.setCheckedKeys([])
          this.$nextTick(() => {
            this.DEFAULT_CHECKED_KEYS.forEach(e => {
              this.$refs.tree.setChecked(e, true, false)
            })
          })
        }
      } catch (error) {
        console.log('获取角色权限 error = ', error)
      }
    },
    // 格式化菜单和功能树
    treeFormat(arr) {
      if (arr && Array.isArray(arr) && arr.length > 0) {
        arr.map(ele => {
          if (
            ele.children &&
            Array.isArray(ele.children) &&
            ele.children.length > 0 &&
            ['', null, undefined].includes(ele.functionDtoList)
          ) {
            ele._fun = false
          } else if (
            ele.functionDtoList &&
            Array.isArray(ele.functionDtoList) &&
            ele.functionDtoList.length > 0
          ) {
            ele.functionDtoList.map(el => {
              el._fun = true
            })
            ele.children = ele.functionDtoList
            ele.functionDtoList = null
            ele._fun = false
          }
          this.treeFormat(ele.children)
        })
      }
    },
    // 获取默认勾选数组
    defaultChecked(data) {
      if (data && Array.isArray(data) && data.length > 0) {
        data.map(el => {
          if (this.role.isSys === 'Y') {
            el.disabled = true
          }
          if (el.isCheck === 'Y') {
            this.DEFAULT_CHECKED_KEYS.push(el.id)
          }
          el.children && this.defaultChecked(el.children)
          return el
        })
      }
    },
    //节点选中状态更改
    onCheckedChange(node, isCheck) {
      node.isCheck = isCheck ? 'Y' : 'N'
    },
    // 获取 菜单和权限 勾选对应列表
    selectedMenuPermissions() {
      let _arr = []
      // 获取半选状态的节点数组getHalfCheckedNodes()
      // 获取全选状态的节点数组getCheckedNodes()
      let _checkedArr = this.$refs.tree
        .getHalfCheckedNodes()
        .concat(this.$refs.tree.getCheckedNodes())
      _checkedArr.map(el => {
        if (!el._fun) {
          _arr.push({
            menuId: el.id,
            functionId: null,
            isCheck: 'Y'
          })
        } else {
          _arr.push({
            menuId: el.menuId,
            functionId: el.id,
            isCheck: 'Y'
          })
        }
      })
      // 返回 获取的菜单和权限 勾选对应列表
      return _arr
    },
    /**
     * @description: tab切换事件
     * @author: Hemingway
     */
    async handleTabClick(tab) {
      console.log('tab', tab)
      if (tab.index === '1') {
        this.activeName = 'second'
        this.getRoleUser(this.role, this.page.pageSize, this.page.pageNum)
      } else {
        this.activeName = 'first'
      }
    },

    /**
     * @description: 获取指定角色下的用户
     * @param {Object} role
     * @author: Hemingway
     */
    async getRoleUser(role, pageSize, pageNum) {
      if (role.bondedTenantCount > 0) {
        this.loading2 = true
        try {
          const res = await getRoleUser({
            roleId: role.relatedRoleId,
            pageSize: pageSize,
            pageNum: pageNum
          })

          this.tableData = res.list.map(el => {
            if (['', undefined, null].includes(el.name)) {
              el.name = `后台返显字段值: ${el.name}`
            }
            return el
          })
          this.page.total = res.total

          this.loading2 = false
        } catch (error) {
          console.log('获取角色用户 error = ', error)
          this.loading2 = false
        }
      } else {
        this.tableData = []
      }
    },

    /**
     * @description: 分配角色权限，提交事件
     * @author: Hemingway
     */
    async onSubmitRoleRights() {
      // 获取 菜单和权限 功能对应列表
      this.privileges = this.selectedMenuPermissions()
      // 创建提交数据
      const params = {
        roleId: this.role.relatedRoleId,
        dataPermissionAddress: this.codeObj.dataPermissionAddress,
        dataPermissionProvince: this.codeObj.dataPermissionProvince,
        dataPermissionCity: this.codeObj.dataPermissionCity,
        dataPermissionArea: this.codeObj.dataPermissionArea,
        privileges: this.privileges
      }
      console.log('this.privileges', this.privileges)
      if (Array.isArray(this.privileges) && this.privileges.length > 0) {
        try {
          this.submitLoading = true
          await submitRoleRight(params)
          this.$notify({
            message: '设置成功',
            type: 'success',
            duration: 2000,
            onClose: () => {
              this.getRoleRight(this.role)
            }
          })
          this.submitLoading = false
        } catch (error) {
          this.submitLoading = false
          console.log('设置角色权限 error = ', error)
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.role-action {
  height: 100%;
  padding: 0 16px;
  border: 1px solid #eee;
  position: relative;

  .inner-tab {
    height: calc(100% - 56px);

    ::v-deep .el-tabs__content {
      max-height: calc(100vh - 381px);
    }
  }

  // 权限设置
  .right-setting {
    .section {
      .title {
        height: 36px;
        line-height: 36px;
        color: #8b8b8b;

        .section-icon {
          margin-right: 12px;
        }
      }
    }

    // 菜单及权限设置
    .section:last-of-type {
      ::v-deep .el-tree {
        .el-tree-node__content {
          height: auto;
          min-height: 40px;
          border-bottom: 1px solid #eee;

          .custom-tree-node {
            .menu-name {
              color: #212121;
            }
            .menu-check {
              margin-left: 16px;
            }

            .func-wrapper {
              display: flex;
              align-items: center;
              flex-wrap: wrap;

              .el-checkbox {
                height: 32px;
                line-height: 32px;
              }
            }

            .el-tree-node__content {
              border: none;
            }
          }
        }
      }
    }
  }

  .empty {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    color: #909399;
  }

  ::v-deep .i-footer {
    left: 16px;
    right: 16px;
  }
}
.dropBtn {
  font-size: 13px;
  width: 300px;
  padding: 10px 10px 10px 20px;
  border-radius: 5px;
  border: 1px solid #ccc;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.popoverBox {
  width: 500px;
}
</style>
