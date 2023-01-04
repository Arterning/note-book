import permission from './btn-permission'

const install = function (Vue) {
  Vue.directive('btn-permission', permission)
}

if (window.Vue) {
  window['btn-permission'] = permission
  Vue.use(install); // eslint-disable-line
}

permission.install = install
export default permission
