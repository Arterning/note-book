<!--
 * @Description  : 登录页面
 * @Autor        : Hemingway
 * @Date         : 2021-12-14 17:20:54
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-15 11:57:10
 * @FilePath     : \blockchain-platform\src\views\login\index.vue
-->
<template>
  <div class="login-container">
    <login-window
      v-if="!dialogFormVisible"
      @setPassword="dialogFormVisible = true"
    >
      <div class="window-title">智农云区块链溯源运营平台</div>
      <div class="window-content">
        <div class="tab-box" v-if="mode === 'login'">
          <div
            class="tab-item"
            :class="{ 'active-item': activeClass === 'passwordLogin' }"
            @click="activeClass = 'passwordLogin'"
          >
            密码登录
          </div>
          <div
            class="tab-item"
            :class="{ 'active-item': activeClass === 'messageLogin' }"
            @click="activeClass = 'messageLogin'"
          >
            短信登录
          </div>
        </div>
        <login-form
          v-if="mode === 'login'"
          :username="username"
          :isPasswordLogin="activeClass === 'passwordLogin'"
          @login-success="next"
          @set-password="dialogFormVisible = true"
          @reset="onReset"
        ></login-form>
        <reset-form
          v-else
          @login="onLogin"
          @reset-success="onLogin"
          :username="username"
        ></reset-form>
      </div>
      <!-- <el-tabs v-model="activeName">
        <el-tab-pane label="用户登录" name="userLogin">
          <login-form
            v-if="mode === 'login'"
            :username="username"
            @login-success="next"
            @set-password="dialogFormVisible = true"
            @reset="onReset"
          ></login-form>
          <reset-form
            v-else
            @login="onLogin"
            @reset-success="onLogin"
            :username="username"
          ></reset-form>
        </el-tab-pane>
      </el-tabs> -->
    </login-window>

    <!-- dialog -->
    <i-dialog
      :showClose="false"
      :visible.sync="dialogFormVisible"
      v-else
      class="dialog"
    >
      <template #title>
        设置密码
        <el-tooltip
          content="密码长度为 8 ~ 20位（含大写字母、小写字母、数字、特殊字符两项及以上）"
          placement="top"
          class="tooltip"
        >
          <i class="el-icon-info"> </i>
        </el-tooltip>
      </template>
      <set-password-form
        @submit-disabled="isSubmitDisabled = $event"
        @set-password-success="next"
        ref="setPassword"
      />

      <div slot="footer" class="dialog-footer">
        <el-button @click="next">下次再说</el-button>
        <el-button
          type="primary"
          @click="$refs.setPassword.submit()"
          :disabled="isSubmitDisabled"
          >确 定</el-button
        >
      </div>
    </i-dialog>
  </div>
</template>

<script>
import LoginWindow from './components/LoginWindow.vue'
import LoginForm from './components/LoginForm.vue'
import ResetForm from './components/ResetForm.vue'
import SetPasswordForm from './components/SetPasswordForm'
// import { getUsernameLTS } from '@/utils/auth'

export default {
  name: 'Login',
  components: {
    LoginWindow,
    LoginForm,
    ResetForm,
    SetPasswordForm
  },
  data() {
    return {
      activeName: 'userLogin',
      dialogFormVisible: false, // 是否弹出设置密码dialog
      isSubmitDisabled: true, // 设置密码dialog，提交按钮是否禁用

      mode: 'login', // login 登录；reset 重置
      username: '',

      otherQuery: {},
      activeClass: 'passwordLogin'
    }
  },

  watch: {
    $route: {
      handler: function (route) {
        const query = route.query
        if (query) {
          this.redirect = query.redirect
          // this.otherQuery = this.getOtherQuery(query)
        }
      },
      immediate: true
    }
  },

  methods: {
    /**
     * @description: 执行跳转首页逻辑
     * @author: Hemingway
     */
    next() {
      // console.log('路由1', this.$store.state.addRoutes)
      this.$router.push('/operational/operational-report')
      // this.$router.push({
      //   path: this.redirect || '/',
      //   // query: this.otherQuery
      // })
    },

    /**
     * @description: 重置密码事件
     * @param {String} e
     * @author: Hemingway
     */
    onReset(e) {
      this.username = e
      this.mode = 'reset'
    },
    /**
     * @description: 登录事件
     * @param {String} e
     * @author: Hemingway
     */
    onLogin(e) {
      this.username = e
      this.mode = 'login'
    },

    getOtherQuery(query) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {})
    }
  }
}
</script>

<style lang="scss" scoped>
$bg: #2d3a4b;
.login-container {
  min-height: 100%;
  width: 100vw;
  background-color: $bg;
  background-image: url('../../assets/images/bg-img.png');
  background-size: cover;
  overflow: hidden;

  .dialog {
    ::v-deep .el-dialog {
      width: 30%;
    }
  }
}
</style>

<style scoped lang="scss">
.login-container {
  background-color: #153a7b;
}

.dataCollect-style {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  height: 249px;
  background-color: #b9c4d8 !important;

  div {
    text-align: center;
    width: 180px;
    height: 180px;
  }
}

.tab-box {
  display: flex;
  justify-content: space-evenly;
  font-size: 16px;
  padding: 0 50px;

  .tab-item {
    padding: 10px 0;
    color: #8e8e8e;
    cursor: pointer;
  }
  .active-item {
    border-bottom: 1px solid #05dbfe;
    color: black;
  }
}

.qrcode-box {
  position: fixed;
  right: 20px;
  bottom: 20px;
  display: flex;
  .qrcode-item {
    width: 125px;
    height: 150px;
    background: white;
    border-radius: 5px;
    margin-left: 20px;
    padding: 10px;
    font-size: 12px;
    text-align: center;

    .img-box {
      width: 100%;
      height: 105px;
      background: #ccc;
    }
    .text-box {
      line-height: 30px;
    }
  }
}

::v-deep .el-tabs__active-bar {
  background-color: #dfe4ed;
}

::v-deep .el-tabs__item.is-active {
  color: #96d045;
}

::v-deep .el-tabs__header .is-top {
  font-size: 16px;
}

::v-deep .el-tabs__item:hover {
  color: #96d045;
}

::v-deep .el-tab-pane {
  background: #fff;
  border-radius: 6px !important;
}

::v-deep .login-window {
  .window-title {
    color: white;
    font-size: 40px;
    padding-bottom: 20px;
  }
  .window-content {
    background: #b9c4d8;
    border-radius: 6px;
    padding: 30px 35px;
  }
}

::v-deep .login-window .el-form-item {
  background-color: #fff;
  // border: 1px solid rgb(222, 220, 220);
  height: 36px;
  margin-bottom: 30px;
  line-height: 36px;
}
::v-deep .login-window .el-input input {
  padding: 4px 12px;
  height: 30px;
}

::v-deep .el-form-item__label:before {
  content: '' !important;
}

::v-deep .el-form-item--medium .el-form-item__label {
  color: #112350;
  font-weight: 100;
}
</style>
