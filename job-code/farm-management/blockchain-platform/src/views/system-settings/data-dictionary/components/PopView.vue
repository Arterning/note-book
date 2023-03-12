<template>
  <i-dialog :title="title" :visible.sync="visible">
    <el-form :model="form" :rules="rules" ref="form">
      <el-form-item label="名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入名称"> </el-input>
      </el-form-item>
      <el-form-item label="编码" prop="code">
        <el-input v-model="form.code" placeholder="请输入编码"> </el-input>
      </el-form-item>

      <!-- 字典项 -->
      <template v-if="isType === 'newDicValue' || isType === 'editDicValue'">
        <el-form-item label="值" prop="dicValue">
          <el-input v-model="form.dicValue" placeholder="请输入值"> </el-input>
        </el-form-item>
      </template>

      <!-- 字典 -->
      <template v-else>
        <el-form-item label="产品" prop="classify">
          <el-select v-model="form.classify" placeholder="请选择产品">
            <el-option
              v-for="item in product"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="form.type" placeholder="请选择类型">
            <el-option
              v-for="item in dictTypeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </template>

      <el-form-item label="可见状态" prop="radio">
        <el-radio v-model="form.radio" label="Y">可见</el-radio>
        <el-radio v-model="form.radio" label="N">不可见</el-radio>
      </el-form-item>

      <el-form-item label="备注" prop="remark">
        <el-input
          :maxlength="200"
          show-word-limit
          type="textarea"
          :autosize="{ minRows: 3 }"
          placeholder="请输入内容"
          v-model="form.remark"
        >
        </el-input>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="onCancel">取 消</el-button>
      <el-button
        @click="onSubmit"
        type="primary"
        :loading="loading"
        :disabled="isDisabled"
        >确 定</el-button
      >
    </template>
  </i-dialog>
</template>

<script>
import { checkCode } from '@/api/systemSettings'

export default {
  name: 'PopView',

  props: {
    // 用来判断是新增、编辑
    isType: {
      type: String,
      default: ''
    },
    popData: {
      type: Object,
      default: () => ({ type: '', code: '', name: '', classify: '' })
    },
    dictTypeList: {
      type: Array,
      default: () => []
    }
  },

  data() {
    return {
      visible: false,
      loading: false,

      product: [
        {
          name: '通用',
          id: '0'
        }
      ],
      form: {
        name: '',
        dicValue: '',
        radio: 'Y',
        type: '',
        code: '',
        classify: '0',
        remark: ''
      },
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        code: [
          {
            required: true,
            trigger: 'blur',
            validator: this.checkCode
          }
        ],
        dicValue: [{ required: true, message: '请输入值', trigger: 'blur' }],
        radio: [{ required: true, message: '请选择可见状态' }],
        classify: [{ required: true, message: '请选择产品' }],
        type: [{ required: true, message: '请选择类型' }]
      }
    }
  },

  computed: {
    title() {
      if (this.isType === 'newDic') {
        return '新增字典'
      } else if (this.isType === 'newDicValue') {
        return '新增值'
      } else if (this.isType === 'editDic') {
        return '编辑字典'
      } else if (this.isType === 'editDicValue') {
        return '编辑值'
      }
      return ''
    },

    isDisabled() {
      const { name, code, dicValue, radio, type, classify } = this.form

      let flag = !name || !code || !radio
      if (flag) {
        return true
      }

      if (this.isType === 'newDic' || this.isType === 'editDic') {
        return !type || !classify
      } else {
        return !dicValue
      }
    }
  },

  watch: {
    popData(e) {
      this.form.name = e.name
      this.form.dicValue = e.value
      this.form.radio = e.enable
      this.form.code = e.code
      this.form.classify = e.classify
      this.form.remark = e.remark

      // 如果是编辑字典，则需要遍历出类型
      if (this.isType === 'editDic') {
        this.form.type = this.dictTypeList.find(
          dictType => dictType.value === this.popData.type
        ).value
      }
    }
  },

  methods: {
    /**
     * @description: 取消事件
     * @author: Hemingway
     */
    onCancel() {
      this.visible = false
      this.$refs.form.resetFields()
    },

    onSubmit() {
      this.$refs.form.validate(async valid => {
        if (valid) {
          const { name, radio, code, classify, remark } = this.form
          const param = {
            name: name,
            enable: radio,
            code: code,
            classify: classify,
            remark: remark,
            isType: this.isType
          }

          if (this.isType === 'newDic' || this.isType === 'editDic') {
            param['type'] = this.form.type
            param['funId'] =
              this.isType === 'newDic' ? 'createDictionary' : 'editDictionary'
          } else if (
            this.isType === 'newDicValue' ||
            this.isType === 'editDicValue'
          ) {
            param['value'] = this.form.dicValue
            param['funId'] =
              this.isType === 'newDicValue'
                ? 'createDictionaryValue'
                : 'editDictionaryValue'
          }

          if (this.isType === 'editDic' || this.isType === 'editDicValue') {
            // 若是编辑，需要传id
            param['id'] = this.popData.id
          }

          // 提出请求
          this.$emit('submit', param)
        }
      })
    },

    async checkCode(rule, value, callback) {
      if (!value) {
        return callback(new Error('请输入编码'))
      }

      try {
        await checkCode({
          code: value,
          isDictionary:
            this.isType === 'newDic' || this.isType === 'editDic' ? '0' : '1'
        })
        return callback()
      } catch (error) {
        return callback(new Error(error))
      }
    }
  }
}
</script>
