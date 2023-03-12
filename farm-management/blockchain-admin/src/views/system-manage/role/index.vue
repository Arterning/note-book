<!--
 * @Description  : 角色管理
 * @Autor        : Hemingway
 * @Date         : 2021-12-22 19:47:44
 * @LastEditors  : SunXiuqiong
 * @LastEditTime : 2022-07-21 22:07:37
 * @FilePath     : \blockchain-admin\src\views\system-manage\role\index.vue
-->
<template>
  <div class="role">
    <i-tab>
      <div class="wrapper">
        <role-list @onChooseRole="onChooseRole" />
        <role-action ref="action" :role="role" />
      </div>
    </i-tab>
  </div>
</template>

<script>
import RoleList from './components/RoleList.vue'
import RoleAction from './components/RoleAction.vue'

export default {
  name: 'system-manage--role',

  components: {
    RoleList,
    RoleAction
  },

  data() {
    return {
      role: {} // 当前处理的角色
    }
  },

  methods: {
    /**
     * @description: 角色选择事件
     * @param {Object} role
     * @author: Hemingway
     */
    onChooseRole(role) {
      this.role = role
    }
  }
}
</script>

<style lang="scss" scoped>
.role {
  height: 100%;
  background-color: #fff;
  position: relative;

  .wrapper {
    background-color: #fff;
    display: grid;
    grid-template-columns: 12fr 12fr;
    grid-template-rows: 100%;
    grid-gap: 32px;
  }
}
</style>
