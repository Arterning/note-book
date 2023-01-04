/*
 * @Description  : 网络请求封装
 * config选项有：url 接口, method 方法, data 参数; isHideLoading 是否隐藏加载动画, isFile 是否属于文件上传, 以及isJson
 * @Autor        : Hemingway
 * @Date         : 2021-05-13 23:01:49
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-18 17:08:10
 * @FilePath     : \blockchain-minicode\src\http\index.js
 */

import axios from 'axios'
import qs from 'qs'
import md5 from 'js-md5'
import store from '@/store'
import { throwErr } from './throwErr'

// 桥接器
const adapter = function(config) {
  return new Promise((resolve, reject) => {
    const settle = require('axios/lib/core/settle')
    const { baseURL, url, data, headers, method, timeout } = config
    uni.request({
      url: baseURL + url,
      data: data,
      header: headers,
      method: method.toUpperCase(),
      timeout: timeout,
      complete: function complete(response) {
        const { statusCode, errMsg, data, header } = response
        settle(resolve, reject, {
          config: config,
          status: statusCode,
          statusText: errMsg,
          data: data,
          headers: header
        })
      }
    })
  })
}

// 创建axios实例
const service = axios.create({
  adapter,
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 20000,
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
    Accept: 'application/json'
  }
})

/**
 * @description: request拦截器
 * @author: Hemingway
 */
service.interceptors.request.use(
  config => {
    // 是否展示“加载中”（默认展示加载动画，不需要时添加该配置）
    if (!config.isHideLoading) {
      uni.showLoading({
        title: '加载中...',
        mask: true
      })
    }

    // 上传文件的请求，修改header
    if (config.isFile) {
      config.headers['Content-Type'] = 'multipart/form-data'
    }

    // 原始请求参数组装
    handleParams(config)

    // 为参数加签名
    signParams(config.data)

    if (config.method === 'post') {
      if (
        config.data &&
        config.headers['Content-Type'] === 'application/x-www-form-urlencoded'
      ) {
        if (config.isJson) {
          config.data = qs.stringify(config.data)
        } else {
          // 格式化数组参数
          config.data = qs.stringify(config.data, {
            arrayFormat: 'indices',
            allowDots: true
          })
        }
      }
    } else {
      if (config.data) {
        config.url = config.url + '?' + qs.stringify(config.data)
      }
    }

    console.log('网络请求 config = ', config)
    return config
  },
  error => {
    Promise.reject(error)
  }
)

/**
 * @description: respone拦截器
 * @author: Hemingway
 */
service.interceptors.response.use(
  response => {
    // 请求响应后关闭加载框
    if (!response.config.isHideLoading) {
      uni.hideLoading()
    }
    console.log('网络响应 response = ', response)

    if (response.status === 200) {
      const res = response.data
      if (res.code === '0' || res.code === 0 || res.length > 0) {
        return Promise.resolve(res)
      } else {
        return throwErr(res)
      }
    } else {
      return Promise.reject('networkRequestError')
    }
  },
  error => {
    // 请求响应后关闭加载框
    uni.hideLoading()

    console.log('网络请求 error = ', error)

    // 超时 或 断网
    if (!error.response) {
      if (error.message.includes('timeout')) {
        uni.showModal({
          title: 'error',
          content: '请求超时，请检查网络',
          showCancel: false
        })
      } else {
        uni.showModal({
          title: 'error',
          content: '请求失败，请检查网络',
          showCancel: false
        })
      }
    }

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

  const { data, isFile, noLogin } = config
  const { sessionId, appType, locale, clientId } = getParams()

  if (isFile) {
    // 上传文件

    if (!noLogin && sessionId) {
      data.append('sessionId', sessionId)
    }
    data.append('appType', appType)
    data.append('locale', locale)
    data.append('clientId', clientId)
  } else {
    // 非文件上传

    if (!noLogin && sessionId) {
      data['sessionId'] = sessionId
    }
    data['appType'] = appType
    data['locale'] = locale
    data['clientId'] = clientId
  }
}

/**
 * @description: 获取公共必要请求参数
 * @author: Hemingway
 */
function getParams() {
  const userInfo = store.state.user.userInfo || {}
  const sessionId = uni.getStorageSync('sessionId')
  const appType = userInfo.appType || '1'
  const locale = userInfo.locale || 'zh_CN'
  const clientId = 'poweb'
  return {
    sessionId,
    appType,
    locale,
    clientId
  }
}

/**
 * @description: 参数签名
 * @param {Object} data
 * @author: Hemingway
 */
function signParams(data) {
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

export default service
