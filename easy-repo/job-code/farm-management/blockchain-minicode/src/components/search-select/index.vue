<template>
  <view style="position: relative;" :class="size === 'small' ? 'small' : ''">
    <view class="uni-flex" :class="{ unishadow: uniShadow }">
      <view class="uni-search-form uni-circular uni-background">
        <input
          v-model="inputVal"
          type="text"
          :placeholder="placeholder"
          @input="handleSearch"
          @focus="onFocus"
        />
        <!-- <icon v-if="isShowClearIcon" type="clear" size="16" class="uni-icon-clear" @tap="clearInputValue"/> -->
      </view>
    </view>
    <view v-if="isShowSelect">
      <view v-if="inputIsShow" class="uni-combox__selector">
        <scroll-view scroll-y="true" class="uni-combox__selector-scroll">
          <view v-if="loading" class="uni-combox__selector-empty ">
            <view>加载中··· </view>
          </view>
          <view
            v-if="!loading && dataListLength === 0"
            class="uni-combox__selector-empty "
            @tap="closeDielog"
          >
            <view>{{ emptyTips }}</view>
          </view>
          <view
            v-for="(item, index) in dataList"
            :key="index"
            class="uni-combox__selector-item"
            @tap="onSelectorClick(index)"
          >
            <text>{{ showValue(item) }}</text>
          </view>
        </scroll-view>
      </view>
    </view>
  </view>
</template>
<script>
export default {
  name: 'SearchSelect',
  props: {
    // placeholder: {
    // 	type: String,
    // 	default: '请输入信息'
    // },
    infoList: {
      type: Array,
      default() {
        return []
      }
    },
    value: {
      type: String,
      default: ''
    },
    showValue: {
      type: Function,
      default: function() {}
    },
    emptyTips: {
      type: String,
      default: '暂无数据'
    },
    loading: {
      type: Boolean,
      default: false
    },
    btnStyleColor: {
      type: Object,
      default() {
        return {}
      }
    },
    uniShadow: {
      type: Boolean,
      default: true
    },
    type: {
      type: String,
      default: 'primary'
    },
    size: {
      type: String,
      default: 'medium'
    },
    isShowSelect: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      inputIsShow: false,
      inputVal: '',
      dataList: [],
      placeholder: '请输入信息'
    }
  },
  computed: {
    dataListLength() {
      return this.dataList.length
    }
  },
  watch: {
    value: {
      handler(newVal) {
        this.inputVal = newVal
      },
      immediate: true
    },
    infoList: {
      handler(newVal) {
        this.dataList = newVal
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    closeDielog() {
      this.inputIsShow = false
    },
    onFocus() {
      if (this.dataList.length) {
        this.inputIsShow = true
      }
    },
    onSelectorClick(index) {
      this.inputIsShow = false
      // this.inputVal = this.showValue
      this.$emit('change', this.dataList[index])
    },

    handleSearch() {
      this.inputIsShow = true
      this.$emit('handleSearch', this.inputVal)
    },
    clearInputValue() {
      this.inputVal = ''
    }
  }
}
</script>
<style scoped lang="scss">
// $selectWidth: 75%; // 下拉选择框宽度
.small {
  transform: scale(0.9, 0.9);
}

.uni-primary {
  background-color: $uni-color-primary;
}

.uni-success {
  background-color: #67c23a;
}

.uni-warning {
  background-color: $uni-color-warning;
}

.uni-error {
  background-color: $uni-color-error;
}

.hover {
  transition: all 0.6s;
  transform: scale(0.8, 0.8);
}

.uni-flex {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10rpx 20rpx;
  box-sizing: border-box;
}

.uni-search-form {
  position: relative;
  width: 100%;
  margin: 10rpx;
  font-size: 24rpx;
  // color: #999999;
}

.uniRound {
  border-radius: 5px;
}
// .uni-circular {
//   border-radius: 100rpx;
// }
.uni-icon-position {
  position: absolute;
  top: 50%;
  left: 26rpx;
  transform: translate(0, -50%);
}

.uni-icon-clear {
  position: absolute;
  top: 50%;
  right: 26rpx;
  transform: translate(0, -50%);
}
// .uni-background {
//   background-color: #f5f5f5;
// }

/* button */
.uni-action {
  width: 120rpx;
  height: 66rpx;
}

.uni-cu-btn {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #fff;
  font-size: 28rpx;
  border: none;
  // background-color: $uni-color-primary;
}

.uni-shadow-blur {
  box-shadow: 0rpx 1rpx 10rpx #c8c7cc;
}

.uni-round {
  border-radius: 100rpx;
}

.uni-combox__selector {
  position: absolute;
  top: 100rpx;
  left: 40rpx;
  box-sizing: border-box;
  width: 75%; // 下拉框宽度
  background-color: #fff;
  border-radius: 6px;
  box-shadow: #ddd 4px 4px 8px, #ddd -4px -4px 8px;
  z-index: 999;
  text-align: center;
}

.uni-combox__selector-scroll {
  max-height: 200px;
  box-sizing: border-box;
}

.uni-combox__selector::before {
  content: '';
  position: absolute;
  width: 0;
  height: 0;
  border-bottom: solid 6px #fff;
  border-right: solid 6px transparent;
  border-left: solid 6px transparent;
  left: 50%;
  top: -6px;
  margin-left: -6px;
}

.uni-combox__selector-empty {
  text-align: center;
  color: #8f8f94;
  padding: 20rpx 0;
  font-size: 28rpx;
}

.uni-combox__selector-item {
  /* #ifdef APP-NVUE */
  display: flex;

  /* #endif */
  font-size: 24rpx;
  margin: 0 10px;
  padding: 20rpx 10rpx;
  color: #808080;
}

.uni-combox__selector-item:hover {
  background-color: #ddd;
}

.uni-combox__selector-empty:last-child,
.uni-combox__selector-item:last-child {
  border-bottom: none;
}
</style>
