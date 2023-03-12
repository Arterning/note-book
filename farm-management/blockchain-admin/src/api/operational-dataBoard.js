import newRequest from '@/utils/newRequest'
import request from '@/utils/request'

// PC-报表模块【报表】- 商品
export function getBrandList(query) {
  return newRequest({
    url: '/blockchainadmin/srvBrandCommodity/queryCommodityByEnterpriseIdList',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】- 原粮总量
export function getGrainGross(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/grainGross',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】- 成品大米总量
export function getFinishedProductRiceGross(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/finishedProductRiceGross',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】- 绑定溯源码的个数
export function getBondSyCodeGross(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/bondSyCodeGross',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】- 被扫码的个数
export function getBeSweptCodeGross(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/beSweptCodeGross',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-扫码的次数
export function getSweptCodeGross(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/sweptCodeGross',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-原粮品种分布
export function getGrainVarietyPie(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/grainVarietyPie',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-原粮品种分布:列表
export function getGrainVarietyList(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/grainVarietyList',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-成品大米品牌商分类饼图
export function getFinishedProductRicePie(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/finishedProductRicePie',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-成品及溯源码情况
export function getBar(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/bar',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-扫码地区分布图
export function getScAreaDistributionPie(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/scAreaDistributionPie',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-扫码数据查询：折线
export function getScDataLine(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/scDataLine',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-扫码数据查询：列表
export function getScDataList(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/scDataList',
    method: 'get',
    params: query
  })
}

// PC-报表模块【报表】-扫码数据查询：导出
export function getExport(query) {
  return newRequest({
    url: '/blockchainadmin/managementPlatformReport/export',
    method: 'post',
    params: query,
    responseType: 'blob'
  })
}
// 字典
export function queryGetDict(query) {
  return request({
    url: '/blockchainadmin/v1/dict/getDict',
    method: 'get',
    params: query
  })
}
