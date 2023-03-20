import request from '@/utils/request'

// 查询角色列表数据
export function getListData() {
  return request({
    url: '/agriculturedbo/adminRole/getRoleList/w/v1',
    method: 'post'
  })
}

// 新增角色
export function createRole(payload) {
  return request({
    url: '/agriculturedbo/adminRole/create/w/v1',
    method: 'post',
    data: payload
  })
}

// 修改角色
export function editRole(payload) {
  return request({
    url: '/agriculturedbo/adminRole/update/w/v1',
    method: 'post',
    data: payload
  })
}

// 删除角色
export function deleteRole(payload) {
  return request({
    url: '/agriculturedbo/adminRole/deleteRole/w/v1',
    method: 'post',
    data: payload
  })
}

// 首页查询公司
export function getCompany(payload) {
  return request({
    url: '/agriculturepoadmin/role/chooseCorp/w/v1',
    method: 'post',
    data: payload
  })
}

// 获取角色对应的菜单权限
export function getRoleRightsList(payload) {
  return request({
    url: '/agriculturedbo/adminRole/getRoleRightsList/w/v1',
    method: 'post',
    data: payload
  })
}

// 保存角色对应的菜单权限
export function saveRoleRightsTree(payload) {
  return request({
    url: '/agriculturedbo/adminRole/saveRoleRightsTree/w/v1',
    method: 'post',
    data: payload
  })
}

// 分页获取角色授权用户
export function getRoleUsers(payload) {
  return request({
    url: '/agriculturedbo/adminRole/getRoleUsers/w/v1',
    method: 'post',
    data: payload
  })
}

// 删除角色与用户关系
export function deleteRoleUsers(payload) {
  return request({
    url: '/agriculturepoadmin/role/deleteRoleUsers/w/v1',
    method: 'post',
    data: payload
  })
}

// 数据权限 分公司数据
export function getOrgIsCorp(payload) {
  return request({
    url: '/agriculturedbo/adminOrg/getOrgIsCorp/w/v1',
    method: 'post',
    data: payload
  })
}

// 权限设置 - 数据权限 ->  获取产品列表
export function queryProductList(payload) {
  return request({
    url: '/agriculturedbo/product/query/w/v1',
    method: 'post',
    data: payload
  })
}
