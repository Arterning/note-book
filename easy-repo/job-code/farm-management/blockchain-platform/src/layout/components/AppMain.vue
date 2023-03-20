<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-09-23 16:15:57
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-23 16:17:23
 * @FilePath     : \blockchain-platform\src\layout\components\AppMain.vue
-->
<template>
  <section class="app-main">
    <div class="page-main" :style="{ width: pageWidth }">
      <transition name="fade-transform" mode="out-in">
        <keep-alive :include="cachedViews">
          <router-view :key="key" />
        </keep-alive>
      </transition>
      <div class="copyright">版权所有 © 2019-2022 中联智慧农业股份有限公司</div>
    </div>
  </section>
</template>

<script>
const { body } = document
export default {
  name: 'AppMain',
  data() {
    return {
      pageWidth: '100%'
    }
  },
  beforeMount() {
    window.addEventListener('resize', this.contractWidth)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.contractWidth)
  },
  computed: {
    cachedViews() {
      return this.$store.state.tagsView.cachedViews
    },
    key() {
      return this.$route.path
    }
  },
  methods: {
    /**
     * @description: 监听屏幕的宽度
     * @author: qinjj
     */
    contractWidth() {
      const rect = body.getBoundingClientRect().width
      if (rect < 1200) {
        this.pageWidth = '1200px'
      } else {
        this.pageWidth = '100%'
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.app-main {
  background: #f5f5f5;
  width: 100%;
  height: 100vh;
  position: relative;
  overflow: hidden;
  background: #f5f5f5;

  .page-main {
    height: 100%;
    padding: 8px 8px 32px;

    .copyright {
      position: absolute;
      left: 50%;
      bottom: 14px;
      transform: translate(-50%, 50%);
      font-size: 12px;
      color: #c7c7c7;
    }
  }
}

.fixed-header + .app-main {
  padding-top: 44px;
}

.hasTagsView {
  .app-main {
    /* 84 = navbar + tags-view = 44 + 40 */
    min-height: calc(100vh - 84px);
  }

  .fixed-header + .app-main {
    padding-top: 84px;
  }
}
</style>
