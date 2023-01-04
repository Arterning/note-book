import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/login', '/auth-redirect'] // no redirect whitelist

router.beforeEach(async (to, from, next) => {
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)

  // determine whether the user has logged in
  const hasToken = getToken() && getToken() === store.getters.token

  if (hasToken) {
    if (to.path === '/login') {
      // if is logged in, redirect to the home page
      next({ path: '/' })
      NProgress.done() // hack: https://github.com/PanJiaChen/vue-element-admin/pull/2939
    } else {
      // determine whether the user has obtained his permission roles through getInfo
      const hasPermission_routes =
        store.getters.permission_routes &&
        store.getters.permission_routes.length > 0
      if (hasPermission_routes) {
        next()
      } else {
        try {
          // generate accessible routes map based on roles
          const accessRoutes = await store.dispatch('permission/generateRoutes')
          // dynamically add accessible routes
          if (accessRoutes) {
            router.addRoutes(accessRoutes)
          }

          // const defaultRoutePath = getDefaultRoutePath(accessRoutes)
          // router.addRoutes([
          //   {
          //     path: '/',
          //     redirect: defaultRoutePath
          //   },
          //   { path: '*', name: '404', redirect: '/404' }
          // ]) // 建立映射
          // router.history.setupListeners()

          // await Promise.all([
          //   store.dispatch('user/getUserInfo'),
          //   store.dispatch('dict/getSystemDic'),
          //   store.dispatch('dict/getBusinessDic')
          // ])

          // hack method to ensure that addRoutes is complete
          // set the replace: true, so the navigation will not leave a history record
          // next({ ...to, replace: true })
          console.log('kankanjinli', accessRoutes)
          let url =
            accessRoutes[0].path + '/' + accessRoutes[0].children[0].path
          next(url)
        } catch (error) {
          // remove token and go to login page to re-login
          await store.dispatch('user/resetToken')
          Message.error(error || 'Has Error')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    /* has no token*/

    if (whiteList.indexOf(to.path) !== -1 || to.path.startsWith('/qrcode/')) {
      // in the free login whitelist, go directly
      next()
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})

/**
 * @description: 计算首页路由
 * @param {Array} routes
 * @author: Hemingway
 */
// function getDefaultRoutePath(routes) {
//   const getRouteLink = routes => {
//     let arr = []

//     arr.push(routes[0])
//     if (routes[0].children) {
//       arr = arr.concat(getRouteLink(routes[0].children))
//     }

//     return arr
//   }

//   const routeLink = getRouteLink(routes)
//   const defaultRoutePath = routeLink.map(route => route.path).join('/')

//   return defaultRoutePath
// }
