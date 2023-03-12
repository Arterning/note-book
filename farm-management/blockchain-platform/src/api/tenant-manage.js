import newRequest from '@/utils/newRequest'
import request from '@/utils/request'
// --------------------------------------- 角色管理 ---------------------------------------

/**
 * @description: 运营平台-服务模式-获取服务模式
 * @param {Object} data searchKey
 * @author: Hemingway
 */
export function getRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/serviceModel/list',
    method: 'get',
    data
  })
}

/**
 * @description: 运营平台-服务模式-获取服务模式
 * @param {Object} data searchKey
 * @author: Hemingway
 */
export function getService(data) {
  return request({
    url: '/blockchaindbo/v1/serviceModel/getServiceModelByEnterpriseType',
    method: 'get',
    data
  })
}

/**
 * @description: 运营平台-服务模式-新增服务模式
 * @param {Object} data 名称、启用状态、备注
 * @author: Hemingway
 */
export function addRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/serviceModel/create',
    method: 'post',
    data
  })
}

/**
 * @description: 运营平台-服务模式-修改服务模式
 * @param {Object} data id、名称、启用状态、备注
 * @author: Hemingway
 */
export function editRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/serviceModel/update',
    method: 'post',
    data
  })
}

/**
 * @description: 运营平台-服务模式-删除服务模式
 * @param {Object} data id
 * @author: Hemingway
 */
export function deleteRole(data) {
  return newRequest({
    url: '/blockchaindbo/v1/serviceModel/delete',
    method: 'post',
    data
  })
}

/**
 * @description: 获取角色菜单功能列表
 * @param {Object} data roleId
 * @author: Hemingway
 */
export function getRoleRight(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/function/list',
    method: 'get',
    data
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

// 运营平台-服务模式-获取服务模式【关联】租户列表
export function getServiceUser(data) {
  return newRequest({
    url: '/blockchaindbo/v1/serviceModel/tenant/list',
    method: 'get',
    data
  })
}

/**
 * @description: 设置指定角色的权限
 * @param {Object} data roleId、dataPermissionType、checkedMenuIdList、checkedFunctionList
 * @author: Hemingway
 */
export function setRoleRight(data) {
  return newRequest({
    url: '/agriculturefarm/v1/role/setRoleRights',
    method: 'post',
    data
  })
}

// --------------------------------------- 员工管理 ---------------------------------------

// /**
//  * @description: 获取员工列表
//  * @param {Object} data keywords
//  * @author: Hemingway
//  */
export function getStaffList(params) {
  return newRequest({
    url: '/blockchainadmin/v1/staff/list',
    method: 'get',
    params
  })
}

// /**
//  * @description: 新建员工
//  * @param {Object} data
//  * @author: Hemingway
//  */
export function createStaff(data) {
  return newRequest({
    url: '/blockchainadmin/v1/staff/create',
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
    url: '/blockchainadmin/v1/staff/modify',
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
    url: '/blockchainadmin/v1/staff/delete',
    method: 'delete',
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
    url: '/blockchainadmin/v1/staff/account/operate',
    method: 'post',
    data
  })
}

// 绑定/更新角色菜单功能权限(提交)
export function submitRoleRight(data) {
  return newRequest({
    url: '/blockchaindbo/v1/role/management/privilege/bond',
    method: 'post',
    data
  })
}
