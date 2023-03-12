const getters = {
  allInfo: state => state.app.allInfo,
  code: state => state.app.code,
  mainInfoData: state => state.app.mainInfoData,
  configurationInformation: state => state.app.configurationInformation,
  colorId: state => state.app.colorId,
  isStaticData: state => state.app.isStaticData,
  allSteps: state => state.app.allSteps,
  packingInfo: state => state.app.packingInfo,
  configInfor: state => state.app.configInfor,
  qualityLevel: state => state.app.qualityLevel
}

export default getters
