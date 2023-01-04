import newRequest from '@/utils/newRequest'
import request from '@/utils/request'

// 获取溯源码批次
export function getChaincodeBatch(query) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/batch/get',
    method: 'get',
    params: query
  })
}

// 溯源码创建
export function createCode(data) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/create',
    method: 'post',
    data
  })
}

//品牌列表
export function getbrandList(data) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/brand/list',
    method: 'post',
    data
  })
}

// 查询所有列表数据
export function getList(data) {
  return newRequest({
    url: `/blockchainadmin/v1/chaincode/list`,
    method: 'post',
    data
  })
}

// 下载文件
export function downFile(data, fileName) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/download',
    method: 'post',
    isDownload: true,
    fileName,
    data
  })
}

// 获取绑定/待绑定数据列表
export function getBindList(data) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/packing/listNew',
    method: 'post',
    data
  })
}

// 获取包装查询条件【商品-品种-加工厂-农场】
export function getConditionList(params) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/packing/conditions/query',
    method: 'get',
    params
  })
}

// 绑定溯源码
export function bondChaincode(data) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/bond',
    method: 'post',
    data
  })
}

// 查询【品牌-商品】可用溯源码批次
export function getVailableCode(params) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/available/get',
    method: 'get',
    params
  })
}

// 绑定溯源码
export function boundChaincode(data) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/bond',
    method: 'post',
    data
  })
}

// 品牌模板查询
export function getTemplate(params) {
  return newRequest({
    url: '/blockchainadmin/v1/chaincode/template/get',
    method: 'get',
    params
  })
}

// 品牌模板修改
export function modifyTemplate(data, brandId) {
  return newRequest({
    url: `/blockchainadmin/v1/chaincode/template/modify?brandId=${brandId}`,
    method: 'post',
    data
  })
}

export function getGoods(params) {
  return request({
    url: '/blockchainadmin/comm/v2/commodityByEnterpriseId',
    method: 'post',
    params
  })
}
