<!--
 * @Description  : IDialog（1.属性：visible、title；2.插槽：title插槽、default插槽、footer插槽；3.事件：close、scroll）
 * @Autor        : Hemingway
 * @Date         : 2022-04-21 14:31:24
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-23 17:42:39
 * @FilePath     : \blockchain-admin\src\components\i-dialog\index.vue
-->
<template>
  <el-dialog
    :visible.sync="dialogVisible"
    ref="dialog"
    @close="$emit('update:visible', dialogVisible)"
    :show-close="showClose"
  >
    <template #title>
      <span class="el-dialog__title">
        <!-- title插槽 -->
        <slot name="title">
          {{ title }}
        </slot>
      </span>
    </template>
    <!-- 默认插槽 -->
    <slot></slot>
    <template #footer>
      <!-- footer插槽 -->
      <slot name="footer"> </slot>
    </template>
  </el-dialog>
</template>

<script>
export default {
  name: 'IDialog',

  props: {
    // 是否可见
    visible: {
      type: Boolean,
      default: false
    },
    // 是否可以关闭弹框
    showClose: {
      type: Boolean,
      default: true
    },
    // 标题
    title: {
      type: String,
      default: ''
    }
  },

  data() {
    return {
      dialogVisible: false,
      hasAddListener: false
    }
  },

  computed: {},

  watch: {
    visible: {
      handler(newVal) {
        {
          this.dialogVisible = newVal

          if (!this.hasAddListener && newVal) {
            this.$nextTick(() => {
              this.$refs.dialog.$el
                .getElementsByClassName('el-dialog__body')[0]
                .addEventListener('scroll', this.onBodyScroll)
            })

            this.hasAddListener = true
          }
        }
      },
      immediate: true
    },
    dialogVisible(newVal) {
      if (!newVal) {
        this.$emit('close')
      }
    }
  },

  methods: {
    /**
     * @description: el-dialog__body 滚动事件
     * @param {Object} event
     * @author: Hemingway
     */
    onBodyScroll(event) { // eslint-disable-line
      // console.log('onBodyScroll: ', event)
      this.$emit('scroll')
    }
  }
}
</script>

<style lang="scss" scoped>
::v-deep .el-dialog {
  width: fit-content;
  min-width: 640px;

  .el-dialog__header {
    border-bottom: 1px solid #eee;
  }
  .el-dialog__body {
    max-height: calc(70vh - 122px);
    overflow-y: auto;
    padding: 32px 80px 32px 40px;
  }
  .el-dialog__body::-webkit-scrollbar {
    display: none;
  }
  .el-dialog__footer {
    border-top: 1px solid #eee;
    text-align: center;
  }

  // others
  .el-form {
    .el-form-item__label {
      width: 112px;
    }
    .el-form-item__content {
      margin-left: 112px;
    }
    .el-input {
      width: 376px;
    }
    .el-textarea {
      width: 376px;
    }
    .el-date-editor {
      width: 376px;
    }

    .tooltip {
      position: absolute;
      top: 18px;
      transform: translateY(-50%);
      right: 0;
      font-size: 16px;
    }
  }
}
</style>
