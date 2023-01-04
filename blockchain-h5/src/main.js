import Vue from 'vue'
import App from './App'
import store from './store'

import uView from 'uview-ui'
Vue.use(uView)

Vue.config.productionTip = false

import './permission'

import { router, RouterMount } from './router'
Vue.use(router)

App.mpType = 'app'

const app = new Vue({
  store,
  ...App
})
RouterMount(app, router, '#app')
