import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home'
import Solution from '../views/Solution'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/solution',
    name: 'solution',
    component: Solution
  }
]

const router = new VueRouter({
  routes
})

export default router
