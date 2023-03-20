<!--
 * @Description  :
 * @Autor        : SunXiuqiong
 * @Date         : 2022-07-20 11:17:28
 * @LastEditors  : Hemingway
 * @LastEditTime : 2022-10-25 16:26:26
 * @FilePath     : \blockchain-platform\src\views\tenant-manage\authorize-manage\components\AddEdit.vue
-->
<template>
  <div>
    <i-section-header title="基本信息" />
    <el-form label-suffix="：" :model="baseInfoFrom" :rules="baseInfoFromRules">
      <el-row>
        <el-col :span="8">
          <el-form-item label="企业名称" prop="name">
            <el-input
              style="width: 220px"
              v-model="baseInfoFrom.name"
            ></el-input></el-form-item
        ></el-col>
        <el-col :span="8">
          <el-form-item label="企业类型" prop="type">
            <el-select
              style="width: 220px"
              v-model="baseInfoFrom.type"
              @change="companyTypeChange"
            >
              <el-option
                v-for="item in companyTypeArr"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
              </el-option>
            </el-select> </el-form-item
        ></el-col>

        <el-col :span="8">
          <el-form-item label="服务模式" prop="serviceModelId">
            <el-select
              style="width: 220px"
              v-model="baseInfoFrom.serviceModelId"
              @change="serviceModeChange"
            >
              <el-option
                v-for="item in serviceModeArr"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item
            label="使用时间"
            required
            v-if="['1'].includes(serverType)"
          >
            <div style="display: flex">
              <el-form-item prop="startDate">
                <el-date-picker
                  type="date"
                  placeholder="选择开始日期"
                  v-model="baseInfoFrom.startDate"
                  format="yyyy-MM-dd"
                  value-format="yyyy-MM-dd"
                  style="width: 150px"
                ></el-date-picker>
              </el-form-item>
              <span style="padding: 0 5px">~</span>
              <el-form-item prop="endDate">
                <el-date-picker
                  type="date"
                  placeholder="选择结束日期"
                  v-model="baseInfoFrom.endDate"
                  format="yyyy-MM-dd"
                  value-format="yyyy-MM-dd"
                  style="width: 150px"
                ></el-date-picker>
              </el-form-item>
            </div>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item
            label="允许使用码范围"
            v-if="['2'].includes(serverType)"
          >
            <el-input-number
              v-model="baseInfoFrom.limit"
              controls-position="right"
              style="width: 220px"
              @change="handleCodeScope"
              :min="0"
              :max="99999999"
            ></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="系统管理员姓名" prop="manager">
            <el-input
              style="width: 220px"
              v-model="baseInfoFrom.manager"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="管理员手机号" prop="phone">
            <el-input
              style="width: 220px"
              v-model="baseInfoFrom.phone"
              @input="adminPhoneChange"
            ></el-input>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>

    <i-section-header
      title="相关附件"
      remark="建议上传合同等相关材料，可上传jpg、png图片和pdf格式的附件，单个附件不超过10M，最多可上传10个"
    />
    <i-upload
      :limit="limit"
      :limitSize="limitSize"
      accept=".jpg,.png,.pdf"
      uploadUrl="/blockchaincomponent/file/uploadFile/w/v1"
      :params="params"
      @change="filelistsChange"
      :files="fileArr"
      @uploadOver="e => (submitLoading = !e)"
    ></i-upload>

    <i-section-header title="地址定位" />
    <el-form style="margin-bottom: 56px">
      <el-form-item label="企业地址" label-suffix="：" required>
        <el-button class="pickmap-btn" @click="pickAddress">{{
          baseInfoFrom.address ? baseInfoFrom.address : '点击跳转到地址选择'
        }}</el-button>
      </el-form-item>
    </el-form>

    <i-footer>
      <el-button plain size="small" style="margin-right: 20px" @click="cancel"
        >关 闭</el-button
      >
      <el-button
        type="primary"
        size="small"
        @click="submit"
        :loading="submitLoading"
        >提 交</el-button
      ></i-footer
    >

    <i-dialog title="企业地址" :visible.sync="dialogVisible" class="map-dialog">
      <i-amap
        class="map_address_container"
        :lng="lng"
        :lat="lat"
        @returnAddress="getDetailedAddress"
      ></i-amap>
    </i-dialog>

    <!-- <el-dialog title="企业地址" :visible="dialogVisible" :show-close="false">
      <i-amap
        class="map_address_container"
        :lng="lng"
        :lat="lat"
        @returnAddress="getDetailedAddress"
      ></i-amap>
    </el-dialog> -->
  </div>
</template>

