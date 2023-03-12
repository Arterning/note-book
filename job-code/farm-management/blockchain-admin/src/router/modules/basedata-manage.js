import Layout from '@/layout'

const basedataManageRouter = {
  id: '14',
  parentId: '0',
  path: '/basedata-manage',
  component: Layout,
  redirect: 'noRedirect',
  name: 'basedata-manage',
  meta: {
    title: '基础数据管理',
    icon: 'el-icon-s-order'
  },
  children: [
    {
      id: '18',
      parentId: '14',
      path: 'companyinfo',
      component: () => import('@/views/basedata-manage/company-info'),
      name: 'company-info',
      meta: {
        title: '企业信息管理',
        icon: 'dot',
        noCache: true
      }
    },
    {
      id: '19',
      parentId: '14',
      path: 'brandcommodity',
      component: () => import('@/views/basedata-manage/brand-commodity'),
      name: 'brand-commodity',
      meta: {
        title: '品牌/商品管理',
        icon: 'dot',
        noCache: true
      }
    },
    {
      parentId: '14',
      path: 'brandcommodity/add/:id',
      component: () =>
        import('@/views/basedata-manage/brand-commodity/details.vue'),
      name: 'details',
      props: true,
      hidden: true,
      meta: {
        title: '商品详情',
        noCache: true
      }
    },
    {
      parentId: '14',
      path: 'brandcommodity/add',
      component: () =>
        import('@/views/basedata-manage/brand-commodity/addcommodity.vue'),
      name: 'addcommodity',
      hidden: true,
      meta: {
        title: '新建商品',
        noCache: true
      }
    },
    // {
    //   path: 'brandcommodity/brandlist',
    //   component: () =>
    //     import('@/views/basedata-manage/brand-commodity/brandlist.vue'),
    //   name: 'brandlist',
    //   hidden: true,
    //   meta: {
    //     title: '品牌列表',
    //     noCache: true
    //   }
    // },
    {
      id: '20',
      parentId: '14',
      path: 'reportupload',
      component: () => import('@/views/basedata-manage/report-upload'),
      name: 'report-upload',
      meta: {
        title: '质检报告上传',
        icon: 'dot',
        noCache: true
      }
    }
  ]
}

export default basedataManageRouter
