import { getSystemDic, queryBusinessDic } from '@/api/dict'
const state = {
  systemDic: {}, // 系统字典
  businessDic: {} // 业务字典
}
const mutations = {
  SET_SYSTEM_DIC(state, systemDic) {
    state.systemDic = systemDic
  },
  SET_BUSINESSDIC(state, businessDic) {
    state.businessDic = businessDic
  }
}

const actions = {
  // 获取系统字典
  async getSystemDic({ commit }) {
    try {
      const { allDictionaryKeyValueMap } = await getSystemDic()
      commit('SET_SYSTEM_DIC', allDictionaryKeyValueMap)

      return Promise.resolve()
    } catch (error) {
      console.log('业务字典查询 error = ', error)
      return Promise.reject()
    }
  },

  // 获取业务字典
  async getBusinessDic({ commit }) {
    try {
      const { dictMap } = await queryBusinessDic()
      commit('SET_BUSINESSDIC', dictMap)

      return Promise.resolve()
    } catch (error) {
      console.log('业务字典查询 error = ', error)
      return Promise.reject()
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
