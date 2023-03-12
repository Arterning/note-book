import ITab from '@/components/i-tab' // 页面tab
import IAction from '@/components/i-action' // 功能区
import IActionItem from '@/components/i-action-item' // 功能区中的筛选项（应用于 IAction 中）
import ISectionHeader from '@/components/i-section-header' // 内容区的分栏标题
import ITable from '@/components/i-table' // 表格
import IFooter from '@/components/i-footer' // 尾部行
import IFooterSelected from '@/components/i-footer-selected' // 尾部行中的批量选择（应用于 IFooter 中）
import IPagination from '@/components/i-pagination' // 表格的翻页（应用于 IFooter 中）

import IDialog from '@/components/i-dialog' // dialog 弹窗
import IUpload from '@/components/i-upload' // 文件上传
import IAmap from '@/components/i-amap' // 地图

const components = {
  ITab,
  IAction,
  IActionItem,
  ISectionHeader,
  ITable,
  IPagination,
  IFooter,
  IFooterSelected,

  IDialog,
  IUpload,
  IAmap
}

export default {
  install(Vue) {
    Object.keys(components).forEach(key => {
      Vue.component(key, components[key])
    })
  }
}
