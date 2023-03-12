import Layout from '@/layout'

const traceabilityMmanageRouter = {
  id: '15',
  parentId: '0',
  path: '/traceability-manage',
  component: Layout,
  redirect: 'noRedirect',
  name: 'traceability-manage',
  meta: {
    title: '溯源码管理',
    icon: 'el-icon-paperclip'
  },
  children: [
    {
      id: '21',
      parentId: '15',
      path: 'template',
      component: () => import('@/views/traceability-manage/template'),
      name: 'traceability-manage--template',
      meta: {
        title: '数据标准配置',
        icon: 'dot',
        noCache: true
      }
    },
    {
      id: '211',
      parentId: '15',
      path: 'create',
      component: () => import('@/views/traceability-manage/create'),
      name: 'traceability-manage--create',
      meta: {
        title: '溯源码创建',
        icon: 'dot',
        noCache: true
      }
    },
    {
      id: '212',
      parentId: '15',
      path: 'bindlist',
      component: () => import('@/views/traceability-manage/bindlist'),
      name: 'traceability-manage--bindlist',
      meta: {
        title: '溯源批次列表',
        icon: 'dot',
        noCache: true
      }
    },
    {
      parentId: '15',
      path: 'bindlist/:id',
      component: () => import('@/views/traceability-manage/clickbind'),
      name: 'clickbind',
      hidden: true,
      // props: true,
      meta: {
        title: '批次绑定',
        icon: 'dot',
        noCache: true
      }
    }
  ]
}

export default traceabilityMmanageRouter
