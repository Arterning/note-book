import request from '@/utils/request'
/**
 * 溯源码列表接口
 * @param {*} payload
 */
export function getTracingCodeList(payload) {
  return request({
    url: '/blockchaindbo/srvTracingCode/apply/list',
    method: 'post',
    data: payload
  })
}
/**
 * 溯源码发货接口
 * @param {*} payload
 */
export function sendout(payload) {
  return request({
    url: '/blockchaindbo/srvTracingCode/sendout',
    method: 'post',
    data: payload
  })
}
/**
 * 溯源码作废接口
 * @param {*} payload
 */
export function tracingCodeCancel(payload) {
  return request({
    url: '/blockchaindbo/srvTracingCode/cancel',
    method: 'post',
    data: payload
  })
}
/**
 * 获取溯源二维码
 * @param {*} syNum
 */
export function getSyCode(syNum) {
  return request({
    url: `/blockchaindbo/srvTracingCode/qrcode/get/${syNum}`,
    method: 'get'
  })
}

/**
 * 下载溯源码接口
 * @param {*} payload
 */
export function downloadSyCode(payload) {
  return request({
    url: `/blockchaindbo/srvTracingCode/download/syCode`,
    method: 'post',
    data: payload,
    isDownload: true,
    fileName: payload.fileName
  })
}
