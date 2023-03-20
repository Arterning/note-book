/**
 * Created by PanJiaChen on 16/11/18.
 */

/**
 * @param {string} path
 * @returns {Boolean}
 */
export function isExternal(path) {
  return /^(https?:|mailto:|tel:)/.test(path)
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validUsername(str) {
  const valid_map = ['admin', 'editor']
  return valid_map.indexOf(str.trim()) >= 0
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validPhoneNumber(str) {
  let reg = /^1[3456789]\d{9}$/
  return reg.test(str)
}

/**
 * @param {Object} rule
 * @param {string} value
 * @param {Function} callback
 * @returns {Promise}
 */
export function validPassword(value) {
  return new Promise(function (resolve, reject) {
    if (value === '') {
      reject('请输入密码')
    } else {
      let lv = 0 // 密码强度等级
      if (value.length >= 8 && value.length <= 20) {
        if (value.match(/[a-z]/g)) {
          lv++
        } // 含有小写字母
        if (value.match(/[A-Z]/g)) {
          lv++
        } // 含有大写字母
        if (value.match(/[0-9]/g)) {
          lv++
        } // 含有数字
        if (value.match(/[-_,!|~`()#$%^&*{}:;"L<>?]/g)) {
          lv++
        } // 含有特殊字符
        if (lv >= 2) {
          // 满足其中两项时,查询密码强度
          resolve()
        } else {
          reject('密码为8-20位，数字/字母/符号两种及以上组合')
        }
      } else {
        reject('请输入长度为 8 ~ 20位密码')
      }
    }
  })
}

/**
 * @param {Object} rule
 * @param {string} value
 * @param {Function} callback
 * @returns {Promise}
 */
export function checkPasswordLevel(rule, value, callback) {
  return new Promise(function (resolve) {
    if (value === '') {
      callback('请输入密码')
    } else {
      let lv = 0 // 密码强度等级
      if (value.length >= 8 && value.length <= 20) {
        if (value.match(/[a-z]/g)) {
          lv++
        } // 含有小写字母
        if (value.match(/[A-Z]/g)) {
          lv++
        } // 含有大写字母
        if (value.match(/[0-9]/g)) {
          lv++
        } // 含有数字
        if (value.match(/[-_,!|~`()#$%^&*{}:;"L<>?]/g)) {
          lv++
        } // 含有特殊字符
        if (lv >= 2) {
          // 满足其中两项时,查询密码强度
          callback()
          resolve()
        } else {
          callback('密码为8-20位，数字/字母/符号两种及以上组合')
        }
      } else {
        callback(new Error('请输入长度为 8 ~ 20位密码'))
      }
      callback()
    }
  })
}

/**
 * @param {string} url
 * @returns {Boolean}
 */
export function validURL(url) {
  const reg =
    /^(https?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/
  return reg.test(url)
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validLowerCase(str) {
  const reg = /^[a-z]+$/
  return reg.test(str)
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validUpperCase(str) {
  const reg = /^[A-Z]+$/
  return reg.test(str)
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function validAlphabets(str) {
  const reg = /^[A-Za-z]+$/
  return reg.test(str)
}

/**
 * @param {string} email
 * @returns {Boolean}
 */
export function validEmail(email) {
  const reg =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
  return reg.test(email)
}

/**
 * @param {string} str
 * @returns {Boolean}
 */
export function isString(str) {
  if (typeof str === 'string' || str instanceof String) {
    return true
  }
  return false
}

/**
 * @param {Array} arg
 * @returns {Boolean}
 */
export function isArray(arg) {
  if (typeof Array.isArray === 'undefined') {
    return Object.prototype.toString.call(arg) === '[object Array]'
  }
  return Array.isArray(arg)
}
