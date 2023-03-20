import request from '@/http'

/**
 * 查询农场地块
 * @param {参数对象} payload
 */
export function QuerySectionSByFarm(payload) {
  let url = '/agriculturegis/v1/Section/QuerySectionSByFarm'
  if (process.env.VUE_APP_URL_PREFIX) {
    url = process.env.VUE_APP_URL_PREFIX + url
  }
  return request({
    url: url,
    method: 'get',
    data: payload
  })
}
