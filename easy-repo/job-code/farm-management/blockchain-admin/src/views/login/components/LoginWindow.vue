<!--
 * @Description  : 窗口组件
 * @Autor        : Hemingway
 * @Date         : 2021-12-14 16:41:03
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-02 14:57:56
 * @FilePath     : \blockchain-admin\src\views\login\components\LoginWindow.vue
-->
<template>
  <div class="login-window"><slot></slot></div>
</template>

<script>
export default {
  name: 'LoginWindow'
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #283443;
$input: #212121;
$cursor: #00c853;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-window .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-window {
  .el-input {
    display: inline-block;
    width: 100%;
    border: 1px solid rgb(222, 220, 220);
    border-radius: 5px;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $input;
      height: 48px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style>

<style lang="scss" scoped>
.login-window {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 480px;
  max-width: 100%;
  // padding: 30px 35px;
  overflow: hidden;
  // background-color: #fff;
}
</style>
