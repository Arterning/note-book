<!--
 * @Description  : 
 * @Autor        : Hemingway
 * @Date         : 2021-07-21 16:54:34
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-07-22 11:42:26
 * @FilePath     : \blockchain-minicode\src\pages\partnership\city.vue
-->
<template>
  <view>
    <city-select
      ref="citys"
      :format-name="formatName"
      :active-city="activeCity"
      :hot-city="hotCity"
      :obtain-citys="obtainCitys"
      :is-search="true"
      @cityClick="cityClick"
    ></city-select>
  </view>
</template>

<script>
import citys from './citys.js'
console.log(citys.length)
import citySelect from '@/components/city-select/city-select.vue'
export default {
  components: {
    citySelect
  },
  data() {
    return {
      //需要构建索引参数的名称（注意：传递的对象里面必须要有这个名称的参数）
      formatName: 'title',
      //当前城市
      activeCity: {},
      //热门城市
      hotCity: [
        {
          cityCode: '430100',
          cityName: '长沙市'
        },
        {
          cityCode: '430200',
          cityName: '株洲市'
        },
        {
          cityCode: '430300',
          cityName: '湘潭市'
        },
        {
          cityCode: '431300',
          cityName: '娄底市'
        },

        {
          cityCode: '460100',
          cityName: '海口市'
        },
        {
          cityCode: '430900',
          cityName: '益阳市'
        }
      ],
      //显示的城市数据
      obtainCitys: []
    }
  },
  onLoad() {
    //动态更新数据
    //修改需要构建索引参数的名称
    this.formatName = 'cityName'
    //修改当前城市

    //修改构建索引数据
    this.obtainCitys = citys
  },
  methods: {
    cityClick(item) {
      uni.navigateTo({
        url: `/subPages/partnership/List?cityName=${item.cityName}`
      })
    }

    // getAuthorizeInfo() {
    //   const that = this
    //   uni.authorize({
    //     scope: 'scope.userLocation',
    //     success() {
    //       // 允许授权
    //       that.getLocationInfo()
    //     },
    //     fail() {
    //       // 拒绝授权
    //       that.openConfirm()
    //       console.log('你拒绝了授权，无法获得周边信息')
    //     }
    //   })
    // },
    // // 获取地理位置
    // getLocationInfo() {
    //   const self = this
    //   uni.getLocation({
    //     geocode: true,
    //     type: 'wgs84',
    //     success(res) {
    //       uni.request({
    //         header: {
    //           'Content-Type': 'application/text'
    //         },
    //         //注意:这里的key值需要高德地图的 web服务生成的key  只有web服务才有逆地理编码
    //         url:
    //           'https://restapi.amap.com/v3/geocode/regeo?output=JSON&location=' +
    //           res.longitude +
    //           ',' +
    //           res.latitude +
    //           '&key=280802ed0116fef931dbcf5e7e9278d7&radius=1000&extensions=all',
    //         success(re) {
    //           self.activeCity = {
    //             cityName: re.data.regeocode.addressComponent.city,
    //             cityCode: self.obtainCitys.find(
    //               el =>
    //                 re.data.regeocode.addressComponent.city.indexOf(
    //                   el.cityName
    //                 ) !== -1
    //             ).cityCode
    //           }

    //           //   if (re.statusCode === 200) {
    //           //     that.citydata = re.data.regeocode.addressComponent.city
    //           //     console.log('获取中文街道地理位置成功', that.citydata)
    //           //   } else {
    //           //     console.log('获取信息失败，请重试！')
    //           //   }
    //         }
    //       })
    //     }
    //   })
    // },

    // // 再次获取授权
    // // 当用户第一次拒绝后再次请求授权
    // openConfirm() {
    //   uni.showModal({
    //     title: '请求授权当前位置',
    //     content: '需要获取您的地理位置，请确认授权',
    //     success: res => {
    //       if (res.confirm) {
    //         uni.openSetting() // 打开地图权限设置
    //       } else if (res.cancel) {
    //         uni.showToast({
    //           title: '你拒绝了授权，无法获得周边信息',
    //           icon: 'none',
    //           duration: 1000
    //         })
    //       }
    //     }
    //   })
    // }
  }
}
</script>
