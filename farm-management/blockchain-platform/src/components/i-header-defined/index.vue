<!--
 * @Description  : 表头自定义
 * @Autor        : WuJing
 * @Date         : 2022-03-24 11:03:09
 * @LastEditors  : WuJing
 * @LastEditTime : 2022-04-28 15:58:37
 * @FilePath     : \i-farm-platform\src\components\zl\i-header-defined\index.vue
-->
<template>
  <el-popover
    placement="bottom"
    ref="popverRef"
    trigger="click"
    @show="showPopver"
  >
    <div style="height: 510px">
      <el-input
        placeholder="搜索字段"
        v-model="filterText"
        prefix-icon="el-icon-search"
        clearable
      >
      </el-input>
      <div class="tree-style">
        <el-tree
          :data="data"
          show-checkbox
          node-key="id"
          default-expand-all
          :default-checked-keys="chooseKeys"
          :filter-node-method="filterNode"
          ref="tree"
        >
        </el-tree>
      </div>
      <div class="bottom">
        <el-button @click="$refs.popverRef.doClose()">取 消</el-button>
        <el-button type="primary" @click="comfimSave">保 存</el-button>
      </div>
    </div>
    <el-button slot="reference" @click.native="chooseTitle">表头设置</el-button>
  </el-popover>
</template>

<script>
import { updateHeader } from '@/api/soil-formula/define-table'
export default {
  name: 'IHeaderDefined',
  data() {
    return {
      //过滤的文字
      filterText: '',
      //选中的key
      chooseKeys: [],
      //选中的label
      chooseLabels: [],
      //表头
      header: []
    }
  },

  props: {
    //传树形数据
    data: {
      type: Array,
      default: () => []
    },
    //传树形默认选中的数据key
    checkedKey: {
      type: Array,
      default: () => []
    },
    //传树形默认选中的数据值
    checkedLabel: {
      type: Array,
      default: () => []
    },
    //页面id
    pageId: {
      type: String,
      default: ''
    },
    //表头数组
    tableHeader: {
      type: Array,
      default: () => []
    },
    //表的数据
    tableList: {
      type: Array,
      default: () => []
    }
  },
  methods: {
    //树形数据过滤方法
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    //表头变化
    getTableHeader() {
      this.header = []
      this.tableList.forEach(item => {
        this.chooseLabels.forEach(index => {
          if (item.label === index) {
            this.header.push(item)
          }
        })
      })
      this.$emit(
        'getHeaderMsg',
        this.header,
        this.chooseKeys,
        this.chooseLabels
      )
    },
    //保存选中的数据
    comfirm() {
      //获取被选中的节点元素，可抽离出id和label
      let res = this.$refs.tree.getCheckedNodes()
      this.chooseKeys = []
      this.chooseLabels = []
      res &&
        res.forEach(item => {
          this.chooseKeys.push(item.id)
          this.chooseLabels.push(item.label)
        })
      this.getTableHeader()
      // this.dialogVisible = false
      //调用关闭
      this.$refs.popverRef.doClose()
    },
    //保存方案,下次直接用这个回显
    async comfimSave() {
      this.comfirm()
      //再调用一下后台接口
      const params = {
        pageId: this.pageId,
        checkKey: this.chooseKeys.join(','),
        checkName: this.chooseLabels.join(',')
      }
      const res = await updateHeader(params)

      if (res.code === '0') {
        this.$message({
          type: 'success',
          message: '保存方案成功'
        })
      }
    },
    //点击表头自定义按钮
    chooseTitle() {
      this.$nextTick(() => {
        this.$refs.tree.setCheckedKeys(this.chooseKeys)
      })
    },
    //弹窗显示监听方法
    showPopver() {
      this.chooseKeys = this.checkedKey
      this.chooseLabels = this.checkedLabel
      this.header = this.tableHeader
    }
  },

  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    },
    //监听传进来的值
    isPopShow(val) {
      console.log(val)
      this.chooseKeys = this.checkedKey
      this.chooseLabels = this.checkedLabel
      this.header = this.tableHeader
    }
  }
}
</script>
<style lang="scss" scoped>
.tree-style {
  margin-top: 10px;
  height: 400px;
  overflow-y: auto;
}
.bottom {
  text-align: center;
  border-top: 1px solid #d9d9d9;
  padding-top: 16px;
  margin-top: 5px;
}
</style>
