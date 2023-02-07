import request from '@/http'

/**
 * @description: 获取员工列表
 * @param {Object} data
 * @author: jiagui_chen
 */
export function getUserList(data) {
  return request({
    url: '/blockchainadmin/sysUser/list',
    method: 'post',
    data
  })
}

/**
 * @description: 获取员工类型下拉框数据
 * @param {Object}
 * @author: jiagui_chen
 */
export function getRoleList() {
  return request({
    url: '/blockchainadmin/sysDict/list',
    method: 'post',
    data: {
      dicType: 'account_type'
    }
  })
}

/**
 * @description: 员工新增
 * @param {Object} data
 * @author: jiagui_chen
 */
export function userAdd(data) {
  return request({
    url: '/blockchainadmin/sysUser/add',
    method: 'post',

    data
  })
}

/**
 * @description: 员工编辑
 * @param {Object} data
 * @author: jiagui_chen
 */
export function updateUser(data) {
  return request({
    url: '/blockchainadmin/sysUser/update',
    method: 'post',

    data
  })
}

/**
 * @description: 员工删除
 * @param {Object} id 员工id
 * @author: jiagui_chen
 */
export function delUsers(id) {
  return request({
    url: `/blockchainadmin/sysUser/delete`,
    method: 'get',
    data: { id },
    headers: {
      'Content-Type': 'application/json',
      sessionId: uni.getStorageSync('sessionId')
    }
  })
}

/**
 * @description: 筛选员工
 * @param {Object} id 员工id
 * @author: guoxi
 */
export function filterUser(name) {
  return request({
    url: `/blockchainadmin/sysUser/userList`,
    method: 'post',
    data: { userNameLike: name }
  })
}
