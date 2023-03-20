import request, { getParams } from '@/utils/request'
// --------------------------------------- 品牌/商品管理 ---------------------------------------

// 品牌下拉框
export function getBrandList(data) {
  return request({
    url: '/blockchainadmin/comm/v2/brandByEnterpriseId',
    method: 'post',
    data
  })
}
// 品牌列表
export function getBrandDialog(data) {
  return request({
    url: '/blockchainadmin/srvBrand/list',
    method: 'post',
    data
  })
}

// 品牌修改
export function updateBrand(params) {
  return request({
    url: '/blockchainadmin/srvBrand/update',
    method: 'post',
    params
  })
}

// 品牌新增
export function addBrand(params) {
  return request({
    url: '/blockchainadmin/srvBrand/add',
    method: 'post',
    params
  })
}

//  商品列表
export function commodityList(params) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/commodityList',
    method: 'get',
    params
  })
}

// -------------商品详情-------------

// 根据商品ID, 查商品详情
export function getBrandCommodityInfo(params) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/getBrandCommodityById',
    method: 'get',
    params
  })
}

// 商品修改接口
export function updateCommodity(data) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/updateCommodity',
    method: 'put',
    data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 商品新增接口
export function saveCommodity(data) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/saveCommodity',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 企业信息查询
export function getCompanyInfo(params) {
  return request({
    url: '/blockchainadmin/enterprise/getUserEnterpriseManagement',
    method: 'get',
    params
  })
}

// 品类-查询字典
export function getType(params) {
  return request({
    url: '/blockchainadmin//v1/dict/getDict',
    method: 'get',
    params
  })
}

// 获取农场下拉框
export function getFarmSelect(params) {
  return request({
    url: '/blockchainadmin/comm/v2/farmRelations/query',
    method: 'post',
    params
  })
}

export function getGoods(params) {
  return request({
    url: '/blockchainadmin/comm/v2/commodityByEnterpriseId',
    method: 'post',
    params
  })
}

// 企业信息-修改
export function updateCompanyInfo(params) {
  return request({
    url: '/blockchainadmin/enterprise/updateUserEnterpriseManagement',
    method: 'put',
    params
  })
}

// -------------质检报告-------------
// 获取查询列表
export function getList(params) {
  return request({
    url: '/blockchainadmin/srvQcInspection/list',
    method: 'get',
    params
  })
}

// 上传质检信息
export function reportUploada(data) {
  let params = { ...data, ...getParams() }
  let formData = new FormData()

  for (let key in params) {
    formData.append(key, params[key])
  }

  return request({
    url: '/blockchainadmin/srvQcInspection/upload',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
