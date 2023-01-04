import store from '@/store'

export const throwErr = res => {
  let message
  if (res.code === '20' || res.code === '21') {
    message =
      res.code === '20' ? '暂无用户信息，请登入' : '用户信息失效，请重新登入'

    uni.showToast({
      title: message,
      icon: 'none',
      duration: 2000,
      complete: () => {
        store.commit('setSessionId', '')
        store.commit('setPermissions', null)
        store.commit('setUserInfo', null)
        uni.removeStorageSync('sessionId')
        uni.reLaunch({ url: '/pages/login/index' })
      }
    })
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

    uni.showToast({
      icon: 'none',
      title: message,
      duration: 3000
    })
  }

  return Promise.reject(message || '页面发生错误')
}
