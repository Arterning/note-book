<!--
 * @Description  : 农事日历
 * @Autor        : Hemingway
 * @Date         : 2022-05-19 17:51:54
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-23 11:07:02
 * @FilePath     : \blockchain-admin\src\views\dashboard\components\farming-calendar\index.vue
-->
<template>
  <farming-map />
</template>

<script>
import FarmingMap from '@/views/gis-farm-calendar'

export default {
  name: 'FarmingCalendar',

  components: {
    FarmingMap
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
