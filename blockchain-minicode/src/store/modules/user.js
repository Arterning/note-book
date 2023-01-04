const sessionId = uni.getStorageSync('sessionId')
const role =
  uni.getStorageSync('role') && JSON.parse(uni.getStorageSync('role'))

const user = {
  state: {
    sessionId: sessionId, // 登录态
    permissions: null, // 权限
    userInfo: null, // 用户信息
    role: role, // 角色
    dictMap: null // 数据字典
  },

  mutations: {
    setSessionId(state, sessionId) {
      state.sessionId = sessionId
    },

    setPermissions(state, permissions) {
      state.permissions = permissions
    },

    setUserInfo(state, userInfo) {
      state.userInfo = userInfo
    },

    setRole(state, role) {
      state.role = role
    },

    setDictMap(state, dictMap) {
      state.dictMap = dictMap
    }
  },

  actions: {}
}

export default user
