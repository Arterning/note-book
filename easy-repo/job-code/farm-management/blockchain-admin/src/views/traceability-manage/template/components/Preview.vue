<!--
 * @Description  : 预览组件
 * @Autor        : SunXiuqiong
 * @Date         : 2022-06-08 11:26:44
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-26 18:21:46
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\template\components\Preview.vue
-->
<template>
  <div class="main">
    <span class="title">预览</span>
    <div class="middle-container">
      <div class="pre-template" :style="{ height: iframeBox + 'px' }">
        <iframe
          name="iframe_Preview"
          scrolling="auto"
          frameborder="0"
          border="0"
          :src="iframeUrl"
          ref="iframe"
          id="h5Page"
        ></iframe>
      </div>

      <span class="template-title">模 板 一</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Preview',
  data() {
    return {
      iframeUrl: null,
      iframeBox: window.innerHeight - 300
    }
  },
  created() {
    this.iframeUrl = process.env.VUE_APP_BASE_URL + '/blockchain/h5/'
    // this.iframeUrl = 'https://localhost:8080/blockchain/h5/1'
  },

  methods: {
    // select-data组件调用此方法改变iframe内部的布局
    postMessage(templateData) {
      const iframe = document.querySelector('#h5Page')
      console.log('pciframe', iframe)
      // 处理兼容行问题
      if (iframe.attachEvent) {
        iframe.attachEvent('onload', () => {
          this.$refs.iframe.contentWindow.postMessage(templateData, '*')
        })
      } else {
        iframe.onload = () => {
          this.$refs.iframe.contentWindow.postMessage(templateData, '*')
        }
      }
      console.log('templateData', templateData)
      this.$refs.iframe.contentWindow.postMessage(templateData, '*')
    }
  }
}
</script>

<style scoped lang="scss">
.main {
  position: relative;
  .title {
    font-weight: bold;
    font-size: 14px;
    display: block;
    margin: 4px;
  }

  .middle-container {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
  }
  .pre-template {
    width: 375px;
    height: 580px;
    border: 1px solid #ccc;

    iframe {
      width: 100%;
      height: 100%;
    }
  }

  .template-title {
    display: block;
    padding: 8px 0;
    text-align: center;
    color: #212121;
    font-size: 14px;
  }
}
</style>
