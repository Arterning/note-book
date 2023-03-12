import { queryBusinessDic } from '@/api/dic'
const state = {
  businessDic: {}
}
const mutations = {
  SET_BUSINESSDIC(state, businessDic) {
    state.businessDic = businessDic
  }
}

const actions = {
  async getBusinessDic({ commit }) {
    try {
      const { data } = await queryBusinessDic()
      commit('SET_BUSINESSDIC', data)

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
