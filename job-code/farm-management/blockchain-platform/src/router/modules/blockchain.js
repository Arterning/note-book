import Layout from '@/layout'

const blockchainRouter = {
  path: '/blockchain',
  component: Layout,
  redirect: 'noRedirect',
  name: 'blockchain',
  meta: {
    title: '区块链溯源管理',
    icon: 'blockchain'
  },
  children: [
    {
      path: 'blockchain-settle-audit',
      component: () => import('@/views/blockchain/settle-audit'),
      name: 'blockchain--blockchain-settle-audit',
      meta: {
        title: '入驻审核',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    },
    {
      path: 'blockchain-traceability-code',
      component: () => import('@/views/blockchain/traceability-code'),
      name: 'blockchain--blockchain-traceability-code',
      meta: {
        title: '溯源码管理',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    },
    {
      path: 'blockchain-role-manage',
      component: () => import('@/views/blockchain/role-manage'),
      name: 'blockchain--blockchain-role-manage',
      meta: {
        title: '角色管理',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    }
  ]
}

export default blockchainRouter
