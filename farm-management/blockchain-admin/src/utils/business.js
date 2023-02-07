import { getToken } from './auth'

// 获取图片URL地址
export function getImgUrl(id) {
  if (process.env.VUE_APP_URL_PREFIX) {
    return `${
      process.env.VUE_APP_BASE_URL + process.env.VUE_APP_URL_PREFIX
    }/agriculturecomponent/file/downloadFile/w/v1?id=${id}&sessionId=${getToken()}&clientId=poweb`
  } else {
    return `${
      process.env.VUE_APP_BASE_URL
    }/agriculturecomponent/file/downloadFile/w/v1?id=${id}&sessionId=${getToken()}&clientId=poweb`
  }
}
