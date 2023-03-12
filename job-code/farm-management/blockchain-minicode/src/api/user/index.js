import request from '@/http'

/**
 * @description: 获取验证码
 * @param {Object} data {phone, smsType} 其中smsType：'00' 注册；'01' 验证码登录；'02' 重置密码
 * @author: Hemingway
 */
export function getMsgCode(data) {
  return request({
    url: '/blockchainauthcenter/verificationCode',
    method: 'post',
    data
  })
}

/**
 * @description: 校验验证码（注册或重置密码）
 * @param {Object} data {phone, smsType, code} 其中smsType：'00' 注册；'01' 验证码登录；'02' 重置密码
 * @author: Hemingway
 */
export function verifyCode(data) {
  return request({
    url: '/blockchainauthcenter/verificationCodeCheck',
    method: 'post',
    data
  })
}

/**
 * @description: 获取登录密码加密公钥
 * @author: Hemingway
 */
export async function getPublicKey() {
  return await request({
    url: '/blockchainauthcenter/publicKey',
    method: 'get'
  })
}

/**
 * @description: 注册
 * @param {Object} data
 * @author: Hemingway
 */
export async function register(data) {
  return await request({
    url: '/blockchainadmin/portal/register/w/v1',
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
 * @description: 账号登录
 * @param {Object} data
 * @author: Hemingway
 */
export async function accountLogin(data) {
  return await request({
    url: '/blockchainadmin/portal/accountLogin/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 短信登录
 * @param {Object} data
 * @author: Hemingway
 */
export async function smsLogin(data) {
  return await request({
    url: '/blockchainadmin/portal/smsLogin/w/v1',
    method: 'post',
    data
  })
}

/**
 * @description: 获取用户信息
 * @author: Hemingway
 */
export async function getUserInfo() {
  return await request({
    url: '/blockchainadmin/user/v2/userInfo/query', // 新接口
    method: 'get'
  })
}

/**
 * @description: 获取用户权限树
 * @author: Hemingway
 */
export async function getUserPermission(data) {
  return await request({
    url: '/blockchainadmin/user/v2/menu/query', // 新接口
    method: 'post',
    data
  })
}

/**
 * @description: 获取数据字典
 * @author: Hemingway
 */
export async function queryDict() {
  return await request({
    url: '/blockchainadmin/v1/dict/query',
    method: 'get',
    isHideLoading: true
  })
}

// ------------------------------------------------入驻--------------------------------------------------

/**
 * @description: 提交入驻申请
 * @param {Object} data
 * @author: Hemingway
 */
export async function submitApplication(data) {
  return await request({
    url: '/blockchainadmin/v2/enterpriseChange/addOrUpdate',
    method: 'post',
    data
  })
}

/**
 * @description: 审核进度列表查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryApplicationList(data) {
  return await request({
    url: '/blockchainadmin/v2/enterpriseChange/query',
    method: 'post',
    data,
    isHideLoading: true
  })
}

/**
 * @description: 审核进度详情查询
 * @param {Object} data
 * @author: Hemingway
 */
export async function queryApplicationDetail(data) {
  return await request({
    url: '/blockchainadmin/v2/enterpriseChange/detail',
    method: 'post',
    data
  })
}

// ------------------------------------------------消息--------------------------------------------------

/**
 * @description: 查询当前用户未读消息数量
 * @author: Hemingway
 */
export async function queryUnreadMessageCount() {
  return await request({
    url: '/blockchaincontent/v1/messageInfo/queryUnreadMessage',
    method: 'post',
    isHideLoading: true
  })
}

/**
 * @description: 查询消息列表
 * @author: Hemingway
 */
export async function queryMessageList(data) {
  return await request({
    url: '/blockchaincontent/v1/messageInfo/query',
    method: 'post',
    data
  })
}

/**
 * @description: 修改消息已读状态
 * @author: Hemingway
 */
export async function readMessageInfo(data) {
  return await request({
    url: '/blockchaincontent/v1/messageInfo/readMessageInfo',
    method: 'post',
    data
  })
}
// ------------------------------------------------素材--------------------------------------------------
/**
 * @description: 新增/修改素材
 * @param {*}
 * @return {*}
 * @author: WuJing
 * @example:
 */
export async function updateMaterial(data) {
  return await request({
    url: '/blockchainadmin/v2/h5/material/add',
    method: 'post',
    data
  })
}
/**
 * @description: 查询商品素材
 * @param {*}
 * @return {*}
 * @author: WuJing
 * @example:
 */
export async function queryMaterial(data) {
  return await request({
    url: '/blockchainadmin/v2/h5/material/query',
    method: 'get',
    data
  })
}
