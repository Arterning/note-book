<!--
 * @Description  : 数据模板配置
 * @Autor        : SunXiuqiong
 * @Date         : 2022-05-31 20:34:12
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-21 09:14:28
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\template\index.vue
-->
<template>
  <div class="template">
    <i-tab>
      <div class="wrapper">
        <select-data
          :readonly="selectDataReadonly"
          :getPrivewInstance="getPrivewInstance"
          ref="select-data"
          @isEditEvent="e => (isEdit = e)"
        ></select-data>
        <preview @hook:mounted="onPreviewMounted" ref="preview"></preview>
      </div>
      <!-- 页脚区域 -->
      <i-footer v-btn-permission:1711>
        <div class="footer-btn" v-if="selectDataReadonly">
          <el-button @click="selectDataReadonly = false" :disabled="isEdit"
            >修改</el-button
          >
        </div>
        <div class="footer-btn" v-else>
          <el-button type="primary" @click="btnOK">确定</el-button>
        </div>
      </i-footer>
    </i-tab>
  </div>
</template>

<script>
import Preview from './components/Preview.vue'
import SelectData from './components/SelectData.vue'
import { modifyTemplate } from '@/api/traceability-manage'

export default {
  name: 'traceability-manage--template',
  components: {
    Preview,
    SelectData
  },
  data() {
    return {
      selectDataReadonly: true,
      isEdit: false
    }
  },

  methods: {
    onPreviewMounted() {
      this.$nextTick(() => {
        this.$refs['select-data'].onPreview()
      })
    },
    // 给select-data组件提供preview组件的实例进行访问
    getPrivewInstance() {
      return this.$refs.preview
    },
    // 递归载入项
    RecursiveLoadTreeItems(menus, treeItems = []) {
      menus.forEach(menu => {
        let tempObj =
          'menuId' in menu && menu.menuId !== menu.moduleId
            ? { menuId: menu.menuId }
            : {}
        tempObj.moduleId = menu.moduleId
        //当isChecked||isDefault 等于Y的时候，均视为选中
        tempObj.isChecked =
          'isChecked' in menu
            ? menu.isChecked
            : menu.isDefault === 'Y'
            ? 'Y'
            : 'N'
        treeItems.push(tempObj)
        this.RecursiveLoadTreeItems(menu.subTree || [], treeItems)
      })
    },
    async btnOK() {
      this.selectDataReadonly = true
      // 所有项
      let treeItems = []
      let { data: templateData, brandId } = this.$refs['select-data'].expose()
      this.RecursiveLoadTreeItems(templateData, treeItems)
      try {
        const { code } = await modifyTemplate(
          { data: treeItems, brandId },
          brandId
        )
        if (code === '0') {
          this.$notify({
            title: '成功',
            message: '修改成功',
            type: 'success'
          })
        }
      } catch (error) {
        console.log('修改error:', error)
      }
    }
  }
}
</script>

<style scoped lang="scss">
.template {
  height: 100%;
  background-color: #fff;
  position: relative;

  .footer-btn {
    text-align: center;
  }

  .wrapper {
    background-color: #fff;
    display: grid;
    grid-template-columns: 6fr 18fr;
    grid-template-rows: 100%;
    grid-gap: 32px;
    height: calc(100% - 56px);
  }
}
</style>
