<!--
 * @Description  : 种植过程（生育期详情） - 为大通湖定制
 * @Autor        : Hemingway
 * @Date         : 2021-12-16 19:09:33
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-06-13 18:18:23
 * @FilePath     : \blockchain-h5\src\pages\plant\datonghu\Stage.vue
-->
<template>
  <view class="stage">
    <steps :steps="steps" :current="steps.length - 1">
      <template #date="{date}">
        <view class="date__up">
          {{ date | dateFormatter('up') }}
        </view>
        <view class="date__down">
          {{ date | dateFormatter('down') }}
        </view>
      </template>
      <template #body="{body}">
        <view v-if="body.desc && body.desc.length > 0" class="desc-wrapper">
          <view
            v-for="(text, index) in body.desc"
            :key="index"
            class="desc-p"
            >{{ text }}</view
          >
        </view>
        <view v-if="body.imgs" class="img-wrapper">
          <img
            v-for="(img, index) in body.imgs"
            :key="img"
            class="img"
            :src="img"
            @click="preViewImg(index, body.imgs)"
          />
        </view>
      </template>
    </steps>
  </view>
</template>

<script>
import Steps from '@/components/steps'
import { preViewImg } from '@/utils/tool'
export default {
  name: 'Plant',

  components: {
    Steps
  },

  filters: {
    /**
     * @description: 根据UI设计稿处理日期展示
     * @param {String} date
     * @param {String} flagStr
     * @return {String}
     * @author: Hemingway
     */
    dateFormatter(date, flagStr) {
      const isDate = !isNaN(Date.parse(date))
      if (isDate) {
        if (flagStr === 'up') {
          return date.slice(5)
        } else {
          return date.slice(0, 4)
        }
      } else {
        if (flagStr === 'up') {
          return date.slice(0, 3)
        } else {
          return date.slice(3)
        }
      }
    }
  },

  data() {
    return {
      steps: [] // 种植步骤时间线
    }
  },

  created() {
    const { name, steps } = this.$Route.query
    uni.setNavigationBarTitle({
      title: name
    })
    this.steps = JSON.parse(decodeURIComponent(steps))
  },

  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    }
  }
}
</script>

<style lang="scss" scoped>
// 时间线
.stage {
  padding: 40rpx 32rpx;
  background-color: #fff;

  .desc-wrapper {
    padding-top: 24rpx;

    .desc-p {
      word-break: break-all;
      color: #8b8b8b;
    }
  }

  .img-wrapper {
    padding: 32rpx 0 0;
    display: flex;
    flex-wrap: wrap;

    .img {
      display: block;
      margin-right: 16rpx;
      width: 120rpx;
      height: 120rpx;
    }
  }
}
</style>
