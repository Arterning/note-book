import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import qs from 'qs'
import md5 from 'js-md5'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
    Accept: 'application/json'
  }
})

service.interceptors.request.use(
  config => {
    // 文件上传
    if (config.isUpload) {
      config.headers['Content-Type'] = 'multipart/form-data'
    }
    // 文件下载
    if (config.isDownload) {
      config.responseType = 'blob'
    }

    // 原始请求参数组装
    handleParams(config)

    // 原始请求参数签名
    signParams(config.data)

    // 请求参数
    if (config.method === 'post') {
      if (
        config.data &&
        config.headers['Content-Type'] === 'application/x-www-form-urlencoded'
      ) {
        config.data = qs.stringify(config.data, {
          arrayFormat: 'indices',
          allowDots: true
        })
      }
    } else {
      if (config.data) {
        config.url = config.url + '?' + qs.stringify(config.data)
      }
    }

    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  async response => {
    const res = response.data

    if (response.config.responseType === 'blob') {
      const checkRes = await _fileToJson(res)
      console.log('blob checkRes: ', res, checkRes)
      if (!checkRes.message) {
        Message({
          message: checkRes.msg || 'Error',
          type: 'error',
          duration: 5 * 1000
        })

        return Promise.reject(new Error(checkRes.msg || 'Error'))
      }

      if (response.config.isDownload) {
        const objUrl = window.URL.createObjectURL(res)

        const aLink = document.createElement('a')
        aLink.style.display = 'none'
        aLink.href = objUrl
        aLink.setAttribute('download', response.config.fileName)
        document.body.appendChild(aLink)
        aLink.click()

        document.body.removeChild(aLink) // 下载完成移除元素
        window.URL.revokeObjectURL(objUrl) // 释放掉blob对象

        return true
      } else {
        return res
      }
    }

    if (res.code !== '0' && res.code !== 0) {
      // '20': 缺少 sessionId 参数; '21': 无效的 sessionId 参数
      if (res.code === '20' || res.code === '21') {
        // to re-login
        MessageBox.confirm(
          '您已注销，您可以取消留在此页面，或重新登录',
          '确认注销',
          {
            confirmButtonText: '重新登录',
            cancelButtonText: '取消',
            type: 'warning'
          }
        ).then(() => {
          store.dispatch('user/resetToken').then(() => {
            location.reload()
          })
        })
      } else {
        let message = res.msg

        if (
          res.subErrors &&
          res.subErrors.length > 0 &&
          res.subErrors[0].message
        ) {
          message = res.subErrors[0].message
        }

        Message({
          message: message || 'Error',
          type: 'error',
          duration: 5 * 1000
        })

        return Promise.reject(new Error(message || 'Error'))
      }
    } else {
      return res
    }
  },
  error => {
    console.log(error)
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

/**
 * @description: 处理原始请求参数
 * @param {Object} config
 * @author: Hemingway
 */
function handleParams(config) {
  // 防止不传参数的情况下，data 为 undefined
  if (!config.data) {
    config.data = {}
  }

  const { data, isUpload, noLogin } = config
  const { sessionId, locale, clientId } = getParams()

  if (isUpload) {
    // 文件上传

    if (!noLogin && sessionId) {
      data.append('sessionId', sessionId)
    }
    data.append('locale', locale)
    data.append('clientId', clientId)
  } else {
    // 非文件上传

    if (!noLogin && sessionId) {
      data['sessionId'] = sessionId
    }
    data['locale'] = locale
    data['clientId'] = clientId
  }
}

/**
 * @description: 获取公共必要请求参数（根据项目情况修改）
 * @author: Hemingway
 */
export function getParams() {
  let sessionId = ''
  if (store.getters.token) {
    sessionId = getToken()
  }

  const locale = 'zh_CN'
  const clientId = 'poweb'

  return {
    sessionId,
    locale,
    clientId
  }
}

/**
 * @description: 参数签名
 * @param {Object} data
 * @author: Hemingway
 */
export function signParams(data) {
  const arr = qs
    .stringify(data, {
      encode: false,
      delimiter: '#%-&^#-W&',
      arrayFormat: 'indices',
      allowDots: true
    })
    .split('#%-&^#-W&')
  const timestamp = new Date().getTime()
  const KEY = '/l/A/8/9/2/w/b/b/s/B/S/N/f/r/5/a/J/2/c/o/V/T/u/T/a/3/d/Q/t'
  const paramStr = KEY + arr.sort().join('&') + timestamp + KEY

  data['timestamp'] = timestamp
  data['sign'] = md5(paramStr)

  return data
}

// 将blob对象转化为json（文件类型调用ajax 取后端的返回值做特殊处理）
function _fileToJson(file) {
  let data = {},
    message = ''

  return new Promise(resolve => {
    const reader = new FileReader()

    // 成功回调
    reader.onload = res => {
      const { result } = res.target // 得到字符串
      try {
        // 解析成json对象
        data = JSON.parse(result)
      } catch (err) {
        message = err.message || err
      }

      resolve({ data, message })
    }

    // 失败回调
    reader.onerror = err => {
      message = err.message || err
      resolve({ data, message })
    }

    reader.readAsText(new Blob([file]), 'utf-8') // 按照utf-8编码解析
  })
}

export default service
