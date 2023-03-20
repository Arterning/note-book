import request from '@/http'

/**
 * @description: 获取品牌列表
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getBrandList(data) {
  return request({
    url: '/blockchainadmin/srvBrand/list',
    method: 'post',

    data
  })
}

/**
 * @description: 新增品牌
 * @param {Object} data
 * @author: jiagui_chen
 */
export function addBrand(data) {
  return request({
    url: '/blockchainadmin/srvBrand/add',
    method: 'post',

    data
  })
}

/**
 * @description: 删除品牌
 * @param {Number} id
 * @author: jiagui_chen
 */
export function delBrand(id) {
  return request({
    url: `/blockchainadmin/srvBrand/delete?id=${id}`,
    method: 'post',
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    }
  })
}

/**
 * @description: 更新品牌
 * @param {Object} data
 * @author: jiagui_chen
 */
export function updateBrand(data) {
  return request({
    url: '/blockchainadmin/srvBrand/update',
    method: 'post',

    data
  })
}

/**
 * @description: 获取商品列表
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getCommodityList(data) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/list',
    method: 'post',

    data
  })
}

/**
 * @description: 新增商品
 * @param {Object} data
 * @author: jiagui_chen
 */
export function addCommodity(data) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/add',
    method: 'post',

    data
  })
}

/**
 * @description: 删除商品
 * @param {Number} id
 * @author: jiagui_chen
 */
export function delCommodity(id) {
  return request({
    url: `/blockchainadmin/srvBrandCommodity/delete?id=${id}`,
    method: 'post'
  })
}

/**
 * @description: 更新商品
 * @param {Object} data
 * @author: jiagui_chen
 */
export function updateCommodity(data) {
  return request({
    url: '/blockchainadmin/srvBrandCommodity/update',
    method: 'post',

    data
  })
}

/**
 * @description: 查看商品详情数据
 * @param {Number} id 商品id
 * @author: jiagui_chen
 */
export function detailCommodity(id) {
  return request({
    url: `/blockchainadmin/srvBrandCommodity/find?id=${id}`,
    method: 'post',
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    }
  })
}
