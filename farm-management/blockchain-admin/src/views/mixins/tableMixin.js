/*
 * @Description  : 翻页table公共逻辑
 * tips1         ：在页面中调用getTableData方法时，需要实现getTableDataImpl方法；需在getTableDataImpl中描述请求接口的逻辑，并返回{ tableData, total }
 * tips2         ：在页面中调用onReset方法时，需要实现onResetImpl方法；需在onResetImpl方法中重置参数
 * @Autor        : Hemingway
 * @Date         : 2021-06-23 11:26:48
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-03-09 14:40:33
 * @FilePath     : \blockchain-admin\src\views\mixins\tableMixin.js
 */
export default {
  data() {
    return {
      loading: false,
      tableData: [],
      current: 1,
      size: 10,
      total: 0,
      selections: [],
      isIndeterminate: false
    }
  },

  watch: {
    current() {
      this.getTableData(false)
    },

    size() {
      this.getTableData(false)
    }
  },

  methods: {
    /**
     * @description: 重置（查询）
     * @author: Hemingway
     */
    onReset() {
      this.onResetImpl()
      // this.getTableData()
    },

    /**
     * @description: 查询tableData
     * @param {Boolean} refresh 是否重载
     * @param {Object} extra 其他参数对象
     * @author: Hemingway
     */
    async getTableData(refresh = true, extra = {}) {
      this.loading = true

      if (refresh) {
        this.current = 1
        this.size = 10
      }
      const payload = { pageNum: this.current, pageSize: this.size, ...extra }
      const res = await this.getTableDataImpl(payload)

      if (res) {
        this.tableData = res.tableData
        this.total = res.total
      } else {
        this.tableData = []
        this.total = 0
      }

      this.loading = false
    },

    /**
     * @description: 列表size事件
     * @param {Number} val
     * @author: Hemingway
     */
    handleSizeChange(val) {
      this.size = val
    },

    /**
     * @description: 列表翻页事件
     * @param {Number} val
     * @author: Hemingway
     */
    handleCurrentChange(val) {
      this.current = val
    },

    /**
     * @description: selections 改变事件
     * @param {Array} selections 被选择项集合
     * @author: Hemingway
     */
    onSelectionChange(selections) {
      this.selections = selections

      if (
        this.selections.length > 0 &&
        this.selections.length < this.tableData.length
      ) {
        this.isIndeterminate = true
      } else if (this.selections.length === 0) {
        this.isIndeterminate = false
        this.$refs.selected.checked = false // 改变radio表现
      } else {
        this.isIndeterminate = false
        console.log('this.$refs.selected = ', this.$refs.selected)
        this.$refs.selected.checked = true // 改变radio表现
      }
    },

    /**
     * @description: 批量操作check事件
     * @param {Boolean} checked
     * @author: Hemingway
     */
    onChange(checked) {
      if (checked) {
        // 全选
        this.$refs.table.$refs.table.clearSelection()
        this.$refs.table.$refs.table.toggleAllSelection()
      } else {
        // 取消全选
        this.$refs.table.$refs.table.clearSelection()
      }
    }
  }
}
