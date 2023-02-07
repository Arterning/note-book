<!--
 * @Description  : 
 * @Autor        : SunXiuqiong
 * @Date         : 2022-06-06 22:47:27
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-08-05 14:29:18
 * @FilePath     : \blockchain-admin\src\views\traceability-manage\clickbind\index.vue
-->
<template>
  <div class="bathBind">
    <i-tab>
      <div class="wrapper">
        <div class="bathInfo">
          <i-section-header>批次信息</i-section-header>
          <el-descriptions>
            <el-descriptions-item label="产品批次">{{
              data.productBatch
            }}</el-descriptions-item>
            <el-descriptions-item label="品种">{{
              data.varietyName
            }}</el-descriptions-item>
            <el-descriptions-item label="商品名称">{{
              data.commodityName
            }}</el-descriptions-item>
            <el-descriptions-item label="所属品牌">
              {{ data.brandName }}
            </el-descriptions-item>

            <el-descriptions-item label="包装日期"
              >{{ data.packingDate }}
            </el-descriptions-item>
            <el-descriptions-item label="包装规格">{{
              data.packingUnit
            }}</el-descriptions-item>
            <el-descriptions-item label="米厂">{{
              data.riceFactoryName
            }}</el-descriptions-item>
            <el-descriptions-item label="来源农场">{{
              data.farmName
            }}</el-descriptions-item>
          </el-descriptions>
        </div>

        <div class="bathInfo">
          <i-section-header>批次绑定</i-section-header>
          <el-descriptions>
            <el-descriptions-item label="该批损坏溯源码">
              <el-input
                placeholder="请输入数量"
                v-model.number="discardChaincodeNum"
                oninput="if(value<0)value=0"
                size="mini"
                :disabled="chaincodeLeft === 0"
              >
                <template slot="append">个</template>
              </el-input>
            </el-descriptions-item>

            <el-descriptions-item label="该批绑定溯源码">
              {{ boundNum }}个
            </el-descriptions-item>
          </el-descriptions>
          <div class="bind-Id">
            <span class="required" style="color: red">* </span>
            <span class="bindid-text">绑定编号：</span>
            <br /><br />

            <el-form :model="formModel" ref="form">
              <template v-for="(row, index) of formModel.rows">
                <div class="formModel__row" :key="row.rowUniqueKey">
                  <el-form-item
                    :prop="'rows.' + index + '.chaincodeStart'"
                    :rules="formRules.chaincodeStart"
                  >
                    <span>SY0DY </span
                    ><el-input
                      type="text"
                      placeholder="请输入开始编号( 8位数字 )"
                      required
                      v-model="row.chaincodeStart"
                      @change="
                        onChaincodeStartChange(row, [
                          'rows.' + index + '.chaincodeStart',
                          'rows.' + index + '.chaincodeEnd'
                        ])
                      "
                      @blur="onChaincodeStartBlur(row)"
                      style="width: 212px"
                      maxlength="8"
                      clearable
                      :disabled="chaincodeLeft === 0"
                    />
                  </el-form-item>
                  <el-form-item
                    :prop="'rows.' + index + '.chaincodeEnd'"
                    :rules="formRules.chaincodeEnd"
                  >
                    <span> ~ SY0DY </span>
                    <el-input
                      type="text"
                      placeholder="请输入结束编号( 8位数字 )"
                      v-model="row.chaincodeEnd"
                      @change="row.chaincodeEnd || (row.chaincodeInRange = 0)"
                      @blur="onChaincodeEndBlur(row)"
                      required
                      style="width: 212px"
                      maxlength="8"
                      clearable
                      :disabled="chaincodeLeft === 0"
                    />
                  </el-form-item>
                  <el-form-item>
                    <div style="width: 3rem; color: red">
                      {{ row.chaincodeInRange || 0 }} 个
                    </div>
                  </el-form-item>
                  <el-form-item>
                    <el-button
                      v-if="index === 0"
                      size="mini"
                      plain
                      type="success"
                      @click="add"
                      :disabled="chaincodeLeft === 0"
                      >新增</el-button
                    >
                    <el-button
                      v-else
                      size="mini"
                      plain
                      type="danger"
                      @click="ondelete(index)"
                      >删除</el-button
                    >
                  </el-form-item>
                </div>
              </template>
            </el-form>
          </div>
        </div>
        <br />
        <div style="padding: 20px 30px; font-size: 14px">
          <div style="margin-bottom: 6px">
            该商品还剩{{ chaincodeLeft || 0 }}个溯源码可使用
            <span v-if="chaincodeLeft && chaincodeLeft != 0">，编号为：</span>
          </div>
          <div v-if="chaincodeLeft && chaincodeLeft != 0">
            {{ cptchaincodeRanges }}
          </div>
        </div>
      </div>

      <i-footer>
        <div style="text-align: center">
          <el-button @click="cancel">取 消</el-button>
          <el-button
            type="primary"
            @click="bonunBtn"
            :disabled="chaincodeLeft === 0"
            >绑 定</el-button
          >
        </div></i-footer
      >
    </i-tab>
  </div>
