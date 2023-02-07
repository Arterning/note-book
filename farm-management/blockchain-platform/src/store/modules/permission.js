import { asyncRoutes, constantRoutes } from '@/router'
import { getUserMenuTree } from '@/api/user'

/**
 * @description: 计算动态路由
 * @param {Array} permissions 菜单权限
 * @param {Array} routes 预设路由
 * @return {Array}
 * @author: Hemingway
 */
// function getAsyncRoutes(permissions, routes) {
//   let hasSetAffix = false

//   const filterAsyncRoutes = (permissions, routes) => {
//     const res = []

//     permissions.forEach(permission => {
//       const cursor = hasPermission(routes, permission)
//       if (cursor !== -1) {
//         // 菜单权限命中路由
//         const route = routes[cursor]
//         route.hidden = permission.isDisplay === 'N' // 是否在侧边栏显示

//         const hasChildren = permission.children && route.children

//         if (!hasSetAffix && !hasChildren) {
//           route.meta.affix = true
//           hasSetAffix = true
//         } // 设置默认访问路由的affix属性

//         if (hasChildren) {
//           route.children = filterAsyncRoutes(
//             permission.children,
//             route.children
//           )
//         } else {
//           delete route.children
//         }

//         res.push(route)
//       }
//     })

//     return res
//   }

//   return filterAsyncRoutes(permissions, routes)
// }

/**
 * @description: 计算动态路由
 * @param {Array} permissions 菜单权限
 * @param {Array} routes 预设路由
 * @return {Array}
 * @author: qinjj
 */
function computeRouter(permissions, routes) {
  // 扁平化数组
  let flatPermissions = flatArr(permissions)
  let flatRoutes = flatArr(routes)
  // 根据配置信息过滤预设路由
  let filterPermissions = flatRoutes.filter(res => {
    res.children = []
    return (
      !res?.id ||
      flatPermissions.find(item => item.id === res.id)?.isCheck === 'Y'
    )
  })
  // 过滤后的数组，树形结构化
  let resultArr = getTree(filterPermissions)
  return resultArr
}

/**
 * @description: 树形扁平化
 * @param {Array} nodes 需要扁平化的树形结构
 * @param {Array} arr 返回的扁平数组
 * @return {Array}
 * @author: qinjj
 */
function flatArr(nodes = [], arr = []) {
  for (let item of nodes) {
    arr.push(item)
    if (item.children && item.children.length) {
      flatArr(item.children, arr)
    }
  }
  return arr
}

/**
 * @description: 扁平树形化，找到顶层节点
 * @param {Array} arr 需要扁平化的数组
 * @return {Array}
 * @author: qinjj
 */
function getTree(arr) {
  const treeData = []
  arr.forEach(v => {
    //找到第一层父节点
    if (v.parentId === '0') {
      treeData.push(v)
      v.children = getChildren(arr, v.id)
    }
  })
  return treeData
}

/**
 * @description: 扁平树形化，找到父节点子节点
 * @param {Array} id 当前节点id
 * @param {Array} arr 需要扁平化的数组
 * @return {Array}
 * @author: qinjj
 */
function getChildren(arr, id) {
  const children = []
  //查找父节点对应的子节点
  arr.forEach(v => {
    if (v.parentId === id) {
      v.children = getChildren(arr, v.id)
      children.push(v)
    }
  })
  return children
}
/**
 * @description: 权限判断
 * @param {Array} routes 预设路由
 * @param {Object} permission
 * @return {Index} 索引
 * @author: Hemingway
 */
// function hasPermission(routes, permission) {
//   return routes.findIndex(route => route.path === permission.code)
// }

const state = {
  addRoutes: [],
  routes: [],
  staticRoute: false,
  permissionBtn: ['btn1', 'btn2']
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  },

  REMOVE_ROUTES: state => {
    state.addRoutes = []
    state.routes = []
  },
  SET_PERMISSION_BTN: (state, permissionBtn) => {
    state.permissionBtn = permissionBtn
  }
}

const actions = {
  generateRoutes({ commit }) {
    return new Promise((resolve, reject) => {
      let userInfo = JSON.parse(
        localStorage.getItem(process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO')
      )
      let params = {
        roleId: userInfo.initRoleId,
        appType: '1'
      }
      // 调接口获取动态路由
      getUserMenuTree(params)
        .then(({ data }) => {
          let resArr = data.menuTree ? data.menuTree[0].children : []
          // 存储按钮权限数据
          commit('SET_PERMISSION_BTN', data.functionIds)
          const accessedRoutes = computeRouter(resArr, asyncRoutes)
          console.log('计算后的路由', accessedRoutes)
          commit('SET_ROUTES', accessedRoutes)
          resolve(accessedRoutes)
        })
        .catch(error => {
          reject(error)
        })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
