/* eslint-disable */
import * as roleManager from '@/api/roleManage'

var roleManagers = {
  actions: {
    getListData({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .getListData(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    createRole({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .createRole(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    editRole({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .editRole(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    deleteRole({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .deleteRole(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    getCompany({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .getCompany(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    getRoleRightsList({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .getRoleRightsList(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    saveRoleRightsTree({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .saveRoleRightsTree(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    getRoleUsers({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .getRoleUsers(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    deleteRoleUsers({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .deleteRoleUsers(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    },
    getOrgIsCorp({ commit, state }, payload) {
      return new Promise((resolve, reject) => {
        roleManager
          .getOrgIsCorp(payload)
          .then(res => {
            resolve(res)
          })
          .catch(error => {
            console.log(error)
          })
      })
    }
  }
}

export default roleManagers
