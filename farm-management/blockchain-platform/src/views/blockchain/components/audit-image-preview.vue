<template>
  <div class="audit-image-preview">
    <span>{{ group.title }}:</span>
    <div class="image-group-body">
      <div
        :key="index"
        v-for="(item, index) in group.imageList"
        class="image-preview-box"
      >
        <el-image
          :src="item.imageUrl"
          :preview-src-list="imageUrlList"
          :z-index="3000"
          style="width: 50px; height: 50px; border-radius: 5px"
        >
        </el-image>
        <span>{{ item.imageName }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AuditImagePreview',
  props: {
    group: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  computed: {
    imageUrlList() {
      return this.group.imageList.map(({ imageUrl }) => imageUrl)
    }
  }
}
</script>

<style lang="scss" scoped>
.audit-image-preview {
  width: 100%;
  margin-top: 10px;

  .image-group-body {
    margin-bottom: 16px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;

    .image-preview-box {
      text-align: center;
      width: 50px;
      margin-right: 20px;
      overflow: visible;

      span {
        margin-top: 5px;
        color: #333;
        font-size: 13px;
      }
    }
  }
}
</style>
