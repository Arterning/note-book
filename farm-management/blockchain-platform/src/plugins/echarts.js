import * as ECharts from 'echarts'
import VueECharts from 'vue-echarts'

export default {
  install(Vue) {
    Vue.prototype.$echarts = ECharts
    Vue.component('v-chart', VueECharts)
  }
}
