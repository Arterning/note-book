import {
  getPublicKey,
  accountLogin,
  smsLogin,
  setPassword,
  updatePassword,
  resetPassword,
  updatePhone,
  logout,
  getUserInfo
} from '@/api/user'
import {
  getToken,
  setToken,
  setName,
  removeToken,
  setUsernameLTS,
  getPwdStatus,
  setPwdStatus,
  removePwdStatus
} from '@/utils/auth'
import { /* router, */ resetRouter } from '@/router'
import JSEncrypt from 'jsencrypt/bin/jsencrypt'

/**
 * @description: rsa加密
 * @param {String} raw 被加密的字符
 * @author: Hemingway
 */
async function encrypt(raw) {
  const encrypt = new JSEncrypt()
  try {
    const publicKey = await getKey()
    if (publicKey) {
      // console.log('publicKey = ', publicKey)
      encrypt.setPublicKey(publicKey)
      return encrypt.encrypt(raw)
    }
  } catch (error) {
    console.log('rsa加密 error = ', error)
  }
}
/**
 * @description: 获取公钥
 * @author: Hemingway
 */
async function getKey() {
  try {
    const { code, publicKey } = await getPublicKey()
    if (code === '0') {
      return publicKey
    }
  } catch (error) {
    console.log('获取加密公钥 error = ', error)
  }
}

const state = {
  token: getToken(),
  hasPassword: getPwdStatus() === 'N',
  name: '',
  username: '',
  orgRoles: [],
  //TODO: 20220706 刷新后管理员信息丢失，做持久化存储
  userInfo: null,
  telphone: ''
  // introduction: '',
  // avatar: ''
}

const getters = {
  //TODO: 20220706 刷新后管理员信息丢失，做持久化存储
  userInfoGetter(state) {
    let userInfo = state.userInfo
    if (!userInfo) {
      try {
        userInfo =
          JSON.parse(
            localStorage.getItem(
              process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO'
            )
          ) || {}
        state.userInfo = userInfo
        // eslint-disable-next-line no-empty
      } catch (error) {
        console.log(error)
      }
    }
    return userInfo
  },
  getTelphone(state) {
    let telphone = state.telphone
    if (!telphone) {
      try {
        telphone = localStorage.getItem(
          process.env.VUE_APP_UNIQUE_PREFIX + '_TELPHONE'
        )
        state.telphone = telphone
        // eslint-disable-next-line no-empty
      } catch (error) {
        console.log(error)
      }
    }
    return telphone
  }
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_PASSWORD_STATUS: (state, hasPassword) => {
    state.hasPassword = hasPassword
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_USERNAME: (state, username) => {
    state.username = username
  },
  SET_ORGROLES: (state, orgRoles) => {
    state.orgRoles = orgRoles
  },
  // SET_INTRODUCTION: (state, introduction) => {
  //   state.introduction = introduction
  // },
  // SET_AVATAR: (state, avatar) => {
  //   state.avatar = avatar
  // }
  //TODO: 20220706 刷新后管理员信息丢失，做持久化存储
  UPDATE_USER_INFO: (state, userInfo) => {
    state.userInfo = userInfo
    if (userInfo) {
      localStorage.setItem(
        process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO',
        JSON.stringify(userInfo)
      )
      return
    }
    //当userInfo为null时，清除持久化存储的用户信息
    localStorage.removeItem(process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO')
  },
  SET_TELPHONE: (state, telphone) => {
    state.telphone = telphone
    if (telphone) {
      localStorage.setItem(
        process.env.VUE_APP_UNIQUE_PREFIX + '_TELPHONE',
        telphone
      )
      return
    }
    //当telphone为null时，清除持久化存储的用户信息
    localStorage.removeItem(process.env.VUE_APP_UNIQUE_PREFIX + '_TELPHONE')
  }
}

const actions = {
  // 用户登录
  async login({ commit }, payload) {
    const { isPasswordLogin, username, password, code } = payload
    let api,
      paramObj = null
    if (isPasswordLogin) {
      api = accountLogin
      paramObj = { username, password: await encrypt(password) }
    } else {
      api = smsLogin
      paramObj = { phone: username, code }
    }

    return new Promise((resolve, reject) => {
      api(paramObj)
        .then(response => {
          //TODO: 20220706 刷新后管理员信息丢失，做持久化存储
          commit('UPDATE_USER_INFO', response)
          //TODO: 20220706 刷新后管理员信息丢失，存储电话号码
          commit('SET_TELPHONE', username)
          const { sessionId, needSetPassword, name } = response

          commit('SET_TOKEN', sessionId)
          setToken(sessionId)

          commit('SET_NAME', name)
          setName(name)

          commit('SET_PASSWORD_STATUS', needSetPassword === 'N')
          setPwdStatus(needSetPassword)
          resolve(needSetPassword)
        })
        .catch(error => {
          reject(error)
        })
    })
  },

  // 获取用户信息
  async getUserInfo({ commit }) {
    try {
      const { organRoles, userName, phone } = await getUserInfo()
      commit('SET_ORGROLES', organRoles)
      commit('SET_NAME', userName)
      commit('SET_USERNAME', phone)
      setUsernameLTS(phone)

      return Promise.resolve()
    } catch (error) {
      console.log('获取用户信息 error = ', error)
      return Promise.reject()
    }
  },

  // 设置密码
  // eslint-disable-next-line
  async setPassword({ commit, state }, payload) {
    const { password1 } = payload
    let password = await encrypt(password1)
    const paramObj = {
      access_token: getToken(),
      onePassword: password,
      twoPassword: password,
      username: state.telphone
    }

    return new Promise((resolve, reject) => {
      setPassword(paramObj)
        .then(() => {
          commit('SET_PASSWORD_STATUS', true)
          resolve()
        })
        .catch(error => {
          reject(error)
        })
    })
  },

  // 更改密码
  // eslint-disable-next-line
  async updatePassword({ commit }, payload) {
    const { oldPassword, newPassword } = payload
    const paramObj = {
      access_token: getToken(),
      oldPassword: await encrypt(oldPassword),
      newPassword: await encrypt(newPassword)
    }

    return new Promise((resolve, reject) => {
      updatePassword(paramObj)
        .then(() => {
          resolve()
        })
        .catch(error => {
          reject(error)
        })
    })
  },

  // 重置密码
  // eslint-disable-next-line
  async resetPassword({ commit }, payload) {
    const { username, password1, certificate } = payload
    const paramObj = {
      phone: username,
      newPassWord: await encrypt(password1),
      certifiCate: certificate
    }

    return new Promise((resolve, reject) => {
      resetPassword(paramObj)
        .then(() => {
          resolve()
        })
        .catch(error => {
          reject(error)
        })
    })
  },

  // 修改手机号
  async updatePhone({ commit }, payload) {
    try {
      await updatePhone({ ...payload, access_token: getToken() })
      commit('SET_USERNAME', payload.phone)
      setUsernameLTS(payload.phone)
      Promise.resolve()
    } catch (error) {
      console.log('user = ', error)
      return Promise.reject(error)
    }
  },

  // user logout
  async logout({ commit, state, dispatch }) {
    try {
      await logout(state.token)

      dispatch('tagsView/delAllViews', null, { root: true }) // reset visited views and cached views

      // 路由置空，避免切换账号登录后，错误复用
      commit('permission/REMOVE_ROUTES', null, { root: true })
      resetRouter()

      // 移除密码状态
      removePwdStatus()

      await dispatch('resetToken')

      commit('UPDATE_USER_INFO', null)

      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    }
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      removeToken()

      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
