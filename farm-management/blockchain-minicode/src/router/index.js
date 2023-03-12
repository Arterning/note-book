import { createRouter } from 'uni-simple-router'

const router = createRouter({
  platform: 'mp-weixin',
  routes: [...ROUTES]
})

export { router }
