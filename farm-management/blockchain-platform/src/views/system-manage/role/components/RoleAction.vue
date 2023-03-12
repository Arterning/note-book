<!--
 * @Description  : 角色人员查看，及角色权限管理
 * @Autor        : Hemingway
 * @Date         : 2021-12-28 17:59:21
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 14:16:25
 * @FilePath     : \blockchain-platform\src\views\system-manage\role\components\RoleAction.vue
-->
<template>
  <div class="role-action">
    <i-section-header>角色权限及角色人员</i-section-header>
    <el-tabs
      type="card"
      v-model="activeName"
      @tab-click="handleTabClick"
      v-if="role.id"
      class="inner-tab"
    >
      <el-tab-pane label="角色权限" name="first" class="right-setting">
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
            :data="treeData"
            :default-expand-all="false"
            :default-expanded-keys="['-1', '-2']"
            node-key="id"
            :expand-on-click-node="true"
            ref="tree"
            show-checkbox
            :props="defaultProps"
            @check-change="onCheckedChange"
            v-loading="loading"
          >
          </el-tree>
        </section>
      </el-tab-pane>
      <el-tab-pane label="角色人员" name="second" class="role-person">
        <el-table
          :data="tableData"
          v-loading="loading2"
          :header-cell-style="{ backgroundColor: '#fafafa' }"
          :cell-style="{ padding: '7px 0' }"
        >
          <el-table-column prop="name" label="姓名"> </el-table-column>
          <el-table-column prop="phone" label="手机号"> </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
    <span class="empty" v-else>选择左侧角色，以设置角色权限和查看角色人员</span>

    <i-footer v-show="showSubmitBtn && activeName === 'first'">
      <el-button
        v-if="!isSys"
        type="primary"
        @click="onSubmitRoleRights"
        :loading="submitLoading"
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
import { getRoleRight, getRoleUser, submitRoleRight } from '@/api/system-manage'
import areaList from './select/linkageSelect/area.js'
// import linkageSelect from './select/linkageSelect/index.vue'

export default {
  name: 'RoleAction',
  // components: {
  //   linkageSelect
  // },
  data() {
    return {
      areaList,
      regionPower: [],
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
      privileges: [],
      // 是否预设
      isSys: false
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
          this.isSys = newRole.isSys === 'Y' ? true : false
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
          roleId: role.id,
          appType: role.appType
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
        if (Array.isArray(data.menuTree) && data.menuTree.length > 0) {
          // 子节点functionDtoList同步到children
          this.handleArr(data.menuTree)
          // 显示提交按钮
          this.showSubmitBtn = true
          // 菜单及功能权限树
          this.treeData = data.menuTree
          // 默认勾选数组
          this.DEFAULT_CHECKED_KEYS = []
          this.getIsChecked(data.menuTree)
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
    /**
     * @description: 最后一层数据同步到children
     * @param {Arr} nodes 需要处理的树形结构
     * @author: qinjj
     */
    handleArr(nodes = []) {
      for (let item of nodes) {
        if (item.children && item.children?.length) {
          this.handleArr(item.children)
        } else if (item.functionDtoList && item.functionDtoList?.length) {
          item.children = item.functionDtoList
        }
      }
      return nodes
    },
    // 获取默认勾选数组
    getIsChecked(data) {
      if (Array.isArray(data) && data.length > 0) {
        data.map(el => {
          if (this.role.isSys === 'Y') {
            el.disabled = true
          }
          if (el.isCheck === 'Y') {
            this.DEFAULT_CHECKED_KEYS.push(el.id)
          }
          if (Array.isArray(el.children) && el.children.length > 0) {
            this.getIsChecked(el.children)
          }
          return el
        })
      }
    },
    //节点选中状态更改
    onCheckedChange(node, isCheck) {
      node.isCheck = isCheck ? 'Y' : 'N'
    },
    //获取所有功能权限的选中状态|表单提交前调用
    getCheckedRights(data) {
      if (Array.isArray(data) && data.length > 0) {
        data.map(el => {
          let _el = {
            menuId: el.id,
            isCheck: el.isCheck,
            functionId: null,
            name: el.name
          }
          // 判断子菜单勾选是否大于1, 大于1显示本级菜单,否则不显示本级菜单
          if (Array.isArray(el.children) && el.children.length > 0) {
            let _count = 0
            el.children.map(n => {
              if (n.isCheck === 'Y') {
                _count++
              }
            })
            if (_count > 1) {
              _el = Object.assign({}, _el, { isCheck: 'Y' })
            } else {
              _el = Object.assign({}, _el, { isCheck: 'N' })
            }
          }
          // 判断本级菜单是否有功能, 如果有功能,则按功能是否显示决定本级菜单是否显示
          if (
            Array.isArray(el.functionDtoList) &&
            el.functionDtoList.length > 0
          ) {
            el.functionDtoList.map(e => {
              let _e = {
                menuId: el.id,
                isCheck: e.isCheck,
                functionId: e.id,
                name: e.name
              }
              _el = Object.assign({}, _el, _e)
              this.privileges.push(_el)
            })
          } else {
            this.privileges.push(_el)
          }
          if (Array.isArray(el.children) && el.children.length > 0) {
            this.getCheckedRights(el.children)
          }
        })
      }
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
      if (role.userCount > 0) {
        this.loading2 = true
        try {
          const res = await getRoleUser({
            roleId: role.id,
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
     * @description: 处理提交的权限数据
     * @author: qinjj
     */
    handleSubmitTree() {
      let resultArr = []
      let checkArr = this.$refs.tree.getCheckedNodes() //返回选中节点
      let halfCheckArr = this.$refs.tree.getHalfCheckedNodes() // 返回半选中节点
      ;[...checkArr, ...halfCheckArr].forEach(res => {
        if (res.children || res.functionDtoList) {
          resultArr.push({
            menuId: res.id,
            functionId: '',
            isCheck: 'Y'
          })
        } else {
          resultArr.push({
            menuId: res.menuId,
            functionId: res.id,
            isCheck: 'Y'
          })
        }
      })
      return resultArr
    },

    /**
     * @description: 分配角色权限，提交事件
     * @author: Hemingway
     */
    async onSubmitRoleRights() {
      this.privileges = this.handleSubmitTree()
      const params = {
        roleId: this.role.id,
        dataPermissionAddress: this.codeObj.dataPermissionAddress,
        dataPermissionProvince: this.codeObj.dataPermissionProvince,
        dataPermissionCity: this.codeObj.dataPermissionCity,
        dataPermissionArea: this.codeObj.dataPermissionArea,
        privileges: this.privileges
      }
      console.log('params', params)
      if (Array.isArray(this.privileges) && this.privileges.length > 0) {
        try {
          // await setRoleRight(payload)
          this.submitLoading = true
          await submitRoleRight(params)
          this.$notify({
            message: '设置成功',
            type: 'success',
            duration: 2000,
            onClose: () => {}
          })
          this.submitLoading = false
        } catch (error) {
          this.submitLoading = false
          console.log('设置角色权限 error = ')
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
