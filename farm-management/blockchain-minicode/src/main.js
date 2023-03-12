import Vue from 'vue'
import App from './App'

import store from './store'
Vue.prototype.$store = store

import { router } from './router'
Vue.use(router)

import './utils/uni-ui'

import './utils/permission'
import '@/common/iconfont/iconfont.css'

Vue.config.productionTip = false
App.mpType = 'app'

Vue.prototype.$bus = new Vue()

const app = new Vue({
  store,
  ...App
})

app.$mount()
