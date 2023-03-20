// import newRequest from '@/utils/newRequest'
import request from '@/utils/request'
// import { getToken } from '@/utils/auth'

const headers = {
  'Content-Type': 'application/json',
  Accept: 'application/json'
  // sessionId: getToken()
}

// --------------------------------------- 角色管理 ---------------------------------------

/**
 * @description: 获取当前用户所在查询一级组织下的角色信息
 * @param {Object} data searchKey
 * @author: Hemingway
 */
export function getRole(data) {
  return request({
    url: '/blockchainadmin/v1/role/list',
    method: 'get',
    data,
    headers
  })
}

/**
 * @description: 创建角色
 * @param {Object} data 名称、启用状态、备注
 * @author: Hemingway
 */
export function addRole(data) {
  return request({
    url: '/blockchainadmin/v1/role/create',
    method: 'post',
    data,
    headers
  })
}

/**
 * @description: 编辑角色
 * @param {Object} data id、名称、启用状态、备注
 * @author: Hemingway
 */
export function editRole(data) {
  return request({
    url: '/blockchainadmin/v1/role/update',
    method: 'post',
    data,
    headers
  })
}

/**
 * @description: 删除角色
 * @param {Object} data id
 * @author: Hemingway
 */
export function deleteRole(data) {
  return request({
    url: '/blockchainadmin/v1/role/delete',
    method: 'post',
    data,
    headers
  })
}

/**
 * @description: 获取角色菜单功能列表
 * @param {Object} data roleId
 * @author: Hemingway
 */
export function getRoleRight(data) {
  return request({
    url: '/blockchainadmin/v1/role/privilege/list',
    method: 'get',
    data,
    headers
  })
}

/**
 * @description: 获取角色人员列表
 * @param {Object} data roleId,pageSize,pageNum
 * @author: Hemingway
 */
export function getRoleUser(data) {
  return request({
    url: '/blockchainadmin/v1/role/user/list',
    method: 'post',
    data,
    headers
  })
}

/**
 * @description: 设置指定角色的权限
 * @param {Object} data roleId、dataPermissionType、checkedMenuIdList、checkedFunctionList
 * @author: Hemingway
 */
export function setRoleRight(data) {
  return request({
    url: '/agriculturefarm/v1/role/setRoleRights',
    method: 'post',
    data,
    headers
  })
}

// 绑定/更新角色菜单功能权限(提交)
export function submitRoleRight(data) {
  return request({
    url: '/blockchainadmin/v1/role/privilege/bond',
    method: 'post',
    data,
    headers
  })
}

// --------------------------------------- 员工管理 ---------------------------------------

// /**
//  * @description: 获取员工列表
//  * @param {Object} data keywords
//  * @author: Hemingway
//  */
export function getStaffList(params) {
  return request({
    url: '/blockchainadmin/v1/staff/list',
    method: 'get',
    params,
    headers
  })
}

// /**
//  * @description: 新建员工
//  * @param {Object} data
//  * @author: Hemingway
//  */
export function createStaff(data) {
  return request({
    url: '/blockchainadmin/v1/staff/create',
    method: 'post',
    data,
    headers
  })
}

// /**
//  * @description: 编辑员工
//  * @param {Object} data
//  * @author: Hemingway
//  */
export function editStaff(data) {
  return request({
    url: '/blockchainadmin/v1/staff/modify',
    method: 'post',
    data,
    headers
  })
}

// /**
//  * @description: 删除员工
//  * @param {Object} data {userId}
//  * @author: Hemingway
//  */
export function deleteStaff(data) {
  return request({
    url: '/blockchainadmin/v1/staff/delete',
    method: 'post',
    data,
    headers
  })
}

// /**
//  * @description: 冻结/解冻员工账号
//  * @param {Object} data {userId}
//  * @author: Hemingway
//  */
export function operateAccount(data) {
  return request({
    url: '/blockchainadmin/v1/staff/account/operate',
    method: 'post',
    data,
    headers
  })
}

// /**
//  * @description: 获取员工列表
//  * @param {Object} data keywords
//  * @author: Hemingway
//  */
// export function getUser(data) {
//   return newRequest({
//     url: '/agriculturefarm/v1/farmManagement/farmUser/list',
//     method: 'post',
//     data
//   })
// }

// /**
//  * @description: 新建用户
//  * @param {Object} data
//  * @author: Hemingway
//  */
// export function addUser(data) {
//   return newRequest({
//     url: '/agriculturefarm/v1/farmManagement/farmUser/add',
//     method: 'post',
//     data
//   })
// }

// /**
//  * @description: 编辑用户
//  * @param {Object} data
//  * @author: Hemingway
//  */
// export function editUser(data) {
//   return newRequest({
//     url: '/agriculturefarm/v1/farmManagement/farmUser/update',
//     method: 'post',
//     data
//   })
// }

// /**
//  * @description: 锁定与解锁用户
//  * @param {Object} data {userIds, isLocked}
//  * @author: Hemingway
//  */
// export function lockOrUnlockUser(data) {
//   return newRequest({
//     url: '/agriculturefarm/v1/farmManagement/farmUser/account/operate',
//     method: 'post',
//     data
//   })
// }

// /**
//  * @description: 删除用户
//  * @param {Object} data {userId}
//  * @author: Hemingway
//  */
// export function deleteUser(data) {
//   return newRequest({
//     url: '/agriculturefarm/v1/farmManagement/farmUser/delete',
//     method: 'post',
//     data
//   })
// }
