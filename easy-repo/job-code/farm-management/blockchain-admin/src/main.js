import Vue from 'vue'
import router from './router'
import store from './store'
import App from './App'

import 'normalize.css/normalize.css'
import './styles/index.scss' // global css
import './icons' // svg-icon

import elementUI from './plugins/element-ui.js'
import echarts from '@/plugins/echarts.js'
import businessComponents from '@/plugins/business-components.js'
import moment from 'moment'
Vue.use(elementUI)
Vue.use(echarts)
Vue.use(businessComponents)
Vue.prototype.$moment = moment

import VueAMap from 'vue-amap'
Vue.use(VueAMap)
VueAMap.initAMapApiLoader({
  key: '4a08e25b3d0da69a985fdc8604bb9f85',
  plugin: [
    'AMap.Autocomplete',
    'AMap.LngLat',
    'AMap.Marker',
    'AMap.Geocoder',
    'AMap.PlaceSearch',
    'AMap.Scale',
    'AMap.OverView',
    'AMap.ToolBar',
    'AMap.MapType',
    'AMap.PolyEditor',
    'MarkerClusterer'
  ],
  v: '1.4.9'
})

import * as filters from './filters'
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.prototype.$bus = new Vue()

import './permission' // 菜单权限
import permission from '@/directive/permission' // 功能权限
Vue.use(permission)
import permission1 from '@/directive/btn-permission' // 按钮权限
Vue.use(permission1)

Vue.config.productionTip = false

//FIXME: 20220706 临时写死左侧菜单，禁用登出权限清空
window.__PROJECT_VERSION__ = '20220706'
window.vm = new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})