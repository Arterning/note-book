import newRequest from '@/utils/newRequest'

//驾驶舱-获取指定农场的信息
export function getFarmInfo(query) {
  return newRequest({
    url: '/agriculturepub/v1/farm/getFarmInfo',
    method: 'get',
    params: query
  })
}

// 获取农场田间设备
export function getDeviceList(data) {
  return newRequest({
    url: '/agriculturepub/v1/equipmentInfo/getDeviceList',
    method: 'get',
    data
  })
}
// 获取地块数据
export function getSectionInfo(data) {
  return newRequest({
    url: '/agriculturepub/v1/section/getSectionInfo',
    method: 'get',
    data
  })
}
// 获取农机数量
export function getDeviceCount(data) {
  return newRequest({
    url: '/agriculturepub/v1/agriculturalMachineryTask/getDeviceCount',
    method: 'get',
    data
  })
}
//查询天气信息
export function getWeatherData(data) {
  return newRequest({
    url: '/agriculturepub/v1/weatherData/queryWeatherInfo',
    method: 'post',
    data: data
  })
}
// 海康摄像头预览URL
export function queryPreviewUrl(data) {
  return newRequest({
    url: '/agriculturepub/v1/cameras/queryPreviewURLs',
    method: 'post',
    data
  })
}
// 根据农场ID查询所有土样检测结果

export function queryByFarmId(data) {
  return newRequest({
    url: '/agriculturepub/v1/soilDetectionInfo/queryByFarmId',
    method: 'post',
    data
  })
}

// 根据地块查询水位与土壤湿度数据
export function getSectionWaterLevel(data) {
  return newRequest({
    url: '/agriculturepub/v1/section/getSectionWaterLevel',
    method: 'get',
    data
  })
}
