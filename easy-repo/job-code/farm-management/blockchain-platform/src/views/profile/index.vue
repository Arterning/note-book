<!--
 * @Description  : 个人中心页面
 * @Autor        : Hemingway
 * @Date         : 2022-03-11 10:11:19
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-15 14:58:34
 * @FilePath     : \blockchain-platform\src\views\profile\index.vue
-->
<template>
  <div class="profile">
    <i-tab>
      <i-section-header title="基本信息" />
      <el-form label-position="top" class="section-form">
        <el-row>
          <el-col :span="8">
            <el-form-item label="姓名">
              <el-input :value="name" disabled /> </el-form-item
          ></el-col>
          <el-col :span="16">
            <el-form-item label="角色">
              <el-input :value="rolesFormatter" disabled /> </el-form-item
          ></el-col>
        </el-row>
      </el-form>

      <i-section-header title="账号设置" />
      <el-form label-position="top" class="section-form">
        <el-row>
          <el-col :span="8">
            <el-form-item label="手机绑定">
              <el-input prop="phone" :value="phone" disabled ref="phone">
                <el-button
                  slot="append"
                  @click="$refs.phoneDialog.visible = true"
                  >更改</el-button
                ></el-input
              ></el-form-item
            ></el-col
          >
          <el-col :span="8">
            <el-form-item label="密码设置">
              <el-input :value="password" disabled type="password">
                <el-button slot="append" @click="$refs.pwdDialog.visible = true"
                  >更改</el-button
                ></el-input
              >
            </el-form-item></el-col
          >
        </el-row>
      </el-form>
      <!-- 手机号设置 dialog -->
      <set-phone ref="phoneDialog" :old-phone="phone" />
      <!-- 设置密码 -->
      <set-password ref="pwdDialog" />
    </i-tab>
  </div>
</template>

<script>
// import { mapGetters } from 'vuex'

export default {
  name: 'Profile',
  components: {
    SetPhone: () => import('./components/SetPhone'),
    SetPassword: () => import('./components/SetPassword')
  },
  data() {
    return {
      password: '........',

      farmName: '',
      showCascader: false,
      orgIdChain: [],
      rolesFormatter: '',
      name: '',
      phone: ''
    }
  },
  mounted() {
    let userInfo = JSON.parse(
      localStorage.getItem(process.env.VUE_APP_UNIQUE_PREFIX + '_USER_INFO')
    )
    this.name = userInfo?.name || ''
    this.rolesFormatter = userInfo?.initRoleName || ''
    this.phone = localStorage.getItem(
      process.env.VUE_APP_UNIQUE_PREFIX + '_PHONE'
    )
  }
}
</script>

<style lang="scss" scoped>
.profile {
  height: 100%;
  background-color: #fff;
  position: relative;

  .section-form {
    padding: 0 16px;

    ::v-deep {
      .el-form-item {
        width: 280px;
      }

      .el-input-group__append:hover {
        background-color: #e6faee;
        color: #00c853;
      }

      .el-cascader {
        width: 280px;
      }
    }
  }
}
</style>
