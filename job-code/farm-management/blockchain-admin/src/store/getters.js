import { constantRoutes } from '@/router'

const getters = {
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
  device: state => state.app.device,
  showOverlay: state => state.app.showOverlay,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  token: state => state.user.token,
  hasPassword: state => state.user.hasPassword,
  name: state => state.user.name,
  username: state => state.user.username,
  orgRoles: state => state.user.orgRoles,
  // introduction: state => state.user.introduction,
  // avatar: state => state.user.avatar,
  // roles: state => state.user.roles,
  permission_routes: state => {
    if (!state.permission.routesFull) {
      return constantRoutes.concat(
        state.permission.addRoutes.filter(
          route => route.path === '/system-manage'
        )
      )
    }
    return state.permission.routes
  },
  hasSetPermissionRoutes: state => state.permission.hasSetPermissionRoutes,
  defaultRoutePath: state => state.permission.defaultRoutePath,
  menuList: state => state.permission.menuList,
  functionList: state => state.permission.functionList,
  businessDic: state => state.dic.businessDic,
  permissionBtn: state => state.permission.permissionBtn
}
export default getters
