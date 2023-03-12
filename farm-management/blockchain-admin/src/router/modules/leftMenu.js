import Layout from '@/layout'

const leftMenuRouter = {
  id: '13',
  parentId: '0',
  path: '/operational',
  component: Layout,
  redirect: 'noRedirect',
  name: 'operational',
  meta: {
    title: '数据驾驶舱',
    icon: 'el-icon-cpu'
  },
  children: [
    {
      id: '17',
      parentId: '13',
      path: 'dataBoard',
      component: () => import('@/views/operational/dataBoard'),
      name: 'dataBoard',
      meta: {
        title: '数据看板',
        icon: 'dot',
        noCache: true
      }
    },
    {
      parentId: '13',
      path: 'dataBoard/report',
      component: () => import('@/views/operational/dataBoard/report.vue'),
      name: 'report',
      hidden: true,
      meta: {
        title: '采收明细',
        noCache: true
      }
    }
  ]
}

export default leftMenuRouter
