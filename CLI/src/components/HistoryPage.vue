<template>
    <div class="history-page">
      <h2 class="page-title">作業記録の照会</h2>
    </div>
  </template>
  
  <script>
  export default {
    name: 'HistoryPage',
  };
  </script>
  
  <style scoped>
  .history-page {
    padding: 20px; /* 页面内边距 */
    height: 100vh; /* 使页面适应手机屏幕 */
    display: flex;
    flex-direction: column;
    justify-content: flex-start; /* 向上对齐 */
  }
  
  .page-title {
    margin: 0; /* 去掉默认的上下间距 */
  }
  </style>
  