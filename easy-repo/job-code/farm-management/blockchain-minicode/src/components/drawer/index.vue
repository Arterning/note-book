<!--
 * @Description  : 抽屉筛选组件
 * @Autor        : Hemingway
 * @Date         : 2021-09-09 10:49:33
 * @LastEditors  : Hemingway
 * @LastEditTime : 2021-11-23 19:55:59
 * @FilePath     : \blockchain-minicode\src\components\drawer\index.vue
-->
<template>
  <uni-drawer ref="drawer" mode="right" :mask-click="true" :width="width">
    <view class="drawer">
      <scroll-view class="screen" scroll-y>
        <slot></slot>

        <view
          v-for="(item, index) in dataArray"
          :key="index"
          class="screen-wrapper"
        >
          <view class="screen-wrapper--inner">
            <view class="title" @click="handleFoldEvent(item, index)"
              ><text>{{ item.title }}</text>
              <view
                v-if="item.tags.length > (item.columnCount || 1)"
                class="icon-wrapper"
              >
                <template v-if="item.fold">
                  <text class="text">展开全部</text
                  ><uni-icons type="arrowdown" size="8"></uni-icons>
                </template>
                <template v-else>
                  <text class="text">收起</text>
                  <uni-icons type="arrowup" size="8"></uni-icons>
                </template> </view
            ></view>
            <view :class="['section', item.fold ? 'fold' : '']">
              <view
                v-for="(tag, i) in item.tags"
                :key="i"
                :class="{
                  'tag-wrapper': true,
                  'col--12': item.columnCount === 2,
                  'col--8': item.columnCount === 3
                }"
                @click="onSelect(tag, item)"
                ><span :class="{ tag: true, selected: tag.selected }">{{
                  tag.text
                }}</span></view
              >
            </view>
          </view>
        </view>
      </scroll-view>
      <view class="footer">
        <button type="plain" class="btn" @click="onReset">
          重置
        </button>
        <button type="primary" class="btn" @click="onSure">
          确定
        </button>
      </view>
    </view>
  </uni-drawer>
</template>

<script>
export default {
  name: 'Drawer',

  props: {
    // 抽屉宽度
    width: {
      type: Number,
      default: 240
    },
    // 筛选数据源，数据结构说明如下
    /* array数据结构 >>
    [
      {
        title: [String], // 标题
        tags: [Array],  // 标签 ->  见下方
        fold: [Boolean], // 是否折叠，默认为否
        isMutliple: [Boolean] // 是否多选，默认为否
        columnCount: [Number] // 每行标签的数量，默认为1（支持2、3）
      }
    ]

    tags数据结构 >>
    [
      {
        text: [String], // 标签的文本
        code: [String], // 标签的值
        selected: [Boolean] // 标签是否被选，默认为否
      }
    ]
    */
    array: {
      type: Array,
      default: () => []
    }
  },

  data() {
    return {
      dataArray: []
    }
  },

  computed: {},

  watch: {
    array(newVal) {
      this.dataArray = newVal
    }
  },

  created() {},

  mounted() {},

  methods: {
    /**
     * @description: 折叠和展开事件
     * @param {Object} item
     * @param {index} index
     * @author: Hemingway
     */
    handleFoldEvent(item, index) {
      item.fold = !item.fold
      this.$set(this.dataArray, index, item)
    },

    /**
     * @description: tag选择事件
     * @param {Object} tag 目标
     * @param {Object} item 父级目标
     * @author: Hemingway
     */
    onSelect(tag, item) {
      if (!tag.selected && !item.isMutliple) {
        item.tags.forEach(tag => {
          tag.selected = false
        })
      }
      tag.selected = !tag.selected

      // 通知父组件筛选项
      const selectedList = item.tags.filter(item => item.selected)
      const codeList = selectedList.map(tag => tag.code)
      const textList = selectedList.map(tag => tag.text)
      this.$emit('onSelect', {
        field: item.field,
        codeList,
        textList
      })
    },

    /**
     * @description: 重置筛选项
     * @author: Hemingway
     */
    onReset() {
      this.dataArray.forEach(item => {
        item.tags &&
          item.tags.forEach(tag => {
            tag.selected = false
          })
      })

      // 通知父组件筛选项
      this.$emit('onReset')
    },

    /**
     * @description: 确认筛选项
     * @author: Hemingway
     */
    onSure() {
      // 关闭抽屉
      this.$refs.drawer.close()

      this.$emit('onSure')
    },

    /**
     * @description: 打开抽屉
     * @author: Hemingway
     */
    open() {
      this.$refs.drawer.open()
    }
  }
}
</script>

<style lang="scss" scoped>
.drawer {
  box-sizing: border-box;
  height: 100%;
  padding-bottom: 144rpx;

  .screen {
    height: 100%;

    .screen-wrapper {
      padding-bottom: 24rpx;
      background-color: $uni-bg-color-page;

      .screen-wrapper--inner {
        padding-bottom: 12rpx;
        background-color: #fff;
        border-radius: 0 0 24rpx 24rpx;

        .title {
          padding: 0 24rpx;
          height: 84rpx;
          display: flex;
          justify-content: space-between;
          line-height: 84rpx;
          font-size: 26rpx;
          font-weight: bold;

          .icon-wrapper {
            display: flex;
            align-items: center;
            color: #8b8b8b;

            .text {
              margin-right: 4rpx;
              font-size: 24rpx;
              font-weight: lighter;
            }
          }
        }

        .section {
          margin: 0 -16rpx;
          padding: 0 24rpx;
          display: flex;
          flex-wrap: wrap;
          align-content: flex-start;

          .tag-wrapper {
            width: 100%;
            height: 68rpx;
            padding: 0 16rpx;
            box-sizing: border-box;

            .tag {
              box-sizing: border-box;
              display: block;
              height: 48rpx;
              padding: 0 16rpx;
              line-height: 46rpx;
              text-align: center;
              font-size: 24rpx;
              border-radius: 24rpx;
              border: solid 1rpx #e0e0e0;
              overflow: hidden;
              white-space: nowrap;
              text-overflow: ellipsis;
            }

            .selected {
              // color: $theme-color;
              // border-color: $theme-color;
              color: #70b913;
              border-color: #70b913;
              background-color: rgba(112, 185, 19, 0.1);
            }
          }

          .col--12 {
            width: 50%;
          }

          .col--8 {
            width: 1/3;
          }
        }

        .fold {
          max-height: 68rpx;
          overflow: hidden;
        }
      }
    }

    .screen-wrapper:last-of-type {
      padding-bottom: 0;

      .screen-wrapper--inner {
        border-radius: 0;
      }
    }
  }

  .screen::-webkit-scrollbar {
    display: none;
  }

  .footer {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    // padding: 32rpx 48rpx;
    padding: 32rpx 0;
    background-color: $uni-bg-color-page;
    display: flex;
    justify-content: space-between;

    .btn {
      flex: 1;
      margin: 0rpx 36rpx;
      // margin: 0;
      // width: 168rpx;
      height: 80rpx;
      font-size: 32rpx;
      background-color: $uni-bg-color-btn;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 40rpx;
    }

    .btn:first-of-type {
      color: $theme-color;
      background-color: #fff;
      border: 1rpx solid $theme-color;
    }

    .btn::after {
      display: none;
    }
  }
}
</style>
