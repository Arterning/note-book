import store from '@/store'
import { router } from '@/router'
import { afterLogin } from '@/utils/auth'

const whiteList = ['/pages/login/index', '/pages/login/LoginExtra']

// 全局路由前置守卫
router.beforeEach(async (to, from, next) => {
  console.log('进：to = ', to, 'from = ', from)

  if (whiteList.includes(to.path)) {
    next()
  } else {
    if (getSessionId()) {
      console.log('---------- 已登录 --------')
      console.log('getPermissions()', getPermissions())
      // 已登陆

      if (getPermissions()) {
        next()
      } else if (to.path === '/pages/select-role/index') {
        next()
      } else {
        console.log('----------0 已登录未配置权限')
        await afterLogin()
        console.log('----------8 已完成整个登录过程')
        next()
      }
    } else {
      console.log('---------- 未登陆 --------')
      // 未登陆

      if (to.path === '/pages/home/index') {
        next()
      } else {
        next({ name: 'login' })
      }
    }
  }
})

function getSessionId() {
  return store.state.user.sessionId
}

function getPermissions() {
  return store.state.user.permissions
}

// 全局路由后置守卫
// router.afterEach((to, from) => {
//   console.log('出：to = ', to, 'from = ', from)
// })
