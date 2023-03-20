import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import qs from 'qs'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['sessionId'] = getToken()
    }

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

    // 请求参数
    if (config.method === 'post') {
      if (
        config.data &&
        config.headers['Content-Type'] === 'application/json'
      ) {
        config.data = JSON.stringify(config.data)
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

  const { data, isUpload, needAppType } = config
  const locale = 'zh_CN'
  const clientId = 'poweb'

  if (isUpload) {
    // 文件上传

    data.append('locale', locale)
    data.append('clientId', clientId)
  } else {
    // 非文件上传

    data['locale'] = locale
    data['clientId'] = clientId
    needAppType ? (data['appType'] = '30') : null // 只有登录接口需要传appType
  }
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
