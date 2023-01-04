import request from '@/http'

/**
 * @description: 获取合作伙伴列表
 * @param {Object} data
 * @author: guoxi
 */
export function getCooParList(data) {
  return request({
    url: '/blockchainadmin/srvCooperativePartner/cooParList',
    method: 'post',

    data
  })
}

export function getChoiceList(data) {
  return request({
    url: '/blockchainadmin/srvCooperativePartner/choiceList',
    method: 'post',

    data
  })
}

export function choiceCooPar(data) {
  return request({
    url: '/blockchainadmin/srvCooperativePartner/choiceCooPar',
    method: 'post',

    data
  })
}
export function cooOrNon(data) {
  return request({
    url: '/blockchainadmin/srvCooperativePartner/cooOrNon',
    method: 'post',

    data
  })
}
export function getEnterTypeList(data) {
  return request({
    url: '/blockchainadmin/sysUser/enterTypeList',
    method: 'get',

    data
  })
}
export function getDictList(data) {
  return request({
    url: '/blockchainadmin/sysDict/list',
    method: 'post',
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    },
    data
  })
}
export function getEnterpriseById(data) {
  return request({
    url: '/blockchainadmin/enterprise/getEnterpriseById/w/v1',
    method: 'post',
    data
  })
}

export function cancelCooPar(data) {
  return request({
    url: '/blockchainadmin/srvCooperativePartner/cancelCooPar',
    method: 'post',
    data
  })
}
