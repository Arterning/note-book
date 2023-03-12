<!--
 * @Description  : 角色选择页面
 * @Autor        : Hemingway
 * @Date         : 2021-09-02 18:45:14
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-18 16:57:15
 * @FilePath     : \blockchain-minicode\src\pages\select-role\index.vue
-->
<template>
  <view class="select-role">
    <view class="prompt">
      <text class="text">点击以选择身份</text>
    </view>
    <view class="role-container">
      <view
        v-for="(item, index) in roles"
        :key="index"
        class="wrapper"
        @click="onChoose(item)"
      >
        <text>{{ item.name }}</text>
        <view v-if="isCurrentRole(item)" class="mark">
          <uni-icons
            type="smallcircle-filled"
            size="12"
            color="#00c853"
            style="margin-right: 4rpx;"
          ></uni-icons>
          <text>当前使用</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { nextStep } from '@/utils/auth'

export default {
  name: 'SelectRole',

  data() {
    return {}
  },

  computed: {
    // 获取角色列表
    roles() {
      return this.$store.state.user.userInfo.roles
    },

    // 获取当前角色
    role() {
      return this.$store.state.user.role
    }
  },

  created() {},

  mounted() {},

  methods: {
    /**
     * @description: 角色选择事件
     * @param {Object} role
     * @author: Hemingway
     */
    async onChoose(role) {
      if (this.isCurrentRole(role)) {
        this.$Router.back()
      } else {
        await nextStep(role)
        console.log('----------7 已确认角色权限信息和字典信息')
        this.$Router.replaceAll({ name: 'home' })
      }
    },

    /**
     * @description: 是否是当前角色
     * @param {Object} role
     * @return {Boolean}
     * @author: Hemingway
     */
    isCurrentRole(role) {
      return this.role ? this.role.name === role.name : false
    }
  }
}
</script>

<style lang="scss" scoped>
.select-role {
  padding: 0 16rpx;

  .prompt {
    padding: 60rpx 0;
    text-align: center;
    position: relative;

    .text {
      font-size: 60rpx;
    }
  }

  .prompt::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 50%;
    transform: translateX(-50%);
    width: 300rpx;
    height: 1rpx;
    background: #eee;
  }

  .role-container {
    margin-top: 60rpx;

    .wrapper {
      margin-bottom: 24rpx;
      padding: 24rpx;
      background-color: #fff;
      font-size: 32rpx;
      border-radius: 16rpx;
      box-shadow: 0 0 12rpx 1rpx rgba(0, 0, 0, 0.05);
      position: relative;

      .mark {
        position: absolute;
        top: 24rpx;
        right: 24rpx;
        font-size: 24rpx;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
  }
}
</style>
