<!--
 * @Description  : 肥力图
 * @Autor        : Hemingway
 * @Date         : 2022-05-19 17:54:09
 * @LastEditors  : Yumanshan
 * @LastEditTime : 2022-05-26 05:35:55
 * @FilePath     : \blockchain-admin\src\views\dashboard\components\fertility-map\index.vue
-->
<template>
  <fertility-analyse ref="map" />
</template>

<script>
import FertilityAnalyse from '@/views/gis/soilFertilize/analyse'

export default {
  name: 'FertilityMap',

  components: {
    FertilityAnalyse
  },

  data() {
    return {}
  },

  computed: {},

  created() {},

  mounted() {
    setTimeout(() => {
      this.$refs.map.refreshMap()
    }, 2000)
  },

  methods: {}
}
</script>

<style lang="scss" scoped></style>
