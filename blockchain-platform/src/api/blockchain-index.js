import newRequest from '@/utils/newRequest'

// 统计年度上链数据
export function getYearly(data) {
  return newRequest({
    url: '/blockchaindbo/v1/chaincode/statistics/yearly',
    method: 'get',
    data
  })
}

// 统计年度区间上链数据
export function getRange(data) {
  return newRequest({
    url: '/blockchaindbo/v1/chaincode/statistics/range',
    method: 'get',
    data
  })
}

// 查询区间上链数据列表
export function getList(data) {
  return newRequest({
    url: '/blockchaindbo/v1/chaincode/detail/list',
    method: 'get',
    data
  })
}

// 查询区间上链历史数据
export function getHistory(data) {
  return newRequest({
    url: '/blockchaindbo/v1/chaincode/detail/history',
    method: 'get',
    data
  })
}

// 字典
export function getDict(data) {
  return newRequest({
    url: '/blockchaindbo/v1/dict/query',
    method: 'get',
    data
  })
}
