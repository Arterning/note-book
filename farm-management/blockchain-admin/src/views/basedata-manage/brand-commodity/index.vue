<!--
 * @Description  :品牌/商品管理
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-01 14:13:03
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-06 19:13:02
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\index.vue
-->
<template>
  <div class="brand-commodity">
    <i-tab>
      <i-action @on-query="getTableData" @on-reset="onReset" show-all>
        <i-action-item label="品牌名称">
          <el-select
            v-model="searchFields.brandId"
            placeholder="请选择品牌名称"
            filterable
            clearable
          >
            <el-option
              v-for="item in brandList"
              :key="item.id"
              :label="item.brandName"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </i-action-item>
        <i-action-item label="商品名称">
          <el-select
            v-model="searchFields.commodityId"
            placeholder="请选择商品名称"
            filterable
            clearable
          >
            <el-option
              v-for="item in goodsArr"
              :label="item.commodityName"
              :value="item.id"
              :key="item.id"
              class="option-style"
            ></el-option>
          </el-select>
        </i-action-item>

        <i-action-item label="品类名称">
          <el-select
            v-model="searchFields.type"
            placeholder="请选择品类名称"
            filterable
            clearable
          >
            <el-option
              v-for="item in typeArr"
              :label="item.dicValue"
              :value="item.dicCode"
              :key="item.id"
              class="option-style"
            ></el-option>
          </el-select>
        </i-action-item>

        <template #action>
          <el-button type="primary" @click="onAdd" v-btn-permission:1705
            >添加商品</el-button
          >
          <el-button @click="checkList">品牌列表</el-button>
        </template>
      </i-action>

      <brand-list
        v-if="drawerVisible"
        :visible.sync="drawerVisible"
        @refresh="refresh"
      />

      <!-- 表格部分 -->
      <i-table border :data="tableData" :loading="loading" :max-height="450">
        <el-table-column
          label="商品名称"
          prop="commodityName"
        ></el-table-column>
        <el-table-column label="所属品牌" prop="brandName"></el-table-column>
        <el-table-column label="所属品类" prop="type">
          <template slot-scope="scope">
            <span>{{ scope.row.type | typeFilter(typeArr) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="规格" prop="netContent"></el-table-column>
        <el-table-column label="质量等级" prop="qualityGrade"></el-table-column>
        <el-table-column label="标准代号" prop="standard"></el-table-column>
        <el-table-column
          label="保质期（天）"
          prop="qualityGuaPeriod"
        ></el-table-column>
        <el-table-column label="创建时间" prop="createDate"></el-table-column>
        <el-table-column label="上传人" prop="creator"></el-table-column>
        <el-table-column label="操作" width="90px">
          <template slot-scope="{ row }">
            <el-button type="text" @click="clickDetail(row.brandCommodityId)"
              >商品详情</el-button
            >
          </template>
        </el-table-column>
      </i-table>

      <!-- footer区域 -->
      <i-footer v-if="tableData && tableData.length > 0">
        <i-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="current"
          :page-size="size"
          :total="total"
        />
      </i-footer>
    </i-tab>
  </div>
</template>

<script>
import tableMixin from '@/views/mixins/tableMixin'
import BrandList from './components/BrandList.vue'
import {
  commodityList,
  getBrandList,
  getType,
  getGoods
} from '@/api/basedata-manage'

export default {
  name: 'brand-commodity',
  mixins: [tableMixin],
  components: { BrandList },
  data() {
    return {
      drawerVisible: false,
      searchFields: {
        commodityId: '',
        type: '',
        brandId: ''
      },
      brandList: [],
      goodsArr: [],
      typeArr: [],
      loading: false,
      tableData: [],
      brand_commodity_type: ''
    }
  },
  filters: {
    typeFilter(val, typeArr) {
      let result = typeArr.filter(res => {
        return val === res.dicCode
      })
      if (result.length === 0) {
        return ''
      }
      return result[0].dicValue
    }
  },
  created() {
    this.getBrandList()
    this.getGoods()
    this.getType()
    this.getTableData()
  },
  methods: {
    refresh() {
      this.drawerVisible = false
      this.getTableData()
    },
    // 获取品牌
    async getBrandList() {
      const { data } = await getBrandList()
      localStorage.setItem('BrandList', JSON.stringify(data))
      this.brandList = data
    },
    //获取商品
    async getGoods() {
      const { data } = await getGoods()
      this.goodsArr = data
    },
    // 获取品类
    async getType() {
      // const { data } = await getType({ dicType: 'brand_commodity_type' })
      const { data } = await getType({ dicType: 'brand_commodity_type' })
      this.typeArr = data
    },

    onAdd() {
      this.$router.push({
        name: 'addcommodity'
      })
    },
    checkList() {
      this.drawerVisible = true
    },

    async getTableDataImpl(payload) {
      const params = {
        ...payload,
        commodityId: this.searchFields.commodityId,
        type: this.searchFields.type,
        brandId: this.searchFields.brandId
      }
      const data = await commodityList(params)
      return { tableData: data.list, total: data.total }
    },

    onResetImpl() {
      this.searchFields.brandId = ''
      this.searchFields.commodityId = ''
      this.searchFields.type = ''
      this.getTableData()
    },

    clickDetail(brandCommodityId) {
      this.$router.push({
        name: 'details',
        params: { id: brandCommodityId }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.brand-commodity {
  height: 100%;
  background-color: #fff;
  position: relative;
}
</style>
