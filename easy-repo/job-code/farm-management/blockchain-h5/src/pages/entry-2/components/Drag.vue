<!--
 * @Description  : 可拖拽屏幕吸附组件
 * @Autor        : Hemingway
 * @Date         : 2021-11-19 14:16:47
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 10:33:43
 * @FilePath     : \blockchain-h5\src\pages\entry-2\components\Drag.vue
-->
<template>
  <view
    ref="drag"
    class="drag"
    @touchstart="onDown"
    @touchmove.prevent="onMove"
    @touchend="onEnd"
    @click="onClick"
    ><text class="text">{{ data.title }}</text>
    <text class="text copy">{{ data.title }}</text></view
  >
</template>

<script>
export default {
  name: 'Drag',

  props: {
    data: {
      type: Object,
      default: () => {}
    },
    qrcode: {
      type: String,
      default: ''
    }
  },

  data() {
    return {
      touch: { x: 0, y: 0 }, // 触摸位置
      dom: { x: 0, y: 0 }, // 触摸时元素位置
      movement: { x: 0, y: 0 }, // 移动量
      isBoundary: true // 是否吸附至边缘
    }
  },

  methods: {
    onDown(e) {
      console.log('---down: ', e)

      this.$el.style.opacity = 0.8
      this.$el.style.transform = `scale3d(1.1, 1.1, 1.1)`

      // 记录触摸位置
      this.touch.x = e.touches[0].clientX
      this.touch.y = e.touches[0].clientY

      // 记录触摸时，元素的位置
      this.dom.x = e.currentTarget.offsetLeft
      this.dom.y = e.currentTarget.offsetTop
      console.log('dom位置 = ', this.dom.x, this.dom.y)
    },

    onMove(e) {
      // console.log('---move: ', e)

      // 移动量
      this.movement.x = e.touches[0].clientX - this.touch.x
      this.movement.y = e.touches[0].clientY - this.touch.y

      // 初始化正常状态
      let leftFlag = true
      let rightFlag = true
      let topFlag = true
      let bottomFlag = true

      // 判断移动量，状态赋予
      if (this.dom.x + this.movement.x < 10) {
        leftFlag = false
      }
      if (
        this.dom.x + this.movement.x >
        document.documentElement.clientWidth - this.$el.clientWidth - 10
      ) {
        rightFlag = false
      }
      if (this.dom.y + this.movement.y < 10) {
        topFlag = false
      }
      if (
        this.dom.y + this.movement.y >
        document.documentElement.clientHeight - this.$el.clientHeight - 10
      ) {
        bottomFlag = false
      }

      // 判断左右出界情况
      if (leftFlag && rightFlag) {
        this.$el.style.left = this.dom.x + this.movement.x + 'px'
      } else if (!leftFlag) {
        this.$el.style.left = 8 + 'px'
      } else {
        this.$el.style.right = 8 + 'px'
      }
      // 判断上下出界情况
      if (topFlag && bottomFlag) {
        this.$el.style.top = this.dom.y + this.movement.y + 'px'
      } else if (!topFlag) {
        this.$el.style.top = 8 + 'px'
      } else {
        this.$el.style.bottom = 8 + 'px'
      }
    },

    onEnd(e) {
      // console.log('---end: ', e)

      let isBoundary = true

      if (
        this.dom.x + this.movement.x >= 10 &&
        this.dom.x + this.movement.x <=
          document.documentElement.clientWidth - this.$el.clientWidth - 10 &&
        this.dom.y + this.movement.y >= 10 &&
        this.dom.y + this.movement.y <=
          document.documentElement.clientHeight - this.$el.clientHeight - 10
      ) {
        isBoundary = false
      }

      if (!isBoundary) {
        // 判断手势倾向
        if (Math.abs(this.movement.x) >= Math.abs(this.movement.y)) {
          // 水平倾向
          if (
            e.currentTarget.offsetLeft + 0.5 * this.$el.clientWidth >=
            0.5 * document.documentElement.clientWidth
          ) {
            this.$el.style.left = ''
            this.$el.style.right = 8 + 'px'
          } else {
            this.$el.style.right = ''
            this.$el.style.left = 8 + 'px'
          }
        } else {
          // 垂直倾向
          if (
            e.currentTarget.offsetTop + 0.5 * this.$el.clientHeight >=
            0.5 * document.documentElement.clientHeight
          ) {
            this.$el.style.top = ''
            this.$el.style.bottom = 8 + 'px'
          } else {
            this.$el.style.bottom = ''
            this.$el.style.top = 8 + 'px'
          }
        }
      }

      this.$el.style.opacity = 1
      this.$el.style.transform = `none`
    },

    onClick() {
      window.open(this.data.imgUrls, '_blank')
      // this.$Router.push({
      //   path: '/graphic',
      //   query: {
      //     id: this.qrcode,
      //     title: this.data.title,
      //     imgUrls: this.data.imgUrls
      //   }
      // })
    }
  }
}
</script>

<style lang="scss" scoped>
.drag {
  width: 196rpx;
  height: 166rpx;
  position: fixed;
  right: 16rpx;
  bottom: 218rpx;
  z-index: 20;
  background-image: url('/blockchain/h5/static/img/home/rice.png');
  background-size: 100% 100%;
  // transition: all 0.1s ease-in-out 0s;

  .text {
    position: absolute;
    bottom: 30rpx;
    left: 18rpx;
    right: 18rpx;
    font-size: 22rpx;
    font-weight: bold;
    color: #d8ad00;
    letter-spacing: 2rpx;
    transform: skew(-6deg);
    text-align: center;
  }

  .copy {
    color: #fff;
    bottom: 34rpx;
    text-align: center;
  }
}
</style>
