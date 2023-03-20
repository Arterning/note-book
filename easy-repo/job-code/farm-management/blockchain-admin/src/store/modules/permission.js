/* eslint-disable */
import { asyncRoutes, constantRoutes } from '@/router'
import { getUserMenuTree } from '@/api/user'

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

const state = {
  addRoutes: [],
  routes: [],
  routesFull: true, // 动态路由是否完整
  defaultRoutePath: '', // 首页路由path
  hasSetPermissionRoutes: false, // 是否已经设置过动态路由
  menuList: [],
  functionList: [],
  permissionBtn: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  },
  SET_ROUTES_FULL: (state, isFullRoutes) => {
    state.routesFull = isFullRoutes
  },
  SET_DEFAULT_ROUTE_PATH: (state, defaultRoutePath) => {
    state.defaultRoutePath = defaultRoutePath
  },
  SET_MENU_LIST: (state, menuList) => {
    state.menuList = menuList
  },
  SET_FUNCTION_LIST: (state, functionList) => {
    state.functionList = functionList
  },
  REMOVE_ROUTES: (state) => {
    //FIXME: 20220706 临时写死左侧菜单，禁用 登出时权限清空操作
    // if(__PROJECT_VERSION__==="20220706")return;
    state.addRoutes = []
    state.routes = []
  },
  REMOVE_MENU_LIST: (state) => {
    state.menuList = []
  },
  REMOVE_FUNCTION_LIST: (state) => {
    state.functionList = []
  },
  SET_PERMISSION_BTN: (state, permissionBtn) => {
    state.permissionBtn = permissionBtn
  }
}

const actions = {
  generateRoutes({ commit,rootGetters }) {
    return new Promise((resolve, reject) => {
      let userInfo = rootGetters['user/userInfoGetter'].userRole[0]
      let params = {
        roleId: userInfo.initRoleId,
        appType: '2'
      }
      // 调接口获取动态路由
      getUserMenuTree(params)
        .then(({ data }) => {
          let resArr = data.menuTree ? data.menuTree[0].children : []
          commit('SET_PERMISSION_BTN', data.functionIds)
          const accessedRoutes = computeRouter(resArr, asyncRoutes)
          commit('SET_ROUTES', accessedRoutes)
          resolve(accessedRoutes)
        })
        .catch(error => {
          reject(error)
        })
    })
  }
  // 生成路由
  // generateRoutes({ commit, rootGetters }) {
  //   return new Promise((resolve, reject) => {
  //     // 调接口获取动态路由
  //     getUserMenuTree()
  //       .then(response => {
  //         // console.log('response = ', response)

  //         let accessedRoutes = [],
  //         menuList = [],
  //         functionList = []

  //         if (response.menuList && response.menuList.length > 0) {
  //           accessedRoutes = getAsyncRoutes(
  //             response.menuList,
  //             asyncRoutes
  //           )

  //           menuList = handleMenuTreeToList(response.menuList)
  //         }
  //         if(response.functionList && response.functionList.length > 0) {
  //           functionList = response.functionList
  //           // console.log("functionList = ", functionList)
  //         }

  //         commit('SET_ROUTES', accessedRoutes)
  //         commit('SET_MENU_LIST', menuList)
  //         commit('SET_FUNCTION_LIST', functionList)
  //         resolve(accessedRoutes)
  //       })
  //       .catch(error => {
  //         reject(error)
  //       })
  //   })
  // },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
