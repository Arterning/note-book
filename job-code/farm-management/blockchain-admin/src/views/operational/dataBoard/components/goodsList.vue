<template>
  <div class="">
    <!-- 表格部分  -->
    <i-table border :data="tableData" :loading="loading" :max-height="450">
      <el-table-column
        label="日期"
        prop="scDate"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="被扫码数"
        prop="beSweptCodeCount"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="扫码次数"
        prop="sweptCodeCount"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column
        label="所属商品"
        prop="commodityName"
        show-overflow-tooltip
      ></el-table-column>
    </i-table>
    <!-- footer区域  -->
    <!-- <i-footer :title="title" v-if="tableData && tableData.length > 0">
      <i-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="pageNum"
        :page-size="pageSize"
        :total="total"
      />
    </i-footer> -->
  </div>
</template>
<script>
import { getScDataList } from '@/api/operational-dataBoard'
import tableMixin from '@/views/mixins/tableMixin'
export default {
  name: '',
  mixins: [tableMixin],
  components: {},
  props: {
    timesData: {
      type: Array
    },
    plantYear: {
      type: String
    },
    varietyType: {
      type: String
    },
    commodityId: {
      type: String
    },
    commodityName: {
      type: String
    }
  },
  data() {
    return {
      tableData: [],
      loading: false,
      title: '表',
      pageNum: 1,
      pageSize: 10,
      total: 0
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {
    this.init()
  },
  methods: {
    async init() {
      let params = JSON.parse(
        JSON.stringify({
          plantYear: this.plantYear !== '' ? this.plantYear : undefined,
          varietyType: this.varietyType !== '' ? this.varietyType : undefined,
          commodityId: this.commodityId !== '' ? this.commodityId : undefined,
          startDate: this.timesData[0] !== '' ? this.timesData[0] : undefined,
          endDate: this.timesData[1] !== '' ? this.timesData[1] : undefined,
          commodityName:
            this.commodityName !== '' ? this.commodityName : undefined,
          pageNum: this.pageNum,
          pageSize: this.pageSize
        })
      )
      this.tableData = []
      this.total = 0
      try {
        let { data } = await getScDataList(params)
        if (!data?.list || data?.list?.length === 0) {
          this.$emit('tableDataChange', true)
        } else {
          this.$emit('tableDataChange', false)
        }
        console.log('列表查询-getScDataList', data)
        if (data.list && Array.isArray(data.list) && data.list.length > 0) {
          let _beNumber = []
          let _number = []
          data.list.map(el => {
            _beNumber.push(el.beSweptCodeCount)
            _number.push(el.sweptCodeCount)
          })
          this.tableData = data.list
          this.total = data.total
          if (data.list.length < 1) return
          this.tableData.push({
            scDate: `总计`,
            beSweptCodeCount: eval(_beNumber.join('+')),
            sweptCodeCount: eval(_number.join('+')),
            commodityName: data.list[0].commodityName
          })
        }
      } catch (error) {
        this.$message.error
      }
    },

    handleSizeChange(val) {
      console.log('val', val)
      this.pageSize = val
      this.init()
    },
    handleCurrentChange(val) {
      console.log('val', val)
      this.pageNum = val
      this.init()
    }
  }
}
</script>
<style scoped lang="scss"></style>
