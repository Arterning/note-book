<!--
 * @Description  : 应用preview
 * @Autor        : Hemingway
 * @Date         : 2022-06-13 18:00:59
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 16:45:12
 * @FilePath     : \blockchain-h5\src\pages\preview\index.vue
-->
<template>
  <view class="preview"></view>
</template>

<script>
import { getMainInfo, getTemplateConfig, chainCodeStatus } from '@/api/myRice'
import templateConfig from '@/static/json/templateConfig.js'
export default {
  name: 'PrewView',

  async onLoad() {
    // 判断页面在不在iframe里面
    if (self.location !== top.location) {
      this.type = ''
      this.$store.commit('app/SET_CONFIGURATION_INFORMATION', templateConfig)
      this.$store.commit('app/SET_IS_STATIC_DATA', true)
      return
    }
    this.type = '1'
    const code = this.$Route.query.qrcode
    if (code) {
      this.qrcode = code
      this.$store.commit('app/SET_CODE', code)
      await this.getPageData()
      await this.getConfig()
      this.queryTemplate(code)
    }
  },
  data() {
    return {
      qrcode: '',
      mainInfoData: {},
      firstInto: true,
      // 模板编号
      pageSn: '1',
      type: '' // 类型为1时新增查询记录
    }
  },
  computed: {
    packingInfo() {
      return this.$store.getters.packingInfo
    }
  },
  mounted() {
    // 获取pc端传过来的信息
    let that = this
    window.addEventListener(
      'message',
      async function(e) {
        if (e.data.type === 'webpackOk' || e.data.type === 'webpackClose')
          return
        console.log('传递接收', e.data)
        // 模板数据，不需要请求接口
        if (e.data instanceof Array) {
          that.$store.commit('app/SET_CONFIGURATION_INFORMATION', e.data)
          // 首次进来才进路由
          if (that.firstInto) {
            that.firstInto = false
            that.$Router.push({
              path: `/entry-1`,
              query: {
                id: that.qrcode
              }
            })
          }
        } else {
          // 判断是否是绑定页面跳转或者大屏项目进入
          if (e.data.machPackingId || e.data?.qrCode) {
            that.$store.commit('app/SET_IS_STATIC_DATA', false)
            that.$store.commit('app/SET_PACKING_INFO', e.data)
            // 设置溯源码
            that.qrcode = e.data?.qrCode || ''
            that.$store.commit('app/SET_CODE', that.qrcode)
            await that.getPageData()
            await that.getConfig()
            that.$Router.push({
              path: `/entry-${that.pageSn || '1'}`,
              query: {
                id: that.qrcode
              }
            })
          }
          // 判断是否是创建页面
          if (e.data.templateId) {
            templateConfig[0].templateId = e.data.templateId
            templateConfig[0].colorId = e.data.colorId
            that.$store.commit(
              'app/SET_CONFIGURATION_INFORMATION',
              templateConfig
            )
            if (that.pageSn === e.data.templateId) return
            that.pageSn = e.data.templateId
            that.$Router.push({
              path: `/entry-${that.pageSn}`,
              query: {
                id: that.qrcode
              }
            })
          }
        }
      },
      false
    )
  },

  methods: {
    /**
     * @description: 查询模板信息
     * @author: Hemingway
     */
    queryTemplate() {
      // 判断如果pageSn为空值，那么跳转到空白页
      if (JSON.stringify(this.mainInfoData) !== '{}') {
        // ...传递编号，调接口获取h5模板序号
        this.$Router.push({
          path: `/entry-${this.pageSn}`,
          query: {
            id: this.qrcode
          }
        })
      } else {
        this.$Router.push({
          path: '/blank',
          query: {
            id: this.qrcode
          }
        })
      }
    },
    /**
     * @description: 获取当前扫码位置信息
     * @author: qjj
     */
    getLocation() {
      return new Promise(resolve => {
        uni.getLocation({
          type: 'wgs84',
          geocode: true,
          success: function(item) {
            console.log('获取位置信息成功')
            resolve(item)
          },
          fail: function() {
            console.log('获取位置信息失败')
            uni.showToast({
              title: '获取位置信息失败',
              icon: 'none'
            })
            resolve(null)
          }
        })
      })
    },
    /**
     * @description: 获取溯源码主要信息
     * @author: qjj
     */
    async getPageData() {
      // 判断该溯源码是否被禁用
      let status = await chainCodeStatus({
        qrCode: this.qrcode,
        machPackingId: this.packingInfo && this.packingInfo.machPackingId,
        packingBatchId: this.packingInfo && this.packingInfo.packingBatchId
      })
      if (status.data.status === '1') {
        this.$Router.push({
          path: '/forbidden',
          query: {
            id: this.qrcode,
            reason: status.data.disableReason
          }
        })
        return Promise.reject()
      }
      //TODO 谷歌浏览器定位需要连国外服务器，获取不到定位信息
      let result = await this.getLocation()
      // let result = null
      const { data } = await getMainInfo({
        qrCode: this.qrcode,
        machPackingId: this.packingInfo && this.packingInfo.machPackingId,
        packingBatchId: this.packingInfo && this.packingInfo.packingBatchId,
        location: result?.latitude
          ? result?.latitude + ',' + result?.longitude
          : '',
        type: this.type
      })
      if (!data) return
      this.mainInfoData = data || {}
      this.$store.commit('app/SET_MAIN_INFO_DATA', data)
    },
    /**
     * @description: 获取页面配置信息
     * @author: qjj
     */
    async getConfig() {
      const { data } = await getTemplateConfig({
        qrCode: this.packingInfo?.qrCode || this.qrcode,
        brandId: this.mainInfoData.brandId || ''
      })
      if (!data) return
      this.pageSn = data[0].templateId
      // this.pageSn = '2'
      this.$store.commit('app/SET_CONFIGURATION_INFORMATION', data)
    }
  }
}
</script>
