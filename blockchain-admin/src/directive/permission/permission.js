import store from '@/store'

function checkPermission(el, binding, vnode) {
  const { arg } = binding

  const menuList = store.getters && store.getters.menuList
  const functionList = store.getters && store.getters.functionList

  const start = vnode.context.$route.path.lastIndexOf('/') + 1
  const menu = menuList.find(
    menu => menu.code === vnode.context.$route.path.slice(start)
  )
  const actions = functionList
    .filter(func => func.menuId === menu.id)
    .map(func => func.code)
  if (!actions.includes(arg)) {
    el.parentNode && el.parentNode.removeChild(el)
  }
}

export default {
  inserted(el, binding, vnode) {
    checkPermission(el, binding, vnode)
  }
  // update(el, binding, vnode) {
  //   checkPermission(el, binding, vnode)
  // }
}
