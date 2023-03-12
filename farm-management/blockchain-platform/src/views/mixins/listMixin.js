/*
 * @Description  : 翻页table公共逻辑
 * tips1         ：在页面中调用getListData方法时，需要实现getListDataImpl方法；需在getListDataImpl中描述请求接口的逻辑，并返回{ listData }
 * tips2         ：在页面中调用onReset方法时，需要实现onResetImpl方法；需在onResetImpl方法中重置参数
 * @Autor        : Hemingway
 * @Date         : 2021-06-23 11:26:48
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-03-09 14:49:15
 * @FilePath     : \i-farm-admin\src\views\mixins\listMixin.js
 */
export default {
  data() {
    return {
      loading: false,
      listData: []
    }
  },

  methods: {
    /**
     * @description: 重置（查询）
     * @author: Hemingway
     */
    onReset() {
      this.onResetImpl()
      this.getListData()
    },

    /**
     * @description: 查询listData
     * @param {Object} params 参数对象
     * @author: Hemingway
     */
    async getListData(params = {}) {
      this.loading = true

      const payload = { ...params }
      const res = await this.getListDataImpl(payload)

      if (res) {
        this.listData = res.listData
      } else {
        this.listData = []
      }

      this.loading = false
    }
  }
}
