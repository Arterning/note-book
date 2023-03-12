/*
 * @Description  : 提取公用js逻辑：获取用户信息/角色权限/拉取字典
 * @Autor        : Hemingway
 * @Date         : 2021-09-02 13:54:32
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-18 17:05:58
 * @FilePath     : \blockchain-minicode\src\utils\auth.js
 */
import { router } from '@/router'
import store from '@/store'
import { getUserInfo, getUserPermission, queryDict } from '@/api/user'

/**
 * @description: 用户登录后逻辑
 * @author: Hemingway
 */
export async function afterLogin() {
  console.log('----------1')
  const roleDtos = await checkUserInfo()
  console.log('----------3 已确认用户信息')

  let role = null
  const rolesName = roleDtos.map(role => role.name)

  if (store.state.user.role && rolesName.includes(store.state.user.role.name)) {
    console.log('----------4 正在设置当前角色信息（命中缓存）')
    // 角色命中缓存

    role = store.state.user.role
    await nextStep(role)
    console.log('----------7 已确认角色权限信息和字典信息')
  } else {
    console.log('----------4 正在设置当前角色信息')
    if (roleDtos.length === 1) {
      // 单角色用户

      role = roleDtos[0] // 使用唯一的角色

      await nextStep(role)
      console.log('----------7 已确认角色权限信息和字典信息')
    } else {
      // 多角色用户，选择一个角色

      router.push({ name: 'selectRole' })
    }
  }
}

/**
 * @description: 检查用户信息
 * @author: Hemingway
 */
async function checkUserInfo() {
  console.log('----------2')
  try {
    const { code, name, enterpriseName, roleDtos } = await getUserInfo()
    if (code === '0' && name && roleDtos && roleDtos.length > 0) {
      store.commit('setUserInfo', {
        name,
        companyName: enterpriseName || '',
        roles: roleDtos
      })

      return Promise.resolve(roleDtos)
    } else {
      return Promise.reject()
    }
  } catch (error) {
    console.log('获取用户信息 error = ', error)
    return Promise.reject()
  }
}

/**
 * @description: 根据角色拉取对应权限，并拉取字典
 * @param {Object} role
 * @return {Promise}
 * @author: Hemingway
 */
export async function nextStep(role) {
  console.log('----------5')
  // 更新角色缓存，并获取权限
  store.commit('setRole', role)
  uni.setStorageSync('role', JSON.stringify(role))

  const res = await getPermission(role)
  if (res) {
    getDict() // 请求字典

    return Promise.resolve()
  } else {
    return Promise.reject()
  }
}

/**
 * @description: 根据角色获取权限
 * @param {Object} role 角色
 * @author: Hemingway
 */
async function getPermission(role) {
  console.log('----------6')
  try {
    console.log('----------6-1')
    const { code, menuTree } = await getUserPermission({ roleId: role.id })
    console.log('----------6-2')
    if (code === '0') {
      store.commit('setPermissions', menuTree)
      console.log('setPermissions---', store.state.user.permissions)

      return Promise.resolve(true)
    } else {
      return Promise.resolve(false)
    }
  } catch (error) {
    console.log('获取角色权限 error = ', error)
    return Promise.resolve(false)
  }
}

/**
 * @description: 获取数据字典
 * @author: Hemingway
 */
async function getDict() {
  try {
    const { code, dictMap } = await queryDict()
    if (code === '0') {
      store.commit('setDictMap', dictMap)
    }
  } catch (error) {
    console.log('获取数据字典 error = ', error)
  }
}
