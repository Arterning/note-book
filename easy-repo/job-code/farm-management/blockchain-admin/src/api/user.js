import request from '@/utils/request'
import newRequest from '@/utils/newRequest'

// --------------------------------------- blockchainauthcenter ---------------------------------------

/**
 * @description: 获取加密公钥（密码登录或重置密码）
 * @author: Hemingway
 */
export function getPublicKey() {
  return request({
    url: '/blockchainadmin/portal/publicKey',
    method: 'get'
  })
}

/**
 * @description: 发送验证码
 * @param {Object} data {phone, smsType}
 * @author: Hemingway
 */
export function sendCode(data) {
  return request({
    url: '/blockchainadmin/portal/verificationCode',
    method: 'post',
    data
  })
}

/**
 * @description: 校验验证码（注册或重置密码）
 * @param {Object} data {phone, smsType, code}
 * @author: Hemingway
 */
export function validateCode(data) {
  return request({
    url: '/blockchainadmin/user/checkVerifyCode/w/v1',
    method: 'post',
    data
  })
}

// --------------------------------------- blockchainadmin ---------------------------------------

/**
 * @description: 账号登录
 * @param {Object} data {username, password}
 * @author: Hemingway
 */
export function accountLogin(data) {
  return request({
    url: '/blockchainadmin/portal/accountLogin/w/v1',
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
    url: '/blockchainadmin/portal/smsLogin/w/1',
    method: 'post',
    data,
    needAppType: true
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
 * @description: 重置密码
 * @param {Object} data
 * @author: Hemingway
 */
export async function resetPassword(data) {
  return await request({
    url: '/blockchainadmin/user/forgetPassword/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 设置密码
 * @param {Object} data {access_token, password}
 * @author: Hemingway
 */
export async function setPassword(data) {
  return await request({
    url: '/blockchainadmin/user/setPassword/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 更改密码
 * @param {Object} data {access_token, oldPassword, updatePassword}
 * @author: Hemingway
 */
export async function updatePassword(data) {
  return await request({
    url: '/blockchainadmin/user/changeUserPassword/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 登出账号
 * @author: Hemingway
 */
export function logout() {
  return request({
    url: '/blockchainadmin/user/logout/w/v1',
    method: 'get'
  })
}

// --------------------------------------- blockchainpub ---------------------------------------

/**
 * @description: 用户是否存在，以及是否需要设置密码（未设置密码时为Y）
 * @param {Object} data {username}
 * @author: Hemingway
 */
export function checkUsername(data) {
  return newRequest({
    url: '/blockchainadmin/portal/needSetPassword',
    method: 'get',
    data
  })
}

/**
 * @description: 获取用户权限
 * @author: Hemingway
 */
export function getUserMenuTree(data) {
  return newRequest({
    url: '/blockchainadmin/v1/role/privilege/list',
    method: 'get',
    data
  })
}

// --------------------------------------- ??? ---------------------------------------

/**
 * @description: 获取用户信息
 * @param {Object} data {phone}
 * @author: Hemingway
 */
export function getUserInfo(data) {
  return newRequest({
    url: '/blockchainfarm/v1/userCenter/basicInfo/query',
    method: 'get',
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
    url: '/blockchainadmin/user/changePhone/w/v1',
    method: 'post',
    data
  })
}

// --------------------------------------- 以下暂未联调 ---------------------------------------

/**
 * @description: 获取用户信息
 * @author: Hemingway
 */
export function getInfo() {
  return request({
    url: '/blockchaingateway/current_user',
    method: 'get'
  })
}
