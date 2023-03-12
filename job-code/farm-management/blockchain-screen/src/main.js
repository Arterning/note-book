import Vue from 'vue'
import * as ECharts from 'echarts'
import VueECharts from 'vue-echarts'
import './plugins/element.js'
import './style/index.scss'
Vue.config.productionTip = false
Vue.prototype.$echarts = ECharts
Vue.component('v-chart', VueECharts)

import router from './router'
import App from './App.vue'

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
