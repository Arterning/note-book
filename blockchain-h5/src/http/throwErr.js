export const throwErr = res => {
  let message
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

  uni.showToast({
    icon: 'none',
    title: message,
    duration: 3000
  })

  return Promise.reject(message || '页面发生错误')
}
