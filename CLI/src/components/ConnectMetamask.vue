<template>
  <div id="app">
    <el-button v-if="!account" @click="enableEthereum" class="enableEthereumButton" type="success">MetaMaskに接続</el-button>
    <h2 v-if="account">
      <el-alert type="info" title="Account public-key:" :description="account" />
      <el-alert type="success" title="ウォレットアドレスの認証：" description="この接続により、サービスがあなたのウォレットアドレスを一意のIDとして識別するための認証が完了しました。" />
      <el-alert type="info" title="ネットワークの互換性チェック：" description="現在接続しているブロックチェーンネットワーク（例: Ethereum Mainnet)が、このサービスと互換性があると確認されました。" />
      <el-alert type="warning" title="プライバシー保護：" description="このサービスは、あなたのウォレットアドレスのみを参照しており、プライベートキーやその他の個人情報には一切アクセスしません。" />
    </h2>
  </div>
</template>

<script>
import detectEthereumProvider from "@metamask/detect-provider";
// import WalletConnect from "@walletconnect/client";
// import QRCodeModal from "@walletconnect/qrcode-modal";

export default {
  data() {
    return {
      account: null,
    };
  },
  methods: {
    async enableEthereum() {
      const provider = await detectEthereumProvider();

      // Determine the platform
      const userAgent = navigator.userAgent || navigator.vendor || window.opera;
      let isMobile = /android/i.test(userAgent) || /iPad|iPhone|iPod/.test(userAgent) && !window.MSStream;

      if (isMobile) {
        // Redirect to MetaMask mobile app
        const dappUrl = encodeURIComponent(window.location.href); // URL encode the current page
        const metamaskAppDeepLink = `https://metamask.app.link/dapp/${dappUrl}`;
        window.location.href = metamaskAppDeepLink;
      } else if (provider && provider === window.ethereum) {
        try {
          const accounts = await window.ethereum.request({
            method: "eth_requestAccounts",
          });
          this.account = accounts[0];
          console.log("Connected Account:", this.account);
        } catch (error) {
          console.error("Error connecting to MetaMask:", error);
        }
      } else {
        alert("MetaMaskをインストールしてください!");
      }
    },
  },
};
</script>

<style scoped>
.enableEthereumButton {
  width: 100%;
  padding: 10px;
  /* background-color: #007bff;  Match the color of the "Primary" button  */
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.enableEthereumButton:hover {
  /* background-color: #0056b3; Hover color similar to the "Primary" button */
}
</style>