<script>
import { getService } from '@/api/tenant-manage'
import {
  getTenantAdd,
  getTenantUpdate,
  getDictQuery,
  getEnterpriseName
} from '@/api/authorize-manage'
import { empDictValues } from '@/utils/smallTools'

export default {
  name: 'AddEdit',

  props: {
    rows: {
      //数据类型有 String、Number、Boolean、Array ...
      type: Object
    }
  },
  data() {
    // form表单自定义验证规则，开始时间不可大于结束时间
    let startRule = (rule, value, callback) => {
      if (!value) {
        callback(new Error('请选择开始时间'))
      } else {
        if (!['', undefined, null].includes(this.baseInfoFrom.endDate)) {
          const _startTime = new Date(value).getTime()
          const _endTime = new Date(this.baseInfoFrom.endDate).getTime()
          if (_startTime > _endTime) {
            callback(new Error('开始时间必须小于结束时间'))
          } else {
            callback()
          }
        } else {
          callback()
        }
      }
    }

    // 结束时间不可小于开始时间
    let endRule = (rule, value, callback) => {
      if (!value) {
        callback(new Error('请选择结束时间'))
      } else {
        if (!['', undefined, null].includes(this.baseInfoFrom.startDate)) {
          const _startTime = new Date(this.baseInfoFrom.startDate).getTime()
          const _endTime = new Date(value).getTime()
          if (_endTime < _startTime) {
            callback(new Error('结束时间必须大于开始时间'))
          } else {
            callback()
          }
        } else {
          callback()
        }
      }
    }

    return {
      serverType: '0',
      limit: 10,
      limitSize: 10,
      params: { fileClass: 'enterpriseImages' },
      dialogVisible: false,
      fileLists: [],
      fileArr: [],
      baseInfoFrom: {
        name: '',
        type: '',
        serviceModelId: '',
        startDate: '',
        endDate: '',
        limit: '',
        manager: '',
        phone: '',
        attachmentIds: [],
        provinceCode: '',
        cityCode: '',
        areaCode: '',
        address: '',
        latitude: '',
        longitude: ''
      },
      baseInfoFromRules: {
        name: [{ required: true, message: '请填写企业名称', trigger: 'blur' }],
        type: [
          { required: true, message: '请选择企业类型', trigger: 'change' }
        ],
        serviceModelId: [
          { required: true, message: '请选择服务模式', trigger: 'change' }
        ],
        startDate: [
          { required: true, validator: startRule, trigger: 'change' }
        ],
        endDate: [{ required: true, validator: endRule, trigger: 'change' }],
        limit: [
          { required: true, message: '请输入使用码范围', trigger: 'blur' }
        ],
        manager: [
          { required: true, message: '请输入系统管理员姓名', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '请输入管理员手机号', trigger: 'blur' }
        ]
      },
      // 服务模式
      serviceModeArr: [],
      // 企业类型
      companyTypeArr: [
        {
          id: 3,
          name: '企业类型1'
        },
        {
          id: 4,
          name: '企业类型2'
        }
      ],
      // el-amap组件 经 纬 度
      lng: '112.879456',
      lat: '28.227257',
      submitLoading: false
    }
  },
  created() {},
  watch: {
    rows: {
      async handler(newRows) {
        if (newRows.model === 'edit') {
          // 回显字段同步
          newRows.row.phone = newRows.row.linkmanPhone
          newRows.row.manager = newRows.row.linkman
          newRows.row.limit = newRows.row.codeLimit
          newRows.row.startDate = newRows.row.serviceStartDate
          newRows.row.endDate = newRows.row.serviceEndDate
          // 回显服务模式下拉数据
          await this.getRole(newRows.row.type)
          // 回显服务模式下限制类型
          this.serverType = this.serviceModeArr.filter(
            res => res.id === newRows.row.serviceModelId
          )[0].type
          // 回显图片
          newRows.row.imageIds &&
            (await this.getCompanyInfo(JSON.parse(newRows.row.imageIds)))
          console.log('newRows.model', newRows)
          if (['edit'].includes(newRows.model)) {
            this.baseInfoFrom = Object.assign(
              {},
              this.baseInfoFrom,
              newRows.row
            )
            if (
              ['', undefined, null].includes(this.baseInfoFrom.longitude) &&
              ['', undefined, null].includes(this.baseInfoFrom.latitude)
            ) {
              this.lng = this.baseInfoFrom.longitude
              this.lat = this.baseInfoFrom.latitude
            }
          } else {
            this.baseInfoFrom = this.$options.data().baseInfoFrom
          }
          console.log('this.baseInfoFrom', this.baseInfoFrom)
        }
      },
      deep: true
      // immediate: true
    }
  },
  mounted() {
    this.getDictQuery()
  },
  methods: {
    // 管理员手机号改变事件
    async adminPhoneChange(val) {
      var reg =
        /^1((34[0-8])|(8\d{2})|(([35][0-35-9]|4[579]|66|7[35678]|9[1389])\d{1}))\d{7}$/

      if (reg.test(val)) {
        let resut = await getEnterpriseName({ phone: val })
        if (resut.data) {
          this.baseInfoFrom.name = resut.data.name
        }
      }
    },
    // 回显图片信息
    async getCompanyInfo(imgIdArr) {
      this.fileArr = []
      imgIdArr.forEach(item => {
        if (item) {
          this.fileArr.push({
            name: item + '.jpg',
            url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${item}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`,
            response: {
              code: '0',
              msg: 'OK',
              id: item,
              url: `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${item}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
            },
            fileId: item
          })
        }

        console.log('123')
      })
    },
    // 获取地图组件地址数据
    getDetailedAddress(addressObj) {
      console.log(`addressObj`, addressObj)
      this.dialogVisible = false
      let {
        provinceCode,
        cityCode,
        districtCode,
        fullAddress,
        lng,
        lat,
        province,
        city,
        district
      } = addressObj
      let _addressObj = {
        provinceCode: provinceCode,
        cityCode: cityCode,
        areaCode: districtCode,
        address: fullAddress,
        longitude: lng,
        latitude: lat,
        province,
        city,
        district
      }
      this.baseInfoFrom = Object.assign({}, this.baseInfoFrom, _addressObj)
    },

    // 地图弹框
    pickAddress() {
      this.dialogVisible = true
    },
    // 提交按钮
    submit() {
      let api = () => {}
      if (['add'].includes(this.rows.model)) {
        api = getTenantAdd
      } else if (['edit'].includes(this.rows.model)) {
        api = getTenantUpdate
      }
      let _params = {
        attachmentIds: this.fileLists.map(el => el.fileId)
      }
      this.baseInfoFrom = Object.assign({}, this.baseInfoFrom, _params)
      this.submitLoading = true
      api(this.baseInfoFrom)
        .then(() => {
          this.$notify({
            message: '提交成功',
            type: 'success',
            duration: 2000,
            onClose: () => {
              this.cancel()
            }
          })
          this.submitLoading = false
        })
        .catch(err => {
          console.log('提交异常 error = ', err)
          this.submitLoading = false
        })
    },
    // 取消按钮
    cancel() {
      this.$store.dispatch('tagsView/delView', this.$route).then(() => {
        // 恢复原始数据
        this.baseInfoFrom = this.$options.data().baseInfoFrom
        // 返回上一步路由
        this.$router.go(-1)
      })
    },
    // 使用码范围改变事件
    handleCodeScope() {},
    // 文件列表变更事件
    filelistsChange(e) {
      console.log(`fileListsChange`, e.fileList)
      this.fileLists = e.fileList
    },
    // 企业类型字典
    async getDictQuery() {
      try {
        let { dictMap } = await getDictQuery()
        this.companyTypeArr = empDictValues(dictMap.TENANT_TYPE)
      } catch (error) {
        console.log('error', error)
      }
    },
    // 企业类型改变事件
    async companyTypeChange(item) {
      if (!item) {
        this.serviceModeArr = []
        return
      }
      this.baseInfoFrom.serviceModelId = ''
      this.getRole(item)
    },
    // 服务模式改变事件
    serviceModeChange(val) {
      if (!val) {
        this.serverType = '0'
        return
      }
      let resultArr = this.serviceModeArr.filter(res => {
        return res.id === val
      })
      this.serverType = resultArr[0].type
    },
    // 服务模式列表
    async getRole(enterpriseType) {
      try {
        let { data } = await getService({
          enterpriseType,
          enable: 'Y'
        })
        this.serviceModeArr = data
      } catch (error) {
        console.log('error', error)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.map_address_container {
  margin-top: -20px;
}
.pickmap-btn {
  width: 500px;
  height: 36px;
  padding: 0 10px;
  line-height: 36px;
  text-align: left;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

// .map_address {
//   /* height:100vh; */
//   height: 550px;
// }
// .address-wrapper {
//   display: flex;
//   flex-direction: column;
// }
// .search-box {
//   /* width:100vw; */
//   width: 400px;
//   height: 40px;
// }
// .amap-demo {
//   flex: 1;
//   /* height:100vh; */
//   height: 550px;
// }

// map-dialog
.map-dialog ::v-deep {
  .el-dialog {
    margin: 5vh auto !important;
    .el-dialog__body {
      max-height: none;
      padding: 10px;
    }
  }
}
</style>
