<template>
  <Editor
    :initialValue="initialValue"
    :initialEditType="initialEditType"
    :options="options"
    :height="height"
    :previewStyle="previewStyle"
    @change="onEditorChange"
    @blur="onEditorBlur"
    ref="toastuiEditor"
  />
</template>

<script>
import defaultOptions from './default-options'

import '@toast-ui/editor/dist/toastui-editor.css'
import { Editor } from '@toast-ui/vue-editor'

export default {
  name: 'MarkdownEditor',
  components: {
    Editor: Editor
  },
  props: {
    initialValue: {
      type: String,
      default: ''
    },
    initialEditType: {
      type: String,
      default: 'markdown'
    },
    options: {
      type: Object,
      default() {
        return defaultOptions
      }
    },
    height: {
      type: String,
      required: false,
      default: '300px'
    },
    previewStyle: {
      type: String,
      default: 'vertical'
    }
  },
  data() {
    return {}
  },
  mounted() {
    this.onEditorChange()
  },
  methods: {
    onEditorChange() {
      this.$emit('input', this.$refs.toastuiEditor.invoke('getMarkdown'))
    },
    onEditorBlur() {
      this.$emit('blur')
    }
  }
}
</script>
