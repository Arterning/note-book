import Layout from '@/layout'

const leftMenuRouter = [
  // {
  //   path: '/',
  //   component: Layout,
  //   redirect: '/operational/operational-report'
  // },
  {
    id: '1',
    parentId: '0',
    path: '/operational',
    component: Layout,
    redirect: 'noRedirect',
    name: 'operational',
    meta: {
      title: '运营数据',
      icon: 'el-icon-cpu'
    },
    children: [
      {
        id: '6',
        parentId: '1',
        path: 'operational-report',
        component: () => import('@/views/operational/report'),
        name: 'operational--operational-report',
        meta: {
          title: '运营数据报表',
          icon: 'dot',
          noCache: true
        }
      }
    ]
  }
]

export default leftMenuRouter
