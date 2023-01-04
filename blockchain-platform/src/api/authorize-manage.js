import newRequest from '@/utils/newRequest'
import request from '@/utils/request'

// 运营平台-租户管理-获取租户列表
export function getTenantList(data) {
  return newRequest({
    url: '/blockchaindbo/v1/tenant/list',
    method: 'post',
    data
  })
}

// 运营平台-租户管理-创建租户
export function getTenantAdd(data) {
  return newRequest({
    url: '/blockchaindbo/v1/tenant/create',
    method: 'post',
    data
  })
}

// 运营平台-租户管理-修改租户信息
export function getTenantUpdate(data) {
  return newRequest({
    url: '/blockchaindbo/v1/tenant/update',
    method: 'post',
    data
  })
}

// 运营平台-租户管理-冻结/解冻租户
export function getTenantOperate(data) {
  return newRequest({
    url: '/blockchaindbo/v1/tenant/operate',
    method: 'post',
    data
  })
}

// 运营平台-公共【字典】- 数据字典查询(企业类型)
export function getDictQuery(data) {
  return newRequest({
    url: '/blockchaindbo/v1/dict/query',
    method: 'get',
    data
  })
}

// 根据手机号码查询企业名称
export function getEnterpriseName(data) {
  return request({
    url: '/blockchaindbo/v1/tenant/getTenantByPhone',
    method: 'post',
    data
  })
}
