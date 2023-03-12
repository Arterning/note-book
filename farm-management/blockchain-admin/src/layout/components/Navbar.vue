<template>
  <div class="navbar">
    <div class="left-side">
      <hamburger
        id="hamburger-container"
        :is-active="sidebar.opened"
        class="hamburger-container"
        @toggleClick="toggleSideBar"
      />
    </div>

    <div class="right-menu">
      <template v-if="device !== 'mobile'">
        <screenfull id="screenfull" class="right-menu-item hover-effect" />
      </template>

      <el-dropdown
        class="avatar-container right-menu-item hover-effect"
        trigger="click"
      >
        <div class="avatar-wrapper">
          <span>{{ `欢迎您，${computedUserInfo.name}` }}</span>
          <i class="el-icon-caret-bottom" />
        </div>
        <el-dropdown-menu slot="dropdown">
          <router-link to="/profile/index">
            <el-dropdown-item>个人中心</el-dropdown-item>
          </router-link>
          <el-dropdown-item @click.native="logout" divided>
            退 出
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Hamburger from '@/components/Hamburger'
import Screenfull from '@/components/Screenfull'

export default {
  components: {
    Hamburger,
    Screenfull
  },

  data() {
    return {}
  },

  computed: {
    ...mapGetters([
      'sidebar',
      // 'avatar',
      'device',
      'name'
    ]),
    //TODO: 20220706 刷新后管理员信息丢失，做持久化存储
    ...mapGetters('user', { computedUserInfo: 'userInfoGetter' })
  },

  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },

    async logout() {
      await this.$store.dispatch('user/logout')
      this.$router.push(`/login?redirect=${this.$route.fullPath}`)
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 44px;
  overflow: hidden;
  background: #fff;
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: relative;

  .hamburger-container {
    padding: 0 16px;
    height: 44px;
    line-height: 44px;
    float: left;
    cursor: pointer;
    transition: background 0.3s;
    -webkit-tap-highlight-color: transparent;
    display: flex;
    align-items: center;

    &:hover {
      background: rgba(0, 0, 0, 0.025);
    }
  }

  .right-menu {
    height: 100%;
    line-height: 44px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      padding: 0 16px;
      display: inline-block;
      height: 100%;
      font-size: 16px;
      color: #5a5e66;

      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;

        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }

    .avatar-container {
      .avatar-wrapper {
        position: relative;
        margin-right: 20px;
      }
    }
  }

  .el-icon-caret-bottom {
    cursor: pointer;
    position: absolute;
    right: -20px;
    top: 50%;
    transform: translateY(-50%);
    font-size: 12px;
  }
}
</style>
