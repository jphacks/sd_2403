<template>
  <div class="container">
    <h1 class="title">「フードチェーン」 へ<br> ようこそ</h1>
    <h3>サプライチェーン事業者ぺージ</h3>
    <el-button type="primary" size="large" class="custom-button" @click="goToLogin">MetaMaskでログイン</el-button>
    <el-button type="success" size="large" class="custom-button" @click="goToRegister">新規登録</el-button>
    <el-button type="warning" size="large" class="custom-button" @click="openMetaMask">MetaMaskをインストール</el-button>
    <el-button v-if="(isIOS || isAndroid) && !inMetaMaskApp" type="warning" size="large" class="custom-button" @click="goToMobileCoverLayer">
      スマホ端末専用<br>(MetaMask必要)
    </el-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isIOS: false,
      isAndroid: false,
      inMetaMaskApp: false,
    };
  },
  mounted() {
    this.isIOS = /iPad|iPhone|iPod/.test(navigator.userAgent) && !window.MSStream;
    this.isAndroid = /Android/.test(navigator.userAgent);
    this.inMetaMaskApp = window.ethereum && window.ethereum.isMetaMask;
  },
  methods: {
    goToLogin() {
      this.$router.push('/login');
    },
    goToRegister() {
      this.$router.push('/register');
    },
    openMetaMask() {
      window.open('https://metamask.io/ja/download/', '_blank');
    },
    goToMobileCoverLayer() {
      this.$router.push('/mobile-cover-layer');
    },
  },
};
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100vh;
  justify-content: center;
  max-width: 480px;
  margin: 0 auto;
  padding: 20px;
}

.title {
  font-size: 24px;
  margin-bottom: 40px;
  text-align: center;
}

.custom-button {
  width: 80%;
  height: 60px;
  font-size: 20px;
  margin-bottom: 20px;
}

/* Override Element Plus button styles */
:deep(.el-button--large) {
  padding: 15px 20px;
}
</style>
