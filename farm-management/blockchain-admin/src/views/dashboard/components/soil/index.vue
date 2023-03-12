<!--
 * @Description  : 土壤
 * @Autor        : Hemingway
 * @Date         : 2022-05-19 17:54:09
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-23 11:33:32
 * @FilePath     : \blockchain-admin\src\views\dashboard\components\soil\index.vue
-->
<template>
  <soil-map />
</template>

<script>
import SoilMap from '@/views/gis/soilFertilize/testing.vue'

export default {
  name: 'Soil',

  components: {
    SoilMap
  },

  data() {
    return {}
  },

  computed: {},

  created() {},

  mounted() {},

  methods: {}
}
</script>

<style lang="scss" scoped></style>
