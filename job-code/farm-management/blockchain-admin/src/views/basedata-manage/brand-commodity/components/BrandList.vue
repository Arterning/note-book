<!--
 * @Description  : 品牌列表弹窗
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-04 10:06:54
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-07 09:57:57
 * @FilePath     : \blockchain-admin\src\views\basedata-manage\brand-commodity\components\BrandList.vue
-->
<template>
  <el-drawer
    :visible="visible"
    :direction="direction"
    :show-close="false"
    size="375px"
  >
    <template slot="title">
      <p class="title-detail">品牌列表</p>
    </template>

    <!-- container部分 -->
    <el-input
      v-model="brandName"
      placeholder="请输入品牌名称"
      suffix-icon="el-icon-search"
      @input="onInput"
    >
    </el-input>

    <div class="brand-item" v-for="(row, index) in list" :key="row.id">
      <div>
        <el-input
          v-model="saveIpt"
          size="mini"
          v-if="editStateObj.index == index"
          placeholder="请输入品牌名称"
        ></el-input>
        <p v-else>{{ row.brandName }}</p>
        <span>共{{ row.commodityCount || 0 }}个商品</span>
      </div>
      <div>
        <div class="save-cancel" v-if="editStateObj.index == index">
          <el-button type="text" @click="onSave(row.id)">保存</el-button>
          <el-button
            type="text"
            style="color: red; margin-right: 9px"
            @click="onCancel(row.id)"
            >取消</el-button
          >
        </div>
        <el-button type="text" v-else @click="onEditstate(index)"
          >编辑</el-button
        >
      </div>
    </div>
    <!-- container部分 -->

    <div class="dialogFooter">
      <el-button plain size="small" @click="btnClose">关 闭</el-button>
      <el-button type="primary" size="small" @click="onAdd()">新 增</el-button>
    </div>
  </el-drawer>
</template>

<script>
import { getBrandDialog, updateBrand, addBrand } from '@/api/basedata-manage'

export default {
  name: 'BrandList',
  props: { visible: { type: Boolean, default: false } },
  data() {
    return {
      direction: 'rtl',
      editStateObj: { index: null },
      brandName: null,
      saveIpt: '',
      list: []
    }
  },
  created() {
    this.getBrandList()
  },
  methods: {
    async getBrandList() {
      try {
        const { list } = await getBrandDialog({ brandNameLike: this.brandName })
        this.list = list
      } catch (error) {
        console.log('品牌列表 err:', error)
      }
    },

    // 输入框联想建议
    onInput() {
      if (this.brandName) {
        this.getBrandList()
      } else {
        this.getBrandList()
      }
    },

    // 点击保存按钮(编辑)
    async onSave(id) {
      console.log('id', id)
      if (this.saveIpt) {
        try {
          if (id) {
            // 编辑保存
            const { code } = await updateBrand({
              brandName: this.saveIpt,
              id: id
            })
            if (code === 0 || code === '0') {
              this.getBrandList()
              this.editStateObj.index = null
            }
          } else {
            // 新增保存
            const { code } = await addBrand({
              brandName: this.saveIpt
            })
            if (code === 0 || code === '0') {
              this.getBrandList()
              this.editStateObj.index = null
            }
          }
        } catch (error) {
          console.log('编辑新增保存：', error)
        }
      } else {
        this.$message.warning('品牌名称不能为空！')
      }
    },
    onCancel(id) {
      if (!id) {
        this.list.pop()
      }
      this.editStateObj.index = null
      // if (this.saveIpt.length === 0) {
      //   this.list.pop()
      // }
      // this.editStateObj.index = null
    },
    // 点击编辑按钮
    onEditstate(index) {
      this.editStateObj.index = index
      this.saveIpt = this.list[index].brandName
    },

    // 点击底部关闭按钮
    btnClose() {
      this.$emit('refresh', false)
    },

    // 点击新增按钮事件
    onAdd() {
      if (this.editStateObj.index) return
      this.editStateObj.index = this.list.length
      this.saveIpt = ''
      this.list.push({ brandName: '', commodityCount: '0' })
    }
  }
}
</script>

<style scoped lang="scss">
.el-drawer__wrapper {
  height: 667px;
  margin-top: 100px;
}

::v-deep .el-drawer__header {
  border-bottom: 1px solid #ccc;
  height: 9% !important;

  .title-detail {
    margin-top: 10px;
    font-size: 16px;
    color: #000 !important;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif !important;
  }
}

::v-deep .el-drawer__open .el-drawer.rtl {
  border-radius: 1%;
}

::v-deep .el-drawer__body {
  margin: -9% 0 16% 0;
  padding: 14px;
  overflow: auto;
  &::-webkit-scrollbar {
    display: none;
  }

  .el-input {
    margin-bottom: 20px;
    padding-right: 6px;
  }
}

.dialogFooter {
  height: 9%;
  width: 100%;
  position: absolute;
  bottom: 0;
  left: 0;
  text-align: center;
  line-height: 56px;
  border-top: 1px solid #ccc;
}
</style>

<style scoped lang="scss">
.brand-item {
  height: 80px;
  border-bottom: 1px solid rgb(223, 215, 215);
  padding: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  div:nth-child(2) {
    .save-cancel {
      display: flex;
      flex-direction: column;
      margin-right: -10px;
    }
  }

  span {
    font-size: 14px;
    color: #a7a8a5;
  }
  p:nth-child(1) {
    font-size: 16px;
  }

  p:nth-child(2) {
    color: #27d581;
  }
}
</style>
