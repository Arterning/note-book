<!--
 * @Description  : 个人中心页面
 * @Autor        : Hemingway
 * @Date         : 2022-03-11 10:11:19
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-15 14:56:23
 * @FilePath     : \blockchain-admin\src\views\profile\index.vue
-->
<template>
  <div class="profile">
    <i-tab>
      <i-section-header title="基本信息" />
      <el-form label-position="top" class="section-form">
        <el-row>
          <el-col :span="8">
            <el-form-item label="姓名">
              <el-input
                :value="computedUserInfo.name"
                disabled
              /> </el-form-item
          ></el-col>
          <el-col :span="16">
            <el-form-item label="所属企业">
              <el-input
                :value="computedUserInfo.companyName"
                disabled
              /> </el-form-item
          ></el-col>
        </el-row>
      </el-form>

      <i-section-header title="账号设置" />
      <el-form label-position="top" class="section-form">
        <el-row>
          <el-col :span="8">
            <el-form-item label="手机绑定">
              <el-input prop="phone" :value="telphone" disabled ref="phone">
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
              <el-input
                :value="password"
                :placeholder="passwordPlaceholder"
                disabled
                type="password"
              >
                <el-button
                  slot="append"
                  @click="$refs.pwdDialog.visible = true"
                  >{{ hasPassword ? '更改' : '设置' }}</el-button
                ></el-input
              >
            </el-form-item></el-col
          >
        </el-row>
      </el-form>
      <!-- 手机号设置 dialog -->
      <set-phone ref="phoneDialog" :old-phone="telphone" />
      <!-- 设置密码 -->
      <set-password ref="pwdDialog" />
    </i-tab>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getPersonalCenter } from '@/api/profile'

export default {
  name: 'Profile',
  components: {
    SetPhone: () => import('./components/SetPhone'),
    SetPassword: () => import('./components/SetPassword')
  },
  data() {
    return {
      phone: '',
      password: '......',
      passwordPlaceholder: '',

      farmName: '',
      defaultFarmPlaceholder: '',
      showCascader: false,
      orgIdChain: [],

      userInfo: {}
    }
  },

  computed: {
    ...mapGetters(['username', 'hasPassword', 'name', 'orgRoles']),
    //TODO: 20220706 刷新后管理员信息丢失，做持久化存储
    ...mapGetters('user', { computedUserInfo: 'userInfoGetter' }),
    ...mapGetters('user', { telphone: 'getTelphone' }),
    // 格式化组织/角色
    orgRolesFormatter() {
      return this.orgRoles
        .map(item => item.organizationName + ' / ' + item.roleName)
        .join('，')
    }
  },

  watch: {
    username: {
      handler(newV) {
        this.phone = newV
      },
      immediate: true
    },
    hasPassword: {
      handler(newV) {
        if (newV) {
          this.password = '........'
        } else {
          this.passwordPlaceholder = '请设置密码'
        }
      },
      immediate: true
    }
  },
  created() {
    // this.getPersonalCenter()
  },
  methods: {
    async getPersonalCenter() {
      const { userLatestUseDto } = JSON.parse(
        localStorage.getItem('SUYUANMA_BLOBK_ADMIN_USER_INFO')
      )

      const { data } = await getPersonalCenter({
        userId: userLatestUseDto.userId
      })
      this.userInfo = data
    }
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
