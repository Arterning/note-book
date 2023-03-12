import {
  getPublicKey,
  accountLogin,
  smsLogin,
  logout,
  getUserInfo,
  updatePhone,
  updatePassword,
  setPassword
} from '@/api/user'
import {
  getToken,
  setToken,
  removeToken,
  getUserId,
  setName,
  setPwdStatus
  // setUserId
} from '@/utils/auth'
import { /* router, */ resetRouter } from '@/router'
// import md5 from 'js-md5'
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
  name: '',
  username: '',
  roles: [],
  userId: getUserId(),
  userInfo: {},
  phone: ''
  // introduction: '',
  // avatar: ''
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_USERNAME: (state, username) => {
    state.username = username
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_USER_ID: (state, userId) => {
    state.userId = userId
  },
  //TODO: 20220808 刷新后管理员信息丢失，做持久化存储
  UPDATE_USER_INFO: (state, userInfo) => {
    state.userInfo = userInfo
    if (userInfo) {
      localStorage.setItem(
        process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO',
        JSON.stringify(userInfo)
      )
      console.log(
        '持久化存储',
        localStorage.getItem(process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO')
      )
      return
    }

    //当userInfo为null时，清除持久化存储的用户信息
    localStorage.removeItem(process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO')
  },
  SET_PHONE: (state, phone) => {
    state.phone = phone
    if (phone) {
      localStorage.setItem(process.env.VUE_APP_UNIQUE_PREFIX + '_PHONE', phone)
      return
    }
    //当phone为null时，清除持久化存储的用户信息
    localStorage.removeItem(process.env.VUE_APP_UNIQUE_PREFIX + '_PHONE')
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
      paramObj = { phone: username, code, username, msgType: '01' }
    }

    return new Promise((resolve, reject) => {
      api(paramObj)
        .then(response => {
          //TODO: 20220808 刷新后管理员信息丢失，做持久化存储
          if (response) {
            localStorage.setItem(
              process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO',
              JSON.stringify(response)
            )
          } else {
            localStorage.removeItem(
              process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO'
            )
          }
          //TODO: 20220808 持久化存储登录用户手机号信息
          localStorage.setItem(
            process.env.VUE_APP_UNIQUE_PREFIX + '_PHONE',
            username
          )

          const { sessionId, isSetPassWord, name } = response

          commit('SET_TOKEN', sessionId)
          setToken(sessionId)

          commit('SET_NAME', name)
          setName(name)

          commit('SET_PASSWORD_STATUS', isSetPassWord === 'N')
          setPwdStatus(isSetPassWord)
          resolve(isSetPassWord)
        })
        .catch(error => {
          reject(error)
        })
    })
  },

  // user login
  // login({ commit }, payload) {
  //   return new Promise((resolve, reject) => {
  //     payload.password = md5(payload.password).toUpperCase()
  //     login(payload)
  //       .then(response => {
  //         // console.log('response = ', response)
  //         commit('SET_TOKEN', response.sessionId)
  //         setToken(response.sessionId)
  //         commit('SET_USER_ID', response.id)
  //         setUserId(response.id)
  //         resolve()
  //       })
  //       .catch(error => {
  //         reject(error)
  //       })
  //   })
  // },

  // 获取用户信息
  async getUserInfo({ commit }) {
    try {
      const { roleDtos, name, phone } = await getUserInfo()
      commit('SET_ROLES', roleDtos)
      commit('SET_NAME', name)
      commit('SET_USERNAME', phone)

      return Promise.resolve()
    } catch (error) {
      console.log('获取用户信息 error = ', error)
      return Promise.reject()
    }
  },

  // 设置密码
  // eslint-disable-next-line
  async setPassword({ commit }, payload) {
    const { password1 } = payload
    let password = await encrypt(password1)
    const paramObj = {
      // access_token: getToken(),
      onePassword: password,
      twoPassword: password
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

  // 修改手机号
  async updatePhone({ commit }, payload) {
    try {
      await updatePhone({ ...payload, access_token: getToken() })
      commit('SET_USERNAME', payload.phone)
      Promise.resolve()
    } catch (error) {
      console.log('user = ', error)
      return Promise.reject(error)
    }
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

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      logout(state.token)
        .then(() => {
          dispatch('tagsView/delAllViews', null, { root: true }) // reset visited views and cached views

          // 路由置空，避免切换账号登录后，错误复用
          commit('permission/REMOVE_ROUTES', null, { root: true })
          resetRouter()

          dispatch('resetToken')

          resolve()
        })
        .catch(error => {
          reject(error)
        })
    })
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
  mutations,
  actions
}
