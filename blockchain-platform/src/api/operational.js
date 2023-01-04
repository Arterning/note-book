import newRequest from '@/utils/newRequest'

// --------------------------------------- 运营数据报表 ---------------------------------------

/**
 * @description: 品类下拉框的值
 * @author: qinjj
 */
export function getDict(data) {
  return newRequest({
    url: '/blockchaindbo/commonQuery/getDict',
    method: 'get',
    data
  })
}

/**
 * @description: 合作企业（家）
 * @author: qinjj
 */
export function cooperativeEnterpriseNumber(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/cooperativeEnterpriseNumber',
    method: 'get',
    data
  })
}

/**
 * @description: 创建品牌商
 * @author: qinjj
 */
export function brandNumber(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/brandNumber',
    method: 'get',
    data
  })
}

/**
 * @description: 创建溯源码个数
 * @author: qinjj
 */
export function syCodeNumber(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/syCodeNumber',
    method: 'get',
    data
  })
}

/**
 * @description: 绑定溯源码个数
 * @author: qinjj
 */
export function bondSyCodeNumber(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/bondSyCodeNumber',
    method: 'get',
    data
  })
}

/**
 * @description: 扫码个数
 * @author: qinjj
 */
export function scanCodeNumber(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/scanCodeNumber',
    method: 'get',
    data
  })
}

/**
 * @description: 服务模式下拉框
 * @author: qinjj
 */
export function serviceModelList(data) {
  return newRequest({
    url: '/blockchaindbo/commonQuery/serviceModelList',
    method: 'get',
    data
  })
}

/**
 * @description: 商品下拉框
 * @author: qinjj
 */
export function commodityList(data) {
  return newRequest({
    url: '/blockchaindbo/commonQuery/commodityList',
    method: 'get',
    data
  })
}

/**
 * @description: 企业下拉框
 * @author: qinjj
 */
export function enterpriseList(data) {
  return newRequest({
    url: '/blockchaindbo/commonQuery/enterpriseList',
    method: 'get',
    data
  })
}

/**
 * @description: 品牌下拉框
 * @author: qinjj
 */
export function brandList() {
  return newRequest({
    url: '/blockchaindbo/commonQuery/brandList',
    method: 'get'
  })
}

/**
 * @description: 品牌商下拉框
 * @author: qinjj
 */
export function enterpriseRelations() {
  return newRequest({
    url: '/blockchaindbo/commonQuery/enterpriseList',
    method: 'get'
  })
}

/**
 * @description: 合作企业-柱状图
 * @author: qinjj
 */
export function cooperativeEnterpriseBar(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/cooperativeEnterpriseBar',
    method: 'get',
    data
  })
}

/**
 * @description: 消费者已扫码-饼图
 * @params{string}: brandId 品牌id
 * @author: qinjj
 */
export function scanCodePie(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/scanCodePie',
    method: 'get',
    data
  })
}

/**
 * @description: 查询数据表格
 * @author: qinjj
 */
export function dimensionalityList(data) {
  return newRequest({
    url: '/blockchaindbo/dboPlatformReport/dimensionalityList',
    method: 'get',
    data
  })
}
