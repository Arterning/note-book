export const permissionMap = [
  {
    title: '账号管理',
    name: 'accoutManagement',
    list: [
      {
        text: '申请入驻',
        name: 'application',
        icon: 'file_add_outlined',
        url: '/pages/application/index'
      },
      {
        text: '审核进度',
        name: 'reviewProgress',
        icon: 'hourglass_outlined',
        url: '/pages/review-progress/index'
      },
      {
        text: '员工列表',
        name: 'staffManagement',
        icon: 'yuangongguanli',
        url: '/pages/employee/index'
      },
      {
        text: '合作米厂',
        name: 'partner',
        icon: 'briefcase_outlined',
        url: '/subPages/partnership/index'
      }
    ]
  },
  {
    title: '数据收集',
    name: 'dataManagement',
    list: [
      {
        text: '品牌/商品',
        name: 'commodity',
        icon: 'flag_outlined',
        url: '/pages/commodity/BrandGoodsManagement'
      },
      {
        text: '收粮清单',
        name: 'grainIn',
        icon: 'spa_outlined',
        url: '/subPages/grain-in/index'
      },
      {
        text: '烘干信息',
        name: 'drying',
        icon: 'brightness__outlined',
        url: '/subPages/drying/index'
      },

      {
        text: '加工信息',
        name: 'machining',
        icon: 'settings_vs_outlined',
        url: '/subPages/machining/index'
      },
      {
        text: '包装信息',
        name: 'packing',
        icon: 'packing',
        url: '/subPages/packing/index'
      },
      {
        text: '检测报告',
        name: 'examiningReport',
        icon: 'vertical_align_top',
        url: '/pages/report/List'
      }
    ]
  },
  {
    title: '溯源管理',
    name: 'traceManagement',
    list: [
      {
        text: '申请溯源码',
        name: 'codeApplication',
        icon: 'focus',
        url: '/pages/appSourceCode/index'
      },
      {
        text: '激活溯源码',
        name: 'codeActivation',
        icon: 'activation',
        url: '/pages/appSourceCode/activation/index'
      },
      {
        text: '数据统计',
        name: 'traceData',
        icon: 'canvas_text',
        url: '/subPages/statistics/index'
      }
    ]
  }
]
