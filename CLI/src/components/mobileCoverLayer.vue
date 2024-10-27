<template>
  <div class="cover-layer">
    <div class="info-frame">
      <el-alert type="info" title="MetaMask Appにページジャンプします。" />
    </div>
    <el-button type="primary" size="large" class="custom-button" @click="connectToMetaMask">MetaMaskに接続する</el-button>
    <el-button type="danger" class="close-button" @click="closeLayer">x</el-button>
  </div>
</template>

<script>
export default {
  methods: {
    connectToMetaMask() {
      const helloPageUrl = encodeURIComponent(window.location.origin + '/'); // Assuming HelloPage is at the root
      const metamaskAppDeepLink = `https://metamask.app.link/dapp/${helloPageUrl}`;
      window.location.href = metamaskAppDeepLink;
    },
    closeLayer() {
      this.$router.push('/');
    },
  },
};
</script>

<style scoped>
.cover-layer {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
}

.custom-button {
  width: 80%;
  height: 60px;
  font-size: 20px;
  margin-bottom: 20px;
}

.info-frame {
  width: 80%;
  margin-bottom: 20px;
}

/* Override Element Plus button styles */
:deep(.el-button--large) {
  padding: 15px 20px;
}
</style>
