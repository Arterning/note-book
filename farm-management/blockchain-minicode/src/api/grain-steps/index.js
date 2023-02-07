/*
 * @Description  : 收粮 -> 列表查询/详情查询/新增提交/修改提交；烘干/加工/包装 -> 列表查询/新增提交
 * @Autor        : Hemingway
 * @Date         : 2021-09-10 11:55:02
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-10-11 21:40:11
 * @FilePath     : \blockchain-minicode\src\api\grain-steps\index.js
 */
import request from '@/http'

// -----------------------------------------mixin-----------------------------------------------

/**
 * @description: 根据用户企业id，获取合作方的企业id合集
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryEnterpriseRelations(data) {
  return await request({
    url: '/blockchainadmin/comm/v2/enterpriseRelations/query',
    method: 'post',
    data
  })
}

/**
 * @description: 调用农场列表查询接口，获取作物品种id合集
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryFarmList(data) {
  return await request({
    url: '/blockchainadmin/comm/v2/farm/query',
    method: 'post',
    data
  })
}

// -----------------------------------------收粮-----------------------------------------------

/**
 * @description: 列表查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryGrainInList(data) {
  return await request({
    url: '/blockchainadmin/v2/feedList/query',
    method: 'get',
    data
  })
}

/**
 * @description: 收粮清单 - 新建提交
 * @param {Object} data
 * @author: Hemingway
 */
export async function addGrainIn(data) {
  return await request({
    url: '/blockchainadmin/v2/feedList/add',
    method: 'post',
    data
  })
}

/**
 * @description: 收粮清单 - 修改提交
 * @param {Object} data
 * @author: Hemingway
 */
export async function updateGrainIn(data) {
  return await request({
    url: '/blockchainadmin/v2/feedList/modify',
    method: 'post',
    data
  })
}

/**
 * @description: 收粮清单 - 详情查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryGrainInDetail(data) {
  return await request({
    url: '/blockchainadmin/v2/feedList/detail',
    method: 'get',
    data
  })
}

// -----------------------------------------烘干-----------------------------------------------

/**
 * @description: 列表查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryDryingList(data) {
  return await request({
    url: '/blockchainadmin/v2/dryInfo/query',
    method: 'get',
    data
  })
}

/**
 * @description: 烘干待选收粮清单列表查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryToDryingList(data) {
  return await request({
    url: '/blockchainadmin/v2/dryInfo/riceDetail',
    method: 'get',
    data
  })
}

/**
 * @description: 烘干表单 - 新建提交
 * @param {Object} data
 * @author: Hemingway
 */
export async function addDrying(data) {
  return await request({
    url: '/blockchainadmin/v2/dryInfo/add',
    method: 'post',
    data
  })
}

// -----------------------------------------加工-----------------------------------------------

/**
 * @description: 列表查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryMachiningList(data) {
  return await request({
    url: '/blockchainadmin/v2/processInfo/query',
    method: 'get',
    data
  })
}

/**
 * @description: 加工表单 - 新建提交
 * @param {Object} data
 * @author: Hemingway
 */
export async function addMachining(data) {
  return await request({
    url: '/blockchainadmin/v2/processInfo/add',
    method: 'post',
    data
  })
}

// -----------------------------------------包装-----------------------------------------------

/**
 * @description: 列表查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryPackingList(data) {
  return await request({
    url: '/blockchainadmin/v2/packingInfo/query',
    method: 'get',
    data
  })
}

/**
 * @description: 根据品牌商企业id，查询商品列表
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryCommodityList(data) {
  return await request({
    url: '/blockchainadmin/v2/packingInfo/commodity/query',
    method: 'get',
    data
  })
}

/**
 * @description: 包装表单 - 新建提交
 * @param {Object} data
 * @author: Hemingway
 */
export async function addPacking(data) {
  return await request({
    url: '/blockchainadmin/v2/packingInfo/add',
    method: 'post',
    data
  })
}
