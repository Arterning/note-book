/*
 * @Description  : 【单列表】查询公用逻辑 - 查询/翻页/刷新/模糊
 * @Autor        : Hemingway
 * @Date         : 2021-09-17 14:25:33
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-22 16:10:30
 * @FilePath     : \blockchain-minicode\src\mixins\listMixin.js
 */
import { stringifyDate } from '@/utils/tool'

export default {
  data() {
    return {
      scrollTop: 0, // 页面滚动距离
      status: 'more', // 列表加载状态
      pageNum: 1,
      list: [],
      // 观测对象
      watchObj: {
        name: '' // 搜索框keyword
      }
    }
  },

  watch: {
    watchObj: {
      handler(newV, oldV) {
        console.log(newV, oldV)
        this.onRefresh()
      },
      deep: true
    }
  },

  mounted() {
    this.queryList() // 首次加载
  },

  /**
   * @description: 监听页面滚动事件
   * @param {Object} event
   * @author: Hemingway
   */
  onPageScroll(event) {
    this.scrollTop = event.scrollTop
  },

  /**
   * @description:触底翻页
   * @author: Hemingway
   */
  onReachBottom() {
    console.log('到底了')

    this.queryList() // 加载下一页
  },

  /**
   * @description: 下拉刷新
   * @author: Hemingway
   */
  onPullDownRefresh() {
    console.log('下拉刷新')
    this.onRefresh()
  },

  methods: {
    /**
     * @description: 搜索框搜索事件
     * @param {String} e 搜索关键词
     * @author: Hemingway
     */
    onSearch(e) {
      this.watchObj.name = e
    },

    /**
     * @description: 搜索框清空事件
     * @author: Hemingway
     */
    onClear() {
      this.watchObj.name = ''
    },

    /**
     * @description: 列表刷新查询
     * @author: Hemingway
     */
    onRefresh() {
      this.status = 'more'
      this.pageNum = 1
      this.list = []
      this.queryList() // 首次加载
    },

    /**
     * @description: 格式化日期
     * @param {*} date
     * @author: Hemingway
     */
    formatDate(date) {
      return stringifyDate(date)
    }
  }
}
