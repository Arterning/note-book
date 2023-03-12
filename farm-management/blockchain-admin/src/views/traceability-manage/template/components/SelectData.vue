<!--
 * @Description  : 数据模板勾选组件
 * @Autor        : SunXiuqiong
 * @Date         : 2022-06-08 13:46:25
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-21 09:16:05
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\template\components\SelectData.vue
-->
<template>
  <div class="main">
    <span class="title">勾选展示的数据</span>
    <div class="main-left">
      <div class="one">
        <span>品牌</span>
        <el-select
          size="mini"
          v-model="brandId"
          filterable
          placeholder="请选择品牌"
          @change="onChange"
        >
          <el-option
            v-for="item in brandList"
            :key="item.brandName"
            :label="item.brandName"
            :value="item.brandId"
          ></el-option>
        </el-select>
        <el-button size="mini" type="primary" @click="onPreview"
          >查看预览</el-button
        >
      </div>
      <span v-if="data.length > 0" class="label-name">勾选需要展示的模块</span>
      <div class="module-box" v-if="data.length > 0">
        <el-checkbox-group :disabled="readonly" v-model="checkList">
          <el-checkbox
            v-for="item in data"
            :key="item.moduleId"
            :label="item.moduleId"
            :disabled="item.isDefault === 'Y'"
            @change="onFirstLevelMenuChange(item, $event)"
          >
            {{ item.menuName }}
          </el-checkbox>
        </el-checkbox-group>
      </div>
      <span class="label-name">勾选各模块需要展示的内容</span>
      <div class="three">
        <div class="ipt-search">
          <el-input
            placeholder="内容搜索"
            v-model="filterText"
            size="mini"
            prefix-icon="el-icon-search"
          >
          </el-input>
        </div>
        <el-tree
          style="margin-top: 4px"
          class="filter-tree"
          :data="menuList"
          show-checkbox
          node-key="menuId"
          :check-strictly="true"
          :props="defaultProps"
          default-expand-all
          :filter-node-method="filterNode"
          :default-checked-keys="defaultCheckIds"
          @check-change="onCheckChange"
          ref="tree"
        >
        </el-tree>
      </div>
    </div>
  </div>
</template>

<script>
import { getbrandList, getTemplate } from '@/api/traceability-manage'
import { mapGetters } from 'vuex'
export default {
  name: 'SelectData',
  props: { readonly: Boolean, getPrivewInstance: Function },
  data() {
    let _this = this
    return {
      filterText: '',
      brandId: '',
      brandList: [],
      data: [],
      cacheRequestData: [],
      checkList: [], //选择模块
      defaultCheckIds: [], //默认选择的id列表
      defaultProps: {
        children: 'subTree',
        label: 'menuName',
        disabled: menu => (_this.readonly ? true : menu.isDefault === 'Y')
      }
    }
  },
  computed: {
    menuList() {
      // console.log(this.checkList)
      return this.data.filter(item => this.checkList.includes(item.moduleId))
    },
    ...mapGetters('user', { computedUserInfo: 'userInfoGetter' })
  },
  created() {
    ;(async () => {
      await this.getbrandList()
      this.getTemplate()
    })()
  },

  methods: {
    // 获取所有品牌
    async getbrandList() {
      const { data } = await getbrandList({
        pageSize: -1,
        enterpriseId: this.computedUserInfo.enterpriseId
      })
      this.brandList = data
      if (!this.brandList || this.brandList.length === 0) {
        this.$emit('isEditEvent', true)
      }
      this.brandList?.length && (this.brandId = this.brandList[0].brandId)
    },

    // 预览，获取预览组件实例，然后调用postMessage方法
    onPreview() {
      this.$nextTick(() => {
        let data = JSON.parse(JSON.stringify(this.data)) || []
        let cacheData = JSON.parse(JSON.stringify(this.cacheRequestData))
        console.log('123', cacheData)
        console.log('234', data)
        data.forEach((item, index) => {
          item.isChecked = this.checkList.includes(item.moduleId) ? 'Y' : 'N'
          if (item.menuId === cacheData[index].moduleId) {
            ;['isDefault', 'isChecked'].forEach(key => {
              cacheData[index][key] = item[key]
            })
            cacheData[index].subTree = cacheData[index].subTree ? item : null
          }
        })
        console.log('12', cacheData)
        this.getPrivewInstance()?.postMessage(cacheData)
      })
    },
    // 暴露给父组件，获取内部data,brandId
    expose() {
      return { data: this.data, brandId: this.brandId }
    },

    // 模板查询树形数据
    async getTemplate() {
      try {
        const { data } = await getTemplate({
          brandId: this.brandId
        })
        this.cacheRequestData = data
        let templateData = data.map(item => ({
          menuName: item.moduleName,
          menuId: item.moduleId,
          ...(item.subTree ? item.subTree : item)
        }))

        let defaultCheckList = []
        this.RecursiveProcess(templateData || [], defaultCheckList)

        this.defaultCheckIds = defaultCheckList
        this.data = templateData
        this.checkList = data
          .filter(item => item.isChecked === 'Y')
          .map(item => item.moduleId)

        this.onPreview()
      } catch (error) {
        console.log('树形查询 err:' + error)
      }
    },

    // 设置是否选中
    // eslint-disable-next-line no-unused-vars
    onCheckChange(menu, flag, halfCheck) {
      console.log(333, menu, flag, halfCheck)
      if ('isChecked' in menu) {
        menu.isChecked = flag ? 'Y' : 'N'
      }
      // this.RecursiveToggleNodeCheckedStatus(menu, flag ? 'Y' : 'N')
    },

    // 递归处理tree
    RecursiveProcess(menus, checkList = []) {
      menus.forEach(menu => {
        if (menu.isChecked === 'Y' || menu.isDefault === 'Y') {
          checkList.push(menu.menuId)
        }
        this.RecursiveProcess(menu.subTree || [], checkList)
      })
    },

    async onChange() {
      this.getTemplate()
    },
    RecursiveToggleNodeCheckedStatus(menu, isChecked) {
      menu.isChecked = isChecked
      if (menu.subTree) {
        menu.subTree.forEach(item =>
          this.RecursiveToggleNodeCheckedStatus(item, isChecked)
        )
      }
    },
    onFirstLevelMenuChange(item, flag) {
      item.isChecked = flag ? 'Y' : 'N'
      // this.RecursiveToggleNodeCheckedStatus(item, flag ? 'Y' : 'N')
    },

    filterNode(value, data) {
      if (!value) return true
      return data.menuName.indexOf(value) !== -1
    }
  },

  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  }
}
</script>

