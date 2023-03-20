<!--
 * @Description  : 我的
 * @Autor        : Hemingway
 * @Date         : 2021-06-22 11:18:55
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-24 09:45:07
 * @FilePath     : \blockchain-minicode\src\pages\personal\index.vue
-->
<template>
  <view class="personal">
    <card :data="userData">
      <template #title>
        <view class="card__title">
          <img class="avatar" mode="aspectFit" src="/static/img/avatar.png" />
          <text>{{ userData.title }}</text>
        </view>
      </template>
      <template slot-scope="{ needDecoration }">
        <template v-if="needDecoration">
          <button
            class="role-btn"
            @click="$Router.push({ name: 'selectRole' })"
          >
            切换身份
          </button>
        </template>
      </template>
    </card>
    <uni-list :border="false">
      <uni-list-item
        title="消息中心"
        :show-badge="unRead > 0"
        :badge-text="unRead"
        :show-arrow="true"
        :border="false"
        :clickable="true"
        @click="$Router.push({ name: 'message' })"
      >
        <template #header>
          <text class="iconfont icon-notifications_outlined"></text>
        </template>
      </uni-list-item>
      <uni-list-item
        v-if="role.enterpriseType === '2'"
        title="溯源素材"
        :show-arrow="true"
        :border="true"
        :clickable="true"
        @click="$Router.push({ name: 'material' })"
      >
        <template #header>
          <text class="iconfont .icon-file_add_outlined"></text>
        </template>
      </uni-list-item>
      <uni-list-item
        class="layout-center"
        title="注销账号"
        :border="false"
        :clickable="true"
        @click="onCloseAccount"
      />
      <uni-list-item
        class="layout-center"
        title="退出"
        :border="false"
        :clickable="true"
        @click="onLogout"
      />
    </uni-list>
  </view>
</template>

<script>
import Card from '@/components/card'
import { queryUnreadMessageCount } from '@/api/user'

export default {
  name: 'Personal',

  components: {
    Card
  },

  data() {
    return {
      unRead: 0 // 消息未读数
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

    // 获取当前角色
    role() {
      return this.$store.state.user.role
    },
    // 当前用户信息
    userData() {
      const obj = {
        title: '',
        items: [{ text: `所属公司：` }, { text: `我的身份：` }]
      }

      if (this.userInfo) {
        // 已登录

        obj.title = this.userInfo.name
        obj.items[0].text += this.userInfo.companyName || '[暂无]'

        if (this.userInfo.roles.length > 1) {
          obj.items[1].needDecoration = true // 需要装饰
        }
      }

      if (this.role) {
        obj.items[1].text += this.role.name || ''
      }

      return obj
    }
  },

  created() {
    // 未登录用户点击tab到访事件
    if (!this.isLogin) {
      this.$Router.replace({ name: 'login' })
    }
  },

  onShow() {
    if (this.isLogin) {
      this.queryMessage()
    }
  },

  methods: {
    /**
     * @description: 查询消息未读数量
     * @author: Hemingway
     */
    async queryMessage() {
      try {
        const { code, unReadNumber } = await queryUnreadMessageCount()
        if (code === '0') {
          this.unRead = unReadNumber
        }
      } catch (error) {
        console.log('查询未读消息数量 error = ', error)
      }
    },

    /**
     * @description: 注销事件
     * @author: Hemingway
     */
    onCloseAccount() {
      uni.showModal({
        title: '提示',
        content: '数据无法恢复，请谨慎',
        success: res => {
          if (res.confirm) {
            // this.$store.commit('setSessionId', '')
            // this.$store.commit('setPermissions', null)
            // this.$store.commit('setUserInfo', null)
            // this.$store.commit('setRole', null)
            // uni.removeStorageSync('sessionId')
            // uni.removeStorageSync('role')
            // uni.reLaunch({ url: '/pages/login/index' })
            uni.showToast({
              title: '功能静待开放',
              icon: 'none',
              duration: 2000
            })
          }
        }
      })
    },

    /**
     * @description: 登出事件
     * @author: Hemingway
     */
    onLogout() {
      uni.showModal({
        title: '提示',
        content: '退出登录',
        success: res => {
          if (res.confirm) {
            this.$store.commit('setSessionId', '')
            this.$store.commit('setPermissions', null)
            this.$store.commit('setUserInfo', null)
            this.$store.commit('setRole', null)
            uni.removeStorageSync('sessionId')
            uni.removeStorageSync('role')
            uni.reLaunch({ url: '/pages/login/index' })
          }
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.personal {
  // 个人卡片
  .card__title {
    display: flex;
    align-items: center;

    .avatar {
      width: 80rpx;
      height: 80rpx;
      margin-right: 24rpx;
    }
  }

  .role-btn {
    height: 48rpx;
    line-height: 48rpx;
    font-size: 26rpx;
    color: #fff;
    background-color: $uni-bg-color-btn;
    border-radius: 44rpx;
  }

  .role-btn::after {
    display: none;
  }

  // 列表
  ::v-deep {
    .uni-list {
      background-color: $uni-bg-color-page;
    }

    .uni-badge--x {
      display: block;

      .uni-badge {
        background-color: #fa5050;
      }
    }

    .iconfont {
      font-size: 40rpx;
      margin-right: 8rpx;
      color: $theme-color;
    }
  }

  // 居中对齐
  .layout-center {
    margin-top: 16rpx;
    text-align: center;
  }
}
</style>
