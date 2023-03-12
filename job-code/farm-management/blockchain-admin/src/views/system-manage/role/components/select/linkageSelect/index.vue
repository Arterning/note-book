<template>
  <div id="linkageSelect">
    <div class="province">
      <div
        v-for="(item, index) in provinceList"
        :key="index"
        :class="activeProve && activeProve.code === item.code ? 'active' : ''"
        @click="provinceClick(item, index)"
      >
        {{ item.name }}
      </div>
    </div>
    <div class="city">
      <div
        v-for="(item, index) in cityList"
        :key="index"
        :class="activeCity && activeCity.code === item.code ? 'active' : ''"
        @click="cityClick(item, index)"
      >
        {{ item.name }}
      </div>
    </div>
    <div class="area">
      <div
        v-for="(item, index) in areaList"
        :key="index"
        :class="activeArea && activeArea.code === item.code ? 'active' : ''"
        @click="areaClick(item, index)"
      >
        {{ item.name }}
      </div>
    </div>
  </div>
</template>

<script>
import areaList from './area'
export default {
  name: 'linkageSelect',
  props: {
    value: {
      type: Boolean,
      default: true
    },
    codeObj: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      provinceList: [],
      cityList: [],
      areaList: [],
      activeProve: null,
      activeCity: null,
      activeArea: null
    }
  },
  components: {},
  beforeCreate() {
    // console.log(areaList)
  },
  created() {},
  beforeMount() {
    if (this.codeObj.dataPermissionProvince) {
      this.findNameByCode()
    }
    this.provinceList = areaList.reduce(
      (prev, cur) => {
        prev.push({
          name: cur.name,
          code: cur.code
        })
        return prev
      },
      [
        {
          name: '全国',
          code: 'all'
        }
      ]
    )
    // console.log(this.provinceList)
    // console.log(areaList)
  },
  mounted() {},
  methods: {
    findNameByCode() {
      this.activeProve = areaList.find(
        item => item.code === this.codeObj.dataPermissionProvince
      )
      this.cityList = this.activeProve.children
      this.activeCity = this.cityList.find(
        item => item.code === this.codeObj.dataPermissionCity
      )
      this.areaList = this.activeCity.children || []
      this.activeArea = this.areaList.find(
        item => item.code === this.codeObj.dataPermissionArea
      )
    },
    provinceClick(item) {
      this.clearData()
      this.activeProve = item
      let activeProve = areaList.find(prove => item.code === prove.code)
      if (activeProve && activeProve.children && activeProve.children.length) {
        this.cityList = activeProve.children
      } else {
        this.changeOver()
      }
    },
    cityClick(item) {
      this.activeCity = item
      let activeCity = this.cityList.find(city => item.code === city.code)
      if (activeCity.children) {
        this.areaList = activeCity.children
      } else {
        this.changeOver()
      }
    },
    areaClick(item) {
      this.activeArea = item
      this.changeOver()
    },
    clearData() {
      this.activeCity = null
      this.cityList = []
      this.activeArea = null
      this.areaList = []
    },
    changeOver() {
      let addressInfo = {
        province: this.activeProve,
        city: this.activeCity,
        area: this.activeArea
      }
      console.log('addressInfo', addressInfo)
      this.$emit('changeOver', addressInfo)
      this.$emit('input', false)
    }
  },
  watch: {},
  computed: {},
  beforeDestroy() {},
  destroyed() {}
}
</script>

<style scoped lang="scss" type="text/less">
#linkageSelect {
  display: flex;
  align-items: center;
  > div {
    flex: 1;
    height: 300px;
    overflow: auto;
    > div {
      padding: 6px 10px;
      &:hover,
      &.active {
        background: #ccc;
      }
    }
  }
}
</style>
