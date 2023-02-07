import Layout from '@/layout'

const tenantMangeRouter = {
  id: '2',
  parentId: '0',
  path: '/tenant-manage',
  component: Layout,
  redirect: 'noRedirect',
  name: 'tenant-manage',
  meta: {
    title: '租户管理',
    icon: 'el-icon-s-operation'
  },
  children: [
    {
      id: '7',
      parentId: '2',
      path: 'blockchain-service-pattern',
      component: () => import('@/views/tenant-manage/service-pattern'),
      name: 'blockchain-service-pattern',
      meta: {
        title: '服务模式管理',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    },
    {
      id: '8',
      parentId: '2',
      path: 'blockchain-authorize-manage',
      component: () => import('@/views/tenant-manage/authorize-manage'),
      name: 'blockchain-authorize-manage',
      meta: {
        title: '租户授权管理',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    },
    {
      parentId: '2',
      path: 'blockchain-authorize-manage/add',
      component: () => import('@/views/tenant-manage/authorize-manage/add.vue'),
      name: 'add-tenement',
      meta: {
        title: '新建租户',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      },
      hidden: true
    },
    {
      parentId: '2',
      path: 'blockchain-authorize-manage/edit',
      component: () =>
        import('@/views/tenant-manage/authorize-manage/edit.vue'),
      name: 'edit-tenement',
      meta: {
        title: '编辑租户',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false,
        hiddenHead: true
      },
      hidden: true
    }
  ]
}

export default tenantMangeRouter
