<!--
 * @Description  : 
 * @Autor        : guoxi
 * @Date         : 2021-08-04 16:32:37
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-08-05 11:42:50
 * @FilePath     : \blockchain-minicode\src\components\i-uni-f2\f2.vue
-->
<template>
  <canvas
    type="2d"
    class="f2-canvas"
    bindtouchstart="touchStart"
    bindtouchmove="touchMove"
    bindtouchend="touchEnd"
  >
  </canvas>
</template>

<script>
import F2 from '@antv/f2'

function wrapEvent(e) {
  if (!e) return
  if (!e.preventDefault) {
    e.preventDefault = function() {}
  }
  return e
}

export default {
  name: 'F2',

  /**
   * 组件的属性列表
   */
  props: {
    init: {
      type: Function,
      default: () => {}
    }
  },

  mounted() {
    const query = uni.createSelectorQuery().in(this)
    query
      .select('.f2-canvas')
      .fields({
        node: true,
        size: true
      })
      .exec(res => {
        const { node, width, height } = res[0]
        const context = node.getContext('2d')
        const pixelRatio = uni.getSystemInfoSync().pixelRatio
        // 高清设置
        node.width = width * pixelRatio
        node.height = height * pixelRatio

        const config = {
          context,
          width,
          height,
          pixelRatio
        }
        this.init(F2, config).then(data => {
          if (data) {
            this.chart = data
            this.canvasEl = data.get('el')
          }
        })
      })
  },

  /**
   * 组件的方法列表
   */
  methods: {
    touchStart(e) {
      const canvasEl = this.canvasEl
      if (!canvasEl) {
        return
      }
      canvasEl.dispatchEvent('touchstart', wrapEvent(e))
    },
    touchMove(e) {
      const canvasEl = this.canvasEl
      if (!canvasEl) {
        return
      }
      canvasEl.dispatchEvent('touchmove', wrapEvent(e))
    },
    touchEnd(e) {
      const canvasEl = this.canvasEl
      if (!canvasEl) {
        return
      }
      canvasEl.dispatchEvent('touchend', wrapEvent(e))
    }
  }
}
</script>

<style scoped>
.f2-canvas {
  width: 100%;
  height: 100%;
}
</style>
