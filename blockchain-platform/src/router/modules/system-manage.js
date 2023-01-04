import Layout from '@/layout'

const systemManageRouter = {
  id: '5',
  parentId: '0',
  path: '/system-manage',
  component: Layout,
  redirect: 'noRedirect',
  name: 'system-manage',
  meta: {
    title: '系统管理',
    icon: 'el-icon-s-tools'
  },
  children: [
    {
      id: '11',
      parentId: '5',
      path: 'role',
      component: () => import('@/views/system-manage/role'),
      name: 'system-manage--role',
      meta: {
        title: '角色管理',
        icon: 'dot',
        noCache: true
      }
    },
    {
      id: '12',
      parentId: '5',
      path: 'user',
      component: () => import('@/views/system-manage/user'),
      name: 'system-manage--user',
      meta: {
        title: '员工管理',
        icon: 'dot',
        noCache: true
      }
    }
  ]
}

export default systemManageRouter
