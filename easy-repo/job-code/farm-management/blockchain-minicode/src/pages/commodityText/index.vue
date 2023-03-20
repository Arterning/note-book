<!--
 * @Description  : 商品文字介绍
 * @Autor        : WuJing
 * @Date         : 2021-11-10 10:57:54
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-11-22 09:52:35
 * @FilePath     : \blockchain-minicode\src\pages\commodityText\index.vue
-->
<template>
  <view class="commodityText">
    <textarea
      maxlength="500"
      placeholder="请输入商品文字介绍，如：中联数字大米是......"
      :value="text"
      @input="inputValue"
    >
    </textarea>
    <view class="number"
      ><text>{{ fontNum }}</text
      >/500</view
    >
    <view class="bottom">
      <bottom-button text="提交" @click="onSubmit"></bottom-button>
    </view>
  </view>
</template>

<script>
import BottomButton from '@/components/bottom-button'
import { updateMaterial } from '@/api/user'
export default {
  components: {
    BottomButton
  },
  data() {
    return {
      fontNum: 0, //字数
      text: '' //输入内容
    }
  },
  computed: {
    // 获取当前角色
    role() {
      return this.$store.state.user.role
    }
  },
  mounted() {
    console.log('路由传参文字描述')
    this.textQuery()
  },
  methods: {
    inputValue(e) {
      this.text = e.detail.value
      this.fontNum = this.text.length //获取字数显示
    },
    /**
     * @description: 商品文字介绍提交事件
     * @author: Hemingway
     */
    async onSubmit() {
      uni.showModal({
        title: '提示',
        content: '确认信息填写无误？',
        success: async res => {
          if (res.confirm) {
            try {
              const { code } = await updateMaterial({
                enterpriseId: this.role.enterpriseId,
                commodityDesc: this.text //商品文字介绍
              })
              if (code === '0') {
                uni.showToast({
                  title: '提交成功',
                  icon: 'success',
                  duration: 2000,
                  complete: () => {
                    setTimeout(() => {
                      this.$bus.$emit('refresh')
                      this.$Router.back()
                    }, 1500)
                  }
                })
              }
            } catch (error) {
              console.log('商品文字介绍提交 error = ', error)
            }
          }
        }
      })
    },
    /**
     * @description: 商品素材查询
     * @param {*}
     * @return {*}
     * @author: WuJing
     * @example:
     */

    textQuery() {
      this.text = this.$Route.query.commodityText
      this.fontNum = this.text.length //获取字数显示
    }
  }
}
</script>

<style lang="scss" scoped>
.commodityText {
  margin-top: 20rpx;

  textarea {
    background: #fff;
    margin: auto;
    width: 86%;
    height: 500rpx;
    border-radius: 10rpx;
    padding: 30rpx;
    font-size: 26rpx;
  }

  .number {
    position: relative;
    right: 30rpx;
    bottom: 30rpx;
    float: right;
    font-size: 20rpx;
    color: #808080;
  }

  .bottom {
    //定位到底部按钮
    position: absolute;
    bottom: 0;
    width: 100%;
  }
}
</style>
