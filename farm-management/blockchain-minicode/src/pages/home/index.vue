<!--
 * @Description  : 首页
 * @Autor        : Hemingway
 * @Date         : 2021-09-08 16:38:22
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-06 10:23:55
 * @FilePath     : \blockchain-minicode\src\pages\home\index.vue
-->
<template>
  <view class="home">
    <view
      class="banner-wrapper"
      @click="isLogin ? null : $Router.push({ name: 'login' })"
    >
      <img class="banner" mode="aspectFit" src="/static/img/banner.png" />
    </view>

    <template v-if="isLogin">
      <view
        v-for="(item, index) in permissionsFilter"
        :key="index"
        class="permission"
      >
        <view class="title">{{ item.title }}</view>
        <view class="container">
          <view
            v-for="(ite, idx) in item.list"
            :key="idx"
            class="container-item"
            @tap="onActionClick(ite)"
          >
            <text :class="['iconfont', 'icon-' + ite.icon]"></text>
            <text class="container-text">{{ ite.text }}</text>
          </view>
        </view>
      </view>

      <!-- 登录未入驻弹窗 -->
      <uni-popup ref="application" :mask-click="false">
        <popup
          title="申请入驻"
          content="您的账号未关联到企业，需要申请企业入驻"
          confirm-text="立即申请"
          @close="close('application')"
          @confirm="onConfirm('application')"
        />
      </uni-popup>
    </template>
    <template v-else>
      <view class="login-tip-wrapper">
        <text class="login-text" @click="$Router.push({ name: 'login' })"
          >立即登录</text
        >
        <text>，查看更多功能...</text>
      </view>
    </template>
  </view>
</template>

<script>
import { queryApplicationList } from '@/api/user'
import { permissionMap } from './permissionMap'
import popup from '../../components/Popup.vue'

export default {
  name: 'Home',

  components: {
    popup
  },

  data() {
    return {
      activeApplication: null, // 流程相关的入驻申请信息
      isChecked: false // 入驻检查弹窗
    }
  },

  computed: {
    // 是否已登录
    isLogin() {
      return this.$store.state.user.sessionId
    },

    // 获取用户信息
    userInfo() {
      return this.$store.state.user.userInfo
    },

    // 权限控制，返回用户角色在首页可用的功能入口
    permissionsFilter() {
      return this.filterAsyncPermisson(this.$store.state.user.permissions)
    }
  },

  mounted() {},

  onShow() {
    this.checkApplication()
  },

  methods: {
    /**
     * @description: 首页登录后业务逻辑：入驻检查
     * @author: Hemingway
     */
    async checkApplication() {
      if (this.isChecked) {
        return
      }

      if (!this.isLogin) {
        return
      }

      const { total, active } = await this.queryList()
      // 存储生效的入驻记录
      this.activeApplication = active

      // 入驻弹窗
      if (
        this.userInfo.roles.length === 1 &&
        this.userInfo.roles[0].name === '初始角色' &&
        total === 0
      ) {
        this.openPopup('application')
      }

      this.isChecked = true
    },

    /**
     * @description: 查询审核列表
     * @author: Hemingway
     */
    async queryList() {
      try {
        const { code, total, list } = await queryApplicationList()
        if (code === '0') {
          return Promise.resolve({
            total,
            active: (list || []).find(item => item.checkStatus === '1')
          })
        } else {
          return Promise.reject()
        }
      } catch (error) {
        console.log('入驻申请概况查询 error = ', error)
        return Promise.reject()
      }
    },

    /**
     * @description: 动态权限过滤
     * @param {Array} permissions 接口返回的用户权限
     * @return {Array} 根据角色权限匹配首页功能入口
     * @author: Hemingway
     */
    filterAsyncPermisson(permissions) {
      if (!permissions || permissions.length === 0) {
        return []
      }

      const arr = []

      permissionMap.forEach(permission => {
        const tempPermission = Object.assign({}, permission)

        if (permissions.map(item => item.name).includes(permission.title)) {
          const obj = permissions.find(item => item.name === permission.title)
          if (obj.children && obj.children.length > 0) {
            tempPermission.list = permission.list.filter(item =>
              obj.children.map(child => child.name).includes(item.text)
            )

            arr.push(tempPermission)
          }
        }
      })

      return arr
    },

    /**
     * @description: 功能点击事件
     * @param {String} name
     * @param {String} url
     * @author: Hemingway
     */
    onActionClick({ url }) {
      if (url) {
        if (url === '/pages/application/index' && this.activeApplication) {
          // 入驻申请
          this.$Router.push({
            path: '/application',
            query: {
              id: this.activeApplication.id,
              handleType: 'append'
            }
          })
        } else {
          uni.navigateTo({ url })
        }
      }
    },

    /**
     * @description: 打开指定弹窗
     * @param {String} ref
     * @author: Hemingway
     */
    openPopup(ref) {
      this.$refs[ref].open()
    },

    /**
     * @description: 关闭指定弹窗
     * @param {String} ref
     * @author: Hemingway
     */
    close(ref) {
      this.$refs[ref].close()
    },

    /**
     * @description: 确认指定弹窗
     * @param {String} ref
     * @author: Hemingway
     */
    onConfirm(ref) {
      this.$refs[ref].close()
      uni.navigateTo({ url: `/pages/application/index` })
    }
  }
}
</script>

<style lang="scss" scoped>
.home {
  .banner-wrapper {
    box-sizing: border-box;
    width: 100%;
    height: 418rpx;
    padding: 16rpx 16rpx 0;

    .banner {
      width: 100%;
      height: 100%;
      border-radius: 16rpx;
    }
  }

  .permission {
    width: 750rpx;

    .title {
      height: 80rpx;
      line-height: 80rpx;
      font-size: 24rpx;
      padding-left: 32rpx;
    }

    .container {
      padding: 24rpx 35rpx;
      background-color: #fff;
      display: flex;
      flex-wrap: wrap;

      .container-item {
        box-sizing: border-box;
        width: 170rpx;
        height: 156rpx;
        padding: 24rpx 0;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;

        .iconfont {
          font-size: 48rpx;
          color: $theme-color;
        }

        .container-text {
          margin-top: 20rpx;
          color: #212121;
          font-size: 24rpx;
        }
      }
    }
  }

  .login-tip-wrapper {
    height: calc(100vh - 418rpx);
    display: flex;
    justify-content: center;
    align-items: center;

    .login-text {
      color: $theme-color;
    }
  }
}
</style>
