<template>
  <div class="data-dict">
    <i-tab>
      <el-row :gutter="20">
        <el-col :span="12"
          ><left-col
            :dict-type-list="dictTypeList"
            @pop="onDialogVisible"
            @row-click="onRowClick"
            @row-delete="onRowDelete"
            ref="leftCol"
            class="col"
        /></el-col>
        <el-col :span="12"
          ><right-col
            :dictionary-id="dictionaryId"
            @pop="onDialogVisible"
            @row-delete="onRowDelete"
            ref="rightCol"
            class="col"
        /></el-col>
      </el-row>

      <pop-view
        ref="dialog"
        :is-type="isType"
        :pop-data="popData"
        :dict-type-list="dictTypeList"
        @submit="onSubmit"
      />
    </i-tab>
  </div>
</template>

<script>
import LeftCol from './components/LeftCol'
import RightCol from './components/RightCol'
import PopView from './components/PopView'
import {
  createDictionary,
  editDictionary,
  createDictionaryValue,
  editDictionaryValue,
  deleteDictionary,
  deleteDictionaryValue
} from '@/api/systemSettings'

export default {
  name: 'DataDictionary',

  components: { LeftCol, RightCol, PopView },

  data() {
    return {
      // left：筛选条件
      type: '', // 类型筛选
      searchKey: '', // 搜索

      leftList: [],
      totalLeftList: [],
      rightList: [],
      isType: '',
      popData: {}, // 传值给popView
      dictionaryId: ''
    }
  },

  computed: {
    dictType() {
      return this.$store.getters.systemDic.dictType
    },
    dictTypeList() {
      // 组装类型数组
      const dicType = this.dictType

      if (dicType && dicType.length !== 0) {
        const keys = Object.keys(dicType)
        const dics = []
        keys.forEach(key => {
          const dic = {}
          dic['value'] = key
          dic['label'] = dicType[key]
          dics.push(dic)
        })
        return dics
      }
      return []
    }
  },

  methods: {
    onRowClick(data) {
      this.dictionaryId = data.id // 保存选择的字典ID，之后创建字典值的时候用
    },

    // 新建/编辑
    onDialogVisible(row) {
      this.isType = row.isType
      this.popData = row
      this.$refs.dialog.visible = true
    },

    // 新建/编辑 提交
    async onSubmit(param) {
      if (param.isType === 'newDicValue' || param.isType === 'editDicValue') {
        // 字典值
        param['dictionaryId'] = this.dictionaryId
      }

      const apiMap = {
        createDictionary,
        editDictionary,
        createDictionaryValue,
        editDictionaryValue
      }

      try {
        this.$refs.dialog.loading = true
        await apiMap[param.funId](param)
        this.$refs.dialog.loading = false
        this.$refs.dialog.$refs.form.resetFields()

        if (param.isType === 'newDic' || param.isType === 'editDic') {
          // 字典
          this.$refs.leftCol.getTableData()
        } else if (
          // 字典项
          param.isType === 'newDicValue' ||
          param.isType === 'editDicValue'
        ) {
          this.$refs.rightCol.getTableData()
        }
      } catch (error) {
        console.log('新建/编辑 字典/字典项 error: ', error)
        this.$refs.dialog.loading = false
        this.$refs.dialog.resetForm()
      }

      this.$refs.dialog.visible = false
    },

    // 删除行
    async onRowDelete(params) {
      let api, param
      if (params.isType === 'DIC') {
        // 删除字典
        api = deleteDictionary
        param = { dictionaryId: params.id }
      } else {
        // 删除字典项
        api = deleteDictionaryValue
        param = { dictionaryItemId: params.id }
      }

      try {
        await api(param)
        if (params.isType === 'DIC') {
          this.$refs.leftCol.getTableData()
        } else {
          this.$refs.rightCol.getTableData()
        }
      } catch (error) {
        console.log('删除字典/字典项 error', error)
      }
    }
  }
}
</script>

<style scoped lang="scss">
.data-dict {
  height: 100%;
  background-color: #fff;
  position: relative;

  .col {
    height: calc(100vh - 204px);
    position: relative;
    padding-bottom: 56px;

    ::v-deep {
      .body {
        height: 100%;
        overflow-y: auto;
      }
      .body::-webkit-scrollbar {
        display: none;
      }
    }
  }

  ::v-deep .i-footer {
    left: 0;
    right: 0;
    bottom: -16px;
  }
}
</style>
