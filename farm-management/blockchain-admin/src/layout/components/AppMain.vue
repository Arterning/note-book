<template>
  <section class="app-main">
    <div class="page-main" :style="{ width: pageWidth }">
      <transition name="fade-transform" mode="out-in">
        <keep-alive :include="cachedViews">
          <router-view :key="key" v-if="isRouterAlive" />
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
      isRouterAlive: true,
      pageWidth: '100%'
    }
  },

  provide() {
    // 父组件中通过provide来提供变量，在子组件中通过inject来注入变量
    return {
      reload: this.reload
    }
  },

  computed: {
    cachedViews() {
      return this.$store.state.tagsView.cachedViews
    },
    key() {
      return this.$route.path
    }
  },

  beforeMount() {
    window.addEventListener('resize', this.contractWidth)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.contractWidth)
  },

  methods: {
    /**
     * @description: 页面刷新的实现
     * @author: Hemingway
     */
    reload() {
      this.isRouterAlive = false
      this.$nextTick(function () {
        this.isRouterAlive = true
      })
    },
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

  .no-padding {
    padding: 0;
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
