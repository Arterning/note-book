<!--
 * @Description  : IAction（1.属性：loading、show-all、show-query；2.插槽：default插槽、action插槽；3.事件：on-query、on-reset）
 * @Autor        : Hemingway
 * @Date         : 2022-03-02 14:45:38
 * @LastEditors  : WuJing
 * @LastEditTime : 2022-06-15 11:16:29
 * @FilePath     : \i-farm-admin\src\components\i-action\index.vue
-->
<template>
  <div class="i-action">
    <div class="left-side" ref="leftSide">
      <!-- 默认：查询区域 -->
      <template v-if="showQuery">
        <slot></slot>
        <div class="query-wrapper">
          <div class="query">
            <el-button
              type="primary"
              @click="$emit('on-query')"
              :loading="loading"
              >查 询</el-button
            >
          </div>
          <el-button @click="$emit('on-reset')">重 置</el-button>

          <el-button
            type="text"
            @click="toggleHidden"
            v-if="showDrawerWrapper"
            >{{ !show ? '展开' : '收起' }}</el-button
          >
        </div>
      </template>
      <!-- 自定义左侧内容区 -->
      <div style="height: 52px; padding-bottom: 16px" v-else>
        <slot></slot>
      </div>
    </div>
    <div class="right-side">
      <div class="action-wrapper">
        <slot name="action"></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IAction',

  props: {
    // 查询按钮加载状态
    loading: {
      type: Boolean,
      default: false
    },
    // 是否展示全部筛选项
    showAll: {
      type: Boolean,
      default: false
    },
    // 是否展示查询按钮组合
    showQuery: {
      type: Boolean,
      default: true
    }
  },

  data() {
    return {
      show: undefined, // 是否展现全部筛选项
      leftSideItem: undefined, //
      actionItems: undefined, // i-action-item 集合
      queryWrapperItem: undefined,
      count: -1 // 首行可以放下的 i-action-item 数量
    }
  },

  computed: {
    showDrawerWrapper() {
      return (this.actionItems && this.actionItems.length) > this.count
    }
  },

  watch: {
    showAll: {
      handler(newV) {
        this.show = newV // 初始化show
      },
      immediate: true
    },

    count(newV, oldV) {
      if (oldV !== -1 && !this.show) {
        this.initHidden()
      }
    }
  },

  mounted() {
    this.init()

    window.addEventListener('resize', this.computeCount)
  },

  beforeDestroy() {
    window.removeEventListener('resize', this.computeCount)
  },

  methods: {
    init() {
      if (!this.showQuery) {
        return
      }

      this.actionItems =
        this.$refs.leftSide.getElementsByClassName('i-action-item')

      this.leftSideItem = this.$refs.leftSide

      this.queryWrapperItem =
        this.$refs.leftSide.getElementsByClassName('query-wrapper')[0]

      setTimeout(() => {
        this.computeCount()

        if (!this.show) {
          this.initHidden()
        }
      })
    },

    /**
     * @description: 计算首行可排布 i-action-item 的数量
     * @author: Hemingway
     */
    computeCount() {
      if (this.showQuery) {
        this.count = Math.floor(
          (this.leftSideItem.getBoundingClientRect().width -
            this.queryWrapperItem.getBoundingClientRect().width) /
            420
        )
      }
    },

    /**
     * @description: 初始化筛选项展现
     * @author: Hemingway
     */
    initHidden() {
      if (this.count >= 1) {
        Array.prototype.forEach.call(this.actionItems, (item, index) => {
          if (index >= this.count) {
            item.classList.add('hidden')
          } else {
            item.classList.remove('hidden')
          }
        })
      } else {
        Array.prototype.forEach.call(this.actionItems, (item, index) => {
          if (index > this.count) {
            item.classList.add('hidden')
          } else {
            item.classList.remove('hidden')
          }
        })
      }
    },

    /**
     * @description: 设置筛选项展现
     * @author: Hemingway
     */
    toggleHidden() {
      this.show = !this.show
      if (this.count >= 1) {
        Array.prototype.forEach.call(this.actionItems, (item, index) => {
          if (index >= this.count) {
            item.classList.toggle('hidden')
          }
        })
      } else {
        Array.prototype.forEach.call(this.actionItems, (item, index) => {
          if (index > this.count) {
            item.classList.toggle('hidden')
          }
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.i-action {
  display: flex;
  justify-content: flex-end;
  align-items: flex-end;
  position: relative;

  .left-side {
    flex: 1;
    display: flex;
    flex-wrap: wrap;

    .hidden {
      display: none;
    }

    .query-wrapper {
      height: 52px;
      padding-bottom: 16px;

      display: flex;

      .query {
        display: inline-block;
        width: 86px;
        margin-right: 10px;
        text-align: right;
      }
    }
  }

  .right-side {
    height: 52px;
    padding-bottom: 16px;
    display: flex;
    align-items: flex-end;

    .action-wrapper {
      margin-left: 96px;
      display: flex;
    }
  }

  // others
  ::v-deep {
    .el-input {
      width: 280px;
    }

    .el-date-editor {
      width: 280px;
    }
  }
}
</style>
