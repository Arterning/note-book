<!--
 * @Description  : 投诉商品详情
 * @Autor        : qinjj
 * @Date         : 2022-07-19 16:59:00
 * @LastEditors  : Your Name
 * @LastEditTime : 2022-09-07 17:41:24
 * @FilePath     : \blockchain-platform\src\views\traceability-supervision\complaint-goods\detail.vue
-->
<template>
  <div class="authorize-manage">
    <i-tab>
      <i-section-header title="商品信息" class="title-box" />
      <div class="good-box">
        <div class="good-item">
          <span>商品名称：</span>
          <span>{{ rowInfo.commodityName }}</span>
        </div>
        <div class="good-item">
          <span>溯源编码：</span>
          <span>{{ rowInfo.tracingCode }}</span>
        </div>
        <div class="good-item">
          <span>所属品牌：</span>
          <span>{{ rowInfo.brandName }}</span>
        </div>
        <div class="good-item">
          <span>所属企业：</span>
          <span>{{ rowInfo.enterpriseName }}</span>
        </div>
        <div class="good-item">
          <span>企业联系电话：</span>
          <span>{{ rowInfo.enterprisePhone }}</span>
        </div>
        <div class="good-item">
          <span>企业联系地址：</span>
          <span>{{ rowInfo.enterpriseAddress }}</span>
        </div>
      </div>
      <i-section-header title="投诉信息" class="title-box" />
      <div class="complaint-info">
        <div class="complaint-item">
          <span>投诉人姓名：</span>
          <span>{{ rowInfo.complainant }}</span>
        </div>
        <div class="complaint-item">
          <span>联系电话：</span>
          <span>{{ rowInfo.complainantsHotline }}</span>
        </div>
        <div class="complaint-item">
          <span>联系邮箱：</span>
          <span>{{ rowInfo.mailbox }}</span>
        </div>
      </div>
      <div class="complaint-reson">
        <div>投诉理由：</div>
        <el-input
          type="textarea"
          v-model="rowInfo.complaintsReasons"
          disabled
          :rows="5"
        ></el-input>
      </div>
      <div class="complaint-image">
        <div>相关附件：</div>
        <div class="image-box">
          <el-image
            v-for="(url, index) in filesArr"
            :key="index"
            style="width: 100px; height: 100px"
            :src="url"
            :preview-src-list="filesArr"
          >
          </el-image>
        </div>
      </div>
      <i-section-header title="处理结果" class="title-box" />
      <div class="handle-result">
        <el-input
          type="textarea"
          v-model="rowInfo.handleResult"
          disabled
          :rows="5"
        ></el-input>
      </div>
      <div class="footer-box">
        <el-button size="small" @click="clickRetrunBTn">返回</el-button>
      </div>
    </i-tab>
  </div>
</template>

<script>
export default {
  name: 'blockchain-complaint-goods-detail',
  data() {
    return {
      rowInfo: {},
      filesArr: []
    }
  },
  mounted() {
    this.rowInfo = this.$route.params.row
    this.echoPic(this.rowInfo.complaintImages)
  },
  methods: {
    clickRetrunBTn() {
      this.$router.back()
      this.$store.dispatch('tagsView/delView', this.$route)
    },
    //回显图片
    echoPic(complaintImages) {
      if (!complaintImages) return
      let picIdArr = complaintImages.split(',')
      for (var i = 0; i < picIdArr.length; i++) {
        this.filesArr.push(
          `${process.env.VUE_APP_BASE_URL}/blockchaincomponent/file/downloadFile/w/v1?id=${picIdArr[i]}&isAbbreviation=N&sessionId=${this.$store.state.user.sessionId}&clientId=poweb`
        )
      }
    },
    onAdd() {
      this.$router.push({
        name: 'add-tenement'
      })
    },
    detailBtn(row) {
      this.$router.push({
        name: 'edit-tenement',
        params: { id: row.id }
      })
    }
  },
  computed: {}
}
</script>

<style scoped lang="scss">
.title-box {
  ::v-deep .left-side > .title {
    font-weight: 700;
  }
}
.authorize-manage {
  height: 100%;
  background-color: #fff;
  position: relative;

  .footer-box {
    padding: 20px;
    text-align: center;
    border-top: 1px solid #ccc;
    margin-top: 25px;
  }

  .good-box {
    display: flex;
    flex-wrap: wrap;
    padding-left: 13px;
    margin-bottom: 25px;
    .good-item {
      width: 33.3%;
    }
    .good-item:nth-child(-n + 3) {
      margin-bottom: 20px;
    }
  }

  .complaint-info {
    display: flex;
    flex-wrap: wrap;
    padding-left: 13px;
    .complaint-item {
      width: 33.3%;
      margin-bottom: 20px;
    }
  }

  .complaint-reson {
    padding-left: 13px;
    margin-bottom: 20px;
    div:nth-child(-n + 3) {
      margin-bottom: 10px;
    }
  }

  .complaint-image {
    padding-left: 13px;
    margin-bottom: 25px;
    div:nth-child(-n + 3) {
      margin-bottom: 10px;
    }
  }

  .handle-result {
    padding-left: 13px;
  }
}
</style>
