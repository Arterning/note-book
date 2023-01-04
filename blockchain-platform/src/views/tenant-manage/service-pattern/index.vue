<!--
 * @Description  : 角色管理
 * @Autor        : Hemingway
 * @Date         : 2021-12-22 19:47:44
 * @LastEditors  : SunXiuqiong
 * @LastEditTime : 2022-07-26 12:29:26
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\service-pattern\index.vue
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
  name: 'blockchain-service-pattern',

  components: {
    RoleList,
    RoleAction
  },

  data() {
    return {
      // 当前处理的角色
      role: {}
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
