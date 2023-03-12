import request from '@/utils/request'

// 获取菜单树形结构
export function queryMenuTreeList(payload) {
  return request({
    url: '/blockchaindbo/menu/getMenuTree/w/v1',
    method: 'post',
    data: payload
  })
}

// 查询角色列表数据
export function getRoleListData(payload) {
  return request({
    url: '/blockchaindbo/adminRole/queryAllRoles/w/v1',
    method: 'post',
    data: payload
  })
}

// 查询角色权限数据
export function getRoleJurisdictionData(payload) {
  return request({
    url: '/blockchaindbo/adminRole/getRoleRightsList/w/v1',
    method: 'post',
    data: payload
  })
}

// 新建角色（不包含菜单权限）
export function createRole(payload) {
  return request({
    url: '/blockchaindbo/adminRole/create/w/v1',
    method: 'post',
    data: {
      enable: 'Y',
      ...payload
    }
  })
}

// 更新角色（不包含菜单权限）
export function updateRole(payload) {
  return request({
    url: '/blockchaindbo/adminRole/update/w/v1',
    method: 'post',
    data: payload
  })
}

// 删除角色
export function deleteRole(payload) {
  return request({
    url: '/blockchaindbo/adminRole/deleteRole/w/v1',
    method: 'post',
    data: payload
  })
}

// 保存角色菜单权限
export function saveRoleRightsTree(payload) {
  return request({
    url: '/blockchaindbo/adminRole/saveRoleRightsTree/w/v1',
    method: 'post',
    data: payload
  })
}
