import newRequest from '@/utils/newRequest'

// --------------------------------------- 溯源批次查询 ---------------------------------------

/**
 * @description: 企业信息查询
 * @author: qinjj
 */
export function enterpriseInfoList(data) {
  return newRequest({
    url: '/blockchaindbo/commonQuery/enterpriseInfoList',
    method: 'get',
    data
  })
}

/**
 * @description: 商品列表查询
 * @author: qinjj
 */
export function violationCommodityList(data) {
  return newRequest({
    url: '/blockchaindbo/violationCommodity/violationCommodityList',
    method: 'get',
    data
  })
}

/**
 * @description: 禁用接口
 * @author: qinjj
 */
export function updateSyCodeStatus(data) {
  return newRequest({
    url: '/blockchaindbo/violationCommodity/updateSyCodeStatus',
    method: 'put',
    data
  })
}
