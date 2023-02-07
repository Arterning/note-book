import { Message } from 'element-ui'

export const throwErr = res => {
  let message
  if (res.code === '20' || res.code === '21') {
    message =
      res.code === '20' ? '暂无用户信息，请登入' : '用户信息失效，请重新登入'

    // do some thing
  } else {
    if (
      res.subErrors &&
      res.subErrors.length !== 0 &&
      res.subErrors[0].message &&
      res.subErrors[0].message.length !== 0 &&
      res.subErrors[0].code
    ) {
      message = res.subErrors[0].message
    } else {
      message = res.msg
    }

    Message({
      message,
      type: 'error',
      duration: 2000
    })
  }

  return Promise.reject(message || '页面发生错误')
}
