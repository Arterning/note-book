import newRequest from '@/utils/newRequest'

// --------------------------------------- 角色管理 ---------------------------------------

/**
 * @description: 获取角色列表
 * @param {Object} data searchKey
 * @author: Hemingway
 */
export function getRole(params) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/list',
    method: 'get',
    params
  })
}

/**
 * @description: 创建角色
 * @param {Object} data 名称、启用状态、备注
 * @author: Hemingway
 */
export function addRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/create',
    method: 'post',
    data
  })
}

/**
 * @description: 编辑角色
 * @param {Object} data id、名称、启用状态、备注
 * @author: Hemingway
 */
export function editRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/update',
    method: 'post',
    data
  })
}

/**
 * @description: 删除角色
 * @param {Object} data id
 * @author: Hemingway
 */
export function deleteRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/delete',
    // method: 'delete',
    method: 'post',
    data
  })
}

/**
 * @description: 获取角色菜单功能列表
 * @param {Object} data roleId
 * @author: Hemingway
 */
export function getRoleRight(params) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/function/list',
    method: 'get',
    params
  })
}

/**
 * @description: 获取角色人员列表
 * @param {Object} data roleId,pageSize,pageNum
 * @author: Hemingway
 */
export function getRoleUser(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/user/list',
    method: 'post',
    data
  })
}

/**
 * @description: 设置指定角色的权限
 * @param {Object} data roleId、dataPermissionType、checkedMenuIdList、checkedFunctionList
 * @author: Hemingway
 */
export function submitRoleRight(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/privilege/bond',
    method: 'post',
    data
  })
}

// --------------------------------------- 员工管理 ---------------------------------------

// /**
//  * @description: 获取用户列表
//  * @param {Object} data keywords
//  * @author: Hemingway
//  */
export function getStaffList(data) {
  return newRequest({
    url: '/blockchaindbo/v1/user/list',
    // method: 'get',
    method: 'post',
    data
  })
}

// /**
//  * @description: 新建员工
//  * @param {Object} data
//  * @author: Hemingway
//  */
export function createStaff(data) {
  return newRequest({
    url: '/blockchaindbo/v1/user/create',
    method: 'post',
    data
  })
}

// /**
//  * @description: 编辑员工
//  * @param {Object} data
//  * @author: Hemingway
//  */
export function editStaff(data) {
  return newRequest({
    url: '/blockchaindbo/v1/user/update',
    method: 'post',
    data
  })
}

// /**
//  * @description: 删除员工
//  * @param {Object} data {userId}
//  * @author: Hemingway
//  */
export function deleteStaff(data) {
  return newRequest({
    url: '/blockchaindbo/v1/user/delete',
    method: 'post',
    data
  })
}

// /**
//  * @description: 冻结/解冻员工账号
//  * @param {Object} data {userId}
//  * @author: Hemingway
//  */
export function operateAccount(data) {
  return newRequest({
    url: '/blockchaindbo/v1/user/account/operate',
    method: 'post',
    data
  })
}
