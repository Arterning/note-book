<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-11 17:17:55
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-07 10:08:56
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\components\GoodsBaseInfoItem.vue
-->
<template>
  <div class="base-info">
    <i-section-header title="商品基本信息">
      <div slot="action" v-if="showEditBtn === true">
        <el-row v-if="isWrite === true">
          <el-button type="text" size="mini" @click="clickCancel"
            >取消</el-button
          >
          <el-button type="text" size="mini" @click="clickSave">保存</el-button>
        </el-row>
        <el-button
          slot="action"
          type="text"
          size="mini"
          v-else
          @click="clickWrite"
          >编辑</el-button
        >
      </div>
    </i-section-header>
    <slot></slot>
    <el-form
      label-suffix="："
      inline
      :model="data"
      :rules="rules"
      ref="formData"
    >
      <div>
        <el-form-item
          prop="commodityName"
          style="
            padding-right: 40px;
            margin-right: 0;
            margin-bottom: 10px;
            width: 33.33%;
          "
          label="商品名称"
        >
          <el-input
            v-model="data.commodityName"
            placeholder="请输入"
            :disabled="inputDisable"
            size="mini"
          ></el-input>
        </el-form-item>
        <el-form-item
          prop="brandId"
          style="
            padding-right: 40px;
            margin-right: 0;
            margin-bottom: 10px;
            width: 33.33%;
          "
          label="所属品牌"
        >
          <el-select
            v-model="data.brandId"
            size="mini"
            :disabled="inputDisable"
          >
            <el-option
              v-for="item in BrandList"
              :key="item.id"
              :label="item.brandName"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          prop="type"
          style="
            padding-right: 40px;
            margin-right: 0;
            margin-bottom: 10px;
            width: 33.33%;
          "
          label="所属品类"
        >
          <el-select
            v-model="data.type"
            placeholder="请选择品类"
            filterable
            clearable
            size="mini"
            :disabled="inputDisable"
          >
            <el-option
              v-for="item in typeArr"
              :label="item.dicValue"
              :value="item.dicCode"
              :key="item.id"
              class="option-style"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item
          prop="qualityGrade"
          style="
            padding-right: 40px;
            margin-right: 0;
            margin-bottom: 10px;
            width: 33.33%;
          "
          label="质量等级"
        >
          <el-input
            v-model="data.qualityGrade"
            placeholder="请输入"
            :disabled="inputDisable"
            size="mini"
          ></el-input>
        </el-form-item>
        <el-form-item
          style="
            padding-right: 40px;
            margin-right: 0;
            margin-bottom: 10px;
            width: 33.33%;
          "
          label="标准代号"
        >
          <el-input
            v-model="data.standard"
            placeholder="请输入"
            :disabled="inputDisable"
            size="mini"
          ></el-input>
        </el-form-item>
        <el-form-item
          prop="qualityGuaPeriod"
          style="
            padding-right: 40px;
            margin-right: 0;
            margin-bottom: 10px;
            width: 33.33%;
          "
          label="保质期"
        >
          <el-input
            v-model="data.qualityGuaPeriod"
            placeholder="请输入"
            :disabled="inputDisable"
            size="mini"
            style="width: 200px"
          >
            <template slot="append">天</template>
          </el-input>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
import { getBrandList, updateCommodity, getType } from '@/api/basedata-manage'

export default {
  name: 'GoodsBaseInfoItem',
  props: {
    showEditBtn: {
      type: Boolean,
      default: false
    },
    inputDisable: {
      type: Boolean,
      default: false
    },
    data: {
      type: Object,
      default: null,
      required: false
    }
  },
  data() {
    return {
      isWrite: false,
      BrandList: [], //获取从index页面存储到本地的品牌列表数据
      typeArr: [],
      rules: {
        commodityName: [
          { required: true, message: '请填写商品名称', trigger: 'blur' }
        ],
        brandId: [
          { required: true, message: '请选择所属品牌', trigger: 'change' }
        ],
        type: [
          { required: true, message: '请选择所属品类', trigger: 'change' }
        ],
        qualityGrade: [
          { required: true, message: '请填写质量等级', trigger: 'blur' }
        ],
        qualityGuaPeriod: [
          { required: true, message: '请填写保质期', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getBrandList()
    this.getType()
  },
  methods: {
    async getBrandList() {
      const { data } = await getBrandList()
      this.BrandList = data.map(res => {
        res.id = res.id.toString()
        return res
      })
    },

    // 获取品类
    async getType() {
      // const { data } = await getType({ dicType: 'brand_commodity_type' })
      const { data } = await getType({ dicType: 'brand_commodity_type' })

      this.typeArr = data
    },

    clickWrite() {
      this.isWrite = true
      this.$emit('update:inputDisable', false)
    },
    clickCancel() {
      this.isWrite = false
      this.$emit('update:inputDisable', true)
    },
    async clickSave() {
      this.$refs['formData'].validate(async valid => {
        if (valid) {
          try {
            console.log('this.data', this.data)
            const { code } = await updateCommodity(this.data)
            if (code === 0 || code === '0') {
              this.$notify({
                title: '成功',
                message: '修改成功！',
                type: 'success',
                duration: 2000
              })
              this.isWrite = false
              this.$emit('update:inputDisable', true)
            }
          } catch (error) {
            console.log('编辑保存失败', error)
          }
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
.base-info {
  ::v-deep .el-form-item--medium .el-form-item__label {
    font-size: 13px;
    color: #000;
    font-weight: lighter;
  }
  ::v-deep .el-form-item__error {
    top: 80%;
  }
}
</style>
