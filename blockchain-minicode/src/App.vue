<!--
 * @Description  : 
 * @Autor        : Hemingway
 * @Date         : 2021-06-30 11:49:38
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2021-09-06 17:09:02
 * @FilePath     : \blockchain-minicode\src\App.vue
-->
<script>
export default {
  onLaunch: () => {
    console.log('App Launch')
  },
  onShow: () => {
    console.log('App Show')
  },
  onHide: () => {
    console.log('App Hide')
  }
}
</script>

<style lang="scss">
/* 每个页面公共css */
page {
  background-color: $uni-bg-color-page;
}

// .content {
//   background: #fff;
// }

.img-box {
  width: 100%;
  height: 176rpx;

  .upload-container {
    display: flex;
    flex-wrap: wrap;

    .upload-img {
      position: relative;

      image {
        display: inline-block;
        width: 112rpx;
        height: 112rpx;
        margin-right: 10rpx;
      }

      .upload-icon {
        position: absolute;

        /* font-size: 121rpx; */
        right: 0;
        top: -16rpx;
      }
    }

    .upload {
      width: 112rpx;
      height: 112rpx;
      border: 1rpx solid #eee;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}

.no-more {
  width: 718rpx;
  margin: 40rpx auto;
  font-size: 24rpx;
  color: #c7c7c7;
  position: relative;

  .text {
    width: 184rpx;
    text-align: center;
    display: inline-block;
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    background: #eee;
  }

  .line {
    position: absolute;
    top: 16rpx;
    width: 100%;
    height: 1rpx;
    background-color: #c7c7c7;
  }
}

.search ::v-deep .uni-easyinput__content-input {
  text-align: left !important;
}
</style>
