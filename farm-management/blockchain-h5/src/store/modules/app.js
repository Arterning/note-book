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
  // ģ��������Ϣ
  SET_MAIN_INFO_DATA(state, mainInfoData) {
    state.mainInfoData = mainInfoData
  },
  // ģ��������Ϣ
  SET_CONFIGURATION_INFORMATION(state, configurationInformation) {
    state.configurationInformation = configurationInformation
  },
  // ������Ϣ
  SET_COLOR_ID(state, colorId) {
    state.colorId = colorId
  },
  // �ж��Ƿ�ʹ�þ�̬����
  SET_IS_STATIC_DATA(state, isStaticData) {
    state.isStaticData = isStaticData
  },
  // ��ֲȫ������
  SET_ALL_STEPS(state, allSteps) {
    state.allSteps = allSteps
  },
  // ��װ��Ϣ
  SET_PACKING_INFO(state, packingInfo) {
    state.packingInfo = packingInfo
  },
  // ��������������Ϣ
  SET_CONFIG_INFOR(state, configInfor) {
    state.configInfor = configInfor
  },
  // �����ȼ� ����ģ���
  SET_QUALITY_LEVEL(state, qualityLevel) {
    state.qualityLevel = qualityLevel
  },
  // ���󵽵���ֲ��Ϣ
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
