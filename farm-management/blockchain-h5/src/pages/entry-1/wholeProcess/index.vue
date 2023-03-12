<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-07-07 17:09:37
 * @FilePath     : \blockchain-h5\src\pages\entry-1\wholeProcess\index.vue
-->
<template>
  <view class="query-record">
    <view class="common-box">
      <view class="process-content">
        <steps :steps="steps" :current="steps.length - 1">
          <template #date="{date}">
            <view class="date__up">
              {{ date | dateFormatter('up') }}
            </view>
            <view class="date__down">
              {{ date | dateFormatter('down') }}
            </view>
          </template>
          <template #scene="{scene}">
            <view class="weather-box">
              <view class="weather-icon" v-html="scene.weather"> </view>
              <text class="temper">{{ scene.temper }}℃</text>
            </view>
          </template>
          <template #body="{body}">
            <view v-if="body.desc" class="desc-wrapper">
              <view
                v-for="(text, index) in body.desc"
                :key="index"
                class="desc-p"
                >{{ text }}</view
              >
            </view>
            <view v-if="body.imgs" class="img-wrapper">
              <image
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
    </view>
    <view class="back-btn" @click="handleBack">
      <u-icon name="close" size="20" color="white" class="icon"></u-icon>
    </view>
  </view>
</template>

<script>
import Steps from '@/components/steps'
import { preViewImg, getSteps } from '@/utils/tool'
import { getPlantInfo } from '@/api/myRice'
export default {
  name: 'wholeProcess',
  components: { Steps },
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
      steps: [], // 加工步骤时间线
      farmRecords: []
    }
  },

  computed: {
    allSteps() {
      return this.$store.getters.allSteps
    }
  },

  watch: {
    '$store.state.app.configurationInformation': {
      handler(newValue) {
        let menuTrue = newValue.find(res => res.moduleId === '4')?.subTree
        getSteps(menuTrue).then(res => {
          this.steps = res
        })
      },
      deep: true
    }
  },

  created() {
    this.steps = this.allSteps
  },

  mounted() {},

  methods: {
    preViewImg(i, images) {
      preViewImg(i, images)
    },
    /**
     * @description: 获取页面初始化信息
     * @author: qjj
     */
    async plantInfo() {
      const { data } = await getPlantInfo({
        qrCode: this.qrcode
      })
      if (!data) return
      this.farmRecords = data.farmRecords || []
    },
    /**
     * @description: // 回退按钮
     * @author: qjj
     */
    handleBack() {
      window.history.go(-1)
    }
  }
}
</script>

<style lang="scss" scoped>
.query-record {
  background-color: #fafafa;
  padding: 16rpx;

  .common-box {
    border: 1px solid #eeeeee;
    background-color: white;
    border-radius: 16rpx;
    width: 100%;
    box-sizing: border-box;
    padding: 38rpx;

    // 时间线
    .process-content {
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

      .weather-box {
        display: inline-block;
        margin-left: 10rpx;
        color: #8b8b8b;

        .weather-icon {
          width: 40rpx;
          height: 40rpx;
          vertical-align: middle;
          display: inline-block;
        }
        .temper {
          font-size: 28rpx;
          margin-left: 8rpx;
        }
      }
    }
  }

  .back-btn {
    position: fixed;
    right: 0;
    bottom: 0;
    background: red;
    width: 80rpx;
    height: 80rpx;
    border-radius: 80rpx 0 0 0;
    .icon {
      position: absolute;
      /* top: 4px; */
      right: 12rpx;
      bottom: 12rpx;
    }
  }
}
</style>
