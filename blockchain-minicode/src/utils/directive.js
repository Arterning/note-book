/*
 * @Description  : Action 权限指令
 * @Autor        : Hemingway
 * @Date         : 2021-06-16 10:58:45
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-06-16 13:51:23
 * @FilePath     : \blockchain-minicode\src\utils\directive.js
 */

import Vue from 'vue'
import store from '@/store'

/**
 * Action 权限指令
 * 指令用法：
 *  - 在需要控制 action 级别权限的组件上使用 v-action:[method] , 如下：
 *    <view v-action:edit @click="edit(record)">修改</view>
 *  - 当前用户没有权限时，组件上使用了该指令则会被隐藏
 */
const action = Vue.directive('action', {
  inserted: function(el, binding) {
    // 根据指令，检索对应的用户权限
    const permissions = store.state.user.permissions
    const actions = permissions.find(
      item => item.name === el.parentNode.previousSibling.textContent
    ).children
    const actionSet = actions && actions.map(item => item.name)

    // 如果无权限，隐藏节点
    const actionName = binding.arg
    if (actionSet && !actionSet.includes(actionName)) {
      ;(el.parentNode && el.parentNode.removeChild(el)) ||
        (el.style.display = 'none')
    }
  }
})

export default action
