<!--
 * @Description  : 货源加密货币
 * @Autor        : SunXiuqiong
 * @Date         : 2022-05-31 17:35:43
 * @LastEditors  : SunXiuqiong
 * @LastEditTime : 2022-06-29 18:14:59
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\bindlist\index.vue
-->
<template>
  <div class="bind-list">
    <i-tab
      :activeName="activeName"
      :label="['待绑定', '已绑定']"
      @tab-click="onTabClick"
    >
      <template #0>
        <template v-if="activeName == '0'">
          <unbind-list />
        </template>
      </template>
      <template #1>
        <template v-if="activeName == '1'">
          <bind-list />
        </template>
      </template>
    </i-tab>
  </div>
</template>

<script>
import UnbindList from './components/UnbindList.vue'
import BindList from './components/BindList.vue'

export default {
  name: 'traceability-manage--bindlist',
  components: {
    UnbindList,
    BindList
  },
  data() {
    return {
      activeName: '0'
    }
  },
  methods: {
    onTabClick(tabpane) {
      this.activeName = tabpane.index
    }
  }
}
</script>

<style lang="scss" scoped>
.bind-list {
  height: 100%;
  background-color: #fff;
  position: relative;
}
</style>
