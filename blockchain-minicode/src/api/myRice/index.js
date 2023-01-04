import request from '@/http'

/**
 * @description: 获取收割列表
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getReapBatchList(data) {
  return request({
    url: '/blockchainadmin/srvQrCode/reapBatchList',
    method: 'post',

    data
  })
}

/**
 * @description: 收割批次激活二维码
 * @param {
 *  businessId: String, // 收割批次ID
 *  qrCode: String // 扫描到的二维码
 * } data
 * @author: jiagui_chen
 */
export function actQrCode(data) {
  return request({
    url: '/blockchainadmin/srvQrCode/reapBatchActQrCode',
    method: 'post',

    data
  })
}

/**
 * @description: 获取品种下拉框列表数据
 * @author: jiagui_chen
 */
export function getVarietiesList() {
  return request({
    url: '/blockchainadmin/sysDict/varietiesList',
    method: 'get',
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    }
  })
}

/**
 * @description: 获取流程列表
 * @author: jiagui_chen
 */
export function getProcessList(data) {
  return request({
    url: '/blockchainadmin/srvQrCode/progressList',
    method: 'post',

    data
  })
}

/**
 * @description: 获流程详情
 * @author: jiagui_chen
 */
export function getProgressSrvInfo(data) {
  return request({
    url: '/blockchainadmin/srvQrCode/progressSrvInfo',
    method: 'post',

    data
  })
}

/**
 * @description: 获种植信息
 * @author: guoxi
 */
export function geth5SrvInfo(data) {
  return request({
    url: '/blockchainadmin/srvH5/progressPlantInfo',
    method: 'post',
    data
  })
}
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
 * @description: 获种h5溯源信息加工id
 * @author: guoxi
 */
export function getSrvInfoByJgBatch(data) {
  return request({
    url: '/blockchainadmin/srvH5/srvInfoByJgBatch',
    method: 'post',

    data
  })
}
