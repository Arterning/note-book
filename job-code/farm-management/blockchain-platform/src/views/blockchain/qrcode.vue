<template>
  <div v-if="imgPath" class="qrcode">
    <div class="qrcode-img">
      <img :src="imgPath" alt="" srcset="" />
    </div>
    <div id="qrcode">{{ qrcode }}</div>
  </div>
</template>

<script>
import { getSyCode } from '@/api/blockchain'
export default {
  name: 'Qrcode',
  data() {
    return {
      imgPath: '',
      qrcode: '',
      showQrcode: ''
    }
  },

  created() {
    const id = this.$route.params.id
    this.qrcode = 'SY****' + id.substring(id.length - 8)
    getSyCode(id).then(res => {
      if (res.code === '0') this.imgPath = res.data
    })
  }
}
</script>

<style lang="scss">
.qrcode {
  width: 300px;
  height: 300px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.qrcode-img {
  width: 200px;
  height: 200px;

  img {
    width: 100%;
    height: 100%;
  }
}
</style>
