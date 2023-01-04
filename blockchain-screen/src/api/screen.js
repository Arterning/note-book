import request from '../utils/request'
//查询数据指挥舱数据
export function queryCommand() {
  return request({
    url: '/blockchainadmin/statistics/majority/query',
    method: 'get'
  })
}
//查询农场地图数据
export function queryFarmMap(data) {
  return request({
    url: '/blockchainadmin/statistics/farmMap/query',
    method: 'get',
    data
  })
}
//上链数据分布查询
export function queryLink() {
  return request({
    url: '/blockchainadmin/statistics/uploads/query',
    method: 'get'
  })
}
//入驻企业分布查询
export function queryEnterprise() {
  return request({
    url: '/blockchainadmin/statistics/enterprise/query',
    method: 'get'
  })
}
//扫码数据查询
export function queryScanData(data) {
  return request({
    url: '/blockchainadmin/statistics/scanner/query',
    method: 'get',
    data
  })
}
