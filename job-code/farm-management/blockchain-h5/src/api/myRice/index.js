import request from '@/http'

/**
 * @description: 获种h5溯源信息
 * @author: guoxi
 */
export function geth5AllInfo(data) {
  return request({
    url: '/blockchainadmin/srvH5/srvInfo',
    method: 'post',
    data
  })
}

/**
 * @description: 获种h5溯源信息
 * @author: guoxi
 */
export function getMaterial(data) {
  return request({
    url: '/blockchainadmin/v2/h5/material/query',
    method: 'get',
    data
  })
}

/**
 * @description:溯源统计
 * @author: guoxi
 */
export function addAccRecord(data) {
  return request({
    url: '/blockchainadmin/srvH5AccRecReport/addAccRecord',
    method: 'post',
    data
  })
}

/**
 * @description:获取物联网数据
 * @author: guoxi
 */
export function getFarmWeatherInfo(farmId) {
  return request({
    url: '/blockchainadmin/comm/v2/farmWeather/query?farmId=' + farmId,
    method: 'post'
  })
}

/**
 * @description: 获取溯源模板主体信息
 * @author: qjj
 */
export function getMainInfo(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/mainInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取扫码记录
 * @author: qjj
 */
export function getScanInfo(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/scanInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取溯源产品信息
 * @author: qjj
 */
export function getProductInfo(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/productInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取溯源产品质检信息
 * @author: qjj
 */
export function getCheckInfo(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/checkInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取溯源产品种植信息
 * @author: qjj
 */
export function getPlantInfo(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/plantInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取溯源加工信息
 * @author: qjj
 */
export function getProcessInfo(data) {
  return request({
    url: '/blockchainadmin//v2/h5/template/processInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取溯源企业信息
 * @author: qjj
 */
export function getEnterpriseInfo(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/enterpriseInfo',
    method: 'get',
    data
  })
}

/**
 * @description: 获取页面配置信息
 * @author: qjj
 */
export function getTemplateConfig(data) {
  return request({
    url: '/blockchainadmin/v1/chaincode/template/get',
    method: 'get',
    data
  })
}

/**
 * @description: 提交好评消息
 * @author: qjj
 */
export function submitGoodComment(data) {
  return request({
    url: '/blockchainadmin/v2/goodDetails/add',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

/**
 * @description: 提交投诉消息
 * @author: qjj
 */
export function submitComplaintDetails(data) {
  return request({
    url: '/blockchainadmin/v2/complaintDetails/add',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

/**
 * @description: 获取溯源码状态
 * @author: qjj
 */
export function chainCodeStatus(data) {
  return request({
    url: '/blockchainadmin/v2/h5/template/chainCodeStatus',
    method: 'get',
    data
  })
}
