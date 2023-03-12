export const state = {
  allInfo: '',
  code: '',
  mainInfoData: '',
  configurationInformation: '',
  colorId: '',
  isStaticData: false,
  allSteps: [],
  packingInfo: null,
  configInfor: {
    product: {
      isShow: true,
      subTree: []
    },
    quality: {
      isShow: true,
      subTree: []
    },
    plant: {
      isShow: true,
      subTree: []
    },
    machining: {
      isShow: true,
      subTree: []
    },
    enterprise: {
      isShow: true,
      subTree: []
    }
  },
  qualityLevel: '',
  farmRecords: []
}

const mutations = {
  SET_CODE(state, code) {
    state.code = code
  },
  SET_ALL_INFO(state, allInfo) {
    state.allInfo = allInfo
  },
  // 模板主体信息
  SET_MAIN_INFO_DATA(state, mainInfoData) {
    state.mainInfoData = mainInfoData
  },
  // 模板配置信息
  SET_CONFIGURATION_INFORMATION(state, configurationInformation) {
    state.configurationInformation = configurationInformation
  },
  // 主题信息
  SET_COLOR_ID(state, colorId) {
    state.colorId = colorId
  },
  // 判断是否使用静态数据
  SET_IS_STATIC_DATA(state, isStaticData) {
    state.isStaticData = isStaticData
  },
  // 种植全部过程
  SET_ALL_STEPS(state, allSteps) {
    state.allSteps = allSteps
  },
  // 包装信息
  SET_PACKING_INFO(state, packingInfo) {
    state.packingInfo = packingInfo
  },
  // 处理过后的配置信息
  SET_CONFIG_INFOR(state, configInfor) {
    state.configInfor = configInfor
  },
  // 质量等级 用于模板二
  SET_QUALITY_LEVEL(state, qualityLevel) {
    state.qualityLevel = qualityLevel
  },
  // 请求到的种植信息
  SET_FARM_RECORDS(state, farmRecords) {
    state.farmRecords = farmRecords
  }
}

const actions = {}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
