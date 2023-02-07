import store from '@/store'
function checkPermission(el, binding) {
  const permissionBtn = store.getters && store.getters.permissionBtn

  const { arg } = binding

  if (!permissionBtn.includes(arg)) {
    el.style.display = 'none'
  }
}

export default {
  inserted(el, binding) {
    checkPermission(el, binding)
  }
}
