<!--
 * @Description  : 高德地图组件
 * @Autor        : Hemingway
 * @Date         : 2021-12-27 17:04:06
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-10-25 16:31:24
 * @FilePath     : \blockchain-platform\src\components\i-amap\index.vue
-->
<template>
  <div class="i-amap">
    <div class="i-amap-container">
      <el-amap :center="mapCenter" :zoom="zoom" :events="events">
        <el-amap-marker v-if="visible" :position="marker" :visible="visible" />
      </el-amap>
    </div>
    <search-box
      :search-option="searchOption"
      :on-search-result="onSearchResult"
      class="search-box"
      placeholder="请输入地址"
    />

    <div class="address-wrapper">
      <div class="address">
        <span class="prefix">当前选择地址为:</span>
        {{ addressObj.fullAddress || '' }}
      </div>
      <el-button @click="onConfirm" type="primary">确 定</el-button>
    </div>
  </div>
</template>

<script>
import { lazyAMapApiLoaderInstance } from 'vue-amap'
import SearchBox from './search-box'

export default {
  name: 'IAmap',

  components: { SearchBox },

  props: {
    lng: {
      type: String,
      default: ''
    },
    lat: {
      type: String,
      default: ''
    }
  },

  data() {
    return {
      addressObj: {},
      marker: [],
      visible: false,
      mapCenter: ['112.9349', '28.22495'], // 默认地图中心
      zoom: 12,

      searchOption: {
        pageSize: 1,
        pageIndex: 1,
        city: '',
        citylimit: false
      },

      events: {
        init: this.init,
        click: e => {
          this.mapCenter = [e.lnglat.lng.toFixed(6), e.lnglat.lat.toFixed(6)]
        }
      },
      inited: false // 地图是否已初始化
    }
  },

  watch: {
    mapCenter() {
      this.handleAddress()
    }
  },

  methods: {
    /**
     * @description: 设置地图中心
     * @author: Hemingway
     */
    init() {
      this.inited = this.inited || true
      this.mapCenter = [this.lng, this.lat]
    },

    /**
     * @description: 地图 PlaceSearch
     * @param {Array} pois POI数据
     * @author: Hemingway
     */
    onSearchResult(pois) {
      const { lng, lat } = pois[0]
      this.mapCenter = [lng, lat]
    },

    /**
     * @description: 根据坐标，处理地图显示及位置数据
     * @param {Array} center [lng, lat]
     * @author: Hemingway
     */
    handleAddress() {
      this.getAddress(this.mapCenter)
        .then(res => {
          this.marker = this.mapCenter
          this.visible = true
          this.zoom = 15
          this.addressObj = res
        })
        .catch(error => {
          console.log('以经纬度获取地址信息 error = ', error)
        })
    },

    /**
     * @description: 使用经纬度查询地址详情
     * @param {Array} center [lng, lat]
     * @return {Promise}
     * @author: Hemingway
     */
    getAddress(center) {
      let geocoder = null

      return new Promise((resolve, reject) => {
        // 定义resolveMethod函数逻辑
        const resolveMethod = () => {
          let res = null

          geocoder.getAddress(
            new AMap.LngLat(center[0], center[1]),
            (status, result) => {
              if (status === 'complete' && result.info === 'OK') {
                const { addressComponent, formattedAddress } = result.regeocode
                const {
                  province,
                  city,
                  district,
                  township,
                  street,
                  streetNumber,
                  adcode
                } = addressComponent

                // 计算省市区编码
                const [provinceCode, cityCode, districtCode] = Array(3)
                  .fill(adcode)
                  .map((code, i) => code.slice(0, 2 * (i + 1)).padEnd(6, 0))

                // 计算详细地址
                const formattedAddressDetail = this.getAddressDetail(
                  formattedAddress,
                  [province, city, district]
                )

                res = {
                  lng: center[0].toString(),
                  lat: center[1].toString(),
                  province: province,
                  city: city || province, // 高德地图直辖市获取不到city，取province
                  district: district,
                  provinceCode,
                  cityCode,
                  districtCode,
                  addressDetail: township + street + streetNumber,
                  formattedAddress,
                  formattedAddressDetail,
                  fullAddress: `${formattedAddress}(${
                    township + street + streetNumber
                  })`
                }

                resolve(res)
              } else {
                reject(result)
              }
            }
          )
        }

        if (!geocoder) {
          lazyAMapApiLoaderInstance.load().then(() => {
            geocoder = new AMap.Geocoder({})
            resolveMethod()
          })
        } else {
          resolveMethod()
        }
      })
    },

    /**
     * @description: 省市区名称截取，获取详细地址
     * @param {String} raw 源字符串
     * @param {Arrays} words [province, city, district]
     * @return {String}
     * @author: Hemingway
     */
    getAddressDetail(raw, words) {
      let res = raw
      const replaceAll = (FindText, RepText) => {
        const regExp = new RegExp(FindText, 'g')
        return RepText.replace(regExp, '')
      }

      words.forEach(w => {
        res = replaceAll(w, res)
      })

      return res
    },

    /**
     * @description: 提交事件
     * @author: Hemingway
     */
    onConfirm() {
      if (!this.addressObj.province) {
        this.$message.error('请输入详细地址')
        return
      }

      this.$emit('returnAddress', this.addressObj)
    }
  }
}
</script>

<style scoped lang="scss">
.i-amap {
  position: relative;
  width: 90vw;
  height: calc(90vh - 75px);
  display: flex;
  flex-direction: column;

  &-container {
    flex: 1;
  }

  .address-wrapper {
    height: 60px;
    padding: 0 16px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .prefix {
      font-weight: bold;
    }
  }

  .search-box {
    position: absolute;
    top: 25px;
    left: 20px;
  }
}
</style>
