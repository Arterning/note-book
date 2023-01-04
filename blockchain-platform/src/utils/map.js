import { TextToCode } from 'element-china-area-data'

/**
 * @description: 省市区名称匹配，获取省市区编码
 * @param {String} province
 * @param {String} city
 * @param {String} district
 * @return {Object} {provinceCode, cityCode, districtCode}
 * @author: Hemingway
 */
export function getAddressCode(province, city, district) {
  let provinceCode,
    cityCode,
    districtCode = ''

  provinceCode = TextToCode[province].code
  cityCode = TextToCode[province][city].code
  districtCode = TextToCode[province][city][district].code

  return { provinceCode, cityCode, districtCode }
}
