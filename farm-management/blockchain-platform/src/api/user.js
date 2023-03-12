import request from '@/utils/request'
import newRequest from '@/utils/newRequest'

import { getToken } from '@/utils/auth'

/**
 * @description: 获取加密公钥（密码登录或重置密码）
 * @author: Hemingway
 */
export function getPublicKey() {
  return request({
    url: '/blockchaindbo/adminUser/publicKey',
    method: 'get'
  })
}
/**
 * @description: 账号登录
 * @param {Object} data {username, password}
 * @author: Hemingway
 */
export function accountLogin(data) {
  return request({
    url: '/blockchaindbo/adminUser/accountLogin/w/v1',
    method: 'post',
    data,
    needAppType: true
  })
}

/**
 * @description: 短信登录
 * @param {Object} data {phone, code}
 * @author: Hemingway
 */
export function smsLogin(data) {
  return request({
    url: '/blockchaindbo/adminUser/smsLogin/w/v1',
    method: 'post',
    data,
    needAppType: true
  })
}

// export function login(data) {
//   return request({
//     url: '/blockchaindbo/adminUser/accountLogin/w/v1',
//     method: 'post',
//     data
//   })
// }

/**
 * @description: 设置密码
 * @param {Object} data {access_token, password}
 * @author: Hemingway
 */
export async function setPassword(data) {
  return await request({
    url: '/blockchaindbo/adminUser/setPassword/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 发送验证码
 * @param {Object} data {phone}
 * @author: Hemingway
 */
export function sendCode(data) {
  return request({
    url: '/blockchaindbo/adminUser/sendVerifyCode/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 校验验证码
 * @param {Object} data {phone, code}
 * @author: Hemingway
 */
export function validateCode(data) {
  return request({
    url: '/blockchaindbo/adminUser/codeCheck/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 验证用户手机号是否有效
 * @param {Object} data {phone}
 * @author: Hemingway
 */
export function validatePhone(data) {
  return request({
    url: 'blockchainadmin/user/validatePhone/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 修改手机号
 * @param {Object} data {phone，oldCertificate, certificate, access_token}
 * @author: Hemingway
 */
export function updatePhone(data) {
  return request({
    url: '/blockchaindbo/adminUser/modifyPhone/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 更改密码
 * @param {Object} data {access_token, newPassword, oldPassword}
 * @author: Hemingway
 */
export async function updatePassword(data) {
  return await request({
    url: '/blockchaindbo/adminUser/changePassword/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 用户是否存在，以及是否需要设置密码（未设置密码时为Y）
 * @param {Object} data {username}
 * @author: Hemingway
 */
export function checkUsername(data) {
  return newRequest({
    url: '/blockchaindbo/adminUser/account/exists',
    method: 'get',
    data
  })
}

/**
 * @description: 获取用户信息
 * @author: Hemingway
 */
export function getUserInfo() {
  return request({
    url: '/agriculturedbo/adminUser/getUserInfo/w/v1',
    method: 'post'
  })
}

export function getUserMenuTree(data) {
  return request({
    url: '/blockchaindbo/v1/role/management/function/list',
    method: 'get',
    data,
    headers: {
      'Content-Type': 'application/json',
      sessionId: getToken()
    }
  })
}

export function logout() {
  return request({
    url: '/blockchaindbo/adminUser/logout/w/v1',
    method: 'get'
  })
}
