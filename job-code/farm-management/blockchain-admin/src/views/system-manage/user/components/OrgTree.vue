<!--
 * @Description  : 组织结构树查询
 * @Autor        : Hemingway
 * @Date         : 2021-12-28 17:15:18
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-17 08:24:03
 * @FilePath     : \blockchain-admin\src\views\system-manage\user\components\OrgTree.vue
-->
<template>
  <div class="org-tree">
    <div class="tree-wrapper">
      <i-section-header>组织架构</i-section-header>

      <el-tree
        accordion
        :data="treeData"
        :default-expand-all="true"
        :expand-on-click-node="true"
        ref="tree"
        highlight-current
        :props="defaultProps"
        @node-click="onChooseOrg"
        v-loading="loading"
      >
        <div class="custom-tree-node" slot-scope="{ data }">
          <span>{{ data.label }} </span>
          <span class="remark">{{ `（${data.orgInfo.userCount}人）` }} </span>
        </div>
      </el-tree>
    </div>
  </div>
</template>

<script>
import { getOrg } from '@/api/system-manage'
import handleTreeData from './mixin'

export default {
  name: 'OrgTree',

  data() {
    return {
      loading: false,
      treeData: [], // 组织结构树
      defaultProps: {
        label: 'label', // 节点标签
        children: 'childList' // 子树
      }
    }
  },

  async mounted() {
    await this.getOrgInfo()
    this.onChooseOrg(this.treeData[0])
  },

  methods: {
    /**
     * @description: 查询组织架构树
     * @author: Hemingway
     */
    async getOrgInfo() {
      this.loading = true
      try {
        const { orgInfoList } = await getOrg()
        handleTreeData(orgInfoList)
        this.treeData = [orgInfoList]
        this.loading = false
      } catch (error) {
        console.log('查询组织架构树 error = ', error)
        this.loading = false
      }
    },

    /**
     * @description: 组织选择事件
     * @param {Object} node
     * @author: Hemingway
     */
    onChooseOrg(node) {
      this.$emit('onChooseOrg', node.orgInfo)
    }
  }
}
</script>

<style lang="scss" scoped>
.org-tree {
  height: calc(100vh - 204px);
  border: 1px solid #eee;

  .tree-wrapper {
    height: 100%;
    padding: 0 16px;

    background-color: #fff;
    font-size: 14px;
    position: relative;

    ::v-deep {
      .el-tree {
        height: calc(100% - 56px);
        padding-bottom: 16px;
        overflow-y: auto;

        .el-tree-node__content {
          height: 44px;
          border-bottom: 1px solid #eee;

          .custom-tree-node {
            color: #212121;

            .remark {
              color: #8b8b8b;
            }
          }
        }
      }

      .el-tree::-webkit-scrollbar {
        display: none;
      }
    }
  }
}
</style>