</template>

<script>
import { getVailableCode, boundChaincode } from '@/api/traceability-manage'

let uniqueKeyGetter = (prefix = 'unique') => {
  let id = 0
  return () => prefix + '_' + ++id
}

export default {
  name: 'clickbind',
  data() {
    let _this = this
    return {
      data: {},
      chaincodeLeft: 0, //剩余可用
      chaincodeBatchs: [], //溯源码批次信息
      discardChaincodeNum: null, //损坏溯源码个数
      chaincodeRanges: [], //可用溯源编号
      formModel: {
        rows: []
      },
      formRules: {
        chaincodeStart: [
          { required: true, message: '请输入起始编号', trigger: 'blur' },
          {
            trigger: 'blur',
            validator(rule, value, callback) {
              if (!/^\d{8}$/.test(value || '')) {
                rule.message = '编号格式有误(仅为8位数字)'
                callback(new Error())
                return
              }

              let intValue = parseInt(value)
              let index = _this.computedChaincodeRanges.findIndex(
                range => intValue >= range[0] && intValue <= range[1]
              )
              if (index === -1) {
                rule.message = '填写的溯源码不在范围内，请先创建溯源码'
                callback(new Error())
                return
              }
              callback()
            }
          }
        ],
        chaincodeEnd: [
          { required: true, message: '请输入结束编号', trigger: 'blur' },
          {
            trigger: 'blur',
            validator(rule, value, callback, field) {
              let chaincodeStart =
                _this.formModel.rows[
                  Object.keys(field)[0].replace(/[^\d]/g, '')
                ].chaincodeStart

              if (!/^\d{8}$/.test(chaincodeStart || '')) {
                rule.message = '请先输入正确的起始编号'
                callback(new Error())
                return
              }

              if (!/^\d{8}$/.test(value || '')) {
                rule.message = '编号格式有误(仅为8位数字)'
                callback(new Error())
                return
              }

              let intChaincodeStart = parseInt(chaincodeStart)
              //根据前面已输入开始编号，找到对应的可输入区间
              let chaincodeRange = _this.computedChaincodeRanges.find(
                range =>
                  intChaincodeStart >= range[0] && intChaincodeStart <= range[1]
              )

              let intValue = parseInt(value)
              if (intValue < intChaincodeStart) {
                rule.message = '结束编号需大于等于开始编号'
                callback(new Error())
                return
              }

              if (intValue > chaincodeRange[1]) {
                rule.message = `编号不在(${chaincodeRange
                  .map(code => (code + '').padStart(8, '0'))
                  .join('~')})范围内`
                callback(new Error())
                return
              }

              callback()
            }
          }
        ]
      }
    }
  },

  created() {
    this.getLocalData()
    this.getVailableCode()
    this.formModelUniqueKeyGetter = uniqueKeyGetter('formModel-rows')
    this.add()
  },

  computed: {
    cptchaincodeRanges() {
      return this.chaincodeRanges.toString().replace(',', '，')
    },
    computedChaincodeBatchs() {
      return this.chaincodeBatchs.map(item => {
        return {
          chaincodeBatch: item.chaincodeBatch,
          chaincodeBatchId: item.chaincodeBatchId,
          chaincodeRanges: (item.chaincodeRanges || '')
            .split('~')
            .map(str => str.replace(item.chaincodePrefix, ''))
        }
      })
    },
    computedChaincodeRanges() {
      return this.computedChaincodeBatchs.map(batchs =>
        batchs.chaincodeRanges.map(str => parseInt(str))
      )
    },
    // 溯源码绑定数量：损坏+包装件数
    boundNum() {
      return Number(this.discardChaincodeNum) + Number(this.data.packingNumber)
    },

    // 填写区间编号的总个数
    totalNum() {
      let sum = 0
      this.formModel.rows.forEach(item => {
        sum += item.chaincodeInRange
      })
      return sum
    }
  },

  methods: {
    getLocalData() {
      const obj = localStorage.getItem('rowData')
      console.log('初始化', JSON.parse(obj))
      this.data = JSON.parse(obj)
    },

    // 获取剩余溯源码及编号
    async getVailableCode() {
      try {
        const params = {
          brandId: this.data.brandId,
          commodityId: this.data.commodityId
        }
        const {
          data: { chaincodeLeft, chaincodeBatchs }
        } = await getVailableCode(params)

        this.chaincodeLeft = chaincodeLeft

        this.chaincodeBatchs = chaincodeBatchs

        this.chaincodeBatchs.forEach(e => {
          if (e.chaincodeRanges) {
            this.chaincodeRanges.push(e.chaincodeRanges)
          }
        })
      } catch (error) {
        console.log('查询剩余码 err:', error)
      }
    },

    onChaincodeStartBlur(row) {
      let value = row.chaincodeStart
      if (/^\d{8}$/.test(value || '')) {
        let intValue = parseInt(value)
        let index = this.computedChaincodeRanges.findIndex(
          range => intValue >= range[0] && intValue <= range[1]
        )
        if (index !== -1) {
          let batch = this.computedChaincodeBatchs[index]
          row.chaincodeBatch = batch.chaincodeBatch
          row.chaincodeBatchId = batch.chaincodeBatchId
          row.range_ = this.computedChaincodeRanges[index]
        } else {
          row.chaincodeBatch = null
          row.chaincodeBatchId = null
          row.range_ = []
        }
      }
    },

    onChaincodeStartChange(row, props) {
      if (!row.chaincodeStart) {
        row.chaincodeEnd = null
        row.chaincodeBatch = null
        row.chaincodeBatchId = null
        row.chaincodeInRange = 0
        this.$nextTick(() => void this.$refs.form.clearValidate(props))
      }
    },

    onChaincodeEndBlur(row) {
      let value = row.chaincodeEnd
      //先判断是否输入开始编号(具有range_属性,且长度为2)
      if (/^\d{8}$/.test(row.chaincodeStart || '') && row.range_.length === 2) {
        let intChaincodeStart = parseInt(row.chaincodeStart)
        let intValue = parseInt(value)
        let flag = [
          //判断结束编号是否符合规则
          /^\d{8}$/.test(value || ''),
          //判断结束编号是否在区间内
          intValue >= intChaincodeStart,
          intValue <= row.range_[1]
        ].every(it => it)
        row.chaincodeInRange = flag ? intValue - intChaincodeStart + 1 : 0
      }
    },

    add() {
      this.formModel.rows.push({
        rowUniqueKey: this.formModelUniqueKeyGetter(),
        chaincodeStart: null, //区间开始
        chaincodeEnd: null, //区间结束
        chaincodeInRange: null, //该区间个数
        chaincodeBatch: null, //区块批次
        chaincodeBatchId: null, //区块批次Id
        range_: []
      })
    },

    ondelete(index) {
      this.formModel.rows.splice(index, 1)
    },

    bonunBtn() {
      this.$refs.form.validate(async flag => {
        if (flag) {
          console.log('123')
          // 各区间相加必须等于上面的溯源码绑定个数。若未等于则提示“填写的溯源码少于/超过需绑定数量），需等于后才能进行提交操作。
          if (this.totalNum > this.boundNum) {
            this.$message.error('填写的溯源码超过需绑定的数量!')
            return
          }
          if (this.totalNum < this.boundNum) {
            this.$message.error('填写的溯源码少于需绑定的数量!')
            return
          }
          let payload = JSON.parse(JSON.stringify(this.formModel.rows))
          payload.forEach(item => {
            delete item.rowUniqueKey
            delete item.range_
          })

          try {
            const params = {
              packingBatch: this.data.packingBatch,
              machPackingId: this.data.packingInfoId,
              discardChaincodeNum: String(
                this.discardChaincodeNum ? this.discardChaincodeNum : 0
              ),
              totalBondChainCode: String(this.boundNum),
              bondRanges: payload
            }
            const { code } = await boundChaincode(params)
            if (code === '0') {
              this.$notify({
                title: '成功',
                message: '绑定成功！',
                type: 'success',
                duration: 2000
              })
              this.$store.dispatch('tagsView/delView', this.$route)
              this.$router.back()
            }
          } catch (error) {
            console.log('绑定溯源码提交 err:' + error)
          }
        }
      })
    },

    cancel() {
      this.$router.back()
      this.$store.dispatch('tagsView/delView', this.$route)
    }
  }
}
</script>

<style lang="scss" scoped>
.bathBind {
  height: 100%;
  background-color: #fff;
  position: relative;

  .wrapper {
    background-color: #fff;
    padding-bottom: 100px;

    .bind-Id {
      width: 100%;
      padding: 20px 30px 20px 30px;
      margin-bottom: -40px;

      .bindid-text {
        color: #606266;
        font-size: 14px !important;
        margin-bottom: 20px;
      }

      .formModel__row {
        display: flex;
      }
      .formModel__row > * {
        margin-right: 10px;
      }
      .formModel__row > *:nth-last-child(1) {
        margin-right: 0;
      }

      span {
        font-size: 14px;
      }
    }
  }
}
</style>

<style>
.el-descriptions-item__content {
  color: #000;
  font-weight: bold;
}

.el-descriptions--medium:not(.is-bordered) .el-descriptions-item__cell {
  padding: 0 30px 10px 30px;
}

el-descriptions-item__label.has-colon::after {
  content: '：';
  position: relative;
  top: -0.5px;
}
</style>
