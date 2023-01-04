import request from '@/http'

/**
 * @description: 获取申请列表
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getApplyList(data) {
  return request({
    url: '/blockchainadmin/srvTracingCode/apply/list',
    method: 'post',

    data
  })
}

/**
 * @description: 查看申请详情
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getApplyDetail(data) {
  return request({
    url: '/blockchainadmin/srvTracingCode/apply/detail',
    method: 'post',

    data
  })
}

/**
 * @description: 申请溯源码接口
 * @param {Object} data
 * @author: jiagui_chen
 */
export function applySourceCode(data) {
  return request({
    url: `/blockchainadmin/srvTracingCode/apply`,
    method: 'post',

    data
  })
}

/**
 * @description: 溯源码接口
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getCodeList(data) {
  return request({
    url: `/blockchainadmin/srvMachPacking/list`,
    method: 'post',

    data
  })
}

/**
 * @description: 绑定溯源码接口
 * @param {Object} data
 * @author: guoxi
 */
export function bindSourceCode(data) {
  return request({
    url: `/blockchainadmin/srvMachPacking/bindSourceCode`,
    method: 'post',

    data
  })
}

/**
 * @description: 获取起始溯源码接口(结尾溯源码默认为起始溯源码+申请数量)
 * @author: jiagui_chen
 */
export function getMincode() {
  return request({
    url: '/blockchainadmin/srvTracingCode/get/mincode',
    method: 'get',
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    }
  })
}
