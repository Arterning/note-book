<!--
 * @Description  : 
 * @Autor        : Your Name
 * @Date         : 2022-06-14 11:26:28
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-25 20:33:34
 * @FilePath     : \blockchain-h5\src\pages\module\machining.vue
-->
<template>
  <view>
    <!-- 加工信息 -->
    <view class="card-info common-box" :class="themeLinear">
      <view class="card-title" :class="themeColor">
        <text>加工信息</text>
        <view class="certified" :class="themeBackground">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 61.26 64"
            class="status-img"
          >
            <g id="图层_2" data-name="图层 2">
              <g id="图层_1-2" data-name="图层 1">
                <path
                  id="认证"
                  d="M61.23,5.66a2.54,2.54,0,0,0-.89-1.58c-.1-.1-1.68-1.29-4.15-3.17A2.54,2.54,0,0,0,53.82.62C40.27,6.45,32.46.82,32.06.52a2.25,2.25,0,0,0-2.87,0c-.09.1-8,5.93-21.75.1A2.54,2.54,0,0,0,5.07.91C2.6,2.79,1,4,.92,4.08A2.49,2.49,0,0,0,0,5.66,2.33,2.33,0,0,0,.52,7.44c.1.1,5.14,6.63,2.77,16.51C1.9,29.79.42,35.82,1.71,41.46,3.19,48.18,8.23,52.93,16.93,56H17A25.85,25.85,0,0,1,28.7,63.11a2.43,2.43,0,0,0,3.76,0A25,25,0,0,1,44.13,56h.09C53,52.93,58,48.18,59.45,41.46c1.29-5.74-.2-11.77-1.58-17.51A20.31,20.31,0,0,1,60.64,7.44,2.09,2.09,0,0,0,61.23,5.66Zm-43,6.82c1.38,1.88,2.77,4.06,4.25,6.43-1.09.59-2.47,1.39-4.25,2.47-1.29-2.27-2.57-4.45-3.86-6.52ZM15.15,41.26A3.83,3.83,0,0,0,16,38.69V27.51H13.28V23h7.51V36.81l3.07-2.47c.09,1,.39,2.57.79,4.74-2.18,1.58-4.65,3.46-7.42,5.84Zm32.54,4.05H23.06V40.66h2.87V23.26h4.65V40.57h3.76V19.41H24.15V14.86H47.29v4.55H39.08v7.81H46.5v4.64H39.08v8.8h8.61Z"
                />
              </g>
            </g>
          </svg>
          已认证
        </view>
      </view>
      <view class="card-content-box">
        <view class="content-info">
          <view class="content-info-row">
            <text space="emsp" decode="true" class="label" :class="themeColor"
              >加工企业: {{ '' }}</text
            >
            <text class="value">{{ processInfoData.riceFactoryName }}</text>
          </view>
          <view class="content-info-row">
            <text space="emsp" decode="true" class="label" :class="themeColor"
              >加工地址: {{ '' }}</text
            >
            <text class="value">{{ processInfoData.riceFactoryLocation }}</text>
          </view>
        </view>
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
    </view>
  </view>
</template>

<script>
import { preViewImg, getImageUrl } from '../../utils/tool'
import Steps from '@/components/steps'
import { getProcessInfo } from '@/api/myRice'
export default {
  name: 'Machining',
  components: { Steps },
  props: ['menuTree'],
  filters: {
    /**
     * @description: 根据UI设计稿处理日期展示
     * @param {String} date
     * @param {String} flagStr
     * @return {String}
     * @author: Hemingway
     */
    dateFormatter(date, flagStr) {
      if (!date) return
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
      processInfoData: {},
      processData: [],
      excludeArr: ['加工企业', '加工地址']
    }
  },

  computed: {
    // 溯源码编号
    qrcode() {
      return this.$store.getters.code
    },
    isStaticData() {
      return this.$store.getters.isStaticData
    },
    packingInfo() {
      return this.$store.getters.packingInfo
    },
    themeColor() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'theme1-color'
      } else if (colorId === '3') {
        return 'theme2-color'
      } else {
        return ''
      }
    },
    themeBackground() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'theme1-background'
      } else if (colorId === '3') {
        return 'theme2-background'
      } else {
        return ''
      }
    },
    themeLinear() {
      let colorId = this.$store.getters.colorId.toString()
      if (colorId === '2') {
        return 'theme1-linear'
      } else if (colorId === '3') {
        return 'theme2-linear'
      } else {
        return ''
      }
    }
  },

  watch: {
    // 深度监听，及时更新页面
    menuTree: {
      handler() {
        this.getSteps()
      },
      deep: true
    }
  },

  async created() {
    if (this.isStaticData) {
      let { machining } = require('../../static/json/static.json')
      this.processInfoData = machining || {}
      this.processData = machining.processData || []
      this.getSteps()
      return
    }
    await this.processInfo()
    this.getSteps()
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
    async processInfo() {
      const { data } = await getProcessInfo({
        qrCode: this.qrcode,
        machPackingId: this.packingInfo && this.packingInfo.machPackingId,
        packingBatchId: this.packingInfo && this.packingInfo.packingBatchId
      })
      if (!data) return
      this.processInfoData = data || {}
      this.processData = data.processData || []
    },
    /**
     * @description: 根据配置信息获取steps数据
     * @author: qjj
     */
    getSteps() {
      let newArr = []
      if (this.isStaticData) {
        let fitArr = this.menuTree.subTree.filter(item => {
          return (
            !this.excludeArr.includes(item.menuName) && item.isChecked === 'Y'
          )
        })
        newArr = fitArr.map(res => {
          return {
            date: '2022-01-02',
            title: res.menuName,
            body: {
              imgs: ['../../static/img/template/machning/包装.jpg']
            }
          }
        })
      } else {
        this.processData.forEach(res => {
          let currentConfig = this.menuTree.subTree.find(item => {
            return item.menuId.toString() === res.menuId.toString()
          })
          // 如果主配置项不显示，直接返回
          if (
            typeof currentConfig === 'undefined' ||
            currentConfig?.isChecked !== 'Y'
          )
            return
          newArr.push({
            date:
              currentConfig.menuId === '79'
                ? '具体见外包装'
                : res.operationDate && res.operationDate.split(' ')[0],
            title: res.operationType,
            body: {
              imgs:
                (res.images &&
                  res.images.split(',').map(id => getImageUrl(id))) ||
                []
            }
          })
        })
      }

      this.steps = newArr
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../static/style/template1.scss';
.card-content-box {
  .section {
    margin-top: 48rpx;
  }

  .content-info {
    .content-info-row {
      margin-bottom: 25rpx;
    }

    .label {
      color: #eabc02;
    }
  }

  // 时间线
  .process-content {
    background-color: #fff;
    margin-top: 50rpx;
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
}
</style>
