<!--
 * @Description  : 单页面内容区组件
 * @Autor        : Hemingway
 * @Date         : 2022-03-01 19:19:58
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-22 16:00:55
 * @FilePath     : \blockchain-platform\src\components\i-tab\index.vue
-->
<template>
  <div style="height: 100%; position: ralative">
    <!-- 父组件可监听 tab-click 事件 -->
    <el-tabs
      @tab-click="$emit('tab-click', $event)"
      class="i-tab"
      v-model="tabName"
    >
      <template v-if="labels.length === 1">
        <el-tab-pane :label="item" v-for="(item, index) in labels" :key="index">
          <!-- 当tab长度为1时，使用默认插槽 -->
          <slot>
            <slot :name="index"></slot>
          </slot>
        </el-tab-pane>
      </template>
      <template v-else>
        <el-tab-pane
          :label="item"
          v-for="(item, index) in labels"
          :key="index"
          :name="index + ''"
        >
          <!-- 当tab长度大于1时，使用labels数组的下标标注插槽名称 -->
          <slot :name="index"></slot>
        </el-tab-pane>
      </template>
    </el-tabs>

    <div class="right-side">
      <!-- 注意，请使用高度为32px的控件（size选用small类型） -->
      <slot name="action"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ITab',

  data() {
    return {
      tabName: ''
    }
  },

  props: {
    // 最多支持4个tab；当tab长度为1时，label可传数组或字符串，如 ['label']，或 'label'；当tab长度大于1时，label数组必传，如 ['label1', 'label2']
    label: {
      type: [Array, String],
      default: '',
      validator: function (value) {
        const arr = value instanceof Array ? value : [value]
        return arr.length < 5
      }
    },
    activeName: {
      type: String,
      default: '0'
    }
  },

  computed: {
    labels() {
      return this.label instanceof Array
        ? this.label
        : [this.label || this.$route.meta.title]
    }
  },

  watch: {
    activeName: {
      handler: function () {
        this.tabName = this.activeName
      },
      immediate: true
    }
  }
}
</script>

<style lang="scss" scoped>
.i-tab {
  height: 100%;
  padding: 0 16px;

  ::v-deep {
    .el-tabs__header {
      margin-bottom: 20px;

      .el-tabs__item {
        height: 44px;
        line-height: 44px;
      }
    }

    .el-tabs__content {
      height: calc(100% - 64px);

      .el-tab-pane {
        height: 100%;
        padding-bottom: 16px;
        overflow-y: scroll;
      }
      .el-tab-pane::-webkit-scrollbar {
        display: none;
      }
    }
  }
}

.right-side {
  position: absolute;
  top: 6px;
  right: 16px;
}
</style>
