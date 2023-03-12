import Layout from '@/layout'

const traceabilitySupervisionRouter = {
  id: '4',
  parentId: '0',
  path: '/traceability-supervision',
  component: Layout,
  redirect: 'noRedirect',
  name: 'traceability-supervision',
  meta: {
    title: '溯源监管',
    icon: 'el-icon-s-platform'
  },
  children: [
    {
      id: '4101',
      parentId: '4',
      path: 'blockchain-complaint-goods',
      component: () =>
        import('@/views/traceability-supervision/complaint-goods'),
      name: 'blockchain-complaint-goods',
      meta: {
        title: '投诉商品',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    },
    {
      id: '10',
      parentId: '4',
      path: 'blockchain-commodity-handling',
      component: () =>
        import('@/views/traceability-supervision/commodity-handling'),
      name: 'blockchain-commodity-handling',
      meta: {
        title: '溯源批次查询',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      }
    },
    {
      parentId: '4',
      path: 'blockchain-complaint-goods/detail',
      component: () =>
        import('@/views/traceability-supervision/complaint-goods/detail.vue'),
      name: 'blockchain-complaint-goods-detail',
      meta: {
        title: '投诉商品详情',
        icon: 'dot',
        noCache: true,
        breadcrumb: true,
        affix: false
      },
      hidden: true
    }
  ]
}

export default traceabilitySupervisionRouter
