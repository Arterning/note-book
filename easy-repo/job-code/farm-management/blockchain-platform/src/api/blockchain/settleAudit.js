import request from '@/utils/request'

// 查询角色列表数据
export function getListData(payload) {
  return request({
    url: '/blockchaindbo/v2/enterpriseChange/query',
    method: 'post',
    data: payload
  })
}

// 查询角色详情数据
export function getAuditDetail(payload) {
  return request({
    url: '/blockchaindbo/v2/enterpriseChange/detail',
    method: 'post',
    data: payload
  })
}

// 查询参数类型
export function getInfoDict() {
  return request({
    url: '/blockchaindbo/v1/dict/query',
    method: 'post'
  })
}

// 入驻审核
export function checkEnterprise(data) {
  return request({
    url: '/blockchaindbo/v2/enterpriseChange/check',
    method: 'post',
    data
  })
}
