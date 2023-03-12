<!--
 * @Description  : 搜索栏组件
 * @Autor        : Hemingway
 * @Date         : 2021-07-21 17:29:23
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-16 17:13:23
 * @FilePath     : \blockchain-minicode\src\components\search-bar\index.vue
-->
<template>
  <view :class="{ 'search-bar': true, shadow }">
    <uni-search-bar
      v-model="searchValue"
      :placeholder="placeholder"
      cancel-button="none"
      class="search"
      @confirm="onSearch"
      @blur="onBlur"
      @focus="onFocus"
      @input="onInput"
      @clear="onClear"
    />
    <slot></slot>
  </view>
</template>

<script>
export default {
  name: 'SearchBar',

  options: {
    styleIsolation: 'shared'
  },

  props: {
    placeholder: {
      type: String,
      default: ''
    },
    shadow: {
      type: Boolean,
      default: false
    }
  }, // https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/wxml-wxss.html

  data() {
    return {
      searchValue: ''
    }
  },

  computed: {},

  created() {},

  mounted() {},

  methods: {
    onSearch(e) {
      this.searchValue = e.value.trim()
      this.$emit('search', this.searchValue)
    },
    onBlur() {
      this.$emit('blur')
    },
    onFocus() {
      this.$emit('focus')
    },
    onInput() {
      this.$emit('input', this.searchValue)
    },
    onClear() {
      this.$emit('clear')
    }
  }
}
</script>

<style lang="scss" scoped>
.search-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1;
  padding: 12rpx 32rpx;
  background-color: #fff;
  display: flex;
  align-items: center;

  .search {
    flex: 1;

    ::v-deep .uni-searchbar {
      padding: 0;

      .uni-searchbar__box {
        height: 64rpx;
        padding: 0 16rpx 0 0;
        background-color: #f5f5f5 !important;
        border: none;
        border-radius: 32rpx !important;
      }
    }
  }
}

.shadow {
  box-shadow: 0 24rpx 12rpx 1rpx rgba(0, 0, 0, 0.05);
}
</style>
