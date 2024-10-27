<template>
    <div class="register-container">
      <h2>新規登録</h2>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="unit_id">法人番号:</label>
          <input type="text" v-model="form.unit_id" required />
        </div>
        <div class="form-group">
          <label for="unit_name">会社名:</label>
          <input type="text" v-model="form.unit_name" required />
        </div>
        <div class="form-group">
          <label for="post_code">郵便番号:</label>
          <input type="text" v-model="form.post_code" required />
        </div>
        <div class="form-group">
          <label for="unit_address">会社住所:</label>
          <input type="text" v-model="form.unit_address" required />
        </div>
        <div class="form-group">
          <label for="unit_tel">電話番号:</label>
          <input type="tel" v-model="form.unit_tel" required />
        </div>
        <div class="form-group">
          <label for="unit_mail">メールアドレス:</label>
          <input type="email" v-model="form.unit_mail" required />
        </div>
        <div class="form-group">
          <label for="unit_type">サプライチェーン事業者の業種:</label>
          <select v-model="form.unit_type" required>
            <option value="" disabled>業種を選択してください</option>
            <option value="生産">生産</option>
            <option value="運輸">運輸</option>
            <option value="セール">セール</option>
          </select>
        </div>
  
        <!-- 按钮区域 -->
        <div class="button-group">
          <!-- <el-button type="primary" @click="openMetaMask">MetaMaskをインストール</el-button> -->
          <ConnectMetamask />
          <el-button type="success" @click="handleSubmit">送信</el-button>
          <el-button type="danger" @click="goToHelloPage">戻る</el-button>
        </div>
      </form>
    </div>
  </template>
  
  <script>
  import ConnectMetamask from './ConnectMetamask.vue';
  export default {
    components: {
        ConnectMetamask
    },
    data() {
      return {
        form: {
          unit_id: '',
          unit_name: '',
          post_code: '',
          unit_address: '',
          unit_tel: '',
          unit_mail: '',
          unit_type: '',
        },
      };
    },
    methods: {
      handleSubmit() {
        // 这里可以添加表单提交逻辑，比如发送到后端
        console.log('提交的表单数据：', this.form);
        alert('登録成功！');
        // 可以重置表单
        this.resetForm();
      },
      resetForm() {
        this.form = {
          unit_id: '',
          unit_name: '',
          post_code: '',
          unit_address: '',
          unit_tel: '',
          unit_mail: '',
          unit_type: '',
        };
      },
      openMetaMask() {
        window.open('https://metamask.io/ja/download/', '_blank');
      },
      goToHelloPage() {
        this.$router.push('/'); // 跳转到 HelloPage
      },
    },
  };
  </script>
  
  <style scoped>
  .register-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
    background-color: #f9f9f9;
  }
  
  h2 {
    text-align: center;
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  label {
    display: block;
    margin-bottom: 5px;
  }
  
  input,
  select {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .button-group {
    display: flex;
    flex-direction: column;
    gap: 10px; /* 按钮间距 */
  }
  
  .metamask-button,
  .submit-button,
  .back-button {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  .metamask-button {
    background-color: #f6851b; /* MetaMask橙色 */
    color: white;
  }
  
  .metamask-button:hover {
    background-color: #e2761b;
  }
  
  .submit-button {
    background-color: #007bff; /* 送信按钮颜色 */
    color: white;
  }
  
  .submit-button:hover {
    background-color: #0056b3;
  }
  
  .back-button {
    background-color: #dc3545; /* 返回按钮颜色 */
    color: white;
  }
  
  .back-button:hover {
    background-color: #c82333;
  }
  </style>
  
