<!--
 * @Description  : 田间监测
 * @Autor        : Hemingway
 * @Date         : 2022-05-19 17:51:54
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-05-23 19:11:27
 * @FilePath     : \blockchain-admin\src\views\dashboard\components\field-equipment\index.vue
-->
<template>
  <equipment-map />
</template>

<script>
import EquipmentMap from '@/views/gis/field-equipment/components/EquipmentMap.vue'

export default {
  name: 'FieldEquipment',

  components: {
    EquipmentMap
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
