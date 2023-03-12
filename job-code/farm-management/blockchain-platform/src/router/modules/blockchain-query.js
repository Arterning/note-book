import Layout from '@/layout'

const blockchainQueryRouter = {
  id: '3',
  parentId: '0',
  path: '/blockchain-query',
  component: Layout,
  redirect: 'noRedirect',
  name: 'blockchain-query',
  meta: {
    title: '区块链数据查询',
    icon: 'el-icon-s-data'
  },
  children: [
    {
      id: '9',
      parentId: '3',
      path: 'blockchain-index',
      component: () => import('@/views/blockchain-query/blockchain-index'),
      name: 'blockchain-index',
      meta: {
        title: '区块链数据查询',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    }
  ]
}

export default blockchainQueryRouter
