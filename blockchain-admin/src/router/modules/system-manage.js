import Layout from '@/layout'

const systemManageRouter = {
  id: '16',
  parentId: '0',
  path: '/system-manage',
  component: Layout,
  redirect: 'noRedirect',
  name: 'system-manage',
  meta: {
    title: '组织管理',
    icon: 'system-management'
  },
  children: [
    {
      id: '22',
      parentId: '16',
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
      id: '23',
      parentId: '16',
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
