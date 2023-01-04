import { createRouter, RouterMount } from 'uni-simple-router'

const router = createRouter({
  platform: 'h5',
  h5: {
    mode: 'history'
  },
  routes: [
    ...ROUTES
    // {
    //   path: '*',
    //   // eslint-disable-next-line
    //   redirect: to => {
    //     return { name: '404' }
    //   }
    // }
  ]
})

export { router, RouterMount }
