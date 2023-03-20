import { router } from '@/router'

// 全局路由前置守卫
router.beforeEach((to, from, next) => {
  console.log('进：to = ', to, 'from = ', from)

  if (to.name === from.name && !['index', 'previewImage'].includes(to.name)) {
    let params = {}
    if (location.href.indexOf('?') !== -1) {
      let paramsArr = location.href
        .slice(location.href.indexOf('?') + 1)
        .split('&')
      paramsArr.forEach(res => {
        let obj = res.split('=')
        params[obj[0]] = obj[1]
      })
    }

    next({
      path: '/',
      // query: { qrcode: location.href.slice(start, start + 13) }
      query: { qrcode: params?.id || 110 }
    })
  } else {
    next()
  }
})

// 全局路由后置守卫
// router.afterEach((to, from) => {
//   console.log('出：to = ', to, 'from = ', from)
// })
