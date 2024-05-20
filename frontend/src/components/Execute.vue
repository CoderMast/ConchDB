<script setup lang="ts">
import { ref } from 'vue'
import { Execute } from "../../wailsjs/go/services/Connection";

import { useMessage } from 'naive-ui';

const message = useMessage();

// 执行 SQL
const sql = ref("select * from users;");
const sqlResult = ref();
const sqlResultHead = ref()
const sqlResultBody = ref()


function execute() {
  // 只要执行，就记录日志
  addSqlExecuteLog();
  // 实际请求
  Execute(sql.value).then(result => {
    if (result.substring(0, 5) != "Error") {
      console.log(result)
      // 将 JSON 字符串转换为二维数组
      const jsonResult = JSON.parse(result);
      console.log(jsonResult)
      // 将二维数组赋值给页面变量
      sqlResult.value = jsonResult;
      sqlResultHead.value = jsonResult[0];
      sqlResultBody.value = jsonResult.slice(1);


      // 执行成功提示
      message.success("执行成功！");
    } else {
      // 执行失败的提示信息
      message.error(result)
      console.log(result)
    }
  })
}

// SQL 语句执行记录
let executeList = ref(new Array<string>());

// 增加 SQL 执行记录
function addSqlExecuteLog() {
  executeList.value.push(sql.value)
}
</script>

<template>

  <n-row style="width: 80vw;margin: 0 auto;">
    <n-col :span="8" style="margin: auto">请输入你要执行的SQL语句：</n-col>
    <n-col :span="16" style="display: flex;margin: auto">
      <n-input type="text" v-model="sql" />
      <n-button type="primary" @click="execute">执行SQL</n-button>
    </n-col>
  </n-row>


  <!-- 表单渲染 -->
  <!-- <n-data-table :columns="sqlResultBody[0]" :data="sqlResultBody" style="width: 100%" /> -->

  <!-- 日志渲染 -->
  <!-- <n-card style="width: 100vw">
    <p style="text-align: center;" v-for="(item, index) in executeList" :key="index" class="text item">{{ item }}</p>
  </n-card> -->
</template>

<style scoped></style>