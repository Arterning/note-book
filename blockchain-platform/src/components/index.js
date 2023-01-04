import ITab from './i-tab' // 页面tab
import IAction from './i-action' // 功能区
import IActionItem from './i-action-item' // 功能区中的项（如筛选项）
import ISectionHeader from './i-section-header' // 内容区的分栏标题
import ITable from './i-table' // 表格
import IPagination from './i-pagination' // 表格的翻页（应用于 IFooter 中）
import IFooter from './i-footer' // 尾部栏
import IFooterSelected from './i-footer-selected' // 尾部栏中的批量选择

import IDialog from './i-dialog' // dialog 弹窗
import IUpload from './i-upload' // 文件上传

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
  IUpload
}

const install = function (Vue) {
  Object.keys(components).forEach(key => {
    Vue.component(key, components[key])
  })
}

if (typeof window !== 'undefined' && window.Vue) {
  install(window.Vue)
}

export default {
  components,
  install
}
