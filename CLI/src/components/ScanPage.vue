<template>
  <div class="scan-page">
    <h2 class="page-title">生鮮のお問い合わせ</h2>
    <qrcode-stream v-if="isScanning" @decode="onDecode" @init="onInit" />
    <div v-if="errorMessage" class="error-message">{{ errorMessage }}</div>
    <button v-if="!isScanning" @click="startScanning">重新扫描</button>
  </div>
</template>

<script>
import { QrcodeStream } from 'vue-qrcode-reader';

export default {
  name: 'ScanPage',
  components: {
    QrcodeStream,
  },
  data() {
    return {
      errorMessage: '',
      isScanning: true, // 控制扫描器的显示
    };
  },
  methods: {
    onDecode(decodedString) {
      alert(`扫码结果: ${decodedString}`);
      // 处理扫码结果，比如发送请求等
    },
    onInit(success, error) {
      if (error) {
        this.errorMessage = error.message;
        this.isScanning = false; // 停止扫描
      } else {
        this.errorMessage = '';
      }
    },
    startScanning() {
      this.isScanning = true; // 重新开始扫描
    }
  },
};
</script>

<style scoped>
.scan-page {
  padding: 20px; /* 页面内边距 */
  height: 100vh; /* 使页面适应手机屏幕 */
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 向上对齐 */
}

.page-title {
  margin: 0; /* 去掉默认的上下间距 */
}

.error-message {
  color: red; /* 错误信息显示为红色 */
  margin-top: 10px; /* 增加一些顶部间距 */
}
</style>
