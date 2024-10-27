<template>
    <div class="login">
      <div class="content">
        <h2 v-if="showWelcomeMessage" class="welcome-message">ようこそ！</h2> <!-- 添加欢迎语，条件渲染 -->
        <router-view></router-view> <!-- 渲染子组件 -->
      </div>
      <div class="tab-bar">
        <el-button type="primary" @click="goTo('home')">ホーム</el-button>
        <el-button type="success" @click="goTo('history')">My履歴</el-button>
        <el-button type="warning" @click="goTo('scan')">スキャン</el-button>
        <el-button type="danger" @click="goTo('settings')">設置</el-button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'LoginPage',
    data() {
      return {
        showWelcomeMessage: true, // 控制欢迎语的显示
      };
    },
    methods: {
      goTo(section) {
        this.showWelcomeMessage = false; // 切换分区时隐藏欢迎语
        this.$router.push({ name: section });
      },
    },
  };
  </script>
  
  <style scoped>
  .login {
    display: flex;
    flex-direction: column;
    height: 100vh;
    max-width: 480px; /* 适应手机屏幕宽度 */
    margin: 0 auto;
  }
  
  .content {
    flex: 1;
    padding: 20px;
    overflow-y: auto; /* 确保内容可以滚动 */
  }
  
  .welcome-message {
    text-align: center;
    font-size: 18px;
    margin-bottom: 20px;
    color: #333;
  }
  
  .tab-bar {
    display: flex;
    justify-content: space-around;
    background-color: #f1f1f1;
    padding: 10px 0;
  }
  
  /* 设置按钮样式 */
  button {
    flex: 1;
    padding: 10px;
    border: none;
    color: white;
    font-weight: bold;
    font-size: 14px;
    cursor: pointer;
  }
  
  /* 为每个按钮分配不同颜色 */
  .home-button {
    background-color: #007bff; /* 蓝色 */
  }
  
  .history-button {
    background-color: #28a745; /* 绿色 */
  }
  
  .scan-button {
    background-color: #ffc107; /* 黄色 */
  }
  
  .settings-button {
    background-color: #dc3545; /* 红色 */
  }
  
  /* 添加悬停效果 */
  button:hover {
    opacity: 0.8;
  }
  </style>
  
