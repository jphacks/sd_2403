<template>
  <div class="table-container">
    <el-table :data="tableData" style="width: 100%">
      <el-table-column label="種類" prop="category" width="100" :header-cell-style="{ textAlign: 'center' }"></el-table-column>
      <el-table-column label="時間" prop="time" width="150"></el-table-column>
      <el-table-column label="都道府県" prop="unit_toponym" width="100"></el-table-column>
      <el-table-column label="地址详细" prop="unit_address" width="200"></el-table-column>
      <el-table-column label="会社名" prop="unit_name" width="120"></el-table-column>
      <el-table-column label="運輸" prop="transport_state" width="120"></el-table-column>
      <el-table-column label="保存" prop="fresh_state" width="120"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: [
        {
          category: "野菜",
          time: "2024-10-25 15:04:05",
          unit_toponym: "千葉",
          unit_address: "茂原市",
          unit_name: "公司A",
          transport_state: "生産",
          fresh_state: "常温",
        },
        {
          category: "野菜",
          time: "2024-10-25 18:34:57",
          unit_toponym: "千葉",
          unit_address: "茂原市",
          unit_name: "公司A",
          transport_state: "出庫",
          fresh_state: "常温",
        },
        {
          category: "野菜",
          time: "2024-10-25 20:29:45",
          unit_toponym: "千葉",
          unit_address: "市川市",
          unit_name: "会社B",
          transport_state: "入庫",
          fresh_state: "常温",
        },
        {
          category: "野菜",
          time: "2024-10-25 22:15:41",
          unit_toponym: "千葉",
          unit_address: "市川市",
          unit_name: "会社B",
          transport_state: "出庫",
          fresh_state: "常温",
        },
        {
          category: "野菜",
          time: "2024-10-26 12:26:39",
          unit_toponym: "宮城",
          unit_address: "仙台市青葉区",
          unit_name: "会社B",
          transport_state: "入庫",
          fresh_state: "常温",
        },
        {
          category: "野菜",
          time: "2024-10-26 16:34:48",
          unit_toponym: "宮城",
          unit_address: "仙台市青葉区",
          unit_name: "会社B",
          transport_state: "出庫",
          fresh_state: "常温",
        },
        {
          category: "野菜",
          time: "2024-10-26 18:07:20",
          unit_toponym: "宮城",
          unit_address: "仙台市青葉区",
          unit_name: "会社C",
          transport_state: "入庫",
          fresh_state: "冷蔵",
        },
      ],
    };
  },
};
</script>

<style scoped>
.table-container {
  width: 100%;
  overflow-x: auto; /* 允许在小屏幕上水平滚动 */
}

.el-table {
  width: 100%;
}

/* 响应式样式 */
@media (max-width: 600px) {
  .el-table {
    display: block; /* 使表格为块级元素 */
  }

  .el-table th,
  .el-table td {
    padding: 8px 5px; /* 减少小屏幕上的内边距 */
    font-size: 14px; /* 更小的字体大小 */
    display: block; /* 将单元格变为块级 */
    width: 100%; /* 让单元格全宽 */
  }

  .el-table th {
    position: relative;
    padding-top: 10px;
    padding-bottom: 10px;
    border-bottom: 1px solid #e4e7ed; /* 添加分隔线 */
    text-align: left; /* 让表头文本左对齐 */
  }

  .el-table td {
    border-bottom: 1px solid #e4e7ed; /* 添加分隔线 */
  }
}
</style>
