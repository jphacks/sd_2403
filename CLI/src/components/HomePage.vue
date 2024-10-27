<template>
    <div class="home-page">
      <h2 class="page-title">リスト作成</h2>
      <div class="button-container">
        <el-button type="primary" @click="goToInbound">入庫リスト</el-button>
        <el-button type="danger" @click="goToOutbound">出庫リスト</el-button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'HomePage',
    methods: {
      goToInbound() {
        this.$router.push({ name: 'inbound' }); // 导航到入库子页面
      },
      goToOutbound() {
        this.$router.push({ name: 'outbound' }); // 导航到出库子页面
      }
    }
  };
  </script>
  
  <style scoped>
  .home-page {
    display: flex;
    flex-direction: column;
    align-items: flex-start; /* 左对齐 */
    padding: 20px; /* 页面内边距 */
    height: 100vh; /* 使页面适应手机屏幕 */
    box-sizing: border-box; /* 包含padding在内的总宽度和高度计算 */
  }
  
  .page-title {
    margin: 0; /* 去掉默认的上下间距 */
    /* 添加额外的样式确保标题靠左 */
    text-align: left; /* 确保文本左对齐 */
  }
  
  .button-container {
    display: flex; /* 使用 flexbox */
    flex-direction: column; /* 设置为垂直排列 */
    align-items: center; /* 居中对齐按钮 */
    margin-top: 20px; /* 给按钮增加上边距 */
    width: 100%; /* 使按钮容器占满全宽 */
  }
  
  button {
    width: 120px; /* 设置按钮宽度 */
    padding: 10px;
    margin: 10px 0; /* 垂直间距 */
    border: none;
    color: white;
    font-size: 18px;
    font-weight: bold;
    border-radius: 8px;
    cursor: pointer;
  }
  
  .inbound-button {
    background-color: #007bff; /* 入库按钮颜色 */
  }
  
  .outbound-button {
    background-color: #dc3545; /* 出库按钮颜色 */
  }
  
  /* 添加悬停效果 */
  button:hover {
    opacity: 0.8;
  }
  </style>
