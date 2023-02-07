// import { constantRoutes } from '../router/index.js'
import { constantRoutes } from '@/router'
const getters = {
  sidebar: state => state.app.sidebar,
  size: state => state.app.size,
  device: state => state.app.device,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  token: state => state.user.token,
  userId: state => state.user.userId,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  username: state => state.user.username,
  roles: state => state.user.roles,
  introduction: state => state.user.introduction,
  // roles: state => state.user.roles,
  permission_routes: state => {
    console.log('路由', constantRoutes)
    if (state.permission.staticRoute) {
      return constantRoutes
    }
    return state.permission.routes
  },
  systemDic: state => state.dict.systemDic,
  businessDic: state => state.dict.businessDic,
  permissionBtn: state => state.permission.permissionBtn
}
export default getters
