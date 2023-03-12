<!--
 * @Description  : markdown业务组件【监听blur事件，一键获得markdown内容中的图片及（剔除掉图片之后的）文本】
 * @Autor        : Hemingway
 * @Date         : 2022-01-11 13:23:14
 * @LastEditors  : Tyne
 * @LastEditTime : 2022-02-22 10:07:58
 * @FilePath     : /i-farm-platform/src/components/zl/i-markdown/index.vue
-->
<template>
  <div class="i-markdown">
    <markdown-editor
      v-model="content"
      :height="height"
      :initial-value="initialValue"
      initialEditType="wysiwyg"
      :options="{ hideModeSwitch: true, language: 'zh-CN' }"
      ref="markdownEditor"
      @blur="onEditorBlur"
    />
    <!-- 副本，用于跟踪文本（剔除图片内容） -->
    <markdown-editor
      style="
        visibility: hidden;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
      "
      v-model="result.content"
      :height="height"
      initialEditType="wysiwyg"
      :options="{ hideModeSwitch: true }"
      ref="markdownEditorCopy"
    />
  </div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
import '@toast-ui/editor/dist/i18n/zh-cn.js'

export default {
  name: 'IMarkdown',
  components: { MarkdownEditor },
  props: {
    // 容器高度
    height: {
      type: String,
      default: '300px'
    },
    // 默认值
    initialValue: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      // 富文本内容
      content: '',

      // 富文本解析结果
      result: {
        imageArr: [],
        content: ''
      }
    }
  },
  methods: {
    /**
     * @description: 解析html文本，返回图片及文本
     * @return {Array, String} imageArr, content
     * @author: Hemingway
     */
    onEditorBlur() {
      const html =
        this.$refs.markdownEditor.$refs.toastuiEditor.invoke('getHTML')
      const div = document.createElement('div')
      div.innerHTML = html

      const imgs = div.getElementsByTagName('img')
      const imageArr = []
      const delImgArr = []

      // image
      for (let i = 0; i < imgs.length; i++) {
        let item = imgs[i]
        if (item.src) {
          imageArr.push(item.src)
        }

        delImgArr.push(item)
      }
      this.result.imageArr = imageArr

      // content
      delImgArr.forEach(item => {
        item.parentNode.removeChild(item)
      })
      this.$refs.markdownEditorCopy.$refs.toastuiEditor.invoke(
        'setHTML',
        div.innerHTML,
        false
      )

      this.$emit('blur', {
        content: html,
        result: this.result
      })
    }
  }
}
</script>

<style scoped lang="scss">
.i-markdown {
  position: relative;
}
</style>