<style scoped lang="scss">
.main {
  color: #212121;
  font-size: 12px;
  height: 100%;
  display: flex;
  flex-direction: column;
  .title {
    font-weight: bold !important;
    font-size: 14px !important;
  }
  span {
    display: block;
    margin: 4px;
  }
  .main-left {
    background-color: #fff;
    // min-height: 485px;
    // max-height: 1000px;
    border: 1px solid rgb(234, 230, 230);
    display: flex;
    flex-direction: column;
    flex: 1;
    overflow-y: auto;

    .one {
      background: #fff;
      display: flex;
      padding: 10px 6px;
      justify-content: space-around;
      align-items: center;
      > * {
        margin-right: 10px;
      }
      > *:nth-child(1),
      > * :nth-last-child(1) {
        white-space: nowrap;
        flex-shrink: 0;
      }
      > *:nth-last-child(1) {
        margin-right: 0;
      }
    }
    .module-box {
      background: #fff;
      padding: 6px;
      display: flex;
    }

    .three {
      position: relative;
      // height: 500px;
      flex: 1;
      background: #fff;
      overflow-y: auto;
      padding-bottom: 10px;

      ::v-deep .filter-tree {
        padding: 0 6px;

        // 隐藏filter-tree的一级node
        > .el-tree-node > .el-tree-node__content > .el-checkbox {
          display: none;
        }

        $treemaxlevel: 4; //调整filter-tree的node间距 当前支持4级，可调整
        @for $i from 1 through $treemaxlevel {
          $attr: "style *= 'padding-left: " + 18px * $i + ";'";
          @if ($i == 1) {
            .el-tree-node__content[#{$attr}] {
              padding-left: 0px !important;
            }
          } @else {
            .el-tree-node__content[#{$attr}] {
              padding-left: 10px * $i !important;
            }
          }
        }
      }

      .ipt-search {
        position: sticky;
        z-index: 999;
        top: 0;
        width: 100%;
        background: #ffffff;
        padding: 6px 6px 2px 6px;
        box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.03);
      }
    }
  }
}
</style>
<style>
.el-checkbox__input.is-checked + .el-checkbox__label {
  color: #212121;
}
.el-checkbox__input.is-disabled + span.el-checkbox__label {
  color: #212121;
}
.el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color: #bcbfc3;
}
.el-checkbox__input.is-disabled.is-checked .el-checkbox__inner::after {
  border-color: #fff;
}
</style>
<style lang="scss"></style>
