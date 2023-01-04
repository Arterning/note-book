/*
 * @Description  : 【双列表】查询公用逻辑 - 查询/翻页/刷新/模糊
 * @Autor        : Hemingway
 * @Date         : 2021-09-17 14:25:33
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-25 08:54:08
 * @FilePath     : \blockchain-minicode\src\mixins\doubleListMixin.js
 */
import { queryEnterpriseRelations, queryFarmList } from '@/api/grain-steps'
import { stringifyDate } from '@/utils/tool'

export default {
  data() {
    return {
      current: 0, // 分段器页码号
      scrollTop: 0, // 页面滚动距离

      // 待处理
      unhandled: {
        status: 'more', // 列表加载状态
        pageNum: 1,
        list: []
      },

      // 已处理
      handled: {
        isInit: false,
        status: 'more', // 列表加载状态
        pageNum: 1,
        list: []
      },

      // 筛选源对象
      screen: {
        relationList: [],
        cropList: []
      },

      // 观测对象
      watchObj: {
        name: '', // 农场名称
        range: [], // 日期范围
        relation: '', //  品牌商/一体化米厂
        crop: '' // 作物品种
      }
    }
  },

  computed: {
    /**
     * @description: 操作类型
     * @return {String}
     * @author: Hemingway
     */
    type() {
      return this.current === 0 ? 'unhandled' : 'handled'
    },

    // 获取当前角色
    role() {
      return this.$store.state.user.role
    },

    // 判断登录用户角色是否是品牌商一方
    isBrandOwner() {
      return this.role.name.includes('品牌商')
    },

    /**
     * @description: 筛选器数据源（多选）
     * @return {Array}
     * @author: Hemingway
     */
    dataArray() {
      const arr = []
      arr.push({
        title: `${this.isBrandOwner ? '一体化米厂' : '品牌商'}`,
        field: 'relation', // 映射watchObj
        tags: this.screen.relationList,
        isMutliple: true
      })
      arr.push({
        title: '作物品种',
        field: 'crop', // 映射watchObj
        tags: this.screen.cropList,
        isMutliple: true,
        columnCount: 2
      })

      return arr
    },

    /**
     * @description: 筛选器数据源（单选）
     * @return {Array}
     * @author: Hemingway
     */
    dataSingerArray() {
      const arr = []
      arr.push({
        title: `${this.isBrandOwner ? '一体化米厂' : '品牌商'}`,
        field: 'relation', // 映射watchObj
        tags: this.screen.relationList,
        isMutliple: false
      })
      arr.push({
        title: '作物品种',
        field: 'crop', // 映射watchObj
        tags: this.screen.cropList,
        isMutliple: false,
        columnCount: 2
      })

      return arr
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
     * @description: 打开抽屉
     * @author: Hemingway
     */
    async onRightOpen() {
      if (this.screen.relationList.length === 0) {
        // 初始化筛选器
        const [relationList, cropList] = await Promise.all([
          this.getEnterpriseRelationList(),
          this.getCropList()
        ])

        this.screen.relationList = relationList
        this.screen.cropList = cropList

        this.$refs.drawer.open()
      } else {
        this.$refs.drawer.open()
      }
    },

    /**
     * @description: 获取 品牌商或一体化米厂 id合集
     * @return {Promise}
     * @author: Hemingway
     */
    async getEnterpriseRelationList() {
      try {
        const { code, relations } = await queryEnterpriseRelations({
          enterpriseId: this.role.enterpriseId,
          enterpriseType: this.role.enterpriseType
        })
        if (code === '0') {
          const relationList = (relations.enterprises || []).map(item => ({
            code: this.isBrandOwner ? item.cooEnterpriseId : item.enterpriseId,
            text: this.isBrandOwner
              ? item.cooEnterpriseName
              : item.enterpriseName,
            selected: false
          }))

          return Promise.resolve(relationList)
        } else {
          return Promise.reject()
        }
      } catch (error) {
        console.log('查询合作企业列表 error = ', error)
        return Promise.reject()
      }
    },

    /**
     * @description: 获取 种植品种 id集合
     * @return {Promise}
     * @author: Hemingway
     */
    async getCropList() {
      try {
        const { code, list } = await queryFarmList({ pageSize: 0 })
        if (code === '0') {
          const allList = list
            .filter(item => item.varietyId && item.varietyName)
            .map(item => ({
              code: item.varietyId,
              text: item.varietyName,
              selected: false
            }))
          const cropList = this.duplicateRemoval(allList)

          return Promise.resolve(cropList)
        } else {
          return Promise.reject()
        }
      } catch (error) {
        console.log('查询农场列表 error = ', error)
        return Promise.reject()
      }
    },

    /**
     * @description: 品种去重算法
     * @param {Array} origin
     * @return {Array}
     * @author: Hemingway
     */
    duplicateRemoval(origin) {
      const result = []
      const set = new Set()
      for (const i of origin) {
        if (!set.has(i.code)) {
          result.push(i)
          set.add(i.code)
        }
      }
      return result
    },

    /**
     * @description: 筛选项点击事件
     * @param {Object} e
     * @author: Hemingway
     */
    onSelect(e) {
      this.watchObj[e.field] = e.codeList.join()
    },

    /**
     * @description: 重置筛选项
     * @author: Hemingway
     */
    onReset() {
      Object.keys(this.watchObj).forEach(key => {
        if (key !== 'name') {
          if (key === 'range') {
            this.watchObj[key] = []
          } else {
            this.watchObj[key] = ''
          }
        }
      })
    },

    /**
     * @description: 分段器切换事件
     * @param {Object} e
     * @author: Hemingway
     */
    onSegmentChange(e) {
      console.log('~~~~~~~~~~~~e', e)
      this.current = e.currentIndex

      // 确认列表，切换时才加载
      if (this.current === 1 && !this.handled.isInit) {
        this.queryList() // 首次加载
        this.handled.isInit = true
      }
    },

    /**
     * @description: 列表刷新查询
     * @author: Hemingway
     */
    onRefresh() {
      this[this.type].status = 'more'
      this[this.type].pageNum = 1
      this[this.type].list = []
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
