<!--
 * @Description  : 施肥
 * @Autor        : Hemingway
 * @Date         : 2022-05-19 17:51:54
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-23 11:26:50
 * @FilePath     : \blockchain-admin\src\views\dashboard\components\fertilization\index.vue
-->
<template>
  <fertilization-map />
</template>

<script>
import FertilizationMap from '@/views/gis/soilFertilize/fertilizer'

export default {
  name: 'Fertilization',

  components: {
    FertilizationMap
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
