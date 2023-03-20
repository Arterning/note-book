/*
 * @Description  : 常用工具函数
 * @Autor        : Hemingway
 * @Date         : 2021-05-13 22:01:46
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 15:14:44
 * @FilePath     : \blockchain-minicode\src\utils\tool.js
 */

/**
 * @description: 格式化日期为指定字符串类型
 * @param {Date | Number | String} time 时间：日期类型或时间戳
 * @param {String} cFormat
 * @return {String}
 * @author: Hemingway
 */
export function stringifyDate(time, cFormat) {
  if (arguments.length === 0) {
    return null
  }
  const format = cFormat || '{y}-{m}-{d}'
  let date
  if (typeof time === 'object') {
    date = time
  } else {
    if (isNaN(Date.parse(time))) {
      date = new Date(time.replace(/-/g, '/')) // 兼容iOS
    } else {
      date = new Date(time)
    }
  }
  const formatObj = {
    y: date.getFullYear(),
    m: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    i: date.getMinutes(),
    s: date.getSeconds(),
    a: date.getDay()
  }
  const time_str = format.replace(/{(y|m|d|h|i|s|a)+}/g, (result, key) => {
    let value = formatObj[key]
    if (key === 'a')
      return ['一', '二', '三', '四', '五', '六', '日'][value - 1]
    if (result.length > 0 && value < 10) {
      value = '0' + value
    }
    return value || 0
  })
  return time_str
}

/**
 * @description: 扫码
 * @return {promise}
 * @author: guoxi
 */
export function scanCode() {
  let timer = null
  const promise = new Promise((resolve, reject) => {
    uni.scanCode({
      onlyFromCamera: true,
      success: function(res) {
        resolve(res)
      },
      complete: function() {
        uni.showLoading({
          title: '识别中...'
        })
        timer = setTimeout(() => {
          uni.hideLoading()
          clearTimeout(timer)
        }, 500)
      },
      fail: function(res) {
        reject(res)
        timer = setTimeout(() => {
          uni.hideLoading()
          clearTimeout(timer)
        })
      }
    })
  })
  return promise
}

/**
 * @description: 格式化日期为指定字符串类型
 * @param {Number} section 年限区间
 * @param {String} type 类型：开始 结束
 * @return {Object}
 * @author: jiagui_chen
 */
export function getDateFormat(type, section) {
  const date = new Date()

  let year = date.getFullYear()
  let month = date.getMonth() + 1
  let day = date.getDate()

  if (type === 'start') {
    year = year - section
  } else if (type === 'end') {
    year = year + section
  }
  month = month > 9 ? month : '0' + month
  day = day > 9 ? day : '0' + day

  return `${year}-${month}-${day}`
}

/**
 * @description: uni绘制canvas路径
 * @return {promise}
 * @author: guoxi
 */
export function canvasToTempFilePath(params) {
  const promise = new Promise((resolve, reject) => {
    uni.canvasToTempFilePath({
      canvasId: params.canvasId,
      destWidth: params.destWidth, //分享图片尺寸=画布尺寸1*缩放比0.5*像素比2
      destHeight: params.destHeight,
      quality: params.quality,
      fileType: params.fileType,
      canvas: params.canvas,
      success: function(res) {
        resolve(res)
      },
      fail: function(error) {
        reject(error)
      }
    })
  })
  return promise
}

/**
 * @description: uni获取系统信息
 * @return {promise}
 * @author: guoxi
 */
export function getSystemInfo() {
  return new Promise((resolve, reject) => {
    uni.getSystemInfo({
      success: function(res) {
        resolve(res)
      },
      fail: function(error) {
        reject(error)
      }
    })
  })
}

export function chooseImg(formData, imgMaxNum, imgPaths, that) {
  uni.chooseImage({
    count: 4, //默认9
    sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
    sourceType: ['album', 'camera'], //从相册选择和拍照
    success: res => {
      const sessionId = uni.getStorageSync('sessionId')
      console.log('res:', res)
      let tempFilePaths = res.tempFilePaths
      if (formData.images.length + tempFilePaths.length > imgMaxNum) {
        uni.showToast({
          title: '超过上传图片的最大数量',
          icon: 'none'
        })
      } else {
        imgPaths = [...imgPaths, ...tempFilePaths]
        uni.showLoading({
          title: '加载中...',
          mask: true
        })
        for (var i = 0; i < tempFilePaths.length; i++) {
          uni.uploadFile({
            url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/uploadFile/w/v1`,
            method: 'POST',
            filePath: tempFilePaths[i],
            name: 'file',
            formData: {
              fileClass: 'blockchain',
              sessionId: sessionId,
              clientId: 'poweb',
              isOpen: 'Y'
            },
            success: res => {
              const id = JSON.parse(res.data).id
              formData.ids && formData.ids.push({ id })
              formData.images.push(
                // `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${sessionId}&clientId=poweb`

                `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v2?id=${id}&isAbbreviation=N`
              )
              // 强制刷新
              that.$forceUpdate()
            }
          })
        }
        uni.hideLoading()
      }
    }
  })
}

export function preViewImg(i, imgPaths) {
  uni.previewImage({
    current: i,
    urls: imgPaths,
    loop: true,
    success: result => {
      console.log(result)
    },
    fail: error => {
      console.log(error)
    }
  })
}

export function getImageUrl(id) {
  // const sessionId = uni.getStorageSync('sessionId')
  // return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${sessionId}&clientId=poweb`
  return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v2?id=${id}&isAbbreviation=Y`
}

export function getZNImageUrl(id) {
  // const sessionId = uni.getStorageSync('sessionId')
  // return `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${id}&isAbbreviation=N&sessionId=${sessionId}&clientId=poweb`
  return `${process.env.VUE_APP_BASE_URL}/agriculturecomponent/file/downloadFile/w/v2?id=${id}&isAbbreviation=N`
}

export function onlyNumOnePoint(number_only) {
  //先把非数字的都替换掉，除了数字和小数点
  number_only = number_only.replace(/[^\d.]/g, '')
  //必须保证第一个为数字而不是小数点
  number_only = number_only.replace(/^\./g, '')
  //保证只有出现一个小数点而没有多个小数点
  number_only = number_only.replace(/\.{2,}/g, '.')
  //保证小数点只出现一次，而不能出现两次以上
  number_only = number_only
    .replace('.', '$#$')
    .replace(/\./g, '')
    .replace('$#$', '.')
  //保证只能输入两个小数
  // number_only = number_only.replace(/^(\-)*(\d+)\.(\d\d).*$/, '$1$2.$3')
  return number_only
}

/**
 * 生成uuid方法
 * @returns {string}
 */
export const createUUID = function() {
  var d = new Date().getTime()
  // if (window.performance && typeof window.performance.now === 'function') {
  //   d += performance.now() // use high-precision timer if available
  // }
  var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(
    c
  ) {
    var r = (d + Math.random() * 16) % 16 | 0
    d = Math.floor(d / 16)
    return (c === 'x' ? r : (r & 0x3) | 0x8).toString(16)
  })
  return uuid
}
